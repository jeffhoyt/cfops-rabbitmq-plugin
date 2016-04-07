package plugin

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pivotalservices/cfbackup"
	"github.com/pivotalservices/cfops/plugin/cfopsplugin"
	"github.com/xchapter7x/lo"
)

const (
	productName               = "p-rabbitmq"
	jobName                   = "rabbitmq-server"
	serverAdminCredentialsKey = "server_admin_credentials"
	rabbitDefinitionsFile     = "rabbit-definitions.json"
)

// Backup - method to execute backup
func (plugin *RabbitMQPlugin) Backup() (err error) {
	lo.G.Debug("Starting RabbitMQ backup")
	//fmt.Printf("Installation Settings: %v\n", plugin.InstallationSettings)
	//fmt.Printf("Detected Rabbit API URL %s\n", GetAPIInformationFromInstallationSettings(plugin.InstallationSettings))

	definitionBytes, err := plugin.RabbitClient.GetServerDefinitions()

	reader := bytes.NewReader(definitionBytes)
	var writer io.WriteCloser
	writer, err = plugin.PivotalCF.NewArchiveWriter(rabbitDefinitionsFile)
	if err == nil {
		defer writer.Close()
		written, err := io.Copy(writer, reader)
		if err != nil {
			lo.G.Errorf("Failed to write backup file: %s", err)
		} else {
			lo.G.Debugf("Wrote %d bytes for backup file.", written)
		}
	}

	lo.G.Debug("Completed RabbitMQ backup")
	return
}

// Restore - method to execute Restore
func (plugin *RabbitMQPlugin) Restore() (err error) {
	lo.G.Debug("Starting RabbitMQ restore")

	var reader io.ReadCloser
	if reader, err = plugin.PivotalCF.NewArchiveReader(rabbitDefinitionsFile); err == nil {
		defer reader.Close()
		bytes, err := ioutil.ReadAll(reader)
		if err == nil {
			lo.G.Debugf("Read %d bytes from backup file.", len(bytes))
			err = plugin.RabbitClient.RestoreDefinitions(bytes)
			if err != nil {
				lo.G.Errorf("Failed to restore definitions to rabbit: %s\n", err.Error())
			}
		}
	}
	return
}

//GetMeta Get meta data of the plugin
func (plugin *RabbitMQPlugin) GetMeta() cfopsplugin.Meta {
	return plugin.Meta
}

//Setup -- Pass in the installation information
func (plugin *RabbitMQPlugin) Setup(pcf cfopsplugin.PivotalCF) (err error) {
	lo.G.Debug("Setting up plugin...")
	plugin.PivotalCF = pcf
	plugin.InstallationSettings = pcf.GetInstallationSettings()
	if plugin.RabbitClient == nil {
		clientData, _ := GetAPIInformationFromInstallationSettings(plugin.InstallationSettings)
		lo.G.Debugf("** Setting Rabbit Client %s\n", clientData.URL)
		plugin.RabbitClient = &clientData
	}
	return
}

// GetAPIInformationFromInstallationSettings sifts through the installation settings data to find the right
// API endpint for the rabbit API.
func GetAPIInformationFromInstallationSettings(installationSettings cfbackup.InstallationSettings) (clientData RabbitClientData, err error) {
	IPs, err := installationSettings.FindIPsByProductAndJob(productName, jobName)
	if err == nil {
		clientData.URL = fmt.Sprintf("http://%s:15672/api/", IPs[0])
		propMap, err := installationSettings.FindPropertyValues(productName, jobName, serverAdminCredentialsKey)
		if err == nil {
			clientData.Username = propMap["identity"]
			clientData.Password = propMap["password"]
		}
	}
	return
}

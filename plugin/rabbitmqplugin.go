package plugin

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pivotalservices/cfbackup"
	"github.com/pivotalservices/cfops/plugin/cfopsplugin"
	"github.com/xchapter7x/lo"
)

const (
	productName               = "p-rabbitmq"
	jobName                   = "rabbitmq-server"
	serverAdminCredentialsKey = "server_admin_credentials"
	rabbitUsersFile           = "rabbitmq-users.json"
)

// Backup - method to execute backup
func (plugin *RabbitMQPlugin) Backup() (err error) {
	lo.G.Debug("Starting RabbitMQ backup")
	//fmt.Printf("Installation Settings: %v\n", plugin.InstallationSettings)
	//fmt.Printf("Detected Rabbit API URL %s\n", GetAPIInformationFromInstallationSettings(plugin.InstallationSettings))

	userBytes, err := plugin.RabbitClient.GetUsersFile()

	reader := bytes.NewReader(userBytes)
	var writer io.WriteCloser
	if writer, err = plugin.PivotalCF.NewArchiveWriter(rabbitUsersFile); err == nil {
		defer writer.Close()
		_, err = io.Copy(writer, reader)
	}

	lo.G.Debug("Completed RabbitMQ backup")
	return
}

// Restore - method to execute Restore
func (plugin *RabbitMQPlugin) Restore() (err error) {
	lo.G.Debug("Starting RabbitMQ restore")
	return
}

//GetMeta Get meta data of the plugin
func (plugin *RabbitMQPlugin) GetMeta() cfopsplugin.Meta {
	return plugin.Meta
}

//Setup -- Pass in the installation information
func (plugin *RabbitMQPlugin) Setup(pcf cfopsplugin.PivotalCF) (err error) {
	plugin.PivotalCF = pcf
	plugin.InstallationSettings = pcf.GetInstallationSettings()
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

/*
if vmCredentials, err = s.InstallationSettings.FindVMCredentialsByProductAndJob(productName, jobName); err == nil
*/

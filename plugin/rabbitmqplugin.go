package plugin

import (
	"fmt"

	"github.com/pivotalservices/cfops/plugin/cfopsplugin"
)

// Backup - method to execute backup
func (plugin *RabbitMQPlugin) Backup() (err error) {
	fmt.Println("Doing a backup")

	fmt.Printf("HostDetails : %v\n", plugin.PivotalCF.GetHostDetails())
	fmt.Printf("Installation Settings: %v\n", plugin.PivotalCF.GetInstallationSettings())
	return
}

// Restore - method to execute Restore
func (plugin *RabbitMQPlugin) Restore() (err error) {
	fmt.Println("Doing a restore")
	return
}

//GetMeta Get meta data of the plugin
func (plugin *RabbitMQPlugin) GetMeta() cfopsplugin.Meta {
	return plugin.Meta
}

//Setup -- Pass in the installation information
func (plugin *RabbitMQPlugin) Setup(pcf cfopsplugin.PivotalCF) error {
	plugin.PivotalCF = pcf
	return nil
}

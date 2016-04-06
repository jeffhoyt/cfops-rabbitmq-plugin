package main

import (
	"fmt"

	"github.com/pivotalservices/cfops-rabbitmq-plugin/plugin"
	cfopsplugin "github.com/pivotalservices/cfops/plugin/cfopsplugin"
)

// WARNING - do not emit anything to STDOUT from within the main function
// or it will mess up the plugin detection.
func main() {
	meta := cfopsplugin.Meta{Name: "rabbitmq", Role: "backup-and-restore-rabbitmq"}
	rabbitplugin := plugin.NewRabbitMQPlugin(meta)
	augmentPlugin(rabbitplugin)
	cfopsplugin.Start(rabbitplugin)
}

func augmentPlugin(rabbitplugin *plugin.RabbitMQPlugin) {
	rabbitplugin.PivotalCF.GetInstallationSettings()
	clientData, err := plugin.GetAPIInformationFromInstallationSettings(rabbitplugin.InstallationSettings)
	rabbitplugin.RabbitClient = &clientData
	if err != nil {
		panic(fmt.Sprintf("Couldn't locate RabbitMQ server configuration in installation settings: %s\n", err.Error()))
	}
}

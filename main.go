package main

import (
	"github.com/pivotalservices/cfops-rabbitmq-plugin/plugin"
	cfopsplugin "github.com/pivotalservices/cfops/plugin/cfopsplugin"
)

// WARNING - do not emit anything to STDOUT from within the main function
// or it will mess up the plugin detection.
func main() {
	meta := cfopsplugin.Meta{Name: "rabbitmq", Role: "backup-and-restore-rabbitmq"}
	rabbitplugin := plugin.NewRabbitMQPlugin(meta)

	cfopsplugin.Start(rabbitplugin)
}

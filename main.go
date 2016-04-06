package main

import (
	"fmt"

	"github.com/pivotalservices/cfops-rabbitmq-plugin/plugin"
	cfopsplugin "github.com/pivotalservices/cfops/plugin/cfopsplugin"
)

func main() {
	meta := cfopsplugin.Meta{Name: "rabbitmq", Role: "backup-and-restore-rabbitmq"}
	plugin := plugin.NewRabbitMQPlugin(meta)
	fmt.Printf("%+v\n", meta)
	cfopsplugin.Start(plugin)
}

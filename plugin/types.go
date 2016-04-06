package plugin

import (
	"github.com/pivotalservices/cfbackup"
	"github.com/pivotalservices/cfops/plugin/cfopsplugin"
)

// RabbitMQPlugin anchor for methods
type RabbitMQPlugin struct {
	PivotalCF            cfopsplugin.PivotalCF
	InstallationSettings cfbackup.InstallationSettings
	Meta                 cfopsplugin.Meta
}

// NewRabbitMQPlugin creates a new RabbitMQ Plugin struct
func NewRabbitMQPlugin(meta cfopsplugin.Meta) (plugin *RabbitMQPlugin) {
	plugin = &RabbitMQPlugin{Meta: meta}
	return plugin
}

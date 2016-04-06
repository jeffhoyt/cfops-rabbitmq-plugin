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
	RabbitClient         RabbitClient
}

// Exchanges, Queues, Virtual Hosts, Policies and Users

// RabbitClient represents the set of functions we need to perform against
// rabbit to perform backup and restore operations
type RabbitClient interface {
	GetServerDefinitions() (definitionsFile []byte, err error)
	RestoreDefinitions(definitionsFile []byte) (err error)
}

// RabbitClientData represents the information required to connect to the Rabbit API
type RabbitClientData struct {
	URL      string
	Username string
	Password string
}

// NewRabbitMQPlugin creates a new RabbitMQ Plugin struct
func NewRabbitMQPlugin(meta cfopsplugin.Meta) (plugin *RabbitMQPlugin) {
	plugin = &RabbitMQPlugin{Meta: meta}
	return plugin
}

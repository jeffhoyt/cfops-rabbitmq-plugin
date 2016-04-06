package integrations_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/pivotalservices/cfbackup"
	plugin "github.com/pivotalservices/cfops-rabbitmq-plugin/plugin"
	"github.com/pivotalservices/cfops/plugin/cfopsplugin"
	"github.com/pivotalservices/cfops/tileregistry"
)

func TestIntegration(t *testing.T) {
	fmt.Println("placeholder")

}

func setupPlugin(installationSettingsPath string) (rabbitplugin *plugin.RabbitMQPlugin) {
	controlTmpDir, _ := ioutil.TempDir("", "unit-test")
	rabbitplugin = &plugin.RabbitMQPlugin{
		Meta: cfopsplugin.Meta{
			Name: "rabbitmq-tile",
		},
	}
	configParser := cfbackup.NewConfigurationParser(installationSettingsPath)
	pivotalCF := cfopsplugin.NewPivotalCF(configParser.InstallationSettings, tileregistry.TileSpec{
		ArchiveDirectory: controlTmpDir,
	})
	rabbitplugin.PivotalCF.GetInstallationSettings()
	clientData := &plugin.RabbitClientData{
		URL:      "fix",
		Username: "fix",
		Password: "fix",
	}

	rabbitplugin.RabbitClient = clientData
	rabbitplugin.Setup(pivotalCF)
	return
}

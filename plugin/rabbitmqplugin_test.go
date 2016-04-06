package plugin

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pivotalservices/cfbackup"
	fakes "github.com/pivotalservices/cfops-rabbitmq-plugin/fakes"
	"github.com/pivotalservices/cfops/plugin/cfopsplugin"
	"github.com/pivotalservices/cfops/tileregistry"
)

var (
	controlTmpDir string
)

func TestPluginSetupRetrievesRabbitEndpoint(t *testing.T) {
	plugin := setupPlugin("../fixtures/settings-1.6-aws.json")
	if plugin.Meta.Name != "rabbitmq-tile" {
		t.Errorf("Didn't properly configure plugin. Got %+v\n", plugin)
	}

	clientData, err := GetAPIInformationFromInstallationSettings(plugin.InstallationSettings)
	if err != nil {
		t.Errorf("Failed to get client data, %s", err.Error())
	}
	if clientData.URL != "http://10.0.16.17:15672/api/" {
		t.Errorf("Did not get appropriate management URL. Received %s\n", clientData.URL)
	}
	if clientData.Username != "srao" {
		t.Errorf("Did not get the right username, got %s\n", clientData.URL)
	}
	if clientData.Password != "srao" {
		t.Errorf("Did not get the right password, got %s\n", clientData.Password)
	}
	cleanTempDirectory(plugin.PivotalCF.GetHostDetails().ArchiveDirectory)
}

func TestPluginSetupFailsWhenRabbitIsNotInInstallationSettings(t *testing.T) {
	plugin := setupPlugin("../fixtures/settings-1.6-aws-no-rabbit.json")

	clientData, err := GetAPIInformationFromInstallationSettings(plugin.InstallationSettings)
	if err == nil {
		t.Errorf("Expected to get an error when rabbit isn't installed or configured, but didn't: %s\n", clientData.URL)
	}
	cleanTempDirectory(plugin.PivotalCF.GetHostDetails().ArchiveDirectory)
}

func TestBackupInvokesRabbitClient(t *testing.T) {
	//body := []byte("{\n  \"gridsize\": 19,\n  \"playerWhite\": \"bob\",\n  \"playerBlack\": \"alfred\"\n}")
	plugin := setupPlugin("../fixtures/settings-1.6-aws.json")
	if plugin.Meta.Name != "rabbitmq-tile" {
		t.Errorf("Didn't properly configure plugin. Got %+v\n", plugin)
	}
	err := plugin.Backup()
	if err != nil {
		t.Errorf("Failed to perform backup: %s\n", err.Error())
	}

	fakeclient := plugin.RabbitClient.(*fakes.FakeRabbitClient)
	if fakeclient.GetUsersCallCount != 1 {
		t.Errorf("Rabbit client didn't call get users 1x, called %d\n", fakeclient.GetUsersCallCount)
	}

	usersFile := fmt.Sprintf("%s/%s", plugin.PivotalCF.GetHostDetails().ArchiveDirectory, rabbitUsersFile)
	if isEmpty(usersFile) {
		t.Errorf("Should have backed up some data from the fake, but did not.")
	}

	cleanTempDirectory(plugin.PivotalCF.GetHostDetails().ArchiveDirectory)
}

func cleanTempDirectory(dir string) {
	os.RemoveAll(dir)
}

func setupPlugin(installationSettingsPath string) (plugin *RabbitMQPlugin) {
	controlTmpDir, _ := ioutil.TempDir("", "unit-test")
	plugin = &RabbitMQPlugin{
		Meta: cfopsplugin.Meta{
			Name: "rabbitmq-tile",
		},
		RabbitClient: fakes.NewFakeRabbitClient(),
	}
	configParser := cfbackup.NewConfigurationParser(installationSettingsPath)
	pivotalCF := cfopsplugin.NewPivotalCF(configParser.InstallationSettings, tileregistry.TileSpec{
		ArchiveDirectory: controlTmpDir,
	})
	plugin.Setup(pivotalCF)
	return
}

func isEmpty(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true
	}
	return false // Either not empty or error, suits both cases
}

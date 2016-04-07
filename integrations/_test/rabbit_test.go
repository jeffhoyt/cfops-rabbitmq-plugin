package integrations_test

import (
	"fmt"
	"os"
	"testing"

	plugin "github.com/pivotalservices/cfops-rabbitmq-plugin/plugin"
)

func TestIntegration(t *testing.T) {
	fmt.Println("placeholder")

	rabbitclient := setupClient()
	definitionsBytes, err := rabbitclient.GetServerDefinitions()

	if err != nil {
		t.Errorf("Failed to download server definitions: %s\n", err.Error())
		return
	}

	fmt.Printf("Downloaded %d bytes.\n", len(definitionsBytes))

	err = rabbitclient.RestoreDefinitions(definitionsBytes)
	if err != nil {
		t.Errorf("Failed to restore server definitions: %s\n", err.Error())
	}

}

func setupClient() (client *plugin.RabbitClientData) {

	rabbitHost := os.Getenv("RABBITMQ_PORT_15672_TCP_ADDR")
	rabbitAdminURL := fmt.Sprintf("http://%s:15672/api/", rabbitHost)
	client = &plugin.RabbitClientData{
		URL:      rabbitAdminURL,
		Username: "guest",
		Password: "guest",
	}
	fmt.Printf("Rabbit client set up for host %s\n", rabbitHost)
	return
}

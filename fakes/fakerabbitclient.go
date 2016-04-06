package fakes

// FakeRabbitClient used to mimic rabbit client to test plugin functionality
type FakeRabbitClient struct {
	GetDefinitionsCallCount  int
	PostDefinitionsCallCount int
}

// GetServerDefinitions fake for retrieving user file
func (client *FakeRabbitClient) GetServerDefinitions() (definitionsFile []byte, err error) {

	//definitionsFile := []byte("{\"drone_id\":\"drone666\", \"battery\": 72, \"uptime\": 6941, \"core_temp\": 21 }")
	definitionsFile = []byte("{\"rabbit_version\":\"3.5.7\",\"users\":[],\"vhosts\":[], \"permissions\":[],\"parameters\":[],\"queues\":[],\"exchanges\":[],\"bindings\":[]}")
	client.GetDefinitionsCallCount++
	return
}

// RestoreDefinitions fake for restoring definitions
func (client *FakeRabbitClient) RestoreDefinitions(definitionsFile []byte) (err error) {
	client.PostDefinitionsCallCount++
	return
}

// NewFakeRabbitClient creates a new fake client for testing
func NewFakeRabbitClient() (fake *FakeRabbitClient) {
	fake = &FakeRabbitClient{}
	return
}

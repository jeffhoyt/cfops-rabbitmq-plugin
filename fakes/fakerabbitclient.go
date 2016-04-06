package fakes

// FakeRabbitClient used to mimic rabbit client to test plugin functionality
type FakeRabbitClient struct {
	GetUsersCallCount     int
	RestoreUsersCallCount int
}

// GetUsersFile fake for retrieving user file
func (client *FakeRabbitClient) GetUsersFile() (userFile []byte, err error) {
	client.GetUsersCallCount++
	return
}

// RestoreUsersFile fake for restoring users file
func (client *FakeRabbitClient) RestoreUsersFile(userFile []byte) (err error) {
	client.RestoreUsersCallCount++
	return
}

// NewFakeRabbitClient creates a new fake client for testing
func NewFakeRabbitClient() (fake *FakeRabbitClient) {
	fake = &FakeRabbitClient{}
	return
}

package plugin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/xchapter7x/lo"
)

// GetServerDefinitions retrieves definitions from the server
func (cd *RabbitClientData) GetServerDefinitions() (definitionsFile []byte, err error) {
	definitionsFile, err = doGet(fmt.Sprintf("%sdefinitions", cd.URL), cd.Username, cd.Password)
	return
}

// RestoreDefinitions updates definitions on the server
func (cd *RabbitClientData) RestoreDefinitions(definitionsFile []byte) (err error) {
	doPost(fmt.Sprintf("%sdefinitions", cd.URL), cd.Username, cd.Password, definitionsFile)
	return
}

func doPost(URL string, username string, password string, body []byte) (err error) {
	lo.G.Debugf("About to POST %s\n", URL)
	httpclient := &http.Client{}

	req, _ := http.NewRequest("POST", URL, bytes.NewReader(body))
	req.SetBasicAuth(username, password)

	resp, err := httpclient.Do(req)

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Did not get HTTP 200 from server, got %d.", resp.StatusCode)
		return
	}

	if err != nil {
		lo.G.Error("Errored when posting request to the server")
		return
	}

	return
}

func doGet(URL string, username string, password string) (results []byte, err error) {
	lo.G.Debugf("About to GET %s\n", URL)
	httpclient := &http.Client{}

	req, _ := http.NewRequest("GET", URL, nil)
	req.SetBasicAuth(username, password)

	resp, err := httpclient.Do(req)

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Did not get HTTP 200 from server, got %d.", resp.StatusCode)
		return
	}

	if err != nil {
		fmt.Println("Errored when getting request to the server")
		return
	}

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed when reading response body from server.")
	}
	results = payload
	return
}

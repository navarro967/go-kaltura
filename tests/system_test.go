package kaltura

import (
	"testing"

	kaltura "github.com/kaltura/kaltura-client-go"
)

func TestPing(t *testing.T) {
	var serviceUrl string = kaltura.DefaultServiceUrl
	//create client
	client := kaltura.NewClient(&serviceUrl)

	//call ping
	pingResponse, _, err := client.System.Ping()
	if err != nil {
		t.Errorf("Error calling ping: %v", err)
	} else if pingResponse != true {
		t.Errorf("Ping response is not true: %v", pingResponse)
	} else {
		t.Logf("Ping response is true: %v", pingResponse)
	}
}

func TestPingDatabase(t *testing.T) {
	var serviceUrl string = kaltura.DefaultServiceUrl
	//create client
	client := kaltura.NewClient(&serviceUrl)

	//call PingDatabase
	pingResponse, _, err := client.System.PingDatabase()
	if err != nil {
		t.Errorf("Error calling pingDatabase: %v", err)
	} else if pingResponse != true {
		t.Errorf("Ping response is not true: %v", pingResponse)
	} else {
		t.Logf("PingDatabase response is true: %v", pingResponse)
	}
}

func TestGetVersion(t *testing.T) {
	var serviceUrl string = kaltura.DefaultServiceUrl
	//create client
	client := kaltura.NewClient(&serviceUrl)

	//call GetVersion
	versionResponse, _, err := client.System.GetVersion()
	if err != nil {
		t.Errorf("Error calling getVersion: %v", err)
	}
	t.Logf("Version response is: %v", versionResponse)
}

func TestGetTime(t *testing.T) {
	var serviceUrl string = kaltura.DefaultServiceUrl
	//create client
	client := kaltura.NewClient(&serviceUrl)

	//call GetTime
	timeResponse, _, err := client.System.GetTime()
	if err != nil {
		t.Errorf("Error calling getTime: %v", err)
	}
	t.Logf("Time response is: %v", timeResponse)
}

func TestGetHealthCheck(t *testing.T) {
	var serviceUrl string = kaltura.DefaultServiceUrl
	//create client
	client := kaltura.NewClient(&serviceUrl)

	//call GetHealthCheck
	healthCheckResponse, _, err := client.System.GetHealthCheck()
	if err != nil {
		t.Errorf("Error calling getHealthCheck: %v", err)
	}
	t.Logf("HealthCheck response is: %s", healthCheckResponse)
}

package kaltura

// import (
// 	"testing"

// 	"github.com/kaltura/kaltura-client-go"
// )

// //test client with ping function

// func TestRequest(t *testing.T) {
// 	var serviceUrl string = kaltura.DefaultServiceUrl
// 	//create client
// 	client, err := kaltura.NewClient(&serviceUrl)
// 	if err != nil {
// 		t.Errorf("Error creating client: %v", err)
// 	}

// 	url, err := kaltura.BuildQuery("system", "ping", nil)
// 	if err != nil {
// 		t.Errorf("Error building query: %v", err)
// 	}

// 	req, err := client.NewRequest("GET", url, nil)

// 	if err != nil {
// 		t.Errorf("Error creating request: %v", err)
// 	}

// 	t.Logf("Request: %v", req)

// 	var result = false

// 	res, err := client.Do(req, &result)
// 	if err != nil {
// 		t.Errorf("Error calling ping: %v", err)
// 	}

// 	t.Logf("\nResponse: %v\n %v", result, res.Request.URL)
// }

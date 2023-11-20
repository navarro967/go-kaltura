package kaltura

import (
	"testing"

	kaltura "github.com/kaltura/kaltura-client-go"
	"github.com/kaltura/kaltura-client-go/session"
)

func TestSesionGen(t *testing.T) {
	//var serviceUrl string = "https://api.nvq2.ovp.kaltura.com/api_v3/"
	//create client
	// client, err := kaltura.NewClient(&serviceUrl)
	// if err != nil {
	// 	t.Errorf("Error creating client: %v", err)
	// }

	//call Get
	ks := session.NewSession("",
		"admin",
		kaltura.SESSION_TYPE_ADMIN,
		0,
		86400,
		"disableentitlement",
	).GenerateSession()

	t.Logf("Start request is: %v", ks.Ks)
}

func TestSesionGenV2(t *testing.T) {
	//var serviceUrl string = "https://api.nvq2.ovp.kaltura.com/api_v3/"
	//create client
	// client, err := kaltura.NewClient(&serviceUrl)
	// if err != nil {
	// 	t.Errorf("Error creating client: %v", err)
	// }

	ks := session.NewSession("",
		"admin",
		kaltura.SESSION_TYPE_ADMIN,
		0,
		86400,
		"disableentitlement",
	).GenerateSessionV2()

	t.Logf("Start request is: %v", ks)
}

package kaltura

import (
	"encoding/json"
	"testing"

	"github.com/kaltura/kaltura-client-go"
	"github.com/kaltura/kaltura-client-go/session"
)

func TestUiConfGet(t *testing.T) {
	ks := session.NewSession("",
		"admin",
		kaltura.SESSION_TYPE_ADMIN,
		0,
		86400,
		"disableentitlement",
	).GenerateSessionV2()
	var serviceUrl string = "https://api.nvq2.ovp.kaltura.com/api_v3/"
	//create client
	client := kaltura.NewClient(&serviceUrl).WithKS(ks.Ks)

	uiConfResponse, _, err := client.UiConf.Get(23467818)
	if err != nil {
		t.Errorf("Error calling Get: %v", err)
	}
	//print each item in struct
	json, err := json.MarshalIndent(uiConfResponse, "", "  ")
	if err != nil {
		t.Errorf("Error marshalling response: %v", err)
	}
	t.Logf("Get response is: %v", string(json))
}

func TestUiConfList(t *testing.T) {
	ks := session.NewSession("",
		"admin",
		kaltura.SESSION_TYPE_ADMIN,
		0,
		86400,
		"disableentitlement",
	).GenerateSessionV2()
	var serviceUrl string = "https://api.nvq2.ovp.kaltura.com/api_v3/"
	//create client
	client := kaltura.NewClient(&serviceUrl).WithKS(ks.Ks)

	uiConfFilter := &kaltura.UiConfFilter{
		KalturaFilter: kaltura.KalturaFilter{
			Pager: kaltura.KalturaFilterPager{
				PageSize:  10,
				PageIndex: 1,
			},
		},
		Filter: kaltura.UiConfBaseFilter{
			NameLike: "KEA",
			KalturaBaseFilter: kaltura.KalturaBaseFilter{
				OrderBy: kaltura.CREATED_AT_DESC,
			},
		},
	}

	uiConfResponse, _, err := client.UiConf.List(uiConfFilter)
	if err != nil {
		t.Errorf("Error calling Get: %v", err)
	}

	for _, uiconf := range uiConfResponse.Objects {
		t.Logf("UiConf: %v:%v:%v", uiconf.Id, uiconf.Name, uiconf.Tags)
	}
	t.Logf("Total Count: %v", uiConfResponse.TotalCount)
}

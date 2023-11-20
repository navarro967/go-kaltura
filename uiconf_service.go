package kaltura

type UiConfService service

// ObjTypes
const (
	UICONF_OBJ_TYPE_PLAYER                    int8 = 1
	UICONF_OBJ_TYPE_CONTRIBUTION_WIZARD       int8 = 2
	UICONF_OBJ_TYPE_SIMPLE_EDITOR             int8 = 3
	UICONF_OBJ_TYPE_ADVANCED_EDITOR           int8 = 4
	UICONF_OBJ_TYPE_PLAYLIST                  int8 = 5
	UICONF_OBJ_TYPE_APP_STUDIO                int8 = 6
	UICONF_OBJ_TYPE_KRECORD                   int8 = 7
	UICONF_OBJ_TYPE_PLAYER_V3                 int8 = 8
	UICONF_OBJ_TYPE_KMC_ACCOUNT               int8 = 9
	UICONF_OBJ_TYPE_KMC_ANALYTICS             int8 = 10
	UICONF_OBJ_TYPE_KMC_CONTENT               int8 = 11
	UICONF_OBJ_TYPE_KMC_DASHBOARD             int8 = 12
	UICONF_OBJ_TYPE_KMC_LOGIN                 int8 = 13
	UICONF_OBJ_TYPE_PLAYER_SL                 int8 = 14
	UICONF_OBJ_TYPE_CLIENTSIDE_ENCODER        int8 = 15
	UICONF_OBJ_TYPE_KMC_GENERAL               int8 = 16
	UICONF_OBJ_TYPE_KMC_ROLES_AND_PERMISSIONS int8 = 17
	UICONF_OBJ_TYPE_CLIPPER                   int8 = 18
	UICONF_OBJ_TYPE_KSR                       int8 = 19
	UICONF_OBJ_TYPE_KUPLOAD                   int8 = 20
	UICONF_OBJ_TYPE_WEBCASTING                int8 = 21
)

// CreationModes
const (
	UICONF_CREATION_MODE_WIZARD   int8 = 2
	UICONF_CREATION_MODE_ADVANCED int8 = 3
	UICONF_CREATION_MODE_SYSTEM   int8 = 4
)

type UiConf struct {
	Id               int64  `json:"id,omitempty"`
	ConfFile         string `json:"confFile,omitempty"`
	ConfFileFeatures string `json:"confFileFeatures,omitempty"`
	ConfFileParams   string `json:"confFileParams,omitempty"`
	ConfVars         string `json:"confVars,omitempty"`
	Config           string `json:"config,omitempty"`
	CreatedAt        int64  `json:"createdAt,omitempty"`
	CreationMode     int64  `json:"creationMode,omitempty"`
	Description      string `json:"description,omitempty"`
	Height           string `json:"height,omitempty"`
	Html5Url         string `json:"html5Url,omitempty"`
	HtmlParams       string `json:"htmlParams,omitempty"`
	Name             string `json:"name,omitempty"`
	ObjType          int8   `json:"objType,omitempty"`
	ObjTypeAsString  string `json:"objTypeAsString,omitempty"`
	PartnerId        int64  `json:"partnerId,omitempty"`
	PartnerTags      string `json:"partnerTags,omitempty"`
	SwfUrl           string `json:"swfUrl,omitempty"`
	SwfUrlVersion    string `json:"swfUrlVersion,omitempty"`
	Tags             string `json:"tags,omitempty"`
	UpdatedAt        int64  `json:"updatedAt,omitempty"`
	UseCdn           int    `json:"useCdn,omitempty"`
	Version          string `json:"version,omitempty"`
	Width            string `json:"width,omitempty"`
}

type UiConfListResponse struct {
	ClientResponse
	ObjectType string    `json:"objectType,omitempty"`
	Objects    []*UiConf `json:"objects,omitempty"`
}

type UiConfBaseFilter struct {
	KalturaBaseFilter
	CreatedAtGreaterThanOrEqual int64  `json:"createdAtGreaterThanOrEqual,omitempty"`
	CreatedAtLessThanOrEqual    int64  `json:"createdAtLessThanOrEqual,omitempty"`
	CreationModeEqual           int64  `json:"creationModeEqual,omitempty"`
	CreationModeIn              string `json:"creationModeIn,omitempty"`
	IdEqual                     int64  `json:"idEqual,omitempty"`
	IdIn                        string `json:"idIn,omitempty"`
	NameLike                    string `json:"nameLike,omitempty"`
	ObjTypeEqual                int8   `json:"objTypeEqual,omitempty"`
	ObjTypeIn                   string `json:"objTypeIn,omitempty"`
	ObjectType                  string `json:"objectType,omitempty"`
	PartnerIdEqual              int64  `json:"partnerIdEqual,omitempty"`
	PartnerIdIn                 string `json:"partnerIdIn,omitempty"`
	PartnerTagsMultiLikeAnd     string `json:"partnerTagsMultiLikeAnd,omitempty"`
	PartnerTagsMultiLikeOr      string `json:"partnerTagsMultiLikeOr,omitempty"`
	TagsMultiLikeAnd            string `json:"tagsMultiLikeAnd,omitempty"`
	TagsMultiLikeOr             string `json:"tagsMultiLikeOr,omitempty"`
	UpdatedAtGreaterThanOrEqual int64  `json:"updatedAtGreaterThanOrEqual,omitempty"`
	UpdatedAtLessThanOrEqual    int64  `json:"updatedAtLessThanOrEqual,omitempty"`
	VersionEqual                string `json:"versionEqual,omitempty"`
	VersionMultiLikeAnd         string `json:"versionMultiLikeAnd,omitempty"`
	VersionMultiLikeOr          string `json:"versionMultiLikeOr,omitempty"`
}

type UiConfFilter struct {
	KalturaFilter
	Filter     UiConfBaseFilter `json:"filter"`
	ObjectType string           `json:"objectType,omitempty"`
}

func NewUiConfFilter() *UiConfFilter {
	ui := UiConfFilter{}
	return &ui
}

func (s *UiConfService) Get(id int64) (*UiConf, *ClientResponse, error) {
	uiconf := &UiConf{Id: id}
	u, err := buildQuery("uiconf", "get", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, uiconf)
	if err != nil {
		return nil, nil, err
	}

	var result = &UiConf{}
	res, err := s.client.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, res, nil
}

func (s *UiConfService) Add(uiconf *UiConf) (*UiConf, *ClientResponse, error) {
	u, err := buildQuery("uiconf", "add", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, uiconf)
	if err != nil {
		return nil, nil, err
	}

	var result = &UiConf{}
	res, err := s.client.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, res, nil
}

func (s *UiConfService) Update(uiconf *UiConf) (*UiConf, *ClientResponse, error) {
	u, err := buildQuery("uiconf", "update", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, uiconf)
	if err != nil {
		return nil, nil, err
	}

	var result = &UiConf{}
	res, err := s.client.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, res, nil
}

func (s *UiConfService) Delete(id int64) (*ClientResponse, error) {
	uiconf := &UiConf{Id: id}
	u, err := buildQuery("uiconf", "delete", nil)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("POST", u, uiconf)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *UiConfService) List(filter *UiConfFilter) (*UiConfListResponse, *ClientResponse, error) {
	u, err := buildQuery("uiconf", "list", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, filter)
	if err != nil {
		return nil, nil, err
	}

	var result = &UiConfListResponse{}
	res, err := s.client.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, res, nil
}

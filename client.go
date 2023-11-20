package kaltura

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	DefaultServiceUrl string = "https://www.kaltura.com/api_v3/"
	ClientTag         string = "kaltura-client-go"
)

type service struct {
	client *Client
}

type Client struct {
	ServiceURL     string
	HTTPClient     *http.Client
	KalturaSession *Session
	ClientTag      string

	common service

	System  *SystemService
	Session *SessionService
	UiConf  *UiConfService
}

type ClientRequest struct {
	Service   string "json:\"service\""
	Action    string "json:\"action\""
	Params    interface{}
	ClientTag string "json:\"clientTag\""
	Format    int    "json:\"format\""
}

type ClientResponse struct {
	*http.Response

	NextPage  int
	PrevPage  int
	FirstPage int
	LastPage  int

	CurrPage   int
	TotalCount int
}

//Define Error

type KalturaApiException struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Args    string `json:"args"`
}

func (e KalturaApiException) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("%v: %v {%v}", e.Code, e.Message, e.Args)
	}
	return e.Error()
}

func NewResponse(r *http.Response) *ClientResponse {
	response := &ClientResponse{Response: r}
	return response
}

func NewClient(serviceUrl *string) *Client {
	c := Client{
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
		ClientTag: ClientTag,
	}
	strServiceUrl := ""
	if serviceUrl == nil {
		strServiceUrl = DefaultServiceUrl
	} else {
		strServiceUrl = *serviceUrl
	}
	c.ServiceURL = strServiceUrl
	c.initalize()

	return &c
}

func (c *Client) initalize() {
	url, err := url.Parse(c.ServiceURL)
	if err != nil {
		log.Fatal(err)
	}
	c.ServiceURL = url.String()
	c.common.client = c
	c.System = (*SystemService)(&c.common)
	c.Session = (*SessionService)(&c.common)
	c.UiConf = (*UiConfService)(&c.common)
}

func (c *Client) SetKs(ks string) {
	c.KalturaSession = &Session{
		Ks: ks,
	}
}

func (c *Client) WithKS(ks string) *Client {
	c.SetKs(ks)
	return c
}

func buildQuery(service string, action string, params interface{}) (string, error) {
	var base string = "service/" + service + "/action/" + action

	value := reflect.ValueOf(params)
	if value.Kind() == reflect.Ptr && value.IsNil() {
		return base, nil
	}

	url, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	queryString, err := query.Values(params)
	if err != nil {
		return "", err
	}
	queryString.Add("format", "1")

	url.RawQuery = queryString.Encode()
	return url.String(), nil
}

func (c *Client) NewRequest(method string, path string, body interface{}) (*http.Request, error) {
	url := c.ServiceURL + path

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.ClientTag)
	return req, nil
}

func (c *Client) RawDo(req *http.Request) (*ClientResponse, error) {
	if c.KalturaSession != nil {
		q := req.URL.Query()
		q.Add("ks", c.KalturaSession.Ks)
		req.URL.RawQuery = q.Encode()
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	response := NewResponse(res)

	if res.StatusCode != http.StatusOK {
		return response, nil
	}

	return response, nil
}

func checkResponse(buffer []byte) error {
	var result KalturaApiException
	err := json.Unmarshal(buffer, &result)
	if err != nil {
		return result
	}

	return nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*ClientResponse, error) {
	res, err := c.RawDo(req)
	if err != nil {
		return res, err
	}

	defer res.Body.Close()

	//copy response to to string
	buf := new(bytes.Buffer)

	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return nil, err
	}

	err = checkResponse(buf.Bytes())
	if err != nil {
		return nil, err
	}

	json.Unmarshal(buf.Bytes(), &v)

	return res, nil
}

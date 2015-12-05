package twiliogo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const ROOT = "https://api.twilio.com"
const VERSION = "2010-04-01"
const IP_MESSAGING_ROOT = "https://ip-messaging.twilio.com"
const IP_MESSAGING_VERSION = "v1"
const IP_MESSAGING_ROOT_URL = IP_MESSAGING_ROOT + "/" + IP_MESSAGING_VERSION

type Client interface {
	AccountSid() string
	AuthToken() string
	RootUrl() string
	get(url.Values, string) ([]byte, error)
	post(url.Values, string) ([]byte, error)
	getIP(url.Values, string) ([]byte, error)
	postIP(url.Values, string) ([]byte, error)
	deleteIP(string) error
}

type TwilioClient struct {
	accountSid string
	authToken  string
	rootUrl    string
}

func NewClient(accountSid, authToken string) *TwilioClient {
	rootUrl := "/" + VERSION + "/Accounts/" + accountSid
	return &TwilioClient{accountSid, authToken, rootUrl}
}

func (client *TwilioClient) post(values url.Values, uri string) ([]byte, error) {
	if !strings.HasPrefix(uri, "http") {
		uri = ROOT + uri
	}
	req, err := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.AccountSid(), client.AuthToken())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return body, Error{"Server Error"}
		} else {
			twilioError := new(TwilioError)
			json.Unmarshal(body, twilioError)
			return body, twilioError
		}
	}

	return body, err
}

func (client *TwilioClient) postIP(values url.Values, uri string) ([]byte, error) {
	if !strings.HasPrefix(uri, "http") {
		uri = IP_MESSAGING_ROOT_URL + uri
	}
	req, err := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.AccountSid(), client.AuthToken())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return body, Error{"Server Error"}
		} else {
			twilioError := new(TwilioError)
			json.Unmarshal(body, twilioError)
			return body, twilioError
		}
	}

	return body, err
}

func (client *TwilioClient) get(queryParams url.Values, uri string) ([]byte, error) {
	if !strings.HasPrefix(uri, "http") {
		uri = ROOT + uri
	}
	var params *strings.Reader

	if queryParams == nil {
		queryParams = url.Values{}
	}

	params = strings.NewReader(queryParams.Encode())
	req, err := http.NewRequest("GET", uri, params)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.AccountSid(), client.AuthToken())
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return body, Error{"Server Error"}
		} else {
			twilioError := new(TwilioError)
			json.Unmarshal(body, twilioError)
			return body, twilioError
		}
	}

	return body, err
}

func (client *TwilioClient) getIP(queryParams url.Values, uri string) ([]byte, error) {
	if !strings.HasPrefix(uri, "http") {
		uri = IP_MESSAGING_ROOT_URL + uri
	}
	var params *strings.Reader

	if queryParams == nil {
		queryParams = url.Values{}
	}

	params = strings.NewReader(queryParams.Encode())
	req, err := http.NewRequest("GET", uri, params)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.AccountSid(), client.AuthToken())
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return body, err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		if res.StatusCode == 500 {
			return body, Error{"Server Error"}
		} else {
			twilioError := new(TwilioError)
			json.Unmarshal(body, twilioError)
			return body, twilioError
		}
	}

	return body, err
}

func (client *TwilioClient) deleteIP(uri string) error {
	if !strings.HasPrefix(uri, "http") {
		uri = IP_MESSAGING_ROOT_URL + uri
	}
	req, err := http.NewRequest("DELETE", uri, nil)

	if err != nil {
		return err
	}

	req.SetBasicAuth(client.AccountSid(), client.AuthToken())
	httpClient := &http.Client{}

	res, err := httpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 204 {
		return fmt.Errorf("Non-204 returned from server for DELETE: %d", res.StatusCode)
	}

	return nil
}

func (client *TwilioClient) AccountSid() string {
	return client.accountSid
}

func (client *TwilioClient) AuthToken() string {
	return client.authToken
}

func (client *TwilioClient) RootUrl() string {
	return client.rootUrl
}

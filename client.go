package twilio

import (
  "net/http"
  "strings"
  "net/url"
  "io/ioutil"
  "encoding/json"
)

const ROOT = "https://api.twilio.com"
const VERSION = "2010-04-01"

type Client struct {
  AccountSid string
  AuthToken string
  RootUrl string
}

func newClient(accountSid, authToken string) *Client {
  rootUrl := "/" + VERSION + "/Accounts/" + accountSid
  return &Client{accountSid, authToken, rootUrl}
}

func (client *Client) post(formValues url.Values, uri string) (*http.Response, error) {
  req, err := http.NewRequest("POST", ROOT + client.RootUrl + uri, strings.NewReader(formValues.Encode()))

  if err != nil {
    return nil, err
  }

  req.SetBasicAuth(client.AccountSid, client.AuthToken)
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  httpClient := &http.Client{}
  return httpClient.Do(req)
}

func (client *Client) get(queryParams url.Values, uri string) ([]byte, error) {
  var params *strings.Reader

  if queryParams == nil {
    queryParams = url.Values{}
  }

  params = strings.NewReader(queryParams.Encode())
  req, err := http.NewRequest("GET", ROOT + uri, params)

  if err != nil {
    return nil, err
  }

  req.SetBasicAuth(client.AccountSid, client.AuthToken)
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

  return body, err
}

func (client *Client) GetMessageList() (*MessageList, error) {
  var messageList *MessageList

  body, err := client.get(nil, client.RootUrl + "/SMS/Messages.json")

  if err != nil {
    return messageList, err
  }

  messageList = new(MessageList)
  messageList.Client = client
  err = json.Unmarshal(body, messageList)

  return messageList, err
}

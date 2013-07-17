package twilio

import (
  "net/http"
  "strings"
  "net/url"
  "io/ioutil"
  "encoding/json"
)

const VERSION = "2010-04-01"

type Client struct {
  AccountSid string
  AuthToken string
  RootUrl string
}

func newClient(accountSid, authToken string) *Client {
  rootUrl := "https://api.twilio.com/" + VERSION + "/Accounts/" + accountSid
  return &Client{accountSid, authToken, rootUrl}
}

func (client *Client) post(formValues url.Values, uri string) (*http.Response, error) {
  req, err := http.NewRequest("POST", client.RootUrl + uri, strings.NewReader(formValues.Encode()))

  if err != nil {
    return nil, err
  }

  req.SetBasicAuth(client.AccountSid, client.AuthToken)
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  httpClient := &http.Client{}
  return httpClient.Do(req)
}

func (client *Client) get(queryParams url.Values, uri string) (*http.Response, error) {
  var params *strings.Reader

  if queryParams == nil {
    queryParams = url.Values{}
  }

  params = strings.NewReader(queryParams.Encode())
  req, err := http.NewRequest("GET", client.RootUrl + uri, params)

  if err != nil {
    return nil, err
  }

  req.SetBasicAuth(client.AccountSid, client.AuthToken)
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  httpClient := &http.Client{}
  return httpClient.Do(req)
}

func (client *Client) GetMessageList() (*MessageList, error) {
  var messageList *MessageList

  res, err:= client.get(nil, "/SMS/Messages.json")

  if err != nil {
    return messageList, err
  }

  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)

  if err != nil {
    return messageList, err
  }

  messageList = new(MessageList)
  err = json.Unmarshal(body, messageList)

  return messageList, err
}

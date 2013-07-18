package twilio

import (
  "encoding/json"
)

type MessageList struct {
  Client *Client
  Start int `json:"start"`
  Total int `json:"total"`
  NumPages int `json:"num_pages"`
  Page int `json:"page"`
  PageSize int `json:"page_size"`
  End int `json:"end"`
  Uri string `json:"uri"`
  FirstPageUri string `json:"first_page_uri"`
  LastPageUri string `json:"last_page_uri"`
  NextPageUri string `json:"next_page_uri"`
  PreviousPageUri string `json"previous_page_uri"`
  Messages []Message `json:"sms_messages"`
}

func (m *MessageList) list() []Message {
  return m.Messages
}

func (currentMessageList *MessageList) NextPage() (*MessageList, error) {
  if currentMessageList.NextPageUri == "" {
    return nil, Error {"No next page"}
  }

  return currentMessageList.getPage(currentMessageList.NextPageUri)
}

func (currentMessageList *MessageList) PreviousPage() (*MessageList, error) {
  if currentMessageList.PreviousPageUri == "" {
    return nil, Error {"No previous page"}
  }

  return currentMessageList.getPage(currentMessageList.NextPageUri)
}

func (currentMessageList *MessageList) FirstPage() (*MessageList, error) {
  if currentMessageList.FirstPageUri == "" {
    return nil, Error {"No first page"}
  }

  return currentMessageList.getPage(currentMessageList.FirstPageUri)
}

func (currentMessageList *MessageList) LastPage() (*MessageList, error) {
  if currentMessageList.FirstPageUri == "" {
    return nil, Error {"No last page"}
  }

  return currentMessageList.getPage(currentMessageList.LastPageUri)
}

func (currentMessageList *MessageList) getPage(uri string) (*MessageList, error) {
  var messageList *MessageList

  client := currentMessageList.Client

  body, err := client.get(nil, uri)

  if err != nil {
    return messageList, err
  }

  messageList = new(MessageList)
  messageList.Client = client
  err = json.Unmarshal(body, messageList)

  return messageList, err
}

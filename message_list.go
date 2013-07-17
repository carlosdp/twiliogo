package twilio

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

func (m MessageList) list() []Message {
  return m.Messages
}

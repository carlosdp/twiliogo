package twilio

type Message struct {
  Sid string `json:"sid"`
  DateCreated string `json:"date_created"`
  DateUpdated string `json:"date_updated"`
  DateSent string `json:"date_sent"`
  AccountSid string `json:"account_sid"`
  From string `json:"from"`
  To string `json:"to"`
  Body string `json:"body"`
  Status string `json:"status"`
  Direction string `json:"direction"`
  Price string `json:"price"`
  PriceUnit string `json:"price_unit"`
  ApiVersion string `json:"api_version"`
  uri string `json:"uri"`
}

package twilio

type Optional interface {
  GetParam() (string, string)
}

type Callback string

func (callback Callback) GetParam() (string, string) {
  return "Url", string(callback)
}

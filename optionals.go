package twiliogo

type Optional interface {
  GetParam() (string, string)
}

type Callback string

func (callback Callback) GetParam() (string, string) {
  return "Url", string(callback)
}

type ApplicationSid string

func (applicationSid ApplicationSid) GetParam() (string, string) {
  return "ApplicationSid", string(applicationSid)
}

type Method string

func (method Method) GetParams() (string, string) {
  return "Method", string(method)
}

type FallbackUrl string

func (fallbackUrl FallbackUrl) GetParams() (string, string) {
  return "FallbackUrl", string(fallbackUrl)
}

type FallbackMethod string

func (fallbackMethod FallbackMethod) GetParams() (string, string) {
  return "FallbackMethod", string(fallbackMethod)
}

type StatusCallback string

func (statusCallback StatusCallback) GetParams() (string, string) {
  return "StatusCallback", string(statusCallback)
}

type StatusCallbackMethod string

func (statusCallbackMethod StatusCallbackMethod) GetParams() (string, string) {
  return "StatusCallbackMethod", string(statusCallbackMethod)
}

type SendDigits string

func (sendDigits SendDigits) GetParams() (string, string) {
  return "SendDigits", string(sendDigits)
}

type IfMachine string

func (ifMachine IfMachine) GetParams() (string, string) {
  return "IfMachine", string(ifMachine)
}

type Timeout string

func (timeout Timeout) GetParams() (string, string) {
  return "Timeout", string(timeout)
}

type Record string

func (record Record) GetParams() (string, string) {
  return "Record", string(record)
}

type To string

func (to To) GetParams() (string, string) {
  return "To", string(to)
}

type From string

func (from From) GetParams() (string, string) {
  return "From", string(from)
}

type Status string

func (status Status) GetParams() (string, string) {
  return "Status", string(status)
}

type StartTime string

func (startTime StartTime) GetParams() (string, string) {
  return "StartTime", string(startTime)
}

type ParentCallSid string

func (parentCallSid ParentCallSid) GetParams() (string, string) {
  return "ParentCallSid", string(parentCallSid)
}

type DateSent string

func (dateSent DateSent) GetParams() (string, string) {
  return "DateSent", string(dateSent)
}

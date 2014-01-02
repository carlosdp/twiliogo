package twiliogo

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "net/url"
  "encoding/json"
)

var testNumber = IncomingPhoneNumber {
  Sid: "testsid",
  AccountSid: "AC3TestAccount",
  FriendlyName: "testname",
  PhoneNumber: "+1 (444) 444-4444",
  VoiceUrl: "http://test.com",
  VoiceMethod: "POST",
  VoiceFallbackUrl: "http://fail.com",
  VoiceFallbackMethod: "GET",
  StatusCallback: "http://status.com",
  StatusCallbackMethod: "GET",
  SmsUrl: "http://sms.com",
  SmsMethod: "GET",
  DateCreated: "2013-05-11",
  DateUpdated: "2013-05-11",
  ApiVersion: "2008-04-01",
  Uri: "/2010-04-01/Accounts/AC3TestAccount/Messages/testsid.json",
}

func TestBuyPhoneNumber(t *testing.T) {
  client := new(MockClient)

  numberJson, _ := json.Marshal(testNumber)

  params := url.Values{}
  params.Set("PhoneNumber", "4444444444")

  client.On("post", params, client.RootUrl() + "/IncomingPhoneNumbers.json").Return(numberJson, nil)

  number, err := BuyPhoneNumber(client, PhoneNumber("4444444444"))

  client.Mock.AssertExpectations(t)

  if assert.Nil(t, err, "Error unmarshaling number") {
    assert.Equal(t, number.Sid, "testsid", "Number malformed")
  }
}

func TestGetIncomingPhoneNumber(t *testing.T) {
  client := new(MockClient)

  numberJson, _ := json.Marshal(testNumber)

  client.On("get", url.Values{}, client.RootUrl() + "/IncomingPhoneNumbers/testsid.json").Return(numberJson, nil)

  number, err := GetIncomingPhoneNumber(client, "testsid")

  client.Mock.AssertExpectations(t)

  if assert.Nil(t, err, "Error unmarshaling number") {
    assert.Equal(t, number.Sid, "testsid", "Number malformed")
  }
}


package twilio

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCallList(t *testing.T) {
  client := newClient(API_KEY, API_TOKEN)

  callList, err := GetCallList(client)

  if assert.Nil(t, err, "Failed to retrieve call list") {
    calls := callList.GetCalls()

    assert.NotNil(t, calls, "Failed to retrieve calls")
  }
}

func TestMakingCall(t *testing.T) {
  client := newClient(TEST_KEY, TEST_TOKEN)

  message, err := MakeCall(client, TEST_FROM_NUMBER, TO_NUMBER, Callback("http://test.com"))

  if assert.Nil(t, err, "Failed to make call") {
    assert.Equal(t, message.Status, "queued", "Making Call failed, status: " + message.Status)
  }
}

func TestCallListNextPage(t *testing.T) {
  client := newClient(API_KEY, API_TOKEN)

  callList, err := GetCallList(client)

  if assert.Nil(t, err, "Failed to retrieve call list") {
    nextPageCallList, err := callList.NextPage()

    if assert.Nil(t, err, "Failed to retrieve next page") {
      assert.Equal(t, nextPageCallList.Page, 1, "Page incorrect on next page")
    }
  }
}

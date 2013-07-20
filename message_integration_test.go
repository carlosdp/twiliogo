package twilio

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestMessageList(t *testing.T) {
  client := newClient(API_KEY, API_TOKEN)

  messageList, err := GetMessageList(client)

  assert.Nil(t, err, "Failed to retrieve message list")

  messages := messageList.Messages

  assert.NotNil(t, messages, "Failed to retrieve messages")
}

func TestSendSMS(t *testing.T) {
  client := newClient(TEST_KEY, TEST_TOKEN)

  message, err := SendMessage(client, TEST_FROM_NUMBER, TO_NUMBER, "Test Message")

  assert.Nil(t, err, "Failed to Send SMS")

  assert.Equal(t, message.Status, "queued", "Sending SMS failed, status: " + message.Status)
}

func TestMessageListNextPage(t *testing.T) {
  client := newClient(API_KEY, API_TOKEN)

  messageList, err := GetMessageList(client)

  assert.Nil(t, err, "Failed to retrieve message list")

  nextPageMessageList, err := messageList.NextPage()

  assert.Nil(t, err, "Failed to retrieve message list")

  assert.Equal(t, nextPageMessageList.Page, 1, "Page incorrect on next page")
}

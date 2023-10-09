package test

import (
	"secure-messaging-system/pkg"
	"testing"
	"time"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

const (
	senderIDMismatchErr      = "SenderID mismatch"
	receiverIDMismatchErr    = "ReceiverID mismatch"
	encryptedTextMismatchErr = "EncryptedText mismatch"
)

func Test_NewMessage(t *testing.T) {
	t.Run("happy path: create a new message", func(t *testing.T) {
		senderID := fake.CharactersN(10)
		receiverID := fake.CharactersN(10)
		encryptedText := fake.Sentence()

		message := pkg.NewMessage(senderID, receiverID, encryptedText)

		assert.Equal(t, senderID, message.SenderID, senderIDMismatchErr)
		assert.Equal(t, receiverID, message.ReceiverID, receiverIDMismatchErr)
		assert.Equal(t, encryptedText, message.EncryptedText, encryptedTextMismatchErr)

	})

}

func TestMessageBuilder(t *testing.T) {
	t.Run("happy path: create a new message using builder", func(t *testing.T) {
		senderID := fake.CharactersN(10)
		receiverID := fake.CharactersN(10)
		encryptedText := fake.Sentence()

		message, err := pkg.NewMessageBuilder().
			WithSenderID(senderID).
			WithReceiverID(receiverID).
			WithEncryptedText(encryptedText).
			Build()

		assert.NoError(t, err, "Error building the message")

		assert.Equal(t, senderID, message.SenderID, senderIDMismatchErr)
		assert.Equal(t, receiverID, message.ReceiverID, receiverIDMismatchErr)
		assert.Equal(t, encryptedText, message.EncryptedText, encryptedTextMismatchErr)

		// Make sure that the timestamp was set
		assert.False(t, message.Timestamp.IsZero(), "Timestamp was not set")
	})
}

func TestMessageConversion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message *pkg.Message
		wantErr bool
	}{
		{
			name: "happy path: conversion to JSON and from JSON",
			message: &pkg.Message{
				SenderID:      fake.CharactersN(10),
				ReceiverID:    fake.CharactersN(10),
				Timestamp:     time.Now(),
				EncryptedText: fake.Paragraph(),
			},
			wantErr: false,
		},
		{
			name: "error path: invalid UTF-8 sequence",
			message: &pkg.Message{
				SenderID:      fake.CharactersN(10),
				ReceiverID:    fake.CharactersN(10),
				Timestamp:     time.Now(),
				EncryptedText: string([]byte{0x80, 0x81, 0x82, 0x83}),
			},
			wantErr: true,
		},
		{
			name:    "error path: nil Message reference",
			message: nil,
			wantErr: true,
		},
		{
			name: "incomplete Message",
			message: &pkg.Message{
				SenderID:  fake.CharactersN(10),
				Timestamp: time.Now(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Convert to JSON
			jsonData, err := tt.message.ToJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			// Convert from JSON
			msgFromJSON, err := pkg.MessageFromJSON(jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.True(t, messagesAreEquivalent(tt.message, msgFromJSON))
			}
		})
	}
}

// Helper function to compare two messages
func messagesAreEquivalent(msg1, msg2 *pkg.Message) bool {
	return msg1.SenderID == msg2.SenderID &&
		msg1.ReceiverID == msg2.ReceiverID &&
		msg1.EncryptedText == msg2.EncryptedText &&
		msg1.Timestamp.Unix() == msg2.Timestamp.Unix()
}

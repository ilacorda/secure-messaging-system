package test

import (
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"secure-messaging-system/pkg"
	pb "secure-messaging-system/proto"
	"testing"
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

		assert.Equal(t, senderID, message.SenderId, senderIDMismatchErr)
		assert.Equal(t, receiverID, message.ReceiverId, receiverIDMismatchErr)
		assert.Equal(t, encryptedText, message.EncryptedText, encryptedTextMismatchErr)

	})

}

func TestMessageBuilder(t *testing.T) {
	t.Run("happy path: create a new message using the builder", func(t *testing.T) {
		senderID := fake.CharactersN(10)
		receiverID := fake.CharactersN(10)
		encryptedText := fake.Sentence()

		message, err := pkg.NewMessageBuilder().
			WithSenderID(senderID).
			WithReceiverID(receiverID).
			WithEncryptedText(encryptedText).
			Build()

		assert.NoError(t, err, "Error building the message")

		assert.Equal(t, senderID, message.SenderId, senderIDMismatchErr)
		assert.Equal(t, receiverID, message.ReceiverId, receiverIDMismatchErr)
		assert.Equal(t, encryptedText, message.EncryptedText, encryptedTextMismatchErr)

		// Make sure that the timestamp was set
		assert.NotEqual(t, int64(0), message.Timestamp, "Timestamp was not set")
	})
}

func TestMessageConversion(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		message *pb.Message
		wantErr bool
	}{
		{
			name: "happy path: conversion to JSON and from JSON",
			message: &pb.Message{
				SenderId:      fake.CharactersN(10),
				ReceiverId:    fake.CharactersN(10),
				Timestamp:     int64(0),
				EncryptedText: fake.Paragraph(),
			},
			wantErr: false,
		},
		{
			name: "error path: invalid UTF-8 sequence",
			message: &pb.Message{
				SenderId:      fake.CharactersN(10),
				ReceiverId:    fake.CharactersN(10),
				Timestamp:     int64(0),
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
			message: &pb.Message{
				SenderId:  fake.CharactersN(10),
				Timestamp: int64(0),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			jsonData, err := pkg.ToJSON(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

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

// Helper function

// messagesAreEquivalent compares two pb.Message objects and determines if they are equivalent.
func messagesAreEquivalent(msg1, msg2 *pb.Message) bool {
	// Compare individual fields of the messages.
	return msg1.SenderId == msg2.SenderId &&
		msg1.ReceiverId == msg2.ReceiverId &&
		msg1.EncryptedText == msg2.EncryptedText &&
		msg1.Timestamp == msg2.Timestamp
}

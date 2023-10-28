package test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"secure-messaging-system/pkg"
	pb "secure-messaging-system/proto"
	"testing"
	"time"
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
		timestamp := time.Now()

		message := pkg.NewMessage(senderID, receiverID, encryptedText, &timestamp)

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
		timestamp := time.Now()

		message, err := pkg.NewMessageBuilder().
			WithSenderID(senderID).
			WithReceiverID(receiverID).
			WithEncryptedText(encryptedText).
			WithTimestamp(timestamp).
			Build()

		assert.NoError(t, err, "Error building the message")

		assert.Equal(t, senderID, message.SenderId, "Mismatched Sender ID")
		assert.Equal(t, receiverID, message.ReceiverId, "Mismatched Receiver ID")
		assert.Equal(t, encryptedText, message.EncryptedText, "Mismatched Encrypted Text")

		assert.Equal(t, timestamppb.New(timestamp), message.Timestamp, "Timestamp mismatch")
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
				Timestamp:     timestamppb.New(time.Now()),
				EncryptedText: fake.Paragraph(),
			},
			wantErr: false,
		},
		{
			name: "error path: invalid UTF-8 sequence",
			message: &pb.Message{
				SenderId:      fake.CharactersN(10),
				ReceiverId:    fake.CharactersN(10),
				Timestamp:     timestamppb.New(time.Now()),
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
				Timestamp: timestamppb.New(time.Now()),
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

			msgFromJSON, err := pkg.FromJSON(jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessageFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			diff := cmp.Diff(
				tt.message,
				msgFromJSON,
				cmpopts.IgnoreUnexported(pb.Message{}, timestamppb.Timestamp{}),
			)
			if diff != "" {
				t.Errorf("Mismatch (-Original +FromJSON):\n%s", diff)
			}
		})
	}
}

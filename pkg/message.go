package pkg

import (
	"encoding/json"
	"fmt"
	"time"
	"unicode/utf8"
)

type Message struct {
	SenderID      string    `json:"sender_id"`
	ReceiverID    string    `json:"receiver_id"`
	Timestamp     time.Time `json:"timestamp"`
	EncryptedText string    `json:"encrypted_text"`
}

type MessageBuilder struct {
	message *Message
}

// NewMessageBuilder creates a new MessageBuilder instance.
func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		message: &Message{},
	}
}

// WithSenderID sets the SenderID of the message.
func (b *MessageBuilder) WithSenderID(senderID string) *MessageBuilder {
	b.message.SenderID = senderID
	return b
}

// WithReceiverID sets the ReceiverID of the message.
func (b *MessageBuilder) WithReceiverID(receiverID string) *MessageBuilder {
	b.message.ReceiverID = receiverID
	return b
}

// WithEncryptedText sets the EncryptedText of the message.
func (b *MessageBuilder) WithEncryptedText(encryptedText string) *MessageBuilder {
	b.message.EncryptedText = encryptedText
	return b
}

// WithTimestamp sets the Timestamp of the message. If not set, it defaults to the current time when building.
func (b *MessageBuilder) WithTimestamp(timestamp time.Time) *MessageBuilder {
	b.message.Timestamp = timestamp
	return b
}

// Build finalizes the building process and returns the constructed Message.
func (b *MessageBuilder) Build() (*Message, error) {
	if b.message.Timestamp.IsZero() {
		b.message.Timestamp = time.Now()
	}

	// You can add additional validation checks here if needed.

	return b.message, nil
}

func NewMessage(senderID, receiverID, encryptedText string) *Message {
	return &Message{
		SenderID:      senderID,
		ReceiverID:    receiverID,
		Timestamp:     time.Now(),
		EncryptedText: encryptedText,
	}
}

func (m *Message) ToJSON() ([]byte, error) {
	if m == nil {
		return nil, fmt.Errorf("nil Message reference")
	}
	if !utf8.ValidString(m.EncryptedText) {
		return nil, fmt.Errorf("invalid UTF-8 sequence in EncryptedText")
	}
	return json.Marshal(m)
}

func MessageFromJSON(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return &msg, err
}

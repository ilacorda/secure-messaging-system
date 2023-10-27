package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/rs/zerolog/log"
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
	if senderID == "" {
		log.Warn().Msg("Provided empty SenderID")
		return b
	}
	b.message.SenderID = senderID
	return b
}

// WithReceiverID sets the ReceiverID of the message.
func (b *MessageBuilder) WithReceiverID(receiverID string) *MessageBuilder {
	if receiverID == "" {
		log.Warn().Msg("Provided empty ReceiverID")
		return b
	}
	b.message.ReceiverID = receiverID
	return b
}

// WithEncryptedText sets the EncryptedText of the message.
func (b *MessageBuilder) WithEncryptedText(encryptedText string) *MessageBuilder {
	if encryptedText == "" {
		log.Warn().Msg("Provided empty EncryptedText")
		return b
	}
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
	if b.message.SenderID == "" || b.message.ReceiverID == "" || b.message.EncryptedText == "" {
		return nil, errors.New("message is incomplete")
	}

	if b.message.Timestamp.IsZero() {
		b.message.Timestamp = time.Now()
	}

	return b.message, nil
}

func NewMessage(senderID, receiverID, encryptedText string) *Message {
	if senderID == "" || receiverID == "" || encryptedText == "" {
		log.Error().Msg("Cannot create a new message with empty fields")
		return nil
	}

	return &Message{
		SenderID:      senderID,
		ReceiverID:    receiverID,
		Timestamp:     time.Now(),
		EncryptedText: encryptedText,
	}
}

func (m *Message) ToJSON() ([]byte, error) {
	if m == nil {
		log.Error().Msg("Attempted to convert a nil Message to JSON")
		return nil, fmt.Errorf("nil Message reference")
	}
	if !utf8.ValidString(m.EncryptedText) {
		log.Error().Msg("EncryptedText contains invalid UTF-8 sequence")
		return nil, fmt.Errorf("invalid UTF-8 sequence in EncryptedText")
	}
	return json.Marshal(m)
}

func MessageFromJSON(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal JSON data to Message")
	}
	return &msg, err
}

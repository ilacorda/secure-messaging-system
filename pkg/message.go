package pkg

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "secure-messaging-system/proto"
	"time"
	"unicode/utf8"
)

type MessageBuilder struct {
	message *pb.Message
}

// NewMessageBuilder creates a new MessageBuilder instance.
func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		message: &pb.Message{},
	}
}

// WithSenderID sets the SenderID of the message.
func (b *MessageBuilder) WithSenderID(senderID string) *MessageBuilder {
	if senderID == "" {
		log.Warn().Msg("Provided empty SenderId")
		return b
	}
	b.message.SenderId = senderID
	return b
}

// WithReceiverID sets the ReceiverID of the message.
func (b *MessageBuilder) WithReceiverID(receiverID string) *MessageBuilder {
	if receiverID == "" {
		log.Warn().Msg("Provided empty ReceiverId")
		return b
	}
	b.message.ReceiverId = receiverID
	return b
}

func (b *MessageBuilder) WithTimestamp(timeStamp time.Time) *MessageBuilder {
	b.message.Timestamp = timestamppb.New(timeStamp)
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

// Build finalizes the building process and returns the constructed Message.
func (b *MessageBuilder) Build() (*pb.Message, error) {
	if b.message.SenderId == "" ||
		b.message.ReceiverId == "" ||
		b.message.EncryptedText == "" ||
		b.message.Timestamp == nil {
		return nil, errors.New("message is incomplete")
	}

	return b.message, nil
}

func NewMessage(senderID, receiverID, encryptedText string, timestamp *time.Time) *pb.Message {
	if senderID == "" ||
		receiverID == "" ||
		encryptedText == "" ||
		timestamp == nil {
		log.Error().Msg("Cannot create a new message with empty fields")
		return nil
	}

	return &pb.Message{
		SenderId:      senderID,
		ReceiverId:    receiverID,
		Timestamp:     timestamppb.New(*timestamp),
		EncryptedText: encryptedText,
	}
}

// ToJSON converts a message to its JSON representation.
func ToJSON(msg *pb.Message) ([]byte, error) {
	if msg == nil {
		return nil, errors.New("nil message provided")
	}
	if !utf8.ValidString(msg.EncryptedText) {
		return nil, errors.New("invalid UTF-8 sequence in EncryptedText")
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func FromJSON(data []byte) (*pb.Message, error) {
	var msg pb.Message

	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, errors.New("failed to unmarshal JSON into Message struct: " + err.Error())
	}

	return &msg, nil
}

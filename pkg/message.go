package pkg

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	pb "secure-messaging-system/proto"
	"time"
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
		log.Warn().Msg("Provided empty SenderID")
		return b
	}
	b.message.SenderId = senderID
	return b
}

// WithReceiverID sets the ReceiverID of the message.
func (b *MessageBuilder) WithReceiverID(receiverID string) *MessageBuilder {
	if receiverID == "" {
		log.Warn().Msg("Provided empty ReceiverID")
		return b
	}
	b.message.ReceiverId = receiverID
	return b
}

func (b *MessageBuilder) WithTimestamp(timestamp time.Time) *MessageBuilder {
	b.message.Timestamp = timestamp.Unix()
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
	if b.message.SenderId == "" || b.message.ReceiverId == "" || b.message.EncryptedText == "" {
		return nil, errors.New("message is incomplete")
	}

	if b.message.Timestamp == 0 {
		b.message.Timestamp = time.Now().Unix()
	}

	return b.message, nil
}

func NewMessage(senderID, receiverID, encryptedText string) *pb.Message {
	if senderID == "" || receiverID == "" || encryptedText == "" {
		log.Error().Msg("Cannot create a new message with empty fields")
		return nil
	}

	return &pb.Message{
		SenderId:      senderID,
		ReceiverId:    receiverID,
		Timestamp:     time.Now().Unix(),
		EncryptedText: encryptedText,
	}
}

// ToJSON converts a pb.Message to a JSON byte array.
func ToJSON(msg *pb.Message) ([]byte, error) {
	return json.Marshal(msg)
}

// MessageFromJSON converts a JSON byte array to a pb.Message.
func MessageFromJSON(data []byte) (*pb.Message, error) {
	var msg pb.Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

package db

import "time"

type MessageEntity struct {
	ID            int       `db:"id"`
	SenderID      string    `db:"sender_id"`
	ReceiverID    string    `db:"receiver_id"`
	Timestamp     time.Time `db:"timestamp"`
	EncryptedText string    `db:"encrypted_text"`
}

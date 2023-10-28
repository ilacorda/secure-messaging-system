CREATE TABLE Messages (
                          id SERIAL PRIMARY KEY,
                          sender_id VARCHAR(255) NOT NULL,
                          receiver_id VARCHAR(255) NOT NULL,
                          timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
                          encrypted_text TEXT NOT NULL
);
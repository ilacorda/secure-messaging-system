## Secure Messaging System
This repository contains a secure messaging system implementation, designed to send and receive encrypted messages using the AES (Advanced Encryption Standard) algorithm. 
The system adheres to secure coding practices and robust design principles, ensuring the security and integrity of your messages.

### Instructions for the task
For this assignment, the goal of this task is to design and implement a secure messaging system. In this system, a sender will be able to send an encrypted message to a receiver, who can then decrypt the message.

#### Requirements:

- Messages should be sent as JSON payloads. Each message should at least include the following fields: sender ID, receiver ID, timestamp, and the encrypted message.
- Use AES (Advanced Encryption Standard) for encryption and decryption of messages. The encryption key should be securely generated and stored.
- Implement a mechanism to securely exchange the encryption key between the sender and the receiver (for example, Diffie-Hellman key exchange or public key cryptography).
- Transfer of messages may occur either over a network or any other method (usb pendrive, dropbox, cloud service)
- Use secure coding best practices: for instance, validate input, employ proper error handling
- Include a brief documentation explaining the design decisions, the choice of cryptographic algorithms, and how to build and run / build and test the code.

#### Deliverables:
A GitHub repository, containing the following:
- Source code.
- A README file explaining the design decisions, how to build and run the code.
- A basic GitHub CI Action to ensure the application builds and any unit tests pass.

### Design Choices

- **Simplified Directory Structure**: To make the application more accessible and easier to navigate, a rather flat directory structure has been adopted. I have included the encryption and key exchange logic together with the `Message` struct and its related functions in the `pkg` directory. 
The `cmd` directory contains the main entry point of the application. `test` contains the unit tests for the application.

- **AES Encryption**: AES, a widely recognized symmetric encryption standard, is used for its efficiency and strong security guarantees. I implemented a validation for 16, 24, or 32 bytes long to match AES-128, AES-192, or AES-256 respectively. 
The implementation also uses CFB Mode, which operates by encrypting the current block and using parts of the output as feedback for the next block.

- **Key Exchange**: I employed the Diffie-Hellman key exchange mechanism to securely exchange the encryption key between the sender and the receiver as offered by the `crypto/elliptic` package. Because of the time constraints and ver basic knowledge of such mechanism, I just showcased it in the `main.go` function

- **JSON Payloads**: Messages are structured as JSON payloads, ensuring compatibility and easy integration with other systems. As per the specification, the message payload contains the following fields:
    - `sender_id`: The sender ID of the message.
    - `recipient_id`: The recipient ID of the message.
    - `encrypted_text`: The encrypted textual message.
    - `timestamp`: The timestamp of the message.

Protocolbuffers would normall be the preferable choice because of efficiency, fast serialisation and deserialisation (as well as built-in validations), but I chose JSON for its simplicity and ease of use and to align with the recommended time of completion for the task. 

### Structure
```
├── cmd/
│   ├── main.go                  # CLI application's main entry.
├── pkg/
│   ├── encryption.go            # AES-GCM encryption & decryption logic.
│   ├── key_exchange.go          # Diffie-Hellman key exchange mechanism.
│   └── message.go               # Message creation, validation & handling and basic builder.
│
├── test/
│   ├── encryption_test.go           # Tests for encryption functionalities.
│   └── message_test.go              # Tests for message-related utilities.
│
├── README.md                        # README file with instructions and design choices.
│
├── .github/
│   └── workflows/
│       └── build.yml                # GitHub Actions CI/CD workflow.
│
├── go.mod
├── go.sum
```

### How to Build and Run the Application

`cd` into the `cmd` directory and run the following: 

`go build -o secure-messaging-system`

Run the application:

`./secure-messaging-system`

### How to Run the Tests
To run tests, from the root level, run the following command:

`go test ./...`

Test flags can be added to the command above to generate coverage reports, run exclusively unit tests, etc. Considering that there are only have unit tests, I did not include them


### GitHub CI Action
A basic GitHub CI action is set up to ensure that the application builds successfully and all unit tests pass. 
This CI action can be found in the `.github/workflows/build.yml` file.

### Future work and expected improvements
- **Key Management**: The application showcases the Elliptic Curve Diffie-Hellman (ECDH) key exchange mechanism to dynamically generating and deriving encryption keys for sessions. In future iterations, I might consider integrating with advanced key management systems like HashiCorp Vault for further security enhancements.
- **Message Validation**: The application currently validates the message payload to ensure that it is a valid JSON object. However, I would like to improve this by using Protobuf to define and validate the message structure.
- **Message Integrity**: The application currently does not provide any guarantees about the integrity of the message. I would like to improve this by adding a message authentication code (MAC) to the message payload.
- **Data Persistence**: The application currently does not actively persist messages. I would like to improve this by adding a database to the application, so that messages can be stored and retrieved.
- **Message Delivery**: The application currently does not provide any guarantees about the delivery of the message. I would like to improve this by adding a message queue to the application, so that messages can be queued and delivered.
- **API implementation**: The application currently does not provide any API endpoints. I would like to improve this by adding API endpoints to the application, so that messages can be sent and received via API calls. Tracing and logging can also be added to the API endpoints. In particular, the usage of OpenTelemetry and OpenTracing can be explored.
- **Test coverage**: Test coverage can be improved by adding more unit tests and integration tests. If persistence is implemented, I can also add repository tests to test the database layer.
- **Improve CICD pipeline**: The CICD pipeline can be improved by adding more stages, such as linting and security scanning, coverage reports (e.g. codecov), etc.

### Task Duration
The task was time-boxed and this currently version was reached in approximately 1.5 hours. While this could have been expanded, I wanted to align with the guidlines in the instructions.

### Subsequent PRs
- Added structured logging with [zerolog](https://github.com/rs/zerolog) whose basic setup for Console logging is defined in `logger.go` and called in `main.go`
- ADD linter to Github actions and include Makefile with targets
- Move to proto definitions as opposed to JSON
# Pub/Sub Application in Go

Pub/Sub (Publish/Subscribe) application implemented in Go. It demonstrates how to create a basic Pub/Sub system using goroutines and channels.

## Features

- Subscribe to topics and receive messages asynchronously
- Publish messages to specific topics
- Decoupled communication between publishers and subscribers

## Getting Started

To run the application, you need to have Go installed on your system. Then, follow these steps:


1. Run the main Go file:

    ```bash
    go run cmd/main.go
    ```

## Usage

1. **Subscribe**: Create subscribers to listen to specific topics by calling the `Subscribe` method on the `PubSub` struct.

2. **Publish**: Send messages to subscribers by calling the `Publish` method on the `PubSub` struct with the appropriate topic and message.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

by **Eduardo Raider** - Software Engineer
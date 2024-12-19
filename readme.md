# 🚀 Simple Messenger API

WebSocket-based messenger API built with **Golang** and **Gorilla WebSocket**. Connect multiple clients and experience real-time messaging with ease.

---

## ✨ Features

- 📡 **Real-time Messaging**: Instant message delivery across connected clients.

---

## 📋 Requirements

- **Go** 1.18+
- **Gorilla WebSocket** Library

---

## 🛠️ Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/adityasuryadi/messenger.git
   cd messenger
   ```

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up configuration**:

   Create a `config.yaml` file in the root directory with the following structure:

   ```yaml
   service:
     port: ":8080"
   jwt:
     secretJWT: "rahasia"
     ttl: 30 #in minutes
   ```

4. **Run the server**:

   ```bash
   go run main.go
   ```

5. The WebSocket server will start at `http://localhost:8080`.

---

## 📡 API Endpoint

### WebSocket Connection

- **Endpoint**: `/ws`
- **Method**: `GET`

**Connect using any WebSocket client**, like:

- postman:
  ```bash
  websocat ws://localhost:8080/ws
  ```

#### JSON Body Example

```json
{
    "from": "Adit",
    "message": "Hai, bro"
}
```

---

## 🚀 How It Works

1. **Connect Clients**: Multiple clients establish WebSocket connections to the server.
2. **Send Messages**: A client sends a message via the WebSocket connection.
3. **Broadcast**: All other clients receive the message in real-time, except the sender.

### Example Flow

- **Client A** sends:

  ```
  Hello from Client A
  ```

- **Clients B and C** receive:

  ```
  Hello from Client A
  ```

- **Client A** does not receive its own message.

---

## 🧪 Example Test

1. Start the server:

   ```bash
   go run main.go
   ```

2. Connect multiple WebSocket clients:

   ```bash
   websocat ws://localhost:8080/ws
   ```

3. Send and receive messages in real-time 🎉.

---

## 📜 License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.

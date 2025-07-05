# **Real-time Stock Data gRPC Service** 📈

This project showcases a robust and scalable real-time stock data service built with Go and gRPC. It's designed to demonstrate various gRPC communication patterns, including Unary, Server Streaming, Client Streaming, and Bidirectional Streaming, providing a comprehensive example of building high-performance, resilient microservices. Integrated with the Alpha Vantage API, it offers up-to-the-minute stock price information and simulated historical data processing.

## 🚀 Installation

Getting this project up and running locally is straightforward. Just follow these steps:

### 📥 Clone the Repository

First, grab the source code from GitHub:

```bash
git clone https://github.com/samueltuoyo15/Live-Stock-Grpc.git
cd Live-Stock-Grpc
```

### ⚙️ Generate Protocol Buffer Code

The gRPC service definitions in `proto/stock.proto` need to be compiled into Go code. A handy Makefile rule takes care of this:

```bash
make proto
```

This command generates the necessary Go client and server stubs in `proto/generated/stockpb`.

### 🔑 Configure Alpha Vantage API Key

This service integrates with the Alpha Vantage API to fetch real-time stock prices. You'll need an API key:

*   Go to [Alpha Vantage](https://www.alphavantage.co/) and sign up for a free API key.
*   Create a `.env` file in the project's root directory (`Live-Stock-Grpc/`) and add your API key:

    ```env
    ALPHA_VANTAGE_API_KEY=YOUR_ALPHA_VANTAGE_API_KEY
    ```

### 📦 Install Dependencies

Ensure all Go dependencies are installed:

```bash
go mod tidy
```

## 🏃 Usage

Once set up, you can run the server and then the client to see the gRPC communication patterns in action.

### ⬆️ Start the gRPC Server

Open your terminal in the project root and start the server:

```bash
go run server/main.go
```

You should see output indicating the server is listening:

```
INFO Grpc Server is listening on port :8080
```

### ⬇️ Run the gRPC Client

In a *separate* terminal, run the client application:

```bash
go run client/main.go
```

The client will automatically perform a series of gRPC calls demonstrating each communication type:

#### Unary RPC (Get Stock)

The client makes a single request for a stock symbol (e.g., AAPL) and receives a single response with its current price.

```
INFO Starting Unary rpc
INFO Stock Response symbol=AAPL price=$... timestamp=...
```

#### Server Streaming RPC (Watch Stock)

The client requests continuous updates for a specific stock (e.g., GOOGL). The server will stream multiple price updates over time.

```
INFO Starting Server Streaming Rpc
INFO Stock Response symbol=GOOGL price=$... timestamp=...
INFO Stock Response symbol=GOOGL price=$... timestamp=...
...
INFO Stream Ended
```

#### Client Streaming RPC (Upload Stock History)

The client sends a stream of stock requests (e.g., multiple MSFT symbols), simulating an upload of historical data. The server processes these and sends a single summary response when done.

```
INFO Starting client streaming RPC
INFO sending stock symbol=MSFT
INFO sending stock symbol=MSFT
INFO sending stock symbol=MSFT
INFO upload summary message="Upload Complete" count=3
```

#### Bidirectional Streaming RPC (Chat Stock)

The client and server send messages concurrently. The client sends stock symbols (e.g., AMZN, TSLA, NVDA), and the server responds with immediate updates for each, showcasing a real-time interactive session.

```
INFO Starting Bidirectional Streaming RPC
INFO Sending stock symbol=AMZN
INFO Received Stock Update symbol=AMZN price=$... timestamp=...
INFO Sending stock symbol=TSLA
INFO Received Stock Update symbol=TSLA price=$... timestamp=...
...
INFO Stream ended
INFO ChatStock completed
```

## ✨ Features

This project highlights several key capabilities:

*   **Diverse gRPC Communication Patterns**: Demonstrates Unary, Server Streaming, Client Streaming, and Bidirectional Streaming RPCs for various interaction models.
*   **Real-time Stock Price Integration**: Fetches live stock prices using the Alpha Vantage API, providing practical data integration.
*   **Robust Error Handling**: Includes basic error handling for network requests and gRPC calls.
*   **Structured Logging**: Utilizes Go's `slog` package for clear and structured logging, making it easier to monitor and debug.
*   **Environment Variable Management**: Securely handles API keys and other configurations using `godotenv`.
*   **Clean Architecture**: Separates client and server logic, promoting modularity and maintainability.
*   **Makefile Automation**: Simplifies the protobuf code generation process.

## 🛠️ Technologies Used

| Technology             | Description                                          | Link                                                  |
| :--------------------- | :--------------------------------------------------- | :---------------------------------------------------- |
| **Go (Golang)**        | Primary programming language for service development | [golang.org](https://golang.org/)                     |
| **gRPC**               | High-performance RPC framework                       | [grpc.io](https://grpc.io/)                           |
| **Protocol Buffers**   | Language-agnostic data serialization format          | [developers.google.com/protocol-buffers](https://developers.google.com/protocol-buffers) |
| **Alpha Vantage API**  | External API for real-time stock market data         | [alphavantage.co](https://www.alphavantage.co/)       |
| **`slog`**             | Go's built-in structured logging package             | [pkg.go.dev/log/slog](https://pkg.go.dev/log/slog)      |
| **`godotenv`**         | Loads environment variables from `.env` files        | [github.com/joho/godotenv](https://github.com/joho/godotenv) |
| **Makefile**           | Build automation and script management               | [gnu.org/software/make](https://www.gnu.org/software/make) |

## 🤝 Contributing

Contributions are warmly welcome! If you have suggestions for improvements, new features, or bug fixes, please feel free to:

*   ✨ Fork the repository.
*   🌳 Create a new branch (`git checkout -b feature/AmazingFeature`).
*   📝 Make your changes and commit them (`git commit -m 'Add some AmazingFeature'`).
*   ⬆️ Push to the branch (`git push origin feature/AmazingFeature`).
*   🚀 Open a pull request.

Please ensure your code adheres to standard Go best practices and includes appropriate tests if applicable.

## 📝 License

This project is open-source and available for use and modification.

## 👤 Author

**Samuel Tuoyo**
*   GitHub: [@samueltuoyo15](https://github.com/samueltuoyo15)
*   LinkedIn: [Your LinkedIn Profile](https://linkedin.com/in/your-profile)
*   Twitter: [@YourTwitterHandle](https://twitter.com/YourTwitterHandle)

---

[![Go Version](https://img.shields.io/badge/Go-1.24.1-00ADD8?logo=go)](https://golang.org/)
[![gRPC](https://img.shields.io/badge/gRPC-v1.72.0-blue?logo=grpc)](https://grpc.io/)
[![Alpha Vantage API](https://img.shields.io/badge/API-Alpha%20Vantage-green?logo=react)](https://www.alphavantage.co/)
[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)
**Real-time Stock Data Streaming with Go gRPC 📈**

Dive into `Live-Stock-Grpc`, a robust demonstration of real-time stock data exchange using Go and gRPC. This project showcases the power of various gRPC communication patterns, from simple unary calls to complex bidirectional streaming, all while fetching live stock prices. It's a perfect example of building performant and scalable microservices.

## Installation

Setting up `Live-Stock-Grpc` is straightforward. Just follow these steps to get the project running on your local machine.

*   **Prerequisites**:
    *   Go (version 1.22 or higher, as defined in `go.mod`)
    *   Protocol Buffers Compiler `protoc`
*   **Steps**:
    1.  **Clone the repository**:
        ```bash
        git clone https://github.com/samueltuoyo15/Live-Stock-Grpc.git
        cd Live-Stock-Grpc
        ```
    2.  **Install `protoc` and its Go plugins**:
        Ensure you have `protoc` installed on your system (e.g., via `brew install protobuf` on macOS, `sudo apt install protobuf-compiler` on Debian/Ubuntu, or from source).
        Then, install the necessary Go plugins:
        ```bash
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
        ```
        Make sure your `GOPATH/bin` directory is included in your system's `PATH` environment variable so `protoc` can find the plugins.
    3.  **Generate Go code from Protocol Buffers definitions**:
        This step compiles the `.proto` file into Go source code for gRPC service definitions and message structs.
        ```bash
        protoc \
            --go_out=./proto/generated --go_opt=paths=source_relative \
            --go-grpc_out=./proto/generated --go-grpc_opt=paths=source_relative \
            ./proto/stock.proto
        ```
    4.  **Install Go dependencies**:
        Fetch all required Go modules.
        ```bash
        go mod tidy
        ```
    5.  **Configure Alpha Vantage API Key**:
        Create a `.env` file in the project's root directory and add your Alpha Vantage API key. You can obtain a free API key from [Alpha Vantage](https://www.alphavantage.co/).
        ```
        ALPHA_VANTAGE_API_KEY=YOUR_ALPHA_VANTAGE_API_KEY
        ```

## Usage

This project showcases a gRPC server and client interacting to fetch and stream real-time stock data.

1.  **Start the gRPC Server**:
    Open your terminal, navigate to the project root directory, and run the server application:
    ```bash
    go run ./server/main.go
    ```
    You should see an informational log indicating the server has started:
    ```
    INFO Grpc Server is listening on port :8080
    ```

2.  **Run the gRPC Client**:
    Open a *new* terminal window, navigate to the project root, and execute the client application:
    ```bash
    go run ./client/main.go
    ```
    The client will then sequentially demonstrate the four core gRPC communication patterns:

    *   **Unary RPC (`GetStock`)**: The client makes a single request for "AAPL" stock data and receives a single response.
        ```
        INFO Starting Unary rpc
        INFO Stock Response symbol=AAPL price=$185.00 timestamp=...
        ```
    *   **Server Streaming RPC (`WatchStock`)**: The client requests to "watch" "GOOGL" stock. The server then continuously streams updates back to the client.
        ```
        INFO Starting Server Streaming Rpc
        INFO Stock Response symbol=GOOGL price=$150.25 timestamp=...
        INFO Stock Response symbol=GOOGL price=$150.30 timestamp=...
        ... (multiple updates follow) ...
        INFO Stream Ended
        ```
    *   **Client Streaming RPC (`UploadStockHistory`)**: The client sends a series of "MSFT" stock requests to the server, and the server sends back a single summary response once all requests are received.
        ```
        INFO Starting client streaming RPC
        INFO sending stock symbol=MSFT
        INFO sending stock symbol=MSFT
        INFO sending stock symbol=MSFT
        INFO upload summary message="Upload Complete" count=3
        ```
    *   **Bidirectional Streaming RPC (`ChatStock`)**: The client and server exchange streams of messages concurrently. The client sends multiple stock symbols ("AMZN", "TSLA", "NVDA"), and the server responds with updates for each symbol as it receives them.
        ```
        INFO Starting Bidirectional Streaming RPC
        INFO Sending stock symbol=AMZN
        INFO Received Stock Update symbol=AMZN price=$175.10 timestamp=...
        INFO Sending stock symbol=TSLA
        INFO Received Stock Update symbol=TSLA price=$200.50 timestamp=...
        ... (sending and receiving interleaved) ...
        INFO Stream ended
        INFO ChatStock completed
        ```

## Features

*   🚀 **Comprehensive gRPC Demonstrations**: Implements all four gRPC communication patterns: Unary, Server Streaming, Client Streaming, and Bidirectional Streaming.
*   📈 **Real-time Stock Data Integration**: Fetches live stock prices using the Alpha Vantage API, providing up-to-date financial information.
*   🛠️ **Protocol Buffers**: Defines clear and efficient service contracts and data structures for seamless cross-language communication.
*   📦 **Structured Logging**: Utilizes Go's `slog` package for structured, context-rich logging, enhancing observability and debugging.
*   🐳 **Modular Design**: Separates client and server logic into distinct packages, promoting a clean and maintainable codebase.
*   ⚙️ **Automated Code Generation**: Includes a `Makefile` to streamline the generation of Go gRPC stubs from `.proto` definitions.

## Technologies Used

| Technology             | Description                                                                                                   | Link                                                                            |
| :--------------------- | :------------------------------------------------------------------------------------------------------------ | :------------------------------------------------------------------------------ |
| **Go**                 | The primary programming language used for developing both the server and client applications.                 | [golang.org](https://golang.org/)                                               |
| **gRPC**               | A high-performance, open-source universal RPC framework used for inter-service communication.                 | [grpc.io](https://grpc.io/)                                                     |
| **Protocol Buffers**   | Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data.            | [developers.google.com/protocol-buffers](https://developers.google.com/protocol-buffers) |
| **Alpha Vantage API**  | Provides real-time and historical financial data APIs for stocks, used to fetch live prices.                  | [alphavantage.co](https://www.alphavantage.co/)                                 |
| **`slog`**             | Go's standard library for structured, leveled logging, improving log readability and analysis.              | [pkg.go.dev/log/slog](https://pkg.go.dev/log/slog)                              |
| **`github.com/joho/godotenv`** | A simple Go library to load environment variables from a `.env` file, used for API key management. | [github.com/joho/godotenv](https://github.com/joho/godotenv)                    |

## Contributing

Contributions are always welcome and highly appreciated! If you're looking to contribute to `Live-Stock-Grpc`, please consider the following guidelines:

*   ✨ **Fork the repository** and create your branch from `main`.
*   🐛 **Report bugs** by opening a new issue with detailed steps to reproduce the problem.
*   💡 **Suggest enhancements** or new features by opening an issue, explaining your idea clearly.
*   📝 **Ensure your code adheres** to Go's best practices and the project's existing style guidelines.
*   🧪 **Write clear, concise commit messages** that explain the purpose of your changes.

Your efforts help make this project even better for everyone!

## License

This project is open-source.

## Author Info

**Samuel Tuoyo**
*   GitHub: [@samueltuoyo15](https://github.com/samueltuoyo15)
*   LinkedIn: [Your LinkedIn Profile](https://www.linkedin.com/in/your-profile)
*   Portfolio: [Your Portfolio Link](https://your-portfolio.com)

---

[![Go Version](https://img.shields.io/github/go-mod/go-version/samueltuoyo15/Live-Stock-Grpc)](https://golang.org/)
[![GitHub repo size](https://img.shields.io/github/repo-size/samueltuoyo15/Live-Stock-Grpc)](https://github.com/samueltuoyo15/Live-Stock-Grpc)
![GitHub stars](https://img.shields.io/github/stars/samueltuoyo15/Live-Stock-Grpc?style=social)
![GitHub forks](https://img.shields.io/github/forks/samueltuoyo15/Live-Stock-Grpc?style=social)

[![Readme was generated by Dokugen](https://img.shields.io/badge/Readme%20was%20generated%20by-Dokugen-brightgreen)](https://www.npmjs.com/package/dokugen)
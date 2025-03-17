# Caching Proxy

A lightweight Go CLI tool that starts a caching proxy server to forward requests to an origin server and cache the responses. It improves performance by serving cached responses for repeated requests.

[Roadmap.sh](https://roadmap.sh/projects/caching-server)

## Features

- Forward HTTP requests to an origin server
- Cache responses for improved performance
- Custom port configuration
- Cache status indicators in response headers
- Command to clear the cache

## Installation

### Prerequisites

- Go 1.18 or higher

### Option 1: Install using Go

```bash
go install github.com/yourusername/caching-proxy@latest
```

### Option 2: Build from source

```bash
# Clone the repository
git clone https://github.com/yourusername/caching-proxy.git

# Navigate to the project directory
cd caching-proxy

# Build the project
go build -o caching-proxy

# Move the binary to your PATH (optional)
mv caching-proxy /usr/local/bin/
```

### Option 3: Download pre-built binary

Download the appropriate binary for your operating system from the [releases page](https://github.com/hackermanpeter/caching-proxy/releases).

## Usage

### Starting the proxy server

```bash
caching-proxy --port <number> --origin <url>
```

#### Parameters:

- `--port`: The port on which the caching proxy server will run
- `--origin`: The URL of the server to which the requests will be forwarded

#### Example:

```bash
caching-proxy --port 3000 --origin http://dummyjson.com
```

This will start the proxy server on port 3000 and forward requests to `http://dummyjson.com`.

### Making requests

Once the proxy server is running, you can make requests to it using the proxy's address.

#### Example:

```bash
curl http://localhost:3000/products
```

This request will be forwarded to `http://dummyjson.com/products`, and the response will be cached.

### Cache indicators

The proxy adds a header to indicate whether the response came from the cache:

- `X-Cache: HIT` - Response served from cache
- `X-Cache: MISS` - Response fetched from origin server

### Clearing the cache

To clear the cache:

```bash
caching-proxy --clear-cache
```

## How it works

1. When a request is received, the proxy checks if the response is already cached
2. If cached, it returns the cached response with `X-Cache: HIT` header
3. If not cached, it forwards the request to the origin server, caches the response, and returns it with `X-Cache: MISS` header

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

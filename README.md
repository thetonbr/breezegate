## BreezeGate - Modern Load Balancer with Go

[![CI](https://github.com/thetonbr/breezegate/workflows/CI/badge.svg)](https://github.com/thetonbr/breezegate/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/thetonbr/breezegate)](https://goreportcard.com/report/github.com/thetonbr/breezegate)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/thetonbr/breezegate?status.svg)](https://godoc.org/github.com/thetonbr/breezegate)

BreezeGate is a modern, high-performance load balancer written in Go designed to handle incoming HTTP/HTTPS requests, distribute them across backend servers, and automatically manage TLS certificates via Let's Encrypt (using the ACME protocol). This project uses the `lego` library to handle TLS certificate management and supports reverse proxy functionalities with built-in health checking.

### Features

- **Dynamic Backend Management**: Easily configure backend servers and routes through a JSON configuration file
- **Health Checks**: Periodic health checks with configurable intervals to ensure traffic is routed only to healthy backend servers
- **Automatic TLS Certificates**: Automatically generate and manage SSL certificates using Let's Encrypt with DNS-01 challenge support
- **Round Robin Load Balancing**: Distribute requests evenly across healthy backend servers
- **Reverse Proxy**: Forward requests to backend servers seamlessly using Go's built-in `httputil.ReverseProxy`
- **Concurrent Processing**: Built with Go's concurrency patterns for high performance
- **Comprehensive Testing**: Full test suite with race condition detection and coverage reporting
- **Docker Support**: Ready-to-use Docker container for easy deployment
- **CI/CD Integration**: GitHub Actions workflows for automated testing, building, and releasing

### Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Development](#development)
- [Makefile Commands](#makefile-commands)
- [Docker](#docker)
- [CI/CD](#cicd)
- [Future Improvements](#future-improvements)

---

### Installation

1. **Clone the Repository**:
   
   ```bash
   git clone https://github.com/thetonbr/breezegate.git
   cd breezegate
   ```
   
2. **Install Dependencies**:

    Make sure you have Go installed on your system. To install the required libraries, run:
      ```bash
      go mod tidy
      ```

3. **Install `godoc` (optional for documentation):**

    To generate Go documentation locally, you may need to install `godoc`:
    ```bash
    go install golang.org/x/tools/cmd/godoc@latest
    ```

4. **Build the Project:**

    Build the project using the `make` command:
    ```bash
    make build
    ```
    This will compile the Go binary into the `./build/` directory.

---
### Configuration

  **BreezeGate** uses a JSON configuration file to define the routing paths, backend servers, and TLS settings. Below is an example of the `config.json` file:

  ```json
  {
    "port": ":80",
    "healthCheckInterval": "10s",
    "domains": [
      {
        "domainName": "example.com",
        "email": "admin@example.com",
        "useTLS": true,
        "routes": [
          {
            "path": "/api/v1",
            "backends": [
              {"url": "http://localhost:8081", "healthy": true},
              {"url": "http://localhost:8082", "healthy": true}
            ]
          }
        ]
      }
    ]
  }
  ```
### Configuration Fields:

- **port**: The port on which BreezeGate will listen for incoming traffic.
- **healthCheckInterval**: How often to check the health of backend servers.
- **domains**: List of domains BreezeGate will handle. Each domain can have its own email for Let's Encrypt and separate routes.
  - **domainName**: The domain name to be managed.
  - **email**: The admin email for Let's Encrypt registration.
  - **useTLS**: A boolean indicating if TLS should be used.
  - **routes**: Define URL paths and associated backend servers.
    - **path**: The URL path to be routed.
    - **backends**: A list of backend servers for the path.
      - **url**: The URL of the backend server.
      - **healthy**: Initial health status of the backend server (true = healthy).

---

### Usage

Once you've set up the `config.json` file, you can run the load balancer.

1. **Run the Application**:

   To run BreezeGate, execute the following command:

   ```bash
   make run
   ```
    The application will start listening on the configured port, serve requests, and manage backend routing.

2. **Health Checks**:

    BreezeGate performs periodic health checks on backend servers. If a backend server becomes unhealthy, it is temporarily removed from the rotation until it becomes healthy again.

3. **TLS Management**:

    If `useTLS` is set to `true` for a domain, BreezeGate will automatically handle TLS certificates using Let's Encrypt. Certificates are stored locally, and BreezeGate will automatically renew them before they expire.

4. **Monitoring**:

    BreezeGate provides comprehensive health checking and monitoring of backend servers. Failed backends are automatically removed from the rotation until they recover.

---

### Development

#### Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile commands)
- Docker (optional, for containerized deployment)

#### Getting Started

1. **Clone the repository**:
   ```bash
   git clone https://github.com/thetonbr/breezegate.git
   cd breezegate
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Run tests**:
   ```bash
   make test
   # or
   go test ./...
   ```

4. **Run linter**:
   ```bash
   make lint
   # or
   golangci-lint run
   ```

5. **Build the application**:
   ```bash
   make build
   # or
   go build -o build/breezegate ./cmd/app
   ```

#### Code Quality

This project maintains high code quality standards through:

- **Comprehensive Testing**: Unit tests, integration tests, and race condition detection
- **Linting**: Using golangci-lint with strict configuration
- **Code Formatting**: Automatically formatted with `gofmt`
- **Continuous Integration**: Automated testing on multiple Go versions
- **Security Scanning**: Gosec security analysis

### Makefile Commands

BreezeGate comes with a comprehensive Makefile for easy project management:

| Command | Description |
|---------|-------------|
| `make build` | Build the application and generate a binary in the `./build/` directory |
| `make run` | Build and run the application |
| `make clean` | Clean up build artifacts |
| `make test` | Run the tests with coverage reporting |
| `make coverage` | Display test coverage report in terminal |
| `make coverage-html` | Generate HTML coverage report |
| `make lint` | Run golangci-lint to check for code style issues |
| `make fmt` | Automatically format the Go code |
| `make docker-build` | Build a Docker image for the application |
| `make docker-run` | Run the application in a Docker container |
| `make doc` | Start a local Go documentation server at [http://localhost:6060](http://localhost:6060) |
| `make gendoc` | Generate Markdown documentation for the project |

### Docker

BreezeGate supports Docker for easy deployment:

```bash
# Build the Docker image
make docker-build

# Run the container
make docker-run
```

Or use Docker directly:

```bash
# Build
docker build -t breezegate:latest .

# Run
docker run -p 80:80 -p 443:443 --name breezegate breezegate:latest
```

### CI/CD

This project includes comprehensive CI/CD pipelines using GitHub Actions:

#### Continuous Integration (CI)
- **Multi-version testing**: Tests on Go 1.21.x, 1.22.x, and 1.23.x
- **Code quality checks**: Linting, formatting, and security scanning
- **Cross-platform builds**: Linux, macOS, and Windows builds
- **Docker image testing**: Automated container build and test

#### Continuous Deployment (CD)
- **Automated releases**: Tagged releases with pre-built binaries
- **Multi-platform binaries**: Support for multiple architectures
- **Docker Hub integration**: Automated Docker image publishing
- **Changelog generation**: Automated release notes

---

### Future Improvements

- **Database Integration**:
   - Store domain and backend configurations in a database for dynamic updates without requiring a service restart
   - Enable real-time configuration changes via a REST API or web UI
   - Add support for configuration versioning and rollback capabilities

- **WebSockets Support**:
   - Add support for WebSocket connections to handle real-time applications
   - Implement proper WebSocket load balancing with sticky sessions
   - Support for WebSocket health checks and monitoring

- **Security Enhancements**:
   - Implement rate limiting per IP address or user
   - Add IP whitelisting and blacklisting capabilities
   - Integrate with external authentication providers (OAuth, JWT)
   - Add DDoS protection and request filtering

- **Monitoring and Observability**:
   - Integration with Prometheus for metrics collection
   - Add structured logging with configurable log levels
   - Implement distributed tracing support
   - Add alerting for backend failures and performance issues

- **Performance Optimizations**:
   - Implement connection pooling for backend connections
   - Add request/response compression support
   - Implement caching layer for static content
   - Add support for HTTP/2 and HTTP/3

- **Additional Load Balancing Algorithms**:
   - Weighted round-robin
   - Least connections
   - IP hash-based routing
   - Geographic routing based on client location

---

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

### Contributing

We welcome contributions! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

#### Commit Convention

This project uses [Conventional Commits](https://www.conventionalcommits.org/) for automatic semantic versioning. Please ensure your commits follow the format:

```
<type>(<scope>): <subject>
```

Examples:
- `feat(api): add new endpoint for health checks`
- `fix(loadbalancer): resolve memory leak in connection pool`
- `docs(readme): update installation instructions`

Feel free to suggest additional features in the issue tracker or help us improve the code quality by submitting a pull request!

---

### Contact

For questions, issues, or feature requests, please contact me at [thetonbr@gmail.com].


## BreezeGate - Load Balancer with Go

BreezeGate is a load balancer written in Go designed to handle incoming HTTP/HTTPS requests, distribute them across backend servers, and automatically manage TLS certificates via Let's Encrypt (using the ACME protocol). This project uses the `lego` library to handle TLS certificate management and supports reverse proxy functionalities.

### Features

- **Dynamic Backend Management**: Easily configure backend servers and routes through a JSON configuration file.
- **Health Checks**: Periodic health checks to ensure traffic is routed only to healthy backend servers.
- **Automatic TLS Certificates**: Automatically generate and manage SSL certificates using Let's Encrypt.
- **Rate Limiting**: Built-in request rate limiting to prevent abuse.
- **Reverse Proxy**: Forward requests to backend servers seamlessly using Go's built-in `httputil.ReverseProxy`.

### Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [Makefile Commands](#makefile-commands)
- [Future Improvements](#future-improvements)

---

### Installation

1. **Clone the Repository**:
   
   ```bash
   git clone https://github.com/your-username/breezegate.git
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

4. **Rate Limiting**:

    BreezeGate includes a rate-limiting feature to prevent abuse of the service. You can adjust the rate-limiting logic in the `rate_limiting.go` file.

---

### Makefile Commands

BreezeGate comes with a Makefile for easy project management. Here are the available commands:

- **`make build`**: Build the application and generate a binary in the `./build/` directory.
- **`make run`**: Build and run the application.
- **`make clean`**: Clean up build artifacts.
- **`make test`**: Run the tests for the application.
- **`make lint`**: Run the linter to check for code style issues.
- **`make fmt`**: Automatically format the Go code.
- **`make docker-build`**: Build a Docker image for the application.
- **`make docker-run`**: Run the application in a Docker container.
- **`make doc`**: Start a local Go documentation server at [http://localhost:6060](http://localhost:6060).
- **`make gendoc`**: Generate Markdown documentation for the project.

---

### Future Improvements

- **Database Integration**:
   - **TODO**: Store domain and backend configurations in a database for dynamic updates without requiring a service restart. This would allow real-time configuration changes via a UI or API.

- **WebSockets Support**:
   - **TODO**: Add support for WebSockets to allow BreezeGate to handle real-time applications like chats, dashboards, or gaming servers. This would involve modifying the reverse proxy logic to handle the WebSocket protocol.

- **Security Enhancements**:
   - **TODO**: Implement additional security features such as API authentication, an application-level firewall, and protection against DDoS attacks. This could also include rate-limiting per IP, IP whitelisting, and configurable security rules.

---

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

### Contributing

We welcome contributions! Please open an issue or submit a pull request if you have any improvements or bug fixes.

Feel free to suggest additional features in the issue tracker or help us improve the code quality by submitting a pull request!

---

### Contact

For questions, issues, or feature requests, please contact me at [thetonbr@gmail.com].


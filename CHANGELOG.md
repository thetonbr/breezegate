# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive CI/CD pipeline with GitHub Actions
- Security policy and vulnerability reporting process
- Issue and PR templates for better collaboration
- Branch protection rules and governance structure
- Dependabot configuration for automated dependency updates
- Code owners file for review assignments
- Enhanced .gitignore for better repository hygiene
- Contribution guidelines and code of conduct
- Release process documentation

### Changed
- Updated all Go dependencies to latest versions
- Modernized codebase with better error handling
- Improved health check service with context timeouts
- Enhanced documentation with development guides

### Security
- Added HTTP timeouts to prevent slowloris attacks
- Implemented proper context handling in health checks
- Enhanced TLS configuration with secure defaults

### Removed
- Unused internal packages (rate_limiting, file_utils, reverse_proxy)

## [1.0.0] - 2024-01-09

### Added
- Initial release of BreezeGate load balancer
- HTTP/HTTPS load balancing with round-robin algorithm
- Automatic TLS certificate management via Let's Encrypt
- Health checking for backend servers
- JSON-based configuration
- Docker support
- Comprehensive test suite

[Unreleased]: https://github.com/thetonbr/breezegate/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/thetonbr/breezegate/releases/tag/v1.0.0
# Contributing to BreezeGate

First off, thank you for considering contributing to BreezeGate! It's people like you that make BreezeGate such a great tool.

## Code of Conduct

This project and everyone participating in it is governed by the [BreezeGate Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [thetonbr@gmail.com](mailto:thetonbr@gmail.com).

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check [this list](#before-submitting-a-bug-report) as you might find out that you don't need to create one. When you are creating a bug report, please [include as many details as possible](#how-do-i-submit-a-good-bug-report). Fill out [the required template](.github/ISSUE_TEMPLATE/bug_report.md), the information it asks for helps us resolve issues faster.

#### Before Submitting A Bug Report

* **Check the [documentation](README.md)** for a list of common questions and problems.
* **Perform a [cursory search](https://github.com/thetonbr/breezegate/issues)** to see if the problem has already been reported. If it has **and the issue is still open**, add a comment to the existing issue instead of opening a new one.

#### How Do I Submit A (Good) Bug Report?

Bugs are tracked as [GitHub issues](https://github.com/thetonbr/breezegate/issues). Create an issue and provide the following information by filling in [the template](.github/ISSUE_TEMPLATE/bug_report.md).

### Suggesting Enhancements

Enhancement suggestions are tracked as [GitHub issues](https://github.com/thetonbr/breezegate/issues). Create an issue and provide the following information by filling in [the template](.github/ISSUE_TEMPLATE/feature_request.md).

### Your First Code Contribution

Unsure where to begin contributing to BreezeGate? You can start by looking through these `beginner` and `help-wanted` issues:

* [Beginner issues][beginner] - issues which should only require a few lines of code, and a test or two.
* [Help wanted issues][help-wanted] - issues which should be a bit more involved than `beginner` issues.

### Pull Requests

Please follow these steps to have your contribution considered by the maintainers:

1. Follow all instructions in [the template](.github/PULL_REQUEST_TEMPLATE.md)
2. Follow the [styleguides](#styleguides)
3. After you submit your pull request, verify that all [status checks](https://help.github.com/articles/about-status-checks/) are passing

## Development Process

### Branch Strategy

We use a simplified Git Flow strategy:

```
main (stable)
  └── develop (integration)
       ├── feature/your-feature-name
       ├── bugfix/issue-number-description
       ├── hotfix/critical-fix-description
       └── release/v1.2.3
```

#### Branch Types

- **main**: Production-ready code. Protected branch.
- **develop**: Integration branch for features. All feature branches merge here first.
- **feature/**: New features. Branch from `develop`.
- **bugfix/**: Non-critical bug fixes. Branch from `develop`.
- **hotfix/**: Critical fixes. Branch from `main`, merge to both `main` and `develop`.
- **release/**: Release preparation. Branch from `develop`, merge to `main`.

#### Branch Naming Convention

- `feature/short-description` (e.g., `feature/add-prometheus-metrics`)
- `bugfix/issue-number-short-description` (e.g., `bugfix/42-fix-memory-leak`)
- `hotfix/critical-issue-description` (e.g., `hotfix/security-vulnerability`)
- `release/vX.Y.Z` (e.g., `release/v1.2.3`)

### Development Workflow

1. **Fork the repository** and create your branch from `develop`.
2. **Make your changes** and ensure they follow our coding standards.
3. **Write or update tests** as needed.
4. **Run the test suite** to ensure nothing is broken.
5. **Update documentation** if you're changing functionality.
6. **Submit a pull request** to the `develop` branch.

### Setup Development Environment

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/breezegate.git
cd breezegate

# Add upstream remote
git remote add upstream https://github.com/thetonbr/breezegate.git

# Install dependencies
go mod download

# Install development tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run tests
make test

# Run linter
make lint
```

## Styleguides

### Git Commit Messages

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification for automatic semantic versioning:

```
<type>(<scope>): <subject>

<body>

<footer>
```

#### Types
- `feat`: New feature (triggers minor release)
- `fix`: Bug fix (triggers patch release)
- `docs`: Documentation changes (no release unless scope is README)
- `style`: Code style changes (no release)
- `refactor`: Code refactoring (no release)
- `perf`: Performance improvements (triggers patch release)
- `test`: Test changes (no release)
- `build`: Build system changes (no release)
- `ci`: CI configuration changes (no release)
- `chore`: Other changes (no release)
- `revert`: Reverts a previous commit (triggers patch release)

#### Breaking Changes
- Add `BREAKING CHANGE:` in the commit body or footer
- Or add `!` after the type: `feat!: breaking change`
- This triggers a major version release

#### Examples

```
feat(loadbalancer): add weighted round-robin algorithm

Implement weighted round-robin load balancing to allow
different weights for backend servers based on their capacity.

Closes #45
```

```
fix(healthcheck): prevent goroutine leak on shutdown

Add proper context cancellation to health check goroutines
to ensure they terminate when the application shuts down.

Fixes #78
```

```
feat!: change config file format to YAML

BREAKING CHANGE: Configuration files must now be in YAML format.
JSON configuration files are no longer supported.

Migration guide: https://docs.breezegate.dev/migration/v2
```

### Go Styleguide

* Follow the [Effective Go](https://golang.org/doc/effective_go.html) guidelines
* Use `gofmt` to format your code
* Run `golangci-lint` before submitting
* Write meaningful variable and function names
* Add comments for exported functions and types
* Keep functions small and focused
* Handle errors explicitly
* Write tests for new functionality

### Testing

* Write unit tests for all new functionality
* Ensure tests are deterministic and don't depend on external services
* Use table-driven tests where appropriate
* Test edge cases and error conditions
* Aim for at least 80% code coverage for new code

### Documentation

* Update the README.md if you change functionality
* Document all exported functions and types
* Include examples in documentation where helpful
* Keep documentation concise and clear

## Security

### Reporting Security Vulnerabilities

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, please report them to the project maintainers at [thetonbr@gmail.com](mailto:thetonbr@gmail.com). You should receive a response within 48 hours. If for some reason you do not, please follow up via email to ensure we received your original message.

Please include the following information:

* Type of issue (e.g., buffer overflow, SQL injection, cross-site scripting, etc.)
* Full paths of source file(s) related to the manifestation of the issue
* The location of the affected source code (tag/branch/commit or direct URL)
* Any special configuration required to reproduce the issue
* Step-by-step instructions to reproduce the issue
* Proof-of-concept or exploit code (if possible)
* Impact of the issue, including how an attacker might exploit the issue

### Security Best Practices for Contributors

* Never commit secrets, API keys, or credentials
* Always validate and sanitize input
* Use prepared statements for database queries
* Implement proper error handling without exposing sensitive information
* Follow the principle of least privilege
* Keep dependencies up to date
* Use secure communication protocols (HTTPS, TLS)

## Code Review Process

The core team looks at Pull Requests on a regular basis. After feedback has been given, we expect responses within two weeks. After two weeks, we may close the pull request if it isn't showing any activity.

### Review Criteria

* **Code Quality**: Is the code clean, readable, and maintainable?
* **Testing**: Are there adequate tests? Do they cover edge cases?
* **Documentation**: Is the code well-documented? Is user documentation updated?
* **Performance**: Does the change impact performance? Are there benchmarks?
* **Security**: Does the code follow security best practices?
* **Compatibility**: Does the change break backward compatibility?

## Community

* Join our [Discussions](https://github.com/thetonbr/breezegate/discussions) for general questions
* Follow our [Blog](https://breezegate.dev/blog) for updates (when available)
* Chat with us on [Discord](https://discord.gg/breezegate) (when available)

## Recognition

Contributors who have made significant contributions will be recognized in our [CONTRIBUTORS.md](CONTRIBUTORS.md) file.

## Questions?

Don't hesitate to ask questions if something is unclear. You can:

1. Open a [Discussion](https://github.com/thetonbr/breezegate/discussions)
2. Contact the maintainers at [thetonbr@gmail.com](mailto:thetonbr@gmail.com)
3. Ask in a comment on your issue or pull request

Thank you for contributing to BreezeGate! 🎉

[beginner]:https://github.com/thetonbr/breezegate/labels/beginner
[help-wanted]:https://github.com/thetonbr/breezegate/labels/help%20wanted
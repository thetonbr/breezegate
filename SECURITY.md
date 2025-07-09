# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Which versions are eligible for receiving such patches depends on the CVSS v3.0 Rating:

| Version | Supported          |
| ------- | ------------------ |
| 1.x.x   | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

We take the security of BreezeGate seriously. If you believe you have found a security vulnerability, please report it to us as described below.

### Please do NOT:
- Open a public GitHub issue for security vulnerabilities
- Post about it on social media
- Attempt to exploit the vulnerability on production systems

### Please DO:
- Email us directly at: thetonbr@gmail.com
- Encrypt sensitive information using our PGP key (if available)
- Allow us reasonable time to respond before public disclosure

### What to include in your report:

1. **Description**: Clear description of the vulnerability
2. **Impact**: What can an attacker achieve?
3. **Steps to reproduce**: Detailed steps to reproduce the issue
4. **Affected versions**: Which versions are vulnerable?
5. **Mitigation**: Any temporary workarounds?
6. **References**: Links to relevant resources

### What to expect:

1. **Acknowledgment**: We'll acknowledge receipt within 48 hours
2. **Initial Assessment**: Within 7 days, we'll provide an initial assessment
3. **Resolution Timeline**: We'll work with you to understand and resolve the issue
4. **Disclosure**: Once fixed, we'll work on responsible disclosure

## Security Best Practices for Users

### Configuration Security

1. **TLS Configuration**:
   - Always use TLS in production
   - Keep certificates up to date
   - Use strong cipher suites

2. **Backend Security**:
   - Use HTTPS for backend connections when possible
   - Implement proper authentication between load balancer and backends
   - Regularly update backend servers

3. **Access Control**:
   - Restrict access to configuration files
   - Use proper file permissions (600 for sensitive files)
   - Don't expose the configuration endpoint publicly

### Operational Security

1. **Monitoring**:
   - Monitor logs for suspicious activity
   - Set up alerts for failed health checks
   - Track certificate expiration dates

2. **Updates**:
   - Keep BreezeGate updated to the latest version
   - Subscribe to security announcements
   - Test updates in staging before production

3. **Network Security**:
   - Use firewalls to restrict access
   - Implement rate limiting where appropriate
   - Use private networks for backend communication

## Security Features

BreezeGate includes several security features:

1. **Automatic TLS**: Let's Encrypt integration for automatic certificate management
2. **Health Checks**: Automatic detection and removal of unhealthy backends
3. **Secure Defaults**: Secure configuration defaults out of the box
4. **No Sensitive Data Logging**: Careful to not log sensitive information

## Security Acknowledgments

We would like to thank the following individuals for responsibly disclosing security issues:

* None yet - be the first!

## Contact

For any security-related questions or concerns, please contact:
- Email: thetonbr@gmail.com
- PGP Key: [Coming soon]

Thank you for helping keep BreezeGate and its users safe!
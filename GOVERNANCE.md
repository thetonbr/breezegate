# BreezeGate Governance

This document defines the governance structure for the BreezeGate project.

## Project Structure

### Roles

#### Maintainer
- **Current**: @thetonbr
- **Responsibilities**:
  - Project direction and roadmap
  - Release management
  - Security response
  - Final decision on controversial changes
  - Community management

#### Core Contributors
- Regular contributors with commit access
- Can review and merge PRs
- Help with issue triage
- Participate in roadmap discussions

#### Contributors
- Anyone who contributes to the project
- Through code, documentation, testing, or other means

### Decision Making

1. **Consensus-based**: Most decisions are made through consensus in GitHub issues and PRs
2. **Maintainer decision**: For controversial or time-sensitive issues, the maintainer makes the final decision
3. **Community input**: Major changes are discussed in GitHub Discussions before implementation

### Becoming a Core Contributors

To become a Core Contributors, you should:

1. Have made significant contributions to the project
2. Show commitment to the project's success
3. Demonstrate understanding of the codebase
4. Be responsive to issues and PRs
5. Follow the project's code of conduct

Nominations can be made by existing Core Contributors or the Maintainer.

## Release Process

### Version Numbering

We follow [Semantic Versioning](https://semver.org/):
- MAJOR version for incompatible API changes
- MINOR version for backwards-compatible functionality additions
- PATCH version for backwards-compatible bug fixes

### Release Cycle

- **Patch releases**: As needed for critical bugs
- **Minor releases**: Approximately every 2-3 months
- **Major releases**: As needed, with significant planning

### Release Process

1. Create release branch from `develop`
2. Update version numbers and changelog
3. Create release candidate (RC) for testing
4. After testing period, merge to `main`
5. Tag release and create GitHub release
6. Merge back to `develop`

## Contribution Process

1. **Issues**: All work should start with an issue
2. **Discussion**: Significant changes should be discussed before implementation
3. **Development**: Work happens in feature branches
4. **Review**: All changes require code review
5. **Testing**: Changes must include tests
6. **Documentation**: Update docs as needed
7. **Merge**: PRs are merged to `develop` first

## Communication Channels

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: General discussions and questions
- **Email**: thetonbr@gmail.com for private matters

## Code of Conduct

All participants must follow our Code of Conduct. Violations should be reported to thetonbr@gmail.com.

## Changes to Governance

This governance model can be changed through the normal PR process, with approval from the Maintainer and input from Core Contributors.
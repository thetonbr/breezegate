# Release Process

This document describes the release process for BreezeGate.

## Release Schedule

- **Patch releases**: As needed for critical fixes
- **Minor releases**: Every 2-3 months
- **Major releases**: As needed with advance notice

## Release Workflow

### 1. Preparation (1 week before release)

```bash
# Create release branch from develop
git checkout develop
git pull origin develop
git checkout -b release/vX.Y.Z

# Update version in relevant files
# - Update version in code if needed
# - Update CHANGELOG.md
# - Update README.md if needed
```

### 2. Testing Phase (3-5 days)

- Run full test suite
- Deploy to staging environment
- Perform manual testing
- Security scan
- Performance testing

### 3. Release Candidate

```bash
# Tag release candidate
git tag -a vX.Y.Z-rc1 -m "Release candidate 1 for vX.Y.Z"
git push origin vX.Y.Z-rc1
```

### 4. Final Release

```bash
# Merge to main
git checkout main
git merge --no-ff release/vX.Y.Z

# Tag final release
git tag -a vX.Y.Z -m "Release vX.Y.Z"
git push origin main
git push origin vX.Y.Z

# Merge back to develop
git checkout develop
git merge --no-ff main
git push origin develop
```

### 5. GitHub Release

1. Go to GitHub Releases
2. Create new release from tag
3. Add release notes from CHANGELOG
4. Upload binary artifacts
5. Publish release

### 6. Post-Release

- Announce release in Discussions
- Update documentation site
- Monitor for issues
- Plan next release

## Version Numbering

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR.MINOR.PATCH**
- **MAJOR**: Breaking changes
- **MINOR**: New features, backwards compatible
- **PATCH**: Bug fixes, backwards compatible

## Release Checklist

- [ ] All tests passing
- [ ] Security scan clean
- [ ] Performance benchmarks acceptable
- [ ] CHANGELOG.md updated
- [ ] Documentation updated
- [ ] Version numbers updated
- [ ] Release notes written
- [ ] Binaries built for all platforms
- [ ] Docker images built and tested
- [ ] Release branch created
- [ ] RC testing completed
- [ ] Final merge to main
- [ ] Tag created
- [ ] GitHub release published
- [ ] Merged back to develop
- [ ] Release announced

## Emergency Releases

For critical security fixes:

1. Create hotfix branch from main
2. Apply fix with tests
3. Fast-track testing
4. Release as patch version
5. Merge to both main and develop

## Release Notes Template

```markdown
## vX.Y.Z - YYYY-MM-DD

### Breaking Changes
- List any breaking changes

### New Features
- New feature 1
- New feature 2

### Improvements
- Improvement 1
- Improvement 2

### Bug Fixes
- Fixed issue #123: Description
- Fixed issue #456: Description

### Security
- Security fix description (with CVE if applicable)

### Dependencies
- Updated dependency X to vY.Z

### Contributors
- @username1
- @username2

**Full Changelog**: https://github.com/thetonbr/breezegate/compare/vX.Y.Y...vX.Y.Z
```
# Branching Strategy

BreezeGate follows a simplified Git Flow branching model to ensure code quality and stable releases.

## Branch Types

### Protected Branches

#### `main`
- **Purpose**: Production-ready code
- **Protection**: 
  - Requires PR with approval
  - Must pass all CI checks
  - No direct commits
  - No force pushes
  - Requires up-to-date branch

#### `develop`
- **Purpose**: Integration branch for next release
- **Protection**:
  - Requires PR with approval
  - Must pass CI checks
  - No force pushes

### Working Branches

#### `feature/*`
- **Purpose**: New features
- **Naming**: `feature/short-description`
- **Example**: `feature/add-prometheus-metrics`
- **Flow**: 
  ```
  develop → feature/* → develop
  ```

#### `bugfix/*`
- **Purpose**: Non-critical bug fixes
- **Naming**: `bugfix/issue-number-description`
- **Example**: `bugfix/42-fix-memory-leak`
- **Flow**: 
  ```
  develop → bugfix/* → develop
  ```

#### `hotfix/*`
- **Purpose**: Critical production fixes
- **Naming**: `hotfix/description`
- **Example**: `hotfix/security-vulnerability`
- **Flow**: 
  ```
  main → hotfix/* → main + develop
  ```

#### `release/*`
- **Purpose**: Release preparation
- **Naming**: `release/vX.Y.Z`
- **Example**: `release/v1.2.0`
- **Flow**: 
  ```
  develop → release/* → main + develop
  ```

## Workflow Examples

### Feature Development

```bash
# Start from develop
git checkout develop
git pull origin develop

# Create feature branch
git checkout -b feature/my-awesome-feature

# Work on feature
# ... make changes ...
git add .
git commit -m "Add awesome feature"

# Keep up to date with develop
git checkout develop
git pull origin develop
git checkout feature/my-awesome-feature
git rebase develop

# Push and create PR
git push origin feature/my-awesome-feature
# Create PR to develop branch
```

### Bug Fix

```bash
# Start from develop
git checkout develop
git pull origin develop

# Create bugfix branch
git checkout -b bugfix/123-fix-connection-timeout

# Fix bug
# ... make changes ...
git add .
git commit -m "Fix connection timeout issue

Fixes #123"

# Push and create PR
git push origin bugfix/123-fix-connection-timeout
# Create PR to develop branch
```

### Hotfix

```bash
# Start from main for critical fixes
git checkout main
git pull origin main

# Create hotfix branch
git checkout -b hotfix/critical-security-fix

# Apply fix
# ... make changes ...
git add .
git commit -m "Fix critical security vulnerability"

# Push and create PR
git push origin hotfix/critical-security-fix
# Create PR to main branch

# After merge to main, also merge to develop
git checkout develop
git pull origin develop
git merge main
git push origin develop
```

### Release

```bash
# Start release from develop
git checkout develop
git pull origin develop

# Create release branch
git checkout -b release/v1.2.0

# Prepare release
# - Update version numbers
# - Update CHANGELOG.md
# - Final testing
git add .
git commit -m "Prepare release v1.2.0"

# Push and create PR
git push origin release/v1.2.0
# Create PR to main branch

# After merge to main
git checkout main
git pull origin main
git tag -a v1.2.0 -m "Release v1.2.0"
git push origin v1.2.0

# Merge back to develop
git checkout develop
git merge main
git push origin develop
```

## Commit Message Guidelines

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc)
- `refactor`: Code refactoring
- `perf`: Performance improvements
- `test`: Test additions or modifications
- `build`: Build system changes
- `ci`: CI configuration changes
- `chore`: Other changes that don't affect src or test files

### Examples

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

## Pull Request Guidelines

1. **Base Branch**: Choose the correct base branch
   - Features/Bugs → `develop`
   - Hotfixes → `main`
   - Releases → `main`

2. **Title**: Clear and descriptive
   - Good: "Add Prometheus metrics endpoint"
   - Bad: "Update code"

3. **Description**: Use the PR template

4. **Size**: Keep PRs small and focused

5. **Tests**: Include tests for new code

6. **Reviews**: Address reviewer feedback promptly

## Branch Cleanup

- Delete feature/bugfix branches after merge
- Keep release branches for history
- Never delete main or develop

## Conflict Resolution

1. Always rebase feature branches on develop
2. Resolve conflicts locally
3. Test after resolving conflicts
4. Force push to your branch only (never to protected branches)

## FAQ

**Q: When should I use a hotfix vs bugfix branch?**
A: Use hotfix for critical production issues that can't wait for the next release. Use bugfix for non-critical issues that can be included in the next regular release.

**Q: Can I commit directly to develop?**
A: No, all changes must go through a PR, even for maintainers.

**Q: How often should I rebase my feature branch?**
A: Rebase whenever develop has significant changes, and always before creating a PR.

**Q: What if my feature takes a long time to develop?**
A: Consider breaking it into smaller features, or regularly rebase from develop to avoid conflicts.
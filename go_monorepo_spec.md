# Azure Dev Monorepo Specification (Go Modules)

## Overview
This document defines the structure, module layout, versioning, and workflow for the **Azure Dev Monorepo** project written in Go. The repo follows a multi-module architecture using `go.work` to support cohesive local development while maintaining isolated, versioned module releases.

---

## Folder Structure

```
/
├─ go.work                         # Workspace for local + CI integration builds
│
├─ core/
│  ├─ cli/                         # CLI module (consumes sdk + internal)
│  │  └─ go.mod    -> module github.com/azure/azure-dev/core/cli
│  │
│  ├─ sdk/                         # Public SDK module (consumed by CLI + extensions)
│  │  └─ go.mod    -> module github.com/azure/azure-dev/core/sdk
│  │
│  └─ internal/                    # Shared-private helpers (CLI + SDK only)
│     └─ go.mod   -> module github.com/azure/azure-dev/core/internal
│
└─ extensions/
   ├─ extension1/
   │  └─ go.mod   -> module github.com/azure/azure-dev/extensions/extension1
   └─ extension2/
      └─ go.mod   -> module github.com/azure/azure-dev/extensions/extension2
```

### Visibility Rules
| Module | Importable by | Description |
|---------|----------------|--------------|
| `core/internal` | CLI + SDK | Shared-private utilities, hidden from extensions |
| `core/sdk` | CLI + Extensions | Public SDK surface for extensions |
| `core/cli` | CLI only | CLI binary + internal logic |
| `extensions/*` | End users | External extension modules |

---

## Workspace Configuration

### go.work

Committed at the root of the repo:
```txt
go 1.25

use (
  ./core/cli
  ./core/sdk
  ./core/internal
  ./extensions/extension1
  ./extensions/extension2
)
```
- Used for local development and CI integration builds.
- Excludes any `replace` directives — workspace replaces are implicit.
- For release validation and pinned version tests, use `GOWORK=off`.

---

## Module Definitions

### core/sdk/go.mod
```go
module github.com/azure/azure-dev/core/sdk

go 1.25

require (
    github.com/azure/azure-dev/core/internal v1.0.0
)
```

### core/cli/go.mod
```go
module github.com/azure/azure-dev/core/cli

go 1.25

require (
    github.com/azure/azure-dev/core/sdk v1.6.0
    github.com/azure/azure-dev/core/internal v1.1.0
)
```

### extensions/extension1/go.mod
```go
module github.com/azure/azure-dev/extensions/extension1

go 1.25

require (
    github.com/azure/azure-dev/core/sdk v1.6.0
)
```

---

## Versioning & Tagging

Each module uses subdirectory-prefixed Git tags to support independent semantic versioning.

| Module | Tag format | Example |
|---------|-------------|----------|
| `core/cli` | `core/cli/vX.Y.Z` | `core/cli/v1.27.0` |
| `core/sdk` | `core/sdk/vX.Y.Z` | `core/sdk/v1.6.0` |
| `core/internal` | `core/internal/vX.Y.Z` | `core/internal/v1.1.0` |
| `extensions/extension1` | `extensions/extension1/vA.B.C` | `extensions/extension1/v0.9.0` |
| `extensions/extension2` | `extensions/extension2/vA.B.C` | `extensions/extension2/v0.5.2` |

**Tagging Commands**
```bash
git tag core/sdk/v1.6.0
git push origin core/sdk/v1.6.0

git tag core/cli/v1.27.0
git push origin core/cli/v1.27.0
```

### Major Version Changes (SIV)
When a module introduces breaking changes:
- Update the module path and tag:
  - `core/sdk/go.mod` → `module github.com/azure/azure-dev/core/sdk/v2`
  - Tag → `core/sdk/v2.0.0`
  - Consumers import → `github.com/azure/azure-dev/core/sdk/v2/...`

---

## CI/CD Workflow

### Integration Build (Workspace Mode)
For pull requests and integration tests:
```bash
go work sync
go test ./...
```
Runs all modules together using workspace-local replacements.

### Module Release Validation (Pinned Mode)
For release pipelines:
```bash
cd core/sdk
GOWORK=off go test -mod=readonly ./...
```
Ensures the module builds and tests successfully against pinned versions only.

### Environment Variables
| Var | Purpose |
|-----|----------|
| `GOWORK=off` | Disables workspace usage for pinned builds |
| `-mod=readonly` | Prevents accidental go.mod edits in CI |

---

## Release Flow

### Example: New SDK Minor + CLI Update
```bash
# After merging feature PR

git tag core/sdk/v1.6.0
git push origin core/sdk/v1.6.0

# Bump CLI to use the new SDK
# core/cli/go.mod → require github.com/azure/azure-dev/core/sdk v1.6.0

git tag core/cli/v1.27.0
git push origin core/cli/v1.27.0
```

### Example: Breaking SDK Release
```bash
# Feature PR: change module path to /v2 and fix imports

git tag core/sdk/v2.0.0
git push origin core/sdk/v2.0.0

# CLI update to consume v2 SDK
git tag core/cli/v2.0.0
git push origin core/cli/v2.0.0
```

---

## Best Practices
- ✅ Commit `go.work` for shared local + CI consistency.
- ✅ Use subdirectory-prefixed tags for all modules.
- ✅ Always tag **upstream before downstream** (`internal` → `sdk` → `cli` → `extensions`).
- ✅ Use `GOWORK=off` and `-mod=readonly` in release validation.
- 🚫 Never re-tag existing versions (immutable history).
- 🚫 Avoid `replace` directives in committed `go.mod` files.

---

## Summary
This layout ensures:
- Clean separation of internal, SDK, and CLI logic.
- Deterministic versioning per module.
- Seamless local development via `go.work`.
- Strict visibility and import boundaries enforced by Go’s `internal` rules.

It aligns with Go’s recommended multi-module monorepo design and supports modular CI/CD pipelines for each logical component of the Azure Dev ecosystem.


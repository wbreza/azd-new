# Azure Dev Monorepo Implementation

This repository implements the Go monorepo specification defined in `go_monorepo_spec.md`.

## Project Structure

```
/
├─ go.work                         # Workspace configuration
├─ core/
│  ├─ cli/                        # Main Azure Dev CLI (Cobra app)
│  │  ├─ main.go                  # CLI entry point
│  │  ├─ cmd/                     # Cobra command definitions
│  │  └─ internal/                # CLI-specific internal packages
│  ├─ sdk/                        # Public SDK for extensions
│  │  ├─ client.go                # SDK client implementation
│  │  └─ go.mod
│  └─ internal/                   # Shared internal utilities
│     ├─ utils.go                 # Common utilities
│     └─ go.mod
└─ extensions/
   ├─ extension1/                 # Sample extension (Cobra app)
   │  ├─ main.go                  # Extension entry point
   │  ├─ cmd/                     # Extension-specific commands
   │  └─ internal/                # Extension-specific internals
   └─ extension2/                 # Sample monitoring extension
      ├─ main.go
      ├─ cmd/
      └─ internal/

```

## Module Dependencies

- `core/internal`: Base utilities (no dependencies)
- `core/sdk`: Public SDK (depends on core/internal)
- `core/cli`: Main CLI (depends on core/sdk, core/internal, cobra)
- `extensions/*`: Extensions (depend on core/sdk, cobra)

## CLI Applications

### Core CLI (`azd`)
- **Commands**: `deploy`, `init`, `provision`
- **Usage**: `go run ./core/cli deploy myapp`

### Extension 1 (`azd-ext1`)
- **Commands**: `custom`, `integrate`
- **Usage**: `go run ./extensions/extension1 custom myresource`

### Extension 2 (`azd-ext2`)
- **Commands**: `monitor`, `analyze`
- **Usage**: `go run ./extensions/extension2 monitor myapp`

## Building and Running

### Development (Workspace Mode)
```bash
# Run the main CLI
go run ./core/cli deploy myapp

# Run extension1
go run ./extensions/extension1 custom myresource

# Run extension2
go run ./extensions/extension2 monitor myapp

# Test all modules
go test ./...
```

### Production Builds
```bash
# Build main CLI
cd core/cli && go build -o azd.exe .

# Build extensions
cd extensions/extension1 && go build -o azd-ext1.exe .
cd extensions/extension2 && go build -o azd-ext2.exe .
```

## Versioning & Releases

This repository follows the semantic versioning scheme defined in the specification:

- `core/cli/vX.Y.Z` - Main CLI releases
- `core/sdk/vX.Y.Z` - SDK releases  
- `core/internal/vX.Y.Z` - Internal utilities releases
- `extensions/extension1/vX.Y.Z` - Extension1 releases
- `extensions/extension2/vX.Y.Z` - Extension2 releases

### Example Release Flow
```bash
# Tag and release SDK first
git tag core/sdk/v1.6.0
git push origin core/sdk/v1.6.0

# Then update and release CLI
git tag core/cli/v1.27.0
git push origin core/cli/v1.27.0
```

## Notes

- All CLI applications use the Cobra framework
- The workspace handles local module resolution during development
- Extensions can only import from the public SDK, not internal utilities
- Each CLI app follows the pattern: `main.go` → `cmd/root.go` → `internal/commands/`
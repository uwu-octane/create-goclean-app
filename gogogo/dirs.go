package main

// baseDirs returns the static folder layout taken from github.com/evrone/go-clean-template
// (only top-level paths; each will contain a .gitkeep so the directory is versioned).
func baseDirs() []string {
	return []string{
		"cmd/app", "cmd/grpc", "cmd/worker",
		"internal/config",
		"internal/app",
		"internal/controller/http/v1",
		"internal/controller/grpc",
		"internal/entity",
		"internal/usecase",
		"internal/repo",
		"internal/repo/redis",
		"internal/mq",
		"migrations",
		"docs/openapi", "docs/proto",
		"pkg/logger",
	}
}

package version

var (
	// Version is set via -ldflags at build time.
	Version = "dev"
	// Commit is the git commit hash set at build time.
	Commit = ""
)
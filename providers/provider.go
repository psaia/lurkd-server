package providers

// ReleaseLabel based on semver.
type ReleaseLabel int64

// Each ReleaseLabel interval.
const (
	MAJOR ReleaseLabel = iota
	MINOR
	PATCH
)

// VersionResult for GetVersion()
type VersionResult struct {
	Version string
	Change  ReleaseLabel
}

// Provider for an extending service.
type Provider interface {
	// GetVersion will get the version of the module in context. This is used by
	// the other methods.
	GetVersion() (string, error)

	// Initialize will be called upon initially setting up the provider.
	// It will just the version that gets stored and any error.
	Initialize() (string, error)

	// Check the version and
	Check(cur ReleaseLabel) (VersionResult, error)
}

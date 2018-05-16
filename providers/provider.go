package providers

type Settings struct {
}

// Interface for an extending provider.
type Interface interface {
	Initialize()
	GetVersion()
}

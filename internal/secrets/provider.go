package secrets

// Todo

type Provider interface {
	GetSecret(secretPath string) map[string]any
}

type Secrets struct {
	Provider Provider
}

func New(provider Provider) *Secrets {
	return &Secrets{Provider: provider}
}

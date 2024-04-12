package secrets

// Todo

type Provider interface {
	GetSecret(v string) (any, error)
}

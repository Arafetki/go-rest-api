package secrets

import "github.com/Arafetki/my-portfolio-api/internal/vault"

type Store struct {
	Provider interface {
		ReadString(secretPath, key string) (string, error)
	}
}

func NewStore(v *vault.Vault) *Store {
	return &Store{Provider: v}
}

package secrets

import (
	"context"
	"fmt"
	"os"
	"time"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
)

type Vault struct {
	path   string
	client *vault.Logical
}

func NewVault(path string) (*Vault, error) {
	config := &vault.Config{
		Address: os.Getenv("VAULT_ADDR"),
		Timeout: 10 * time.Second,
	}
	client, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}
	appRoleAuth, err := auth.NewAppRoleAuth(
		os.Getenv("VAULT_APPROLE_ROLE_ID"),
		&auth.SecretID{FromString: os.Getenv("VAULT_APPROLE_SECRET_ID")},
		auth.WithMountPath(os.Getenv("VAULT_APPROLE_PATH")),
	)
	if err != nil {
		return nil, err
	}
	_, err = client.Auth().Login(context.Background(), appRoleAuth)
	if err != nil {
		return nil, err
	}

	return &Vault{path: path, client: client.Logical()}, nil
}

func (v *Vault) GetSecret(secretPath string) map[string]string {
	secret, err := v.client.Read(fmt.Sprintf("%s/data/%s", v.path, secretPath))
	if err != nil {
		panic(err)
	}
	return secret.Data["data"].(map[string]string)
}

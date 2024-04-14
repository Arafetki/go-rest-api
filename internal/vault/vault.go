package vault

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
)

type Vault struct {
	Path   string
	Client *api.Logical
}

func NewVault(path string) (*Vault, error) {
	config := &api.Config{
		Address: os.Getenv("VAULT_ADDR"),
		Timeout: 10 * time.Second,
	}
	client, err := api.NewClient(config)
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

	return &Vault{Path: path, Client: client.Logical()}, nil
}

func (v Vault) ReadString(secretPath string, key string) (string, error) {
	secret, err := v.Client.Read(fmt.Sprintf("%s/data/%s", v.Path, secretPath))
	if err != nil {
		return "", err
	}
	secretValue, exist := secret.Data["data"].(map[string]any)[key]
	if !exist {
		return "", fmt.Errorf("secret %s not found", key)
	}
	return secretValue.(string), nil
}

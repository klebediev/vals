package stringmapprovider

import (
	"fmt"
	"github.com/klebediev/vals/pkg/api"
	"github.com/klebediev/vals/pkg/providers/awssecrets"
	"github.com/klebediev/vals/pkg/providers/sops"
	"github.com/klebediev/vals/pkg/providers/ssm"
	"github.com/klebediev/vals/pkg/providers/vault"
)

func New(provider api.StaticConfig) (api.LazyLoadedStringMapProvider, error) {
	tpe := provider.String("name")

	switch tpe {
	case "ssm":
		return ssm.New(provider), nil
	case "vault":
		return vault.New(provider), nil
	case "awssecrets":
		return awssecrets.New(provider), nil
	case "sops":
		return sops.New(provider), nil
	}

	return nil, fmt.Errorf("failed initializing string-map provider from config: %v", provider)
}

package cloudprovider

import (
	"errors"

	"github.com/kubermatic/machine-controller/pkg/cloudprovider/instance"
	"github.com/kubermatic/machine-controller/pkg/cloudprovider/provider/digitalocean"
	"github.com/kubermatic/machine-controller/pkg/machines/v1alpha1"
)

var (
	ErrProviderNotFound = errors.New("cloudprovider not found")

	providers = map[string]CloudProvider{
		"digitalocean": digitalocean.New(),
	}
)

func ForProvider(p string) (CloudProvider, error) {
	if p, found := providers[p]; found {
		return p, nil
	}
	return nil, ErrProviderNotFound
}

type CloudProvider interface {
	Validate(machinespec v1alpha1.MachineSpec) error
	Get(machine *v1alpha1.Machine) (instance.Instance, error)
	Create(machine *v1alpha1.Machine, userdata string, authorizedkey []byte) (instance.Instance, error)
	Delete(machine *v1alpha1.Machine) error
}
package backend

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.com/gridio/test-assignment/internal"
	"gitlab.com/gridio/test-assignment/pkg/chargeamps/identity"
)

type Identity interface {
	// AccessToken provides access token string to be used for device list and status requests
	AccessToken() string

	// IsUnauthorized should return true if the token is expired or invalid
	IsUnauthorized() bool
}

type Backend struct {
	tokenAgent internal.SecretAgent
	userID     string
	id         Identity

	// Whatever fields you need
}

// 1. Write integration factory function that produces integrations with the following signature

func Factory(log logrus.FieldLogger) func(string, internal.SecretAgent) internal.DeviceListProvider {
	f := func(userID string, sa internal.SecretAgent) internal.DeviceListProvider {
		bck := Backend{
			tokenAgent: sa,
			userID:     userID,
			id:         identity.CreateFromSecretAgent(log.WithField("user_id", userID), sa),
		}

		return &bck
	}

	return f
}

// 2. Implement interface internal.ChargerBackend

func (b *Backend) DoDeviceListRequest(ctx context.Context) ([]internal.DeviceMetadata, error) {
	// TODO implement me
	panic("implement me")
}

func (b *Backend) IsUnauthorized() bool {
	// TODO implement me
	panic("implement me")
}

func (b *Backend) DoChargerStatusRequest(ctx context.Context, id internal.PhysicalID) (*internal.ChargerStatus, error) {
	// TODO implement me
	panic("implement me")
}

func (b *Backend) StartCharge(ctx context.Context, id internal.PhysicalID, p internal.Power) error {
	// TODO implement me
	panic("implement me")
}

func (b *Backend) Stop(ctx context.Context, id internal.PhysicalID) error {
	// TODO implement me
	panic("implement me")
}

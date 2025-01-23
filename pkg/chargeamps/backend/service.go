package backend

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	"gitlab.com/gridio/test-assignment/internal"
	httphelper "gitlab.com/gridio/test-assignment/pkg/chargeamps/http"
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

func Factory(log logrus.FieldLogger) func(string, internal.SecretAgent) (internal.DeviceListProvider, error) {
	f := func(userID string, sa internal.SecretAgent) (internal.DeviceListProvider, error) {
		id, err := identity.CreateFromSecretAgent(log.WithField("user_id", userID), sa)
		if err != nil {
			return nil, err
		}

		bck := Backend{
			tokenAgent: sa,
			userID:     userID,
			id:         id,
		}

		return &bck, nil
	}

	return f
}

// 2. Implement interface internal.ChargerBackend

func (b *Backend) DoDeviceListRequest(ctx context.Context) ([]internal.DeviceMetadata, error) {
	const url = "https://eapi.charge.space/api/v5/chargepoints/owned"
	token := b.tokenAgent.ProvideSecret()

	req, err := httphelper.CreateRequest("GET", url, "Authorization", fmt.Sprintf("Bearer %s", token[1:len(token)-1]), nil)
	if err != nil {
		return nil, err
	}

	body, err := httphelper.GetResponseBody(req)
	if err != nil {
		return nil, err
	}

	var deviceMetadata []internal.DeviceMetadata
	err = json.Unmarshal(body, &deviceMetadata)
	if err != nil {
		return nil, err
	}

	return deviceMetadata, nil
}

func (b *Backend) IsUnauthorized() bool {
	// TODO implement me
	panic("implement me")
}

func (b *Backend) DoChargerStatusRequest(ctx context.Context, id internal.PhysicalID) (*internal.ChargePointStatus, error) {
	url := fmt.Sprintf("https://eapi.charge.space/api/v5/chargepoints/%s/status", id)
	token := b.tokenAgent.ProvideSecret()

	req, err := httphelper.CreateRequest("GET", url, "Authorization", fmt.Sprintf("Bearer %s", token[1:len(token)-1]), nil)
	if err != nil {
		return nil, err
	}

	body, err := httphelper.GetResponseBody(req)
	if err != nil {
		return nil, err
	}

	var chargerStatus internal.ChargePointStatus
	err = json.Unmarshal(body, &chargerStatus)
	if err != nil {
		return nil, err
	}

	return &chargerStatus, nil
}

func (b *Backend) StartCharge(ctx context.Context, id internal.PhysicalID, p internal.Power) error {
	// TODO implement me
	panic("implement me")
}

func (b *Backend) Stop(ctx context.Context, id internal.PhysicalID) error {
	// TODO implement me
	panic("implement me")
}

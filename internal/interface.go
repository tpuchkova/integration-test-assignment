package internal

import (
	"context"
	"time"
)

type DeviceListProvider interface {
	DoDeviceListRequest(ctx context.Context) ([]DeviceMetadata, error)
	IsUnauthorized() bool
}

type ChargerBackend interface {
	DeviceListProvider
	DoChargerStatusRequest(ctx context.Context, id PhysicalID) (*ChargePointStatus, error)
	StartCharge(ctx context.Context, id PhysicalID, p Power) error
	Stop(ctx context.Context, id PhysicalID) error
}

// SecretAgent is general integration access token storage system. It will store whatever (most likely JSON)
// marshalled into string and provide the same string back when needed.
type SecretAgent interface {
	// ProvideSecret provides token structure marshalled into string
	ProvideSecret() string

	// UpdateSecret persists token/refresh tokens (most likely a JSON structure) marshalled into string.
	UpdateSecret(string)
}

type PhysicalID string

type Energy float64

type Power float64

type Location struct {
	Latitude  float64
	Longitude float64
}

type DeviceStatus uint8

const (
	StatusUnknown DeviceStatus = iota
	StatusDisconnected
	StatusCharging
	StatusStopped
)

type ChargePointStatusType string

const (
	StatusOnline  ChargePointStatusType = "Online"
	StatusOffline ChargePointStatusType = "Offline"
)

type DeviceMetadata struct {
	ID       PhysicalID
	Name     string
	Location Location
}

type ChargerStatus struct {
	Timestamp    time.Time
	Power        Power
	Energy       Energy
	SetPoint     Power
	ChargeStatus DeviceStatus
}

type ChargePointStatus struct {
	ID     PhysicalID
	Status ChargePointStatusType
}

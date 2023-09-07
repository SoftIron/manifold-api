package sifi

import (
	"context"

	"github.com/softiron/hypercloud-api/client"
	"github.com/softiron/hypercloud-api/cloud"
)

// ZoneService owns the /zone methods.
type ZoneService struct {
	*client.Client
}

// DeleteZone deletes the zone with the given ID.
func (s ZoneService) DeleteZone(ctx context.Context, id int) error {
	p := path(cloud.ZonePath, id)

	return s.Delete(ctx, p, nil)
}

// Zone returns information about a zone.
func (s ZoneService) Zone(ctx context.Context, id int) (*cloud.ListZoneResponse, error) {
	var resp cloud.ListZoneResponse

	p := path(cloud.ZonePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Zones returns information about all zones.
func (s ZoneService) Zones(ctx context.Context) (*cloud.ListZonesResponse, error) {
	var resp cloud.ListZonesResponse

	p := path(cloud.ZonePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ZonesRaftStatus returns the raft status for all zones.
func (s ZoneService) ZonesRaftStatus(ctx context.Context) (*cloud.ListZonesRaftStatusResponse, error) {
	var resp cloud.ListZonesRaftStatusResponse

	p := path(cloud.ZonePath, "raft")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateZone updates the zone with the given ID.
func (s ZoneService) UpdateZone(ctx context.Context, id int, req cloud.UpdateZoneRequest) (*cloud.UpdateZoneResponse, error) {
	var resp cloud.UpdateZoneResponse

	p := path(cloud.ZonePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableZone enables the zone with the given ID.
func (s ZoneService) EnableZone(ctx context.Context, id int, req cloud.EnableZoneRequest) (*cloud.EnableZoneResponse, error) {
	var resp cloud.EnableZoneResponse

	p := path(cloud.ZonePath, id, "enable")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameZone renames the zone with the given ID.
func (s ZoneService) RenameZone(ctx context.Context, id int, req cloud.RenameZoneRequest) (*cloud.RenameZoneResponse, error) {
	var resp cloud.RenameZoneResponse

	p := path(cloud.ZonePath, id, "name")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateZone creates a new zone.
func (s ZoneService) CreateZone(ctx context.Context, req cloud.CreateZoneRequest) (*cloud.CreateZoneResponse, error) {
	var resp cloud.CreateZoneResponse

	p := path(cloud.ZonePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

package sifi

import (
	"context"

	"github.com/softiron/manifold-api/client"
	"github.com/softiron/manifold-api/deprecated/v1/hc"
)

// ZoneService owns the /zone methods.
type ZoneService struct {
	*client.Client
}

// DeleteZone deletes the zone with the given ID.
func (s ZoneService) DeleteZone(ctx context.Context, id int) error {
	p := path(hc.ZonePath, id)

	return s.Delete(ctx, p, nil)
}

// Zone returns information about a zone.
func (s ZoneService) Zone(ctx context.Context, id int) (*hc.ListZoneResponse, error) {
	var resp hc.ListZoneResponse

	p := path(hc.ZonePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Zones returns information about all zones.
func (s ZoneService) Zones(ctx context.Context) (*hc.ListZonesResponse, error) {
	var resp hc.ListZonesResponse

	p := path(hc.ZonePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ZonesRaftStatus returns the raft status for all zones.
func (s ZoneService) ZonesRaftStatus(ctx context.Context) (*hc.ListZonesRaftStatusResponse, error) {
	var resp hc.ListZonesRaftStatusResponse

	p := path(hc.ZonePath, "raft")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateZone updates the zone with the given ID.
func (s ZoneService) UpdateZone(ctx context.Context, id int, req hc.UpdateZoneRequest) (*hc.UpdateZoneResponse, error) {
	var resp hc.UpdateZoneResponse

	p := path(hc.ZonePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableZone enables the zone with the given ID.
func (s ZoneService) EnableZone(ctx context.Context, id int, req hc.EnableZoneRequest) (*hc.EnableZoneResponse, error) {
	var resp hc.EnableZoneResponse

	p := path(hc.ZonePath, id, "enable")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameZone renames the zone with the given ID.
func (s ZoneService) RenameZone(ctx context.Context, id int, req hc.RenameZoneRequest) (*hc.RenameZoneResponse, error) {
	var resp hc.RenameZoneResponse

	p := path(hc.ZonePath, id, "name")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateZone creates a new zone.
func (s ZoneService) CreateZone(ctx context.Context, req hc.CreateZoneRequest) (*hc.CreateZoneResponse, error) {
	var resp hc.CreateZoneResponse

	p := path(hc.ZonePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

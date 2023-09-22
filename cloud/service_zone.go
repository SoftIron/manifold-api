package cloud

import (
	"context"
)

// ZoneService owns the /cloud/zone methods.
type ZoneService struct {
	*service
}

// DeleteZone deletes the zone with the given ID.
func (s ZoneService) DeleteZone(ctx context.Context, id int) error {
	p := s.path(ZonePath, id)

	return s.Delete(ctx, p, nil)
}

// Zone returns information about a zone.
func (s ZoneService) Zone(ctx context.Context, id int) (*ZoneResponse, error) {
	var resp ZoneResponse

	p := s.path(ZonePath, id)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Zones returns information about all zones.
func (s ZoneService) Zones(ctx context.Context) (*ZonesResponse, error) {
	var resp ZonesResponse

	p := s.path(ZonePath)

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ZonesRaftStatus returns the raft status for all zones.
func (s ZoneService) ZonesRaftStatus(ctx context.Context) (*ZonesRaftStatusResponse, error) {
	var resp ZonesRaftStatusResponse

	p := s.path(ZonePath, "raft")

	if err := s.Get(ctx, p, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateZone updates the zone with the given ID.
func (s ZoneService) UpdateZone(ctx context.Context, id int, req UpdateZoneRequest) (*UpdateZoneResponse, error) {
	var resp UpdateZoneResponse

	p := s.path(ZonePath, id)

	if err := s.Patch(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EnableZone enables the zone with the given ID.
func (s ZoneService) EnableZone(ctx context.Context, id int, req EnableZoneRequest) (*EnableZoneResponse, error) {
	var resp EnableZoneResponse

	p := s.path(ZonePath, id, "enable")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// RenameZone renames the zone with the given ID.
func (s ZoneService) RenameZone(ctx context.Context, id int, req RenameZoneRequest) (*RenameZoneResponse, error) {
	var resp RenameZoneResponse

	p := s.path(ZonePath, id, "name")

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateZone creates a new zone.
func (s ZoneService) CreateZone(ctx context.Context, req CreateZoneRequest) (*CreateZoneResponse, error) {
	var resp CreateZoneResponse

	p := s.path(ZonePath)

	if err := s.Post(ctx, p, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

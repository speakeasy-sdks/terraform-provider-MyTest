// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ZoneCredential struct {
	ID *int64 `json:"id,omitempty"`
}

type ZoneGroups struct {
	AccountID *int64  `json:"accountId,omitempty"`
	ID        *int64  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// ZoneZoneType - Map containing the Cloud (zone) code name. See the zone-types API to fetch a list of all available Cloud (zone) types and their codes.
type ZoneZoneType struct {
	Code *string `json:"code,omitempty"`
}

type Zone struct {
	AccountID     *int64             `json:"accountId,omitempty"`
	Code          *string            `json:"code,omitempty"`
	Config        *ZoneVcenterConfig `json:"config,omitempty"`
	Credential    *ZoneCredential    `json:"credential,omitempty"`
	Enabled       *bool              `json:"enabled,omitempty"`
	Groups        []ZoneGroups       `json:"groups,omitempty"`
	ID            *int64             `json:"id,omitempty"`
	Name          *string            `json:"name,omitempty"`
	ScalePriority *int64             `json:"scalePriority,omitempty"`
	Visibility    *string            `json:"visibility,omitempty"`
	// Map containing the Cloud (zone) code name. See the zone-types API to fetch a list of all available Cloud (zone) types and their codes.
	ZoneType *ZoneZoneType `json:"zoneType,omitempty"`
}

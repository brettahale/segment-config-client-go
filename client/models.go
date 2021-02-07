package client

import (
    "encoding/json"
    "time"
 )

type Workspace struct {
	Name        string    `json:"name,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	ID          string    `json:"id,omitempty"`
	CreateTime  time.Time `json:"create_time,omitempty"`
}

// Sources defines the struct for the sources object
type Sources struct {
	Sources []Source `json:"sources,omitempty"`
}

// Source defines the struct for the source object
type Source struct {
	Name            string              `json:"name,omitempty"`
	CatalogName     string              `json:"catalog_name,omitempty"`
	Parent          string              `json:"parent,omitempty"`
	WriteKeys       []string            `json:"write_keys,omitempty"`
	LibraryConfig   LibraryConfig       `json:"library_config,omitempty"`
	FunctionConfig  []FunctionConfig    `json:"function_config,omitempty"`
	CreateTime      time.Time           `json:"create_time,omitempty"`
}

type SourceCreate struct {
    Source          Source              `json:"source"`
}

type SourceUpdate struct {
	Source          Source              `json:"source"`
	UpdateMask		UpdateMask			`json:"update_mask"`
}

type CreateSourceSchemaConfiguration struct {
    SchemaConfig    SourceSchemaConfiguration   `json:"schema_config"`
}

type SourceSchemaConfiguration struct {
        AllowUnplannedTrackEvents           bool    `json:"allow_unplanned_track_events"`
        AllowUnplannedIdentifyTraits        bool    `json:"allow_unplanned_identify_traits"`
        AllowUnplannedGroupTraits           bool    `json:"allow_unplanned_group_traits"`
        ForwardBlockedEventsTo              string  `json:"forwarding_blocked_events_to"`
        AllowUnplannedTrackEventProperties  bool    `json:"allow_unplanned_track_event_properties"`
        AllowTrackEventOnViolations         bool    `json:"allow_track_event_on_violations"`
        AllowIdentifyTraitsOnViolations     bool    `json:"allow_identify_traits_on_violations"`
        AllowGroupTraitsOnViolations        bool    `json:"allow_group_traits_on_violations"`
        ForwardingViolationsTo              string  `json:"forwarding_violations_to"`
        AllowTrackPropertiesOnViolations    bool    `json:"allow_track_properties_on_violations"`
        CommonTrackEventOnViolations        string  `json:"common_track_event_on_violations"`
    	CommonIdentifyEventOnViolations     string  `json:"common_identify_event_on_violations"`
    	CommonGroupEventOnViolations        string  `json:"common_group_event_on_violations" `
}

// LibraryConfig contains information about a source's library
type LibraryConfig struct {
	MetricsEnabled          bool         `json:"metrics_enabled,omitempty"`
	RetryQueue              bool         `json:"retry_queue,omitempty"`
	CrossDomainIDEnabled    bool         `json:"cross_domain_id_enabled,omitempty"`
	APIHost                 string       `json:"api_host,omitempty"`
}

type FunctionConfig struct {
	Name        string      `json:"name,omitempty"`
	Value       string      `json:"value,omitempty"`
	Type        string      `json:"type,omitempty"`
}

type TrackingPlanCreate struct {
    TrackingPlan          TrackingPlan              `json:"tracking_plan"`
}


type TrackingPlanUpdate struct {
	TrackingPlan    TrackingPlan        `json:"tracking_plan"`
	UpdateMask		UpdateMask			`json:"update_mask"`
}

type TrackingPlan struct {
    Name                string      `json:"name,omitempty"`
    DisplayName         string      `json:"display_name,omitempty"`
    Rules               Rules       `json:"rules,omitempty"`
    CreateTime          time.Time   `json:"create_time,omitempty"`
    UpdateTime          string       `json:"update_time,omitempty"`
}

type TrackingPlanSourceConnection struct {
    SourceName      string      `json:"source_name"`
    TrackingPlanId  string      `json:"tracking_plan_id"`
}

type Rules struct {
    Global           json.RawMessage    `json:"global,omitempty"`
    Events           []Event            `json:"events,omitempty"`
    Identify         json.RawMessage    `json:"identify,omitempty"`
    Group            json.RawMessage    `json:"group,omitempty"`
}

type Event struct {
    Name            string              `json:"name,omitempty"`
    Version         int                 `json:"version,omitempty"`
    Description     string              `json:"description,omitempty"`
    Rules           json.RawMessage     `json:"rules,omitempty"`
}


type DestinationCreate struct {
    Destination          Destination             `json:"destination"`
}

type Destinations struct {
	Destinations []Destination `json:"destinations,omitempty"`
}

type Destination struct {
	Name           string              `json:"name,omitempty"`
	Parent         string              `json:"parent,omitempty"`
	DisplayName    string              `json:"display_name,omitempty"`
	Enabled        bool                `json:"enabled,omitempty"`
	ConnectionMode string              `json:"connection_mode,omitempty"`
	Configs        []DestinationConfig `json:"config,omitempty"`
	CreateTime     time.Time           `json:"create_time,omitempty"`
	UpdateTime     time.Time           `json:"update_time,omitempty"`
}

type DestinationConfig struct {
	Name        string      `json:"name,omitempty"`
	DisplayName string      `json:"display_name,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	Type        string      `json:"type,omitempty"`
}

type CatalogSource struct {
    Name            string      `json:"name,omitempty"`
	DisplayName     string      `json:"display_name,omitempty"`
	Catagories      []string    `json:"categories,omitempty"`
	Description     string      `json:"description,omitempty"`
}

type CatalogDestination struct {
    Name            string                  `json:"name,omitempty"`
	DisplayName     string                  `json:"display_name,omitempty"`
	Description     string                  `json:"description,omitempty"`
	Type            string                  `json:"type,omitempty"`
	Status          string                  `json:"status,omitempty"`
    Logos           DestinationLogos        `json:"logos,omitempty"`
    Settings        []DestinationSetting    `json:"settings,omitempty"`
    Categories      DestinationCategory     `json:"categories,omitempty"`
    Compoenents     []DestinationComponent  `json:"components,omitempty"`
}

type DestinationLogos struct {
    Logo            string                  `json:"logo,omitempty"`
    Mark            string                  `json:"mark,omitempty"`
}

type DestinationSetting struct {
    Name                string                  `json:"name,omitempty"`
	DisplayName         string                  `json:"display_name,omitempty"`
	Type                string                  `json:"type,omitempty"`
	Deprecated          bool                    `json:"deprecated,omitempty"`
	Required            bool                    `json:"required,omitempty"`
	StringValidators    StringValidator         `json:"string_validators,omitempty"`
	MapValidators       MapValidator            `json:"map_validators,omitempty"`
	Settings            []string                `json:"settings,omitempty"`
}

type DestinationFilterCreate struct {
    Filter DestinationFilter `json:"filter,omitempty"`
}

// does not work with field_list action
type DestinationFilter struct {
    Name            string                  `json:"name,omitempty"`
    Enabled         bool                    `json:"enabled,omitempty"`
	Tile            string                  `json:"title,omitempty"`
	Actions         []FilterAction          `json:actions,omitempty`
	Identify         string                 `json:"if,omitempty"`
}

type FilterAction struct {
    Type        string                          `json:"name,omitempty"`
    Percent     float32                         `json:"percent,omitempty"`
    Path        string                          `json:"path,omitempty"`
    Fields      map[string][]string             `json:"fields,omitempty"`
}

type StringValidator struct {
    Regexp      string      `json:"regexp,omitempty"`
}

type MapValidator struct {
    Regexp          string      `json:"regexp,omitempty"`
    Min             int         `json:"min,omitempty"`
    Max             int         `json:"max,omitempty"`
    MapPrefix       string      `json:"map_prefix,omitempty"`
    SelectOptions   []string    `json:"select_options,omitempty"`
}

type DestinationCategory struct {
    Primary             string                  `json:"primary,omitempty"`
    Secondary           string                  `json:"secondary,omitempty"`
    Additional          []string                `json:"additional,omitempty"`
}

type DestinationComponent struct {
    Type        string      `json:"type,omitempty"`
}

type Function struct {
    Id              string                  `json:"id,omitempty"`
	DisplayName     string                  `json:"display_name,omitempty"`
	Type            string                  `json:"type,omitempty"`
	Code            string                  `json:"code,omitempty"`
	WorkspaceId     string                  `json:"workspace_id,omitempty"`
	Description     string                  `json:"description,omitempty"`
	Status          string                  `json:"status,omitempty"`
	CreatedAt       time.Time               `json:"created_at,omitempty"`
    CatalogId       string                  `json:"catalog_id,omitempty"`
	LogoUri         string                  `json:"logo_uri,omitempty"`
    Settings        []FunctionSettings      `json:"settings,omitempty"`
}

type FunctionSettings struct {
    Name            string                `json:"name,omitempty"`
	Label           string                `json:"label,omitempty"`
	Type            string                `json:"type,omitempty"`
	Sensitive       bool                  `json:"sensitive,omitempty"`
	Description     string                `json:"description,omitempty"`
}

type IamPolicy struct {
    Name            string                `json:"name,omitempty"`
	DisplayName     string                `json:"display_name,omitempty"`
	Actions         []string              `json:"actions,omitempty"`
	Description     string                `json:"description,omitempty"`
}

type IamRole struct {
    Name            string                `json:"name,omitempty"`
	Resource        string                `json:"resource,omitempty"`
	Subject         string                `json:"subject,omitempty"`
}

type IamInvite struct {
    Email            string                 `json:"name,omitempty"`
	CreateTime       time.Time              `json:"create_time,omitempty"`
	ExpireTime       time.Time              `json:"expire_time,omitempty"`
	Policies         []RoleResourceMap      `json:"policies,omitempty"`
}

type RoleResourceMap struct {
    Role            string      `json:"role,omitempty"`
    Resource        string      `json:"resource,omitempty"`
}

// UpdateMask contains information for updating Destinations
type UpdateMask struct {
	Paths []string `json:"paths,omitempty"`
}

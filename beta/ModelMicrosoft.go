// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// MicrosoftStoreForBusinessApp Microsoft Store for Business Apps. This class does not support Create, Delete, or Update.
type MicrosoftStoreForBusinessApp struct {
	// MobileApp is the base model of MicrosoftStoreForBusinessApp
	MobileApp
	// LicenseType The app license type
	LicenseType *MicrosoftStoreForBusinessLicenseType `json:"licenseType,omitempty"`
	// LicensingType The supported License Type.
	LicensingType *VPPLicensingType `json:"licensingType,omitempty"`
	// PackageIdentityName The app package identifier
	PackageIdentityName *string `json:"packageIdentityName,omitempty"`
	// ProductKey The app product key
	ProductKey *string `json:"productKey,omitempty"`
	// TotalLicenseCount The total number of Microsoft Store for Business licenses.
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// UsedLicenseCount The number of Microsoft Store for Business licenses in use.
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// ContainedApps undocumented
	ContainedApps []MobileContainedApp `json:"containedApps,omitempty"`
}

// MicrosoftStoreForBusinessAppAssignmentSettings Contains properties used to assign an Microsoft Store for Business mobile app to a group.
type MicrosoftStoreForBusinessAppAssignmentSettings struct {
	// MobileAppAssignmentSettings is the base model of MicrosoftStoreForBusinessAppAssignmentSettings
	MobileAppAssignmentSettings
	// UseDeviceContext Whether or not to use device execution context for Microsoft Store for Business mobile app.
	UseDeviceContext *bool `json:"useDeviceContext,omitempty"`
}

// MicrosoftStoreForBusinessContainedApp A class that represents a contained app of a MicrosoftStoreForBusinessApp.
type MicrosoftStoreForBusinessContainedApp struct {
	// MobileContainedApp is the base model of MicrosoftStoreForBusinessContainedApp
	MobileContainedApp
	// AppUserModelID The app user model ID of the contained app of a MicrosoftStoreForBusinessApp.
	AppUserModelID *string `json:"appUserModelId,omitempty"`
}

// MicrosoftTunnelConfiguration Entity that represents a collection of Microsoft Tunnel settings
type MicrosoftTunnelConfiguration struct {
	// Entity is the base model of MicrosoftTunnelConfiguration
	Entity
	// AdvancedSettings Additional settings that may be applied to the server
	AdvancedSettings []KeyValuePair `json:"advancedSettings,omitempty"`
	// DefaultDomainSuffix The Default Domain appendix that will be used by the clients
	DefaultDomainSuffix *string `json:"defaultDomainSuffix,omitempty"`
	// Description The MicrosoftTunnelConfiguration's description
	Description *string `json:"description,omitempty"`
	// DisplayName The MicrosoftTunnelConfiguration's display name
	DisplayName *string `json:"displayName,omitempty"`
	// DNSServers The DNS servers that will be used by the clients
	DNSServers []string `json:"dnsServers,omitempty"`
	// LastUpdateDateTime When the MicrosoftTunnelConfiguration was last updated
	LastUpdateDateTime *time.Time `json:"lastUpdateDateTime,omitempty"`
	// ListenPort The port that both TCP and UPD will listen over on the server
	ListenPort *int `json:"listenPort,omitempty"`
	// Network The subnet that will be used to allocate virtual address for the clients
	Network *string `json:"network,omitempty"`
	// RoleScopeTagIDs List of Scope Tags for this Entity instance.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// RoutesExclude Subsets of the routes that will not be routed by the server
	RoutesExclude []string `json:"routesExclude,omitempty"`
	// RoutesInclude The routs that will be routed by the server
	RoutesInclude []string `json:"routesInclude,omitempty"`
	// SplitDNS The domains that will be resolved using the provided dns servers
	SplitDNS []string `json:"splitDNS,omitempty"`
}

// MicrosoftTunnelHealthThreshold Entity that represents the health thresholds of a health metric.
type MicrosoftTunnelHealthThreshold struct {
	// Entity is the base model of MicrosoftTunnelHealthThreshold
	Entity
	// DefaultHealthyThreshold The default threshold for being healthy
	DefaultHealthyThreshold *int `json:"defaultHealthyThreshold,omitempty"`
	// DefaultUnhealthyThreshold The default threshold for being unhealthy
	DefaultUnhealthyThreshold *int `json:"defaultUnhealthyThreshold,omitempty"`
	// HealthyThreshold The threshold for being healthy
	HealthyThreshold *int `json:"healthyThreshold,omitempty"`
	// UnhealthyThreshold The threshold for being unhealthy
	UnhealthyThreshold *int `json:"unhealthyThreshold,omitempty"`
}

// MicrosoftTunnelServer Entity that represents a single Microsoft Tunnel server
type MicrosoftTunnelServer struct {
	// Entity is the base model of MicrosoftTunnelServer
	Entity
	// DisplayName The MicrosoftTunnelServer's display name
	DisplayName *string `json:"displayName,omitempty"`
	// LastCheckinDateTime When the MicrosoftTunnelServer last checked in
	LastCheckinDateTime *time.Time `json:"lastCheckinDateTime,omitempty"`
	// TunnelServerHealthStatus The MicrosoftTunnelServer's health status
	TunnelServerHealthStatus *MicrosoftTunnelServerHealthStatus `json:"tunnelServerHealthStatus,omitempty"`
}

// MicrosoftTunnelServerLogCollectionResponse Entity that stores the server log collection status.
type MicrosoftTunnelServerLogCollectionResponse struct {
	// Entity is the base model of MicrosoftTunnelServerLogCollectionResponse
	Entity
	// EndDateTime The end time of the logs collected
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// ExpiryDateTime The time when the log collection is expired
	ExpiryDateTime *time.Time `json:"expiryDateTime,omitempty"`
	// RequestDateTime The time when the log collection was requested
	RequestDateTime *time.Time `json:"requestDateTime,omitempty"`
	// ServerID ID of the server the log collection is requested upon
	ServerID *string `json:"serverId,omitempty"`
	// SizeInBytes The size of the logs in bytes
	SizeInBytes *int `json:"sizeInBytes,omitempty"`
	// StartDateTime The start time of the logs collected
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Status The status of log collection
	Status *MicrosoftTunnelLogCollectionStatus `json:"status,omitempty"`
}

// MicrosoftTunnelSite Entity that represents a Microsoft Tunnel site
type MicrosoftTunnelSite struct {
	// Entity is the base model of MicrosoftTunnelSite
	Entity
	// Description The MicrosoftTunnelSite's description
	Description *string `json:"description,omitempty"`
	// DisplayName The MicrosoftTunnelSite's display name
	DisplayName *string `json:"displayName,omitempty"`
	// PublicAddress The MicrosoftTunnelSite's public domain name or IP address
	PublicAddress *string `json:"publicAddress,omitempty"`
	// RoleScopeTagIDs List of Scope Tags for this Entity instance.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// MicrosoftTunnelConfiguration undocumented
	MicrosoftTunnelConfiguration *MicrosoftTunnelConfiguration `json:"microsoftTunnelConfiguration,omitempty"`
	// MicrosoftTunnelServers undocumented
	MicrosoftTunnelServers []MicrosoftTunnelServer `json:"microsoftTunnelServers,omitempty"`
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ComanagementEligibleDevice Device Co-Management eligibility state
type ComanagementEligibleDevice struct {
	// Entity is the base model of ComanagementEligibleDevice
	Entity
	// ClientRegistrationStatus ClientRegistrationStatus
	ClientRegistrationStatus *DeviceRegistrationState `json:"clientRegistrationStatus,omitempty"`
	// DeviceName DeviceName
	DeviceName *string `json:"deviceName,omitempty"`
	// DeviceType DeviceType
	DeviceType *DeviceType `json:"deviceType,omitempty"`
	// EntitySource EntitySource
	EntitySource *int `json:"entitySource,omitempty"`
	// ManagementAgents ManagementAgents
	ManagementAgents *ManagementAgentType `json:"managementAgents,omitempty"`
	// ManagementState ManagementState
	ManagementState *ManagementState `json:"managementState,omitempty"`
	// Manufacturer Manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`
	// MDMStatus MDMStatus
	MDMStatus *string `json:"mdmStatus,omitempty"`
	// Model Model
	Model *string `json:"model,omitempty"`
	// OsDescription OSDescription
	OsDescription *string `json:"osDescription,omitempty"`
	// OsVersion OSVersion
	OsVersion *string `json:"osVersion,omitempty"`
	// OwnerType OwnerType
	OwnerType *OwnerType `json:"ownerType,omitempty"`
	// ReferenceID ReferenceId
	ReferenceID *string `json:"referenceId,omitempty"`
	// SerialNumber SerialNumber
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Status ComanagementEligibleStatus
	Status *ComanagementEligibleType `json:"status,omitempty"`
	// Upn UPN
	Upn *string `json:"upn,omitempty"`
	// UserEmail UserEmail
	UserEmail *string `json:"userEmail,omitempty"`
	// UserID UserId
	UserID *string `json:"userId,omitempty"`
	// UserName UserName
	UserName *string `json:"userName,omitempty"`
}

// ComanagementEligibleDevicesSummary undocumented
type ComanagementEligibleDevicesSummary struct {
	// Object is the base model of ComanagementEligibleDevicesSummary
	Object
	// ComanagedCount Count of devices already Co-Managed
	ComanagedCount *int `json:"comanagedCount,omitempty"`
	// EligibleButNotAzureAdJoinedCount Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
	EligibleButNotAzureAdJoinedCount *int `json:"eligibleButNotAzureAdJoinedCount,omitempty"`
	// EligibleCount Count of devices fully eligible for Co-Management
	EligibleCount *int `json:"eligibleCount,omitempty"`
	// IneligibleCount Count of devices ineligible for Co-Management
	IneligibleCount *int `json:"ineligibleCount,omitempty"`
	// NeedsOsUpdateCount Count of devices that will be eligible for Co-Management after an OS update
	NeedsOsUpdateCount *int `json:"needsOsUpdateCount,omitempty"`
}
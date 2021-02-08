// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Directory undocumented
type Directory struct {
	// Object is the base model of Directory
	Object
	// AdministrativeUnits undocumented
	AdministrativeUnits []AdministrativeUnit `json:"administrativeUnits,omitempty"`
	// DeletedItems undocumented
	DeletedItems []DirectoryObject `json:"deletedItems,omitempty"`
	// SharedEmailDomains undocumented
	SharedEmailDomains []SharedEmailDomain `json:"sharedEmailDomains,omitempty"`
	// FeatureRolloutPolicies undocumented
	FeatureRolloutPolicies []FeatureRolloutPolicy `json:"featureRolloutPolicies,omitempty"`
}

// DirectoryAudit undocumented
type DirectoryAudit struct {
	// Entity is the base model of DirectoryAudit
	Entity
	// ActivityDateTime undocumented
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// ActivityDisplayName undocumented
	ActivityDisplayName *string `json:"activityDisplayName,omitempty"`
	// AdditionalDetails undocumented
	AdditionalDetails []KeyValue `json:"additionalDetails,omitempty"`
	// Category undocumented
	Category *string `json:"category,omitempty"`
	// CorrelationID undocumented
	CorrelationID *string `json:"correlationId,omitempty"`
	// InitiatedBy undocumented
	InitiatedBy *AuditActivityInitiator `json:"initiatedBy,omitempty"`
	// LoggedByService undocumented
	LoggedByService *string `json:"loggedByService,omitempty"`
	// OperationType undocumented
	OperationType *string `json:"operationType,omitempty"`
	// Result undocumented
	Result *OperationResult `json:"result,omitempty"`
	// ResultReason undocumented
	ResultReason *string `json:"resultReason,omitempty"`
	// TargetResources undocumented
	TargetResources []TargetResource `json:"targetResources,omitempty"`
}

// DirectoryDefinition undocumented
type DirectoryDefinition struct {
	// Entity is the base model of DirectoryDefinition
	Entity
	// Discoverabilities undocumented
	Discoverabilities *DirectoryDefinitionDiscoverabilities `json:"discoverabilities,omitempty"`
	// DiscoveryDateTime undocumented
	DiscoveryDateTime *time.Time `json:"discoveryDateTime,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Objects undocumented
	Objects []ObjectDefinition `json:"objects,omitempty"`
	// ReadOnly undocumented
	ReadOnly *bool `json:"readOnly,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
}

// DirectoryObject Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type DirectoryObject struct {
	// Entity is the base model of DirectoryObject
	Entity
	// DeletedDateTime undocumented
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
}

// DirectoryObjectPartnerReference undocumented
type DirectoryObjectPartnerReference struct {
	// DirectoryObject is the base model of DirectoryObjectPartnerReference
	DirectoryObject
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalPartnerTenantID undocumented
	ExternalPartnerTenantID *UUID `json:"externalPartnerTenantId,omitempty"`
	// ObjectType undocumented
	ObjectType *string `json:"objectType,omitempty"`
}

// DirectoryRole undocumented
type DirectoryRole struct {
	// DirectoryObject is the base model of DirectoryRole
	DirectoryObject
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// RoleTemplateID undocumented
	RoleTemplateID *string `json:"roleTemplateId,omitempty"`
	// Members undocumented
	Members []DirectoryObject `json:"members,omitempty"`
	// ScopedMembers undocumented
	ScopedMembers []ScopedRoleMembership `json:"scopedMembers,omitempty"`
}

// DirectoryRoleAccessReviewPolicy undocumented
type DirectoryRoleAccessReviewPolicy struct {
	// Entity is the base model of DirectoryRoleAccessReviewPolicy
	Entity
	// Settings undocumented
	Settings *AccessReviewScheduleSettings `json:"settings,omitempty"`
}

// DirectoryRoleTemplate undocumented
type DirectoryRoleTemplate struct {
	// DirectoryObject is the base model of DirectoryRoleTemplate
	DirectoryObject
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// DirectorySetting undocumented
type DirectorySetting struct {
	// Entity is the base model of DirectorySetting
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// Values undocumented
	Values []SettingValue `json:"values,omitempty"`
}

// DirectorySettingTemplate undocumented
type DirectorySettingTemplate struct {
	// DirectoryObject is the base model of DirectorySettingTemplate
	DirectoryObject
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Values undocumented
	Values []SettingTemplateValue `json:"values,omitempty"`
}

// DirectorySizeQuota undocumented
type DirectorySizeQuota struct {
	// Object is the base model of DirectorySizeQuota
	Object
	// Total undocumented
	Total *int `json:"total,omitempty"`
	// Used undocumented
	Used *int `json:"used,omitempty"`
}

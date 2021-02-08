// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// TeamsApp undocumented
type TeamsApp struct {
	// Entity is the base model of TeamsApp
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DistributionMethod undocumented
	DistributionMethod *TeamsAppDistributionMethod `json:"distributionMethod,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// AppDefinitions undocumented
	AppDefinitions []TeamsAppDefinition `json:"appDefinitions,omitempty"`
}

// TeamsAppDefinition undocumented
type TeamsAppDefinition struct {
	// Entity is the base model of TeamsAppDefinition
	Entity
	// AzureADAppID undocumented
	AzureADAppID *string `json:"azureADAppId,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// PublishingState undocumented
	PublishingState *TeamsAppPublishingState `json:"publishingState,omitempty"`
	// Shortdescription undocumented
	Shortdescription *string `json:"shortdescription,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
	// Bot undocumented
	Bot *TeamworkBot `json:"bot,omitempty"`
}

// TeamsAppInstallation undocumented
type TeamsAppInstallation struct {
	// Entity is the base model of TeamsAppInstallation
	Entity
	// TeamsApp undocumented
	TeamsApp *TeamsApp `json:"teamsApp,omitempty"`
	// TeamsAppDefinition undocumented
	TeamsAppDefinition *TeamsAppDefinition `json:"teamsAppDefinition,omitempty"`
}

// TeamsAsyncOperation undocumented
type TeamsAsyncOperation struct {
	// Entity is the base model of TeamsAsyncOperation
	Entity
	// AttemptsCount undocumented
	AttemptsCount *int `json:"attemptsCount,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Error undocumented
	Error *OperationError `json:"error,omitempty"`
	// LastActionDateTime undocumented
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	// OperationType undocumented
	OperationType *TeamsAsyncOperationType `json:"operationType,omitempty"`
	// Status undocumented
	Status *TeamsAsyncOperationStatus `json:"status,omitempty"`
	// TargetResourceID undocumented
	TargetResourceID *string `json:"targetResourceId,omitempty"`
	// TargetResourceLocation undocumented
	TargetResourceLocation *string `json:"targetResourceLocation,omitempty"`
}

// TeamsDeviceUsageDistributionUserCounts undocumented
type TeamsDeviceUsageDistributionUserCounts struct {
	// Entity is the base model of TeamsDeviceUsageDistributionUserCounts
	Entity
	// AndroidPhone undocumented
	AndroidPhone *int `json:"androidPhone,omitempty"`
	// IOS undocumented
	IOS *int `json:"ios,omitempty"`
	// Mac undocumented
	Mac *int `json:"mac,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// Windows undocumented
	Windows *int `json:"windows,omitempty"`
	// WindowsPhone undocumented
	WindowsPhone *int `json:"windowsPhone,omitempty"`
}

// TeamsDeviceUsageUserCounts undocumented
type TeamsDeviceUsageUserCounts struct {
	// Entity is the base model of TeamsDeviceUsageUserCounts
	Entity
	// AndroidPhone undocumented
	AndroidPhone *int `json:"androidPhone,omitempty"`
	// IOS undocumented
	IOS *int `json:"ios,omitempty"`
	// Mac undocumented
	Mac *int `json:"mac,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// Windows undocumented
	Windows *int `json:"windows,omitempty"`
	// WindowsPhone undocumented
	WindowsPhone *int `json:"windowsPhone,omitempty"`
}

// TeamsDeviceUsageUserDetail undocumented
type TeamsDeviceUsageUserDetail struct {
	// Entity is the base model of TeamsDeviceUsageUserDetail
	Entity
	// DeletedDate undocumented
	DeletedDate *Date `json:"deletedDate,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// UsedAndroidPhone undocumented
	UsedAndroidPhone *bool `json:"usedAndroidPhone,omitempty"`
	// UsediOS undocumented
	UsediOS *bool `json:"usediOS,omitempty"`
	// UsedMac undocumented
	UsedMac *bool `json:"usedMac,omitempty"`
	// UsedWeb undocumented
	UsedWeb *bool `json:"usedWeb,omitempty"`
	// UsedWindows undocumented
	UsedWindows *bool `json:"usedWindows,omitempty"`
	// UsedWindowsPhone undocumented
	UsedWindowsPhone *bool `json:"usedWindowsPhone,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// TeamsTab undocumented
type TeamsTab struct {
	// Entity is the base model of TeamsTab
	Entity
	// Configuration undocumented
	Configuration *TeamsTabConfiguration `json:"configuration,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// MessageID undocumented
	MessageID *string `json:"messageId,omitempty"`
	// SortOrderIndex undocumented
	SortOrderIndex *string `json:"sortOrderIndex,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// TeamsApp undocumented
	TeamsApp *TeamsApp `json:"teamsApp,omitempty"`
}

// TeamsTabConfiguration undocumented
type TeamsTabConfiguration struct {
	// Object is the base model of TeamsTabConfiguration
	Object
	// ContentURL undocumented
	ContentURL *string `json:"contentUrl,omitempty"`
	// EntityID undocumented
	EntityID *string `json:"entityId,omitempty"`
	// RemoveURL undocumented
	RemoveURL *string `json:"removeUrl,omitempty"`
	// WebsiteURL undocumented
	WebsiteURL *string `json:"websiteUrl,omitempty"`
}

// TeamsTemplate undocumented
type TeamsTemplate struct {
	// Entity is the base model of TeamsTemplate
	Entity
}

// TeamsUserActivityCounts undocumented
type TeamsUserActivityCounts struct {
	// Entity is the base model of TeamsUserActivityCounts
	Entity
	// Calls undocumented
	Calls *int `json:"calls,omitempty"`
	// Meetings undocumented
	Meetings *int `json:"meetings,omitempty"`
	// PrivateChatMessages undocumented
	PrivateChatMessages *int `json:"privateChatMessages,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// TeamChatMessages undocumented
	TeamChatMessages *int `json:"teamChatMessages,omitempty"`
}

// TeamsUserActivityUserCounts undocumented
type TeamsUserActivityUserCounts struct {
	// Entity is the base model of TeamsUserActivityUserCounts
	Entity
	// Calls undocumented
	Calls *int `json:"calls,omitempty"`
	// Meetings undocumented
	Meetings *int `json:"meetings,omitempty"`
	// OtherActions undocumented
	OtherActions *int `json:"otherActions,omitempty"`
	// PrivateChatMessages undocumented
	PrivateChatMessages *int `json:"privateChatMessages,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// TeamChatMessages undocumented
	TeamChatMessages *int `json:"teamChatMessages,omitempty"`
}

// TeamsUserActivityUserDetail undocumented
type TeamsUserActivityUserDetail struct {
	// Entity is the base model of TeamsUserActivityUserDetail
	Entity
	// AssignedProducts undocumented
	AssignedProducts []string `json:"assignedProducts,omitempty"`
	// CallCount undocumented
	CallCount *int `json:"callCount,omitempty"`
	// DeletedDate undocumented
	DeletedDate *Date `json:"deletedDate,omitempty"`
	// HasOtherAction undocumented
	HasOtherAction *bool `json:"hasOtherAction,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// MeetingCount undocumented
	MeetingCount *int `json:"meetingCount,omitempty"`
	// PrivateChatMessageCount undocumented
	PrivateChatMessageCount *int `json:"privateChatMessageCount,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// TeamChatMessageCount undocumented
	TeamChatMessageCount *int `json:"teamChatMessageCount,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Office365ActivationCounts undocumented
type Office365ActivationCounts struct {
	// Entity is the base model of Office365ActivationCounts
	Entity
	// Android undocumented
	Android *int `json:"android,omitempty"`
	// IOS undocumented
	IOS *int `json:"ios,omitempty"`
	// Mac undocumented
	Mac *int `json:"mac,omitempty"`
	// ProductType undocumented
	ProductType *string `json:"productType,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Windows undocumented
	Windows *int `json:"windows,omitempty"`
	// Windows10Mobile undocumented
	Windows10Mobile *int `json:"windows10Mobile,omitempty"`
}

// Office365ActivationsUserCounts undocumented
type Office365ActivationsUserCounts struct {
	// Entity is the base model of Office365ActivationsUserCounts
	Entity
	// Activated undocumented
	Activated *int `json:"activated,omitempty"`
	// Assigned undocumented
	Assigned *int `json:"assigned,omitempty"`
	// ProductType undocumented
	ProductType *string `json:"productType,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SharedComputerActivation undocumented
	SharedComputerActivation *int `json:"sharedComputerActivation,omitempty"`
}

// Office365ActivationsUserDetail undocumented
type Office365ActivationsUserDetail struct {
	// Entity is the base model of Office365ActivationsUserDetail
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// UserActivationCounts undocumented
	UserActivationCounts []UserActivationCounts `json:"userActivationCounts,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// Office365ActiveUserCounts undocumented
type Office365ActiveUserCounts struct {
	// Entity is the base model of Office365ActiveUserCounts
	Entity
	// Exchange undocumented
	Exchange *int `json:"exchange,omitempty"`
	// Office365 undocumented
	Office365 *int `json:"office365,omitempty"`
	// OneDrive undocumented
	OneDrive *int `json:"oneDrive,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SharePoint undocumented
	SharePoint *int `json:"sharePoint,omitempty"`
	// SkypeForBusiness undocumented
	SkypeForBusiness *int `json:"skypeForBusiness,omitempty"`
	// Teams undocumented
	Teams *int `json:"teams,omitempty"`
	// Yammer undocumented
	Yammer *int `json:"yammer,omitempty"`
}

// Office365ActiveUserDetail undocumented
type Office365ActiveUserDetail struct {
	// Entity is the base model of Office365ActiveUserDetail
	Entity
	// AssignedProducts undocumented
	AssignedProducts []string `json:"assignedProducts,omitempty"`
	// DeletedDate undocumented
	DeletedDate *Date `json:"deletedDate,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExchangeLastActivityDate undocumented
	ExchangeLastActivityDate *Date `json:"exchangeLastActivityDate,omitempty"`
	// ExchangeLicenseAssignDate undocumented
	ExchangeLicenseAssignDate *Date `json:"exchangeLicenseAssignDate,omitempty"`
	// HasExchangeLicense undocumented
	HasExchangeLicense *bool `json:"hasExchangeLicense,omitempty"`
	// HasOneDriveLicense undocumented
	HasOneDriveLicense *bool `json:"hasOneDriveLicense,omitempty"`
	// HasSharePointLicense undocumented
	HasSharePointLicense *bool `json:"hasSharePointLicense,omitempty"`
	// HasSkypeForBusinessLicense undocumented
	HasSkypeForBusinessLicense *bool `json:"hasSkypeForBusinessLicense,omitempty"`
	// HasTeamsLicense undocumented
	HasTeamsLicense *bool `json:"hasTeamsLicense,omitempty"`
	// HasYammerLicense undocumented
	HasYammerLicense *bool `json:"hasYammerLicense,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// OneDriveLastActivityDate undocumented
	OneDriveLastActivityDate *Date `json:"oneDriveLastActivityDate,omitempty"`
	// OneDriveLicenseAssignDate undocumented
	OneDriveLicenseAssignDate *Date `json:"oneDriveLicenseAssignDate,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SharePointLastActivityDate undocumented
	SharePointLastActivityDate *Date `json:"sharePointLastActivityDate,omitempty"`
	// SharePointLicenseAssignDate undocumented
	SharePointLicenseAssignDate *Date `json:"sharePointLicenseAssignDate,omitempty"`
	// SkypeForBusinessLastActivityDate undocumented
	SkypeForBusinessLastActivityDate *Date `json:"skypeForBusinessLastActivityDate,omitempty"`
	// SkypeForBusinessLicenseAssignDate undocumented
	SkypeForBusinessLicenseAssignDate *Date `json:"skypeForBusinessLicenseAssignDate,omitempty"`
	// TeamsLastActivityDate undocumented
	TeamsLastActivityDate *Date `json:"teamsLastActivityDate,omitempty"`
	// TeamsLicenseAssignDate undocumented
	TeamsLicenseAssignDate *Date `json:"teamsLicenseAssignDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// YammerLastActivityDate undocumented
	YammerLastActivityDate *Date `json:"yammerLastActivityDate,omitempty"`
	// YammerLicenseAssignDate undocumented
	YammerLicenseAssignDate *Date `json:"yammerLicenseAssignDate,omitempty"`
}

// Office365GroupsActivityCounts undocumented
type Office365GroupsActivityCounts struct {
	// Entity is the base model of Office365GroupsActivityCounts
	Entity
	// ExchangeEmailsReceived undocumented
	ExchangeEmailsReceived *int `json:"exchangeEmailsReceived,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// YammerMessagesLiked undocumented
	YammerMessagesLiked *int `json:"yammerMessagesLiked,omitempty"`
	// YammerMessagesPosted undocumented
	YammerMessagesPosted *int `json:"yammerMessagesPosted,omitempty"`
	// YammerMessagesRead undocumented
	YammerMessagesRead *int `json:"yammerMessagesRead,omitempty"`
}

// Office365GroupsActivityDetail undocumented
type Office365GroupsActivityDetail struct {
	// Entity is the base model of Office365GroupsActivityDetail
	Entity
	// ExchangeMailboxStorageUsedInBytes undocumented
	ExchangeMailboxStorageUsedInBytes *int `json:"exchangeMailboxStorageUsedInBytes,omitempty"`
	// ExchangeMailboxTotalItemCount undocumented
	ExchangeMailboxTotalItemCount *int `json:"exchangeMailboxTotalItemCount,omitempty"`
	// ExchangeReceivedEmailCount undocumented
	ExchangeReceivedEmailCount *int `json:"exchangeReceivedEmailCount,omitempty"`
	// ExternalMemberCount undocumented
	ExternalMemberCount *int `json:"externalMemberCount,omitempty"`
	// GroupDisplayName undocumented
	GroupDisplayName *string `json:"groupDisplayName,omitempty"`
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
	// GroupType undocumented
	GroupType *string `json:"groupType,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// MemberCount undocumented
	MemberCount *int `json:"memberCount,omitempty"`
	// OwnerPrincipalName undocumented
	OwnerPrincipalName *string `json:"ownerPrincipalName,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SharePointActiveFileCount undocumented
	SharePointActiveFileCount *int `json:"sharePointActiveFileCount,omitempty"`
	// SharePointSiteStorageUsedInBytes undocumented
	SharePointSiteStorageUsedInBytes *int `json:"sharePointSiteStorageUsedInBytes,omitempty"`
	// SharePointTotalFileCount undocumented
	SharePointTotalFileCount *int `json:"sharePointTotalFileCount,omitempty"`
	// YammerLikedMessageCount undocumented
	YammerLikedMessageCount *int `json:"yammerLikedMessageCount,omitempty"`
	// YammerPostedMessageCount undocumented
	YammerPostedMessageCount *int `json:"yammerPostedMessageCount,omitempty"`
	// YammerReadMessageCount undocumented
	YammerReadMessageCount *int `json:"yammerReadMessageCount,omitempty"`
}

// Office365GroupsActivityFileCounts undocumented
type Office365GroupsActivityFileCounts struct {
	// Entity is the base model of Office365GroupsActivityFileCounts
	Entity
	// Active undocumented
	Active *int `json:"active,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
}

// Office365GroupsActivityGroupCounts undocumented
type Office365GroupsActivityGroupCounts struct {
	// Entity is the base model of Office365GroupsActivityGroupCounts
	Entity
	// Active undocumented
	Active *int `json:"active,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
}

// Office365GroupsActivityStorage undocumented
type Office365GroupsActivityStorage struct {
	// Entity is the base model of Office365GroupsActivityStorage
	Entity
	// MailboxStorageUsedInBytes undocumented
	MailboxStorageUsedInBytes *int `json:"mailboxStorageUsedInBytes,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SiteStorageUsedInBytes undocumented
	SiteStorageUsedInBytes *int `json:"siteStorageUsedInBytes,omitempty"`
}

// Office365ServicesUserCounts undocumented
type Office365ServicesUserCounts struct {
	// Entity is the base model of Office365ServicesUserCounts
	Entity
	// ExchangeActive undocumented
	ExchangeActive *int `json:"exchangeActive,omitempty"`
	// ExchangeInactive undocumented
	ExchangeInactive *int `json:"exchangeInactive,omitempty"`
	// Office365Active undocumented
	Office365Active *int `json:"office365Active,omitempty"`
	// Office365Inactive undocumented
	Office365Inactive *int `json:"office365Inactive,omitempty"`
	// OneDriveActive undocumented
	OneDriveActive *int `json:"oneDriveActive,omitempty"`
	// OneDriveInactive undocumented
	OneDriveInactive *int `json:"oneDriveInactive,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SharePointActive undocumented
	SharePointActive *int `json:"sharePointActive,omitempty"`
	// SharePointInactive undocumented
	SharePointInactive *int `json:"sharePointInactive,omitempty"`
	// SkypeForBusinessActive undocumented
	SkypeForBusinessActive *int `json:"skypeForBusinessActive,omitempty"`
	// SkypeForBusinessInactive undocumented
	SkypeForBusinessInactive *int `json:"skypeForBusinessInactive,omitempty"`
	// TeamsActive undocumented
	TeamsActive *int `json:"teamsActive,omitempty"`
	// TeamsInactive undocumented
	TeamsInactive *int `json:"teamsInactive,omitempty"`
	// YammerActive undocumented
	YammerActive *int `json:"yammerActive,omitempty"`
	// YammerInactive undocumented
	YammerInactive *int `json:"yammerInactive,omitempty"`
}

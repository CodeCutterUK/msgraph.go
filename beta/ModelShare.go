// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ShareAction undocumented
type ShareAction struct {
	// Object is the base model of ShareAction
	Object
	// Recipients undocumented
	Recipients []IdentitySet `json:"recipients,omitempty"`
}

// SharePointActivityPages undocumented
type SharePointActivityPages struct {
	// Entity is the base model of SharePointActivityPages
	Entity
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// VisitedPageCount undocumented
	VisitedPageCount *int `json:"visitedPageCount,omitempty"`
}

// SharePointActivityUserCounts undocumented
type SharePointActivityUserCounts struct {
	// Entity is the base model of SharePointActivityUserCounts
	Entity
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SharedExternally undocumented
	SharedExternally *int `json:"sharedExternally,omitempty"`
	// SharedInternally undocumented
	SharedInternally *int `json:"sharedInternally,omitempty"`
	// Synced undocumented
	Synced *int `json:"synced,omitempty"`
	// ViewedOrEdited undocumented
	ViewedOrEdited *int `json:"viewedOrEdited,omitempty"`
	// VisitedPage undocumented
	VisitedPage *int `json:"visitedPage,omitempty"`
}

// SharePointActivityUserDetail undocumented
type SharePointActivityUserDetail struct {
	// Entity is the base model of SharePointActivityUserDetail
	Entity
	// AssignedProducts undocumented
	AssignedProducts []string `json:"assignedProducts,omitempty"`
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
	// SharedExternallyFileCount undocumented
	SharedExternallyFileCount *int `json:"sharedExternallyFileCount,omitempty"`
	// SharedInternallyFileCount undocumented
	SharedInternallyFileCount *int `json:"sharedInternallyFileCount,omitempty"`
	// SyncedFileCount undocumented
	SyncedFileCount *int `json:"syncedFileCount,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// ViewedOrEditedFileCount undocumented
	ViewedOrEditedFileCount *int `json:"viewedOrEditedFileCount,omitempty"`
	// VisitedPageCount undocumented
	VisitedPageCount *int `json:"visitedPageCount,omitempty"`
}

// SharePointSiteUsageDetail undocumented
type SharePointSiteUsageDetail struct {
	// Entity is the base model of SharePointSiteUsageDetail
	Entity
	// ActiveFileCount undocumented
	ActiveFileCount *int `json:"activeFileCount,omitempty"`
	// FileCount undocumented
	FileCount *int `json:"fileCount,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *Date `json:"lastActivityDate,omitempty"`
	// OwnerDisplayName undocumented
	OwnerDisplayName *string `json:"ownerDisplayName,omitempty"`
	// OwnerPrincipalName undocumented
	OwnerPrincipalName *string `json:"ownerPrincipalName,omitempty"`
	// PageViewCount undocumented
	PageViewCount *int `json:"pageViewCount,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// RootWebTemplate undocumented
	RootWebTemplate *string `json:"rootWebTemplate,omitempty"`
	// SiteID undocumented
	SiteID *UUID `json:"siteId,omitempty"`
	// SiteURL undocumented
	SiteURL *string `json:"siteUrl,omitempty"`
	// StorageAllocatedInBytes undocumented
	StorageAllocatedInBytes *int `json:"storageAllocatedInBytes,omitempty"`
	// StorageUsedInBytes undocumented
	StorageUsedInBytes *int `json:"storageUsedInBytes,omitempty"`
	// VisitedPageCount undocumented
	VisitedPageCount *int `json:"visitedPageCount,omitempty"`
}

// SharePointSiteUsageFileCounts undocumented
type SharePointSiteUsageFileCounts struct {
	// Entity is the base model of SharePointSiteUsageFileCounts
	Entity
	// Active undocumented
	Active *int `json:"active,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SiteType undocumented
	SiteType *string `json:"siteType,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
}

// SharePointSiteUsagePages undocumented
type SharePointSiteUsagePages struct {
	// Entity is the base model of SharePointSiteUsagePages
	Entity
	// PageViewCount undocumented
	PageViewCount *int `json:"pageViewCount,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SiteType undocumented
	SiteType *string `json:"siteType,omitempty"`
}

// SharePointSiteUsageSiteCounts undocumented
type SharePointSiteUsageSiteCounts struct {
	// Entity is the base model of SharePointSiteUsageSiteCounts
	Entity
	// Active undocumented
	Active *int `json:"active,omitempty"`
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SiteType undocumented
	SiteType *string `json:"siteType,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
}

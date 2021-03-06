// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AppCatalogs undocumented
type AppCatalogs struct {
	// Object is the base model of AppCatalogs
	Object
	// TeamsApps undocumented
	TeamsApps []TeamsApp `json:"teamsApps,omitempty"`
}

// AppConfigurationSettingItem Contains properties for App configuration setting item.
type AppConfigurationSettingItem struct {
	// Object is the base model of AppConfigurationSettingItem
	Object
	// AppConfigKey app configuration key.
	AppConfigKey *string `json:"appConfigKey,omitempty"`
	// AppConfigKeyType app configuration key type.
	AppConfigKeyType *MDMAppConfigKeyType `json:"appConfigKeyType,omitempty"`
	// AppConfigKeyValue app configuration key value.
	AppConfigKeyValue *string `json:"appConfigKeyValue,omitempty"`
}

// AppConsentApprovalRoute undocumented
type AppConsentApprovalRoute struct {
	// Entity is the base model of AppConsentApprovalRoute
	Entity
	// AppConsentRequests undocumented
	AppConsentRequests []AppConsentRequestObject `json:"appConsentRequests,omitempty"`
}

// AppConsentRequestObject undocumented
type AppConsentRequestObject struct {
	// Entity is the base model of AppConsentRequestObject
	Entity
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// ConsentType undocumented
	ConsentType *string `json:"consentType,omitempty"`
	// PendingScopes undocumented
	PendingScopes []AppConsentRequestScope `json:"pendingScopes,omitempty"`
	// UserConsentRequests undocumented
	UserConsentRequests []UserConsentRequestObject `json:"userConsentRequests,omitempty"`
}

// AppConsentRequestScope undocumented
type AppConsentRequestScope struct {
	// Object is the base model of AppConsentRequestScope
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// AppHostedMediaConfig undocumented
type AppHostedMediaConfig struct {
	// MediaConfig is the base model of AppHostedMediaConfig
	MediaConfig
	// Blob undocumented
	Blob *string `json:"blob,omitempty"`
}

// AppIdentity undocumented
type AppIdentity struct {
	// Object is the base model of AppIdentity
	Object
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ServicePrincipalID undocumented
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty"`
	// ServicePrincipalName undocumented
	ServicePrincipalName *string `json:"servicePrincipalName,omitempty"`
}

// AppListItem Represents an app in the list of managed applications
type AppListItem struct {
	// Object is the base model of AppListItem
	Object
	// AppID The application or bundle identifier of the application
	AppID *string `json:"appId,omitempty"`
	// AppStoreURL The Store URL of the application
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// Name The application name
	Name *string `json:"name,omitempty"`
	// Publisher The publisher of the application
	Publisher *string `json:"publisher,omitempty"`
}

// AppLogCollectionDownloadDetails undocumented
type AppLogCollectionDownloadDetails struct {
	// Object is the base model of AppLogCollectionDownloadDetails
	Object
	// AppLogDecryptionAlgorithm DecryptionAlgorithm for Content
	AppLogDecryptionAlgorithm *AppLogDecryptionAlgorithm `json:"appLogDecryptionAlgorithm,omitempty"`
	// DecryptionKey DecryptionKey as string
	DecryptionKey *string `json:"decryptionKey,omitempty"`
	// DownloadURL Download SAS Url for completed AppLogUploadRequest
	DownloadURL *string `json:"downloadUrl,omitempty"`
}

// AppLogCollectionRequestObject AppLogCollectionRequest Entity.
type AppLogCollectionRequestObject struct {
	// Entity is the base model of AppLogCollectionRequestObject
	Entity
	// CompletedDateTime Time at which the upload log request reached a terminal state
	CompletedDateTime *time.Time `json:"completedDateTime,omitempty"`
	// CustomLogFolders List of log folders.
	CustomLogFolders []string `json:"customLogFolders,omitempty"`
	// ErrorMessage Error message if any during the upload process
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Status Log upload status
	Status *AppLogUploadState `json:"status,omitempty"`
}

// AppMetadata undocumented
type AppMetadata struct {
	// Object is the base model of AppMetadata
	Object
	// Data undocumented
	Data []AppMetadataEntry `json:"data,omitempty"`
	// Version undocumented
	Version *int `json:"version,omitempty"`
}

// AppMetadataEntry undocumented
type AppMetadataEntry struct {
	// Object is the base model of AppMetadataEntry
	Object
	// Key undocumented
	Key *string `json:"key,omitempty"`
	// Value undocumented
	Value *Binary `json:"value,omitempty"`
}

// AppRole undocumented
type AppRole struct {
	// Object is the base model of AppRole
	Object
	// AllowedMemberTypes undocumented
	AllowedMemberTypes []string `json:"allowedMemberTypes,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Origin undocumented
	Origin *string `json:"origin,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

// AppRoleAssignment undocumented
type AppRoleAssignment struct {
	// Entity is the base model of AppRoleAssignment
	Entity
	// AppRoleID undocumented
	AppRoleID *UUID `json:"appRoleId,omitempty"`
	// CreationTimestamp undocumented
	CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
	// PrincipalDisplayName undocumented
	PrincipalDisplayName *string `json:"principalDisplayName,omitempty"`
	// PrincipalID undocumented
	PrincipalID *UUID `json:"principalId,omitempty"`
	// PrincipalType undocumented
	PrincipalType *string `json:"principalType,omitempty"`
	// ResourceDisplayName undocumented
	ResourceDisplayName *string `json:"resourceDisplayName,omitempty"`
	// ResourceID undocumented
	ResourceID *UUID `json:"resourceId,omitempty"`
}

// AppScope undocumented
type AppScope struct {
	// Entity is the base model of AppScope
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// AppVulnerabilityManagedDevice An app vulnerability managed device.
type AppVulnerabilityManagedDevice struct {
	// Entity is the base model of AppVulnerabilityManagedDevice
	Entity
	// DisplayName The device name.
	DisplayName *string `json:"displayName,omitempty"`
	// LastSyncDateTime The created date.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// ManagedDeviceID The Intune managed device ID.
	ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
}

// AppVulnerabilityMobileApp An app vulnerability mobile app.
type AppVulnerabilityMobileApp struct {
	// Entity is the base model of AppVulnerabilityMobileApp
	Entity
	// CreatedDateTime The created date.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName The device name.
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime The last modified date.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// MobileAppID The Intune mobile app ID.
	MobileAppID *string `json:"mobileAppId,omitempty"`
	// MobileAppType The app type.
	MobileAppType *string `json:"mobileAppType,omitempty"`
	// Version The app version.
	Version *string `json:"version,omitempty"`
}

// AppVulnerabilityTask An app vulnerability task.
type AppVulnerabilityTask struct {
	// DeviceAppManagementTask is the base model of AppVulnerabilityTask
	DeviceAppManagementTask
	// AppName The app name.
	AppName *string `json:"appName,omitempty"`
	// AppPublisher The app publisher.
	AppPublisher *string `json:"appPublisher,omitempty"`
	// AppVersion The app version.
	AppVersion *string `json:"appVersion,omitempty"`
	// Insights Information about the mitigation.
	Insights *string `json:"insights,omitempty"`
	// ManagedDeviceCount The number of vulnerable devices.
	ManagedDeviceCount *int `json:"managedDeviceCount,omitempty"`
	// MitigationType The mitigation type.
	MitigationType *AppVulnerabilityTaskMitigationType `json:"mitigationType,omitempty"`
	// MobileAppCount The number of vulnerable mobile apps.
	MobileAppCount *int `json:"mobileAppCount,omitempty"`
	// Remediation The remediation steps.
	Remediation *string `json:"remediation,omitempty"`
	// ManagedDevices undocumented
	ManagedDevices []AppVulnerabilityManagedDevice `json:"managedDevices,omitempty"`
	// MobileApps undocumented
	MobileApps []AppVulnerabilityMobileApp `json:"mobileApps,omitempty"`
}

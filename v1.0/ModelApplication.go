// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Application undocumented
type Application struct {
	// DirectoryObject is the base model of Application
	DirectoryObject
	// AddIns undocumented
	AddIns []AddIn `json:"addIns,omitempty"`
	// API undocumented
	API *APIApplication `json:"api,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// ApplicationTemplateID undocumented
	ApplicationTemplateID *string `json:"applicationTemplateId,omitempty"`
	// AppRoles undocumented
	AppRoles []AppRole `json:"appRoles,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GroupMembershipClaims undocumented
	GroupMembershipClaims *string `json:"groupMembershipClaims,omitempty"`
	// IdentifierUris undocumented
	IdentifierUris []string `json:"identifierUris,omitempty"`
	// Info undocumented
	Info *InformationalURL `json:"info,omitempty"`
	// IsDeviceOnlyAuthSupported undocumented
	IsDeviceOnlyAuthSupported *bool `json:"isDeviceOnlyAuthSupported,omitempty"`
	// IsFallbackPublicClient undocumented
	IsFallbackPublicClient *bool `json:"isFallbackPublicClient,omitempty"`
	// KeyCredentials undocumented
	KeyCredentials []KeyCredential `json:"keyCredentials,omitempty"`
	// Logo undocumented
	Logo *Stream `json:"logo,omitempty"`
	// Notes undocumented
	Notes *string `json:"notes,omitempty"`
	// OAuth2RequirePostResponse undocumented
	OAuth2RequirePostResponse *bool `json:"oauth2RequirePostResponse,omitempty"`
	// OptionalClaims undocumented
	OptionalClaims *OptionalClaims `json:"optionalClaims,omitempty"`
	// ParentalControlSettings undocumented
	ParentalControlSettings *ParentalControlSettings `json:"parentalControlSettings,omitempty"`
	// PasswordCredentials undocumented
	PasswordCredentials []PasswordCredential `json:"passwordCredentials,omitempty"`
	// PublicClient undocumented
	PublicClient *PublicClientApplication `json:"publicClient,omitempty"`
	// PublisherDomain undocumented
	PublisherDomain *string `json:"publisherDomain,omitempty"`
	// RequiredResourceAccess undocumented
	RequiredResourceAccess []RequiredResourceAccess `json:"requiredResourceAccess,omitempty"`
	// SignInAudience undocumented
	SignInAudience *string `json:"signInAudience,omitempty"`
	// Spa undocumented
	Spa *SpaApplication `json:"spa,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// TokenEncryptionKeyID undocumented
	TokenEncryptionKeyID *UUID `json:"tokenEncryptionKeyId,omitempty"`
	// Web undocumented
	Web *WebApplication `json:"web,omitempty"`
	// CreatedOnBehalfOf undocumented
	CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`
	// ExtensionProperties undocumented
	ExtensionProperties []ExtensionProperty `json:"extensionProperties,omitempty"`
	// HomeRealmDiscoveryPolicies undocumented
	HomeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicy `json:"homeRealmDiscoveryPolicies,omitempty"`
	// Owners undocumented
	Owners []DirectoryObject `json:"owners,omitempty"`
	// TokenIssuancePolicies undocumented
	TokenIssuancePolicies []TokenIssuancePolicy `json:"tokenIssuancePolicies,omitempty"`
	// TokenLifetimePolicies undocumented
	TokenLifetimePolicies []TokenLifetimePolicy `json:"tokenLifetimePolicies,omitempty"`
}

// ApplicationEnforcedRestrictionsSessionControl undocumented
type ApplicationEnforcedRestrictionsSessionControl struct {
	// ConditionalAccessSessionControl is the base model of ApplicationEnforcedRestrictionsSessionControl
	ConditionalAccessSessionControl
}

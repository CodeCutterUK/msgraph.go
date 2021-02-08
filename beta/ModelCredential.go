// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Credential undocumented
type Credential struct {
	// Object is the base model of Credential
	Object
	// FieldID undocumented
	FieldID *string `json:"fieldId,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

// CredentialSingleSignOnExtension Represents a Credential-type Single Sign-On extension profile.
type CredentialSingleSignOnExtension struct {
	// SingleSignOnExtension is the base model of CredentialSingleSignOnExtension
	SingleSignOnExtension
	// Configurations Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
	Configurations []KeyTypedValuePair `json:"configurations,omitempty"`
	// Domains Gets or sets a list of hosts or domain names for which the app extension performs SSO.
	Domains []string `json:"domains,omitempty"`
	// ExtensionIdentifier Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
	ExtensionIdentifier *string `json:"extensionIdentifier,omitempty"`
	// Realm Gets or sets the case-sensitive realm name for this profile.
	Realm *string `json:"realm,omitempty"`
	// TeamIdentifier Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
	TeamIdentifier *string `json:"teamIdentifier,omitempty"`
}

// CredentialUsageSummary undocumented
type CredentialUsageSummary struct {
	// Entity is the base model of CredentialUsageSummary
	Entity
	// AuthMethod undocumented
	AuthMethod *UsageAuthMethod `json:"authMethod,omitempty"`
	// FailureActivityCount undocumented
	FailureActivityCount *int `json:"failureActivityCount,omitempty"`
	// Feature undocumented
	Feature *FeatureType `json:"feature,omitempty"`
	// SuccessfulActivityCount undocumented
	SuccessfulActivityCount *int `json:"successfulActivityCount,omitempty"`
}

// CredentialUserRegistrationCount undocumented
type CredentialUserRegistrationCount struct {
	// Entity is the base model of CredentialUserRegistrationCount
	Entity
	// TotalUserCount undocumented
	TotalUserCount *int `json:"totalUserCount,omitempty"`
	// UserRegistrationCounts undocumented
	UserRegistrationCounts []UserRegistrationCount `json:"userRegistrationCounts,omitempty"`
}

// CredentialUserRegistrationDetails undocumented
type CredentialUserRegistrationDetails struct {
	// Entity is the base model of CredentialUserRegistrationDetails
	Entity
	// AuthMethods undocumented
	AuthMethods []RegistrationAuthMethod `json:"authMethods,omitempty"`
	// IsCapable undocumented
	IsCapable *bool `json:"isCapable,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// IsMFARegistered undocumented
	IsMFARegistered *bool `json:"isMfaRegistered,omitempty"`
	// IsRegistered undocumented
	IsRegistered *bool `json:"isRegistered,omitempty"`
	// UserDisplayName undocumented
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

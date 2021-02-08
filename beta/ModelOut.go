// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OutOfBoxExperienceSettings Out of box experience setting
type OutOfBoxExperienceSettings struct {
	// Object is the base model of OutOfBoxExperienceSettings
	Object
	// DeviceUsageType AAD join authentication type
	DeviceUsageType *WindowsDeviceUsageType `json:"deviceUsageType,omitempty"`
	// HideEscapeLink If set to true, then the user can't start over with different account, on company sign-in
	HideEscapeLink *bool `json:"hideEscapeLink,omitempty"`
	// HideEULA Show or hide EULA to user
	HideEULA *bool `json:"hideEULA,omitempty"`
	// HidePrivacySettings Show or hide privacy settings to user
	HidePrivacySettings *bool `json:"hidePrivacySettings,omitempty"`
	// SkipKeyboardSelectionPage If set, then skip the keyboard selection page if Language and Region are set
	SkipKeyboardSelectionPage *bool `json:"skipKeyboardSelectionPage,omitempty"`
	// UserType Type of user
	UserType *WindowsUserType `json:"userType,omitempty"`
}

// OutOfOfficeSettings undocumented
type OutOfOfficeSettings struct {
	// Object is the base model of OutOfOfficeSettings
	Object
	// IsOutOfOffice undocumented
	IsOutOfOffice *bool `json:"isOutOfOffice,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
}

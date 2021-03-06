// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AdminConsent Admin consent information.
type AdminConsent struct {
	// Object is the base model of AdminConsent
	Object
	// ShareAPNSData The admin consent state of sharing user and device data to Apple.
	ShareAPNSData *AdminConsentState `json:"shareAPNSData,omitempty"`
	// ShareUserExperienceAnalyticsData Gets or sets the admin consent for sharing User experience analytics data.
	ShareUserExperienceAnalyticsData *AdminConsentState `json:"shareUserExperienceAnalyticsData,omitempty"`
}

// AdminConsentRequestPolicy undocumented
type AdminConsentRequestPolicy struct {
	// Entity is the base model of AdminConsentRequestPolicy
	Entity
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// NotifyReviewers undocumented
	NotifyReviewers *bool `json:"notifyReviewers,omitempty"`
	// RemindersEnabled undocumented
	RemindersEnabled *bool `json:"remindersEnabled,omitempty"`
	// RequestDurationInDays undocumented
	RequestDurationInDays *int `json:"requestDurationInDays,omitempty"`
	// Reviewers undocumented
	Reviewers []AccessReviewScope `json:"reviewers,omitempty"`
	// Version undocumented
	Version *int `json:"version,omitempty"`
}

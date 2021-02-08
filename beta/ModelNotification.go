// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Notification undocumented
type Notification struct {
	// Entity is the base model of Notification
	Entity
	// DisplayTimeToLive undocumented
	DisplayTimeToLive *int `json:"displayTimeToLive,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// GroupName undocumented
	GroupName *string `json:"groupName,omitempty"`
	// Payload undocumented
	Payload *PayloadTypes `json:"payload,omitempty"`
	// Priority undocumented
	Priority *Priority `json:"priority,omitempty"`
	// TargetHostName undocumented
	TargetHostName *string `json:"targetHostName,omitempty"`
	// TargetPolicy undocumented
	TargetPolicy *TargetPolicyEndpoints `json:"targetPolicy,omitempty"`
}

// NotificationMessageTemplate Notification messages are messages that are sent to end users who are determined to be not-compliant with the compliance policies defined by the administrator. Administrators choose notifications and configure them in the Intune Admin Console using the compliance policy creation page under the “Actions for non-compliance” section. Use the notificationMessageTemplate object to create your own custom notifications for administrators to choose while configuring actions for non-compliance.
type NotificationMessageTemplate struct {
	// Entity is the base model of NotificationMessageTemplate
	Entity
	// BrandingOptions The Message Template Branding Options. Branding is defined in the Intune Admin Console.
	BrandingOptions *NotificationTemplateBrandingOptions `json:"brandingOptions,omitempty"`
	// DefaultLocale The default locale to fallback onto when the requested locale is not available.
	DefaultLocale *string `json:"defaultLocale,omitempty"`
	// DisplayName Display name for the Notification Message Template.
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// RoleScopeTagIDs List of Scope Tags for this Entity instance.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// LocalizedNotificationMessages undocumented
	LocalizedNotificationMessages []LocalizedNotificationMessage `json:"localizedNotificationMessages,omitempty"`
}

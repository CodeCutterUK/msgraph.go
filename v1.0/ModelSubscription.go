// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Subscription undocumented
type Subscription struct {
	// Entity is the base model of Subscription
	Entity
	// ApplicationID undocumented
	ApplicationID *string `json:"applicationId,omitempty"`
	// ChangeType undocumented
	ChangeType *string `json:"changeType,omitempty"`
	// ClientState undocumented
	ClientState *string `json:"clientState,omitempty"`
	// CreatorID undocumented
	CreatorID *string `json:"creatorId,omitempty"`
	// EncryptionCertificate undocumented
	EncryptionCertificate *string `json:"encryptionCertificate,omitempty"`
	// EncryptionCertificateID undocumented
	EncryptionCertificateID *string `json:"encryptionCertificateId,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// IncludeResourceData undocumented
	IncludeResourceData *bool `json:"includeResourceData,omitempty"`
	// LatestSupportedTLSVersion undocumented
	LatestSupportedTLSVersion *string `json:"latestSupportedTlsVersion,omitempty"`
	// LifecycleNotificationURL undocumented
	LifecycleNotificationURL *string `json:"lifecycleNotificationUrl,omitempty"`
	// NotificationURL undocumented
	NotificationURL *string `json:"notificationUrl,omitempty"`
	// Resource undocumented
	Resource *string `json:"resource,omitempty"`
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// SharingDetail undocumented
type SharingDetail struct {
	// Object is the base model of SharingDetail
	Object
	// SharedBy undocumented
	SharedBy *InsightIdentity `json:"sharedBy,omitempty"`
	// SharedDateTime undocumented
	SharedDateTime *time.Time `json:"sharedDateTime,omitempty"`
	// SharingReference undocumented
	SharingReference *ResourceReference `json:"sharingReference,omitempty"`
	// SharingSubject undocumented
	SharingSubject *string `json:"sharingSubject,omitempty"`
	// SharingType undocumented
	SharingType *string `json:"sharingType,omitempty"`
}

// SharingInvitation undocumented
type SharingInvitation struct {
	// Object is the base model of SharingInvitation
	Object
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// InvitedBy undocumented
	InvitedBy *IdentitySet `json:"invitedBy,omitempty"`
	// RedeemedBy undocumented
	RedeemedBy *string `json:"redeemedBy,omitempty"`
	// SignInRequired undocumented
	SignInRequired *bool `json:"signInRequired,omitempty"`
}

// SharingLink undocumented
type SharingLink struct {
	// Object is the base model of SharingLink
	Object
	// Application undocumented
	Application *Identity `json:"application,omitempty"`
	// PreventsDownload undocumented
	PreventsDownload *bool `json:"preventsDownload,omitempty"`
	// Scope undocumented
	Scope *string `json:"scope,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// WebHTML undocumented
	WebHTML *string `json:"webHtml,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Mention undocumented
type Mention struct {
	// Entity is the base model of Mention
	Entity
	// Application undocumented
	Application *string `json:"application,omitempty"`
	// ClientReference undocumented
	ClientReference *string `json:"clientReference,omitempty"`
	// CreatedBy undocumented
	CreatedBy *EmailAddress `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DeepLink undocumented
	DeepLink *string `json:"deepLink,omitempty"`
	// Mentioned undocumented
	Mentioned *EmailAddress `json:"mentioned,omitempty"`
	// MentionText undocumented
	MentionText *string `json:"mentionText,omitempty"`
	// ServerCreatedDateTime undocumented
	ServerCreatedDateTime *time.Time `json:"serverCreatedDateTime,omitempty"`
}

// MentionAction undocumented
type MentionAction struct {
	// Object is the base model of MentionAction
	Object
	// Mentionees undocumented
	Mentionees []IdentitySet `json:"mentionees,omitempty"`
}

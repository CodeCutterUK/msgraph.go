// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AadUserConversationMember undocumented
type AadUserConversationMember struct {
	// ConversationMember is the base model of AadUserConversationMember
	ConversationMember
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// TenantID undocumented
	TenantID *string `json:"tenantId,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// User undocumented
	User *User `json:"user,omitempty"`
}

// AadUserConversationMemberResult undocumented
type AadUserConversationMemberResult struct {
	// ActionResultPart is the base model of AadUserConversationMemberResult
	ActionResultPart
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

// AadUserNotificationRecipient undocumented
type AadUserNotificationRecipient struct {
	// TeamworkNotificationRecipient is the base model of AadUserNotificationRecipient
	TeamworkNotificationRecipient
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

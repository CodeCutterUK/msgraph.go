// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// InvitedUserMessageInfo undocumented
type InvitedUserMessageInfo struct {
	// Object is the base model of InvitedUserMessageInfo
	Object
	// CcRecipients undocumented
	CcRecipients []Recipient `json:"ccRecipients,omitempty"`
	// CustomizedMessageBody undocumented
	CustomizedMessageBody *string `json:"customizedMessageBody,omitempty"`
	// MessageLanguage undocumented
	MessageLanguage *string `json:"messageLanguage,omitempty"`
}

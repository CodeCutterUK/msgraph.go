// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TypedEmailAddress undocumented
type TypedEmailAddress struct {
	EmailAddress
	// Type undocumented
	Type *EmailType `json:"type,omitempty"`
	// OtherLabel undocumented
	OtherLabel *string `json:"otherLabel,omitempty"`
}
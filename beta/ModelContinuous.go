// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ContinuousAccessEvaluationPolicy undocumented
type ContinuousAccessEvaluationPolicy struct {
	// Entity is the base model of ContinuousAccessEvaluationPolicy
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Groups undocumented
	Groups []string `json:"groups,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Users undocumented
	Users []string `json:"users,omitempty"`
}

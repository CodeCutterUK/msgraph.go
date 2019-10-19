// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppliedConditionalAccessPolicy undocumented
type AppliedConditionalAccessPolicy struct {
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EnforcedGrantControls undocumented
	EnforcedGrantControls []string `json:"enforcedGrantControls,omitempty"`
	// EnforcedSessionControls undocumented
	EnforcedSessionControls []string `json:"enforcedSessionControls,omitempty"`
	// ConditionsSatisfied undocumented
	ConditionsSatisfied *ConditionalAccessConditions `json:"conditionsSatisfied,omitempty"`
	// ConditionsNotSatisfied undocumented
	ConditionsNotSatisfied *ConditionalAccessConditions `json:"conditionsNotSatisfied,omitempty"`
	// Result undocumented
	Result *AppliedConditionalAccessPolicyResult `json:"result,omitempty"`
}
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AutoLabeling undocumented
type AutoLabeling struct {
	// Object is the base model of AutoLabeling
	Object
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// SensitiveTypeIDs undocumented
	SensitiveTypeIDs []string `json:"sensitiveTypeIds,omitempty"`
}

// AutoReviewSettings undocumented
type AutoReviewSettings struct {
	// Object is the base model of AutoReviewSettings
	Object
	// NotReviewedResult undocumented
	NotReviewedResult *string `json:"notReviewedResult,omitempty"`
}

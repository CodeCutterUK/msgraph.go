// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// LabelActionBase undocumented
type LabelActionBase struct {
	// Object is the base model of LabelActionBase
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

// LabelDetails undocumented
type LabelDetails struct {
	// ParentLabelDetails is the base model of LabelDetails
	ParentLabelDetails
}

// LabelPolicy undocumented
type LabelPolicy struct {
	// Object is the base model of LabelPolicy
	Object
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

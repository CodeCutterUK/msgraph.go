// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// IPNamedLocation undocumented
type IPNamedLocation struct {
	// NamedLocation is the base model of IPNamedLocation
	NamedLocation
	// IPRanges undocumented
	IPRanges []IPRange `json:"ipRanges,omitempty"`
	// IsTrusted undocumented
	IsTrusted *bool `json:"isTrusted,omitempty"`
}

// IPRange undocumented
type IPRange struct {
	// Object is the base model of IPRange
	Object
}

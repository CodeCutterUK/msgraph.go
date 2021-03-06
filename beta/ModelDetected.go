// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DetectedApp A managed or unmanaged app that is installed on a managed device. Unmanaged apps will only appear for devices marked as corporate owned.
type DetectedApp struct {
	// Entity is the base model of DetectedApp
	Entity
	// DeviceCount The number of devices that have installed this application
	DeviceCount *int `json:"deviceCount,omitempty"`
	// DisplayName Name of the discovered application. Read-only
	DisplayName *string `json:"displayName,omitempty"`
	// SizeInByte Discovered application size in bytes. Read-only
	SizeInByte *int `json:"sizeInByte,omitempty"`
	// Version Version of the discovered application. Read-only
	Version *string `json:"version,omitempty"`
	// ManagedDevices undocumented
	ManagedDevices []ManagedDevice `json:"managedDevices,omitempty"`
}

// DetectedSensitiveContent undocumented
type DetectedSensitiveContent struct {
	// DetectedSensitiveContentBase is the base model of DetectedSensitiveContent
	DetectedSensitiveContentBase
	// ClassificationAttributes undocumented
	ClassificationAttributes []ClassificationAttribute `json:"classificationAttributes,omitempty"`
	// ClassificationMethod undocumented
	ClassificationMethod *ClassificationMethod `json:"classificationMethod,omitempty"`
	// Matches undocumented
	Matches []SensitiveContentLocation `json:"matches,omitempty"`
	// Scope undocumented
	Scope *SensitiveTypeScope `json:"scope,omitempty"`
	// SensitiveTypeSource undocumented
	SensitiveTypeSource *SensitiveTypeSource `json:"sensitiveTypeSource,omitempty"`
}

// DetectedSensitiveContentBase undocumented
type DetectedSensitiveContentBase struct {
	// Object is the base model of DetectedSensitiveContentBase
	Object
	// Confidence undocumented
	Confidence *int `json:"confidence,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// RecommendedConfidence undocumented
	RecommendedConfidence *int `json:"recommendedConfidence,omitempty"`
	// UniqueCount undocumented
	UniqueCount *int `json:"uniqueCount,omitempty"`
}

// DetectedSensitiveContentWrapper undocumented
type DetectedSensitiveContentWrapper struct {
	// Object is the base model of DetectedSensitiveContentWrapper
	Object
	// Classification undocumented
	Classification []DetectedSensitiveContent `json:"classification,omitempty"`
}

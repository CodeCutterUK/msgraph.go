// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SettingSource undocumented
type SettingSource struct {
	// Object is the base model of SettingSource
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

// SettingStateDeviceSummary Device Compilance Policy and Configuration for a Setting State summary
type SettingStateDeviceSummary struct {
	// Entity is the base model of SettingStateDeviceSummary
	Entity
	// CompliantDeviceCount Device Compliant count for the setting
	CompliantDeviceCount *int `json:"compliantDeviceCount,omitempty"`
	// ConflictDeviceCount Device conflict error count for the setting
	ConflictDeviceCount *int `json:"conflictDeviceCount,omitempty"`
	// ErrorDeviceCount Device error count for the setting
	ErrorDeviceCount *int `json:"errorDeviceCount,omitempty"`
	// InstancePath Name of the InstancePath for the setting
	InstancePath *string `json:"instancePath,omitempty"`
	// NonCompliantDeviceCount Device NonCompliant count for the setting
	NonCompliantDeviceCount *int `json:"nonCompliantDeviceCount,omitempty"`
	// NotApplicableDeviceCount Device Not Applicable count for the setting
	NotApplicableDeviceCount *int `json:"notApplicableDeviceCount,omitempty"`
	// RemediatedDeviceCount Device Compliant count for the setting
	RemediatedDeviceCount *int `json:"remediatedDeviceCount,omitempty"`
	// SettingName Name of the setting
	SettingName *string `json:"settingName,omitempty"`
	// UnknownDeviceCount Device Unkown count for the setting
	UnknownDeviceCount *int `json:"unknownDeviceCount,omitempty"`
}

// SettingTemplateValue undocumented
type SettingTemplateValue struct {
	// Object is the base model of SettingTemplateValue
	Object
	// DefaultValue undocumented
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// SettingValue undocumented
type SettingValue struct {
	// Object is the base model of SettingValue
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

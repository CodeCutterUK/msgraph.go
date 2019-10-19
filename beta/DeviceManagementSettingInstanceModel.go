// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementSettingInstance Base type for a setting instance
type DeviceManagementSettingInstance struct {
	Entity
	// DefinitionID The ID of the setting definition for this instance
	DefinitionID *string `json:"definitionId,omitempty"`
	// ValueJSON JSON representation of the value
	ValueJSON *string `json:"valueJson,omitempty"`
}
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// UpdateRecordingStatusOperation undocumented
type UpdateRecordingStatusOperation struct {
	// CommsOperation is the base model of UpdateRecordingStatusOperation
	CommsOperation
}

// UpdateWindowsDeviceAccountActionParameter undocumented
type UpdateWindowsDeviceAccountActionParameter struct {
	// Object is the base model of UpdateWindowsDeviceAccountActionParameter
	Object
	// CalendarSyncEnabled undocumented
	CalendarSyncEnabled *bool `json:"calendarSyncEnabled,omitempty"`
	// DeviceAccount undocumented
	DeviceAccount *WindowsDeviceAccount `json:"deviceAccount,omitempty"`
	// DeviceAccountEmail undocumented
	DeviceAccountEmail *string `json:"deviceAccountEmail,omitempty"`
	// ExchangeServer undocumented
	ExchangeServer *string `json:"exchangeServer,omitempty"`
	// PasswordRotationEnabled undocumented
	PasswordRotationEnabled *bool `json:"passwordRotationEnabled,omitempty"`
	// SessionInitiationProtocalAddress undocumented
	SessionInitiationProtocalAddress *string `json:"sessionInitiationProtocalAddress,omitempty"`
}

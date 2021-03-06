// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// VulnerableManagedDevice This entity represents a device associated with a task.
type VulnerableManagedDevice struct {
	// Entity is the base model of VulnerableManagedDevice
	Entity
	// DisplayName The device name.
	DisplayName *string `json:"displayName,omitempty"`
	// LastSyncDateTime The last sync date.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// ManagedDeviceID The Intune managed device ID.
	ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
}

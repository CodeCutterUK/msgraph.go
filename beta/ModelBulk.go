// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BulkManagedDeviceActionResult undocumented
type BulkManagedDeviceActionResult struct {
	// Object is the base model of BulkManagedDeviceActionResult
	Object
	// FailedDeviceIDs Failed devices
	FailedDeviceIDs []string `json:"failedDeviceIds,omitempty"`
	// NotFoundDeviceIDs Not found devices
	NotFoundDeviceIDs []string `json:"notFoundDeviceIds,omitempty"`
	// NotSupportedDeviceIDs Not supported devices
	NotSupportedDeviceIDs []string `json:"notSupportedDeviceIds,omitempty"`
	// SuccessfulDeviceIDs Successful devices
	SuccessfulDeviceIDs []string `json:"successfulDeviceIds,omitempty"`
}

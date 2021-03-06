// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EBookInstallSummary Contains properties for the installation summary of a book for a device.
type EBookInstallSummary struct {
	// Entity is the base model of EBookInstallSummary
	Entity
	// FailedDeviceCount Number of Devices that have failed to install this book.
	FailedDeviceCount *int `json:"failedDeviceCount,omitempty"`
	// FailedUserCount Number of Users that have 1 or more device that failed to install this book.
	FailedUserCount *int `json:"failedUserCount,omitempty"`
	// InstalledDeviceCount Number of Devices that have successfully installed this book.
	InstalledDeviceCount *int `json:"installedDeviceCount,omitempty"`
	// InstalledUserCount Number of Users whose devices have all succeeded to install this book.
	InstalledUserCount *int `json:"installedUserCount,omitempty"`
	// NotInstalledDeviceCount Number of Devices that does not have this book installed.
	NotInstalledDeviceCount *int `json:"notInstalledDeviceCount,omitempty"`
	// NotInstalledUserCount Number of Users that did not install this book.
	NotInstalledUserCount *int `json:"notInstalledUserCount,omitempty"`
}

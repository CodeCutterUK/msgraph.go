// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSVppAppAssignmentSettings undocumented
type IOSVppAppAssignmentSettings struct {
	MobileAppAssignmentSettings
	// UseDeviceLicensing Whether or not to use device licensing.
	UseDeviceLicensing *bool `json:"useDeviceLicensing,omitempty"`
	// VpnConfigurationID The VPN Configuration Id to apply for this app.
	VpnConfigurationID *string `json:"vpnConfigurationId,omitempty"`
}
// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSExtensionsConfiguration MacOS extensions configuration profile.
type MacOSExtensionsConfiguration struct {
	DeviceConfiguration
	// KernelExtensionOverridesAllowed If set to true, users can approve additional kernel extensions not explicitly allowed by configurations profiles.
	KernelExtensionOverridesAllowed *bool `json:"kernelExtensionOverridesAllowed,omitempty"`
	// KernelExtensionAllowedTeamIdentifiers All kernel extensions validly signed by the team identifiers in this list will be allowed to load.
	KernelExtensionAllowedTeamIdentifiers []string `json:"kernelExtensionAllowedTeamIdentifiers,omitempty"`
	// KernelExtensionsAllowed A list of kernel extensions that will be allowed to load. . This collection can contain a maximum of 500 elements.
	KernelExtensionsAllowed []MacOSKernelExtension `json:"kernelExtensionsAllowed,omitempty"`
}
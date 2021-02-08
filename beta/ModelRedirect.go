// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RedirectSingleSignOnExtension Represents an Apple Single Sign-On Extension.
type RedirectSingleSignOnExtension struct {
	// SingleSignOnExtension is the base model of RedirectSingleSignOnExtension
	SingleSignOnExtension
	// Configurations Gets or sets a list of typed key-value pairs used to configure Credential-type profiles. This collection can contain a maximum of 500 elements.
	Configurations []KeyTypedValuePair `json:"configurations,omitempty"`
	// ExtensionIdentifier Gets or sets the bundle ID of the app extension that performs SSO for the specified URLs.
	ExtensionIdentifier *string `json:"extensionIdentifier,omitempty"`
	// TeamIdentifier Gets or sets the team ID of the app extension that performs SSO for the specified URLs.
	TeamIdentifier *string `json:"teamIdentifier,omitempty"`
	// URLPrefixes One or more URL prefixes of identity providers on whose behalf the app extension performs single sign-on. URLs must begin with http:// or https://. All URL prefixes must be unique for all profiles.
	URLPrefixes []string `json:"urlPrefixes,omitempty"`
}
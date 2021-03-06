// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// IosikEv2VpnConfiguration By providing the configurations in this profile you can instruct the iOS device to connect to desired IKEv2 VPN endpoint.
type IosikEv2VpnConfiguration struct {
	// IOSVPNConfiguration is the base model of IosikEv2VpnConfiguration
	IOSVPNConfiguration
	// AllowDefaultChildSecurityAssociationParameters Allows the use of child security association parameters by setting all parameters to the device's default unless explicitly specified.
	AllowDefaultChildSecurityAssociationParameters *bool `json:"allowDefaultChildSecurityAssociationParameters,omitempty"`
	// AllowDefaultSecurityAssociationParameters Allows the use of security association parameters by setting all parameters to the device's default unless explicitly specified.
	AllowDefaultSecurityAssociationParameters *bool `json:"allowDefaultSecurityAssociationParameters,omitempty"`
	// AlwaysOnConfiguration AlwaysOn Configuration
	AlwaysOnConfiguration *AppleVPNAlwaysOnConfiguration `json:"alwaysOnConfiguration,omitempty"`
	// ChildSecurityAssociationParameters Child Security Association Parameters
	ChildSecurityAssociationParameters *IOSVPNSecurityAssociationParameters `json:"childSecurityAssociationParameters,omitempty"`
	// ClientAuthenticationType Type of Client Authentication the VPN client will use.
	ClientAuthenticationType *VPNClientAuthenticationType `json:"clientAuthenticationType,omitempty"`
	// DeadPeerDetectionRate Determine how often to check if a peer connection is still active.
	DeadPeerDetectionRate *VPNDeadPeerDetectionRate `json:"deadPeerDetectionRate,omitempty"`
	// DisableMobilityAndMultihoming Disable MOBIKE
	DisableMobilityAndMultihoming *bool `json:"disableMobilityAndMultihoming,omitempty"`
	// DisableRedirect Disable Redirect
	DisableRedirect *bool `json:"disableRedirect,omitempty"`
	// EnableAlwaysOnConfiguration Determines if Always on VPN is enabled
	EnableAlwaysOnConfiguration *bool `json:"enableAlwaysOnConfiguration,omitempty"`
	// EnableCertificateRevocationCheck Enables a best-effort revocation check; server response timeouts will not cause it to fail
	EnableCertificateRevocationCheck *bool `json:"enableCertificateRevocationCheck,omitempty"`
	// EnableEAP Enables EAP only authentication
	EnableEAP *bool `json:"enableEAP,omitempty"`
	// EnablePerfectForwardSecrecy Enable Perfect Forward Secrecy (PFS).
	EnablePerfectForwardSecrecy *bool `json:"enablePerfectForwardSecrecy,omitempty"`
	// EnableUseInternalSubnetAttributes Enable Use Internal Subnet Attributes.
	EnableUseInternalSubnetAttributes *bool `json:"enableUseInternalSubnetAttributes,omitempty"`
	// LocalIdentifier Method of identifying the client that is trying to connect via VPN.
	LocalIdentifier *VPNLocalIdentifier `json:"localIdentifier,omitempty"`
	// MtuSizeInBytes Maximum transmission unit. Valid values 1280 to 1400
	MtuSizeInBytes *int `json:"mtuSizeInBytes,omitempty"`
	// RemoteIdentifier Address of the IKEv2 server. Must be a FQDN, UserFQDN, network address, or ASN1DN
	RemoteIdentifier *string `json:"remoteIdentifier,omitempty"`
	// SecurityAssociationParameters Security Association Parameters
	SecurityAssociationParameters *IOSVPNSecurityAssociationParameters `json:"securityAssociationParameters,omitempty"`
	// ServerCertificateCommonName Common name of the IKEv2 Server Certificate used in Server Authentication
	ServerCertificateCommonName *string `json:"serverCertificateCommonName,omitempty"`
	// ServerCertificateIssuerCommonName Issuer Common name of the IKEv2 Server Certificate issuer used in Authentication
	ServerCertificateIssuerCommonName *string `json:"serverCertificateIssuerCommonName,omitempty"`
	// ServerCertificateType The type of certificate the VPN server will present to the VPN client for authentication.
	ServerCertificateType *VPNServerCertificateType `json:"serverCertificateType,omitempty"`
	// SharedSecret Used when Shared Secret Authentication is selected
	SharedSecret *string `json:"sharedSecret,omitempty"`
	// TLSMaximumVersion The maximum TLS version to be used with EAP-TLS authentication
	TLSMaximumVersion *string `json:"tlsMaximumVersion,omitempty"`
	// TLSMinimumVersion The minimum TLS version to be used with EAP-TLS authentication
	TLSMinimumVersion *string `json:"tlsMinimumVersion,omitempty"`
}

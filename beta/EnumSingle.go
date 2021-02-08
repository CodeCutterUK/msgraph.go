// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SingleSignOnMode undocumented
type SingleSignOnMode string

const (
	// SingleSignOnModeVNone undocumented
	SingleSignOnModeVNone SingleSignOnMode = "none"
	// SingleSignOnModeVOnPremisesKerberos undocumented
	SingleSignOnModeVOnPremisesKerberos SingleSignOnMode = "onPremisesKerberos"
	// SingleSignOnModeVSaml undocumented
	SingleSignOnModeVSaml SingleSignOnMode = "saml"
	// SingleSignOnModeVPingHeaderBased undocumented
	SingleSignOnModeVPingHeaderBased SingleSignOnMode = "pingHeaderBased"
	// SingleSignOnModeVAadHeaderBased undocumented
	SingleSignOnModeVAadHeaderBased SingleSignOnMode = "aadHeaderBased"
	// SingleSignOnModeVUnknownFutureValue undocumented
	SingleSignOnModeVUnknownFutureValue SingleSignOnMode = "unknownFutureValue"
)

var (
	// SingleSignOnModePNone is a pointer to SingleSignOnModeVNone
	SingleSignOnModePNone = &_SingleSignOnModePNone
	// SingleSignOnModePOnPremisesKerberos is a pointer to SingleSignOnModeVOnPremisesKerberos
	SingleSignOnModePOnPremisesKerberos = &_SingleSignOnModePOnPremisesKerberos
	// SingleSignOnModePSaml is a pointer to SingleSignOnModeVSaml
	SingleSignOnModePSaml = &_SingleSignOnModePSaml
	// SingleSignOnModePPingHeaderBased is a pointer to SingleSignOnModeVPingHeaderBased
	SingleSignOnModePPingHeaderBased = &_SingleSignOnModePPingHeaderBased
	// SingleSignOnModePAadHeaderBased is a pointer to SingleSignOnModeVAadHeaderBased
	SingleSignOnModePAadHeaderBased = &_SingleSignOnModePAadHeaderBased
	// SingleSignOnModePUnknownFutureValue is a pointer to SingleSignOnModeVUnknownFutureValue
	SingleSignOnModePUnknownFutureValue = &_SingleSignOnModePUnknownFutureValue
)

var (
	_SingleSignOnModePNone               = SingleSignOnModeVNone
	_SingleSignOnModePOnPremisesKerberos = SingleSignOnModeVOnPremisesKerberos
	_SingleSignOnModePSaml               = SingleSignOnModeVSaml
	_SingleSignOnModePPingHeaderBased    = SingleSignOnModeVPingHeaderBased
	_SingleSignOnModePAadHeaderBased     = SingleSignOnModeVAadHeaderBased
	_SingleSignOnModePUnknownFutureValue = SingleSignOnModeVUnknownFutureValue
)

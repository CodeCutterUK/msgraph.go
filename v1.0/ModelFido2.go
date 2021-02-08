// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Fido2AuthenticationMethod undocumented
type Fido2AuthenticationMethod struct {
	// AuthenticationMethod is the base model of Fido2AuthenticationMethod
	AuthenticationMethod
	// AaGUID undocumented
	AaGUID *string `json:"aaGuid,omitempty"`
	// AttestationCertificates undocumented
	AttestationCertificates []string `json:"attestationCertificates,omitempty"`
	// AttestationLevel undocumented
	AttestationLevel *AttestationLevel `json:"attestationLevel,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Model undocumented
	Model *string `json:"model,omitempty"`
}

// Fido2AuthenticationMethodConfiguration undocumented
type Fido2AuthenticationMethodConfiguration struct {
	// AuthenticationMethodConfiguration is the base model of Fido2AuthenticationMethodConfiguration
	AuthenticationMethodConfiguration
	// IsAttestationEnforced undocumented
	IsAttestationEnforced *bool `json:"isAttestationEnforced,omitempty"`
	// IsSelfServiceRegistrationAllowed undocumented
	IsSelfServiceRegistrationAllowed *bool `json:"isSelfServiceRegistrationAllowed,omitempty"`
	// KeyRestrictions undocumented
	KeyRestrictions *Fido2KeyRestrictions `json:"keyRestrictions,omitempty"`
	// IncludeTargets undocumented
	IncludeTargets []AuthenticationMethodTarget `json:"includeTargets,omitempty"`
}

// Fido2KeyRestrictions undocumented
type Fido2KeyRestrictions struct {
	// Object is the base model of Fido2KeyRestrictions
	Object
	// AaGuids undocumented
	AaGuids []string `json:"aaGuids,omitempty"`
	// EnforcementType undocumented
	EnforcementType *Fido2RestrictionEnforcementType `json:"enforcementType,omitempty"`
	// IsEnforced undocumented
	IsEnforced *bool `json:"isEnforced,omitempty"`
}

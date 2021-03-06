// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SuggestedEnrollmentLimit The suggestedEnrollmentLimit resource represents the suggested enrollment limit when given an enrollment type.
type SuggestedEnrollmentLimit struct {
	// Object is the base model of SuggestedEnrollmentLimit
	Object
	// SuggestedDailyLimit The suggested enrollment limit within a day
	SuggestedDailyLimit *int `json:"suggestedDailyLimit,omitempty"`
}

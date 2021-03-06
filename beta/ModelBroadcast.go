// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BroadcastMeetingSettings undocumented
type BroadcastMeetingSettings struct {
	// Object is the base model of BroadcastMeetingSettings
	Object
	// AllowedAudience undocumented
	AllowedAudience *BroadcastMeetingAudience `json:"allowedAudience,omitempty"`
	// IsAttendeeReportEnabled undocumented
	IsAttendeeReportEnabled *bool `json:"isAttendeeReportEnabled,omitempty"`
	// IsQuestionAndAnswerEnabled undocumented
	IsQuestionAndAnswerEnabled *bool `json:"isQuestionAndAnswerEnabled,omitempty"`
	// IsRecordingEnabled undocumented
	IsRecordingEnabled *bool `json:"isRecordingEnabled,omitempty"`
	// IsVideoOnDemandEnabled undocumented
	IsVideoOnDemandEnabled *bool `json:"isVideoOnDemandEnabled,omitempty"`
}

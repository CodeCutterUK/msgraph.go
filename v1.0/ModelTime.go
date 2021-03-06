// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// TimeConstraint undocumented
type TimeConstraint struct {
	// Object is the base model of TimeConstraint
	Object
	// ActivityDomain undocumented
	ActivityDomain *ActivityDomain `json:"activityDomain,omitempty"`
	// TimeSlots undocumented
	TimeSlots []TimeSlot `json:"timeSlots,omitempty"`
}

// TimeOff undocumented
type TimeOff struct {
	// ChangeTrackedEntity is the base model of TimeOff
	ChangeTrackedEntity
	// DraftTimeOff undocumented
	DraftTimeOff *TimeOffItem `json:"draftTimeOff,omitempty"`
	// SharedTimeOff undocumented
	SharedTimeOff *TimeOffItem `json:"sharedTimeOff,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

// TimeOffItem undocumented
type TimeOffItem struct {
	// ScheduleEntity is the base model of TimeOffItem
	ScheduleEntity
	// TimeOffReasonID undocumented
	TimeOffReasonID *string `json:"timeOffReasonId,omitempty"`
}

// TimeOffReason undocumented
type TimeOffReason struct {
	// ChangeTrackedEntity is the base model of TimeOffReason
	ChangeTrackedEntity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IconType undocumented
	IconType *TimeOffReasonIconType `json:"iconType,omitempty"`
	// IsActive undocumented
	IsActive *bool `json:"isActive,omitempty"`
}

// TimeOffRequestObject undocumented
type TimeOffRequestObject struct {
	// ScheduleChangeRequestObject is the base model of TimeOffRequestObject
	ScheduleChangeRequestObject
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// TimeOffReasonID undocumented
	TimeOffReasonID *string `json:"timeOffReasonId,omitempty"`
}

// TimeRange undocumented
type TimeRange struct {
	// Object is the base model of TimeRange
	Object
	// EndTime undocumented
	EndTime *TimeOfDay `json:"endTime,omitempty"`
	// StartTime undocumented
	StartTime *TimeOfDay `json:"startTime,omitempty"`
}

// TimeSlot undocumented
type TimeSlot struct {
	// Object is the base model of TimeSlot
	Object
	// End undocumented
	End *DateTimeTimeZone `json:"end,omitempty"`
	// Start undocumented
	Start *DateTimeTimeZone `json:"start,omitempty"`
}

// TimeZoneBase undocumented
type TimeZoneBase struct {
	// Object is the base model of TimeZoneBase
	Object
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

// TimeZoneInformation undocumented
type TimeZoneInformation struct {
	// Object is the base model of TimeZoneInformation
	Object
	// Alias undocumented
	Alias *string `json:"alias,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

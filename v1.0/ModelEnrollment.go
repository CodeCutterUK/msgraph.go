// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EnrollmentConfigurationAssignment Enrollment Configuration Assignment
type EnrollmentConfigurationAssignment struct {
	// Entity is the base model of EnrollmentConfigurationAssignment
	Entity
	// Target Represents an assignment to managed devices in the tenant
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

// EnrollmentTroubleshootingEvent Event representing an enrollment failure.
type EnrollmentTroubleshootingEvent struct {
	// DeviceManagementTroubleshootingEvent is the base model of EnrollmentTroubleshootingEvent
	DeviceManagementTroubleshootingEvent
	// DeviceID Azure AD device identifier.
	DeviceID *string `json:"deviceId,omitempty"`
	// EnrollmentType Type of the enrollment.
	EnrollmentType *DeviceEnrollmentType `json:"enrollmentType,omitempty"`
	// FailureCategory Highlevel failure category.
	FailureCategory *DeviceEnrollmentFailureReason `json:"failureCategory,omitempty"`
	// FailureReason Detailed failure reason.
	FailureReason *string `json:"failureReason,omitempty"`
	// ManagedDeviceIdentifier Device identifier created or collected by Intune.
	ManagedDeviceIdentifier *string `json:"managedDeviceIdentifier,omitempty"`
	// OperatingSystem Operating System.
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// OsVersion OS Version.
	OsVersion *string `json:"osVersion,omitempty"`
	// UserID Identifier for the user that tried to enroll the device.
	UserID *string `json:"userId,omitempty"`
}

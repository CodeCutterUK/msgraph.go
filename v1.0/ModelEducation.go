// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// EducationClass undocumented
type EducationClass struct {
	// Entity is the base model of EducationClass
	Entity
	// ClassCode undocumented
	ClassCode *string `json:"classCode,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// ExternalName undocumented
	ExternalName *string `json:"externalName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// Term undocumented
	Term *EducationTerm `json:"term,omitempty"`
	// Group undocumented
	Group *Group `json:"group,omitempty"`
	// Members undocumented
	Members []EducationUser `json:"members,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// Teachers undocumented
	Teachers []EducationUser `json:"teachers,omitempty"`
}

// EducationOrganization undocumented
type EducationOrganization struct {
	// Entity is the base model of EducationOrganization
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
}

// EducationRoot undocumented
type EducationRoot struct {
	// Entity is the base model of EducationRoot
	Entity
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Me undocumented
	Me *EducationUser `json:"me,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// Users undocumented
	Users []EducationUser `json:"users,omitempty"`
}

// EducationSchool undocumented
type EducationSchool struct {
	// EducationOrganization is the base model of EducationSchool
	EducationOrganization
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// ExternalPrincipalID undocumented
	ExternalPrincipalID *string `json:"externalPrincipalId,omitempty"`
	// Fax undocumented
	Fax *string `json:"fax,omitempty"`
	// HighestGrade undocumented
	HighestGrade *string `json:"highestGrade,omitempty"`
	// LowestGrade undocumented
	LowestGrade *string `json:"lowestGrade,omitempty"`
	// Phone undocumented
	Phone *string `json:"phone,omitempty"`
	// PrincipalEmail undocumented
	PrincipalEmail *string `json:"principalEmail,omitempty"`
	// PrincipalName undocumented
	PrincipalName *string `json:"principalName,omitempty"`
	// SchoolNumber undocumented
	SchoolNumber *string `json:"schoolNumber,omitempty"`
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Users undocumented
	Users []EducationUser `json:"users,omitempty"`
}

// EducationStudent undocumented
type EducationStudent struct {
	// Object is the base model of EducationStudent
	Object
	// BirthDate undocumented
	BirthDate *Date `json:"birthDate,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// Gender undocumented
	Gender *EducationGender `json:"gender,omitempty"`
	// Grade undocumented
	Grade *string `json:"grade,omitempty"`
	// GraduationYear undocumented
	GraduationYear *string `json:"graduationYear,omitempty"`
	// StudentNumber undocumented
	StudentNumber *string `json:"studentNumber,omitempty"`
}

// EducationTeacher undocumented
type EducationTeacher struct {
	// Object is the base model of EducationTeacher
	Object
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// TeacherNumber undocumented
	TeacherNumber *string `json:"teacherNumber,omitempty"`
}

// EducationTerm undocumented
type EducationTerm struct {
	// Object is the base model of EducationTerm
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDate undocumented
	EndDate *Date `json:"endDate,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// StartDate undocumented
	StartDate *Date `json:"startDate,omitempty"`
}

// EducationUser undocumented
type EducationUser struct {
	// Entity is the base model of EducationUser
	Entity
	// AccountEnabled undocumented
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// AssignedLicenses undocumented
	AssignedLicenses []AssignedLicense `json:"assignedLicenses,omitempty"`
	// AssignedPlans undocumented
	AssignedPlans []AssignedPlan `json:"assignedPlans,omitempty"`
	// BusinessPhones undocumented
	BusinessPhones []string `json:"businessPhones,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailingAddress undocumented
	MailingAddress *PhysicalAddress `json:"mailingAddress,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// MiddleName undocumented
	MiddleName *string `json:"middleName,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// PasswordPolicies undocumented
	PasswordPolicies *string `json:"passwordPolicies,omitempty"`
	// PasswordProfile undocumented
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`
	// PreferredLanguage undocumented
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// PrimaryRole undocumented
	PrimaryRole *EducationUserRole `json:"primaryRole,omitempty"`
	// ProvisionedPlans undocumented
	ProvisionedPlans []ProvisionedPlan `json:"provisionedPlans,omitempty"`
	// RefreshTokensValidFromDateTime undocumented
	RefreshTokensValidFromDateTime *time.Time `json:"refreshTokensValidFromDateTime,omitempty"`
	// ResidenceAddress undocumented
	ResidenceAddress *PhysicalAddress `json:"residenceAddress,omitempty"`
	// ShowInAddressList undocumented
	ShowInAddressList *bool `json:"showInAddressList,omitempty"`
	// Student undocumented
	Student *EducationStudent `json:"student,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// Teacher undocumented
	Teacher *EducationTeacher `json:"teacher,omitempty"`
	// UsageLocation undocumented
	UsageLocation *string `json:"usageLocation,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserType undocumented
	UserType *string `json:"userType,omitempty"`
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// User undocumented
	User *User `json:"user,omitempty"`
}

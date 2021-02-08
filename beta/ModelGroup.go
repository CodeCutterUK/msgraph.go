// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Group undocumented
type Group struct {
	// DirectoryObject is the base model of Group
	DirectoryObject
	// AssignedLabels undocumented
	AssignedLabels []AssignedLabel `json:"assignedLabels,omitempty"`
	// AssignedLicenses undocumented
	AssignedLicenses []AssignedLicense `json:"assignedLicenses,omitempty"`
	// Classification undocumented
	Classification *string `json:"classification,omitempty"`
	// CreatedByAppID undocumented
	CreatedByAppID *string `json:"createdByAppId,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// GroupTypes undocumented
	GroupTypes []string `json:"groupTypes,omitempty"`
	// HasMembersWithLicenseErrors undocumented
	HasMembersWithLicenseErrors *bool `json:"hasMembersWithLicenseErrors,omitempty"`
	// InfoCatalogs undocumented
	InfoCatalogs []string `json:"infoCatalogs,omitempty"`
	// IsAssignableToRole undocumented
	IsAssignableToRole *bool `json:"isAssignableToRole,omitempty"`
	// LicenseProcessingState undocumented
	LicenseProcessingState *LicenseProcessingState `json:"licenseProcessingState,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailEnabled undocumented
	MailEnabled *bool `json:"mailEnabled,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// MDMAppID undocumented
	MDMAppID *string `json:"mdmAppId,omitempty"`
	// MembershipRule undocumented
	MembershipRule *string `json:"membershipRule,omitempty"`
	// MembershipRuleProcessingState undocumented
	MembershipRuleProcessingState *string `json:"membershipRuleProcessingState,omitempty"`
	// OnPremisesDomainName undocumented
	OnPremisesDomainName *string `json:"onPremisesDomainName,omitempty"`
	// OnPremisesLastSyncDateTime undocumented
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// OnPremisesNetBiosName undocumented
	OnPremisesNetBiosName *string `json:"onPremisesNetBiosName,omitempty"`
	// OnPremisesProvisioningErrors undocumented
	OnPremisesProvisioningErrors []OnPremisesProvisioningError `json:"onPremisesProvisioningErrors,omitempty"`
	// OnPremisesSamAccountName undocumented
	OnPremisesSamAccountName *string `json:"onPremisesSamAccountName,omitempty"`
	// OnPremisesSecurityIdentifier undocumented
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// OnPremisesSyncEnabled undocumented
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// PreferredDataLocation undocumented
	PreferredDataLocation *string `json:"preferredDataLocation,omitempty"`
	// PreferredLanguage undocumented
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// ProxyAddresses undocumented
	ProxyAddresses []string `json:"proxyAddresses,omitempty"`
	// RenewedDateTime undocumented
	RenewedDateTime *time.Time `json:"renewedDateTime,omitempty"`
	// ResourceBehaviorOptions undocumented
	ResourceBehaviorOptions []string `json:"resourceBehaviorOptions,omitempty"`
	// ResourceProvisioningOptions undocumented
	ResourceProvisioningOptions []string `json:"resourceProvisioningOptions,omitempty"`
	// SecurityEnabled undocumented
	SecurityEnabled *bool `json:"securityEnabled,omitempty"`
	// SecurityIdentifier undocumented
	SecurityIdentifier *string `json:"securityIdentifier,omitempty"`
	// Theme undocumented
	Theme *string `json:"theme,omitempty"`
	// Visibility undocumented
	Visibility *string `json:"visibility,omitempty"`
	// AccessType undocumented
	AccessType *GroupAccessType `json:"accessType,omitempty"`
	// AllowExternalSenders undocumented
	AllowExternalSenders *bool `json:"allowExternalSenders,omitempty"`
	// AutoSubscribeNewMembers undocumented
	AutoSubscribeNewMembers *bool `json:"autoSubscribeNewMembers,omitempty"`
	// HideFromAddressLists undocumented
	HideFromAddressLists *bool `json:"hideFromAddressLists,omitempty"`
	// HideFromOutlookClients undocumented
	HideFromOutlookClients *bool `json:"hideFromOutlookClients,omitempty"`
	// IsFavorite undocumented
	IsFavorite *bool `json:"isFavorite,omitempty"`
	// IsSubscribedByMail undocumented
	IsSubscribedByMail *bool `json:"isSubscribedByMail,omitempty"`
	// UnseenConversationsCount undocumented
	UnseenConversationsCount *int `json:"unseenConversationsCount,omitempty"`
	// UnseenCount undocumented
	UnseenCount *int `json:"unseenCount,omitempty"`
	// UnseenMessagesCount undocumented
	UnseenMessagesCount *int `json:"unseenMessagesCount,omitempty"`
	// MembershipRuleProcessingStatus undocumented
	MembershipRuleProcessingStatus *MembershipRuleProcessingStatus `json:"membershipRuleProcessingStatus,omitempty"`
	// IsArchived undocumented
	IsArchived *bool `json:"isArchived,omitempty"`
	// AppRoleAssignments undocumented
	AppRoleAssignments []AppRoleAssignment `json:"appRoleAssignments,omitempty"`
	// CreatedOnBehalfOf undocumented
	CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`
	// Endpoints undocumented
	Endpoints []Endpoint `json:"endpoints,omitempty"`
	// MemberOf undocumented
	MemberOf []DirectoryObject `json:"memberOf,omitempty"`
	// Members undocumented
	Members []DirectoryObject `json:"members,omitempty"`
	// MembersWithLicenseErrors undocumented
	MembersWithLicenseErrors []DirectoryObject `json:"membersWithLicenseErrors,omitempty"`
	// Owners undocumented
	Owners []DirectoryObject `json:"owners,omitempty"`
	// PermissionGrants undocumented
	PermissionGrants []ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
	// Settings undocumented
	Settings []DirectorySetting `json:"settings,omitempty"`
	// TransitiveMemberOf undocumented
	TransitiveMemberOf []DirectoryObject `json:"transitiveMemberOf,omitempty"`
	// TransitiveMembers undocumented
	TransitiveMembers []DirectoryObject `json:"transitiveMembers,omitempty"`
	// AcceptedSenders undocumented
	AcceptedSenders []DirectoryObject `json:"acceptedSenders,omitempty"`
	// Calendar undocumented
	Calendar *Calendar `json:"calendar,omitempty"`
	// CalendarView undocumented
	CalendarView []Event `json:"calendarView,omitempty"`
	// Conversations undocumented
	Conversations []Conversation `json:"conversations,omitempty"`
	// Events undocumented
	Events []Event `json:"events,omitempty"`
	// Photo undocumented
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// Photos undocumented
	Photos []ProfilePhoto `json:"photos,omitempty"`
	// RejectedSenders undocumented
	RejectedSenders []DirectoryObject `json:"rejectedSenders,omitempty"`
	// Threads undocumented
	Threads []ConversationThread `json:"threads,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Drives undocumented
	Drives []Drive `json:"drives,omitempty"`
	// Sites undocumented
	Sites []Site `json:"sites,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
	// GroupLifecyclePolicies undocumented
	GroupLifecyclePolicies []GroupLifecyclePolicy `json:"groupLifecyclePolicies,omitempty"`
	// Planner undocumented
	Planner *PlannerGroup `json:"planner,omitempty"`
	// Onenote undocumented
	Onenote *Onenote `json:"onenote,omitempty"`
	// Team undocumented
	Team *Team `json:"team,omitempty"`
}

// GroupAssignmentTarget Represents an assignment to a group.
type GroupAssignmentTarget struct {
	// DeviceAndAppManagementAssignmentTarget is the base model of GroupAssignmentTarget
	DeviceAndAppManagementAssignmentTarget
	// GroupID The group Id that is the target of the assignment.
	GroupID *string `json:"groupId,omitempty"`
}

// GroupLifecyclePolicy undocumented
type GroupLifecyclePolicy struct {
	// Entity is the base model of GroupLifecyclePolicy
	Entity
	// AlternateNotificationEmails undocumented
	AlternateNotificationEmails *string `json:"alternateNotificationEmails,omitempty"`
	// GroupLifetimeInDays undocumented
	GroupLifetimeInDays *int `json:"groupLifetimeInDays,omitempty"`
	// ManagedGroupTypes undocumented
	ManagedGroupTypes *string `json:"managedGroupTypes,omitempty"`
}

// GroupMembers undocumented
type GroupMembers struct {
	// UserSet is the base model of GroupMembers
	UserSet
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

// GroupMembershipGovernanceCriteria undocumented
type GroupMembershipGovernanceCriteria struct {
	// GovernanceCriteria is the base model of GroupMembershipGovernanceCriteria
	GovernanceCriteria
	// GroupID undocumented
	GroupID *string `json:"groupId,omitempty"`
}

// GroupPolicyCategory The category entity stores the category of a group policy definition
type GroupPolicyCategory struct {
	// Entity is the base model of GroupPolicyCategory
	Entity
	// DisplayName The string id of the category's display name
	DisplayName *string `json:"displayName,omitempty"`
	// IsRoot Defines if the category is a root category
	IsRoot *bool `json:"isRoot,omitempty"`
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Children undocumented
	Children []GroupPolicyCategory `json:"children,omitempty"`
	// DefinitionFile undocumented
	DefinitionFile *GroupPolicyDefinitionFile `json:"definitionFile,omitempty"`
	// Definitions undocumented
	Definitions []GroupPolicyDefinition `json:"definitions,omitempty"`
	// Parent undocumented
	Parent *GroupPolicyCategory `json:"parent,omitempty"`
}

// GroupPolicyConfiguration The group policy configuration entity contains the configured values for one or more group policy definitions.
type GroupPolicyConfiguration struct {
	// Entity is the base model of GroupPolicyConfiguration
	Entity
	// CreatedDateTime The date and time the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description User provided description for the resource object.
	Description *string `json:"description,omitempty"`
	// DisplayName User provided name for the resource object.
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// RoleScopeTagIDs The list of scope tags for the configuration.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// Assignments undocumented
	Assignments []GroupPolicyConfigurationAssignment `json:"assignments,omitempty"`
	// DefinitionValues undocumented
	DefinitionValues []GroupPolicyDefinitionValue `json:"definitionValues,omitempty"`
}

// GroupPolicyConfigurationAssignment The group policy configuration assignment entity assigns one or more AAD groups to a specific group policy configuration.
type GroupPolicyConfigurationAssignment struct {
	// Entity is the base model of GroupPolicyConfigurationAssignment
	Entity
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Target The type of groups targeted the group policy configuration.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

// GroupPolicyDefinition The entity describes all of the information about a single group policy.
type GroupPolicyDefinition struct {
	// Entity is the base model of GroupPolicyDefinition
	Entity
	// CategoryPath The localized full category path for the policy.
	CategoryPath *string `json:"categoryPath,omitempty"`
	// ClassType Identifies the type of groups the policy can be applied to.
	ClassType *GroupPolicyDefinitionClassType `json:"classType,omitempty"`
	// DisplayName The localized policy name.
	DisplayName *string `json:"displayName,omitempty"`
	// ExplainText The localized explanation or help text associated with the policy. The default value is empty.
	ExplainText *string `json:"explainText,omitempty"`
	// GroupPolicyCategoryID The category id of the parent category
	GroupPolicyCategoryID *UUID `json:"groupPolicyCategoryId,omitempty"`
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// PolicyType Specifies the type of group policy.
	PolicyType *GroupPolicyType `json:"policyType,omitempty"`
	// SupportedOn Localized string used to specify what operating system or application version is affected by the policy.
	SupportedOn *string `json:"supportedOn,omitempty"`
	// Category undocumented
	Category *GroupPolicyCategory `json:"category,omitempty"`
	// DefinitionFile undocumented
	DefinitionFile *GroupPolicyDefinitionFile `json:"definitionFile,omitempty"`
	// Presentations undocumented
	Presentations []GroupPolicyPresentation `json:"presentations,omitempty"`
}

// GroupPolicyDefinitionFile The entity represents an ADMX (Administrative Template) XML file. The ADMX file contains a collection of group policy definitions and their locations by category path. The group policy definition file also contains the languages supported as determined by the language dependent ADML (Administrative Template) language files.
type GroupPolicyDefinitionFile struct {
	// Entity is the base model of GroupPolicyDefinitionFile
	Entity
	// Description The localized description of the policy settings in the ADMX file. The default value is empty.
	Description *string `json:"description,omitempty"`
	// DisplayName The localized friendly name of the ADMX file.
	DisplayName *string `json:"displayName,omitempty"`
	// LanguageCodes The supported language codes for the ADMX file.
	LanguageCodes []string `json:"languageCodes,omitempty"`
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// PolicyType Specifies the type of group policy.
	PolicyType *GroupPolicyType `json:"policyType,omitempty"`
	// Revision The revision version associated with the file.
	Revision *string `json:"revision,omitempty"`
	// TargetNamespace Specifies the URI used to identify the namespace within the ADMX file.
	TargetNamespace *string `json:"targetNamespace,omitempty"`
	// TargetPrefix Specifies the logical name that refers to the namespace within the ADMX file.
	TargetPrefix *string `json:"targetPrefix,omitempty"`
	// Definitions undocumented
	Definitions []GroupPolicyDefinition `json:"definitions,omitempty"`
}

// GroupPolicyDefinitionValue The definition value entity stores the value for a single group policy definition.
type GroupPolicyDefinitionValue struct {
	// Entity is the base model of GroupPolicyDefinitionValue
	Entity
	// ConfigurationType Specifies how the value should be configured. This can be either as a Policy or as a Preference.
	ConfigurationType *GroupPolicyConfigurationType `json:"configurationType,omitempty"`
	// CreatedDateTime The date and time the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Enabled Enables or disables the associated group policy definition.
	Enabled *bool `json:"enabled,omitempty"`
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Definition undocumented
	Definition *GroupPolicyDefinition `json:"definition,omitempty"`
	// PresentationValues undocumented
	PresentationValues []GroupPolicyPresentationValue `json:"presentationValues,omitempty"`
}

// GroupPolicyMigrationReport The Group Policy migration report.
type GroupPolicyMigrationReport struct {
	// Entity is the base model of GroupPolicyMigrationReport
	Entity
	// CreatedDateTime The date and time at which the GroupPolicyMigrationReport was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName The name of Group Policy Object from the GPO Xml Content
	DisplayName *string `json:"displayName,omitempty"`
	// GroupPolicyCreatedDateTime The date and time at which the GroupPolicyMigrationReport was created.
	GroupPolicyCreatedDateTime *time.Time `json:"groupPolicyCreatedDateTime,omitempty"`
	// GroupPolicyLastModifiedDateTime The date and time at which the GroupPolicyMigrationReport was last modified.
	GroupPolicyLastModifiedDateTime *time.Time `json:"groupPolicyLastModifiedDateTime,omitempty"`
	// GroupPolicyObjectID The Group Policy Object GUID from GPO Xml content
	GroupPolicyObjectID *UUID `json:"groupPolicyObjectId,omitempty"`
	// LastModifiedDateTime The date and time at which the GroupPolicyMigrationReport was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// MigrationReadiness The Intune coverage for the associated Group Policy Object file.
	MigrationReadiness *GroupPolicyMigrationReadiness `json:"migrationReadiness,omitempty"`
	// OuDistinguishedName The distinguished name of the OU.
	OuDistinguishedName *string `json:"ouDistinguishedName,omitempty"`
	// SupportedSettingsCount The number of Group Policy Settings supported by Intune.
	SupportedSettingsCount *int `json:"supportedSettingsCount,omitempty"`
	// SupportedSettingsPercent The Percentage of Group Policy Settings supported by Intune.
	SupportedSettingsPercent *int `json:"supportedSettingsPercent,omitempty"`
	// TargetedInActiveDirectory The Targeted in AD property from GPO Xml Content
	TargetedInActiveDirectory *bool `json:"targetedInActiveDirectory,omitempty"`
	// TotalSettingsCount The total number of Group Policy Settings from GPO file.
	TotalSettingsCount *int `json:"totalSettingsCount,omitempty"`
	// GroupPolicySettingMappings undocumented
	GroupPolicySettingMappings []GroupPolicySettingMapping `json:"groupPolicySettingMappings,omitempty"`
	// UnsupportedGroupPolicyExtensions undocumented
	UnsupportedGroupPolicyExtensions []UnsupportedGroupPolicyExtension `json:"unsupportedGroupPolicyExtensions,omitempty"`
}

// GroupPolicyObjectFile The Group Policy Object file uploaded by admin.
type GroupPolicyObjectFile struct {
	// Entity is the base model of GroupPolicyObjectFile
	Entity
	// Content The Group Policy Object file content.
	Content *string `json:"content,omitempty"`
	// CreatedDateTime The date and time at which the GroupPolicy was first uploaded.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// GroupPolicyObjectID The Group Policy Object GUID from GPO Xml content
	GroupPolicyObjectID *UUID `json:"groupPolicyObjectId,omitempty"`
	// LastModifiedDateTime The date and time at which the GroupPolicyObjectFile was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// OuDistinguishedName The distinguished name of the OU.
	OuDistinguishedName *string `json:"ouDistinguishedName,omitempty"`
}

// GroupPolicyOperation The entity represents an group policy operation.
type GroupPolicyOperation struct {
	// Entity is the base model of GroupPolicyOperation
	Entity
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// OperationStatus The group policy operation status.
	OperationStatus *GroupPolicyOperationStatus `json:"operationStatus,omitempty"`
	// OperationType The type of group policy operation.
	OperationType *GroupPolicyOperationType `json:"operationType,omitempty"`
	// StatusDetails The group policy operation status detail.
	StatusDetails *string `json:"statusDetails,omitempty"`
}

// GroupPolicyPresentation The base entity for the display presentation of any of the additional options in a group policy definition.
type GroupPolicyPresentation struct {
	// Entity is the base model of GroupPolicyPresentation
	Entity
	// Label Localized text label for any presentation entity. The default value is empty.
	Label *string `json:"label,omitempty"`
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Definition undocumented
	Definition *GroupPolicyDefinition `json:"definition,omitempty"`
}

// GroupPolicyPresentationCheckBox Represents an ADMX checkBox element and an ADMX boolean element.
type GroupPolicyPresentationCheckBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationCheckBox
	GroupPolicyPresentation
	// DefaultChecked Default value for the check box. The default value is false.
	DefaultChecked *bool `json:"defaultChecked,omitempty"`
}

// GroupPolicyPresentationComboBox Represents an ADMX comboBox element and an ADMX text element.
type GroupPolicyPresentationComboBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationComboBox
	GroupPolicyPresentation
	// DefaultValue Localized default string displayed in the combo box. The default value is empty.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// MaxLength An unsigned integer that specifies the maximum number of text characters for the parameter. The default value is 1023.
	MaxLength *int `json:"maxLength,omitempty"`
	// Required Specifies whether a value must be specified for the parameter. The default value is false.
	Required *bool `json:"required,omitempty"`
	// Suggestions Localized strings listed in the drop-down list of the combo box. The default value is empty.
	Suggestions []string `json:"suggestions,omitempty"`
}

// GroupPolicyPresentationDecimalTextBox Represents an ADMX decimalTextBox element and an ADMX decimal element.
type GroupPolicyPresentationDecimalTextBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationDecimalTextBox
	GroupPolicyPresentation
	// DefaultValue An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
	DefaultValue *int `json:"defaultValue,omitempty"`
	// MaxValue An unsigned integer that specifies the maximum allowed value. The default value is 9999.
	MaxValue *int `json:"maxValue,omitempty"`
	// MinValue An unsigned integer that specifies the minimum allowed value. The default value is 0.
	MinValue *int `json:"minValue,omitempty"`
	// Required Requirement to enter a value in the parameter box. The default value is false.
	Required *bool `json:"required,omitempty"`
	// Spin If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
	Spin *bool `json:"spin,omitempty"`
	// SpinStep An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
	SpinStep *int `json:"spinStep,omitempty"`
}

// GroupPolicyPresentationDropdownList Represents an ADMX dropdownList element and an ADMX enum element.
type GroupPolicyPresentationDropdownList struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationDropdownList
	GroupPolicyPresentation
	// DefaultItem Localized string value identifying the default choice of the list of items.
	DefaultItem *GroupPolicyPresentationDropdownListItem `json:"defaultItem,omitempty"`
	// Items Represents a set of localized display names and their associated values.
	Items []GroupPolicyPresentationDropdownListItem `json:"items,omitempty"`
	// Required Requirement to enter a value in the parameter box. The default value is false.
	Required *bool `json:"required,omitempty"`
}

// GroupPolicyPresentationDropdownListItem undocumented
type GroupPolicyPresentationDropdownListItem struct {
	// Object is the base model of GroupPolicyPresentationDropdownListItem
	Object
	// DisplayName Localized display name for the drop-down list item.
	DisplayName *string `json:"displayName,omitempty"`
	// Value Associated value for the drop-down list item
	Value *string `json:"value,omitempty"`
}

// GroupPolicyPresentationListBox Represents an ADMX listBox element and an ADMX list element.
type GroupPolicyPresentationListBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationListBox
	GroupPolicyPresentation
	// ExplicitValue If this option is specified true the user must specify the registry subkey value and the registry subkey name. The list box shows two columns, one for the name and one for the data. The default value is false.
	ExplicitValue *bool `json:"explicitValue,omitempty"`
	// ValuePrefix undocumented
	ValuePrefix *string `json:"valuePrefix,omitempty"`
}

// GroupPolicyPresentationLongDecimalTextBox Represents an ADMX longDecimalTextBox element and an ADMX longDecimal element.
type GroupPolicyPresentationLongDecimalTextBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationLongDecimalTextBox
	GroupPolicyPresentation
	// DefaultValue An unsigned integer that specifies the initial value for the decimal text box. The default value is 1.
	DefaultValue *int `json:"defaultValue,omitempty"`
	// MaxValue An unsigned long that specifies the maximum allowed value. The default value is 9999.
	MaxValue *int `json:"maxValue,omitempty"`
	// MinValue An unsigned long that specifies the minimum allowed value. The default value is 0.
	MinValue *int `json:"minValue,omitempty"`
	// Required Requirement to enter a value in the parameter box. The default value is false.
	Required *bool `json:"required,omitempty"`
	// Spin If true, create a spin control; otherwise, create a text box for numeric entry. The default value is true.
	Spin *bool `json:"spin,omitempty"`
	// SpinStep An unsigned integer that specifies the increment of change for the spin control. The default value is 1.
	SpinStep *int `json:"spinStep,omitempty"`
}

// GroupPolicyPresentationMultiTextBox Represents an ADMX multiTextBox element and an ADMX multiText element.
type GroupPolicyPresentationMultiTextBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationMultiTextBox
	GroupPolicyPresentation
	// MaxLength An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
	MaxLength *int `json:"maxLength,omitempty"`
	// MaxStrings An unsigned integer that specifies the maximum number of strings. Default value is 0.
	MaxStrings *int `json:"maxStrings,omitempty"`
	// Required Requirement to enter a value in the text box. Default value is false.
	Required *bool `json:"required,omitempty"`
}

// GroupPolicyPresentationText Represents an ADMX text element.
type GroupPolicyPresentationText struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationText
	GroupPolicyPresentation
}

// GroupPolicyPresentationTextBox Represents an ADMX textBox element and an ADMX text element.
type GroupPolicyPresentationTextBox struct {
	// GroupPolicyPresentation is the base model of GroupPolicyPresentationTextBox
	GroupPolicyPresentation
	// DefaultValue Localized default string displayed in the text box. The default value is empty.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// MaxLength An unsigned integer that specifies the maximum number of text characters. Default value is 1023.
	MaxLength *int `json:"maxLength,omitempty"`
	// Required Requirement to enter a value in the text box. Default value is false.
	Required *bool `json:"required,omitempty"`
}

// GroupPolicyPresentationValue The base presentation value entity that stores the value for a single group policy presentation.
type GroupPolicyPresentationValue struct {
	// Entity is the base model of GroupPolicyPresentationValue
	Entity
	// CreatedDateTime The date and time the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime The date and time the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// DefinitionValue undocumented
	DefinitionValue *GroupPolicyDefinitionValue `json:"definitionValue,omitempty"`
	// Presentation undocumented
	Presentation *GroupPolicyPresentation `json:"presentation,omitempty"`
}

// GroupPolicyPresentationValueBoolean The entity represents a Boolean value of a checkbox presentation on a policy definition.
type GroupPolicyPresentationValueBoolean struct {
	// GroupPolicyPresentationValue is the base model of GroupPolicyPresentationValueBoolean
	GroupPolicyPresentationValue
	// Value An boolean value for the associated presentation.
	Value *bool `json:"value,omitempty"`
}

// GroupPolicyPresentationValueDecimal The entity represents an unsigned integer value of a decimal text box presentation on a policy definition.
type GroupPolicyPresentationValueDecimal struct {
	// GroupPolicyPresentationValue is the base model of GroupPolicyPresentationValueDecimal
	GroupPolicyPresentationValue
	// Value An unsigned integer value for the associated presentation.
	Value *int `json:"value,omitempty"`
}

// GroupPolicyPresentationValueList The entity represents a collection of name/value pairs of a list box presentation on a policy definition.
type GroupPolicyPresentationValueList struct {
	// GroupPolicyPresentationValue is the base model of GroupPolicyPresentationValueList
	GroupPolicyPresentationValue
	// Values A list of pairs for the associated presentation.
	Values []KeyValuePair `json:"values,omitempty"`
}

// GroupPolicyPresentationValueLongDecimal The entity represents an unsigned long value of a long decimal text box presentation on a policy definition.
type GroupPolicyPresentationValueLongDecimal struct {
	// GroupPolicyPresentationValue is the base model of GroupPolicyPresentationValueLongDecimal
	GroupPolicyPresentationValue
	// Value An unsigned long value for the associated presentation.
	Value *int `json:"value,omitempty"`
}

// GroupPolicyPresentationValueMultiText The entity represents a string value of a multi-line text box presentation on a policy definition.
type GroupPolicyPresentationValueMultiText struct {
	// GroupPolicyPresentationValue is the base model of GroupPolicyPresentationValueMultiText
	GroupPolicyPresentationValue
	// Values A collection of non-empty strings for the associated presentation.
	Values []string `json:"values,omitempty"`
}

// GroupPolicyPresentationValueText The entity represents a string value for a drop-down list, combo box, or text box presentation on a policy definition.
type GroupPolicyPresentationValueText struct {
	// GroupPolicyPresentationValue is the base model of GroupPolicyPresentationValueText
	GroupPolicyPresentationValue
	// Value A string value for the associated presentation.
	Value *string `json:"value,omitempty"`
}

// GroupPolicySettingMapping The Group Policy setting to MDM/Intune mapping.
type GroupPolicySettingMapping struct {
	// Entity is the base model of GroupPolicySettingMapping
	Entity
	// AdmxSettingDefinitionID Admx Group Policy Id
	AdmxSettingDefinitionID *string `json:"admxSettingDefinitionId,omitempty"`
	// ChildIDList List of Child Ids of the group policy setting.
	ChildIDList []string `json:"childIdList,omitempty"`
	// IntuneSettingDefinitionID The Intune Setting Definition Id
	IntuneSettingDefinitionID *string `json:"intuneSettingDefinitionId,omitempty"`
	// IntuneSettingURIList The list of Intune Setting URIs this group policy setting maps to
	IntuneSettingURIList []string `json:"intuneSettingUriList,omitempty"`
	// IsMDMSupported Indicates if the setting is supported by Intune or not
	IsMDMSupported *bool `json:"isMdmSupported,omitempty"`
	// MDMCspName The CSP name this group policy setting maps to.
	MDMCspName *string `json:"mdmCspName,omitempty"`
	// MDMMinimumOSVersion The minimum OS version this mdm setting supports.
	MDMMinimumOSVersion *int `json:"mdmMinimumOSVersion,omitempty"`
	// MDMSettingURI The MDM CSP URI this group policy setting maps to.
	MDMSettingURI *string `json:"mdmSettingUri,omitempty"`
	// MDMSupportedState Indicates if the setting is supported in Mdm or not
	MDMSupportedState *MDMSupportedState `json:"mdmSupportedState,omitempty"`
	// ParentID Parent Id of the group policy setting.
	ParentID *string `json:"parentId,omitempty"`
	// SettingCategory The category the group policy setting is in.
	SettingCategory *string `json:"settingCategory,omitempty"`
	// SettingDisplayName The display name of this group policy setting.
	SettingDisplayName *string `json:"settingDisplayName,omitempty"`
	// SettingDisplayValue The display value of this group policy setting.
	SettingDisplayValue *string `json:"settingDisplayValue,omitempty"`
	// SettingDisplayValueType The display value type of this group policy setting.
	SettingDisplayValueType *string `json:"settingDisplayValueType,omitempty"`
	// SettingName The name of this group policy setting.
	SettingName *string `json:"settingName,omitempty"`
	// SettingScope The scope of the setting
	SettingScope *GroupPolicySettingScope `json:"settingScope,omitempty"`
	// SettingType The setting type (security or admx) of the Group Policy.
	SettingType *GroupPolicySettingType `json:"settingType,omitempty"`
	// SettingValue The value of this group policy setting.
	SettingValue *string `json:"settingValue,omitempty"`
	// SettingValueDisplayUnits The display units of this group policy setting value
	SettingValueDisplayUnits *string `json:"settingValueDisplayUnits,omitempty"`
	// SettingValueType The value type of this group policy setting.
	SettingValueType *string `json:"settingValueType,omitempty"`
}

// GroupPolicyUploadedDefinitionFile The entity represents an ADMX (Administrative Template) XML file uploaded by Administrator. The ADMX file contains a collection of group policy definitions and their locations by category path. The group policy definition file also contains the languages supported as determined by the language dependent ADML (Administrative Template) language files.
type GroupPolicyUploadedDefinitionFile struct {
	// GroupPolicyDefinitionFile is the base model of GroupPolicyUploadedDefinitionFile
	GroupPolicyDefinitionFile
	// Content The contents of the uploaded ADMX file.
	Content *Binary `json:"content,omitempty"`
	// DefaultLanguageCode The default language of the uploaded ADMX file.
	DefaultLanguageCode *string `json:"defaultLanguageCode,omitempty"`
	// FileName The file name of the uploaded ADML file.
	FileName *string `json:"fileName,omitempty"`
	// GroupPolicyUploadedLanguageFiles The list of ADML files associated with the uploaded ADMX file.
	GroupPolicyUploadedLanguageFiles []GroupPolicyUploadedLanguageFile `json:"groupPolicyUploadedLanguageFiles,omitempty"`
	// Status The upload status of the uploaded ADMX file.
	Status *GroupPolicyUploadedDefinitionFileStatus `json:"status,omitempty"`
	// UploadDateTime The uploaded time of the uploaded ADMX file.
	UploadDateTime *time.Time `json:"uploadDateTime,omitempty"`
	// GroupPolicyOperations undocumented
	GroupPolicyOperations []GroupPolicyOperation `json:"groupPolicyOperations,omitempty"`
}

// GroupPolicyUploadedLanguageFile The entity represents an ADML (Administrative Template language) XML file uploaded by Administrator.
type GroupPolicyUploadedLanguageFile struct {
	// Object is the base model of GroupPolicyUploadedLanguageFile
	Object
	// Content The contents of the uploaded ADML file.
	Content *Binary `json:"content,omitempty"`
	// FileName The file name of the uploaded ADML file.
	FileName *string `json:"fileName,omitempty"`
	// ID Key of the entity.
	ID *string `json:"id,omitempty"`
	// LanguageCode The language code of the uploaded ADML file.
	LanguageCode *string `json:"languageCode,omitempty"`
	// LastModifiedDateTime The date and time the entity was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// GroupPrintUsageSummary undocumented
type GroupPrintUsageSummary struct {
	// Object is the base model of GroupPrintUsageSummary
	Object
	// CompletedJobCount undocumented
	CompletedJobCount *int `json:"completedJobCount,omitempty"`
	// Group undocumented
	Group *Identity `json:"group,omitempty"`
	// GroupDisplayName undocumented
	GroupDisplayName *string `json:"groupDisplayName,omitempty"`
	// GroupMail undocumented
	GroupMail *string `json:"groupMail,omitempty"`
	// IncompleteJobCount undocumented
	IncompleteJobCount *int `json:"incompleteJobCount,omitempty"`
}

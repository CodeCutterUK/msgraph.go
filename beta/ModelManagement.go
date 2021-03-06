// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ManagementCertificateWithThumbprint undocumented
type ManagementCertificateWithThumbprint struct {
	// Object is the base model of ManagementCertificateWithThumbprint
	Object
	// Certificate The Base 64 encoded management certificate
	Certificate *string `json:"certificate,omitempty"`
	// Thumbprint The thumbprint of the management certificate
	Thumbprint *string `json:"thumbprint,omitempty"`
}

// ManagementCondition Management conditions are events that can be triggered dynamically such as geo-fences, time-fences, and network-fences.
type ManagementCondition struct {
	// Entity is the base model of ManagementCondition
	Entity
	// ApplicablePlatforms The applicable platforms for this management condition.
	ApplicablePlatforms []DevicePlatformType `json:"applicablePlatforms,omitempty"`
	// CreatedDateTime The time the management condition was created. Generated service side.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description The admin defined description of the management condition.
	Description *string `json:"description,omitempty"`
	// DisplayName The admin defined name of the management condition.
	DisplayName *string `json:"displayName,omitempty"`
	// ETag ETag of the management condition. Updated service side.
	ETag *string `json:"eTag,omitempty"`
	// ModifiedDateTime The time the management condition was last modified. Updated service side.
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// UniqueName Unique name for the management condition. Used in management condition expressions.
	UniqueName *string `json:"uniqueName,omitempty"`
	// ManagementConditionStatements undocumented
	ManagementConditionStatements []ManagementConditionStatement `json:"managementConditionStatements,omitempty"`
}

// ManagementConditionExpression A management condition expression is an expression that produces a boolean value when evaluated, i.e. one of true or false, indicating that a management condition statement is activated/deactivated. A management condition expression may be composed of a combination of the expression variables and boolean-valued expression operators.
type ManagementConditionExpression struct {
	// Object is the base model of ManagementConditionExpression
	Object
}

// ManagementConditionExpressionModel A management condition expression model is an model representation of a management condition expression.
type ManagementConditionExpressionModel struct {
	// ManagementConditionExpression is the base model of ManagementConditionExpressionModel
	ManagementConditionExpression
}

// ManagementConditionExpressionString A management condition expression string is a string representation of a management condition expression.
type ManagementConditionExpressionString struct {
	// ManagementConditionExpression is the base model of ManagementConditionExpressionString
	ManagementConditionExpression
	// Value The management condition statement expression string value.
	Value *string `json:"value,omitempty"`
}

// ManagementConditionStatement A management condition statement is a group of management conditions that enable/disable device/application configurations when all contained management conditions are met.
type ManagementConditionStatement struct {
	// Entity is the base model of ManagementConditionStatement
	Entity
	// ApplicablePlatforms The applicable platforms for this management condition statement.
	ApplicablePlatforms []DevicePlatformType `json:"applicablePlatforms,omitempty"`
	// CreatedDateTime The time the management condition statement was created. Generated service side.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description The admin defined description of the management condition statement.
	Description *string `json:"description,omitempty"`
	// DisplayName The admin defined name of the management condition statement.
	DisplayName *string `json:"displayName,omitempty"`
	// ETag ETag of the management condition statement. Updated service side.
	ETag *string `json:"eTag,omitempty"`
	// Expression The management condition statement expression used to evaluate if a management condition statement was activated/deactivated.
	Expression *ManagementConditionExpression `json:"expression,omitempty"`
	// ModifiedDateTime The time the management condition statement was last modified. Updated service side.
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// ManagementConditions undocumented
	ManagementConditions []ManagementCondition `json:"managementConditions,omitempty"`
}

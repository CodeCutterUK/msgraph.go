// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BinaryManagementConditionExpression A management condition expression that is evaluated using a binary operation.
type BinaryManagementConditionExpression struct {
	// ManagementConditionExpressionModel is the base model of BinaryManagementConditionExpression
	ManagementConditionExpressionModel
	// FirstOperand The first operand of the binary operation.
	FirstOperand *ManagementConditionExpressionModel `json:"firstOperand,omitempty"`
	// Operator The operator used in the evaluation of the binary operation.
	Operator *BinaryManagementConditionExpressionOperatorType `json:"operator,omitempty"`
	// SecondOperand The second operand of the binary operation.
	SecondOperand *ManagementConditionExpressionModel `json:"secondOperand,omitempty"`
}

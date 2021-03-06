// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Operator undocumented
type Operator string

const (
	// OperatorVNone undocumented
	OperatorVNone Operator = "none"
	// OperatorVAnd undocumented
	OperatorVAnd Operator = "and"
	// OperatorVOr undocumented
	OperatorVOr Operator = "or"
	// OperatorVIsEquals undocumented
	OperatorVIsEquals Operator = "isEquals"
	// OperatorVNotEquals undocumented
	OperatorVNotEquals Operator = "notEquals"
	// OperatorVGreaterThan undocumented
	OperatorVGreaterThan Operator = "greaterThan"
	// OperatorVLessThan undocumented
	OperatorVLessThan Operator = "lessThan"
	// OperatorVBetween undocumented
	OperatorVBetween Operator = "between"
	// OperatorVNotBetween undocumented
	OperatorVNotBetween Operator = "notBetween"
	// OperatorVGreaterEquals undocumented
	OperatorVGreaterEquals Operator = "greaterEquals"
	// OperatorVLessEquals undocumented
	OperatorVLessEquals Operator = "lessEquals"
	// OperatorVDayTimeBetween undocumented
	OperatorVDayTimeBetween Operator = "dayTimeBetween"
	// OperatorVBeginsWith undocumented
	OperatorVBeginsWith Operator = "beginsWith"
	// OperatorVNotBeginsWith undocumented
	OperatorVNotBeginsWith Operator = "notBeginsWith"
	// OperatorVEndsWith undocumented
	OperatorVEndsWith Operator = "endsWith"
	// OperatorVNotEndsWith undocumented
	OperatorVNotEndsWith Operator = "notEndsWith"
	// OperatorVContains undocumented
	OperatorVContains Operator = "contains"
	// OperatorVNotContains undocumented
	OperatorVNotContains Operator = "notContains"
	// OperatorVAllOf undocumented
	OperatorVAllOf Operator = "allOf"
	// OperatorVOneOf undocumented
	OperatorVOneOf Operator = "oneOf"
	// OperatorVNoneOf undocumented
	OperatorVNoneOf Operator = "noneOf"
	// OperatorVSetEquals undocumented
	OperatorVSetEquals Operator = "setEquals"
	// OperatorVOrderedSetEquals undocumented
	OperatorVOrderedSetEquals Operator = "orderedSetEquals"
	// OperatorVSubsetOf undocumented
	OperatorVSubsetOf Operator = "subsetOf"
	// OperatorVExcludesAll undocumented
	OperatorVExcludesAll Operator = "excludesAll"
)

var (
	// OperatorPNone is a pointer to OperatorVNone
	OperatorPNone = &_OperatorPNone
	// OperatorPAnd is a pointer to OperatorVAnd
	OperatorPAnd = &_OperatorPAnd
	// OperatorPOr is a pointer to OperatorVOr
	OperatorPOr = &_OperatorPOr
	// OperatorPIsEquals is a pointer to OperatorVIsEquals
	OperatorPIsEquals = &_OperatorPIsEquals
	// OperatorPNotEquals is a pointer to OperatorVNotEquals
	OperatorPNotEquals = &_OperatorPNotEquals
	// OperatorPGreaterThan is a pointer to OperatorVGreaterThan
	OperatorPGreaterThan = &_OperatorPGreaterThan
	// OperatorPLessThan is a pointer to OperatorVLessThan
	OperatorPLessThan = &_OperatorPLessThan
	// OperatorPBetween is a pointer to OperatorVBetween
	OperatorPBetween = &_OperatorPBetween
	// OperatorPNotBetween is a pointer to OperatorVNotBetween
	OperatorPNotBetween = &_OperatorPNotBetween
	// OperatorPGreaterEquals is a pointer to OperatorVGreaterEquals
	OperatorPGreaterEquals = &_OperatorPGreaterEquals
	// OperatorPLessEquals is a pointer to OperatorVLessEquals
	OperatorPLessEquals = &_OperatorPLessEquals
	// OperatorPDayTimeBetween is a pointer to OperatorVDayTimeBetween
	OperatorPDayTimeBetween = &_OperatorPDayTimeBetween
	// OperatorPBeginsWith is a pointer to OperatorVBeginsWith
	OperatorPBeginsWith = &_OperatorPBeginsWith
	// OperatorPNotBeginsWith is a pointer to OperatorVNotBeginsWith
	OperatorPNotBeginsWith = &_OperatorPNotBeginsWith
	// OperatorPEndsWith is a pointer to OperatorVEndsWith
	OperatorPEndsWith = &_OperatorPEndsWith
	// OperatorPNotEndsWith is a pointer to OperatorVNotEndsWith
	OperatorPNotEndsWith = &_OperatorPNotEndsWith
	// OperatorPContains is a pointer to OperatorVContains
	OperatorPContains = &_OperatorPContains
	// OperatorPNotContains is a pointer to OperatorVNotContains
	OperatorPNotContains = &_OperatorPNotContains
	// OperatorPAllOf is a pointer to OperatorVAllOf
	OperatorPAllOf = &_OperatorPAllOf
	// OperatorPOneOf is a pointer to OperatorVOneOf
	OperatorPOneOf = &_OperatorPOneOf
	// OperatorPNoneOf is a pointer to OperatorVNoneOf
	OperatorPNoneOf = &_OperatorPNoneOf
	// OperatorPSetEquals is a pointer to OperatorVSetEquals
	OperatorPSetEquals = &_OperatorPSetEquals
	// OperatorPOrderedSetEquals is a pointer to OperatorVOrderedSetEquals
	OperatorPOrderedSetEquals = &_OperatorPOrderedSetEquals
	// OperatorPSubsetOf is a pointer to OperatorVSubsetOf
	OperatorPSubsetOf = &_OperatorPSubsetOf
	// OperatorPExcludesAll is a pointer to OperatorVExcludesAll
	OperatorPExcludesAll = &_OperatorPExcludesAll
)

var (
	_OperatorPNone             = OperatorVNone
	_OperatorPAnd              = OperatorVAnd
	_OperatorPOr               = OperatorVOr
	_OperatorPIsEquals         = OperatorVIsEquals
	_OperatorPNotEquals        = OperatorVNotEquals
	_OperatorPGreaterThan      = OperatorVGreaterThan
	_OperatorPLessThan         = OperatorVLessThan
	_OperatorPBetween          = OperatorVBetween
	_OperatorPNotBetween       = OperatorVNotBetween
	_OperatorPGreaterEquals    = OperatorVGreaterEquals
	_OperatorPLessEquals       = OperatorVLessEquals
	_OperatorPDayTimeBetween   = OperatorVDayTimeBetween
	_OperatorPBeginsWith       = OperatorVBeginsWith
	_OperatorPNotBeginsWith    = OperatorVNotBeginsWith
	_OperatorPEndsWith         = OperatorVEndsWith
	_OperatorPNotEndsWith      = OperatorVNotEndsWith
	_OperatorPContains         = OperatorVContains
	_OperatorPNotContains      = OperatorVNotContains
	_OperatorPAllOf            = OperatorVAllOf
	_OperatorPOneOf            = OperatorVOneOf
	_OperatorPNoneOf           = OperatorVNoneOf
	_OperatorPSetEquals        = OperatorVSetEquals
	_OperatorPOrderedSetEquals = OperatorVOrderedSetEquals
	_OperatorPSubsetOf         = OperatorVSubsetOf
	_OperatorPExcludesAll      = OperatorVExcludesAll
)

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AttributeFlowBehavior undocumented
type AttributeFlowBehavior string

const (
	// AttributeFlowBehaviorVFlowWhenChanged undocumented
	AttributeFlowBehaviorVFlowWhenChanged AttributeFlowBehavior = "FlowWhenChanged"
	// AttributeFlowBehaviorVFlowAlways undocumented
	AttributeFlowBehaviorVFlowAlways AttributeFlowBehavior = "FlowAlways"
)

var (
	// AttributeFlowBehaviorPFlowWhenChanged is a pointer to AttributeFlowBehaviorVFlowWhenChanged
	AttributeFlowBehaviorPFlowWhenChanged = &_AttributeFlowBehaviorPFlowWhenChanged
	// AttributeFlowBehaviorPFlowAlways is a pointer to AttributeFlowBehaviorVFlowAlways
	AttributeFlowBehaviorPFlowAlways = &_AttributeFlowBehaviorPFlowAlways
)

var (
	_AttributeFlowBehaviorPFlowWhenChanged = AttributeFlowBehaviorVFlowWhenChanged
	_AttributeFlowBehaviorPFlowAlways      = AttributeFlowBehaviorVFlowAlways
)

// AttributeFlowType undocumented
type AttributeFlowType string

const (
	// AttributeFlowTypeVAlways undocumented
	AttributeFlowTypeVAlways AttributeFlowType = "Always"
	// AttributeFlowTypeVObjectAddOnly undocumented
	AttributeFlowTypeVObjectAddOnly AttributeFlowType = "ObjectAddOnly"
	// AttributeFlowTypeVMultiValueAddOnly undocumented
	AttributeFlowTypeVMultiValueAddOnly AttributeFlowType = "MultiValueAddOnly"
	// AttributeFlowTypeVValueAddOnly undocumented
	AttributeFlowTypeVValueAddOnly AttributeFlowType = "ValueAddOnly"
	// AttributeFlowTypeVAttributeAddOnly undocumented
	AttributeFlowTypeVAttributeAddOnly AttributeFlowType = "AttributeAddOnly"
)

var (
	// AttributeFlowTypePAlways is a pointer to AttributeFlowTypeVAlways
	AttributeFlowTypePAlways = &_AttributeFlowTypePAlways
	// AttributeFlowTypePObjectAddOnly is a pointer to AttributeFlowTypeVObjectAddOnly
	AttributeFlowTypePObjectAddOnly = &_AttributeFlowTypePObjectAddOnly
	// AttributeFlowTypePMultiValueAddOnly is a pointer to AttributeFlowTypeVMultiValueAddOnly
	AttributeFlowTypePMultiValueAddOnly = &_AttributeFlowTypePMultiValueAddOnly
	// AttributeFlowTypePValueAddOnly is a pointer to AttributeFlowTypeVValueAddOnly
	AttributeFlowTypePValueAddOnly = &_AttributeFlowTypePValueAddOnly
	// AttributeFlowTypePAttributeAddOnly is a pointer to AttributeFlowTypeVAttributeAddOnly
	AttributeFlowTypePAttributeAddOnly = &_AttributeFlowTypePAttributeAddOnly
)

var (
	_AttributeFlowTypePAlways            = AttributeFlowTypeVAlways
	_AttributeFlowTypePObjectAddOnly     = AttributeFlowTypeVObjectAddOnly
	_AttributeFlowTypePMultiValueAddOnly = AttributeFlowTypeVMultiValueAddOnly
	_AttributeFlowTypePValueAddOnly      = AttributeFlowTypeVValueAddOnly
	_AttributeFlowTypePAttributeAddOnly  = AttributeFlowTypeVAttributeAddOnly
)

// AttributeMappingSourceType undocumented
type AttributeMappingSourceType string

const (
	// AttributeMappingSourceTypeVAttribute undocumented
	AttributeMappingSourceTypeVAttribute AttributeMappingSourceType = "Attribute"
	// AttributeMappingSourceTypeVConstant undocumented
	AttributeMappingSourceTypeVConstant AttributeMappingSourceType = "Constant"
	// AttributeMappingSourceTypeVFunction undocumented
	AttributeMappingSourceTypeVFunction AttributeMappingSourceType = "Function"
)

var (
	// AttributeMappingSourceTypePAttribute is a pointer to AttributeMappingSourceTypeVAttribute
	AttributeMappingSourceTypePAttribute = &_AttributeMappingSourceTypePAttribute
	// AttributeMappingSourceTypePConstant is a pointer to AttributeMappingSourceTypeVConstant
	AttributeMappingSourceTypePConstant = &_AttributeMappingSourceTypePConstant
	// AttributeMappingSourceTypePFunction is a pointer to AttributeMappingSourceTypeVFunction
	AttributeMappingSourceTypePFunction = &_AttributeMappingSourceTypePFunction
)

var (
	_AttributeMappingSourceTypePAttribute = AttributeMappingSourceTypeVAttribute
	_AttributeMappingSourceTypePConstant  = AttributeMappingSourceTypeVConstant
	_AttributeMappingSourceTypePFunction  = AttributeMappingSourceTypeVFunction
)

// AttributeType undocumented
type AttributeType string

const (
	// AttributeTypeVString undocumented
	AttributeTypeVString AttributeType = "String"
	// AttributeTypeVInteger undocumented
	AttributeTypeVInteger AttributeType = "Integer"
	// AttributeTypeVReference undocumented
	AttributeTypeVReference AttributeType = "Reference"
	// AttributeTypeVBinary undocumented
	AttributeTypeVBinary AttributeType = "Binary"
	// AttributeTypeVBoolean undocumented
	AttributeTypeVBoolean AttributeType = "Boolean"
	// AttributeTypeVDateTime undocumented
	AttributeTypeVDateTime AttributeType = "DateTime"
)

var (
	// AttributeTypePString is a pointer to AttributeTypeVString
	AttributeTypePString = &_AttributeTypePString
	// AttributeTypePInteger is a pointer to AttributeTypeVInteger
	AttributeTypePInteger = &_AttributeTypePInteger
	// AttributeTypePReference is a pointer to AttributeTypeVReference
	AttributeTypePReference = &_AttributeTypePReference
	// AttributeTypePBinary is a pointer to AttributeTypeVBinary
	AttributeTypePBinary = &_AttributeTypePBinary
	// AttributeTypePBoolean is a pointer to AttributeTypeVBoolean
	AttributeTypePBoolean = &_AttributeTypePBoolean
	// AttributeTypePDateTime is a pointer to AttributeTypeVDateTime
	AttributeTypePDateTime = &_AttributeTypePDateTime
)

var (
	_AttributeTypePString    = AttributeTypeVString
	_AttributeTypePInteger   = AttributeTypeVInteger
	_AttributeTypePReference = AttributeTypeVReference
	_AttributeTypePBinary    = AttributeTypeVBinary
	_AttributeTypePBoolean   = AttributeTypeVBoolean
	_AttributeTypePDateTime  = AttributeTypeVDateTime
)

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ClassificationMethod undocumented
type ClassificationMethod string

const (
	// ClassificationMethodVPatternMatch undocumented
	ClassificationMethodVPatternMatch ClassificationMethod = "patternMatch"
	// ClassificationMethodVExactDataMatch undocumented
	ClassificationMethodVExactDataMatch ClassificationMethod = "exactDataMatch"
	// ClassificationMethodVFingerprint undocumented
	ClassificationMethodVFingerprint ClassificationMethod = "fingerprint"
	// ClassificationMethodVMachineLearning undocumented
	ClassificationMethodVMachineLearning ClassificationMethod = "machineLearning"
)

var (
	// ClassificationMethodPPatternMatch is a pointer to ClassificationMethodVPatternMatch
	ClassificationMethodPPatternMatch = &_ClassificationMethodPPatternMatch
	// ClassificationMethodPExactDataMatch is a pointer to ClassificationMethodVExactDataMatch
	ClassificationMethodPExactDataMatch = &_ClassificationMethodPExactDataMatch
	// ClassificationMethodPFingerprint is a pointer to ClassificationMethodVFingerprint
	ClassificationMethodPFingerprint = &_ClassificationMethodPFingerprint
	// ClassificationMethodPMachineLearning is a pointer to ClassificationMethodVMachineLearning
	ClassificationMethodPMachineLearning = &_ClassificationMethodPMachineLearning
)

var (
	_ClassificationMethodPPatternMatch    = ClassificationMethodVPatternMatch
	_ClassificationMethodPExactDataMatch  = ClassificationMethodVExactDataMatch
	_ClassificationMethodPFingerprint     = ClassificationMethodVFingerprint
	_ClassificationMethodPMachineLearning = ClassificationMethodVMachineLearning
)

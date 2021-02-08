// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CallDirection undocumented
type CallDirection string

const (
	// CallDirectionVIncoming undocumented
	CallDirectionVIncoming CallDirection = "incoming"
	// CallDirectionVOutgoing undocumented
	CallDirectionVOutgoing CallDirection = "outgoing"
)

var (
	// CallDirectionPIncoming is a pointer to CallDirectionVIncoming
	CallDirectionPIncoming = &_CallDirectionPIncoming
	// CallDirectionPOutgoing is a pointer to CallDirectionVOutgoing
	CallDirectionPOutgoing = &_CallDirectionPOutgoing
)

var (
	_CallDirectionPIncoming = CallDirectionVIncoming
	_CallDirectionPOutgoing = CallDirectionVOutgoing
)

// CallDisposition undocumented
type CallDisposition string

const (
	// CallDispositionVDefault undocumented
	CallDispositionVDefault CallDisposition = "default"
	// CallDispositionVSimultaneousRing undocumented
	CallDispositionVSimultaneousRing CallDisposition = "simultaneousRing"
	// CallDispositionVForward undocumented
	CallDispositionVForward CallDisposition = "forward"
)

var (
	// CallDispositionPDefault is a pointer to CallDispositionVDefault
	CallDispositionPDefault = &_CallDispositionPDefault
	// CallDispositionPSimultaneousRing is a pointer to CallDispositionVSimultaneousRing
	CallDispositionPSimultaneousRing = &_CallDispositionPSimultaneousRing
	// CallDispositionPForward is a pointer to CallDispositionVForward
	CallDispositionPForward = &_CallDispositionPForward
)

var (
	_CallDispositionPDefault          = CallDispositionVDefault
	_CallDispositionPSimultaneousRing = CallDispositionVSimultaneousRing
	_CallDispositionPForward          = CallDispositionVForward
)

// CallState undocumented
type CallState string

const (
	// CallStateVIncoming undocumented
	CallStateVIncoming CallState = "incoming"
	// CallStateVEstablishing undocumented
	CallStateVEstablishing CallState = "establishing"
	// CallStateVRinging undocumented
	CallStateVRinging CallState = "ringing"
	// CallStateVEstablished undocumented
	CallStateVEstablished CallState = "established"
	// CallStateVHold undocumented
	CallStateVHold CallState = "hold"
	// CallStateVTransferring undocumented
	CallStateVTransferring CallState = "transferring"
	// CallStateVTransferAccepted undocumented
	CallStateVTransferAccepted CallState = "transferAccepted"
	// CallStateVRedirecting undocumented
	CallStateVRedirecting CallState = "redirecting"
	// CallStateVTerminating undocumented
	CallStateVTerminating CallState = "terminating"
	// CallStateVTerminated undocumented
	CallStateVTerminated CallState = "terminated"
	// CallStateVUnknownFutureValue undocumented
	CallStateVUnknownFutureValue CallState = "unknownFutureValue"
)

var (
	// CallStatePIncoming is a pointer to CallStateVIncoming
	CallStatePIncoming = &_CallStatePIncoming
	// CallStatePEstablishing is a pointer to CallStateVEstablishing
	CallStatePEstablishing = &_CallStatePEstablishing
	// CallStatePRinging is a pointer to CallStateVRinging
	CallStatePRinging = &_CallStatePRinging
	// CallStatePEstablished is a pointer to CallStateVEstablished
	CallStatePEstablished = &_CallStatePEstablished
	// CallStatePHold is a pointer to CallStateVHold
	CallStatePHold = &_CallStatePHold
	// CallStatePTransferring is a pointer to CallStateVTransferring
	CallStatePTransferring = &_CallStatePTransferring
	// CallStatePTransferAccepted is a pointer to CallStateVTransferAccepted
	CallStatePTransferAccepted = &_CallStatePTransferAccepted
	// CallStatePRedirecting is a pointer to CallStateVRedirecting
	CallStatePRedirecting = &_CallStatePRedirecting
	// CallStatePTerminating is a pointer to CallStateVTerminating
	CallStatePTerminating = &_CallStatePTerminating
	// CallStatePTerminated is a pointer to CallStateVTerminated
	CallStatePTerminated = &_CallStatePTerminated
	// CallStatePUnknownFutureValue is a pointer to CallStateVUnknownFutureValue
	CallStatePUnknownFutureValue = &_CallStatePUnknownFutureValue
)

var (
	_CallStatePIncoming           = CallStateVIncoming
	_CallStatePEstablishing       = CallStateVEstablishing
	_CallStatePRinging            = CallStateVRinging
	_CallStatePEstablished        = CallStateVEstablished
	_CallStatePHold               = CallStateVHold
	_CallStatePTransferring       = CallStateVTransferring
	_CallStatePTransferAccepted   = CallStateVTransferAccepted
	_CallStatePRedirecting        = CallStateVRedirecting
	_CallStatePTerminating        = CallStateVTerminating
	_CallStatePTerminated         = CallStateVTerminated
	_CallStatePUnknownFutureValue = CallStateVUnknownFutureValue
)

// CallTranscriptionState undocumented
type CallTranscriptionState string

const (
	// CallTranscriptionStateVNotStarted undocumented
	CallTranscriptionStateVNotStarted CallTranscriptionState = "notStarted"
	// CallTranscriptionStateVActive undocumented
	CallTranscriptionStateVActive CallTranscriptionState = "active"
	// CallTranscriptionStateVInactive undocumented
	CallTranscriptionStateVInactive CallTranscriptionState = "inactive"
	// CallTranscriptionStateVUnknownFutureValue undocumented
	CallTranscriptionStateVUnknownFutureValue CallTranscriptionState = "unknownFutureValue"
)

var (
	// CallTranscriptionStatePNotStarted is a pointer to CallTranscriptionStateVNotStarted
	CallTranscriptionStatePNotStarted = &_CallTranscriptionStatePNotStarted
	// CallTranscriptionStatePActive is a pointer to CallTranscriptionStateVActive
	CallTranscriptionStatePActive = &_CallTranscriptionStatePActive
	// CallTranscriptionStatePInactive is a pointer to CallTranscriptionStateVInactive
	CallTranscriptionStatePInactive = &_CallTranscriptionStatePInactive
	// CallTranscriptionStatePUnknownFutureValue is a pointer to CallTranscriptionStateVUnknownFutureValue
	CallTranscriptionStatePUnknownFutureValue = &_CallTranscriptionStatePUnknownFutureValue
)

var (
	_CallTranscriptionStatePNotStarted         = CallTranscriptionStateVNotStarted
	_CallTranscriptionStatePActive             = CallTranscriptionStateVActive
	_CallTranscriptionStatePInactive           = CallTranscriptionStateVInactive
	_CallTranscriptionStatePUnknownFutureValue = CallTranscriptionStateVUnknownFutureValue
)

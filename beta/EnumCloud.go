// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CloudAppSecuritySessionControlType undocumented
type CloudAppSecuritySessionControlType string

const (
	// CloudAppSecuritySessionControlTypeVMcasConfigured undocumented
	CloudAppSecuritySessionControlTypeVMcasConfigured CloudAppSecuritySessionControlType = "mcasConfigured"
	// CloudAppSecuritySessionControlTypeVMonitorOnly undocumented
	CloudAppSecuritySessionControlTypeVMonitorOnly CloudAppSecuritySessionControlType = "monitorOnly"
	// CloudAppSecuritySessionControlTypeVBlockDownloads undocumented
	CloudAppSecuritySessionControlTypeVBlockDownloads CloudAppSecuritySessionControlType = "blockDownloads"
	// CloudAppSecuritySessionControlTypeVUnknownFutureValue undocumented
	CloudAppSecuritySessionControlTypeVUnknownFutureValue CloudAppSecuritySessionControlType = "unknownFutureValue"
)

var (
	// CloudAppSecuritySessionControlTypePMcasConfigured is a pointer to CloudAppSecuritySessionControlTypeVMcasConfigured
	CloudAppSecuritySessionControlTypePMcasConfigured = &_CloudAppSecuritySessionControlTypePMcasConfigured
	// CloudAppSecuritySessionControlTypePMonitorOnly is a pointer to CloudAppSecuritySessionControlTypeVMonitorOnly
	CloudAppSecuritySessionControlTypePMonitorOnly = &_CloudAppSecuritySessionControlTypePMonitorOnly
	// CloudAppSecuritySessionControlTypePBlockDownloads is a pointer to CloudAppSecuritySessionControlTypeVBlockDownloads
	CloudAppSecuritySessionControlTypePBlockDownloads = &_CloudAppSecuritySessionControlTypePBlockDownloads
	// CloudAppSecuritySessionControlTypePUnknownFutureValue is a pointer to CloudAppSecuritySessionControlTypeVUnknownFutureValue
	CloudAppSecuritySessionControlTypePUnknownFutureValue = &_CloudAppSecuritySessionControlTypePUnknownFutureValue
)

var (
	_CloudAppSecuritySessionControlTypePMcasConfigured     = CloudAppSecuritySessionControlTypeVMcasConfigured
	_CloudAppSecuritySessionControlTypePMonitorOnly        = CloudAppSecuritySessionControlTypeVMonitorOnly
	_CloudAppSecuritySessionControlTypePBlockDownloads     = CloudAppSecuritySessionControlTypeVBlockDownloads
	_CloudAppSecuritySessionControlTypePUnknownFutureValue = CloudAppSecuritySessionControlTypeVUnknownFutureValue
)

// CloudPcDeviceImageStatus undocumented
type CloudPcDeviceImageStatus string

const (
	// CloudPcDeviceImageStatusVPending undocumented
	CloudPcDeviceImageStatusVPending CloudPcDeviceImageStatus = "pending"
	// CloudPcDeviceImageStatusVReady undocumented
	CloudPcDeviceImageStatusVReady CloudPcDeviceImageStatus = "ready"
	// CloudPcDeviceImageStatusVFailed undocumented
	CloudPcDeviceImageStatusVFailed CloudPcDeviceImageStatus = "failed"
)

var (
	// CloudPcDeviceImageStatusPPending is a pointer to CloudPcDeviceImageStatusVPending
	CloudPcDeviceImageStatusPPending = &_CloudPcDeviceImageStatusPPending
	// CloudPcDeviceImageStatusPReady is a pointer to CloudPcDeviceImageStatusVReady
	CloudPcDeviceImageStatusPReady = &_CloudPcDeviceImageStatusPReady
	// CloudPcDeviceImageStatusPFailed is a pointer to CloudPcDeviceImageStatusVFailed
	CloudPcDeviceImageStatusPFailed = &_CloudPcDeviceImageStatusPFailed
)

var (
	_CloudPcDeviceImageStatusPPending = CloudPcDeviceImageStatusVPending
	_CloudPcDeviceImageStatusPReady   = CloudPcDeviceImageStatusVReady
	_CloudPcDeviceImageStatusPFailed  = CloudPcDeviceImageStatusVFailed
)

// CloudPcDeviceImageStatusDetails undocumented
type CloudPcDeviceImageStatusDetails string

const (
	// CloudPcDeviceImageStatusDetailsVInternalServerError undocumented
	CloudPcDeviceImageStatusDetailsVInternalServerError CloudPcDeviceImageStatusDetails = "internalServerError"
	// CloudPcDeviceImageStatusDetailsVSourceImageNotFound undocumented
	CloudPcDeviceImageStatusDetailsVSourceImageNotFound CloudPcDeviceImageStatusDetails = "sourceImageNotFound"
)

var (
	// CloudPcDeviceImageStatusDetailsPInternalServerError is a pointer to CloudPcDeviceImageStatusDetailsVInternalServerError
	CloudPcDeviceImageStatusDetailsPInternalServerError = &_CloudPcDeviceImageStatusDetailsPInternalServerError
	// CloudPcDeviceImageStatusDetailsPSourceImageNotFound is a pointer to CloudPcDeviceImageStatusDetailsVSourceImageNotFound
	CloudPcDeviceImageStatusDetailsPSourceImageNotFound = &_CloudPcDeviceImageStatusDetailsPSourceImageNotFound
)

var (
	_CloudPcDeviceImageStatusDetailsPInternalServerError = CloudPcDeviceImageStatusDetailsVInternalServerError
	_CloudPcDeviceImageStatusDetailsPSourceImageNotFound = CloudPcDeviceImageStatusDetailsVSourceImageNotFound
)

// CloudPcOnPremisesConnectionHealthCheckErrorType undocumented
type CloudPcOnPremisesConnectionHealthCheckErrorType string

const (
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVDNSCheckFqdnNotFound undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVDNSCheckFqdnNotFound CloudPcOnPremisesConnectionHealthCheckErrorType = "dnsCheckFqdnNotFound"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVDNSCheckUnknownError undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVDNSCheckUnknownError CloudPcOnPremisesConnectionHealthCheckErrorType = "dnsCheckUnknownError"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckFqdnNotFound undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckFqdnNotFound CloudPcOnPremisesConnectionHealthCheckErrorType = "adJoinCheckFqdnNotFound"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckIncorrectCredentials undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckIncorrectCredentials CloudPcOnPremisesConnectionHealthCheckErrorType = "adJoinCheckIncorrectCredentials"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckOrganizationalUnitNotFound undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckOrganizationalUnitNotFound CloudPcOnPremisesConnectionHealthCheckErrorType = "adJoinCheckOrganizationalUnitNotFound"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckOrganizationalUnitIncorrectFormat undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckOrganizationalUnitIncorrectFormat CloudPcOnPremisesConnectionHealthCheckErrorType = "adJoinCheckOrganizationalUnitIncorrectFormat"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckUnknownError undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckUnknownError CloudPcOnPremisesConnectionHealthCheckErrorType = "adJoinCheckUnknownError"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVEndpointConnectivityCheckURLNotWhitelisted undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVEndpointConnectivityCheckURLNotWhitelisted CloudPcOnPremisesConnectionHealthCheckErrorType = "endpointConnectivityCheckUrlNotWhitelisted"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVEndpointConnectivityCheckUnknownError undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVEndpointConnectivityCheckUnknownError CloudPcOnPremisesConnectionHealthCheckErrorType = "endpointConnectivityCheckUnknownError"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVAadConnectivityCheckUnknownError undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVAadConnectivityCheckUnknownError CloudPcOnPremisesConnectionHealthCheckErrorType = "aadConnectivityCheckUnknownError"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckNoSubnetIP undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckNoSubnetIP CloudPcOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckNoSubnetIP"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckSubscriptionDisabled undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckSubscriptionDisabled CloudPcOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckSubscriptionDisabled"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckUnknownError undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckUnknownError CloudPcOnPremisesConnectionHealthCheckErrorType = "resourceAvailabilityCheckUnknownError"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoSubscriptionReaderRole undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoSubscriptionReaderRole CloudPcOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoSubscriptionReaderRole"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoResourceGroupOwnerRole undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoResourceGroupOwnerRole CloudPcOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoResourceGroupOwnerRole"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoVNetContributorRole undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoVNetContributorRole CloudPcOnPremisesConnectionHealthCheckErrorType = "permissionCheckNoVNetContributorRole"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckUnknownError undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckUnknownError CloudPcOnPremisesConnectionHealthCheckErrorType = "permissionCheckUnknownError"
	// CloudPcOnPremisesConnectionHealthCheckErrorTypeVInternalServerUnknownError undocumented
	CloudPcOnPremisesConnectionHealthCheckErrorTypeVInternalServerUnknownError CloudPcOnPremisesConnectionHealthCheckErrorType = "internalServerUnknownError"
)

var (
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePDNSCheckFqdnNotFound is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVDNSCheckFqdnNotFound
	CloudPcOnPremisesConnectionHealthCheckErrorTypePDNSCheckFqdnNotFound = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePDNSCheckFqdnNotFound
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePDNSCheckUnknownError is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVDNSCheckUnknownError
	CloudPcOnPremisesConnectionHealthCheckErrorTypePDNSCheckUnknownError = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePDNSCheckUnknownError
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckFqdnNotFound is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckFqdnNotFound
	CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckFqdnNotFound = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckFqdnNotFound
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckIncorrectCredentials is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckIncorrectCredentials
	CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckIncorrectCredentials = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckIncorrectCredentials
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckOrganizationalUnitNotFound is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckOrganizationalUnitNotFound
	CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckOrganizationalUnitNotFound = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckOrganizationalUnitNotFound
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckOrganizationalUnitIncorrectFormat is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckOrganizationalUnitIncorrectFormat
	CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckOrganizationalUnitIncorrectFormat = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckOrganizationalUnitIncorrectFormat
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckUnknownError is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckUnknownError
	CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckUnknownError = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckUnknownError
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePEndpointConnectivityCheckURLNotWhitelisted is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVEndpointConnectivityCheckURLNotWhitelisted
	CloudPcOnPremisesConnectionHealthCheckErrorTypePEndpointConnectivityCheckURLNotWhitelisted = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePEndpointConnectivityCheckURLNotWhitelisted
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePEndpointConnectivityCheckUnknownError is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVEndpointConnectivityCheckUnknownError
	CloudPcOnPremisesConnectionHealthCheckErrorTypePEndpointConnectivityCheckUnknownError = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePEndpointConnectivityCheckUnknownError
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePAadConnectivityCheckUnknownError is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVAadConnectivityCheckUnknownError
	CloudPcOnPremisesConnectionHealthCheckErrorTypePAadConnectivityCheckUnknownError = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePAadConnectivityCheckUnknownError
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckNoSubnetIP is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckNoSubnetIP
	CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckNoSubnetIP = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckNoSubnetIP
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckSubscriptionDisabled is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckSubscriptionDisabled
	CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckSubscriptionDisabled = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckSubscriptionDisabled
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckUnknownError is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckUnknownError
	CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckUnknownError = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckUnknownError
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoSubscriptionReaderRole is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoSubscriptionReaderRole
	CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoSubscriptionReaderRole = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoSubscriptionReaderRole
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoResourceGroupOwnerRole is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoResourceGroupOwnerRole
	CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoResourceGroupOwnerRole = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoResourceGroupOwnerRole
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoVNetContributorRole is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoVNetContributorRole
	CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoVNetContributorRole = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoVNetContributorRole
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckUnknownError is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckUnknownError
	CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckUnknownError = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckUnknownError
	// CloudPcOnPremisesConnectionHealthCheckErrorTypePInternalServerUnknownError is a pointer to CloudPcOnPremisesConnectionHealthCheckErrorTypeVInternalServerUnknownError
	CloudPcOnPremisesConnectionHealthCheckErrorTypePInternalServerUnknownError = &_CloudPcOnPremisesConnectionHealthCheckErrorTypePInternalServerUnknownError
)

var (
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePDNSCheckFqdnNotFound                          = CloudPcOnPremisesConnectionHealthCheckErrorTypeVDNSCheckFqdnNotFound
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePDNSCheckUnknownError                          = CloudPcOnPremisesConnectionHealthCheckErrorTypeVDNSCheckUnknownError
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckFqdnNotFound                       = CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckFqdnNotFound
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckIncorrectCredentials               = CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckIncorrectCredentials
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckOrganizationalUnitNotFound         = CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckOrganizationalUnitNotFound
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckOrganizationalUnitIncorrectFormat  = CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckOrganizationalUnitIncorrectFormat
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePAdJoinCheckUnknownError                       = CloudPcOnPremisesConnectionHealthCheckErrorTypeVAdJoinCheckUnknownError
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePEndpointConnectivityCheckURLNotWhitelisted    = CloudPcOnPremisesConnectionHealthCheckErrorTypeVEndpointConnectivityCheckURLNotWhitelisted
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePEndpointConnectivityCheckUnknownError         = CloudPcOnPremisesConnectionHealthCheckErrorTypeVEndpointConnectivityCheckUnknownError
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePAadConnectivityCheckUnknownError              = CloudPcOnPremisesConnectionHealthCheckErrorTypeVAadConnectivityCheckUnknownError
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckNoSubnetIP           = CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckNoSubnetIP
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckSubscriptionDisabled = CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckSubscriptionDisabled
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePResourceAvailabilityCheckUnknownError         = CloudPcOnPremisesConnectionHealthCheckErrorTypeVResourceAvailabilityCheckUnknownError
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoSubscriptionReaderRole       = CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoSubscriptionReaderRole
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoResourceGroupOwnerRole       = CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoResourceGroupOwnerRole
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckNoVNetContributorRole          = CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckNoVNetContributorRole
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePPermissionCheckUnknownError                   = CloudPcOnPremisesConnectionHealthCheckErrorTypeVPermissionCheckUnknownError
	_CloudPcOnPremisesConnectionHealthCheckErrorTypePInternalServerUnknownError                    = CloudPcOnPremisesConnectionHealthCheckErrorTypeVInternalServerUnknownError
)

// CloudPcOnPremisesConnectionStatus undocumented
type CloudPcOnPremisesConnectionStatus string

const (
	// CloudPcOnPremisesConnectionStatusVPending undocumented
	CloudPcOnPremisesConnectionStatusVPending CloudPcOnPremisesConnectionStatus = "pending"
	// CloudPcOnPremisesConnectionStatusVRunning undocumented
	CloudPcOnPremisesConnectionStatusVRunning CloudPcOnPremisesConnectionStatus = "running"
	// CloudPcOnPremisesConnectionStatusVPassed undocumented
	CloudPcOnPremisesConnectionStatusVPassed CloudPcOnPremisesConnectionStatus = "passed"
	// CloudPcOnPremisesConnectionStatusVFailed undocumented
	CloudPcOnPremisesConnectionStatusVFailed CloudPcOnPremisesConnectionStatus = "failed"
	// CloudPcOnPremisesConnectionStatusVUnknownFutureValue undocumented
	CloudPcOnPremisesConnectionStatusVUnknownFutureValue CloudPcOnPremisesConnectionStatus = "unknownFutureValue"
)

var (
	// CloudPcOnPremisesConnectionStatusPPending is a pointer to CloudPcOnPremisesConnectionStatusVPending
	CloudPcOnPremisesConnectionStatusPPending = &_CloudPcOnPremisesConnectionStatusPPending
	// CloudPcOnPremisesConnectionStatusPRunning is a pointer to CloudPcOnPremisesConnectionStatusVRunning
	CloudPcOnPremisesConnectionStatusPRunning = &_CloudPcOnPremisesConnectionStatusPRunning
	// CloudPcOnPremisesConnectionStatusPPassed is a pointer to CloudPcOnPremisesConnectionStatusVPassed
	CloudPcOnPremisesConnectionStatusPPassed = &_CloudPcOnPremisesConnectionStatusPPassed
	// CloudPcOnPremisesConnectionStatusPFailed is a pointer to CloudPcOnPremisesConnectionStatusVFailed
	CloudPcOnPremisesConnectionStatusPFailed = &_CloudPcOnPremisesConnectionStatusPFailed
	// CloudPcOnPremisesConnectionStatusPUnknownFutureValue is a pointer to CloudPcOnPremisesConnectionStatusVUnknownFutureValue
	CloudPcOnPremisesConnectionStatusPUnknownFutureValue = &_CloudPcOnPremisesConnectionStatusPUnknownFutureValue
)

var (
	_CloudPcOnPremisesConnectionStatusPPending            = CloudPcOnPremisesConnectionStatusVPending
	_CloudPcOnPremisesConnectionStatusPRunning            = CloudPcOnPremisesConnectionStatusVRunning
	_CloudPcOnPremisesConnectionStatusPPassed             = CloudPcOnPremisesConnectionStatusVPassed
	_CloudPcOnPremisesConnectionStatusPFailed             = CloudPcOnPremisesConnectionStatusVFailed
	_CloudPcOnPremisesConnectionStatusPUnknownFutureValue = CloudPcOnPremisesConnectionStatusVUnknownFutureValue
)

// CloudPcProvisioningPolicyImageType undocumented
type CloudPcProvisioningPolicyImageType string

const (
	// CloudPcProvisioningPolicyImageTypeVGallery undocumented
	CloudPcProvisioningPolicyImageTypeVGallery CloudPcProvisioningPolicyImageType = "gallery"
	// CloudPcProvisioningPolicyImageTypeVCustom undocumented
	CloudPcProvisioningPolicyImageTypeVCustom CloudPcProvisioningPolicyImageType = "custom"
)

var (
	// CloudPcProvisioningPolicyImageTypePGallery is a pointer to CloudPcProvisioningPolicyImageTypeVGallery
	CloudPcProvisioningPolicyImageTypePGallery = &_CloudPcProvisioningPolicyImageTypePGallery
	// CloudPcProvisioningPolicyImageTypePCustom is a pointer to CloudPcProvisioningPolicyImageTypeVCustom
	CloudPcProvisioningPolicyImageTypePCustom = &_CloudPcProvisioningPolicyImageTypePCustom
)

var (
	_CloudPcProvisioningPolicyImageTypePGallery = CloudPcProvisioningPolicyImageTypeVGallery
	_CloudPcProvisioningPolicyImageTypePCustom  = CloudPcProvisioningPolicyImageTypeVCustom
)

// CloudPcStatus undocumented
type CloudPcStatus string

const (
	// CloudPcStatusVNotProvisioned undocumented
	CloudPcStatusVNotProvisioned CloudPcStatus = "notProvisioned"
	// CloudPcStatusVProvisioning undocumented
	CloudPcStatusVProvisioning CloudPcStatus = "provisioning"
	// CloudPcStatusVProvisioned undocumented
	CloudPcStatusVProvisioned CloudPcStatus = "provisioned"
	// CloudPcStatusVUpgrading undocumented
	CloudPcStatusVUpgrading CloudPcStatus = "upgrading"
	// CloudPcStatusVInGracePeriod undocumented
	CloudPcStatusVInGracePeriod CloudPcStatus = "inGracePeriod"
	// CloudPcStatusVDeprovisioning undocumented
	CloudPcStatusVDeprovisioning CloudPcStatus = "deprovisioning"
	// CloudPcStatusVFailed undocumented
	CloudPcStatusVFailed CloudPcStatus = "failed"
)

var (
	// CloudPcStatusPNotProvisioned is a pointer to CloudPcStatusVNotProvisioned
	CloudPcStatusPNotProvisioned = &_CloudPcStatusPNotProvisioned
	// CloudPcStatusPProvisioning is a pointer to CloudPcStatusVProvisioning
	CloudPcStatusPProvisioning = &_CloudPcStatusPProvisioning
	// CloudPcStatusPProvisioned is a pointer to CloudPcStatusVProvisioned
	CloudPcStatusPProvisioned = &_CloudPcStatusPProvisioned
	// CloudPcStatusPUpgrading is a pointer to CloudPcStatusVUpgrading
	CloudPcStatusPUpgrading = &_CloudPcStatusPUpgrading
	// CloudPcStatusPInGracePeriod is a pointer to CloudPcStatusVInGracePeriod
	CloudPcStatusPInGracePeriod = &_CloudPcStatusPInGracePeriod
	// CloudPcStatusPDeprovisioning is a pointer to CloudPcStatusVDeprovisioning
	CloudPcStatusPDeprovisioning = &_CloudPcStatusPDeprovisioning
	// CloudPcStatusPFailed is a pointer to CloudPcStatusVFailed
	CloudPcStatusPFailed = &_CloudPcStatusPFailed
)

var (
	_CloudPcStatusPNotProvisioned = CloudPcStatusVNotProvisioned
	_CloudPcStatusPProvisioning   = CloudPcStatusVProvisioning
	_CloudPcStatusPProvisioned    = CloudPcStatusVProvisioned
	_CloudPcStatusPUpgrading      = CloudPcStatusVUpgrading
	_CloudPcStatusPInGracePeriod  = CloudPcStatusVInGracePeriod
	_CloudPcStatusPDeprovisioning = CloudPcStatusVDeprovisioning
	_CloudPcStatusPFailed         = CloudPcStatusVFailed
)

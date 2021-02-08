// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DeliveryOptimizationBandwidth Bandwidth restriction types
type DeliveryOptimizationBandwidth struct {
	// Object is the base model of DeliveryOptimizationBandwidth
	Object
}

// DeliveryOptimizationBandwidthAbsolute Bandwidth limits in kilobytes per second.
type DeliveryOptimizationBandwidthAbsolute struct {
	// DeliveryOptimizationBandwidth is the base model of DeliveryOptimizationBandwidthAbsolute
	DeliveryOptimizationBandwidth
	// MaximumDownloadBandwidthInKilobytesPerSecond Specifies the maximum download bandwidth in KiloBytes/second that the device can use across all concurrent download activities using Delivery Optimization. Valid values 0 to 4294967295
	MaximumDownloadBandwidthInKilobytesPerSecond *int `json:"maximumDownloadBandwidthInKilobytesPerSecond,omitempty"`
	// MaximumUploadBandwidthInKilobytesPerSecond Specifies the maximum upload bandwidth in KiloBytes/second that a device will use across all concurrent upload activity using Delivery Optimization (0-4000000). Valid values 0 to 4000000
	MaximumUploadBandwidthInKilobytesPerSecond *int `json:"maximumUploadBandwidthInKilobytesPerSecond,omitempty"`
}

// DeliveryOptimizationBandwidthBusinessHoursLimit Bandwidth business hours and percentages type
type DeliveryOptimizationBandwidthBusinessHoursLimit struct {
	// Object is the base model of DeliveryOptimizationBandwidthBusinessHoursLimit
	Object
	// BandwidthBeginBusinessHours Specifies the beginning of business hours using a 24-hour clock (0-23). Valid values 0 to 23
	BandwidthBeginBusinessHours *int `json:"bandwidthBeginBusinessHours,omitempty"`
	// BandwidthEndBusinessHours Specifies the end of business hours using a 24-hour clock (0-23). Valid values 0 to 23
	BandwidthEndBusinessHours *int `json:"bandwidthEndBusinessHours,omitempty"`
	// BandwidthPercentageDuringBusinessHours Specifies the percentage of bandwidth to limit during business hours (0-100). Valid values 0 to 100
	BandwidthPercentageDuringBusinessHours *int `json:"bandwidthPercentageDuringBusinessHours,omitempty"`
	// BandwidthPercentageOutsideBusinessHours Specifies the percentage of bandwidth to limit outsidse business hours (0-100). Valid values 0 to 100
	BandwidthPercentageOutsideBusinessHours *int `json:"bandwidthPercentageOutsideBusinessHours,omitempty"`
}

// DeliveryOptimizationBandwidthHoursWithPercentage Bandwidth limit as a percentage with business hours.
type DeliveryOptimizationBandwidthHoursWithPercentage struct {
	// DeliveryOptimizationBandwidth is the base model of DeliveryOptimizationBandwidthHoursWithPercentage
	DeliveryOptimizationBandwidth
	// BandwidthBackgroundPercentageHours Background download percentage hours.
	BandwidthBackgroundPercentageHours *DeliveryOptimizationBandwidthBusinessHoursLimit `json:"bandwidthBackgroundPercentageHours,omitempty"`
	// BandwidthForegroundPercentageHours Foreground download percentage hours.
	BandwidthForegroundPercentageHours *DeliveryOptimizationBandwidthBusinessHoursLimit `json:"bandwidthForegroundPercentageHours,omitempty"`
}

// DeliveryOptimizationBandwidthPercentage Bandwidth limits specified as a percentage.
type DeliveryOptimizationBandwidthPercentage struct {
	// DeliveryOptimizationBandwidth is the base model of DeliveryOptimizationBandwidthPercentage
	DeliveryOptimizationBandwidth
	// MaximumBackgroundBandwidthPercentage Specifies the maximum background download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
	MaximumBackgroundBandwidthPercentage *int `json:"maximumBackgroundBandwidthPercentage,omitempty"`
	// MaximumForegroundBandwidthPercentage Specifies the maximum foreground download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
	MaximumForegroundBandwidthPercentage *int `json:"maximumForegroundBandwidthPercentage,omitempty"`
}

// DeliveryOptimizationGroupIDCustom Custom group id type
type DeliveryOptimizationGroupIDCustom struct {
	// DeliveryOptimizationGroupIDSource is the base model of DeliveryOptimizationGroupIDCustom
	DeliveryOptimizationGroupIDSource
	// GroupIDCustom Specifies an arbitrary group ID that the device belongs to
	GroupIDCustom *string `json:"groupIdCustom,omitempty"`
}

// DeliveryOptimizationGroupIDSource GroupId Support Types
type DeliveryOptimizationGroupIDSource struct {
	// Object is the base model of DeliveryOptimizationGroupIDSource
	Object
}

// DeliveryOptimizationGroupIDSourceOptions Group id options type
type DeliveryOptimizationGroupIDSourceOptions struct {
	// DeliveryOptimizationGroupIDSource is the base model of DeliveryOptimizationGroupIDSourceOptions
	DeliveryOptimizationGroupIDSource
	// GroupIDSourceOption Set this policy to restrict peer selection to a specific source.
	GroupIDSourceOption *DeliveryOptimizationGroupIDOptionsType `json:"groupIdSourceOption,omitempty"`
}

// DeliveryOptimizationMaxCacheSize Delivery Optimization max cache size types.
type DeliveryOptimizationMaxCacheSize struct {
	// Object is the base model of DeliveryOptimizationMaxCacheSize
	Object
}

// DeliveryOptimizationMaxCacheSizeAbsolute Delivery Optimization max cache size absolute type.
type DeliveryOptimizationMaxCacheSizeAbsolute struct {
	// DeliveryOptimizationMaxCacheSize is the base model of DeliveryOptimizationMaxCacheSizeAbsolute
	DeliveryOptimizationMaxCacheSize
	// MaximumCacheSizeInGigabytes Specifies the maximum size in GB of Delivery Optimization cache. Valid values 0 to 4294967295
	MaximumCacheSizeInGigabytes *int `json:"maximumCacheSizeInGigabytes,omitempty"`
}

// DeliveryOptimizationMaxCacheSizePercentage Delivery Optimization Max cache size percentage types.
type DeliveryOptimizationMaxCacheSizePercentage struct {
	// DeliveryOptimizationMaxCacheSize is the base model of DeliveryOptimizationMaxCacheSizePercentage
	DeliveryOptimizationMaxCacheSize
	// MaximumCacheSizePercentage Specifies the maximum cache size that Delivery Optimization can utilize, as a percentage of disk size (1-100). Valid values 1 to 100
	MaximumCacheSizePercentage *int `json:"maximumCacheSizePercentage,omitempty"`
}

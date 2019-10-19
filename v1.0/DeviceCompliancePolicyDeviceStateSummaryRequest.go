// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceCompliancePolicyDeviceStateSummaryRequestBuilder is request builder for DeviceCompliancePolicyDeviceStateSummary
type DeviceCompliancePolicyDeviceStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceCompliancePolicyDeviceStateSummaryRequest
func (b *DeviceCompliancePolicyDeviceStateSummaryRequestBuilder) Request() *DeviceCompliancePolicyDeviceStateSummaryRequest {
	return &DeviceCompliancePolicyDeviceStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceCompliancePolicyDeviceStateSummaryRequest is request for DeviceCompliancePolicyDeviceStateSummary
type DeviceCompliancePolicyDeviceStateSummaryRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceCompliancePolicyDeviceStateSummary
func (r *DeviceCompliancePolicyDeviceStateSummaryRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceCompliancePolicyDeviceStateSummary, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceCompliancePolicyDeviceStateSummary
func (r *DeviceCompliancePolicyDeviceStateSummaryRequest) Get() (*DeviceCompliancePolicyDeviceStateSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceCompliancePolicyDeviceStateSummary
func (r *DeviceCompliancePolicyDeviceStateSummaryRequest) Update(reqObj *DeviceCompliancePolicyDeviceStateSummary) (*DeviceCompliancePolicyDeviceStateSummary, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceCompliancePolicyDeviceStateSummary
func (r *DeviceCompliancePolicyDeviceStateSummaryRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}
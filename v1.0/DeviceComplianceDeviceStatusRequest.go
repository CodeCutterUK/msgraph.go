// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceComplianceDeviceStatusRequestBuilder is request builder for DeviceComplianceDeviceStatus
type DeviceComplianceDeviceStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceComplianceDeviceStatusRequest
func (b *DeviceComplianceDeviceStatusRequestBuilder) Request() *DeviceComplianceDeviceStatusRequest {
	return &DeviceComplianceDeviceStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceComplianceDeviceStatusRequest is request for DeviceComplianceDeviceStatus
type DeviceComplianceDeviceStatusRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceComplianceDeviceStatus
func (r *DeviceComplianceDeviceStatusRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceComplianceDeviceStatus, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceComplianceDeviceStatus
func (r *DeviceComplianceDeviceStatusRequest) Get() (*DeviceComplianceDeviceStatus, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceComplianceDeviceStatus
func (r *DeviceComplianceDeviceStatusRequest) Update(reqObj *DeviceComplianceDeviceStatus) (*DeviceComplianceDeviceStatus, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceComplianceDeviceStatus
func (r *DeviceComplianceDeviceStatusRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}
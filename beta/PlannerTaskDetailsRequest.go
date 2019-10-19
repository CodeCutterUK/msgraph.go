// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PlannerTaskDetailsRequestBuilder is request builder for PlannerTaskDetails
type PlannerTaskDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerTaskDetailsRequest
func (b *PlannerTaskDetailsRequestBuilder) Request() *PlannerTaskDetailsRequest {
	return &PlannerTaskDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerTaskDetailsRequest is request for PlannerTaskDetails
type PlannerTaskDetailsRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerTaskDetails
func (r *PlannerTaskDetailsRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerTaskDetails, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for PlannerTaskDetails
func (r *PlannerTaskDetailsRequest) Get() (*PlannerTaskDetails, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for PlannerTaskDetails
func (r *PlannerTaskDetailsRequest) Update(reqObj *PlannerTaskDetails) (*PlannerTaskDetails, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for PlannerTaskDetails
func (r *PlannerTaskDetailsRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}
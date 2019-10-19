// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookChartTitleRequestBuilder is request builder for WorkbookChartTitle
type WorkbookChartTitleRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartTitleRequest
func (b *WorkbookChartTitleRequestBuilder) Request() *WorkbookChartTitleRequest {
	return &WorkbookChartTitleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartTitleRequest is request for WorkbookChartTitle
type WorkbookChartTitleRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookChartTitle
func (r *WorkbookChartTitleRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookChartTitle, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookChartTitle
func (r *WorkbookChartTitleRequest) Get() (*WorkbookChartTitle, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookChartTitle
func (r *WorkbookChartTitleRequest) Update(reqObj *WorkbookChartTitle) (*WorkbookChartTitle, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookChartTitle
func (r *WorkbookChartTitleRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Format is navigation property
func (b *WorkbookChartTitleRequestBuilder) Format() *WorkbookChartTitleFormatRequestBuilder {
	bb := &WorkbookChartTitleFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/format"
	return bb
}
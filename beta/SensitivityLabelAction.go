// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SensitivityLabelCollectionEvaluateRequestParameter undocumented
type SensitivityLabelCollectionEvaluateRequestParameter struct {
	// DiscoveredSensitiveTypes undocumented
	DiscoveredSensitiveTypes []DiscoveredSensitiveType `json:"discoveredSensitiveTypes,omitempty"`
}

//
type SensitivityLabelCollectionEvaluateRequestBuilder struct{ BaseRequestBuilder }

// Evaluate action undocumented
func (b *DataClassificationServiceSensitivityLabelsCollectionRequestBuilder) Evaluate(reqObj *SensitivityLabelCollectionEvaluateRequestParameter) *SensitivityLabelCollectionEvaluateRequestBuilder {
	bb := &SensitivityLabelCollectionEvaluateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Evaluate action undocumented
func (b *InformationProtectionSensitivityLabelsCollectionRequestBuilder) Evaluate(reqObj *SensitivityLabelCollectionEvaluateRequestParameter) *SensitivityLabelCollectionEvaluateRequestBuilder {
	bb := &SensitivityLabelCollectionEvaluateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Evaluate action undocumented
func (b *SensitivityLabelSublabelsCollectionRequestBuilder) Evaluate(reqObj *SensitivityLabelCollectionEvaluateRequestParameter) *SensitivityLabelCollectionEvaluateRequestBuilder {
	bb := &SensitivityLabelCollectionEvaluateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SensitivityLabelCollectionEvaluateRequest struct{ BaseRequest }

//
func (b *SensitivityLabelCollectionEvaluateRequestBuilder) Request() *SensitivityLabelCollectionEvaluateRequest {
	return &SensitivityLabelCollectionEvaluateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SensitivityLabelCollectionEvaluateRequest) Do(method, path string, reqObj interface{}) (resObj *EvaluateLabelJobResponse, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *SensitivityLabelCollectionEvaluateRequest) Post() (*EvaluateLabelJobResponse, error) {
	return r.Do("POST", "", r.requestObject)
}
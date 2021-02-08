// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DataClassificationServiceRequestBuilder is request builder for DataClassificationService
type DataClassificationServiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataClassificationServiceRequest
func (b *DataClassificationServiceRequestBuilder) Request() *DataClassificationServiceRequest {
	return &DataClassificationServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataClassificationServiceRequest is request for DataClassificationService
type DataClassificationServiceRequest struct{ BaseRequest }

// Get performs GET request for DataClassificationService
func (r *DataClassificationServiceRequest) Get(ctx context.Context) (resObj *DataClassificationService, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataClassificationService
func (r *DataClassificationServiceRequest) Update(ctx context.Context, reqObj *DataClassificationService) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataClassificationService
func (r *DataClassificationServiceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DataLossPreventionPolicyRequestBuilder is request builder for DataLossPreventionPolicy
type DataLossPreventionPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataLossPreventionPolicyRequest
func (b *DataLossPreventionPolicyRequestBuilder) Request() *DataLossPreventionPolicyRequest {
	return &DataLossPreventionPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataLossPreventionPolicyRequest is request for DataLossPreventionPolicy
type DataLossPreventionPolicyRequest struct{ BaseRequest }

// Get performs GET request for DataLossPreventionPolicy
func (r *DataLossPreventionPolicyRequest) Get(ctx context.Context) (resObj *DataLossPreventionPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataLossPreventionPolicy
func (r *DataLossPreventionPolicyRequest) Update(ctx context.Context, reqObj *DataLossPreventionPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataLossPreventionPolicy
func (r *DataLossPreventionPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DataPolicyOperationRequestBuilder is request builder for DataPolicyOperation
type DataPolicyOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataPolicyOperationRequest
func (b *DataPolicyOperationRequestBuilder) Request() *DataPolicyOperationRequest {
	return &DataPolicyOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataPolicyOperationRequest is request for DataPolicyOperation
type DataPolicyOperationRequest struct{ BaseRequest }

// Get performs GET request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Get(ctx context.Context) (resObj *DataPolicyOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Update(ctx context.Context, reqObj *DataPolicyOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataPolicyOperation
func (r *DataPolicyOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DataSharingConsentRequestBuilder is request builder for DataSharingConsent
type DataSharingConsentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DataSharingConsentRequest
func (b *DataSharingConsentRequestBuilder) Request() *DataSharingConsentRequest {
	return &DataSharingConsentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DataSharingConsentRequest is request for DataSharingConsent
type DataSharingConsentRequest struct{ BaseRequest }

// Get performs GET request for DataSharingConsent
func (r *DataSharingConsentRequest) Get(ctx context.Context) (resObj *DataSharingConsent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DataSharingConsent
func (r *DataSharingConsentRequest) Update(ctx context.Context, reqObj *DataSharingConsent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DataSharingConsent
func (r *DataSharingConsentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type DataLossPreventionPolicyCollectionEvaluateRequestBuilder struct{ BaseRequestBuilder }

// Evaluate action undocumented
func (b *InformationProtectionDataLossPreventionPoliciesCollectionRequestBuilder) Evaluate(reqObj *DataLossPreventionPolicyCollectionEvaluateRequestParameter) *DataLossPreventionPolicyCollectionEvaluateRequestBuilder {
	bb := &DataLossPreventionPolicyCollectionEvaluateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/evaluate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DataLossPreventionPolicyCollectionEvaluateRequest struct{ BaseRequest }

//
func (b *DataLossPreventionPolicyCollectionEvaluateRequestBuilder) Request() *DataLossPreventionPolicyCollectionEvaluateRequest {
	return &DataLossPreventionPolicyCollectionEvaluateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DataLossPreventionPolicyCollectionEvaluateRequest) Post(ctx context.Context) (resObj *DlpEvaluatePoliciesJobResponse, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type DataClassificationServiceClassifyExactMatchesRequestBuilder struct{ BaseRequestBuilder }

// ClassifyExactMatches action undocumented
func (b *DataClassificationServiceRequestBuilder) ClassifyExactMatches(reqObj *DataClassificationServiceClassifyExactMatchesRequestParameter) *DataClassificationServiceClassifyExactMatchesRequestBuilder {
	bb := &DataClassificationServiceClassifyExactMatchesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/classifyExactMatches"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DataClassificationServiceClassifyExactMatchesRequest struct{ BaseRequest }

//
func (b *DataClassificationServiceClassifyExactMatchesRequestBuilder) Request() *DataClassificationServiceClassifyExactMatchesRequest {
	return &DataClassificationServiceClassifyExactMatchesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DataClassificationServiceClassifyExactMatchesRequest) Post(ctx context.Context) (resObj *ExactMatchClassificationResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type DataSharingConsentConsentToDataSharingRequestBuilder struct{ BaseRequestBuilder }

// ConsentToDataSharing action undocumented
func (b *DataSharingConsentRequestBuilder) ConsentToDataSharing(reqObj *DataSharingConsentConsentToDataSharingRequestParameter) *DataSharingConsentConsentToDataSharingRequestBuilder {
	bb := &DataSharingConsentConsentToDataSharingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/consentToDataSharing"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DataSharingConsentConsentToDataSharingRequest struct{ BaseRequest }

//
func (b *DataSharingConsentConsentToDataSharingRequestBuilder) Request() *DataSharingConsentConsentToDataSharingRequest {
	return &DataSharingConsentConsentToDataSharingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DataSharingConsentConsentToDataSharingRequest) Post(ctx context.Context) (resObj *DataSharingConsent, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

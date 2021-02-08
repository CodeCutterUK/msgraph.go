// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ContinuousAccessEvaluationPolicyRequestBuilder is request builder for ContinuousAccessEvaluationPolicy
type ContinuousAccessEvaluationPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ContinuousAccessEvaluationPolicyRequest
func (b *ContinuousAccessEvaluationPolicyRequestBuilder) Request() *ContinuousAccessEvaluationPolicyRequest {
	return &ContinuousAccessEvaluationPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ContinuousAccessEvaluationPolicyRequest is request for ContinuousAccessEvaluationPolicy
type ContinuousAccessEvaluationPolicyRequest struct{ BaseRequest }

// Get performs GET request for ContinuousAccessEvaluationPolicy
func (r *ContinuousAccessEvaluationPolicyRequest) Get(ctx context.Context) (resObj *ContinuousAccessEvaluationPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ContinuousAccessEvaluationPolicy
func (r *ContinuousAccessEvaluationPolicyRequest) Update(ctx context.Context, reqObj *ContinuousAccessEvaluationPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ContinuousAccessEvaluationPolicy
func (r *ContinuousAccessEvaluationPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
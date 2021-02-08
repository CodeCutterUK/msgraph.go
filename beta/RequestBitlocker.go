// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// BitlockerRequestBuilder is request builder for Bitlocker
type BitlockerRequestBuilder struct{ BaseRequestBuilder }

// Request returns BitlockerRequest
func (b *BitlockerRequestBuilder) Request() *BitlockerRequest {
	return &BitlockerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BitlockerRequest is request for Bitlocker
type BitlockerRequest struct{ BaseRequest }

// Get performs GET request for Bitlocker
func (r *BitlockerRequest) Get(ctx context.Context) (resObj *Bitlocker, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Bitlocker
func (r *BitlockerRequest) Update(ctx context.Context, reqObj *Bitlocker) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Bitlocker
func (r *BitlockerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BitlockerRecoveryKeyRequestBuilder is request builder for BitlockerRecoveryKey
type BitlockerRecoveryKeyRequestBuilder struct{ BaseRequestBuilder }

// Request returns BitlockerRecoveryKeyRequest
func (b *BitlockerRecoveryKeyRequestBuilder) Request() *BitlockerRecoveryKeyRequest {
	return &BitlockerRecoveryKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BitlockerRecoveryKeyRequest is request for BitlockerRecoveryKey
type BitlockerRecoveryKeyRequest struct{ BaseRequest }

// Get performs GET request for BitlockerRecoveryKey
func (r *BitlockerRecoveryKeyRequest) Get(ctx context.Context) (resObj *BitlockerRecoveryKey, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BitlockerRecoveryKey
func (r *BitlockerRecoveryKeyRequest) Update(ctx context.Context, reqObj *BitlockerRecoveryKey) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BitlockerRecoveryKey
func (r *BitlockerRecoveryKeyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

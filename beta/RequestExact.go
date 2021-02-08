// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// ExactMatchDataStoreRequestBuilder is request builder for ExactMatchDataStore
type ExactMatchDataStoreRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExactMatchDataStoreRequest
func (b *ExactMatchDataStoreRequestBuilder) Request() *ExactMatchDataStoreRequest {
	return &ExactMatchDataStoreRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExactMatchDataStoreRequest is request for ExactMatchDataStore
type ExactMatchDataStoreRequest struct{ BaseRequest }

// Get performs GET request for ExactMatchDataStore
func (r *ExactMatchDataStoreRequest) Get(ctx context.Context) (resObj *ExactMatchDataStore, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExactMatchDataStore
func (r *ExactMatchDataStoreRequest) Update(ctx context.Context, reqObj *ExactMatchDataStore) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExactMatchDataStore
func (r *ExactMatchDataStoreRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExactMatchLookupJobRequestBuilder is request builder for ExactMatchLookupJob
type ExactMatchLookupJobRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExactMatchLookupJobRequest
func (b *ExactMatchLookupJobRequestBuilder) Request() *ExactMatchLookupJobRequest {
	return &ExactMatchLookupJobRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExactMatchLookupJobRequest is request for ExactMatchLookupJob
type ExactMatchLookupJobRequest struct{ BaseRequest }

// Get performs GET request for ExactMatchLookupJob
func (r *ExactMatchLookupJobRequest) Get(ctx context.Context) (resObj *ExactMatchLookupJob, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExactMatchLookupJob
func (r *ExactMatchLookupJobRequest) Update(ctx context.Context, reqObj *ExactMatchLookupJob) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExactMatchLookupJob
func (r *ExactMatchLookupJobRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExactMatchSessionRequestBuilder is request builder for ExactMatchSession
type ExactMatchSessionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExactMatchSessionRequest
func (b *ExactMatchSessionRequestBuilder) Request() *ExactMatchSessionRequest {
	return &ExactMatchSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExactMatchSessionRequest is request for ExactMatchSession
type ExactMatchSessionRequest struct{ BaseRequest }

// Get performs GET request for ExactMatchSession
func (r *ExactMatchSessionRequest) Get(ctx context.Context) (resObj *ExactMatchSession, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExactMatchSession
func (r *ExactMatchSessionRequest) Update(ctx context.Context, reqObj *ExactMatchSession) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExactMatchSession
func (r *ExactMatchSessionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExactMatchUploadAgentRequestBuilder is request builder for ExactMatchUploadAgent
type ExactMatchUploadAgentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExactMatchUploadAgentRequest
func (b *ExactMatchUploadAgentRequestBuilder) Request() *ExactMatchUploadAgentRequest {
	return &ExactMatchUploadAgentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExactMatchUploadAgentRequest is request for ExactMatchUploadAgent
type ExactMatchUploadAgentRequest struct{ BaseRequest }

// Get performs GET request for ExactMatchUploadAgent
func (r *ExactMatchUploadAgentRequest) Get(ctx context.Context) (resObj *ExactMatchUploadAgent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExactMatchUploadAgent
func (r *ExactMatchUploadAgentRequest) Update(ctx context.Context, reqObj *ExactMatchUploadAgent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExactMatchUploadAgent
func (r *ExactMatchUploadAgentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type ExactMatchDataStoreLookupRequestBuilder struct{ BaseRequestBuilder }

// Lookup action undocumented
func (b *ExactMatchDataStoreRequestBuilder) Lookup(reqObj *ExactMatchDataStoreLookupRequestParameter) *ExactMatchDataStoreLookupRequestBuilder {
	bb := &ExactMatchDataStoreLookupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/lookup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ExactMatchDataStoreLookupRequest struct{ BaseRequest }

//
func (b *ExactMatchDataStoreLookupRequestBuilder) Request() *ExactMatchDataStoreLookupRequest {
	return &ExactMatchDataStoreLookupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ExactMatchDataStoreLookupRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]string, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []string
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []string
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *ExactMatchDataStoreLookupRequest) PostN(ctx context.Context, n int) ([]string, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

//
func (r *ExactMatchDataStoreLookupRequest) Post(ctx context.Context) ([]string, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

//
type ExactMatchSessionCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumented
func (b *ExactMatchSessionRequestBuilder) Cancel(reqObj *ExactMatchSessionCancelRequestParameter) *ExactMatchSessionCancelRequestBuilder {
	bb := &ExactMatchSessionCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ExactMatchSessionCancelRequest struct{ BaseRequest }

//
func (b *ExactMatchSessionCancelRequestBuilder) Request() *ExactMatchSessionCancelRequest {
	return &ExactMatchSessionCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ExactMatchSessionCancelRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ExactMatchSessionRenewRequestBuilder struct{ BaseRequestBuilder }

// Renew action undocumented
func (b *ExactMatchSessionRequestBuilder) Renew(reqObj *ExactMatchSessionRenewRequestParameter) *ExactMatchSessionRenewRequestBuilder {
	bb := &ExactMatchSessionRenewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/renew"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ExactMatchSessionRenewRequest struct{ BaseRequest }

//
func (b *ExactMatchSessionRenewRequestBuilder) Request() *ExactMatchSessionRenewRequest {
	return &ExactMatchSessionRenewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ExactMatchSessionRenewRequest) Post(ctx context.Context) (resObj *ExactMatchSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ExactMatchSessionCommitRequestBuilder struct{ BaseRequestBuilder }

// Commit action undocumented
func (b *ExactMatchSessionRequestBuilder) Commit(reqObj *ExactMatchSessionCommitRequestParameter) *ExactMatchSessionCommitRequestBuilder {
	bb := &ExactMatchSessionCommitRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/commit"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ExactMatchSessionCommitRequest struct{ BaseRequest }

//
func (b *ExactMatchSessionCommitRequestBuilder) Request() *ExactMatchSessionCommitRequest {
	return &ExactMatchSessionCommitRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ExactMatchSessionCommitRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

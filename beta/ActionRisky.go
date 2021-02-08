// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// RiskyUserCollectionConfirmCompromisedRequestParameter undocumented
type RiskyUserCollectionConfirmCompromisedRequestParameter struct {
	// UserIDs undocumented
	UserIDs []string `json:"userIds,omitempty"`
}

// RiskyUserCollectionDismissRequestParameter undocumented
type RiskyUserCollectionDismissRequestParameter struct {
	// UserIDs undocumented
	UserIDs []string `json:"userIds,omitempty"`
}

// History returns request builder for RiskyUserHistoryItem collection
func (b *RiskyUserRequestBuilder) History() *RiskyUserHistoryCollectionRequestBuilder {
	bb := &RiskyUserHistoryCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/history"
	return bb
}

// RiskyUserHistoryCollectionRequestBuilder is request builder for RiskyUserHistoryItem collection
type RiskyUserHistoryCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RiskyUserHistoryItem collection
func (b *RiskyUserHistoryCollectionRequestBuilder) Request() *RiskyUserHistoryCollectionRequest {
	return &RiskyUserHistoryCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RiskyUserHistoryItem item
func (b *RiskyUserHistoryCollectionRequestBuilder) ID(id string) *RiskyUserHistoryItemRequestBuilder {
	bb := &RiskyUserHistoryItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RiskyUserHistoryCollectionRequest is request for RiskyUserHistoryItem collection
type RiskyUserHistoryCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RiskyUserHistoryItem collection
func (r *RiskyUserHistoryCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RiskyUserHistoryItem, error) {
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
	var values []RiskyUserHistoryItem
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
			value  []RiskyUserHistoryItem
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

// GetN performs GET request for RiskyUserHistoryItem collection, max N pages
func (r *RiskyUserHistoryCollectionRequest) GetN(ctx context.Context, n int) ([]RiskyUserHistoryItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RiskyUserHistoryItem collection
func (r *RiskyUserHistoryCollectionRequest) Get(ctx context.Context) ([]RiskyUserHistoryItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RiskyUserHistoryItem collection
func (r *RiskyUserHistoryCollectionRequest) Add(ctx context.Context, reqObj *RiskyUserHistoryItem) (resObj *RiskyUserHistoryItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

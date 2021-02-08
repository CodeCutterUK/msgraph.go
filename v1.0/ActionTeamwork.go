// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// WorkforceIntegrations returns request builder for WorkforceIntegration collection
func (b *TeamworkRequestBuilder) WorkforceIntegrations() *TeamworkWorkforceIntegrationsCollectionRequestBuilder {
	bb := &TeamworkWorkforceIntegrationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/workforceIntegrations"
	return bb
}

// TeamworkWorkforceIntegrationsCollectionRequestBuilder is request builder for WorkforceIntegration collection
type TeamworkWorkforceIntegrationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkforceIntegration collection
func (b *TeamworkWorkforceIntegrationsCollectionRequestBuilder) Request() *TeamworkWorkforceIntegrationsCollectionRequest {
	return &TeamworkWorkforceIntegrationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkforceIntegration item
func (b *TeamworkWorkforceIntegrationsCollectionRequestBuilder) ID(id string) *WorkforceIntegrationRequestBuilder {
	bb := &WorkforceIntegrationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkWorkforceIntegrationsCollectionRequest is request for WorkforceIntegration collection
type TeamworkWorkforceIntegrationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]WorkforceIntegration, error) {
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
	var values []WorkforceIntegration
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
			value  []WorkforceIntegration
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

// GetN performs GET request for WorkforceIntegration collection, max N pages
func (r *TeamworkWorkforceIntegrationsCollectionRequest) GetN(ctx context.Context, n int) ([]WorkforceIntegration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Get(ctx context.Context) ([]WorkforceIntegration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Add(ctx context.Context, reqObj *WorkforceIntegration) (resObj *WorkforceIntegration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

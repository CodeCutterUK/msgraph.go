// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// Companies returns request builder for Company collection
func (b *FinancialsRequestBuilder) Companies() *FinancialsCompaniesCollectionRequestBuilder {
	bb := &FinancialsCompaniesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/companies"
	return bb
}

// FinancialsCompaniesCollectionRequestBuilder is request builder for Company collection
type FinancialsCompaniesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Company collection
func (b *FinancialsCompaniesCollectionRequestBuilder) Request() *FinancialsCompaniesCollectionRequest {
	return &FinancialsCompaniesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Company item
func (b *FinancialsCompaniesCollectionRequestBuilder) ID(id string) *CompanyRequestBuilder {
	bb := &CompanyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// FinancialsCompaniesCollectionRequest is request for Company collection
type FinancialsCompaniesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Company collection
func (r *FinancialsCompaniesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Company, error) {
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
	var values []Company
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
			value  []Company
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

// GetN performs GET request for Company collection, max N pages
func (r *FinancialsCompaniesCollectionRequest) GetN(ctx context.Context, n int) ([]Company, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Company collection
func (r *FinancialsCompaniesCollectionRequest) Get(ctx context.Context) ([]Company, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Company collection
func (r *FinancialsCompaniesCollectionRequest) Add(ctx context.Context, reqObj *Company) (resObj *Company, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

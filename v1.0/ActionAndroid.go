// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// Apps returns request builder for ManagedMobileApp collection
func (b *AndroidManagedAppProtectionRequestBuilder) Apps() *AndroidManagedAppProtectionAppsCollectionRequestBuilder {
	bb := &AndroidManagedAppProtectionAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/apps"
	return bb
}

// AndroidManagedAppProtectionAppsCollectionRequestBuilder is request builder for ManagedMobileApp collection
type AndroidManagedAppProtectionAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedMobileApp collection
func (b *AndroidManagedAppProtectionAppsCollectionRequestBuilder) Request() *AndroidManagedAppProtectionAppsCollectionRequest {
	return &AndroidManagedAppProtectionAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedMobileApp item
func (b *AndroidManagedAppProtectionAppsCollectionRequestBuilder) ID(id string) *ManagedMobileAppRequestBuilder {
	bb := &ManagedMobileAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AndroidManagedAppProtectionAppsCollectionRequest is request for ManagedMobileApp collection
type AndroidManagedAppProtectionAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedMobileApp collection
func (r *AndroidManagedAppProtectionAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagedMobileApp, error) {
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
	var values []ManagedMobileApp
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
			value  []ManagedMobileApp
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

// GetN performs GET request for ManagedMobileApp collection, max N pages
func (r *AndroidManagedAppProtectionAppsCollectionRequest) GetN(ctx context.Context, n int) ([]ManagedMobileApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagedMobileApp collection
func (r *AndroidManagedAppProtectionAppsCollectionRequest) Get(ctx context.Context) ([]ManagedMobileApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagedMobileApp collection
func (r *AndroidManagedAppProtectionAppsCollectionRequest) Add(ctx context.Context, reqObj *ManagedMobileApp) (resObj *ManagedMobileApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeploymentSummary is navigation property
func (b *AndroidManagedAppProtectionRequestBuilder) DeploymentSummary() *ManagedAppPolicyDeploymentSummaryRequestBuilder {
	bb := &ManagedAppPolicyDeploymentSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deploymentSummary"
	return bb
}

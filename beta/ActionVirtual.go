// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// CloudPCs returns request builder for CloudPC collection
func (b *VirtualEndpointRequestBuilder) CloudPCs() *VirtualEndpointCloudPCsCollectionRequestBuilder {
	bb := &VirtualEndpointCloudPCsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/cloudPCs"
	return bb
}

// VirtualEndpointCloudPCsCollectionRequestBuilder is request builder for CloudPC collection
type VirtualEndpointCloudPCsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CloudPC collection
func (b *VirtualEndpointCloudPCsCollectionRequestBuilder) Request() *VirtualEndpointCloudPCsCollectionRequest {
	return &VirtualEndpointCloudPCsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CloudPC item
func (b *VirtualEndpointCloudPCsCollectionRequestBuilder) ID(id string) *CloudPCRequestBuilder {
	bb := &CloudPCRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// VirtualEndpointCloudPCsCollectionRequest is request for CloudPC collection
type VirtualEndpointCloudPCsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CloudPC collection
func (r *VirtualEndpointCloudPCsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CloudPC, error) {
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
	var values []CloudPC
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
			value  []CloudPC
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

// GetN performs GET request for CloudPC collection, max N pages
func (r *VirtualEndpointCloudPCsCollectionRequest) GetN(ctx context.Context, n int) ([]CloudPC, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CloudPC collection
func (r *VirtualEndpointCloudPCsCollectionRequest) Get(ctx context.Context) ([]CloudPC, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CloudPC collection
func (r *VirtualEndpointCloudPCsCollectionRequest) Add(ctx context.Context, reqObj *CloudPC) (resObj *CloudPC, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceImages returns request builder for CloudPcDeviceImage collection
func (b *VirtualEndpointRequestBuilder) DeviceImages() *VirtualEndpointDeviceImagesCollectionRequestBuilder {
	bb := &VirtualEndpointDeviceImagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceImages"
	return bb
}

// VirtualEndpointDeviceImagesCollectionRequestBuilder is request builder for CloudPcDeviceImage collection
type VirtualEndpointDeviceImagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CloudPcDeviceImage collection
func (b *VirtualEndpointDeviceImagesCollectionRequestBuilder) Request() *VirtualEndpointDeviceImagesCollectionRequest {
	return &VirtualEndpointDeviceImagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CloudPcDeviceImage item
func (b *VirtualEndpointDeviceImagesCollectionRequestBuilder) ID(id string) *CloudPcDeviceImageRequestBuilder {
	bb := &CloudPcDeviceImageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// VirtualEndpointDeviceImagesCollectionRequest is request for CloudPcDeviceImage collection
type VirtualEndpointDeviceImagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CloudPcDeviceImage collection
func (r *VirtualEndpointDeviceImagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CloudPcDeviceImage, error) {
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
	var values []CloudPcDeviceImage
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
			value  []CloudPcDeviceImage
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

// GetN performs GET request for CloudPcDeviceImage collection, max N pages
func (r *VirtualEndpointDeviceImagesCollectionRequest) GetN(ctx context.Context, n int) ([]CloudPcDeviceImage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CloudPcDeviceImage collection
func (r *VirtualEndpointDeviceImagesCollectionRequest) Get(ctx context.Context) ([]CloudPcDeviceImage, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CloudPcDeviceImage collection
func (r *VirtualEndpointDeviceImagesCollectionRequest) Add(ctx context.Context, reqObj *CloudPcDeviceImage) (resObj *CloudPcDeviceImage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// OnPremisesConnections returns request builder for CloudPcOnPremisesConnection collection
func (b *VirtualEndpointRequestBuilder) OnPremisesConnections() *VirtualEndpointOnPremisesConnectionsCollectionRequestBuilder {
	bb := &VirtualEndpointOnPremisesConnectionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/onPremisesConnections"
	return bb
}

// VirtualEndpointOnPremisesConnectionsCollectionRequestBuilder is request builder for CloudPcOnPremisesConnection collection
type VirtualEndpointOnPremisesConnectionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CloudPcOnPremisesConnection collection
func (b *VirtualEndpointOnPremisesConnectionsCollectionRequestBuilder) Request() *VirtualEndpointOnPremisesConnectionsCollectionRequest {
	return &VirtualEndpointOnPremisesConnectionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CloudPcOnPremisesConnection item
func (b *VirtualEndpointOnPremisesConnectionsCollectionRequestBuilder) ID(id string) *CloudPcOnPremisesConnectionRequestBuilder {
	bb := &CloudPcOnPremisesConnectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// VirtualEndpointOnPremisesConnectionsCollectionRequest is request for CloudPcOnPremisesConnection collection
type VirtualEndpointOnPremisesConnectionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CloudPcOnPremisesConnection collection
func (r *VirtualEndpointOnPremisesConnectionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CloudPcOnPremisesConnection, error) {
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
	var values []CloudPcOnPremisesConnection
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
			value  []CloudPcOnPremisesConnection
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

// GetN performs GET request for CloudPcOnPremisesConnection collection, max N pages
func (r *VirtualEndpointOnPremisesConnectionsCollectionRequest) GetN(ctx context.Context, n int) ([]CloudPcOnPremisesConnection, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CloudPcOnPremisesConnection collection
func (r *VirtualEndpointOnPremisesConnectionsCollectionRequest) Get(ctx context.Context) ([]CloudPcOnPremisesConnection, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CloudPcOnPremisesConnection collection
func (r *VirtualEndpointOnPremisesConnectionsCollectionRequest) Add(ctx context.Context, reqObj *CloudPcOnPremisesConnection) (resObj *CloudPcOnPremisesConnection, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ProvisioningPolicies returns request builder for CloudPcProvisioningPolicy collection
func (b *VirtualEndpointRequestBuilder) ProvisioningPolicies() *VirtualEndpointProvisioningPoliciesCollectionRequestBuilder {
	bb := &VirtualEndpointProvisioningPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/provisioningPolicies"
	return bb
}

// VirtualEndpointProvisioningPoliciesCollectionRequestBuilder is request builder for CloudPcProvisioningPolicy collection
type VirtualEndpointProvisioningPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CloudPcProvisioningPolicy collection
func (b *VirtualEndpointProvisioningPoliciesCollectionRequestBuilder) Request() *VirtualEndpointProvisioningPoliciesCollectionRequest {
	return &VirtualEndpointProvisioningPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CloudPcProvisioningPolicy item
func (b *VirtualEndpointProvisioningPoliciesCollectionRequestBuilder) ID(id string) *CloudPcProvisioningPolicyRequestBuilder {
	bb := &CloudPcProvisioningPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// VirtualEndpointProvisioningPoliciesCollectionRequest is request for CloudPcProvisioningPolicy collection
type VirtualEndpointProvisioningPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CloudPcProvisioningPolicy collection
func (r *VirtualEndpointProvisioningPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CloudPcProvisioningPolicy, error) {
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
	var values []CloudPcProvisioningPolicy
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
			value  []CloudPcProvisioningPolicy
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

// GetN performs GET request for CloudPcProvisioningPolicy collection, max N pages
func (r *VirtualEndpointProvisioningPoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]CloudPcProvisioningPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CloudPcProvisioningPolicy collection
func (r *VirtualEndpointProvisioningPoliciesCollectionRequest) Get(ctx context.Context) ([]CloudPcProvisioningPolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CloudPcProvisioningPolicy collection
func (r *VirtualEndpointProvisioningPoliciesCollectionRequest) Add(ctx context.Context, reqObj *CloudPcProvisioningPolicy) (resObj *CloudPcProvisioningPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

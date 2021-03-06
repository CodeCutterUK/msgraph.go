// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// Fido2AuthenticationMethodRequestBuilder is request builder for Fido2AuthenticationMethod
type Fido2AuthenticationMethodRequestBuilder struct{ BaseRequestBuilder }

// Request returns Fido2AuthenticationMethodRequest
func (b *Fido2AuthenticationMethodRequestBuilder) Request() *Fido2AuthenticationMethodRequest {
	return &Fido2AuthenticationMethodRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Fido2AuthenticationMethodRequest is request for Fido2AuthenticationMethod
type Fido2AuthenticationMethodRequest struct{ BaseRequest }

// Get performs GET request for Fido2AuthenticationMethod
func (r *Fido2AuthenticationMethodRequest) Get(ctx context.Context) (resObj *Fido2AuthenticationMethod, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Fido2AuthenticationMethod
func (r *Fido2AuthenticationMethodRequest) Update(ctx context.Context, reqObj *Fido2AuthenticationMethod) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Fido2AuthenticationMethod
func (r *Fido2AuthenticationMethodRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Fido2AuthenticationMethodConfigurationRequestBuilder is request builder for Fido2AuthenticationMethodConfiguration
type Fido2AuthenticationMethodConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns Fido2AuthenticationMethodConfigurationRequest
func (b *Fido2AuthenticationMethodConfigurationRequestBuilder) Request() *Fido2AuthenticationMethodConfigurationRequest {
	return &Fido2AuthenticationMethodConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Fido2AuthenticationMethodConfigurationRequest is request for Fido2AuthenticationMethodConfiguration
type Fido2AuthenticationMethodConfigurationRequest struct{ BaseRequest }

// Get performs GET request for Fido2AuthenticationMethodConfiguration
func (r *Fido2AuthenticationMethodConfigurationRequest) Get(ctx context.Context) (resObj *Fido2AuthenticationMethodConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Fido2AuthenticationMethodConfiguration
func (r *Fido2AuthenticationMethodConfigurationRequest) Update(ctx context.Context, reqObj *Fido2AuthenticationMethodConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Fido2AuthenticationMethodConfiguration
func (r *Fido2AuthenticationMethodConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

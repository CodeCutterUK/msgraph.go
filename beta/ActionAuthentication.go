// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// AuthenticationMethodDisableSmsSignInRequestParameter undocumented
type AuthenticationMethodDisableSmsSignInRequestParameter struct {
}

// AuthenticationMethodEnableSmsSignInRequestParameter undocumented
type AuthenticationMethodEnableSmsSignInRequestParameter struct {
}

// AuthenticationMethodResetPasswordRequestParameter undocumented
type AuthenticationMethodResetPasswordRequestParameter struct {
	// NewPassword undocumented
	NewPassword *string `json:"newPassword,omitempty"`
	// RequireChangeOnNextSignIn undocumented
	RequireChangeOnNextSignIn *bool `json:"requireChangeOnNextSignIn,omitempty"`
}

// EmailMethods returns request builder for EmailAuthenticationMethod collection
func (b *AuthenticationRequestBuilder) EmailMethods() *AuthenticationEmailMethodsCollectionRequestBuilder {
	bb := &AuthenticationEmailMethodsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/emailMethods"
	return bb
}

// AuthenticationEmailMethodsCollectionRequestBuilder is request builder for EmailAuthenticationMethod collection
type AuthenticationEmailMethodsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmailAuthenticationMethod collection
func (b *AuthenticationEmailMethodsCollectionRequestBuilder) Request() *AuthenticationEmailMethodsCollectionRequest {
	return &AuthenticationEmailMethodsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmailAuthenticationMethod item
func (b *AuthenticationEmailMethodsCollectionRequestBuilder) ID(id string) *EmailAuthenticationMethodRequestBuilder {
	bb := &EmailAuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationEmailMethodsCollectionRequest is request for EmailAuthenticationMethod collection
type AuthenticationEmailMethodsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmailAuthenticationMethod collection
func (r *AuthenticationEmailMethodsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EmailAuthenticationMethod, error) {
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
	var values []EmailAuthenticationMethod
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
			value  []EmailAuthenticationMethod
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

// GetN performs GET request for EmailAuthenticationMethod collection, max N pages
func (r *AuthenticationEmailMethodsCollectionRequest) GetN(ctx context.Context, n int) ([]EmailAuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EmailAuthenticationMethod collection
func (r *AuthenticationEmailMethodsCollectionRequest) Get(ctx context.Context) ([]EmailAuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EmailAuthenticationMethod collection
func (r *AuthenticationEmailMethodsCollectionRequest) Add(ctx context.Context, reqObj *EmailAuthenticationMethod) (resObj *EmailAuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Fido2Methods returns request builder for Fido2AuthenticationMethod collection
func (b *AuthenticationRequestBuilder) Fido2Methods() *AuthenticationFido2MethodsCollectionRequestBuilder {
	bb := &AuthenticationFido2MethodsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fido2Methods"
	return bb
}

// AuthenticationFido2MethodsCollectionRequestBuilder is request builder for Fido2AuthenticationMethod collection
type AuthenticationFido2MethodsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Fido2AuthenticationMethod collection
func (b *AuthenticationFido2MethodsCollectionRequestBuilder) Request() *AuthenticationFido2MethodsCollectionRequest {
	return &AuthenticationFido2MethodsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Fido2AuthenticationMethod item
func (b *AuthenticationFido2MethodsCollectionRequestBuilder) ID(id string) *Fido2AuthenticationMethodRequestBuilder {
	bb := &Fido2AuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationFido2MethodsCollectionRequest is request for Fido2AuthenticationMethod collection
type AuthenticationFido2MethodsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Fido2AuthenticationMethod collection
func (r *AuthenticationFido2MethodsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Fido2AuthenticationMethod, error) {
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
	var values []Fido2AuthenticationMethod
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
			value  []Fido2AuthenticationMethod
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

// GetN performs GET request for Fido2AuthenticationMethod collection, max N pages
func (r *AuthenticationFido2MethodsCollectionRequest) GetN(ctx context.Context, n int) ([]Fido2AuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Fido2AuthenticationMethod collection
func (r *AuthenticationFido2MethodsCollectionRequest) Get(ctx context.Context) ([]Fido2AuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Fido2AuthenticationMethod collection
func (r *AuthenticationFido2MethodsCollectionRequest) Add(ctx context.Context, reqObj *Fido2AuthenticationMethod) (resObj *Fido2AuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Methods returns request builder for AuthenticationMethod collection
func (b *AuthenticationRequestBuilder) Methods() *AuthenticationMethodsCollectionRequestBuilder {
	bb := &AuthenticationMethodsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/methods"
	return bb
}

// AuthenticationMethodsCollectionRequestBuilder is request builder for AuthenticationMethod collection
type AuthenticationMethodsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AuthenticationMethod collection
func (b *AuthenticationMethodsCollectionRequestBuilder) Request() *AuthenticationMethodsCollectionRequest {
	return &AuthenticationMethodsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AuthenticationMethod item
func (b *AuthenticationMethodsCollectionRequestBuilder) ID(id string) *AuthenticationMethodRequestBuilder {
	bb := &AuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationMethodsCollectionRequest is request for AuthenticationMethod collection
type AuthenticationMethodsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AuthenticationMethod collection
func (r *AuthenticationMethodsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AuthenticationMethod, error) {
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
	var values []AuthenticationMethod
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
			value  []AuthenticationMethod
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

// GetN performs GET request for AuthenticationMethod collection, max N pages
func (r *AuthenticationMethodsCollectionRequest) GetN(ctx context.Context, n int) ([]AuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AuthenticationMethod collection
func (r *AuthenticationMethodsCollectionRequest) Get(ctx context.Context) ([]AuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AuthenticationMethod collection
func (r *AuthenticationMethodsCollectionRequest) Add(ctx context.Context, reqObj *AuthenticationMethod) (resObj *AuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Operations returns request builder for LongRunningOperation collection
func (b *AuthenticationRequestBuilder) Operations() *AuthenticationOperationsCollectionRequestBuilder {
	bb := &AuthenticationOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// AuthenticationOperationsCollectionRequestBuilder is request builder for LongRunningOperation collection
type AuthenticationOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for LongRunningOperation collection
func (b *AuthenticationOperationsCollectionRequestBuilder) Request() *AuthenticationOperationsCollectionRequest {
	return &AuthenticationOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for LongRunningOperation item
func (b *AuthenticationOperationsCollectionRequestBuilder) ID(id string) *LongRunningOperationRequestBuilder {
	bb := &LongRunningOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationOperationsCollectionRequest is request for LongRunningOperation collection
type AuthenticationOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for LongRunningOperation collection
func (r *AuthenticationOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]LongRunningOperation, error) {
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
	var values []LongRunningOperation
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
			value  []LongRunningOperation
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

// GetN performs GET request for LongRunningOperation collection, max N pages
func (r *AuthenticationOperationsCollectionRequest) GetN(ctx context.Context, n int) ([]LongRunningOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for LongRunningOperation collection
func (r *AuthenticationOperationsCollectionRequest) Get(ctx context.Context) ([]LongRunningOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for LongRunningOperation collection
func (r *AuthenticationOperationsCollectionRequest) Add(ctx context.Context, reqObj *LongRunningOperation) (resObj *LongRunningOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PasswordMethods returns request builder for PasswordAuthenticationMethod collection
func (b *AuthenticationRequestBuilder) PasswordMethods() *AuthenticationPasswordMethodsCollectionRequestBuilder {
	bb := &AuthenticationPasswordMethodsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/passwordMethods"
	return bb
}

// AuthenticationPasswordMethodsCollectionRequestBuilder is request builder for PasswordAuthenticationMethod collection
type AuthenticationPasswordMethodsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PasswordAuthenticationMethod collection
func (b *AuthenticationPasswordMethodsCollectionRequestBuilder) Request() *AuthenticationPasswordMethodsCollectionRequest {
	return &AuthenticationPasswordMethodsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PasswordAuthenticationMethod item
func (b *AuthenticationPasswordMethodsCollectionRequestBuilder) ID(id string) *PasswordAuthenticationMethodRequestBuilder {
	bb := &PasswordAuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationPasswordMethodsCollectionRequest is request for PasswordAuthenticationMethod collection
type AuthenticationPasswordMethodsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PasswordAuthenticationMethod collection
func (r *AuthenticationPasswordMethodsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PasswordAuthenticationMethod, error) {
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
	var values []PasswordAuthenticationMethod
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
			value  []PasswordAuthenticationMethod
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

// GetN performs GET request for PasswordAuthenticationMethod collection, max N pages
func (r *AuthenticationPasswordMethodsCollectionRequest) GetN(ctx context.Context, n int) ([]PasswordAuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PasswordAuthenticationMethod collection
func (r *AuthenticationPasswordMethodsCollectionRequest) Get(ctx context.Context) ([]PasswordAuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PasswordAuthenticationMethod collection
func (r *AuthenticationPasswordMethodsCollectionRequest) Add(ctx context.Context, reqObj *PasswordAuthenticationMethod) (resObj *PasswordAuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PasswordlessMicrosoftAuthenticatorMethods returns request builder for PasswordlessMicrosoftAuthenticatorAuthenticationMethod collection
func (b *AuthenticationRequestBuilder) PasswordlessMicrosoftAuthenticatorMethods() *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequestBuilder {
	bb := &AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/passwordlessMicrosoftAuthenticatorMethods"
	return bb
}

// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequestBuilder is request builder for PasswordlessMicrosoftAuthenticatorAuthenticationMethod collection
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PasswordlessMicrosoftAuthenticatorAuthenticationMethod collection
func (b *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequestBuilder) Request() *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequest {
	return &AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PasswordlessMicrosoftAuthenticatorAuthenticationMethod item
func (b *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequestBuilder) ID(id string) *PasswordlessMicrosoftAuthenticatorAuthenticationMethodRequestBuilder {
	bb := &PasswordlessMicrosoftAuthenticatorAuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequest is request for PasswordlessMicrosoftAuthenticatorAuthenticationMethod collection
type AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PasswordlessMicrosoftAuthenticatorAuthenticationMethod collection
func (r *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PasswordlessMicrosoftAuthenticatorAuthenticationMethod, error) {
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
	var values []PasswordlessMicrosoftAuthenticatorAuthenticationMethod
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
			value  []PasswordlessMicrosoftAuthenticatorAuthenticationMethod
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

// GetN performs GET request for PasswordlessMicrosoftAuthenticatorAuthenticationMethod collection, max N pages
func (r *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequest) GetN(ctx context.Context, n int) ([]PasswordlessMicrosoftAuthenticatorAuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PasswordlessMicrosoftAuthenticatorAuthenticationMethod collection
func (r *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequest) Get(ctx context.Context) ([]PasswordlessMicrosoftAuthenticatorAuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PasswordlessMicrosoftAuthenticatorAuthenticationMethod collection
func (r *AuthenticationPasswordlessMicrosoftAuthenticatorMethodsCollectionRequest) Add(ctx context.Context, reqObj *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) (resObj *PasswordlessMicrosoftAuthenticatorAuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PhoneMethods returns request builder for PhoneAuthenticationMethod collection
func (b *AuthenticationRequestBuilder) PhoneMethods() *AuthenticationPhoneMethodsCollectionRequestBuilder {
	bb := &AuthenticationPhoneMethodsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/phoneMethods"
	return bb
}

// AuthenticationPhoneMethodsCollectionRequestBuilder is request builder for PhoneAuthenticationMethod collection
type AuthenticationPhoneMethodsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PhoneAuthenticationMethod collection
func (b *AuthenticationPhoneMethodsCollectionRequestBuilder) Request() *AuthenticationPhoneMethodsCollectionRequest {
	return &AuthenticationPhoneMethodsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PhoneAuthenticationMethod item
func (b *AuthenticationPhoneMethodsCollectionRequestBuilder) ID(id string) *PhoneAuthenticationMethodRequestBuilder {
	bb := &PhoneAuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationPhoneMethodsCollectionRequest is request for PhoneAuthenticationMethod collection
type AuthenticationPhoneMethodsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PhoneAuthenticationMethod collection
func (r *AuthenticationPhoneMethodsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PhoneAuthenticationMethod, error) {
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
	var values []PhoneAuthenticationMethod
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
			value  []PhoneAuthenticationMethod
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

// GetN performs GET request for PhoneAuthenticationMethod collection, max N pages
func (r *AuthenticationPhoneMethodsCollectionRequest) GetN(ctx context.Context, n int) ([]PhoneAuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PhoneAuthenticationMethod collection
func (r *AuthenticationPhoneMethodsCollectionRequest) Get(ctx context.Context) ([]PhoneAuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PhoneAuthenticationMethod collection
func (r *AuthenticationPhoneMethodsCollectionRequest) Add(ctx context.Context, reqObj *PhoneAuthenticationMethod) (resObj *PhoneAuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

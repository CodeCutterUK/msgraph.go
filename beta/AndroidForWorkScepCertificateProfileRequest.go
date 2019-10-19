// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AndroidForWorkScepCertificateProfileRequestBuilder is request builder for AndroidForWorkScepCertificateProfile
type AndroidForWorkScepCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkScepCertificateProfileRequest
func (b *AndroidForWorkScepCertificateProfileRequestBuilder) Request() *AndroidForWorkScepCertificateProfileRequest {
	return &AndroidForWorkScepCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidForWorkScepCertificateProfileRequest is request for AndroidForWorkScepCertificateProfile
type AndroidForWorkScepCertificateProfileRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidForWorkScepCertificateProfile
func (r *AndroidForWorkScepCertificateProfileRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidForWorkScepCertificateProfile, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidForWorkScepCertificateProfile
func (r *AndroidForWorkScepCertificateProfileRequest) Get() (*AndroidForWorkScepCertificateProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidForWorkScepCertificateProfile
func (r *AndroidForWorkScepCertificateProfileRequest) Update(reqObj *AndroidForWorkScepCertificateProfile) (*AndroidForWorkScepCertificateProfile, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidForWorkScepCertificateProfile
func (r *AndroidForWorkScepCertificateProfileRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *AndroidForWorkScepCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedDeviceCertificateState collection
func (r *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagedDeviceCertificateState
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ManagedDeviceCertificateState
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get() ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(reqObj *ManagedDeviceCertificateState) (*ManagedDeviceCertificateState, error) {
	return r.Do("POST", "", reqObj)
}
// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequestParameter undocumented
type WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// WindowsAutopilotDeploymentProfileAssignRequestParameter undocumented
type WindowsAutopilotDeploymentProfileAssignRequestParameter struct {
	// DeviceIDs undocumented
	DeviceIDs []string `json:"deviceIds,omitempty"`
}

//
type WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceManagementWindowsAutopilotDeploymentProfilesCollectionRequestBuilder) HasPayloadLinks(reqObj *WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequestParameter) *WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequestBuilder {
	bb := &WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequestBuilder) Request() *WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequest {
	return &WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequest) Do(method, path string, reqObj interface{}) (resObj *[]HasPayloadLinkResultItem, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequest) Paging(method, path string, obj interface{}) ([][]HasPayloadLinkResultItem, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]HasPayloadLinkResultItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]HasPayloadLinkResultItem
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

//
func (r *WindowsAutopilotDeploymentProfileCollectionHasPayloadLinksRequest) Get() ([][]HasPayloadLinkResultItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type WindowsAutopilotDeploymentProfileAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *WindowsAutopilotDeploymentProfileRequestBuilder) Assign(reqObj *WindowsAutopilotDeploymentProfileAssignRequestParameter) *WindowsAutopilotDeploymentProfileAssignRequestBuilder {
	bb := &WindowsAutopilotDeploymentProfileAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsAutopilotDeploymentProfileAssignRequest struct{ BaseRequest }

//
func (b *WindowsAutopilotDeploymentProfileAssignRequestBuilder) Request() *WindowsAutopilotDeploymentProfileAssignRequest {
	return &WindowsAutopilotDeploymentProfileAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsAutopilotDeploymentProfileAssignRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *WindowsAutopilotDeploymentProfileAssignRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}
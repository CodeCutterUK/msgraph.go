// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ListRequestBuilder is request builder for List
type ListRequestBuilder struct{ BaseRequestBuilder }

// Request returns ListRequest
func (b *ListRequestBuilder) Request() *ListRequest {
	return &ListRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ListRequest is request for List
type ListRequest struct{ BaseRequest }

// Do performs HTTP request for List
func (r *ListRequest) Do(method, path string, reqObj interface{}) (resObj *List, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for List
func (r *ListRequest) Get() (*List, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for List
func (r *ListRequest) Update(reqObj *List) (*List, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for List
func (r *ListRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Activities returns request builder for ItemActivityOLD collection
func (b *ListRequestBuilder) Activities() *ListActivitiesCollectionRequestBuilder {
	bb := &ListActivitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activities"
	return bb
}

// ListActivitiesCollectionRequestBuilder is request builder for ItemActivityOLD collection
type ListActivitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivityOLD collection
func (b *ListActivitiesCollectionRequestBuilder) Request() *ListActivitiesCollectionRequest {
	return &ListActivitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivityOLD item
func (b *ListActivitiesCollectionRequestBuilder) ID(id string) *ItemActivityOLDRequestBuilder {
	bb := &ItemActivityOLDRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListActivitiesCollectionRequest is request for ItemActivityOLD collection
type ListActivitiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ItemActivityOLD collection
func (r *ListActivitiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ItemActivityOLD, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ItemActivityOLD collection
func (r *ListActivitiesCollectionRequest) Paging(method, path string, obj interface{}) ([]ItemActivityOLD, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ItemActivityOLD
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ItemActivityOLD
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

// Get performs GET request for ItemActivityOLD collection
func (r *ListActivitiesCollectionRequest) Get() ([]ItemActivityOLD, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ItemActivityOLD collection
func (r *ListActivitiesCollectionRequest) Add(reqObj *ItemActivityOLD) (*ItemActivityOLD, error) {
	return r.Do("POST", "", reqObj)
}

// Columns returns request builder for ColumnDefinition collection
func (b *ListRequestBuilder) Columns() *ListColumnsCollectionRequestBuilder {
	bb := &ListColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columns"
	return bb
}

// ListColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection
type ListColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *ListColumnsCollectionRequestBuilder) Request() *ListColumnsCollectionRequest {
	return &ListColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *ListColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListColumnsCollectionRequest is request for ColumnDefinition collection
type ListColumnsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Paging(method, path string, obj interface{}) ([]ColumnDefinition, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ColumnDefinition
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ColumnDefinition
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

// Get performs GET request for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Get() ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Add(reqObj *ColumnDefinition) (*ColumnDefinition, error) {
	return r.Do("POST", "", reqObj)
}

// ContentTypes returns request builder for ContentType collection
func (b *ListRequestBuilder) ContentTypes() *ListContentTypesCollectionRequestBuilder {
	bb := &ListContentTypesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contentTypes"
	return bb
}

// ListContentTypesCollectionRequestBuilder is request builder for ContentType collection
type ListContentTypesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ContentType collection
func (b *ListContentTypesCollectionRequestBuilder) Request() *ListContentTypesCollectionRequest {
	return &ListContentTypesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ContentType item
func (b *ListContentTypesCollectionRequestBuilder) ID(id string) *ContentTypeRequestBuilder {
	bb := &ContentTypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListContentTypesCollectionRequest is request for ContentType collection
type ListContentTypesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ContentType collection
func (r *ListContentTypesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ContentType, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ContentType collection
func (r *ListContentTypesCollectionRequest) Paging(method, path string, obj interface{}) ([]ContentType, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ContentType
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ContentType
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

// Get performs GET request for ContentType collection
func (r *ListContentTypesCollectionRequest) Get() ([]ContentType, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ContentType collection
func (r *ListContentTypesCollectionRequest) Add(reqObj *ContentType) (*ContentType, error) {
	return r.Do("POST", "", reqObj)
}

// Drive is navigation property
func (b *ListRequestBuilder) Drive() *DriveRequestBuilder {
	bb := &DriveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/drive"
	return bb
}

// Items returns request builder for ListItem collection
func (b *ListRequestBuilder) Items() *ListItemsCollectionRequestBuilder {
	bb := &ListItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// ListItemsCollectionRequestBuilder is request builder for ListItem collection
type ListItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ListItem collection
func (b *ListItemsCollectionRequestBuilder) Request() *ListItemsCollectionRequest {
	return &ListItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ListItem item
func (b *ListItemsCollectionRequestBuilder) ID(id string) *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemsCollectionRequest is request for ListItem collection
type ListItemsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ListItem collection
func (r *ListItemsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ListItem, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ListItem collection
func (r *ListItemsCollectionRequest) Paging(method, path string, obj interface{}) ([]ListItem, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ListItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ListItem
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

// Get performs GET request for ListItem collection
func (r *ListItemsCollectionRequest) Get() ([]ListItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ListItem collection
func (r *ListItemsCollectionRequest) Add(reqObj *ListItem) (*ListItem, error) {
	return r.Do("POST", "", reqObj)
}

// Subscriptions returns request builder for Subscription collection
func (b *ListRequestBuilder) Subscriptions() *ListSubscriptionsCollectionRequestBuilder {
	bb := &ListSubscriptionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/subscriptions"
	return bb
}

// ListSubscriptionsCollectionRequestBuilder is request builder for Subscription collection
type ListSubscriptionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Subscription collection
func (b *ListSubscriptionsCollectionRequestBuilder) Request() *ListSubscriptionsCollectionRequest {
	return &ListSubscriptionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Subscription item
func (b *ListSubscriptionsCollectionRequestBuilder) ID(id string) *SubscriptionRequestBuilder {
	bb := &SubscriptionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListSubscriptionsCollectionRequest is request for Subscription collection
type ListSubscriptionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Subscription, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Subscription, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Subscription
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Subscription
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

// Get performs GET request for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Get() ([]Subscription, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Add(reqObj *Subscription) (*Subscription, error) {
	return r.Do("POST", "", reqObj)
}
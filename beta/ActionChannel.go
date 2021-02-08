// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// ChannelCompleteMigrationRequestParameter undocumented
type ChannelCompleteMigrationRequestParameter struct {
}

// FilesFolder is navigation property
func (b *ChannelRequestBuilder) FilesFolder() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/filesFolder"
	return bb
}

// Members returns request builder for ConversationMember collection
func (b *ChannelRequestBuilder) Members() *ChannelMembersCollectionRequestBuilder {
	bb := &ChannelMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// ChannelMembersCollectionRequestBuilder is request builder for ConversationMember collection
type ChannelMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConversationMember collection
func (b *ChannelMembersCollectionRequestBuilder) Request() *ChannelMembersCollectionRequest {
	return &ChannelMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConversationMember item
func (b *ChannelMembersCollectionRequestBuilder) ID(id string) *ConversationMemberRequestBuilder {
	bb := &ConversationMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChannelMembersCollectionRequest is request for ConversationMember collection
type ChannelMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConversationMember collection
func (r *ChannelMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConversationMember, error) {
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
	var values []ConversationMember
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
			value  []ConversationMember
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

// GetN performs GET request for ConversationMember collection, max N pages
func (r *ChannelMembersCollectionRequest) GetN(ctx context.Context, n int) ([]ConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ConversationMember collection
func (r *ChannelMembersCollectionRequest) Get(ctx context.Context) ([]ConversationMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ConversationMember collection
func (r *ChannelMembersCollectionRequest) Add(ctx context.Context, reqObj *ConversationMember) (resObj *ConversationMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Messages returns request builder for ChatMessage collection
func (b *ChannelRequestBuilder) Messages() *ChannelMessagesCollectionRequestBuilder {
	bb := &ChannelMessagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/messages"
	return bb
}

// ChannelMessagesCollectionRequestBuilder is request builder for ChatMessage collection
type ChannelMessagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ChatMessage collection
func (b *ChannelMessagesCollectionRequestBuilder) Request() *ChannelMessagesCollectionRequest {
	return &ChannelMessagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ChatMessage item
func (b *ChannelMessagesCollectionRequestBuilder) ID(id string) *ChatMessageRequestBuilder {
	bb := &ChatMessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChannelMessagesCollectionRequest is request for ChatMessage collection
type ChannelMessagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ChatMessage collection
func (r *ChannelMessagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ChatMessage, error) {
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
	var values []ChatMessage
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
			value  []ChatMessage
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

// GetN performs GET request for ChatMessage collection, max N pages
func (r *ChannelMessagesCollectionRequest) GetN(ctx context.Context, n int) ([]ChatMessage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ChatMessage collection
func (r *ChannelMessagesCollectionRequest) Get(ctx context.Context) ([]ChatMessage, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ChatMessage collection
func (r *ChannelMessagesCollectionRequest) Add(ctx context.Context, reqObj *ChatMessage) (resObj *ChatMessage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Tabs returns request builder for TeamsTab collection
func (b *ChannelRequestBuilder) Tabs() *ChannelTabsCollectionRequestBuilder {
	bb := &ChannelTabsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tabs"
	return bb
}

// ChannelTabsCollectionRequestBuilder is request builder for TeamsTab collection
type ChannelTabsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsTab collection
func (b *ChannelTabsCollectionRequestBuilder) Request() *ChannelTabsCollectionRequest {
	return &ChannelTabsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsTab item
func (b *ChannelTabsCollectionRequestBuilder) ID(id string) *TeamsTabRequestBuilder {
	bb := &TeamsTabRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChannelTabsCollectionRequest is request for TeamsTab collection
type ChannelTabsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsTab, error) {
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
	var values []TeamsTab
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
			value  []TeamsTab
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

// GetN performs GET request for TeamsTab collection, max N pages
func (r *ChannelTabsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsTab, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Get(ctx context.Context) ([]TeamsTab, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsTab collection
func (r *ChannelTabsCollectionRequest) Add(ctx context.Context, reqObj *TeamsTab) (resObj *TeamsTab, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

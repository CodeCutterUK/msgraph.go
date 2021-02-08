// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codecutteruk/msgraph.go/jsonx"
)

// TeamsAppInstallationUpgradeRequestParameter undocumented
type TeamsAppInstallationUpgradeRequestParameter struct {
}

// AppDefinitions returns request builder for TeamsAppDefinition collection
func (b *TeamsAppRequestBuilder) AppDefinitions() *TeamsAppAppDefinitionsCollectionRequestBuilder {
	bb := &TeamsAppAppDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appDefinitions"
	return bb
}

// TeamsAppAppDefinitionsCollectionRequestBuilder is request builder for TeamsAppDefinition collection
type TeamsAppAppDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAppDefinition collection
func (b *TeamsAppAppDefinitionsCollectionRequestBuilder) Request() *TeamsAppAppDefinitionsCollectionRequest {
	return &TeamsAppAppDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAppDefinition item
func (b *TeamsAppAppDefinitionsCollectionRequestBuilder) ID(id string) *TeamsAppDefinitionRequestBuilder {
	bb := &TeamsAppDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamsAppAppDefinitionsCollectionRequest is request for TeamsAppDefinition collection
type TeamsAppAppDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsAppDefinition collection
func (r *TeamsAppAppDefinitionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsAppDefinition, error) {
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
	var values []TeamsAppDefinition
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
			value  []TeamsAppDefinition
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

// GetN performs GET request for TeamsAppDefinition collection, max N pages
func (r *TeamsAppAppDefinitionsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsAppDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsAppDefinition collection
func (r *TeamsAppAppDefinitionsCollectionRequest) Get(ctx context.Context) ([]TeamsAppDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsAppDefinition collection
func (r *TeamsAppAppDefinitionsCollectionRequest) Add(ctx context.Context, reqObj *TeamsAppDefinition) (resObj *TeamsAppDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Bot is navigation property
func (b *TeamsAppDefinitionRequestBuilder) Bot() *TeamworkBotRequestBuilder {
	bb := &TeamworkBotRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/bot"
	return bb
}

// TeamsApp is navigation property
func (b *TeamsAppInstallationRequestBuilder) TeamsApp() *TeamsAppRequestBuilder {
	bb := &TeamsAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsApp"
	return bb
}

// TeamsAppDefinition is navigation property
func (b *TeamsAppInstallationRequestBuilder) TeamsAppDefinition() *TeamsAppDefinitionRequestBuilder {
	bb := &TeamsAppDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsAppDefinition"
	return bb
}

// TeamsApp is navigation property
func (b *TeamsTabRequestBuilder) TeamsApp() *TeamsAppRequestBuilder {
	bb := &TeamsAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsApp"
	return bb
}

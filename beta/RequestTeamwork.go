// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TeamworkRequestBuilder is request builder for Teamwork
type TeamworkRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamworkRequest
func (b *TeamworkRequestBuilder) Request() *TeamworkRequest {
	return &TeamworkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamworkRequest is request for Teamwork
type TeamworkRequest struct{ BaseRequest }

// Get performs GET request for Teamwork
func (r *TeamworkRequest) Get(ctx context.Context) (resObj *Teamwork, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Teamwork
func (r *TeamworkRequest) Update(ctx context.Context, reqObj *Teamwork) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Teamwork
func (r *TeamworkRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamworkBotRequestBuilder is request builder for TeamworkBot
type TeamworkBotRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamworkBotRequest
func (b *TeamworkBotRequestBuilder) Request() *TeamworkBotRequest {
	return &TeamworkBotRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamworkBotRequest is request for TeamworkBot
type TeamworkBotRequest struct{ BaseRequest }

// Get performs GET request for TeamworkBot
func (r *TeamworkBotRequest) Get(ctx context.Context) (resObj *TeamworkBot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamworkBot
func (r *TeamworkBotRequest) Update(ctx context.Context, reqObj *TeamworkBot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamworkBot
func (r *TeamworkBotRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamworkTagRequestBuilder is request builder for TeamworkTag
type TeamworkTagRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamworkTagRequest
func (b *TeamworkTagRequestBuilder) Request() *TeamworkTagRequest {
	return &TeamworkTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamworkTagRequest is request for TeamworkTag
type TeamworkTagRequest struct{ BaseRequest }

// Get performs GET request for TeamworkTag
func (r *TeamworkTagRequest) Get(ctx context.Context) (resObj *TeamworkTag, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamworkTag
func (r *TeamworkTagRequest) Update(ctx context.Context, reqObj *TeamworkTag) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamworkTag
func (r *TeamworkTagRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamworkTagMemberRequestBuilder is request builder for TeamworkTagMember
type TeamworkTagMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamworkTagMemberRequest
func (b *TeamworkTagMemberRequestBuilder) Request() *TeamworkTagMemberRequest {
	return &TeamworkTagMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamworkTagMemberRequest is request for TeamworkTagMember
type TeamworkTagMemberRequest struct{ BaseRequest }

// Get performs GET request for TeamworkTagMember
func (r *TeamworkTagMemberRequest) Get(ctx context.Context) (resObj *TeamworkTagMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamworkTagMember
func (r *TeamworkTagMemberRequest) Update(ctx context.Context, reqObj *TeamworkTagMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamworkTagMember
func (r *TeamworkTagMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TodoRequestBuilder is request builder for Todo
type TodoRequestBuilder struct{ BaseRequestBuilder }

// Request returns TodoRequest
func (b *TodoRequestBuilder) Request() *TodoRequest {
	return &TodoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TodoRequest is request for Todo
type TodoRequest struct{ BaseRequest }

// Get performs GET request for Todo
func (r *TodoRequest) Get(ctx context.Context) (resObj *Todo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Todo
func (r *TodoRequest) Update(ctx context.Context, reqObj *Todo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Todo
func (r *TodoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TodoTaskRequestBuilder is request builder for TodoTask
type TodoTaskRequestBuilder struct{ BaseRequestBuilder }

// Request returns TodoTaskRequest
func (b *TodoTaskRequestBuilder) Request() *TodoTaskRequest {
	return &TodoTaskRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TodoTaskRequest is request for TodoTask
type TodoTaskRequest struct{ BaseRequest }

// Get performs GET request for TodoTask
func (r *TodoTaskRequest) Get(ctx context.Context) (resObj *TodoTask, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TodoTask
func (r *TodoTaskRequest) Update(ctx context.Context, reqObj *TodoTask) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TodoTask
func (r *TodoTaskRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TodoTaskListRequestBuilder is request builder for TodoTaskList
type TodoTaskListRequestBuilder struct{ BaseRequestBuilder }

// Request returns TodoTaskListRequest
func (b *TodoTaskListRequestBuilder) Request() *TodoTaskListRequest {
	return &TodoTaskListRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TodoTaskListRequest is request for TodoTaskList
type TodoTaskListRequest struct{ BaseRequest }

// Get performs GET request for TodoTaskList
func (r *TodoTaskListRequest) Get(ctx context.Context) (resObj *TodoTaskList, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TodoTaskList
func (r *TodoTaskListRequest) Update(ctx context.Context, reqObj *TodoTaskList) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TodoTaskList
func (r *TodoTaskListRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

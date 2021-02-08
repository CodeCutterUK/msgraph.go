// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PrintRequestBuilder is request builder for Print
type PrintRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintRequest
func (b *PrintRequestBuilder) Request() *PrintRequest {
	return &PrintRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintRequest is request for Print
type PrintRequest struct{ BaseRequest }

// Get performs GET request for Print
func (r *PrintRequest) Get(ctx context.Context) (resObj *Print, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Print
func (r *PrintRequest) Update(ctx context.Context, reqObj *Print) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Print
func (r *PrintRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintConnectorRequestBuilder is request builder for PrintConnector
type PrintConnectorRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintConnectorRequest
func (b *PrintConnectorRequestBuilder) Request() *PrintConnectorRequest {
	return &PrintConnectorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintConnectorRequest is request for PrintConnector
type PrintConnectorRequest struct{ BaseRequest }

// Get performs GET request for PrintConnector
func (r *PrintConnectorRequest) Get(ctx context.Context) (resObj *PrintConnector, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintConnector
func (r *PrintConnectorRequest) Update(ctx context.Context, reqObj *PrintConnector) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintConnector
func (r *PrintConnectorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintDocumentRequestBuilder is request builder for PrintDocument
type PrintDocumentRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintDocumentRequest
func (b *PrintDocumentRequestBuilder) Request() *PrintDocumentRequest {
	return &PrintDocumentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintDocumentRequest is request for PrintDocument
type PrintDocumentRequest struct{ BaseRequest }

// Get performs GET request for PrintDocument
func (r *PrintDocumentRequest) Get(ctx context.Context) (resObj *PrintDocument, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintDocument
func (r *PrintDocumentRequest) Update(ctx context.Context, reqObj *PrintDocument) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintDocument
func (r *PrintDocumentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintJobRequestBuilder is request builder for PrintJob
type PrintJobRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintJobRequest
func (b *PrintJobRequestBuilder) Request() *PrintJobRequest {
	return &PrintJobRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintJobRequest is request for PrintJob
type PrintJobRequest struct{ BaseRequest }

// Get performs GET request for PrintJob
func (r *PrintJobRequest) Get(ctx context.Context) (resObj *PrintJob, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintJob
func (r *PrintJobRequest) Update(ctx context.Context, reqObj *PrintJob) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintJob
func (r *PrintJobRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintOperationRequestBuilder is request builder for PrintOperation
type PrintOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintOperationRequest
func (b *PrintOperationRequestBuilder) Request() *PrintOperationRequest {
	return &PrintOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintOperationRequest is request for PrintOperation
type PrintOperationRequest struct{ BaseRequest }

// Get performs GET request for PrintOperation
func (r *PrintOperationRequest) Get(ctx context.Context) (resObj *PrintOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintOperation
func (r *PrintOperationRequest) Update(ctx context.Context, reqObj *PrintOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintOperation
func (r *PrintOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintServiceRequestBuilder is request builder for PrintService
type PrintServiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintServiceRequest
func (b *PrintServiceRequestBuilder) Request() *PrintServiceRequest {
	return &PrintServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintServiceRequest is request for PrintService
type PrintServiceRequest struct{ BaseRequest }

// Get performs GET request for PrintService
func (r *PrintServiceRequest) Get(ctx context.Context) (resObj *PrintService, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintService
func (r *PrintServiceRequest) Update(ctx context.Context, reqObj *PrintService) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintService
func (r *PrintServiceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintServiceEndpointRequestBuilder is request builder for PrintServiceEndpoint
type PrintServiceEndpointRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintServiceEndpointRequest
func (b *PrintServiceEndpointRequestBuilder) Request() *PrintServiceEndpointRequest {
	return &PrintServiceEndpointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintServiceEndpointRequest is request for PrintServiceEndpoint
type PrintServiceEndpointRequest struct{ BaseRequest }

// Get performs GET request for PrintServiceEndpoint
func (r *PrintServiceEndpointRequest) Get(ctx context.Context) (resObj *PrintServiceEndpoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintServiceEndpoint
func (r *PrintServiceEndpointRequest) Update(ctx context.Context, reqObj *PrintServiceEndpoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintServiceEndpoint
func (r *PrintServiceEndpointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintTaskRequestBuilder is request builder for PrintTask
type PrintTaskRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintTaskRequest
func (b *PrintTaskRequestBuilder) Request() *PrintTaskRequest {
	return &PrintTaskRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintTaskRequest is request for PrintTask
type PrintTaskRequest struct{ BaseRequest }

// Get performs GET request for PrintTask
func (r *PrintTaskRequest) Get(ctx context.Context) (resObj *PrintTask, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintTask
func (r *PrintTaskRequest) Update(ctx context.Context, reqObj *PrintTask) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintTask
func (r *PrintTaskRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintTaskDefinitionRequestBuilder is request builder for PrintTaskDefinition
type PrintTaskDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintTaskDefinitionRequest
func (b *PrintTaskDefinitionRequestBuilder) Request() *PrintTaskDefinitionRequest {
	return &PrintTaskDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintTaskDefinitionRequest is request for PrintTaskDefinition
type PrintTaskDefinitionRequest struct{ BaseRequest }

// Get performs GET request for PrintTaskDefinition
func (r *PrintTaskDefinitionRequest) Get(ctx context.Context) (resObj *PrintTaskDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintTaskDefinition
func (r *PrintTaskDefinitionRequest) Update(ctx context.Context, reqObj *PrintTaskDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintTaskDefinition
func (r *PrintTaskDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintTaskTriggerRequestBuilder is request builder for PrintTaskTrigger
type PrintTaskTriggerRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintTaskTriggerRequest
func (b *PrintTaskTriggerRequestBuilder) Request() *PrintTaskTriggerRequest {
	return &PrintTaskTriggerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintTaskTriggerRequest is request for PrintTaskTrigger
type PrintTaskTriggerRequest struct{ BaseRequest }

// Get performs GET request for PrintTaskTrigger
func (r *PrintTaskTriggerRequest) Get(ctx context.Context) (resObj *PrintTaskTrigger, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintTaskTrigger
func (r *PrintTaskTriggerRequest) Update(ctx context.Context, reqObj *PrintTaskTrigger) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintTaskTrigger
func (r *PrintTaskTriggerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintUsageSummaryByPrinterRequestBuilder is request builder for PrintUsageSummaryByPrinter
type PrintUsageSummaryByPrinterRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintUsageSummaryByPrinterRequest
func (b *PrintUsageSummaryByPrinterRequestBuilder) Request() *PrintUsageSummaryByPrinterRequest {
	return &PrintUsageSummaryByPrinterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintUsageSummaryByPrinterRequest is request for PrintUsageSummaryByPrinter
type PrintUsageSummaryByPrinterRequest struct{ BaseRequest }

// Get performs GET request for PrintUsageSummaryByPrinter
func (r *PrintUsageSummaryByPrinterRequest) Get(ctx context.Context) (resObj *PrintUsageSummaryByPrinter, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintUsageSummaryByPrinter
func (r *PrintUsageSummaryByPrinterRequest) Update(ctx context.Context, reqObj *PrintUsageSummaryByPrinter) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintUsageSummaryByPrinter
func (r *PrintUsageSummaryByPrinterRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintUsageSummaryByUserRequestBuilder is request builder for PrintUsageSummaryByUser
type PrintUsageSummaryByUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintUsageSummaryByUserRequest
func (b *PrintUsageSummaryByUserRequestBuilder) Request() *PrintUsageSummaryByUserRequest {
	return &PrintUsageSummaryByUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintUsageSummaryByUserRequest is request for PrintUsageSummaryByUser
type PrintUsageSummaryByUserRequest struct{ BaseRequest }

// Get performs GET request for PrintUsageSummaryByUser
func (r *PrintUsageSummaryByUserRequest) Get(ctx context.Context) (resObj *PrintUsageSummaryByUser, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintUsageSummaryByUser
func (r *PrintUsageSummaryByUserRequest) Update(ctx context.Context, reqObj *PrintUsageSummaryByUser) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintUsageSummaryByUser
func (r *PrintUsageSummaryByUserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type PrintDocumentCreateUploadSessionRequestBuilder struct{ BaseRequestBuilder }

// CreateUploadSession action undocumented
func (b *PrintDocumentRequestBuilder) CreateUploadSession(reqObj *PrintDocumentCreateUploadSessionRequestParameter) *PrintDocumentCreateUploadSessionRequestBuilder {
	bb := &PrintDocumentCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrintDocumentCreateUploadSessionRequest struct{ BaseRequest }

//
func (b *PrintDocumentCreateUploadSessionRequestBuilder) Request() *PrintDocumentCreateUploadSessionRequest {
	return &PrintDocumentCreateUploadSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrintDocumentCreateUploadSessionRequest) Post(ctx context.Context) (resObj *UploadSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type PrintDocumentUploadDataRequestBuilder struct{ BaseRequestBuilder }

// UploadData action undocumented
func (b *PrintDocumentRequestBuilder) UploadData(reqObj *PrintDocumentUploadDataRequestParameter) *PrintDocumentUploadDataRequestBuilder {
	bb := &PrintDocumentUploadDataRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/uploadData"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrintDocumentUploadDataRequest struct{ BaseRequest }

//
func (b *PrintDocumentUploadDataRequestBuilder) Request() *PrintDocumentUploadDataRequest {
	return &PrintDocumentUploadDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrintDocumentUploadDataRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type PrintJobCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumented
func (b *PrintJobRequestBuilder) Cancel(reqObj *PrintJobCancelRequestParameter) *PrintJobCancelRequestBuilder {
	bb := &PrintJobCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrintJobCancelRequest struct{ BaseRequest }

//
func (b *PrintJobCancelRequestBuilder) Request() *PrintJobCancelRequest {
	return &PrintJobCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrintJobCancelRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type PrintJobStartRequestBuilder struct{ BaseRequestBuilder }

// Start action undocumented
func (b *PrintJobRequestBuilder) Start(reqObj *PrintJobStartRequestParameter) *PrintJobStartRequestBuilder {
	bb := &PrintJobStartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/start"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrintJobStartRequest struct{ BaseRequest }

//
func (b *PrintJobStartRequestBuilder) Request() *PrintJobStartRequest {
	return &PrintJobStartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrintJobStartRequest) Post(ctx context.Context) (resObj *PrintJobStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type PrintJobAbortRequestBuilder struct{ BaseRequestBuilder }

// Abort action undocumented
func (b *PrintJobRequestBuilder) Abort(reqObj *PrintJobAbortRequestParameter) *PrintJobAbortRequestBuilder {
	bb := &PrintJobAbortRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/abort"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrintJobAbortRequest struct{ BaseRequest }

//
func (b *PrintJobAbortRequestBuilder) Request() *PrintJobAbortRequest {
	return &PrintJobAbortRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrintJobAbortRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type PrintJobCancelPrintJobRequestBuilder struct{ BaseRequestBuilder }

// CancelPrintJob action undocumented
func (b *PrintJobRequestBuilder) CancelPrintJob(reqObj *PrintJobCancelPrintJobRequestParameter) *PrintJobCancelPrintJobRequestBuilder {
	bb := &PrintJobCancelPrintJobRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cancelPrintJob"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrintJobCancelPrintJobRequest struct{ BaseRequest }

//
func (b *PrintJobCancelPrintJobRequestBuilder) Request() *PrintJobCancelPrintJobRequest {
	return &PrintJobCancelPrintJobRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrintJobCancelPrintJobRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type PrintJobRedirectRequestBuilder struct{ BaseRequestBuilder }

// Redirect action undocumented
func (b *PrintJobRequestBuilder) Redirect(reqObj *PrintJobRedirectRequestParameter) *PrintJobRedirectRequestBuilder {
	bb := &PrintJobRedirectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/redirect"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrintJobRedirectRequest struct{ BaseRequest }

//
func (b *PrintJobRedirectRequestBuilder) Request() *PrintJobRedirectRequest {
	return &PrintJobRedirectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrintJobRedirectRequest) Post(ctx context.Context) (resObj *PrintJob, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type PrintJobStartPrintJobRequestBuilder struct{ BaseRequestBuilder }

// StartPrintJob action undocumented
func (b *PrintJobRequestBuilder) StartPrintJob(reqObj *PrintJobStartPrintJobRequestParameter) *PrintJobStartPrintJobRequestBuilder {
	bb := &PrintJobStartPrintJobRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/startPrintJob"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type PrintJobStartPrintJobRequest struct{ BaseRequest }

//
func (b *PrintJobStartPrintJobRequestBuilder) Request() *PrintJobStartPrintJobRequest {
	return &PrintJobStartPrintJobRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *PrintJobStartPrintJobRequest) Post(ctx context.Context) (resObj *PrintJobStatus, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Package notifications provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package notifications

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	externalRef0 "github.com/openclarity/apiclarity/api3/common"
	externalRef1 "github.com/openclarity/apiclarity/api3/global"
)

// APIClarityNotification defines model for APIClarityNotification.
type APIClarityNotification struct {
	union json.RawMessage
}

// APIFindings A group of findings
type APIFindings struct {
	// Items A list of findings
	Items *[]externalRef0.APIFinding `json:"items,omitempty"`
}

// ApiFindingsNotification defines model for ApiFindingsNotification.
type ApiFindingsNotification struct {
	// Items A list of findings
	Items            *[]externalRef0.APIFinding `json:"items,omitempty"`
	NotificationType string                     `json:"notificationType"`
}

// ApiInfo defines model for ApiInfo.
type ApiInfo struct {
	DestinationNamespace *string `json:"destinationNamespace,omitempty"`
	HasProvidedSpec      *bool   `json:"hasProvidedSpec,omitempty"`
	HasReconstructedSpec *bool   `json:"hasReconstructedSpec,omitempty"`
	Id                   *uint32 `json:"id,omitempty"`

	// Name API name
	Name *string `json:"name,omitempty"`
	Port *int    `json:"port,omitempty"`

	// TraceSourceId Trace Source ID which created this API. Null UUID 0 means it has been created by APIClarity (from the UI for example)
	TraceSourceId *openapi_types.UUID `json:"traceSourceId,omitempty"`
}

// AuthorizationModel defines model for AuthorizationModel.
type AuthorizationModel struct {
	Learning   bool                                       `json:"learning"`
	Operations []externalRef1.AuthorizationModelOperation `json:"operations"`
	SpecType   externalRef1.SpecType                      `json:"specType"`
}

// AuthorizationModelNotification defines model for AuthorizationModelNotification.
type AuthorizationModelNotification struct {
	Learning         bool                                       `json:"learning"`
	NotificationType string                                     `json:"notificationType"`
	Operations       []externalRef1.AuthorizationModelOperation `json:"operations"`
	SpecType         externalRef1.SpecType                      `json:"specType"`
}

// BaseNotification Base Notification all APIClarity notifications must extend
type BaseNotification struct {
	NotificationType string `json:"notificationType"`
}

// NewDiscoveredAPINotification defines model for NewDiscoveredAPINotification.
type NewDiscoveredAPINotification struct {
	DestinationNamespace *string `json:"destinationNamespace,omitempty"`
	HasProvidedSpec      *bool   `json:"hasProvidedSpec,omitempty"`
	HasReconstructedSpec *bool   `json:"hasReconstructedSpec,omitempty"`
	Id                   *uint32 `json:"id,omitempty"`

	// Name API name
	Name             *string `json:"name,omitempty"`
	NotificationType string  `json:"notificationType"`
	Port             *int    `json:"port,omitempty"`

	// TraceSourceId Trace Source ID which created this API. Null UUID 0 means it has been created by APIClarity (from the UI for example)
	TraceSourceId *openapi_types.UUID `json:"traceSourceId,omitempty"`
}

// ShortTestProgress Describes the progress of an ongoing test
type ShortTestProgress struct {
	ApiID *externalRef0.ApiID `json:"apiID,omitempty"`

	// Progress Progress of the test
	Progress int `json:"progress"`

	// Starttime Timestamp of the start of the test
	Starttime int64 `json:"starttime"`
}

// ShortTestReport Short Test Report
type ShortTestReport struct {
	ApiID *externalRef0.ApiID `json:"apiID,omitempty"`

	// HighestSeverity Severity of a finding
	HighestSeverity *externalRef0.Severity `json:"highestSeverity,omitempty"`

	// Starttime Timestamp of the start of the test
	Starttime int64 `json:"starttime"`

	// Status An enumeration.
	Status externalRef1.FuzzingStatusEnum `json:"status"`

	// StatusMessage Message for status details, if any
	StatusMessage *string                          `json:"statusMessage,omitempty"`
	Tags          *[]externalRef1.FuzzingReportTag `json:"tags,omitempty"`
}

// SpecDiffs defines model for SpecDiffs.
type SpecDiffs struct {
	Diffs externalRef1.APIDiffs `json:"diffs"`
}

// SpecDiffsNotification defines model for SpecDiffsNotification.
type SpecDiffsNotification struct {
	Diffs            externalRef1.APIDiffs `json:"diffs"`
	NotificationType string                `json:"notificationType"`
}

// TestProgressNotification defines model for TestProgressNotification.
type TestProgressNotification struct {
	ApiID            *externalRef0.ApiID `json:"apiID,omitempty"`
	NotificationType string              `json:"notificationType"`

	// Progress Progress of the test
	Progress int `json:"progress"`

	// Starttime Timestamp of the start of the test
	Starttime int64 `json:"starttime"`
}

// TestReportNotification defines model for TestReportNotification.
type TestReportNotification struct {
	ApiID *externalRef0.ApiID `json:"apiID,omitempty"`

	// HighestSeverity Severity of a finding
	HighestSeverity  *externalRef0.Severity `json:"highestSeverity,omitempty"`
	NotificationType string                 `json:"notificationType"`

	// Starttime Timestamp of the start of the test
	Starttime int64 `json:"starttime"`

	// Status An enumeration.
	Status externalRef1.FuzzingStatusEnum `json:"status"`

	// StatusMessage Message for status details, if any
	StatusMessage *string                          `json:"statusMessage,omitempty"`
	Tags          *[]externalRef1.FuzzingReportTag `json:"tags,omitempty"`
}

// PostNotificationApiIDJSONRequestBody defines body for PostNotificationApiID for application/json ContentType.
type PostNotificationApiIDJSONRequestBody = APIClarityNotification

// AsApiFindingsNotification returns the union data inside the APIClarityNotification as a ApiFindingsNotification
func (t APIClarityNotification) AsApiFindingsNotification() (ApiFindingsNotification, error) {
	var body ApiFindingsNotification
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromApiFindingsNotification overwrites any union data inside the APIClarityNotification as the provided ApiFindingsNotification
func (t *APIClarityNotification) FromApiFindingsNotification(v ApiFindingsNotification) error {
	v.NotificationType = "ApiFindingsNotification"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeApiFindingsNotification performs a merge with any union data inside the APIClarityNotification, using the provided ApiFindingsNotification
func (t *APIClarityNotification) MergeApiFindingsNotification(v ApiFindingsNotification) error {
	v.NotificationType = "ApiFindingsNotification"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsAuthorizationModelNotification returns the union data inside the APIClarityNotification as a AuthorizationModelNotification
func (t APIClarityNotification) AsAuthorizationModelNotification() (AuthorizationModelNotification, error) {
	var body AuthorizationModelNotification
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthorizationModelNotification overwrites any union data inside the APIClarityNotification as the provided AuthorizationModelNotification
func (t *APIClarityNotification) FromAuthorizationModelNotification(v AuthorizationModelNotification) error {
	v.NotificationType = "AuthorizationModelNotification"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAuthorizationModelNotification performs a merge with any union data inside the APIClarityNotification, using the provided AuthorizationModelNotification
func (t *APIClarityNotification) MergeAuthorizationModelNotification(v AuthorizationModelNotification) error {
	v.NotificationType = "AuthorizationModelNotification"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsNewDiscoveredAPINotification returns the union data inside the APIClarityNotification as a NewDiscoveredAPINotification
func (t APIClarityNotification) AsNewDiscoveredAPINotification() (NewDiscoveredAPINotification, error) {
	var body NewDiscoveredAPINotification
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromNewDiscoveredAPINotification overwrites any union data inside the APIClarityNotification as the provided NewDiscoveredAPINotification
func (t *APIClarityNotification) FromNewDiscoveredAPINotification(v NewDiscoveredAPINotification) error {
	v.NotificationType = "NewDiscoveredAPINotification"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeNewDiscoveredAPINotification performs a merge with any union data inside the APIClarityNotification, using the provided NewDiscoveredAPINotification
func (t *APIClarityNotification) MergeNewDiscoveredAPINotification(v NewDiscoveredAPINotification) error {
	v.NotificationType = "NewDiscoveredAPINotification"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsSpecDiffsNotification returns the union data inside the APIClarityNotification as a SpecDiffsNotification
func (t APIClarityNotification) AsSpecDiffsNotification() (SpecDiffsNotification, error) {
	var body SpecDiffsNotification
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSpecDiffsNotification overwrites any union data inside the APIClarityNotification as the provided SpecDiffsNotification
func (t *APIClarityNotification) FromSpecDiffsNotification(v SpecDiffsNotification) error {
	v.NotificationType = "SpecDiffsNotification"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSpecDiffsNotification performs a merge with any union data inside the APIClarityNotification, using the provided SpecDiffsNotification
func (t *APIClarityNotification) MergeSpecDiffsNotification(v SpecDiffsNotification) error {
	v.NotificationType = "SpecDiffsNotification"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsTestProgressNotification returns the union data inside the APIClarityNotification as a TestProgressNotification
func (t APIClarityNotification) AsTestProgressNotification() (TestProgressNotification, error) {
	var body TestProgressNotification
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTestProgressNotification overwrites any union data inside the APIClarityNotification as the provided TestProgressNotification
func (t *APIClarityNotification) FromTestProgressNotification(v TestProgressNotification) error {
	v.NotificationType = "TestProgressNotification"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTestProgressNotification performs a merge with any union data inside the APIClarityNotification, using the provided TestProgressNotification
func (t *APIClarityNotification) MergeTestProgressNotification(v TestProgressNotification) error {
	v.NotificationType = "TestProgressNotification"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

// AsTestReportNotification returns the union data inside the APIClarityNotification as a TestReportNotification
func (t APIClarityNotification) AsTestReportNotification() (TestReportNotification, error) {
	var body TestReportNotification
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTestReportNotification overwrites any union data inside the APIClarityNotification as the provided TestReportNotification
func (t *APIClarityNotification) FromTestReportNotification(v TestReportNotification) error {
	v.NotificationType = "TestReportNotification"
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTestReportNotification performs a merge with any union data inside the APIClarityNotification, using the provided TestReportNotification
func (t *APIClarityNotification) MergeTestReportNotification(v TestReportNotification) error {
	v.NotificationType = "TestReportNotification"
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(b, t.union)
	t.union = merged
	return err
}

func (t APIClarityNotification) Discriminator() (string, error) {
	var discriminator struct {
		Discriminator string `json:"notificationType"`
	}
	err := json.Unmarshal(t.union, &discriminator)
	return discriminator.Discriminator, err
}

func (t APIClarityNotification) ValueByDiscriminator() (interface{}, error) {
	discriminator, err := t.Discriminator()
	if err != nil {
		return nil, err
	}
	switch discriminator {
	case "ApiFindingsNotification":
		return t.AsApiFindingsNotification()
	case "AuthorizationModelNotification":
		return t.AsAuthorizationModelNotification()
	case "NewDiscoveredAPINotification":
		return t.AsNewDiscoveredAPINotification()
	case "SpecDiffsNotification":
		return t.AsSpecDiffsNotification()
	case "TestProgressNotification":
		return t.AsTestProgressNotification()
	case "TestReportNotification":
		return t.AsTestReportNotification()
	default:
		return nil, errors.New("unknown discriminator value: " + discriminator)
	}
}

func (t APIClarityNotification) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *APIClarityNotification) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// PostNotificationApiID request with any body
	PostNotificationApiIDWithBody(ctx context.Context, apiID int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostNotificationApiID(ctx context.Context, apiID int64, body PostNotificationApiIDJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) PostNotificationApiIDWithBody(ctx context.Context, apiID int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostNotificationApiIDRequestWithBody(c.Server, apiID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostNotificationApiID(ctx context.Context, apiID int64, body PostNotificationApiIDJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostNotificationApiIDRequest(c.Server, apiID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewPostNotificationApiIDRequest calls the generic PostNotificationApiID builder with application/json body
func NewPostNotificationApiIDRequest(server string, apiID int64, body PostNotificationApiIDJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostNotificationApiIDRequestWithBody(server, apiID, "application/json", bodyReader)
}

// NewPostNotificationApiIDRequestWithBody generates requests for PostNotificationApiID with any type of body
func NewPostNotificationApiIDRequestWithBody(server string, apiID int64, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "apiID", runtime.ParamLocationPath, apiID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/notification/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// PostNotificationApiID request with any body
	PostNotificationApiIDWithBodyWithResponse(ctx context.Context, apiID int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostNotificationApiIDResponse, error)

	PostNotificationApiIDWithResponse(ctx context.Context, apiID int64, body PostNotificationApiIDJSONRequestBody, reqEditors ...RequestEditorFn) (*PostNotificationApiIDResponse, error)
}

type PostNotificationApiIDResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *string
}

// Status returns HTTPResponse.Status
func (r PostNotificationApiIDResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostNotificationApiIDResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// PostNotificationApiIDWithBodyWithResponse request with arbitrary body returning *PostNotificationApiIDResponse
func (c *ClientWithResponses) PostNotificationApiIDWithBodyWithResponse(ctx context.Context, apiID int64, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostNotificationApiIDResponse, error) {
	rsp, err := c.PostNotificationApiIDWithBody(ctx, apiID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostNotificationApiIDResponse(rsp)
}

func (c *ClientWithResponses) PostNotificationApiIDWithResponse(ctx context.Context, apiID int64, body PostNotificationApiIDJSONRequestBody, reqEditors ...RequestEditorFn) (*PostNotificationApiIDResponse, error) {
	rsp, err := c.PostNotificationApiID(ctx, apiID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostNotificationApiIDResponse(rsp)
}

// ParsePostNotificationApiIDResponse parses an HTTP response from a PostNotificationApiIDWithResponse call
func ParsePostNotificationApiIDResponse(rsp *http.Response) (*PostNotificationApiIDResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostNotificationApiIDResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest string
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Provide to Apiclarity list of raw input for a given API ID associated with a given timestamp
	// (POST /notification/{apiID})
	PostNotificationApiID(w http.ResponseWriter, r *http.Request, apiID int64)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostNotificationApiID operation middleware
func (siw *ServerInterfaceWrapper) PostNotificationApiID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "apiID" -------------
	var apiID int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "apiID", runtime.ParamLocationPath, chi.URLParam(r, "apiID"), &apiID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "apiID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNotificationApiID(w, r, apiID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/notification/{apiID}", wrapper.PostNotificationApiID)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AAD//830R9R5HwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../common/openapi.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(pathPrefix, "../global/openapi.gen.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for JobStatus.
const (
	ERRORSUBMITFAILED JobStatus = "ERROR_SUBMIT_FAILED"
	QUEUED            JobStatus = "QUEUED"
	SUBMITTED         JobStatus = "SUBMITTED"
)

// Defines values for JobType.
const (
	BALANCEUPDATE  JobType = "BALANCE_UPDATE"
	UPDATEUSERINFO JobType = "UPDATE_USER_INFO"
)

// BalanceUpdateParams defines model for BalanceUpdateParams.
type BalanceUpdateParams struct {
	// Amount The amount to add to the user's balance
	Amount *float32 `json:"amount,omitempty"`

	// UserId The user ID to update
	UserId string `json:"userId"`
}

// BatchFrequency defines model for BatchFrequency.
type BatchFrequency struct {
	// Frequency Times of batch processing per second
	Frequency int `json:"frequency"`
}

// BatchSize defines model for BatchSize.
type BatchSize struct {
	// BatchSize Number of jobs which pass to BatchProcessor
	BatchSize int `json:"batch-size"`
}

// Job defines model for Job.
type Job struct {
	// Id The unique identifier for the job
	Id openapi_types.UUID `json:"id"`

	// Name The name of the job, optional
	Name   string     `json:"name"`
	Params Job_Params `json:"params"`

	// Status The status of the job
	Status JobStatus `json:"status"`
	Type   JobType   `json:"type"`
}

// Job_Params defines model for Job.Params.
type Job_Params struct {
	union json.RawMessage
}

// JobRequest defines model for JobRequest.
type JobRequest struct {
	// Name The name of the job, optional
	Name   string            `json:"name"`
	Params JobRequest_Params `json:"params"`
	Type   JobType           `json:"type"`
}

// JobRequest_Params defines model for JobRequest.Params.
type JobRequest_Params struct {
	union json.RawMessage
}

// JobStatus The status of the job
type JobStatus string

// JobType defines model for JobType.
type JobType string

// UpdateUserInfoParams defines model for UpdateUserInfoParams.
type UpdateUserInfoParams struct {
	// Email The new email address for the user
	Email *string `json:"email,omitempty"`

	// Name The user name to update
	Name *string `json:"name,omitempty"`

	// UserId The user ID to update
	UserId string `json:"userId"`
}

// PostBatchFrequencyJSONRequestBody defines body for PostBatchFrequency for application/json ContentType.
type PostBatchFrequencyJSONRequestBody = BatchFrequency

// UpdateBatchSizeJSONRequestBody defines body for UpdateBatchSize for application/json ContentType.
type UpdateBatchSizeJSONRequestBody = BatchSize

// PostJobJSONRequestBody defines body for PostJob for application/json ContentType.
type PostJobJSONRequestBody = JobRequest

// AsUpdateUserInfoParams returns the union data inside the Job_Params as a UpdateUserInfoParams
func (t Job_Params) AsUpdateUserInfoParams() (UpdateUserInfoParams, error) {
	var body UpdateUserInfoParams
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUpdateUserInfoParams overwrites any union data inside the Job_Params as the provided UpdateUserInfoParams
func (t *Job_Params) FromUpdateUserInfoParams(v UpdateUserInfoParams) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUpdateUserInfoParams performs a merge with any union data inside the Job_Params, using the provided UpdateUserInfoParams
func (t *Job_Params) MergeUpdateUserInfoParams(v UpdateUserInfoParams) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBalanceUpdateParams returns the union data inside the Job_Params as a BalanceUpdateParams
func (t Job_Params) AsBalanceUpdateParams() (BalanceUpdateParams, error) {
	var body BalanceUpdateParams
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBalanceUpdateParams overwrites any union data inside the Job_Params as the provided BalanceUpdateParams
func (t *Job_Params) FromBalanceUpdateParams(v BalanceUpdateParams) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBalanceUpdateParams performs a merge with any union data inside the Job_Params, using the provided BalanceUpdateParams
func (t *Job_Params) MergeBalanceUpdateParams(v BalanceUpdateParams) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Job_Params) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Job_Params) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsUpdateUserInfoParams returns the union data inside the JobRequest_Params as a UpdateUserInfoParams
func (t JobRequest_Params) AsUpdateUserInfoParams() (UpdateUserInfoParams, error) {
	var body UpdateUserInfoParams
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUpdateUserInfoParams overwrites any union data inside the JobRequest_Params as the provided UpdateUserInfoParams
func (t *JobRequest_Params) FromUpdateUserInfoParams(v UpdateUserInfoParams) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUpdateUserInfoParams performs a merge with any union data inside the JobRequest_Params, using the provided UpdateUserInfoParams
func (t *JobRequest_Params) MergeUpdateUserInfoParams(v UpdateUserInfoParams) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsBalanceUpdateParams returns the union data inside the JobRequest_Params as a BalanceUpdateParams
func (t JobRequest_Params) AsBalanceUpdateParams() (BalanceUpdateParams, error) {
	var body BalanceUpdateParams
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBalanceUpdateParams overwrites any union data inside the JobRequest_Params as the provided BalanceUpdateParams
func (t *JobRequest_Params) FromBalanceUpdateParams(v BalanceUpdateParams) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBalanceUpdateParams performs a merge with any union data inside the JobRequest_Params, using the provided BalanceUpdateParams
func (t *JobRequest_Params) MergeBalanceUpdateParams(v BalanceUpdateParams) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t JobRequest_Params) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *JobRequest_Params) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /batch-frequency)
	GetBatchFrequency(c *gin.Context)

	// (POST /batch-frequency)
	PostBatchFrequency(c *gin.Context)
	// get the current batch size
	// (GET /batch-size)
	GetBatchSize(c *gin.Context)
	// update the batch size
	// (POST /batch-size)
	UpdateBatchSize(c *gin.Context)

	// (POST /job)
	PostJob(c *gin.Context)

	// (GET /job/{id})
	GetJobById(c *gin.Context, id openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetBatchFrequency operation middleware
func (siw *ServerInterfaceWrapper) GetBatchFrequency(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetBatchFrequency(c)
}

// PostBatchFrequency operation middleware
func (siw *ServerInterfaceWrapper) PostBatchFrequency(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostBatchFrequency(c)
}

// GetBatchSize operation middleware
func (siw *ServerInterfaceWrapper) GetBatchSize(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetBatchSize(c)
}

// UpdateBatchSize operation middleware
func (siw *ServerInterfaceWrapper) UpdateBatchSize(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateBatchSize(c)
}

// PostJob operation middleware
func (siw *ServerInterfaceWrapper) PostJob(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostJob(c)
}

// GetJobById operation middleware
func (siw *ServerInterfaceWrapper) GetJobById(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameterWithOptions("simple", "id", c.Param("id"), &id, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetJobById(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/batch-frequency", wrapper.GetBatchFrequency)
	router.POST(options.BaseURL+"/batch-frequency", wrapper.PostBatchFrequency)
	router.GET(options.BaseURL+"/batch-size", wrapper.GetBatchSize)
	router.POST(options.BaseURL+"/batch-size", wrapper.UpdateBatchSize)
	router.POST(options.BaseURL+"/job", wrapper.PostJob)
	router.GET(options.BaseURL+"/job/:id", wrapper.GetJobById)
}

type GetBatchFrequencyRequestObject struct {
}

type GetBatchFrequencyResponseObject interface {
	VisitGetBatchFrequencyResponse(w http.ResponseWriter) error
}

type GetBatchFrequency200JSONResponse BatchFrequency

func (response GetBatchFrequency200JSONResponse) VisitGetBatchFrequencyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostBatchFrequencyRequestObject struct {
	Body *PostBatchFrequencyJSONRequestBody
}

type PostBatchFrequencyResponseObject interface {
	VisitPostBatchFrequencyResponse(w http.ResponseWriter) error
}

type PostBatchFrequency200JSONResponse BatchFrequency

func (response PostBatchFrequency200JSONResponse) VisitPostBatchFrequencyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetBatchSizeRequestObject struct {
}

type GetBatchSizeResponseObject interface {
	VisitGetBatchSizeResponse(w http.ResponseWriter) error
}

type GetBatchSize200JSONResponse struct {
	Size *int `json:"size,omitempty"`
}

func (response GetBatchSize200JSONResponse) VisitGetBatchSizeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type UpdateBatchSizeRequestObject struct {
	Body *UpdateBatchSizeJSONRequestBody
}

type UpdateBatchSizeResponseObject interface {
	VisitUpdateBatchSizeResponse(w http.ResponseWriter) error
}

type UpdateBatchSize200JSONResponse BatchSize

func (response UpdateBatchSize200JSONResponse) VisitUpdateBatchSizeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostJobRequestObject struct {
	Body *PostJobJSONRequestBody
}

type PostJobResponseObject interface {
	VisitPostJobResponse(w http.ResponseWriter) error
}

type PostJob201JSONResponse Job

func (response PostJob201JSONResponse) VisitPostJobResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type GetJobByIdRequestObject struct {
	Id openapi_types.UUID `json:"id"`
}

type GetJobByIdResponseObject interface {
	VisitGetJobByIdResponse(w http.ResponseWriter) error
}

type GetJobById200JSONResponse Job

func (response GetJobById200JSONResponse) VisitGetJobByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /batch-frequency)
	GetBatchFrequency(ctx context.Context, request GetBatchFrequencyRequestObject) (GetBatchFrequencyResponseObject, error)

	// (POST /batch-frequency)
	PostBatchFrequency(ctx context.Context, request PostBatchFrequencyRequestObject) (PostBatchFrequencyResponseObject, error)
	// get the current batch size
	// (GET /batch-size)
	GetBatchSize(ctx context.Context, request GetBatchSizeRequestObject) (GetBatchSizeResponseObject, error)
	// update the batch size
	// (POST /batch-size)
	UpdateBatchSize(ctx context.Context, request UpdateBatchSizeRequestObject) (UpdateBatchSizeResponseObject, error)

	// (POST /job)
	PostJob(ctx context.Context, request PostJobRequestObject) (PostJobResponseObject, error)

	// (GET /job/{id})
	GetJobById(ctx context.Context, request GetJobByIdRequestObject) (GetJobByIdResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetBatchFrequency operation middleware
func (sh *strictHandler) GetBatchFrequency(ctx *gin.Context) {
	var request GetBatchFrequencyRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetBatchFrequency(ctx, request.(GetBatchFrequencyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetBatchFrequency")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetBatchFrequencyResponseObject); ok {
		if err := validResponse.VisitGetBatchFrequencyResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostBatchFrequency operation middleware
func (sh *strictHandler) PostBatchFrequency(ctx *gin.Context) {
	var request PostBatchFrequencyRequestObject

	var body PostBatchFrequencyJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostBatchFrequency(ctx, request.(PostBatchFrequencyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostBatchFrequency")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostBatchFrequencyResponseObject); ok {
		if err := validResponse.VisitPostBatchFrequencyResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetBatchSize operation middleware
func (sh *strictHandler) GetBatchSize(ctx *gin.Context) {
	var request GetBatchSizeRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetBatchSize(ctx, request.(GetBatchSizeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetBatchSize")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetBatchSizeResponseObject); ok {
		if err := validResponse.VisitGetBatchSizeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// UpdateBatchSize operation middleware
func (sh *strictHandler) UpdateBatchSize(ctx *gin.Context) {
	var request UpdateBatchSizeRequestObject

	var body UpdateBatchSizeJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.UpdateBatchSize(ctx, request.(UpdateBatchSizeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "UpdateBatchSize")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(UpdateBatchSizeResponseObject); ok {
		if err := validResponse.VisitUpdateBatchSizeResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostJob operation middleware
func (sh *strictHandler) PostJob(ctx *gin.Context) {
	var request PostJobRequestObject

	var body PostJobJSONRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostJob(ctx, request.(PostJobRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostJob")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(PostJobResponseObject); ok {
		if err := validResponse.VisitPostJobResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetJobById operation middleware
func (sh *strictHandler) GetJobById(ctx *gin.Context, id openapi_types.UUID) {
	var request GetJobByIdRequestObject

	request.Id = id

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetJobById(ctx, request.(GetJobByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetJobById")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetJobByIdResponseObject); ok {
		if err := validResponse.VisitGetJobByIdResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xXS4/iRhD+K61KpFy8A5u9+QaBiUC7M4THaTVCbbs8NMLd3u52RgT5v0fVbR4GG7Ka",
	"kEh7wnK3q7766qsHO4hVliuJ0hoId2DiFWbcPfb5hssYF3nCLU645pl7zZNEWKEk30y0ylFbgQbClG8M",
	"BpCfvNoBz1QhLT0laGItcvoOQpivkPkzZhXjSUI/doWsMKh/MSzyniEAu80RQpBFFqGGMgC6MUqaTdIZ",
	"Gw3IWOFAHw0Yq4V8hbIMQOO3QmhMIPy6t/ZyuKeiNcaWHPW5jVePdBllvCWH9djS06MzLCJDw1TKIrLB",
	"cq1iNEbIV5ajZgZjJZMjNCEtvlJwZ9iODlrhzcRfeInMef1gqrM6tCdHJGFbq8iwt5UggNwYIs3ZnHi0",
	"St9GeOKoCeJYRZfgRFvupPhWIBMJSitSgZqlSjtNrFUEAaRKZ9xCCEUhksu8BiB5hs2m6YQirowFTOVe",
	"v01m8oPOlcTnFMKvO/hZYwoh/NQ5lkqnqpOOr44F6UimqqqSMrj+UVNllS9lAMZyWzjv1z4fq2jmL5b7",
	"CG5+MKdr5wl0TDriDq4rgwciWhI7JW0ae5nfHyIN7+G0ovMfsTg7pPuSLp+PE8IgAJRFRj7+WAwXwwEE",
	"MFv0v4zmc/c8nE6fp0v/ZvnYG30eDk78HqndAw93B3uLyaA3Hy4Xs+F0OXp6fIYA+r3Pvaffhkt/1Gio",
	"kfPvGxCYcbFpkQu+MXdM80GjMYeGQE37+zqAGwxOf1dGw382W+iikKm6dPRFxFp9cG2YpsVYRa4ghd3g",
	"xSkE8Cdq4z/8+NB96FIIKkfJcwEhfHroPnxyCrQrx3XHt+va3HpFV8KUEk4YKHr4He3Z8KPATK6k8Vn7",
	"tduln1hJi3688zzfiNiZ6KwNQdpvErfK6MyTI+eS/bjQGqWtBmpav54r0xDFRJmmMFzb6qtke8cIyv+d",
	"rzOe2IobFiFKZoqYpntabDbbSshJBblT3xuuSsMtHu+Mst4J9m4bVo6G8rmlEGeNLpoiy7jeQkjxuO7R",
	"dLFNQ77D1SO+k4BmHnGth1hd4N21dHTcJiPi6KaCTqn2bx3bZ+norKudsLVkx27Q3YPmk62lsUY//pue",
	"mhjt0RRvITLWeFqKaxV1diIprxXiWEX97SjZLxloURu3JwlyRm1/v9yFftGrCys4CebGcl2+3FGELWTN",
	"/dLD3oRdOSWZHGP6Y5Cw0cDRVJZ/BwAA//9kQ1PXvA4AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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

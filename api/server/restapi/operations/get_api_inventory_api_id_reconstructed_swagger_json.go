// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAPIInventoryAPIIDReconstructedSwaggerJSONHandlerFunc turns a function with the right signature into a get API inventory API ID reconstructed swagger JSON handler
type GetAPIInventoryAPIIDReconstructedSwaggerJSONHandlerFunc func(GetAPIInventoryAPIIDReconstructedSwaggerJSONParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAPIInventoryAPIIDReconstructedSwaggerJSONHandlerFunc) Handle(params GetAPIInventoryAPIIDReconstructedSwaggerJSONParams) middleware.Responder {
	return fn(params)
}

// GetAPIInventoryAPIIDReconstructedSwaggerJSONHandler interface for that can handle valid get API inventory API ID reconstructed swagger JSON params
type GetAPIInventoryAPIIDReconstructedSwaggerJSONHandler interface {
	Handle(GetAPIInventoryAPIIDReconstructedSwaggerJSONParams) middleware.Responder
}

// NewGetAPIInventoryAPIIDReconstructedSwaggerJSON creates a new http.Handler for the get API inventory API ID reconstructed swagger JSON operation
func NewGetAPIInventoryAPIIDReconstructedSwaggerJSON(ctx *middleware.Context, handler GetAPIInventoryAPIIDReconstructedSwaggerJSONHandler) *GetAPIInventoryAPIIDReconstructedSwaggerJSON {
	return &GetAPIInventoryAPIIDReconstructedSwaggerJSON{Context: ctx, Handler: handler}
}

/* GetAPIInventoryAPIIDReconstructedSwaggerJSON swagger:route GET /apiInventory/{apiId}/reconstructed_swagger.json getApiInventoryApiIdReconstructedSwaggerJson

Get reconstructed API spec json file

*/
type GetAPIInventoryAPIIDReconstructedSwaggerJSON struct {
	Context *middleware.Context
	Handler GetAPIInventoryAPIIDReconstructedSwaggerJSONHandler
}

func (o *GetAPIInventoryAPIIDReconstructedSwaggerJSON) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetAPIInventoryAPIIDReconstructedSwaggerJSONParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
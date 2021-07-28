// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// NewPutFegNetworkIDSubscriberConfigBaseNamesParams creates a new PutFegNetworkIDSubscriberConfigBaseNamesParams object
// with the default values initialized.
func NewPutFegNetworkIDSubscriberConfigBaseNamesParams() *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	var ()
	return &PutFegNetworkIDSubscriberConfigBaseNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutFegNetworkIDSubscriberConfigBaseNamesParamsWithTimeout creates a new PutFegNetworkIDSubscriberConfigBaseNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutFegNetworkIDSubscriberConfigBaseNamesParamsWithTimeout(timeout time.Duration) *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	var ()
	return &PutFegNetworkIDSubscriberConfigBaseNamesParams{

		timeout: timeout,
	}
}

// NewPutFegNetworkIDSubscriberConfigBaseNamesParamsWithContext creates a new PutFegNetworkIDSubscriberConfigBaseNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutFegNetworkIDSubscriberConfigBaseNamesParamsWithContext(ctx context.Context) *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	var ()
	return &PutFegNetworkIDSubscriberConfigBaseNamesParams{

		Context: ctx,
	}
}

// NewPutFegNetworkIDSubscriberConfigBaseNamesParamsWithHTTPClient creates a new PutFegNetworkIDSubscriberConfigBaseNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutFegNetworkIDSubscriberConfigBaseNamesParamsWithHTTPClient(client *http.Client) *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	var ()
	return &PutFegNetworkIDSubscriberConfigBaseNamesParams{
		HTTPClient: client,
	}
}

/*PutFegNetworkIDSubscriberConfigBaseNamesParams contains all the parameters to send to the API endpoint
for the put feg network ID subscriber config base names operation typically these are written to a http.Request
*/
type PutFegNetworkIDSubscriberConfigBaseNamesParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string
	/*Record
	  Subscriber Config for the Network

	*/
	Record models.BaseNames

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) WithTimeout(timeout time.Duration) *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) WithContext(ctx context.Context) *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) WithHTTPClient(client *http.Client) *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) WithNetworkID(networkID string) *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithRecord adds the record to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) WithRecord(record models.BaseNames) *PutFegNetworkIDSubscriberConfigBaseNamesParams {
	o.SetRecord(record)
	return o
}

// SetRecord adds the record to the put feg network ID subscriber config base names params
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) SetRecord(record models.BaseNames) {
	o.Record = record
}

// WriteToRequest writes these params to a swagger request
func (o *PutFegNetworkIDSubscriberConfigBaseNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if o.Record != nil {
		if err := r.SetBodyParam(o.Record); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
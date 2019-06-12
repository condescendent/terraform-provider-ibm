// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVpcsVpcIDDefaultSecurityGroupParams creates a new GetVpcsVpcIDDefaultSecurityGroupParams object
// with the default values initialized.
func NewGetVpcsVpcIDDefaultSecurityGroupParams() *GetVpcsVpcIDDefaultSecurityGroupParams {
	var ()
	return &GetVpcsVpcIDDefaultSecurityGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVpcsVpcIDDefaultSecurityGroupParamsWithTimeout creates a new GetVpcsVpcIDDefaultSecurityGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVpcsVpcIDDefaultSecurityGroupParamsWithTimeout(timeout time.Duration) *GetVpcsVpcIDDefaultSecurityGroupParams {
	var ()
	return &GetVpcsVpcIDDefaultSecurityGroupParams{

		timeout: timeout,
	}
}

// NewGetVpcsVpcIDDefaultSecurityGroupParamsWithContext creates a new GetVpcsVpcIDDefaultSecurityGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVpcsVpcIDDefaultSecurityGroupParamsWithContext(ctx context.Context) *GetVpcsVpcIDDefaultSecurityGroupParams {
	var ()
	return &GetVpcsVpcIDDefaultSecurityGroupParams{

		Context: ctx,
	}
}

// NewGetVpcsVpcIDDefaultSecurityGroupParamsWithHTTPClient creates a new GetVpcsVpcIDDefaultSecurityGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVpcsVpcIDDefaultSecurityGroupParamsWithHTTPClient(client *http.Client) *GetVpcsVpcIDDefaultSecurityGroupParams {
	var ()
	return &GetVpcsVpcIDDefaultSecurityGroupParams{
		HTTPClient: client,
	}
}

/*GetVpcsVpcIDDefaultSecurityGroupParams contains all the parameters to send to the API endpoint
for the get vpcs vpc ID default security group operation typically these are written to a http.Request
*/
type GetVpcsVpcIDDefaultSecurityGroupParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcID
	  The VPC identifier

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) WithTimeout(timeout time.Duration) *GetVpcsVpcIDDefaultSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) WithContext(ctx context.Context) *GetVpcsVpcIDDefaultSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) WithHTTPClient(client *http.Client) *GetVpcsVpcIDDefaultSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) WithGeneration(generation int64) *GetVpcsVpcIDDefaultSecurityGroupParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithVersion adds the version to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) WithVersion(version string) *GetVpcsVpcIDDefaultSecurityGroupParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) WithVpcID(vpcID string) *GetVpcsVpcIDDefaultSecurityGroupParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the get vpcs vpc ID default security group params
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVpcsVpcIDDefaultSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	// path param vpc_id
	if err := r.SetPathParam("vpc_id", o.VpcID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
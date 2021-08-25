/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/core"
	. "github.com/twilio/twilio-go/rest/serverless/v1"
)

func ResourceServicesAssets() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesAssets,
		ReadContext:   readServicesAssets,
		UpdateContext: updateServicesAssets,
		DeleteContext: deleteServicesAssets,
		Schema: map[string]*schema.Schema{
			"service_sid":   AsString(SchemaRequired),
			"friendly_name": AsString(SchemaRequired),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesAssetsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesAssets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateAssetParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.CreateAsset(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesAssets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ServerlessV1.DeleteAsset(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesAssets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.FetchAsset(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesAssetsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesAssets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateAssetParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.UpdateAsset(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesFunctions() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesFunctions,
		ReadContext:   readServicesFunctions,
		UpdateContext: updateServicesFunctions,
		DeleteContext: deleteServicesFunctions,
		Schema: map[string]*schema.Schema{
			"service_sid":   AsString(SchemaRequired),
			"friendly_name": AsString(SchemaRequired),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesFunctionsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesFunctions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateFunctionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.CreateFunction(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesFunctions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ServerlessV1.DeleteFunction(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesFunctions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.FetchFunction(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesFunctionsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesFunctions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateFunctionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.UpdateFunction(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServices() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServices,
		ReadContext:   readServices,
		UpdateContext: updateServices,
		DeleteContext: deleteServices,
		Schema: map[string]*schema.Schema{
			"friendly_name":       AsString(SchemaRequired),
			"unique_name":         AsString(SchemaRequired),
			"include_credentials": AsBool(SchemaComputedOptional),
			"ui_editable":         AsBool(SchemaComputedOptional),
			"sid":                 AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.ServerlessV1.CreateService(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ServerlessV1.DeleteService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.FetchService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesEnvironmentsVariables() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesEnvironmentsVariables,
		ReadContext:   readServicesEnvironmentsVariables,
		UpdateContext: updateServicesEnvironmentsVariables,
		DeleteContext: deleteServicesEnvironmentsVariables,
		Schema: map[string]*schema.Schema{
			"service_sid":     AsString(SchemaRequired),
			"environment_sid": AsString(SchemaRequired),
			"key":             AsString(SchemaRequired),
			"value":           AsString(SchemaRequired),
			"sid":             AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesEnvironmentsVariablesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesEnvironmentsVariables(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateVariableParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	environmentSid := d.Get("environment_sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.CreateVariable(serviceSid, environmentSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, environmentSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesEnvironmentsVariables(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	environmentSid := d.Get("environment_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ServerlessV1.DeleteVariable(serviceSid, environmentSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesEnvironmentsVariables(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	environmentSid := d.Get("environment_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.FetchVariable(serviceSid, environmentSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesEnvironmentsVariablesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/environment_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("environment_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func updateServicesEnvironmentsVariables(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateVariableParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	environmentSid := d.Get("environment_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ServerlessV1.UpdateVariable(serviceSid, environmentSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

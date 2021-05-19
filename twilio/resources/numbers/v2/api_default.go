/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/twilio/common"
	. "github.com/twilio/twilio-go/rest/numbers/v2"
)

func ResourceRegulatoryComplianceSupportingDocuments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRegulatoryComplianceSupportingDocuments,
		ReadContext:   readRegulatoryComplianceSupportingDocuments,
		UpdateContext: updateRegulatoryComplianceSupportingDocuments,
		DeleteContext: deleteRegulatoryComplianceSupportingDocuments,
		Schema: map[string]*schema.Schema{
			"attributes":    AsString(SchemaOptional),
			"friendly_name": AsString(SchemaOptional),
			"type":          AsString(SchemaOptional),
			"account_sid":   AsString(SchemaComputed),
			"date_created":  AsString(SchemaComputed),
			"date_updated":  AsString(SchemaComputed),
			"mime_type":     AsString(SchemaComputed),
			"sid":           AsString(SchemaComputed),
			"status":        AsString(SchemaComputed),
			"url":           AsString(SchemaComputed),
		},
	}
}

func createRegulatoryComplianceSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSupportingDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.NumbersV2.CreateSupportingDocument(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteRegulatoryComplianceSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NumbersV2.DeleteSupportingDocument(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readRegulatoryComplianceSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.FetchSupportingDocument(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateRegulatoryComplianceSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSupportingDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.UpdateSupportingDocument(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceRegulatoryComplianceEndUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRegulatoryComplianceEndUsers,
		ReadContext:   readRegulatoryComplianceEndUsers,
		UpdateContext: updateRegulatoryComplianceEndUsers,
		DeleteContext: deleteRegulatoryComplianceEndUsers,
		Schema: map[string]*schema.Schema{
			"attributes":    AsString(SchemaOptional),
			"friendly_name": AsString(SchemaOptional),
			"type":          AsString(SchemaOptional),
			"account_sid":   AsString(SchemaComputed),
			"date_created":  AsString(SchemaComputed),
			"date_updated":  AsString(SchemaComputed),
			"sid":           AsString(SchemaComputed),
			"url":           AsString(SchemaComputed),
		},
	}
}

func createRegulatoryComplianceEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateEndUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.NumbersV2.CreateEndUser(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteRegulatoryComplianceEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NumbersV2.DeleteEndUser(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readRegulatoryComplianceEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.FetchEndUser(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateRegulatoryComplianceEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateEndUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.UpdateEndUser(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceRegulatoryComplianceBundles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRegulatoryComplianceBundles,
		ReadContext:   readRegulatoryComplianceBundles,
		UpdateContext: updateRegulatoryComplianceBundles,
		DeleteContext: deleteRegulatoryComplianceBundles,
		Schema: map[string]*schema.Schema{
			"email":           AsString(SchemaOptional),
			"end_user_type":   AsString(SchemaOptional),
			"friendly_name":   AsString(SchemaOptional),
			"iso_country":     AsString(SchemaOptional),
			"number_type":     AsString(SchemaOptional),
			"regulation_sid":  AsString(SchemaOptional),
			"status_callback": AsString(SchemaOptional),
			"account_sid":     AsString(SchemaComputed),
			"date_created":    AsString(SchemaComputed),
			"date_updated":    AsString(SchemaComputed),
			"links":           AsString(SchemaComputed),
			"sid":             AsString(SchemaComputed),
			"status":          AsString(SchemaComputed),
			"url":             AsString(SchemaComputed),
			"valid_until":     AsString(SchemaComputed),
		},
	}
}

func createRegulatoryComplianceBundles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateBundleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.NumbersV2.CreateBundle(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteRegulatoryComplianceBundles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NumbersV2.DeleteBundle(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readRegulatoryComplianceBundles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.FetchBundle(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateRegulatoryComplianceBundles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateBundleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NumbersV2.UpdateBundle(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
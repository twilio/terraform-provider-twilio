/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.25.1
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
	. "github.com/twilio/twilio-go/rest/messaging/v1"
)

func ResourceServicesAlphaSenders() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesAlphaSenders,
		ReadContext:   readServicesAlphaSenders,
		DeleteContext: deleteServicesAlphaSenders,
		Schema: map[string]*schema.Schema{
			"service_sid":  AsString(SchemaForceNewRequired),
			"alpha_sender": AsString(SchemaForceNewRequired),
			"sid":          AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesAlphaSendersImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesAlphaSenders(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateAlphaSenderParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.CreateAlphaSender(serviceSid, &params)
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

func deleteServicesAlphaSenders(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.MessagingV1.DeleteAlphaSender(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesAlphaSenders(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.FetchAlphaSender(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesAlphaSendersImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceServicesPhoneNumbers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesPhoneNumbers,
		ReadContext:   readServicesPhoneNumbers,
		DeleteContext: deleteServicesPhoneNumbers,
		Schema: map[string]*schema.Schema{
			"service_sid":      AsString(SchemaForceNewRequired),
			"phone_number_sid": AsString(SchemaForceNewRequired),
			"sid":              AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesPhoneNumbersImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesPhoneNumbers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreatePhoneNumberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.CreatePhoneNumber(serviceSid, &params)
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

func deleteServicesPhoneNumbers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.MessagingV1.DeletePhoneNumber(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesPhoneNumbers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.FetchPhoneNumber(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesPhoneNumbersImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceServices() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServices,
		ReadContext:   readServices,
		UpdateContext: updateServices,
		DeleteContext: deleteServices,
		Schema: map[string]*schema.Schema{
			"friendly_name":                 AsString(SchemaRequired),
			"area_code_geomatch":            AsBool(SchemaComputedOptional),
			"fallback_method":               AsString(SchemaComputedOptional),
			"fallback_to_long_code":         AsBool(SchemaComputedOptional),
			"fallback_url":                  AsString(SchemaComputedOptional),
			"inbound_method":                AsString(SchemaComputedOptional),
			"inbound_request_url":           AsString(SchemaComputedOptional),
			"mms_converter":                 AsBool(SchemaComputedOptional),
			"scan_message_content":          AsString(SchemaComputedOptional),
			"smart_encoding":                AsBool(SchemaComputedOptional),
			"status_callback":               AsString(SchemaComputedOptional),
			"sticky_sender":                 AsBool(SchemaComputedOptional),
			"synchronous_validation":        AsBool(SchemaComputedOptional),
			"use_inbound_webhook_on_number": AsBool(SchemaComputedOptional),
			"usecase":                       AsString(SchemaComputedOptional),
			"validity_period":               AsInt(SchemaComputedOptional),
			"sid":                           AsString(SchemaComputed),
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

	r, err := m.(*client.Config).Client.MessagingV1.CreateService(&params)
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

	err := m.(*client.Config).Client.MessagingV1.DeleteService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.FetchService(sid)
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

	r, err := m.(*client.Config).Client.MessagingV1.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesShortCodes() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesShortCodes,
		ReadContext:   readServicesShortCodes,
		DeleteContext: deleteServicesShortCodes,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaForceNewRequired),
			"short_code_sid": AsString(SchemaForceNewRequired),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesShortCodesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesShortCodes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateShortCodeParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.CreateShortCode(serviceSid, &params)
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

func deleteServicesShortCodes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.MessagingV1.DeleteShortCode(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesShortCodes(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.FetchShortCode(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesShortCodesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceServicesComplianceUsa2p() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesComplianceUsa2p,
		ReadContext:   readServicesComplianceUsa2p,
		DeleteContext: deleteServicesComplianceUsa2p,
		Schema: map[string]*schema.Schema{
			"messaging_service_sid":    AsString(SchemaForceNewRequired),
			"brand_registration_sid":   AsString(SchemaForceNewRequired),
			"description":              AsString(SchemaForceNewRequired),
			"has_embedded_links":       AsBool(SchemaForceNewRequired),
			"has_embedded_phone":       AsBool(SchemaForceNewRequired),
			"message_samples":          AsList(AsString(SchemaForceNewRequired), SchemaForceNewRequired),
			"us_app_to_person_usecase": AsString(SchemaForceNewRequired),
			"sid":                      AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesComplianceUsa2pImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesComplianceUsa2p(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateUsAppToPersonParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	messagingServiceSid := d.Get("messaging_service_sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.CreateUsAppToPerson(messagingServiceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{messagingServiceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesComplianceUsa2p(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	messagingServiceSid := d.Get("messaging_service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.MessagingV1.DeleteUsAppToPerson(messagingServiceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesComplianceUsa2p(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	messagingServiceSid := d.Get("messaging_service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.MessagingV1.FetchUsAppToPerson(messagingServiceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesComplianceUsa2pImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected messaging_service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("messaging_service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}

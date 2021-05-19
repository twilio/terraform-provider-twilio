/*
 * Twilio - Verify
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
	. "github.com/twilio/twilio-go/rest/verify/v2"
)

func ResourceServicesRateLimits() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesRateLimits,
		ReadContext:   readServicesRateLimits,
		UpdateContext: updateServicesRateLimits,
		DeleteContext: deleteServicesRateLimits,
		Schema: map[string]*schema.Schema{
			"service_sid":  AsString(SchemaRequired),
			"description":  AsString(SchemaOptional),
			"unique_name":  AsString(SchemaOptional),
			"account_sid":  AsString(SchemaComputed),
			"date_created": AsString(SchemaComputed),
			"date_updated": AsString(SchemaComputed),
			"links":        AsString(SchemaComputed),
			"sid":          AsString(SchemaComputed),
			"url":          AsString(SchemaComputed),
		},
	}
}

func createServicesRateLimits(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateRateLimitParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateRateLimit(serviceSid, &params)
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

func deleteServicesRateLimits(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteRateLimit(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesRateLimits(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchRateLimit(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesRateLimits(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateRateLimitParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateRateLimit(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesMessagingConfigurations() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesMessagingConfigurations,
		ReadContext:   readServicesMessagingConfigurations,
		UpdateContext: updateServicesMessagingConfigurations,
		DeleteContext: deleteServicesMessagingConfigurations,
		Schema: map[string]*schema.Schema{
			"service_sid":           AsString(SchemaRequired),
			"country":               AsString(SchemaOptional),
			"messaging_service_sid": AsString(SchemaOptional),
			"account_sid":           AsString(SchemaComputed),
			"date_created":          AsString(SchemaComputed),
			"date_updated":          AsString(SchemaComputed),
			"url":                   AsString(SchemaComputed),
		},
	}
}

func createServicesMessagingConfigurations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateMessagingConfigurationParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateMessagingConfiguration(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.ServiceSid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesMessagingConfigurations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	country := d.Get("country").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteMessagingConfiguration(serviceSid, country)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesMessagingConfigurations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	country := d.Get("country").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchMessagingConfiguration(serviceSid, country)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesMessagingConfigurations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateMessagingConfigurationParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	country := d.Get("country").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateMessagingConfiguration(serviceSid, country, &params)
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
			"code_length":                  AsString(SchemaOptional),
			"custom_code_enabled":          AsString(SchemaOptional),
			"do_not_share_warning_enabled": AsString(SchemaOptional),
			"dtmf_input_required":          AsString(SchemaOptional),
			"friendly_name":                AsString(SchemaOptional),
			"lookup_enabled":               AsString(SchemaOptional),
			"psd2_enabled":                 AsString(SchemaOptional),
			"push.apn_credential_sid":      AsString(SchemaOptional),
			"push.fcm_credential_sid":      AsString(SchemaOptional),
			"push.include_date":            AsString(SchemaOptional),
			"skip_sms_to_landlines":        AsString(SchemaOptional),
			"totp.code_length":             AsString(SchemaOptional),
			"totp.issuer":                  AsString(SchemaOptional),
			"totp.skew":                    AsString(SchemaOptional),
			"totp.time_step":               AsString(SchemaOptional),
			"tts_name":                     AsString(SchemaOptional),
			"account_sid":                  AsString(SchemaComputed),
			"date_created":                 AsString(SchemaComputed),
			"date_updated":                 AsString(SchemaComputed),
			"links":                        AsString(SchemaComputed),
			"push":                         AsString(SchemaComputed),
			"sid":                          AsString(SchemaComputed),
			"totp":                         AsString(SchemaComputed),
			"url":                          AsString(SchemaComputed),
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VerifyV2.CreateService(&params)
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

func deleteServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteService(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesRateLimitsBuckets() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesRateLimitsBuckets,
		ReadContext:   readServicesRateLimitsBuckets,
		UpdateContext: updateServicesRateLimitsBuckets,
		DeleteContext: deleteServicesRateLimitsBuckets,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"rate_limit_sid": AsString(SchemaRequired),
			"interval":       AsString(SchemaOptional),
			"max":            AsString(SchemaOptional),
			"account_sid":    AsString(SchemaComputed),
			"date_created":   AsString(SchemaComputed),
			"date_updated":   AsString(SchemaComputed),
			"sid":            AsString(SchemaComputed),
			"url":            AsString(SchemaComputed),
		},
	}
}

func createServicesRateLimitsBuckets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateBucketParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateBucket(serviceSid, rateLimitSid, &params)
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

func deleteServicesRateLimitsBuckets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteBucket(serviceSid, rateLimitSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesRateLimitsBuckets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchBucket(serviceSid, rateLimitSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesRateLimitsBuckets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateBucketParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateBucket(serviceSid, rateLimitSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesWebhooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesWebhooks,
		ReadContext:   readServicesWebhooks,
		UpdateContext: updateServicesWebhooks,
		DeleteContext: deleteServicesWebhooks,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"event_types":    AsString(SchemaOptional),
			"friendly_name":  AsString(SchemaOptional),
			"status":         AsString(SchemaOptional),
			"webhook_url":    AsString(SchemaOptional),
			"account_sid":    AsString(SchemaComputed),
			"date_created":   AsString(SchemaComputed),
			"date_updated":   AsString(SchemaComputed),
			"sid":            AsString(SchemaComputed),
			"url":            AsString(SchemaComputed),
			"webhook_method": AsString(SchemaComputed),
		},
	}
}

func createServicesWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateWebhook(serviceSid, &params)
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

func deleteServicesWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteWebhook(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchWebhook(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateWebhook(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesEntitiesFactors() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesEntitiesFactors,
		ReadContext:   readServicesEntitiesFactors,
		UpdateContext: updateServicesEntitiesFactors,
		DeleteContext: deleteServicesEntitiesFactors,
		Schema: map[string]*schema.Schema{
			"service_sid":                  AsString(SchemaRequired),
			"identity":                     AsString(SchemaRequired),
			"binding.alg":                  AsString(SchemaOptional),
			"binding.public_key":           AsString(SchemaOptional),
			"binding.secret":               AsString(SchemaOptional),
			"config.alg":                   AsString(SchemaOptional),
			"config.app_id":                AsString(SchemaOptional),
			"config.code_length":           AsString(SchemaOptional),
			"config.notification_platform": AsString(SchemaOptional),
			"config.notification_token":    AsString(SchemaOptional),
			"config.sdk_version":           AsString(SchemaOptional),
			"config.skew":                  AsString(SchemaOptional),
			"config.time_step":             AsString(SchemaOptional),
			"factor_type":                  AsString(SchemaOptional),
			"friendly_name":                AsString(SchemaOptional),
			"account_sid":                  AsString(SchemaComputed),
			"binding":                      AsString(SchemaComputed),
			"config":                       AsString(SchemaComputed),
			"date_created":                 AsString(SchemaComputed),
			"date_updated":                 AsString(SchemaComputed),
			"entity_sid":                   AsString(SchemaComputed),
			"sid":                          AsString(SchemaComputed),
			"status":                       AsString(SchemaComputed),
			"url":                          AsString(SchemaComputed),
		},
	}
}

func createServicesEntitiesFactors(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateNewFactorParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateNewFactor(serviceSid, identity, &params)
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

func deleteServicesEntitiesFactors(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteFactor(serviceSid, identity, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesEntitiesFactors(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchFactor(serviceSid, identity, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesEntitiesFactors(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateFactorParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateFactor(serviceSid, identity, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
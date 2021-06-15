/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
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
	. "github.com/twilio/twilio-go/rest/events/v1"
)

func ResourceSinks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSinks,
		ReadContext:   readSinks,
		UpdateContext: updateSinks,
		DeleteContext: deleteSinks,
		Schema: map[string]*schema.Schema{
			"description":        AsString(SchemaOptional),
			"sink_configuration": AsString(SchemaOptional),
			"sink_type":          AsString(SchemaOptional),
			"date_created":       AsString(SchemaOptional),
			"date_updated":       AsString(SchemaOptional),
			"links":              AsString(SchemaOptional),
			"sid":                AsString(SchemaOptional),
			"status":             AsString(SchemaOptional),
			"url":                AsString(SchemaOptional),
		},
	}
}

func createSinks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSinkParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.EventsV1.CreateSink(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId((*r.Sid))
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteSinks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.EventsV1.DeleteSink(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readSinks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.FetchSink(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateSinks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSinkParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.UpdateSink(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceSubscriptionsSubscribedEvents() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSubscriptionsSubscribedEvents,
		ReadContext:   readSubscriptionsSubscribedEvents,
		UpdateContext: updateSubscriptionsSubscribedEvents,
		DeleteContext: deleteSubscriptionsSubscribedEvents,
		Schema: map[string]*schema.Schema{
			"subscription_sid": AsString(SchemaRequired),
			"schema_version":   AsInt(SchemaOptional),
			"type":             AsString(SchemaOptional),
			"account_sid":      AsString(SchemaOptional),
			"url":              AsString(SchemaOptional),
		},
	}
}

func createSubscriptionsSubscribedEvents(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSubscribedEventParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	subscriptionSid := d.Get("subscription_sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.CreateSubscribedEvent(subscriptionSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId((*r.Type))
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteSubscriptionsSubscribedEvents(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	subscriptionSid := d.Get("subscription_sid").(string)
	type_ := d.Get("type").(string)

	err := m.(*client.Config).Client.EventsV1.DeleteSubscribedEvent(subscriptionSid, type_)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readSubscriptionsSubscribedEvents(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	subscriptionSid := d.Get("subscription_sid").(string)
	type_ := d.Get("type").(string)

	r, err := m.(*client.Config).Client.EventsV1.FetchSubscribedEvent(subscriptionSid, type_)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateSubscriptionsSubscribedEvents(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSubscribedEventParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	subscriptionSid := d.Get("subscription_sid").(string)
	type_ := d.Get("type").(string)

	r, err := m.(*client.Config).Client.EventsV1.UpdateSubscribedEvent(subscriptionSid, type_, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceSubscriptions() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSubscriptions,
		ReadContext:   readSubscriptions,
		UpdateContext: updateSubscriptions,
		DeleteContext: deleteSubscriptions,
		Schema: map[string]*schema.Schema{
			"description":  AsString(SchemaOptional),
			"sink_sid":     AsString(SchemaOptional),
			"types":        AsString(SchemaOptional),
			"account_sid":  AsString(SchemaOptional),
			"date_created": AsString(SchemaOptional),
			"date_updated": AsString(SchemaOptional),
			"links":        AsString(SchemaOptional),
			"sid":          AsString(SchemaOptional),
			"url":          AsString(SchemaOptional),
		},
	}
}

func createSubscriptions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSubscriptionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.EventsV1.CreateSubscription(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId((*r.Sid))
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteSubscriptions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.EventsV1.DeleteSubscription(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readSubscriptions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.FetchSubscription(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateSubscriptions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSubscriptionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.EventsV1.UpdateSubscription(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

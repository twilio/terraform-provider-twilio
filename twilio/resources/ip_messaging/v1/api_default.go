/*
 * Twilio - Ip_messaging
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
	. "github.com/twilio/twilio-go/rest/ip_messaging/v1"
)

func ResourceServices() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServices,
		ReadContext:   readServices,
		UpdateContext: updateServices,
		DeleteContext: deleteServices,
		Schema: map[string]*schema.Schema{
			"friendly_name":                    AsString(SchemaOptional),
			"account_sid":                      AsString(SchemaComputed),
			"consumption_report_interval":      AsInt(SchemaComputed),
			"date_created":                     AsString(SchemaComputed),
			"date_updated":                     AsString(SchemaComputed),
			"default_channel_creator_role_sid": AsString(SchemaComputed),
			"default_channel_role_sid":         AsString(SchemaComputed),
			"default_service_role_sid":         AsString(SchemaComputed),
			"limits":                           AsString(SchemaComputed),
			"links":                            AsString(SchemaComputed),
			"notifications":                    AsString(SchemaComputed),
			"post_webhook_url":                 AsString(SchemaComputed),
			"pre_webhook_url":                  AsString(SchemaComputed),
			"reachability_enabled":             AsBool(SchemaComputed),
			"read_status_enabled":              AsBool(SchemaComputed),
			"sid":                              AsString(SchemaComputed),
			"typing_indicator_timeout":         AsInt(SchemaComputed),
			"url":                              AsString(SchemaComputed),
			"webhook_filters":                  AsList(AsString(SchemaComputed), SchemaComputed),
			"webhook_method":                   AsString(SchemaComputed),
			"webhooks":                         AsString(SchemaComputed),
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.IpMessagingV1.CreateService(&params)
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

	err := m.(*client.Config).Client.IpMessagingV1.DeleteService(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.FetchService(sid)
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

	r, err := m.(*client.Config).Client.IpMessagingV1.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesUsers,
		ReadContext:   readServicesUsers,
		UpdateContext: updateServicesUsers,
		DeleteContext: deleteServicesUsers,
		Schema: map[string]*schema.Schema{
			"service_sid":           AsString(SchemaRequired),
			"attributes":            AsString(SchemaOptional),
			"friendly_name":         AsString(SchemaOptional),
			"identity":              AsString(SchemaOptional),
			"role_sid":              AsString(SchemaOptional),
			"account_sid":           AsString(SchemaComputed),
			"date_created":          AsString(SchemaComputed),
			"date_updated":          AsString(SchemaComputed),
			"is_notifiable":         AsBool(SchemaComputed),
			"is_online":             AsBool(SchemaComputed),
			"joined_channels_count": AsInt(SchemaComputed),
			"links":                 AsString(SchemaComputed),
			"sid":                   AsString(SchemaComputed),
			"url":                   AsString(SchemaComputed),
		},
	}
}

func createServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.CreateUser(serviceSid, &params)
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

func deleteServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV1.DeleteUser(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.FetchUser(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.UpdateUser(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesChannels() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannels,
		ReadContext:   readServicesChannels,
		UpdateContext: updateServicesChannels,
		DeleteContext: deleteServicesChannels,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"attributes":     AsString(SchemaOptional),
			"friendly_name":  AsString(SchemaOptional),
			"type":           AsString(SchemaOptional),
			"unique_name":    AsString(SchemaOptional),
			"account_sid":    AsString(SchemaComputed),
			"created_by":     AsString(SchemaComputed),
			"date_created":   AsString(SchemaComputed),
			"date_updated":   AsString(SchemaComputed),
			"links":          AsString(SchemaComputed),
			"members_count":  AsInt(SchemaComputed),
			"messages_count": AsInt(SchemaComputed),
			"sid":            AsString(SchemaComputed),
			"url":            AsString(SchemaComputed),
		},
	}
}

func createServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.CreateChannel(serviceSid, &params)
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

func deleteServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV1.DeleteChannel(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.FetchChannel(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.UpdateChannel(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesChannelsMembers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsMembers,
		ReadContext:   readServicesChannelsMembers,
		UpdateContext: updateServicesChannelsMembers,
		DeleteContext: deleteServicesChannelsMembers,
		Schema: map[string]*schema.Schema{
			"service_sid":                 AsString(SchemaRequired),
			"channel_sid":                 AsString(SchemaRequired),
			"identity":                    AsString(SchemaOptional),
			"role_sid":                    AsString(SchemaOptional),
			"account_sid":                 AsString(SchemaComputed),
			"date_created":                AsString(SchemaComputed),
			"date_updated":                AsString(SchemaComputed),
			"last_consumed_message_index": AsInt(SchemaComputed),
			"last_consumption_timestamp":  AsString(SchemaComputed),
			"sid":                         AsString(SchemaComputed),
			"url":                         AsString(SchemaComputed),
		},
	}
}

func createServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateMemberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.CreateMember(serviceSid, channelSid, &params)
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

func deleteServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV1.DeleteMember(serviceSid, channelSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.FetchMember(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateMemberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.UpdateMember(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesChannelsMessages() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsMessages,
		ReadContext:   readServicesChannelsMessages,
		UpdateContext: updateServicesChannelsMessages,
		DeleteContext: deleteServicesChannelsMessages,
		Schema: map[string]*schema.Schema{
			"service_sid":  AsString(SchemaRequired),
			"channel_sid":  AsString(SchemaRequired),
			"attributes":   AsString(SchemaOptional),
			"body":         AsString(SchemaOptional),
			"from":         AsString(SchemaOptional),
			"account_sid":  AsString(SchemaComputed),
			"date_created": AsString(SchemaComputed),
			"date_updated": AsString(SchemaComputed),
			"index":        AsInt(SchemaComputed),
			"sid":          AsString(SchemaComputed),
			"to":           AsString(SchemaComputed),
			"url":          AsString(SchemaComputed),
			"was_edited":   AsBool(SchemaComputed),
		},
	}
}

func createServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.CreateMessage(serviceSid, channelSid, &params)
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

func deleteServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV1.DeleteMessage(serviceSid, channelSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.FetchMessage(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.UpdateMessage(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesRoles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesRoles,
		ReadContext:   readServicesRoles,
		UpdateContext: updateServicesRoles,
		DeleteContext: deleteServicesRoles,
		Schema: map[string]*schema.Schema{
			"service_sid":   AsString(SchemaRequired),
			"friendly_name": AsString(SchemaOptional),
			"permission":    AsString(SchemaOptional),
			"type":          AsString(SchemaOptional),
			"account_sid":   AsString(SchemaComputed),
			"date_created":  AsString(SchemaComputed),
			"date_updated":  AsString(SchemaComputed),
			"permissions":   AsList(AsString(SchemaComputed), SchemaComputed),
			"sid":           AsString(SchemaComputed),
			"url":           AsString(SchemaComputed),
		},
	}
}

func createServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.CreateRole(serviceSid, &params)
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

func deleteServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV1.DeleteRole(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.FetchRole(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.UpdateRole(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCredentials,
		ReadContext:   readCredentials,
		UpdateContext: updateCredentials,
		DeleteContext: deleteCredentials,
		Schema: map[string]*schema.Schema{
			"api_key":       AsString(SchemaOptional),
			"certificate":   AsString(SchemaOptional),
			"friendly_name": AsString(SchemaOptional),
			"private_key":   AsString(SchemaOptional),
			"sandbox":       AsBool(SchemaOptional),
			"secret":        AsString(SchemaOptional),
			"type":          AsString(SchemaOptional),
			"account_sid":   AsString(SchemaComputed),
			"date_created":  AsString(SchemaComputed),
			"date_updated":  AsString(SchemaComputed),
			"sid":           AsString(SchemaComputed),
			"url":           AsString(SchemaComputed),
		},
	}
}

func createCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.IpMessagingV1.CreateCredential(&params)
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

func deleteCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV1.DeleteCredential(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.FetchCredential(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV1.UpdateCredential(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

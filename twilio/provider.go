// Package twilio provides bindings for Twilio's REST APIs.
package twilio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	twilio "github.com/twilio/terraform-provider-twilio/client"
	accountsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/accounts/v1"
	apiV2010 "github.com/twilio/terraform-provider-twilio/twilio/resources/api/v2010"
	autopilotV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/autopilot/v1"
	chatV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/chat/v1"
	chatV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/chat/v2"
	conversationsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/conversations/v1"
	eventsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/events/v1"
	faxV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/fax/v1"
	flexV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/flex/v1"
	ip_messagingV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/ip_messaging/v1"
	ip_messagingV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/ip_messaging/v2"
	messagingV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/messaging/v1"
	notifyV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/notify/v1"
	numbersV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/numbers/v2"
	proxyV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/proxy/v1"
	serverlessV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/serverless/v1"
	studioV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/studio/v1"
	studioV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/studio/v2"
	syncV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/sync/v1"
	taskrouterV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/taskrouter/v1"
	trunkingV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/trunking/v1"
	trusthubV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/trusthub/v1"
	verifyV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/verify/v2"
	videoV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/video/v1"
	voiceV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/voice/v1"
	wirelessV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/wireless/v1"
	client "github.com/twilio/twilio-go"
)

// Provider initializes terraform-provider-twilio.
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_ACCOUNT_SID", nil),
				Description: "Your Account SID can be found on the Twilio dashboard at www.twilio.com/console.",
				Required:    true,
			},
			"auth_token": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
				Description: "Your Auth Token can be found on the Twilio dashboard at www.twilio.com/console.",
				Required:    true,
			},
			"subaccount_sid": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_SUBACCOUNT_SID", nil),
				Description: "Your SubAccount SID can be found on the Twilio dashboard at www.twilio.com/console.",
				Optional:    true,
			},
			"edge": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_EDGE", nil),
				Description: "https://www.twilio.com/docs/global-infrastructure/edge-locations#public-edge-locations",
				Optional:    true,
			},
			"region": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_REGION", nil),
				Description: "https://www.twilio.com/docs/global-infrastructure/edge-locations/legacy-regions",
				Optional:    true,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"accounts_credentials_public_keys_v1":                     accountsV1.ResourceCredentialsPublicKeys(),
			"accounts_credentials_v1":                                 accountsV1.ResourceCredentialsAWS(),
			"api_accounts_ip_access_control_lists_v2010":              apiV2010.ResourceAccountsSIPIpAccessControlLists(),
			"api_accounts_domains_v2010":                              apiV2010.ResourceAccountsSIPDomains(),
			"api_accounts_credential_lists_credentials_v2010":         apiV2010.ResourceAccountsSIPCredentialListsCredentials(),
			"api_accounts_ip_access_control_lists_ip_addresses_v2010": apiV2010.ResourceAccountsSIPIpAccessControlListsIpAddresses(),
			"api_accounts_incoming_phone_numbers_v2010":               apiV2010.ResourceAccountsIncomingPhoneNumbers(),
			"api_accounts_messages_v2010":                             apiV2010.ResourceAccountsMessages(),
			"api_accounts_calls_recordings_v2010":                     apiV2010.ResourceAccountsCallsRecordings(),
			"api_accounts_queues_v2010":                               apiV2010.ResourceAccountsQueues(),
			"api_accounts_credential_lists_v2010":                     apiV2010.ResourceAccountsSIPCredentialLists(),
			"api_accounts_applications_v2010":                         apiV2010.ResourceAccountsApplications(),
			"api_accounts_addresses_v2010":                            apiV2010.ResourceAccountsAddresses(),
			"api_accounts_conferences_participants_v2010":             apiV2010.ResourceAccountsConferencesParticipants(),
			"api_accounts_calls_v2010":                                apiV2010.ResourceAccountsCalls(),
			"api_accounts_outgoing_caller_ids_v2010":                  apiV2010.ResourceAccountsOutgoingCallerIds(),
			"api_accounts_signing_keys_v2010":                         apiV2010.ResourceAccountsSigningKeys(),
			"api_accounts_usage_triggers_v2010":                       apiV2010.ResourceAccountsUsageTriggers(),
			"api_accounts_keys_v2010":                                 apiV2010.ResourceAccountsKeys(),
			"autopilot_assistants_field_types_v1":                     autopilotV1.ResourceAssistantsFieldTypes(),
			"autopilot_assistants_tasks_samples_v1":                   autopilotV1.ResourceAssistantsTasksSamples(),
			"autopilot_assistants_webhooks_v1":                        autopilotV1.ResourceAssistantsWebhooks(),
			"autopilot_assistants_v1":                                 autopilotV1.ResourceAssistants(),
			"autopilot_assistants_tasks_v1":                           autopilotV1.ResourceAssistantsTasks(),
			"autopilot_assistants_model_builds_v1":                    autopilotV1.ResourceAssistantsModelBuilds(),
			"autopilot_assistants_queries_v1":                         autopilotV1.ResourceAssistantsQueries(),
			"chat_services_v1":                                        chatV1.ResourceServices(),
			"chat_services_users_v1":                                  chatV1.ResourceServicesUsers(),
			"chat_services_channels_v1":                               chatV1.ResourceServicesChannels(),
			"chat_services_channels_members_v1":                       chatV1.ResourceServicesChannelsMembers(),
			"chat_services_channels_messages_v1":                      chatV1.ResourceServicesChannelsMessages(),
			"chat_services_roles_v1":                                  chatV1.ResourceServicesRoles(),
			"chat_credentials_v1":                                     chatV1.ResourceCredentials(),
			"chat_services_v2":                                        chatV2.ResourceServices(),
			"chat_services_users_v2":                                  chatV2.ResourceServicesUsers(),
			"chat_services_channels_v2":                               chatV2.ResourceServicesChannels(),
			"chat_services_channels_members_v2":                       chatV2.ResourceServicesChannelsMembers(),
			"chat_services_channels_messages_v2":                      chatV2.ResourceServicesChannelsMessages(),
			"chat_services_roles_v2":                                  chatV2.ResourceServicesRoles(),
			"chat_services_channels_webhooks_v2":                      chatV2.ResourceServicesChannelsWebhooks(),
			"chat_credentials_v2":                                     chatV2.ResourceCredentials(),
			"conversations_services_conversations_v1":                 conversationsV1.ResourceServicesConversations(),
			"conversations_conversations_messages_v1":                 conversationsV1.ResourceConversationsMessages(),
			"conversations_services_conversations_participants_v1":    conversationsV1.ResourceServicesConversationsParticipants(),
			"conversations_services_conversations_webhooks_v1":        conversationsV1.ResourceServicesConversationsWebhooks(),
			"conversations_roles_v1":                                  conversationsV1.ResourceRoles(),
			"conversations_users_v1":                                  conversationsV1.ResourceUsers(),
			"conversations_services_roles_v1":                         conversationsV1.ResourceServicesRoles(),
			"conversations_conversations_v1":                          conversationsV1.ResourceConversations(),
			"conversations_services_conversations_messages_v1":        conversationsV1.ResourceServicesConversationsMessages(),
			"conversations_conversations_participants_v1":             conversationsV1.ResourceConversationsParticipants(),
			"conversations_services_users_v1":                         conversationsV1.ResourceServicesUsers(),
			"conversations_conversations_webhooks_v1":                 conversationsV1.ResourceConversationsWebhooks(),
			"conversations_credentials_v1":                            conversationsV1.ResourceCredentials(),
			"events_sinks_v1":                                         eventsV1.ResourceSinks(),
			"events_subscriptions_v1":                                 eventsV1.ResourceSubscriptions(),
			"events_subscriptions_subscribed_events_v1":               eventsV1.ResourceSubscriptionsSubscribedEvents(),
			"fax_faxes_v1":                                            faxV1.ResourceFaxes(),
			"flex_flex_flows_v1":                                      flexV1.ResourceFlexFlows(),
			"flex_web_channels_v1":                                    flexV1.ResourceWebChannels(),
			"ip_messaging_services_v1":                                ip_messagingV1.ResourceServices(),
			"ip_messaging_services_users_v1":                          ip_messagingV1.ResourceServicesUsers(),
			"ip_messaging_services_channels_v1":                       ip_messagingV1.ResourceServicesChannels(),
			"ip_messaging_services_channels_members_v1":               ip_messagingV1.ResourceServicesChannelsMembers(),
			"ip_messaging_services_channels_messages_v1":              ip_messagingV1.ResourceServicesChannelsMessages(),
			"ip_messaging_services_roles_v1":                          ip_messagingV1.ResourceServicesRoles(),
			"ip_messaging_credentials_v1":                             ip_messagingV1.ResourceCredentials(),
			"ip_messaging_services_v2":                                ip_messagingV2.ResourceServices(),
			"ip_messaging_services_users_v2":                          ip_messagingV2.ResourceServicesUsers(),
			"ip_messaging_services_channels_v2":                       ip_messagingV2.ResourceServicesChannels(),
			"ip_messaging_services_channels_members_v2":               ip_messagingV2.ResourceServicesChannelsMembers(),
			"ip_messaging_services_channels_messages_v2":              ip_messagingV2.ResourceServicesChannelsMessages(),
			"ip_messaging_services_roles_v2":                          ip_messagingV2.ResourceServicesRoles(),
			"ip_messaging_services_channels_webhooks_v2":              ip_messagingV2.ResourceServicesChannelsWebhooks(),
			"ip_messaging_credentials_v2":                             ip_messagingV2.ResourceCredentials(),
			"messaging_services_v1":                                   messagingV1.ResourceServices(),
			"notify_services_v1":                                      notifyV1.ResourceServices(),
			"notify_credentials_v1":                                   notifyV1.ResourceCredentials(),
			"numbers_regulatory_compliance_supporting_documents_v2":   numbersV2.ResourceRegulatoryComplianceSupportingDocuments(),
			"numbers_regulatory_compliance_end_users_v2":              numbersV2.ResourceRegulatoryComplianceEndUsers(),
			"numbers_regulatory_compliance_bundles_v2":                numbersV2.ResourceRegulatoryComplianceBundles(),
			"proxy_services_v1":                                       proxyV1.ResourceServices(),
			"proxy_services_short_codes_v1":                           proxyV1.ResourceServicesShortCodes(),
			"proxy_services_phone_numbers_v1":                         proxyV1.ResourceServicesPhoneNumbers(),
			"proxy_services_sessions_v1":                              proxyV1.ResourceServicesSessions(),
			"serverless_services_v1":                                  serverlessV1.ResourceServices(),
			"serverless_services_environments_variables_v1":           serverlessV1.ResourceServicesEnvironmentsVariables(),
			"serverless_services_functions_v1":                        serverlessV1.ResourceServicesFunctions(),
			"serverless_services_assets_v1":                           serverlessV1.ResourceServicesAssets(),
			"studio_flows_executions_v1":                              studioV1.ResourceFlowsExecutions(),
			"studio_flows_v2":                                         studioV2.ResourceFlows(),
			"studio_flows_executions_v2":                              studioV2.ResourceFlowsExecutions(),
			"sync_services_v1":                                        syncV1.ResourceServices(),
			"sync_services_streams_v1":                                syncV1.ResourceServicesStreams(),
			"sync_services_documents_v1":                              syncV1.ResourceServicesDocuments(),
			"sync_services_lists_items_v1":                            syncV1.ResourceServicesListsItems(),
			"sync_services_maps_items_v1":                             syncV1.ResourceServicesMapsItems(),
			"sync_services_lists_v1":                                  syncV1.ResourceServicesLists(),
			"sync_services_maps_v1":                                   syncV1.ResourceServicesMaps(),
			"taskrouter_workspaces_task_queues_v1":                    taskrouterV1.ResourceWorkspacesTaskQueues(),
			"taskrouter_workspaces_tasks_v1":                          taskrouterV1.ResourceWorkspacesTasks(),
			"taskrouter_workspaces_v1":                                taskrouterV1.ResourceWorkspaces(),
			"taskrouter_workspaces_workflows_v1":                      taskrouterV1.ResourceWorkspacesWorkflows(),
			"taskrouter_workspaces_task_channels_v1":                  taskrouterV1.ResourceWorkspacesTaskChannels(),
			"taskrouter_workspaces_workers_v1":                        taskrouterV1.ResourceWorkspacesWorkers(),
			"taskrouter_workspaces_activities_v1":                     taskrouterV1.ResourceWorkspacesActivities(),
			"trunking_trunks_v1":                                      trunkingV1.ResourceTrunks(),
			"trunking_trunks_origination_urls_v1":                     trunkingV1.ResourceTrunksOriginationUrls(),
			"trusthub_customer_profiles_v1":                           trusthubV1.ResourceCustomerProfiles(),
			"trusthub_supporting_documents_v1":                        trusthubV1.ResourceSupportingDocuments(),
			"trusthub_end_users_v1":                                   trusthubV1.ResourceEndUsers(),
			"trusthub_trust_products_v1":                              trusthubV1.ResourceTrustProducts(),
			"verify_services_rate_limits_v2":                          verifyV2.ResourceServicesRateLimits(),
			"verify_services_messaging_configurations_v2":             verifyV2.ResourceServicesMessagingConfigurations(),
			"verify_services_v2":                                      verifyV2.ResourceServices(),
			"verify_services_rate_limits_buckets_v2":                  verifyV2.ResourceServicesRateLimitsBuckets(),
			"verify_services_webhooks_v2":                             verifyV2.ResourceServicesWebhooks(),
			"verify_services_entities_factors_v2":                     verifyV2.ResourceServicesEntitiesFactors(),
			"video_composition_hooks_v1":                              videoV1.ResourceCompositionHooks(),
			"voice_source_ip_mappings_v1":                             voiceV1.ResourceSourceIpMappings(),
			"voice_connection_policies_targets_v1":                    voiceV1.ResourceConnectionPoliciesTargets(),
			"voice_connection_policies_v1":                            voiceV1.ResourceConnectionPolicies(),
			"voice_ip_records_v1":                                     voiceV1.ResourceIpRecords(),
			"voice_byoc_trunks_v1":                                    voiceV1.ResourceByocTrunks(),
			"wireless_rate_plans_v1":                                  wirelessV1.ResourceRatePlans(),
		},
	}

	p.ConfigureContextFunc = providerClient(p)

	return p
}

func providerClient(p *schema.Provider) schema.ConfigureContextFunc {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var TwilioClient *client.RestClient

		username := d.Get("account_sid").(string)
		password := d.Get("auth_token").(string)
		region := d.Get("region").(string)
		edge := d.Get("edge").(string)

		if d.Get("subaccount_sid") != nil {
			params := client.RestClientParams{AccountSid: d.Get("subaccount_sid").(string)}
			TwilioClient = client.NewRestClientWithParams(username, password, params)
		} else {
			TwilioClient = client.NewRestClient(username, password)
		}

		TwilioClient.SetRegion(region)
		TwilioClient.SetEdge(edge)

		config := &twilio.Config{
			Client: TwilioClient,
		}

		return config, nil
	}
}

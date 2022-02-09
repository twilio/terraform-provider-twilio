package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	accountsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/accounts/v1"
	apiV2010 "github.com/twilio/terraform-provider-twilio/twilio/resources/api/v2010"
	autopilotV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/autopilot/v1"
	bulkexportsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/bulkexports/v1"
	chatV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/chat/v1"
	chatV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/chat/v2"
	conversationsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/conversations/v1"
	eventsV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/events/v1"
	flexV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/flex/v1"
	ip_messagingV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/ip_messaging/v1"
	ip_messagingV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/ip_messaging/v2"
	messagingV1 "github.com/twilio/terraform-provider-twilio/twilio/resources/messaging/v1"
	numbersV2 "github.com/twilio/terraform-provider-twilio/twilio/resources/numbers/v2"
)

type TwilioResources struct {
	Map map[string]*schema.Resource
}

func NewTwilioResources() *TwilioResources {
	return &TwilioResources{
		map[string]*schema.Resource{
			"twilio_accounts_credentials_v1":                                                accountsV1.ResourceCredentialsAWS(),
			"twilio_accounts_credentials_public_keys_v1":                                    accountsV1.ResourceCredentialsPublicKeys(),
			"twilio_api_accounts_addresses_v2010":                                           apiV2010.ResourceAccountsAddresses(),
			"twilio_api_accounts_applications_v2010":                                        apiV2010.ResourceAccountsApplications(),
			"twilio_api_accounts_calls_v2010":                                               apiV2010.ResourceAccountsCalls(),
			"twilio_api_accounts_calls_feedback_summary_v2010":                              apiV2010.ResourceAccountsCallsFeedbackSummary(),
			"twilio_api_accounts_calls_recordings_v2010":                                    apiV2010.ResourceAccountsCallsRecordings(),
			"twilio_api_accounts_incoming_phone_numbers_v2010":                              apiV2010.ResourceAccountsIncomingPhoneNumbers(),
			"twilio_api_accounts_incoming_phone_numbers_assigned_add_ons_v2010":             apiV2010.ResourceAccountsIncomingPhoneNumbersAssignedAddOns(),
			"twilio_api_accounts_messages_v2010":                                            apiV2010.ResourceAccountsMessages(),
			"twilio_api_accounts_keys_v2010":                                                apiV2010.ResourceAccountsKeys(),
			"twilio_api_accounts_signing_keys_v2010":                                        apiV2010.ResourceAccountsSigningKeys(),
			"twilio_api_accounts_conferences_participants_v2010":                            apiV2010.ResourceAccountsConferencesParticipants(),
			"twilio_api_accounts_queues_v2010":                                              apiV2010.ResourceAccountsQueues(),
			"twilio_api_accounts_domains_auth_calls_credential_list_mappings_v2010":         apiV2010.ResourceAccountsSIPDomainsAuthCallsCredentialListMappings(),
			"twilio_api_accounts_domains_auth_calls_ip_access_control_list_mappings_v2010":  apiV2010.ResourceAccountsSIPDomainsAuthCallsIpAccessControlListMappings(),
			"twilio_api_accounts_domains_auth_registrations_credential_list_mappings_v2010": apiV2010.ResourceAccountsSIPDomainsAuthRegistrationsCredentialListMappings(),
			"twilio_api_accounts_credential_lists_credentials_v2010":                        apiV2010.ResourceAccountsSIPCredentialListsCredentials(),
			"twilio_api_accounts_credential_lists_v2010":                                    apiV2010.ResourceAccountsSIPCredentialLists(),
			"twilio_api_accounts_domains_credential_list_mappings_v2010":                    apiV2010.ResourceAccountsSIPDomainsCredentialListMappings(),
			"twilio_api_accounts_domains_v2010":                                             apiV2010.ResourceAccountsSIPDomains(),
			"twilio_api_accounts_ip_access_control_lists_v2010":                             apiV2010.ResourceAccountsSIPIpAccessControlLists(),
			"twilio_api_accounts_domains_ip_access_control_list_mappings_v2010":             apiV2010.ResourceAccountsSIPDomainsIpAccessControlListMappings(),
			"twilio_api_accounts_ip_access_control_lists_ip_addresses_v2010":                apiV2010.ResourceAccountsSIPIpAccessControlListsIpAddresses(),
			"twilio_api_accounts_usage_triggers_v2010":                                      apiV2010.ResourceAccountsUsageTriggers(),
			"twilio_autopilot_assistants_v1":                                                autopilotV1.ResourceAssistants(),
			"twilio_autopilot_assistants_tasks_fields_v1":                                   autopilotV1.ResourceAssistantsTasksFields(),
			"twilio_autopilot_assistants_field_types_v1":                                    autopilotV1.ResourceAssistantsFieldTypes(),
			"twilio_autopilot_assistants_field_types_field_values_v1":                       autopilotV1.ResourceAssistantsFieldTypesFieldValues(),
			"twilio_autopilot_assistants_model_builds_v1":                                   autopilotV1.ResourceAssistantsModelBuilds(),
			"twilio_autopilot_assistants_queries_v1":                                        autopilotV1.ResourceAssistantsQueries(),
			"twilio_autopilot_assistants_tasks_samples_v1":                                  autopilotV1.ResourceAssistantsTasksSamples(),
			"twilio_autopilot_assistants_tasks_v1":                                          autopilotV1.ResourceAssistantsTasks(),
			"twilio_autopilot_assistants_webhooks_v1":                                       autopilotV1.ResourceAssistantsWebhooks(),
			"twilio_bulkexports_exports_jobs_v1":                                            bulkexportsV1.ResourceExportsJobs(),
			"twilio_chat_services_channels_v1":                                              chatV1.ResourceServicesChannels(),
			"twilio_chat_credentials_v1":                                                    chatV1.ResourceCredentials(),
			"twilio_chat_services_channels_invites_v1":                                      chatV1.ResourceServicesChannelsInvites(),
			"twilio_chat_services_channels_members_v1":                                      chatV1.ResourceServicesChannelsMembers(),
			"twilio_chat_services_channels_messages_v1":                                     chatV1.ResourceServicesChannelsMessages(),
			"twilio_chat_services_roles_v1":                                                 chatV1.ResourceServicesRoles(),
			"twilio_chat_services_v1":                                                       chatV1.ResourceServices(),
			"twilio_chat_services_users_v1":                                                 chatV1.ResourceServicesUsers(),
			"twilio_chat_services_channels_v2":                                              chatV2.ResourceServicesChannels(),
			"twilio_chat_services_channels_webhooks_v2":                                     chatV2.ResourceServicesChannelsWebhooks(),
			"twilio_chat_credentials_v2":                                                    chatV2.ResourceCredentials(),
			"twilio_chat_services_channels_invites_v2":                                      chatV2.ResourceServicesChannelsInvites(),
			"twilio_chat_services_channels_members_v2":                                      chatV2.ResourceServicesChannelsMembers(),
			"twilio_chat_services_channels_messages_v2":                                     chatV2.ResourceServicesChannelsMessages(),
			"twilio_chat_services_roles_v2":                                                 chatV2.ResourceServicesRoles(),
			"twilio_chat_services_v2":                                                       chatV2.ResourceServices(),
			"twilio_chat_services_users_v2":                                                 chatV2.ResourceServicesUsers(),
			"twilio_conversations_configuration_addresses_v1":                               conversationsV1.ResourceConfigurationAddresses(),
			"twilio_conversations_conversations_v1":                                         conversationsV1.ResourceConversations(),
			"twilio_conversations_conversations_messages_v1":                                conversationsV1.ResourceConversationsMessages(),
			"twilio_conversations_conversations_participants_v1":                            conversationsV1.ResourceConversationsParticipants(),
			"twilio_conversations_conversations_webhooks_v1":                                conversationsV1.ResourceConversationsWebhooks(),
			"twilio_conversations_credentials_v1":                                           conversationsV1.ResourceCredentials(),
			"twilio_conversations_roles_v1":                                                 conversationsV1.ResourceRoles(),
			"twilio_conversations_services_v1":                                              conversationsV1.ResourceServices(),
			"twilio_conversations_services_conversations_v1":                                conversationsV1.ResourceServicesConversations(),
			"twilio_conversations_services_conversations_messages_v1":                       conversationsV1.ResourceServicesConversationsMessages(),
			"twilio_conversations_services_conversations_participants_v1":                   conversationsV1.ResourceServicesConversationsParticipants(),
			"twilio_conversations_services_conversations_webhooks_v1":                       conversationsV1.ResourceServicesConversationsWebhooks(),
			"twilio_conversations_services_roles_v1":                                        conversationsV1.ResourceServicesRoles(),
			"twilio_conversations_services_users_v1":                                        conversationsV1.ResourceServicesUsers(),
			"twilio_conversations_users_v1":                                                 conversationsV1.ResourceUsers(),
			"twilio_events_sinks_v1":                                                        eventsV1.ResourceSinks(),
			"twilio_events_subscriptions_subscribed_events_v1":                              eventsV1.ResourceSubscriptionsSubscribedEvents(),
			"twilio_events_subscriptions_v1":                                                eventsV1.ResourceSubscriptions(),
			"twilio_flex_channels_v1":                                                       flexV1.ResourceChannels(),
			"twilio_flex_flex_flows_v1":                                                     flexV1.ResourceFlexFlows(),
			"twilio_flex_web_channels_v1":                                                   flexV1.ResourceWebChannels(),
			"twilio_ip_messaging_services_channels_v1":                                      ip_messagingV1.ResourceServicesChannels(),
			"twilio_ip_messaging_credentials_v1":                                            ip_messagingV1.ResourceCredentials(),
			"twilio_ip_messaging_services_channels_invites_v1":                              ip_messagingV1.ResourceServicesChannelsInvites(),
			"twilio_ip_messaging_services_channels_members_v1":                              ip_messagingV1.ResourceServicesChannelsMembers(),
			"twilio_ip_messaging_services_channels_messages_v1":                             ip_messagingV1.ResourceServicesChannelsMessages(),
			"twilio_ip_messaging_services_roles_v1":                                         ip_messagingV1.ResourceServicesRoles(),
			"twilio_ip_messaging_services_v1":                                               ip_messagingV1.ResourceServices(),
			"twilio_ip_messaging_services_users_v1":                                         ip_messagingV1.ResourceServicesUsers(),
			"twilio_ip_messaging_services_channels_v2":                                      ip_messagingV2.ResourceServicesChannels(),
			"twilio_ip_messaging_services_channels_webhooks_v2":                             ip_messagingV2.ResourceServicesChannelsWebhooks(),
			"twilio_ip_messaging_credentials_v2":                                            ip_messagingV2.ResourceCredentials(),
			"twilio_ip_messaging_services_channels_invites_v2":                              ip_messagingV2.ResourceServicesChannelsInvites(),
			"twilio_ip_messaging_services_channels_members_v2":                              ip_messagingV2.ResourceServicesChannelsMembers(),
			"twilio_ip_messaging_services_channels_messages_v2":                             ip_messagingV2.ResourceServicesChannelsMessages(),
			"twilio_ip_messaging_services_roles_v2":                                         ip_messagingV2.ResourceServicesRoles(),
			"twilio_ip_messaging_services_v2":                                               ip_messagingV2.ResourceServices(),
			"twilio_ip_messaging_services_users_v2":                                         ip_messagingV2.ResourceServicesUsers(),
			"twilio_messaging_services_alpha_senders_v1":                                    messagingV1.ResourceServicesAlphaSenders(),
			"twilio_messaging_services_phone_numbers_v1":                                    messagingV1.ResourceServicesPhoneNumbers(),
			"twilio_messaging_services_v1":                                                  messagingV1.ResourceServices(),
			"twilio_messaging_services_short_codes_v1":                                      messagingV1.ResourceServicesShortCodes(),
			"twilio_messaging_services_compliance_usa_v1":                                   messagingV1.ResourceServicesComplianceUsa2p(),
			"twilio_numbers_regulatory_compliance_bundles_v2":                               numbersV2.ResourceRegulatoryComplianceBundles(),
			"twilio_numbers_regulatory_compliance_end_users_v2":                             numbersV2.ResourceRegulatoryComplianceEndUsers(),
			"twilio_numbers_regulatory_compliance_bundles_item_assignments_v2":              numbersV2.ResourceRegulatoryComplianceBundlesItemAssignments(),
			"twilio_numbers_regulatory_compliance_supporting_documents_v2":                  numbersV2.ResourceRegulatoryComplianceSupportingDocuments(),
		},
	}
}

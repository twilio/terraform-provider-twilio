terraform-provider-twilio changelog
====================
[2024-05-24] Version 0.18.45
----------------------------
**Api**
- Add ie1 as supported region for UserDefinedMessage and UserDefinedMessageSubscription.

**Flex**
- Adding validated field to `plugin_versions`
- Corrected the data type for `runtime_domain`, `call_recording_webhook_url`, `crm_callback_url`, `crm_fallback_url`, `flex_url` in Flex Configuration
- Making `routing` optional in Create Interactions endpoint

**Intelligence**
- Expose operator authoring apis to public visibility
- Deleted `language_code` parameter from updating service in v2 **(breaking change)**
- Add read_only_attached_operator_sids to v2 services

**Numbers**
- Add API endpoint for GET Porting Webhook Configurations By Account SID
- Remove bulk portability api under version `/v1`. **(breaking change)**
- Removed porting_port_in_fetch.json files and move the content into porting_port_in.json files
- Add API endpoint to deleting Webhook Configurations
- Add Get Phone Number by Port in request SID and Phone Number SID api
- Add Create Porting webhook configuration API
- Added bundle_sid and losing_carrier_information fields to Create PortInRequest api to support Japan porting

**Taskrouter**
- Add back `routing_target` property to tasks
- Add back `ignore_capacity` property to tasks
- Removing `routing_target` property to tasks due to revert
- Removing `ignore_capacity` property to tasks due to revert
- Add `routing_target` property to tasks
- Add `ignore_capacity` property to tasks

**Trusthub**
- Add new field errors to bundle as part of public API response in customer_profile.json and trust_product.json **(breaking change)**
- Add themeSetId field in compliance_tollfree_inquiry.

**Verify**
- Update `friendly_name` description on service docs


[2024-04-18] Version 0.18.44
----------------------------
**Flex**
- Add header `ui_version` to `web_channels` API

**Messaging**
- Redeploy after failed pipeline

**Numbers**
- Add Delete Port In request phone number api and Add Delete Port In request api


[2024-04-04] Version 0.18.43
----------------------------
**Api**
- Correct conference filtering by date_created and date_updated documentation, clarifying that times are UTC.
- Add property `queue_time` to conference participant resource
- Update RiskCheck documentation
- Correct call filtering by start and end time documentation, clarifying that times are UTC.

**Flex**
- Remove optional parameter from `plugins` and it to `plugin_versions`
- Adding optional parameter to `plugins`

**Lookups**
- Add new `pre_fill` package to the lookup response

**Media**
- Remove API: MediaProcessor

**Messaging**
- Cleanup api.messaging.next-gen from Messaging Services endpoints
- Readd Sending-Window after fixing test failure
- Remove Sending-Window due to test failure
- Add Sending-Window as a response property to Messaging Services, gated by a beta feature flag

**Numbers**
- Correct valid_until_date field to be visible in Bundles resource
- Adding port_in_status field to the Port In resource and phone_number_status and sid fields to the Port In Phone Number resource

**Oauth**
- Modified token endpoint response
- Added refresh_token and scope as optional parameter to token endpoint
- Add new APIs for vendor authorize and token endpoints

**Trusthub**
- Add update inquiry endpoint in compliance_registration.
- Add new field in themeSetId in compliance_registration.

**Verify**
- Add `whatsapp.msg_service_sid` and `whatsapp.from` parameters to create, update, get and list of services endpoints

**Voice**
- Correct conference filtering by date_created and date_updated documentation, clarifying that times are UTC.
- Correct call filtering by start and end time documentation, clarifying that times are UTC.


[2024-03-12] Version 0.18.42
----------------------------
**Api**
- Correct precedence documentation for application_sid vs status_callback in message creation
- Mark MaxPrice as deprecated

**Flex**
- Making `plugins` visibility to public

**Messaging**
- Add new `errors` attribute to the Brand Registration resource.
- Mark `brand_feedback` attribute as deprecated.
- Mark `failure_reason` attribute as deprecated.
- The new `errors` attribute is expected to provide additional information about Brand registration failures and feedback (if any has been provided by The Campaign Registry). Consumers should use this attribute instead of `brand_feedback` and `failure_reason`.

**Numbers**
- Correcting mount_name for porting port in fetch API

**Trusthub**
- Add new field in statusCallbackUrl in compliance_registration.
- Add new field in isvRegisteringForSelfOrTenant in compliance_registration.


[2024-02-27] Version 0.18.41
----------------------------
**Api**
- remove feedback and feedback summary from call resource

**Flex**
- Adding `routing_properties` to Interactions Channels Participant

**Lookups**
- Add new `line_status` package to the lookup response
- Remove `live_activity` package from the lookup response **(breaking change)**

**Messaging**
- Add tollfree multiple rejection reasons response array

**Trusthub**
- Add ENUM for businessRegistrationAuthority in compliance_registration. **(breaking change)**
- Add new field in isIsvEmbed in compliance_registration.
- Add additional optional fields in compliance_registration for Individual business type.


[2024-02-09] Version 0.18.40
----------------------------
**Api**
- Updated service base url for connect apps and authorized connect apps APIs **(breaking change)**
- Update documentation to reflect RiskCheck GA
- Added optional parameter `CallToken` for create participant api

**Events**
- Marked as GA

**Flex**
- Adding `flex_instance_sid` to Flex Configuration
- Adding `provisioning_status` for Email Manager
- Adding `offline_config` to Flex Configuration

**Insights**
- add flag to restrict access to unapid customers
- decommission voice-qualitystats-endpoint role

**Intelligence**
- Add text-generation operator (for example conversation summary) results to existing OperatorResults collection.

**Lookups**
- Remove `carrier` field from `sms_pumping_risk` and leave `carrier_risk_category` **(breaking change)**
- Remove carrier information from call forwarding package **(breaking change)**

**Messaging**
- Add update instance endpoints to us_app_to_person api
- Add tollfree edit_allowed and edit_reason fields
- Update Phone Number, Short Code, Alpha Sender, US A2P and Channel Sender documentation
- Add DELETE support to Tollfree Verification resource

**Numbers**
- Add Get Port In request api

**Push**
- Migrated to new Push API V4 with Resilient Notification Delivery.

**Serverless**
- Add node18 as a valid Build runtime

**Taskrouter**
- Add `jitter_buffer_size` param in update reservation
- Add container attribute to task_queue_bulk_real_time_statistics endpoint
- Remove beta_feature check on task_queue_bulk_real_time_statistics endpoint

**Trusthub**
- Add optional field NotificationEmail to the POST /v1/ComplianceInquiries/Customers/Initialize API
- Add additional optional fields in compliance_tollfree_inquiry.json
- Rename did to tollfree_phone_number in compliance_tollfree_inquiry.json
- Add new optional field notification_email to compliance_tollfree_inquiry.json

**Verify**
- `Tags` property added again to Public Docs **(breaking change)**
- Remove `Tags` from Public Docs **(breaking change)**
- Add `VerifyEventSubscriptionEnabled` parameter to service create and update endpoints.
- Add `Tags` optional parameter on Verification creation.
- Update Verify TOTP maturity to GA.


[2024-01-25] Version 0.18.39
----------------------------
**Oauth**
- updated openid discovery endpoint uri **(breaking change)**
- Added device code authorization endpoint
- added oauth JWKS endpoint
- Get userinfo resource
- OpenID discovery resource
- Add new API for token endpoint


[2024-01-14] Version 0.18.38
----------------------------
**Push**
- Migrated to new Push API V4 with Resilient Notification Delivery.


[2023-12-14] Version 0.18.37
----------------------------
**Api**
- Updated service base url for connect apps and authorized connect apps APIs **(breaking change)**

**Events**
- Marked as GA

**Insights**
- decommission voice-qualitystats-endpoint role

**Numbers**
- Add Get Port In request api

**Taskrouter**
- Add `jitter_buffer_size` param in update reservation

**Trusthub**
- Add additional optional fields in compliance_tollfree_inquiry.json

**Verify**
- Remove `Tags` from Public Docs **(breaking change)**


[2023-12-01] Version 0.18.36
----------------------------
**Verify**
- Add `VerifyEventSubscriptionEnabled` parameter to service create and update endpoints.


[2023-11-17] Version 0.18.35
----------------------------
**Api**
- Update documentation to reflect RiskCheck GA

**Messaging**
- Add tollfree edit_allowed and edit_reason fields
- Update Phone Number, Short Code, Alpha Sender, US A2P and Channel Sender documentation

**Taskrouter**
- Add container attribute to task_queue_bulk_real_time_statistics endpoint

**Trusthub**
- Rename did to tollfree_phone_number in compliance_tollfree_inquiry.json
- Add new optional field notification_email to compliance_tollfree_inquiry.json

**Verify**
- Add `Tags` optional parameter on Verification creation.


[2023-11-06] Version 0.18.34
----------------------------
**Flex**
- Adding `provisioning_status` for Email Manager

**Intelligence**
- Add text-generation operator (for example conversation summary) results to existing OperatorResults collection.

**Messaging**
- Add DELETE support to Tollfree Verification resource

**Serverless**
- Add node18 as a valid Build runtime

**Verify**
- Update Verify TOTP maturity to GA.


[2023-10-19] Version 0.18.33
----------------------------
**Accounts**
- Updated Safelist metadata to correct the docs.
- Add Global SafeList API changes

**Api**
- Added optional parameter `CallToken` for create participant api

**Flex**
- Adding `offline_config` to Flex Configuration

**Intelligence**
- Deleted `redacted` parameter from fetching transcript in v2 **(breaking change)**

**Lookups**
- Add new `phone_number_quality_score` package to the lookup response
- Remove `disposable_phone_number_risk` package **(breaking change)**

**Messaging**
- Update US App To Person documentation with current `message_samples` requirements

**Taskrouter**
- Remove beta_feature check on task_queue_bulk_real_time_statistics endpoint
- Add `virtual_start_time` property to tasks
- Updating `task_queue_data` format from `map` to `array` in the response of bulk get endpoint of TaskQueue Real Time Statistics API **(breaking change)**


[2023-10-05] Version 0.18.32
----------------------------
**Lookups**
- Add test api support for Lookup v2


[2023-09-21] Version 0.18.31
----------------------------
**Conversations**
- Enable conversation email bindings, email address configurations and email message subjects

**Flex**
- Adding `console_errors_included` to Flex Configuration field `debugger_integrations`
- Introducing new channel status as `inactive` in modify channel endpoint for leave functionality **(breaking change)**
- Adding `citrix_voice_vdi` to Flex Configuration

**Taskrouter**
- Add Update Queues, Workers, Workflow Real Time Statistics API to flex-rt-data-api-v2 endpoint
- Add Update Workspace Real Time Statistics API to flex-rt-data-api-v2 endpoint


[2023-09-07] Version 0.18.30
----------------------------
**Api**
- Make message tagging parameters public **(breaking change)**

**Flex**
- Adding `agent_conv_end_methods` to Flex Configuration

**Messaging**
- Mark Mesasging Services fallback_to_long_code feature obsolete

**Numbers**
- Add Create Port In request api
- Renaming sid for bulk_hosting_sid and remove account_sid response field in numbers/v2/BulkHostedNumberOrders **(breaking change)**

**Pricing**
- gate resources behind a beta_feature


[2023-08-24] Version 0.18.29
----------------------------
**Api**
- Add new property `RiskCheck` for SMS pumping protection feature only (public beta to be available soon): Include this parameter with a value of `disable` to skip any kind of risk check on the respective message request

**Flex**
- Changing `sid<UO>` path param to `sid<UT>` in interaction channel participant update endpoint **(breaking change)**

**Messaging**
- Add Channel Sender api
- Fixing country code docs and removing Zipwhip references

**Numbers**
- Request status changed in numbers/v2/BulkHostedNumberOrders **(breaking change)**
- Add bulk hosting orders API under version `/v2


[2023-08-10] Version 0.18.28
----------------------------
**Insights**
- Normalize annotations parameters in list summary api to be prefixed

**Numbers**
- Change Bulk_hosted_sid from BHR to BH prefix in HNO and dependent under version `/v2` API's. **(breaking change)**
- Added parameter target_account_sid to portability and account_sid to response body

**Verify**
- Remove beta feature flag to list attempts API.
- Remove beta feature flag to verifications summary attempts API.


[2023-07-27] Version 0.18.27
----------------------------
**Api**
- Added `voice-intelligence`, `voice-intelligence-transcription` and `voice-intelligence-operators` to `usage_record` API.
- Added `tts-google` to `usage_record` API.

**Lookups**
- Add new `disposable_phone_number_risk` package to the lookup response

**Verify**
- Documentation of list attempts API was improved by correcting `date_created_after` and `date_created_before` expected date format.
- Documentation was improved by correcting `date_created_after` and `date_created_before` expected date format parameter on attempts summary API.
- Documentation was improved by adding `WHATSAPP` as optional valid parameter on attempts summary API.


[2023-07-13] Version 0.18.26
----------------------------
**Flex**
- Adding `interaction_context_sid` as optional parameter in Interactions API

**Messaging**
- Making visiblity public for tollfree_verification API

**Numbers**
- Remove Sms capability property from HNO creation under version `/v2` of HNO API. **(breaking change)**
- Update required properties in LOA creation under version `/v2` of Authorization document API. **(breaking change)**

**Taskrouter**
- Add api to fetch task queue statistics for multiple TaskQueues

**Verify**
- Add `RiskCheck` optional parameter on Verification creation.


[2023-06-28] Version 0.18.25
----------------------------
**Lookups**
- Add `reassigned_number` package to the lookup response

**Numbers**
- Add hosted_number_order under version `/v2`.
- Update properties in Porting and Bulk Porting APIs. **(breaking change)**
- Added bulk Portability API under version `/v1`.
- Added Portability API under version `/v1`.


[2023-06-15] Version 0.18.24
----------------------------
**Api**
- Added `content_sid` as conditional parameter
- Removed `content_sid` as optional field **(breaking change)**

**Insights**
- Added `annotation` to list summary output


[2023-06-01] Version 0.18.23
----------------------------
**Api**
- Add `Trim` to create Conference Participant API

**Intelligence**
- First public beta release for Voice Intelligence APIs with client libraries

**Messaging**
- Add new `errors` attribute to us_app_to_person resource. This attribute will provide additional information about campaign registration errors.


[2023-05-18] Version 0.18.22
----------------------------
**Conversations**
- Added  `AddressCountry` parameter to Address Configuration endpoint, to support regional short code addresses
- Added query parameters `start_date`, `end_date` and `state` in list Conversations resource for filtering

**Insights**
- Added annotations parameters to list summary api

**Messaging**
- Add GET domainByMessagingService endpoint to linkShortening service
- Add `disable_https` to link shortening domain_config properties

**Numbers**
- Add bulk_eligibility api under version `/v1`.


[2023-05-04] Version 0.18.21
----------------------------
**Conversations**
- Remove `start_date`, `end_date` and `state` query parameters from list operation on Conversations resource **(breaking change)**


[2023-04-19] Version 0.18.20
----------------------------
**Messaging**
- Remove `messaging_service_sids` and `messaging_service_sid_action` from domain config endpoint **(breaking change)**
- Add error_code and rejection_reason properties to tollfree verification API response

**Numbers**
- Added the new Eligibility API under version `/v1`.


[2023-04-05] Version 0.18.19
----------------------------
**Conversations**
- Expose query parameters `start_date`, `end_date` and `state` in list operation on Conversations resource for sorting and filtering

**Insights**
- Added answered by filter in Call Summaries

**Lookups**
- Remove `disposable_phone_number_risk` package **(breaking change)**

**Messaging**
- Add support for `SOLE_PROPRIETOR` brand type and `SOLE_PROPRIETOR` campaign use case.
- New Sole Proprietor Brands should be created with `SOLE_PROPRIETOR` brand type. Brand registration requests with `STARTER` brand type will be rejected.
- New Sole Proprietor Campaigns should be created with `SOLE_PROPRIETOR` campaign use case. Campaign registration requests with `STARTER` campaign use case will be rejected.
- Add Brand Registrations OTP API


[2023-03-22] Version 0.18.18
----------------------------
**Api**
- Revert Corrected the data type for `friendly_name` in Available Phone Number Local, Mobile and TollFree resources
- Corrected the data type for `friendly_name` in Available Phone Number Local, Mobile and TollFree resources **(breaking change)**

**Messaging**
- Add `linkshortening_messaging_service` resource
- Add new endpoint for GetDomainConfigByMessagingServiceSid
- Remove `validated` parameter and add `cert_in_validation` parameter to Link Shortening API **(breaking change)**


[2023-03-09] Version 0.18.17
----------------------------
**Library - Chore**
- [PR #127](https://github.com/twilio/terraform-provider-twilio/pull/127): upgrade dependencies. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #126](https://github.com/twilio/terraform-provider-twilio/pull/126): update minimum go version for development to 1.18. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Api**
- Add new categories for whatsapp template

**Lookups**
- Remove `validation_results` from the `default_output_properties`

**Supersim**
- Add ESimProfile's `matching_id` and `activation_code` parameters to libraries


[2023-02-22] Version 0.18.16
----------------------------
**Library - Fix**
- [PR #125](https://github.com/twilio/terraform-provider-twilio/pull/125): drop bulk exports resource. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Api**
- Remove `scheduled_for` property from message resource
- Add `scheduled_for` property to message resource


[2023-02-08] Version 0.18.15
----------------------------
**Lookups**
- Add `disposable_phone_number_risk` package to the lookup response
- Add `sms_pumping_risk` package to the lookup response


[2023-01-25] Version 0.18.14
----------------------------
**Api**
- Add `public_application_connect_enabled` param to Application resource

**Messaging**
- Add new tollfree verification API property (ExternalReferenceId)]

**Verify**
- Add `device_ip` parameter and channel `auto` for sna/sms orchestration


[2023-01-11] Version 0.18.13
----------------------------
**Conversations**
- Add support for creating Multi-Channel Rich Content Messages

**Lookups**
- Changed the no data message for match postal code from `no_data` to `data_not_available` in identity match package

**Messaging**
- Add update/edit tollfree verification API


[2022-12-14] Version 0.18.12
----------------------------
**Api**
- Add `street_secondary` param to address create and update
- Make `method` optional for user defined message subscription **(breaking change)**

**Flex**
- Flex Conversations is now Generally Available
- Adding the ie1 mapping for authorization api, updating service base uri and base url response attribute **(breaking change)**
- Change web channels to GA and library visibility to public
- Changing the uri for authorization api from using Accounts to Insights **(breaking change)**

**Media**
- Gate Twilio Live endpoints behind beta_feature for EOS

**Messaging**
- Mark `MessageFlow` as a required field for Campaign Creation **(breaking change)**

**Oauth**
- updated openid discovery endpoint uri **(breaking change)**
- Added device code authorization endpoint

**Supersim**
- Allow filtering the SettingsUpdates resource by `status`


[2022-11-30] Version 0.18.11
----------------------------
**Flex**
- Adding new `assessments` api in version `v1`

**Lookups**
- Add `identity_match` package to the lookup response

**Messaging**
- Added `validated` parameter to Link Shortening API

**Serverless**
- Add node16 as a valid Build runtime
- Add ie1 and au1 as supported regions for all endpoints.


[2022-11-16] Version 0.18.10
----------------------------
**Library - Chore**
- [PR #119](https://github.com/twilio/terraform-provider-twilio/pull/119): upgrade GitHub Actions dependencies. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Api**
- Set the Content resource to have public visibility as Preview

**Flex**
- Adding new parameter `base_url` to 'gooddata' response in version `v1`

**Insights**
- Added `answered_by` field in List Call Summary
- Added `answered_by` field in call summary


[2022-11-10] Version 0.18.9
---------------------------
**Flex**
- Adding two new authorization API 'user_roles' and 'gooddata' in version `v1`

**Messaging**
- Add new Campaign properties (MessageFlow, OptInMessage, OptInKeywords, OptOutMessage, OptOutKeywords, HelpMessage, HelpKeywords)


[2022-10-31] Version 0.18.8
---------------------------
**Api**
- Added `contentSid` and `contentVariables` to Message resource with public visibility as Beta
- Add `UserDefinedMessageSubscription` and `UserDefinedMessage` resource

**Proxy**
- Remove FailOnParticipantConflict param from Proxy Session create and update and Proxy Participant create

**Supersim**
- Update SettingsUpdates resource to remove PackageSid

**Taskrouter**
- Add `Ordering` query parameter to Workers and TaskQueues for sorting by
- Add `worker_sid` query param for list reservations endpoint


[2022-10-19] Version 0.18.7
---------------------------
**Api**
- Make link shortening parameters public **(breaking change)**

**Oauth**
- added oauth JWKS endpoint
- Get userinfo resource
- OpenID discovery resource
- Add new API for token endpoint

**Supersim**
- Add SettingsUpdates resource

**Verify**
- Update Verify Push endpoints to `ga` maturity
- Verify BYOT add Channels property to the Get Templates response


[2022-10-05] Version 0.18.6
---------------------------
**Api**
- Added `virtual-agent` to `usage_record` API.
- Add AMD attributes to participant create request


[2022-09-07] Version 0.18.5
---------------------------
**Flex**
- Removed redundant `close` status from Flex Interactions flow **(breaking change)**
- Adding `debugger_integration` and `flex_ui_status_report` to Flex Configuration

**Messaging**
- Add create, list and get tollfree verification API

**Verify**
- Verify SafeList API endpoints added.

**Video**
- Add `Anonymize` API


[2022-08-24] Version 0.18.4
---------------------------
**Library - Test**
- [PR #116](https://github.com/twilio/terraform-provider-twilio/pull/116): add test-docker rule & Dockerfile. Thanks to [@beebzz](https://github.com/beebzz)!

**Api**
- Remove `beta feature` from scheduling params and remove optimize parameters. **(breaking change)**

**Routes**
- Remove Duplicate Create Method - Update Method will work even if Inbound Processing Region is currently empty/404. **(breaking change)**


[2022-08-10] Version 0.18.3
---------------------------
**Routes**
- Inbound Proccessing Region API - Public GA

**Supersim**
- Allow updating `DataLimit` on a Fleet


[2022-07-21] Version 0.18.2
---------------------------
**Library - Miscellaneous**
- [PR #115](https://github.com/twilio/terraform-provider-twilio/pull/115): migrate to crazy-max/ghaction-import-gpg. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Fix**
- [PR #113](https://github.com/twilio/terraform-provider-twilio/pull/113): mark Terraform params as ForceNew when not part of update operation. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Test**
- [PR #114](https://github.com/twilio/terraform-provider-twilio/pull/114): Adding misc as PR type. Thanks to [@rakatyal](https://github.com/rakatyal)!

**Conversations**
- Allowed to use `identity` as part of Participant's resource **(breaking change)**

**Flex**
- Add `status`, `error_code`, and `error_message` fields to Interaction `Channel`
- Adding `messenger` and `gbm` as supported channels for Interactions API

**Lookups**
- Remove `enhanced_line_type` from the lookup response **(breaking change)**

**Messaging**
- Update alpha_sender docs with new valid characters

**Supersim**
- Add support for `sim_ip_addresses` resource to helper libraries

**Verify**
- Reorder Verification Check parameters so `code` stays as the first parameter **(breaking change)**
- Rollback List Attempts API V2 back to pilot stage.
- Changed summary param `service_sid` to `verify_service_sid` to be consistent with list attempts API **(breaking change)**
- Make `code` optional on Verification check to support `sna` attempts.


[2022-06-29] Version 0.18.1
---------------------------
**Api**
- Added `amazon-polly` to `usage_record` API.

**Insights**
- Added `annotation` field in call summary
- Added new endpoint to fetch/create/update Call Annotations

**Verify**
- Remove `api.verify.totp` beta flag and set maturity to `beta` for Verify TOTP properties and parameters. **(breaking change)**
- Changed summary param `verify_service_sid` to `service_sid` to be consistent with list attempts API **(breaking change)**


[2022-06-15] Version 0.18.0
---------------------------
**Note:** This release contains breaking changes, check our [upgrade guide](./UPGRADE.md#2022-06-01-017x-to-018x) for detailed migration notes.

**Library - Fix**
- [PR #111](https://github.com/twilio/terraform-provider-twilio/pull/111): correct the resource name for certain AWS, SIP, and Usa2p resources. Thanks to [@childish-sambino](https://github.com/childish-sambino)! **(breaking change)**

**Lookups**
- Adding support for Lookup V2 API

**Studio**
- Corrected PII labels to be 30 days and added context to be PII


[2022-05-20] Version 0.17.0
---------------------------
**Note:** This release contains breaking changes, check our [upgrade guide](./UPGRADE.md#2022-05-20-016x-to-017x) for detailed migration notes.

**Library - Fix**
- [PR #110](https://github.com/twilio/terraform-provider-twilio/pull/110): Removing prefix for v2010 apis. Thanks to [@rakatyal](https://github.com/rakatyal)!

**Library - Docs**
- [PR #109](https://github.com/twilio/terraform-provider-twilio/pull/109): remove Twilio insiders program note. Thanks to [@JenniferMah](https://github.com/JenniferMah)!
- [PR #108](https://github.com/twilio/terraform-provider-twilio/pull/108): Adding upgrade guide. Thanks to [@rakatyal](https://github.com/rakatyal)!

**Library - Chore**
- [PR #107](https://github.com/twilio/terraform-provider-twilio/pull/107): Removing v2010 prefix for terraform. Thanks to [@rakatyal](https://github.com/rakatyal)! **(breaking change)**

**Api**
- Add property `media_url` to the recording resources

**Verify**
- Include `silent` as a channel type in the verifications API.


[2022-05-04] Version 0.16.0
---------------------------
**Library - Test**
- [PR #106](https://github.com/twilio/terraform-provider-twilio/pull/106): update cluster tests to allow for parallel execution. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Fix**
- [PR #105](https://github.com/twilio/terraform-provider-twilio/pull/105): unpack the interface type when marshaling response schemas. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Conversations**
- Expose query parameter `type` in list operation on Address Configurations resource

**Supersim**
- Add `data_total_billed` and `billed_units` fields to Super SIM UsageRecords API response.
- Change ESimProfiles `Eid` parameter to optional to enable Activation Code download method support **(breaking change)**

**Verify**
- Deprecate `push.include_date` parameter in create and update service.


[2022-04-20] Version 0.15.1
---------------------------
**Library - Test**
- [PR #103](https://github.com/twilio/terraform-provider-twilio/pull/103): add testing for go-1.18. Thanks to [@childish-sambino](https://github.com/childish-sambino)!


[2022-04-06] Version 0.15.0
---------------------------
**Library - Fix**
- [PR #102](https://github.com/twilio/terraform-provider-twilio/pull/102): use go install instead of go get. Thanks to [@beebzz](https://github.com/beebzz)!

**Library - Feature**
- [PR #101](https://github.com/twilio/terraform-provider-twilio/pull/101): update cluster tests authentication. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Api**
- Updated `provider_sid` visibility to private

**Verify**
- Verify List Attempts API summary endpoint added.
- Update PII documentation for `AccessTokens` `factor_friendly_name` property.

**Voice**
- make annotation parameter from /Calls API private


[2022-03-23] Version 0.14.0
---------------------------
**Library - Fix**
- [PR #100](https://github.com/twilio/terraform-provider-twilio/pull/100): renaming RestClientParams to ClientParams. Thanks to [@rakatyal](https://github.com/rakatyal)!
- [PR #98](https://github.com/twilio/terraform-provider-twilio/pull/98): goimports -> tidy -> test. Thanks to [@beebzz](https://github.com/beebzz)!
- [PR #97](https://github.com/twilio/terraform-provider-twilio/pull/97): run go mod tidy before go vet. Thanks to [@beebzz](https://github.com/beebzz)!

**Api**
- Change `stream` url parameter to non optional
- Add `verify-totp` and `verify-whatsapp-conversations-business-initiated` categories to `usage_record` API

**Chat**
- Added v3 Channel update endpoint to support Public to Private channel migration

**Flex**
- Private Beta release of the Interactions API to support the upcoming release of Flex Conversations at the end of Q1 2022.
- Adding `channel_configs` object to Flex Configuration

**Media**
- Add max_duration param to PlayerStreamer

**Supersim**
- Remove Commands resource, use SmsCommands resource instead **(breaking change)**

**Taskrouter**
- Add limits to `split_by_wait_time` for Cumulative Statistics Endpoint

**Video**
- Change recording `status_callback_method` type from `enum` to `http_method` **(breaking change)**
- Add `status_callback` and `status_callback_method` to composition
- Add `status_callback` and `status_callback_method` to recording


[2022-03-09] Version 0.13.2
---------------------------
**Library - Chore**
- [PR #95](https://github.com/twilio/terraform-provider-twilio/pull/95): push Datadog Release Metric upon deploy success. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Api**
- Add optional boolean include_soft_deleted parameter to retrieve soft deleted recordings

**Chat**
- Add `X-Twilio-Wehook-Enabled` header to `delete` method in UserChannel resource

**Numbers**
- Expose `failure_reason` in the Supporting Documents resources

**Verify**
- Add optional `metadata` parameter to "verify challenge" endpoint, so the SDK/App can attach relevant information from the device when responding to challenges.
- remove beta feature flag to list atempt api operations.
- Add `ttl` and `date_created` properties to `AccessTokens`.


[2022-02-23] Version 0.13.1
---------------------------
**Api**
- Add `uri` to `stream` resource
- Add A2P Registration Fee category (`a2p-registration-fee`) to usage records

**Verify**
- Remove outdated documentation commentary to contact sales. Product is already in public beta.


[2022-02-17] Version 0.13.0
---------------------------
**Library - Fix**
- [PR #94](https://github.com/twilio/terraform-provider-twilio/pull/94): run cluster tests before deployment. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Chore**
- [PR #93](https://github.com/twilio/terraform-provider-twilio/pull/93): update the sonar exclusions. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Api**
- Detected a bug and removed optional boolean include_soft_deleted parameter to retrieve soft deleted recordings. **(breaking change)**
- Add optional boolean include_soft_deleted parameter to retrieve soft deleted recordings.

**Numbers**
- Unrevert valid_until and sort filter params added to List Bundles resource
- Revert valid_until and sort filter params added to List Bundles resource
- Update sorting params added to List Bundles resource in the previous release

**Preview**
- Moved `web_channels` from preview to beta under `flex-api` **(breaking change)**

**Taskrouter**
- Add `ETag` as Response Header to List of Task, Reservation & Worker

**Verify**
- Add optional `metadata` to factors.


[2022-02-09] Version 0.12.0
---------------------------
**Library - Chore**
- [PR #91](https://github.com/twilio/terraform-provider-twilio/pull/91): upgrade supported language versions. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #90](https://github.com/twilio/terraform-provider-twilio/pull/90): upgrade twilio-go and downgrade indirect dependencies. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Api**
- Add `stream` resource

**Conversations**
- Fixed DELETE request to accept "sid_like" params in Address Configuration resources **(breaking change)**
- Expose Address Configuration resource for `sms` and `whatsapp`

**Fax**
- Removed deprecated Programmable Fax Create and Update methods **(breaking change)**

**Insights**
- Rename `call_state` to `call_status` and remove `whisper` in conference participant summary **(breaking change)**

**Numbers**
- Expose valid_until filters as part of provisionally-approved compliance feature on the List Bundles resource

**Supersim**
- Fix typo in Fleet resource docs
- Updated documentation for the Fleet resource indicating that fields related to commands have been deprecated and to use sms_command fields instead.
- Add support for setting and reading `ip_commands_url` and `ip_commands_method` on Fleets resource for helper libraries
- Changed `sim` property in requests to create an SMS Command made to the /SmsCommands to accept SIM UniqueNames in addition to SIDs

**Verify**
- Update list attempts API to include new filters and response fields.


[2022-01-26] Version 0.11.1
---------------------------
**Library - Chore**
- [PR #89](https://github.com/twilio/terraform-provider-twilio/pull/89): add sonarcloud coverage analysis. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Insights**
- Added new endpoint to fetch Conference Participant Summary
- Added new endpoint to fetch Conference Summary

**Messaging**
- Add government_entity parameter to brand apis

**Verify**
- Add Access Token fetch endpoint to retrieve a previously created token.
- Add Access Token payload to the Access Token creation endpoint, including a unique Sid, so it's addressable while it's TTL is valid.


[2022-01-12] Version 0.11.0
---------------------------
**Library - Feature**
- [PR #86](https://github.com/twilio/terraform-provider-twilio/pull/86): add GitHub release step during deploy. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #74](https://github.com/twilio/terraform-provider-twilio/pull/74): delve / goland debugging support. Thanks to [@loafoe](https://github.com/loafoe)!

**Library - Chore**
- [PR #83](https://github.com/twilio/terraform-provider-twilio/pull/83): remove githook dependency from make install. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Api**
- Make fixed time scheduling parameters public **(breaking change)**

**Messaging**
- Add update brand registration API

**Numbers**
- Add API endpoint for List Bundle Copies resource

**Video**
- Enable external storage for all customers


[2021-12-15] Version 0.10.0
---------------------------
**Library - Feature**
- [PR #81](https://github.com/twilio/terraform-provider-twilio/pull/81): run tests before deploying. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Api**
- Add optional boolean send_as_mms parameter to the create action of Message resource **(breaking change)**
- Change team ownership for `call` delete

**Conversations**
- Change wording for `Service Webhook Configuration` resource fields

**Insights**
- Added new APIs for updating and getting voice insights flags by accountSid.

**Media**
- Add max_duration param to MediaProcessor

**Video**
- Add `EmptyRoomTimeout` and `UnusedRoomTimeout` properties to a room; add corresponding parameters to room creation

**Voice**
- Add endpoint to delete archived Calls


[2021-12-01] Version 0.9.2
--------------------------
**Conversations**
- Add `Service Webhook Configuration` resource

**Flex**
- Adding `flex_insights_drilldown` and `flex_url` objects to Flex Configuration

**Messaging**
- Update us_app_to_person endpoints to remove beta feature flag based access

**Supersim**
- Add IP Commands resource

**Verify**
- Add optional `factor_friendly_name` parameter to the create access token endpoint.

**Video**
- Add maxParticipantDuration param to Rooms


[2021-11-17] Version 0.9.1
--------------------------
**Frontline**
- Added `is_available` to User's resource

**Messaging**
- Added GET vetting API

**Verify**
- Add `WHATSAPP` to the attempts API.
- Allow to update `config.notification_platform` from `none` to `apn` or `fcm` and viceversa for Verify Push
- Add `none` as a valid `config.notification_platform` value for Verify Push


[2021-11-03] Version 0.9.0
--------------------------
**Library - Chore**
- [PR #80](https://github.com/twilio/terraform-provider-twilio/pull/80): add ci job timeout. Thanks to [@eshanholtz](https://github.com/eshanholtz)!
- [PR #79](https://github.com/twilio/terraform-provider-twilio/pull/79): only notify once on workflow failure. Thanks to [@eshanholtz](https://github.com/eshanholtz)!
- [PR #78](https://github.com/twilio/terraform-provider-twilio/pull/78): update GitHub Actions test runner. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Api**
- Updated `media_url` property to be treated as PII

**Messaging**
- Added a new enum for brand registration status named DELETED **(breaking change)**
- Add a new K12_EDUCATION use case in us_app_to_person_usecase api transaction
- Added a new enum for brand registration status named IN_REVIEW

**Serverless**
- Add node14 as a valid Build runtime

**Verify**
- Fix typos in Verify Push Factor documentation for the `config.notification_token` parameter.
- Added `TemplateCustomSubstitutions` on verification creation
- Make `TemplateSid` parameter public for Verification resource and `DefaultTemplateSid` parameter public for Service resource. **(breaking change)**


[2021-10-19] Version 0.8.1
--------------------------
**Library - Fix**
- [PR #77](https://github.com/twilio/terraform-provider-twilio/pull/77): tests for revert computed API response fields. Thanks to [@JenniferMah](https://github.com/JenniferMah)!
- [PR #76](https://github.com/twilio/terraform-provider-twilio/pull/76): Revert feat: add computed API response fields (#73). Thanks to [@eshanholtz](https://github.com/eshanholtz)!


[2021-10-18] Version 0.8.0
--------------------------
**Library - Feature**
- [PR #73](https://github.com/twilio/terraform-provider-twilio/pull/73): add computed API response fields. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Api**
- Corrected enum values for `emergency_address_status` values in `/IncomingPhoneNumbers` response. **(breaking change)**
- Clarify `emergency_address_status` values in `/IncomingPhoneNumbers` response.

**Messaging**
- Add PUT and List brand vettings api
- Removes beta feature flag based visibility for us_app_to_person_registered and usecase field.Updates test cases to add POLITICAL usecase. **(breaking change)**
- Add brand_feedback as optional field to BrandRegistrations

**Video**
- Add `AudioOnly` to create room


[2021-10-07] Version 0.7.0
--------------------------
**Library - Fix**
- [PR #72](https://github.com/twilio/terraform-provider-twilio/pull/72): Revert "feat: add computed API response fields". Thanks to [@JenniferMah](https://github.com/JenniferMah)!
- [PR #68](https://github.com/twilio/terraform-provider-twilio/pull/68): force create resources without update on update. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #70](https://github.com/twilio/terraform-provider-twilio/pull/70): auth with accountsid/apikey/apisecret. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Feature**
- [PR #71](https://github.com/twilio/terraform-provider-twilio/pull/71): add computed API response fields. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Chore**
- [PR #69](https://github.com/twilio/terraform-provider-twilio/pull/69): fall back to api_key/api_secret for auth. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Api**
- Add `emergency_address_status` attribute to `/IncomingPhoneNumbers` response.
- Add `siprec` resource

**Conversations**
- Added attachment parameters in configuration for `NewMessage` type of push notifications

**Flex**
- Adding `flex_insights_hr` object to Flex Configuration

**Numbers**
- Add API endpoint for Bundle ReplaceItems resource
- Add API endpoint for Bundle Copies resource

**Serverless**
- Add domain_base field to Service response

**Taskrouter**
- Add `If-Match` Header based on ETag for Worker Delete **(breaking change)**
- Add `If-Match` Header based on Etag for Reservation Update
- Add `If-Match` Header based on ETag for Worker Update
- Add `If-Match` Header based on ETag for Worker Delete
- Add `ETag` as Response Header to Worker

**Trunking**
- Added `transfer_caller_id` property on Trunks.

**Verify**
- Document new pilot `whatsapp` channel.


[2021-09-22] Version 0.6.3
--------------------------
**Library - Chore**
- [PR #67](https://github.com/twilio/terraform-provider-twilio/pull/67): remove sonar scanner. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Docs**
- [PR #66](https://github.com/twilio/terraform-provider-twilio/pull/66): Adjust the API example to use a phone number variable. Thanks to [@dkundel](https://github.com/dkundel)!

**Events**
- Add segment sink

**Messaging**
- Add post_approval_required attribute in GET us_app_to_person_usecase api response
- Add Identity Status, Russell 3000, Tax Exempt Status and Should Skip SecVet fields for Brand Registrations
- Add Should Skip Secondary Vetting optional flag parameter to create Brand API


[2021-09-08] Version 0.6.2
--------------------------
**Api**
- Revert adding `siprec` resource
- Add `siprec` resource

**Messaging**
- Add 'mock' as an optional field to brand_registration api
- Add 'mock' as an optional field to us_app_to_person api
- Adds more Use Cases in us_app_to_person_usecase api transaction and updates us_app_to_person_usecase docs

**Verify**
- Verify List Templates API endpoint added.


[2021-08-25] Version 0.6.1
--------------------------
**Api**
- Add Programmabled Voice SIP Refer call transfers (`calls-transfers`) to usage records
- Add Flex Voice Usage category (`flex-usage`) to usage records

**Conversations**
- Add `Order` query parameter to Message resource read operation

**Insights**
- Added `partial` to enum processing_state_request
- Added abnormal session filter in Call Summaries

**Messaging**
- Add brand_registration_sid as an optional query param for us_app_to_person_usecase api

**Pricing**
- add trunking_numbers resource (v2)
- add trunking_country resource (v2)

**Verify**
- Changed to private beta the `TemplateSid` optional parameter on Verification creation.
- Added the optional parameter `Order` to the list Challenges endpoint to define the list order.


[2021-08-11] Version 0.6.0
--------------------------
**Library - Chore**
- [PR #65](https://github.com/twilio/terraform-provider-twilio/pull/65): integrate with sonarcloud. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Api**
- Corrected the `price`, `call_sid_to_coach`, and `uri` data types for Conference, Participant, and Recording **(breaking change)**
- Made documentation for property `time_limit` in the call api public. **(breaking change)**
- Added `domain_sid` in sip_credential_list_mapping and sip_ip_access_control_list_mapping APIs **(breaking change)**

**Insights**
- Added new endpoint to fetch Call Summaries

**Messaging**
- Add brand_type field to a2p brand_registration api
- Revert brand registration api update to add brand_type field
- Add brand_type field to a2p brand_registration api

**Taskrouter**
- Add `X-Rate-Limit-Limit`, `X-Rate-Limit-Remaining`, and `X-Rate-Limit-Config` as Response Headers to all TaskRouter endpoints

**Verify**
- Add `TemplateSid` optional parameter on Verification creation.
- Include `whatsapp` as a channel type in the verifications API.


[2021-07-28] Version 0.5.1
--------------------------
**Library - Fix**
- [PR #64](https://github.com/twilio/terraform-provider-twilio/pull/64): update underlying client construction. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Conversations**
- Expose ParticipantConversations resource

**Taskrouter**
- Adding `links` to the activity resource

**Verify**
- Added a `Version` to Verify Factors `Webhooks` to add new fields without breaking old Webhooks.


[2021-07-14] Version 0.5.0
--------------------------
**Library - Docs**
- [PR #62](https://github.com/twilio/terraform-provider-twilio/pull/62): Fix typo in terraform example for creating a phone number. Thanks to [@eflann](https://github.com/eflann)!
- [PR #61](https://github.com/twilio/terraform-provider-twilio/pull/61): add details on how to use a locally-built provider. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #60](https://github.com/twilio/terraform-provider-twilio/pull/60): Fix usage registry. Thanks to [@stern-shawn](https://github.com/stern-shawn)!
- [PR #58](https://github.com/twilio/terraform-provider-twilio/pull/58): update examples to pull from the terraform registry. Thanks to [@thinkingserious](https://github.com/thinkingserious)!

**Library - Fix**
- [PR #59](https://github.com/twilio/terraform-provider-twilio/pull/59): convert ints in path params to properly be imported. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Library - Chore**
- [PR #57](https://github.com/twilio/terraform-provider-twilio/pull/57): update readme and issue template. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Conversations**
- Changed `last_read_message_index` and `unread_messages_count` type in User Conversation's resource **(breaking change)**
- Expose UserConversations resource

**Messaging**
- Add brand_score field to brand registration responses


[2021-07-01] Version 0.4.0
--------------------------
**Library - Feature**
- [PR #56](https://github.com/twilio/terraform-provider-twilio/pull/56): add release action to sign terraform. Thanks to [@JenniferMah](https://github.com/JenniferMah)!

**Library - Chore**
- [PR #55](https://github.com/twilio/terraform-provider-twilio/pull/55): add build as a prerequisite to test. Thanks to [@thinkingserious](https://github.com/thinkingserious)!


[2021-06-30] Version 0.3.0
--------------------------
**Library - Docs**
- [PR #52](https://github.com/twilio/terraform-provider-twilio/pull/52): add Event Streams example. Thanks to [@eshanholtz](https://github.com/eshanholtz)!

**Library - Chore**
- [PR #53](https://github.com/twilio/terraform-provider-twilio/pull/53): moving twilio/common to core. Thanks to [@childish-sambino](https://github.com/childish-sambino)!

**Conversations**
- Read-only Conversation Email Binding property `binding`

**Supersim**
- Add Billing Period resource for the Super Sim Pilot
- Add List endpoint to Billing Period resource for Super Sim Pilot
- Add Fetch endpoint to Billing Period resource for Super Sim Pilot

**Taskrouter**
- Update `transcribe` & `transcription_configuration` form params in Reservation update endpoint to have private visibility **(breaking change)**
- Add `transcribe` & `transcription_configuration` form params to Reservation update endpoint


[2021-06-16] Version 0.2.0
--------------------------
**Api**
- Update `status` enum for Messages to include 'canceled'
- Update `update_status` enum for Messages to include 'canceled'

**Trusthub**
- Corrected the sid for policy sid in customer_profile_evaluation.json and trust_product_evaluation.json **(breaking change)**


[2021-06-16] Version 0.1.2
--------------------------
**Library - Fix**
- [PR #35](https://github.com/twilio/terraform-provider-twilio/pull/35): change Computed parameters to optional. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #34](https://github.com/twilio/terraform-provider-twilio/pull/34): convert sid_key ints into strings. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #27](https://github.com/twilio/terraform-provider-twilio/pull/27): fix resource parameter types. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #29](https://github.com/twilio/terraform-provider-twilio/pull/29): correct the 'create' function resource IDs. Thanks to [@childish-sambino](https://github.com/childish-sambino)!
- [PR #33](https://github.com/twilio/terraform-provider-twilio/pull/33): to_snake_case logic. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!
- [PR #31](https://github.com/twilio/terraform-provider-twilio/pull/31): resource property names. Thanks to [@shwetha-manvinkurke](https://github.com/shwetha-manvinkurke)!

**Library - Chore**
- [PR #32](https://github.com/twilio/terraform-provider-twilio/pull/32): remove empty files. Thanks to [@thinkingserious](https://github.com/thinkingserious)!
- [PR #30](https://github.com/twilio/terraform-provider-twilio/pull/30): preserve resource ordering. Thanks to [@childish-sambino](https://github.com/childish-sambino)!


[2021-06-09] Version 0.1.1
--------------------------


[2021-06-02] Version 0.1.0
--------------------------

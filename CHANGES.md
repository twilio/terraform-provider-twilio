terraform-provider-twilio changelog
====================
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

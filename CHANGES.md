terraform-provider-twilio changelog
====================
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

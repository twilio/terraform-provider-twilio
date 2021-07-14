terraform-provider-twilio changelog
====================
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

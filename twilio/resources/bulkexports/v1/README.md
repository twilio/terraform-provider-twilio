
## twilio_bulkexports_exports_jobs_v1

### Parameters

Name | Type | Requirement | Description
--- | --- | --- | ---
**resource_type** | string | **Required** | The type of communication – Messages or Calls, Conferences, and Participants
**end_day** | string | **Required** | The end day for the custom export specified as a string in the format of yyyy-mm-dd. End day is inclusive and must be 2 days earlier than the current UTC day.
**friendly_name** | string | **Required** | The friendly name specified when creating the job
**start_day** | string | **Required** | The start day for the custom export specified as a string in the format of yyyy-mm-dd
**email** | string | Optional | The optional email to send the completion notification to. You can set both webhook, and email, or one or the other. If you set neither, the job will run but you will have to query to determine your job&#39;s status.
**webhook_method** | string | Optional | This is the method used to call the webhook on completion of the job. If this is supplied, &#x60;WebhookUrl&#x60; must also be supplied.
**webhook_url** | string | Optional | The optional webhook url called on completion of the job. If this is supplied, &#x60;WebhookMethod&#x60; must also be supplied. If you set neither webhook nor email, you will have to check your job&#39;s status manually.
**job_sid** | string | *Computed* | The unique string that that we created to identify the Bulk Export job
**details** | string | *Computed* | The details of a job state which is an object that contains a &#x60;status&#x60; string, a day count integer, and list of days in the job
**estimated_completion_time** | string | *Computed* | this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position
**job_queue_position** | string | *Computed* | This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease


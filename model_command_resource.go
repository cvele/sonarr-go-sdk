/*
Sonarr

Sonarr API docs

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// CommandResource struct for CommandResource
type CommandResource struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	CommandName NullableString `json:"commandName,omitempty"`
	Message NullableString `json:"message,omitempty"`
	Body *Command `json:"body,omitempty"`
	Priority *CommandPriority `json:"priority,omitempty"`
	Status *CommandStatus `json:"status,omitempty"`
	Queued *time.Time `json:"queued,omitempty"`
	Started NullableTime `json:"started,omitempty"`
	Ended NullableTime `json:"ended,omitempty"`
	Duration *TimeSpan `json:"duration,omitempty"`
	Exception NullableString `json:"exception,omitempty"`
	Trigger *CommandTrigger `json:"trigger,omitempty"`
	ClientUserAgent NullableString `json:"clientUserAgent,omitempty"`
	StateChangeTime NullableTime `json:"stateChangeTime,omitempty"`
	SendUpdatesToClient *bool `json:"sendUpdatesToClient,omitempty"`
	UpdateScheduledTask *bool `json:"updateScheduledTask,omitempty"`
	LastExecutionTime NullableTime `json:"lastExecutionTime,omitempty"`
}

// NewCommandResource instantiates a new CommandResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandResource() *CommandResource {
	this := CommandResource{}
	return &this
}

// NewCommandResourceWithDefaults instantiates a new CommandResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandResourceWithDefaults() *CommandResource {
	this := CommandResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommandResource) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommandResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CommandResource) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CommandResource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CommandResource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CommandResource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CommandResource) UnsetName() {
	o.Name.Unset()
}

// GetCommandName returns the CommandName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetCommandName() string {
	if o == nil || o.CommandName.Get() == nil {
		var ret string
		return ret
	}
	return *o.CommandName.Get()
}

// GetCommandNameOk returns a tuple with the CommandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetCommandNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommandName.Get(), o.CommandName.IsSet()
}

// HasCommandName returns a boolean if a field has been set.
func (o *CommandResource) HasCommandName() bool {
	if o != nil && o.CommandName.IsSet() {
		return true
	}

	return false
}

// SetCommandName gets a reference to the given NullableString and assigns it to the CommandName field.
func (o *CommandResource) SetCommandName(v string) {
	o.CommandName.Set(&v)
}
// SetCommandNameNil sets the value for CommandName to be an explicit nil
func (o *CommandResource) SetCommandNameNil() {
	o.CommandName.Set(nil)
}

// UnsetCommandName ensures that no value is present for CommandName, not even an explicit nil
func (o *CommandResource) UnsetCommandName() {
	o.CommandName.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *CommandResource) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *CommandResource) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *CommandResource) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *CommandResource) UnsetMessage() {
	o.Message.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CommandResource) GetBody() Command {
	if o == nil || o.Body == nil {
		var ret Command
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetBodyOk() (*Command, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CommandResource) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given Command and assigns it to the Body field.
func (o *CommandResource) SetBody(v Command) {
	o.Body = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CommandResource) GetPriority() CommandPriority {
	if o == nil || o.Priority == nil {
		var ret CommandPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetPriorityOk() (*CommandPriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CommandResource) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given CommandPriority and assigns it to the Priority field.
func (o *CommandResource) SetPriority(v CommandPriority) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CommandResource) GetStatus() CommandStatus {
	if o == nil || o.Status == nil {
		var ret CommandStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetStatusOk() (*CommandStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CommandResource) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CommandStatus and assigns it to the Status field.
func (o *CommandResource) SetStatus(v CommandStatus) {
	o.Status = &v
}

// GetQueued returns the Queued field value if set, zero value otherwise.
func (o *CommandResource) GetQueued() time.Time {
	if o == nil || o.Queued == nil {
		var ret time.Time
		return ret
	}
	return *o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetQueuedOk() (*time.Time, bool) {
	if o == nil || o.Queued == nil {
		return nil, false
	}
	return o.Queued, true
}

// HasQueued returns a boolean if a field has been set.
func (o *CommandResource) HasQueued() bool {
	if o != nil && o.Queued != nil {
		return true
	}

	return false
}

// SetQueued gets a reference to the given time.Time and assigns it to the Queued field.
func (o *CommandResource) SetQueued(v time.Time) {
	o.Queued = &v
}

// GetStarted returns the Started field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetStarted() time.Time {
	if o == nil || o.Started.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Started.Get()
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetStartedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Started.Get(), o.Started.IsSet()
}

// HasStarted returns a boolean if a field has been set.
func (o *CommandResource) HasStarted() bool {
	if o != nil && o.Started.IsSet() {
		return true
	}

	return false
}

// SetStarted gets a reference to the given NullableTime and assigns it to the Started field.
func (o *CommandResource) SetStarted(v time.Time) {
	o.Started.Set(&v)
}
// SetStartedNil sets the value for Started to be an explicit nil
func (o *CommandResource) SetStartedNil() {
	o.Started.Set(nil)
}

// UnsetStarted ensures that no value is present for Started, not even an explicit nil
func (o *CommandResource) UnsetStarted() {
	o.Started.Unset()
}

// GetEnded returns the Ended field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetEnded() time.Time {
	if o == nil || o.Ended.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Ended.Get()
}

// GetEndedOk returns a tuple with the Ended field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetEndedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ended.Get(), o.Ended.IsSet()
}

// HasEnded returns a boolean if a field has been set.
func (o *CommandResource) HasEnded() bool {
	if o != nil && o.Ended.IsSet() {
		return true
	}

	return false
}

// SetEnded gets a reference to the given NullableTime and assigns it to the Ended field.
func (o *CommandResource) SetEnded(v time.Time) {
	o.Ended.Set(&v)
}
// SetEndedNil sets the value for Ended to be an explicit nil
func (o *CommandResource) SetEndedNil() {
	o.Ended.Set(nil)
}

// UnsetEnded ensures that no value is present for Ended, not even an explicit nil
func (o *CommandResource) UnsetEnded() {
	o.Ended.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *CommandResource) GetDuration() TimeSpan {
	if o == nil || o.Duration == nil {
		var ret TimeSpan
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetDurationOk() (*TimeSpan, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *CommandResource) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given TimeSpan and assigns it to the Duration field.
func (o *CommandResource) SetDuration(v TimeSpan) {
	o.Duration = &v
}

// GetException returns the Exception field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetException() string {
	if o == nil || o.Exception.Get() == nil {
		var ret string
		return ret
	}
	return *o.Exception.Get()
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetExceptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exception.Get(), o.Exception.IsSet()
}

// HasException returns a boolean if a field has been set.
func (o *CommandResource) HasException() bool {
	if o != nil && o.Exception.IsSet() {
		return true
	}

	return false
}

// SetException gets a reference to the given NullableString and assigns it to the Exception field.
func (o *CommandResource) SetException(v string) {
	o.Exception.Set(&v)
}
// SetExceptionNil sets the value for Exception to be an explicit nil
func (o *CommandResource) SetExceptionNil() {
	o.Exception.Set(nil)
}

// UnsetException ensures that no value is present for Exception, not even an explicit nil
func (o *CommandResource) UnsetException() {
	o.Exception.Unset()
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *CommandResource) GetTrigger() CommandTrigger {
	if o == nil || o.Trigger == nil {
		var ret CommandTrigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetTriggerOk() (*CommandTrigger, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *CommandResource) HasTrigger() bool {
	if o != nil && o.Trigger != nil {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given CommandTrigger and assigns it to the Trigger field.
func (o *CommandResource) SetTrigger(v CommandTrigger) {
	o.Trigger = &v
}

// GetClientUserAgent returns the ClientUserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetClientUserAgent() string {
	if o == nil || o.ClientUserAgent.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientUserAgent.Get()
}

// GetClientUserAgentOk returns a tuple with the ClientUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetClientUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientUserAgent.Get(), o.ClientUserAgent.IsSet()
}

// HasClientUserAgent returns a boolean if a field has been set.
func (o *CommandResource) HasClientUserAgent() bool {
	if o != nil && o.ClientUserAgent.IsSet() {
		return true
	}

	return false
}

// SetClientUserAgent gets a reference to the given NullableString and assigns it to the ClientUserAgent field.
func (o *CommandResource) SetClientUserAgent(v string) {
	o.ClientUserAgent.Set(&v)
}
// SetClientUserAgentNil sets the value for ClientUserAgent to be an explicit nil
func (o *CommandResource) SetClientUserAgentNil() {
	o.ClientUserAgent.Set(nil)
}

// UnsetClientUserAgent ensures that no value is present for ClientUserAgent, not even an explicit nil
func (o *CommandResource) UnsetClientUserAgent() {
	o.ClientUserAgent.Unset()
}

// GetStateChangeTime returns the StateChangeTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetStateChangeTime() time.Time {
	if o == nil || o.StateChangeTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StateChangeTime.Get()
}

// GetStateChangeTimeOk returns a tuple with the StateChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetStateChangeTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StateChangeTime.Get(), o.StateChangeTime.IsSet()
}

// HasStateChangeTime returns a boolean if a field has been set.
func (o *CommandResource) HasStateChangeTime() bool {
	if o != nil && o.StateChangeTime.IsSet() {
		return true
	}

	return false
}

// SetStateChangeTime gets a reference to the given NullableTime and assigns it to the StateChangeTime field.
func (o *CommandResource) SetStateChangeTime(v time.Time) {
	o.StateChangeTime.Set(&v)
}
// SetStateChangeTimeNil sets the value for StateChangeTime to be an explicit nil
func (o *CommandResource) SetStateChangeTimeNil() {
	o.StateChangeTime.Set(nil)
}

// UnsetStateChangeTime ensures that no value is present for StateChangeTime, not even an explicit nil
func (o *CommandResource) UnsetStateChangeTime() {
	o.StateChangeTime.Unset()
}

// GetSendUpdatesToClient returns the SendUpdatesToClient field value if set, zero value otherwise.
func (o *CommandResource) GetSendUpdatesToClient() bool {
	if o == nil || o.SendUpdatesToClient == nil {
		var ret bool
		return ret
	}
	return *o.SendUpdatesToClient
}

// GetSendUpdatesToClientOk returns a tuple with the SendUpdatesToClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetSendUpdatesToClientOk() (*bool, bool) {
	if o == nil || o.SendUpdatesToClient == nil {
		return nil, false
	}
	return o.SendUpdatesToClient, true
}

// HasSendUpdatesToClient returns a boolean if a field has been set.
func (o *CommandResource) HasSendUpdatesToClient() bool {
	if o != nil && o.SendUpdatesToClient != nil {
		return true
	}

	return false
}

// SetSendUpdatesToClient gets a reference to the given bool and assigns it to the SendUpdatesToClient field.
func (o *CommandResource) SetSendUpdatesToClient(v bool) {
	o.SendUpdatesToClient = &v
}

// GetUpdateScheduledTask returns the UpdateScheduledTask field value if set, zero value otherwise.
func (o *CommandResource) GetUpdateScheduledTask() bool {
	if o == nil || o.UpdateScheduledTask == nil {
		var ret bool
		return ret
	}
	return *o.UpdateScheduledTask
}

// GetUpdateScheduledTaskOk returns a tuple with the UpdateScheduledTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResource) GetUpdateScheduledTaskOk() (*bool, bool) {
	if o == nil || o.UpdateScheduledTask == nil {
		return nil, false
	}
	return o.UpdateScheduledTask, true
}

// HasUpdateScheduledTask returns a boolean if a field has been set.
func (o *CommandResource) HasUpdateScheduledTask() bool {
	if o != nil && o.UpdateScheduledTask != nil {
		return true
	}

	return false
}

// SetUpdateScheduledTask gets a reference to the given bool and assigns it to the UpdateScheduledTask field.
func (o *CommandResource) SetUpdateScheduledTask(v bool) {
	o.UpdateScheduledTask = &v
}

// GetLastExecutionTime returns the LastExecutionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommandResource) GetLastExecutionTime() time.Time {
	if o == nil || o.LastExecutionTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastExecutionTime.Get()
}

// GetLastExecutionTimeOk returns a tuple with the LastExecutionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommandResource) GetLastExecutionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutionTime.Get(), o.LastExecutionTime.IsSet()
}

// HasLastExecutionTime returns a boolean if a field has been set.
func (o *CommandResource) HasLastExecutionTime() bool {
	if o != nil && o.LastExecutionTime.IsSet() {
		return true
	}

	return false
}

// SetLastExecutionTime gets a reference to the given NullableTime and assigns it to the LastExecutionTime field.
func (o *CommandResource) SetLastExecutionTime(v time.Time) {
	o.LastExecutionTime.Set(&v)
}
// SetLastExecutionTimeNil sets the value for LastExecutionTime to be an explicit nil
func (o *CommandResource) SetLastExecutionTimeNil() {
	o.LastExecutionTime.Set(nil)
}

// UnsetLastExecutionTime ensures that no value is present for LastExecutionTime, not even an explicit nil
func (o *CommandResource) UnsetLastExecutionTime() {
	o.LastExecutionTime.Unset()
}

func (o CommandResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.CommandName.IsSet() {
		toSerialize["commandName"] = o.CommandName.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Queued != nil {
		toSerialize["queued"] = o.Queued
	}
	if o.Started.IsSet() {
		toSerialize["started"] = o.Started.Get()
	}
	if o.Ended.IsSet() {
		toSerialize["ended"] = o.Ended.Get()
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Exception.IsSet() {
		toSerialize["exception"] = o.Exception.Get()
	}
	if o.Trigger != nil {
		toSerialize["trigger"] = o.Trigger
	}
	if o.ClientUserAgent.IsSet() {
		toSerialize["clientUserAgent"] = o.ClientUserAgent.Get()
	}
	if o.StateChangeTime.IsSet() {
		toSerialize["stateChangeTime"] = o.StateChangeTime.Get()
	}
	if o.SendUpdatesToClient != nil {
		toSerialize["sendUpdatesToClient"] = o.SendUpdatesToClient
	}
	if o.UpdateScheduledTask != nil {
		toSerialize["updateScheduledTask"] = o.UpdateScheduledTask
	}
	if o.LastExecutionTime.IsSet() {
		toSerialize["lastExecutionTime"] = o.LastExecutionTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCommandResource struct {
	value *CommandResource
	isSet bool
}

func (v NullableCommandResource) Get() *CommandResource {
	return v.value
}

func (v *NullableCommandResource) Set(val *CommandResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandResource(val *CommandResource) *NullableCommandResource {
	return &NullableCommandResource{value: val, isSet: true}
}

func (v NullableCommandResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


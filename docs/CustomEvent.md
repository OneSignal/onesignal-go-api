# CustomEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The identifier or name of the event. Maximum 128 characters. | 
**ExternalId** | Pointer to **NullableString** | The external ID of the user targeted for the event. Either the user&#39;s External ID or OneSignal ID is required. | [optional] 
**OnesignalId** | Pointer to **NullableString** | The OneSignal ID of the user targeted for the event. Either the user&#39;s External ID or OneSignal ID is required. | [optional] 
**Timestamp** | Pointer to **time.Time** | Time the event occurred as an ISO8601 formatted string. Defaults to now if not included or past date provided. | [optional] 
**Payload** | Pointer to **map[string]interface{}** | Properties or data related to the event, like {\&quot;geography\&quot;: \&quot;USA\&quot;} | [optional] 

## Methods

### NewCustomEvent

`func NewCustomEvent(name string, ) *CustomEvent`

NewCustomEvent instantiates a new CustomEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEventWithDefaults

`func NewCustomEventWithDefaults() *CustomEvent`

NewCustomEventWithDefaults instantiates a new CustomEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEvent) SetName(v string)`

SetName sets Name field to given value.


### GetExternalId

`func (o *CustomEvent) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CustomEvent) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CustomEvent) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CustomEvent) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *CustomEvent) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CustomEvent) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetOnesignalId

`func (o *CustomEvent) GetOnesignalId() string`

GetOnesignalId returns the OnesignalId field if non-nil, zero value otherwise.

### GetOnesignalIdOk

`func (o *CustomEvent) GetOnesignalIdOk() (*string, bool)`

GetOnesignalIdOk returns a tuple with the OnesignalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnesignalId

`func (o *CustomEvent) SetOnesignalId(v string)`

SetOnesignalId sets OnesignalId field to given value.

### HasOnesignalId

`func (o *CustomEvent) HasOnesignalId() bool`

HasOnesignalId returns a boolean if a field has been set.

### SetOnesignalIdNil

`func (o *CustomEvent) SetOnesignalIdNil(b bool)`

 SetOnesignalIdNil sets the value for OnesignalId to be an explicit nil

### UnsetOnesignalId
`func (o *CustomEvent) UnsetOnesignalId()`

UnsetOnesignalId ensures that no value is present for OnesignalId, not even an explicit nil
### GetTimestamp

`func (o *CustomEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CustomEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CustomEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CustomEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetPayload

`func (o *CustomEvent) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CustomEvent) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CustomEvent) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *CustomEvent) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



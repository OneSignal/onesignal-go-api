# AuditLogActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email address of the actor. Absent if unavailable. | [optional] 
**Id** | Pointer to **string** | UUID of the actor. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional actor-specific data. | [optional] 
**Name** | Pointer to **string** | Display name of the actor. Absent if unavailable. | [optional] 
**Type** | Pointer to **string** | Actor type (e.g. member, api_key, system). | [optional] 

## Methods

### NewAuditLogActor

`func NewAuditLogActor() *AuditLogActor`

NewAuditLogActor instantiates a new AuditLogActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogActorWithDefaults

`func NewAuditLogActorWithDefaults() *AuditLogActor`

NewAuditLogActorWithDefaults instantiates a new AuditLogActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuditLogActor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuditLogActor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuditLogActor) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuditLogActor) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *AuditLogActor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogActor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogActor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogActor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AuditLogActor) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AuditLogActor) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AuditLogActor) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AuditLogActor) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *AuditLogActor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLogActor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLogActor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditLogActor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AuditLogActor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditLogActor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditLogActor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditLogActor) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)



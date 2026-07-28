# AuditLogTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UUID of the resource. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional resource-specific data. | [optional] 
**Name** | Pointer to **string** | Display name of the resource. Absent if unavailable. | [optional] 
**Type** | Pointer to **string** | Resource type (e.g. notification, segment, journey, app). | [optional] 

## Methods

### NewAuditLogTarget

`func NewAuditLogTarget() *AuditLogTarget`

NewAuditLogTarget instantiates a new AuditLogTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogTargetWithDefaults

`func NewAuditLogTargetWithDefaults() *AuditLogTarget`

NewAuditLogTargetWithDefaults instantiates a new AuditLogTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditLogTarget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditLogTarget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditLogTarget) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditLogTarget) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *AuditLogTarget) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AuditLogTarget) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AuditLogTarget) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AuditLogTarget) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *AuditLogTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditLogTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditLogTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditLogTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AuditLogTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditLogTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditLogTarget) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditLogTarget) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)



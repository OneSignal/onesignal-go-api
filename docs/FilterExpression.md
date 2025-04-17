# FilterExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | Required. Name of the field to use as the first operand in the filter expression. | [optional] 
**Key** | Pointer to **string** | If &#x60;field&#x60; is &#x60;tag&#x60;, this field is *required* to specify &#x60;key&#x60; inside the tags. | [optional] 
**Value** | Pointer to **string** | Constant value to use as the second operand in the filter expression. This value is *required* when the relation operator is a binary operator. | [optional] 
**HoursAgo** | Pointer to **string** | If &#x60;field&#x60; is session-related, this is *required* to specify the number of hours before or after the user&#39;s session. | [optional] 
**Radius** | Pointer to **float32** | If &#x60;field&#x60; is &#x60;location&#x60;, this will specify the radius in meters from a provided location point. Use with &#x60;lat&#x60; and &#x60;long&#x60;. | [optional] 
**Lat** | Pointer to **float32** | If &#x60;field&#x60; is &#x60;location&#x60;, this is *required* to specify the user&#39;s latitude. | [optional] 
**Long** | Pointer to **float32** | If &#x60;field&#x60; is &#x60;location&#x60;, this is *required* to specify the user&#39;s longitude. | [optional] 
**Relation** | Pointer to **string** | Required. Operator of a filter expression. | [optional] 
**Operator** | Pointer to **string** | Strictly, this must be either &#x60;\&quot;OR\&quot;&#x60;, or &#x60;\&quot;AND\&quot;&#x60;.  It can be used to compose Filters as part of a Filters object. | [optional] 

## Methods

### NewFilterExpression

`func NewFilterExpression() *FilterExpression`

NewFilterExpression instantiates a new FilterExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterExpressionWithDefaults

`func NewFilterExpressionWithDefaults() *FilterExpression`

NewFilterExpressionWithDefaults instantiates a new FilterExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FilterExpression) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FilterExpression) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FilterExpression) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FilterExpression) HasField() bool`

HasField returns a boolean if a field has been set.

### GetKey

`func (o *FilterExpression) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FilterExpression) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FilterExpression) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FilterExpression) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *FilterExpression) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FilterExpression) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FilterExpression) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FilterExpression) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHoursAgo

`func (o *FilterExpression) GetHoursAgo() string`

GetHoursAgo returns the HoursAgo field if non-nil, zero value otherwise.

### GetHoursAgoOk

`func (o *FilterExpression) GetHoursAgoOk() (*string, bool)`

GetHoursAgoOk returns a tuple with the HoursAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursAgo

`func (o *FilterExpression) SetHoursAgo(v string)`

SetHoursAgo sets HoursAgo field to given value.

### HasHoursAgo

`func (o *FilterExpression) HasHoursAgo() bool`

HasHoursAgo returns a boolean if a field has been set.

### GetRadius

`func (o *FilterExpression) GetRadius() float32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *FilterExpression) GetRadiusOk() (*float32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *FilterExpression) SetRadius(v float32)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *FilterExpression) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetLat

`func (o *FilterExpression) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *FilterExpression) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *FilterExpression) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *FilterExpression) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLong

`func (o *FilterExpression) GetLong() float32`

GetLong returns the Long field if non-nil, zero value otherwise.

### GetLongOk

`func (o *FilterExpression) GetLongOk() (*float32, bool)`

GetLongOk returns a tuple with the Long field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLong

`func (o *FilterExpression) SetLong(v float32)`

SetLong sets Long field to given value.

### HasLong

`func (o *FilterExpression) HasLong() bool`

HasLong returns a boolean if a field has been set.

### GetRelation

`func (o *FilterExpression) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *FilterExpression) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *FilterExpression) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *FilterExpression) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetOperator

`func (o *FilterExpression) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FilterExpression) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FilterExpression) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *FilterExpression) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



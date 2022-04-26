# FilterExpressions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Name of the field to use as the first operand in the filter expression. | 
**Key** | Pointer to **string** | If &#x60;field&#x60; is &#x60;tag&#x60;, this field is *required* to specify &#x60;key&#x60; inside the tags. | [optional] 
**Value** | Pointer to **string** | Constant value to use as the second operand in the filter expression.  This value is *required* when the relation operator is a binary operator. | [optional] 
**Relation** | **string** | Operator of a filter expression. | 
**Operator** | Pointer to **string** | Strictly, this must be either &#x60;\&quot;OR\&quot;&#x60;, or &#x60;\&quot;AND\&quot;&#x60;.  It can be used to compose Filters as part of a Filters object. | [optional] 

## Methods

### NewFilterExpressions

`func NewFilterExpressions(field string, relation string, ) *FilterExpressions`

NewFilterExpressions instantiates a new FilterExpressions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterExpressionsWithDefaults

`func NewFilterExpressionsWithDefaults() *FilterExpressions`

NewFilterExpressionsWithDefaults instantiates a new FilterExpressions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FilterExpressions) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FilterExpressions) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FilterExpressions) SetField(v string)`

SetField sets Field field to given value.


### GetKey

`func (o *FilterExpressions) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FilterExpressions) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FilterExpressions) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FilterExpressions) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *FilterExpressions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FilterExpressions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FilterExpressions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FilterExpressions) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetRelation

`func (o *FilterExpressions) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *FilterExpressions) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *FilterExpressions) SetRelation(v string)`

SetRelation sets Relation field to given value.


### GetOperator

`func (o *FilterExpressions) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *FilterExpressions) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *FilterExpressions) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *FilterExpressions) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



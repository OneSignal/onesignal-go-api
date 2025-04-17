# PropertiesDeltas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionTime** | Pointer to **int32** |  | [optional] 
**SessionCount** | Pointer to **int32** |  | [optional] 
**Purchases** | Pointer to [**[]Purchase**](Purchase.md) |  | [optional] 

## Methods

### NewPropertiesDeltas

`func NewPropertiesDeltas() *PropertiesDeltas`

NewPropertiesDeltas instantiates a new PropertiesDeltas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertiesDeltasWithDefaults

`func NewPropertiesDeltasWithDefaults() *PropertiesDeltas`

NewPropertiesDeltasWithDefaults instantiates a new PropertiesDeltas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionTime

`func (o *PropertiesDeltas) GetSessionTime() int32`

GetSessionTime returns the SessionTime field if non-nil, zero value otherwise.

### GetSessionTimeOk

`func (o *PropertiesDeltas) GetSessionTimeOk() (*int32, bool)`

GetSessionTimeOk returns a tuple with the SessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTime

`func (o *PropertiesDeltas) SetSessionTime(v int32)`

SetSessionTime sets SessionTime field to given value.

### HasSessionTime

`func (o *PropertiesDeltas) HasSessionTime() bool`

HasSessionTime returns a boolean if a field has been set.

### GetSessionCount

`func (o *PropertiesDeltas) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *PropertiesDeltas) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *PropertiesDeltas) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *PropertiesDeltas) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetPurchases

`func (o *PropertiesDeltas) GetPurchases() []Purchase`

GetPurchases returns the Purchases field if non-nil, zero value otherwise.

### GetPurchasesOk

`func (o *PropertiesDeltas) GetPurchasesOk() (*[]Purchase, bool)`

GetPurchasesOk returns a tuple with the Purchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchases

`func (o *PropertiesDeltas) SetPurchases(v []Purchase)`

SetPurchases sets Purchases field to given value.

### HasPurchases

`func (o *PropertiesDeltas) HasPurchases() bool`

HasPurchases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



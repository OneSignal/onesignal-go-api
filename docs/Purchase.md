# Purchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sku** | **string** | The unique identifier of the purchased item. | 
**Amount** | **string** | The amount, in USD, spent purchasing the item. | 
**Iso** | **string** | The 3-letter ISO 4217 currency code. Required for correct storage and conversion of amount. | 
**Count** | Pointer to **float32** |  | [optional] 

## Methods

### NewPurchase

`func NewPurchase(sku string, amount string, iso string, ) *Purchase`

NewPurchase instantiates a new Purchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseWithDefaults

`func NewPurchaseWithDefaults() *Purchase`

NewPurchaseWithDefaults instantiates a new Purchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSku

`func (o *Purchase) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Purchase) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Purchase) SetSku(v string)`

SetSku sets Sku field to given value.


### GetAmount

`func (o *Purchase) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Purchase) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Purchase) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIso

`func (o *Purchase) GetIso() string`

GetIso returns the Iso field if non-nil, zero value otherwise.

### GetIsoOk

`func (o *Purchase) GetIsoOk() (*string, bool)`

GetIsoOk returns a tuple with the Iso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso

`func (o *Purchase) SetIso(v string)`

SetIso sets Iso field to given value.


### GetCount

`func (o *Purchase) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Purchase) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Purchase) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Purchase) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



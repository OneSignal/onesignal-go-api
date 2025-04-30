# PropertiesObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**TimezoneId** | Pointer to **string** |  | [optional] 
**Lat** | Pointer to **float32** |  | [optional] 
**Long** | Pointer to **float32** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**FirstActive** | Pointer to **int32** |  | [optional] 
**LastActive** | Pointer to **int32** |  | [optional] 
**AmountSpent** | Pointer to **float32** |  | [optional] 
**Purchases** | Pointer to [**[]Purchase**](Purchase.md) |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 

## Methods

### NewPropertiesObject

`func NewPropertiesObject() *PropertiesObject`

NewPropertiesObject instantiates a new PropertiesObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertiesObjectWithDefaults

`func NewPropertiesObjectWithDefaults() *PropertiesObject`

NewPropertiesObjectWithDefaults instantiates a new PropertiesObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *PropertiesObject) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PropertiesObject) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PropertiesObject) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *PropertiesObject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLanguage

`func (o *PropertiesObject) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PropertiesObject) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PropertiesObject) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PropertiesObject) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetTimezoneId

`func (o *PropertiesObject) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *PropertiesObject) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *PropertiesObject) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *PropertiesObject) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetLat

`func (o *PropertiesObject) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *PropertiesObject) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *PropertiesObject) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *PropertiesObject) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLong

`func (o *PropertiesObject) GetLong() float32`

GetLong returns the Long field if non-nil, zero value otherwise.

### GetLongOk

`func (o *PropertiesObject) GetLongOk() (*float32, bool)`

GetLongOk returns a tuple with the Long field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLong

`func (o *PropertiesObject) SetLong(v float32)`

SetLong sets Long field to given value.

### HasLong

`func (o *PropertiesObject) HasLong() bool`

HasLong returns a boolean if a field has been set.

### GetCountry

`func (o *PropertiesObject) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PropertiesObject) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PropertiesObject) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PropertiesObject) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetFirstActive

`func (o *PropertiesObject) GetFirstActive() int32`

GetFirstActive returns the FirstActive field if non-nil, zero value otherwise.

### GetFirstActiveOk

`func (o *PropertiesObject) GetFirstActiveOk() (*int32, bool)`

GetFirstActiveOk returns a tuple with the FirstActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstActive

`func (o *PropertiesObject) SetFirstActive(v int32)`

SetFirstActive sets FirstActive field to given value.

### HasFirstActive

`func (o *PropertiesObject) HasFirstActive() bool`

HasFirstActive returns a boolean if a field has been set.

### GetLastActive

`func (o *PropertiesObject) GetLastActive() int32`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *PropertiesObject) GetLastActiveOk() (*int32, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *PropertiesObject) SetLastActive(v int32)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *PropertiesObject) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetAmountSpent

`func (o *PropertiesObject) GetAmountSpent() float32`

GetAmountSpent returns the AmountSpent field if non-nil, zero value otherwise.

### GetAmountSpentOk

`func (o *PropertiesObject) GetAmountSpentOk() (*float32, bool)`

GetAmountSpentOk returns a tuple with the AmountSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSpent

`func (o *PropertiesObject) SetAmountSpent(v float32)`

SetAmountSpent sets AmountSpent field to given value.

### HasAmountSpent

`func (o *PropertiesObject) HasAmountSpent() bool`

HasAmountSpent returns a boolean if a field has been set.

### GetPurchases

`func (o *PropertiesObject) GetPurchases() []Purchase`

GetPurchases returns the Purchases field if non-nil, zero value otherwise.

### GetPurchasesOk

`func (o *PropertiesObject) GetPurchasesOk() (*[]Purchase, bool)`

GetPurchasesOk returns a tuple with the Purchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchases

`func (o *PropertiesObject) SetPurchases(v []Purchase)`

SetPurchases sets Purchases field to given value.

### HasPurchases

`func (o *PropertiesObject) HasPurchases() bool`

HasPurchases returns a boolean if a field has been set.

### GetIp

`func (o *PropertiesObject) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PropertiesObject) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PropertiesObject) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *PropertiesObject) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



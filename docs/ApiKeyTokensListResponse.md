# ApiKeyTokensListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | Pointer to [**[]ApiKeyToken**](ApiKeyToken.md) |  | [optional] 

## Methods

### NewApiKeyTokensListResponse

`func NewApiKeyTokensListResponse() *ApiKeyTokensListResponse`

NewApiKeyTokensListResponse instantiates a new ApiKeyTokensListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyTokensListResponseWithDefaults

`func NewApiKeyTokensListResponseWithDefaults() *ApiKeyTokensListResponse`

NewApiKeyTokensListResponseWithDefaults instantiates a new ApiKeyTokensListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *ApiKeyTokensListResponse) GetTokens() []ApiKeyToken`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ApiKeyTokensListResponse) GetTokensOk() (*[]ApiKeyToken, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ApiKeyTokensListResponse) SetTokens(v []ApiKeyToken)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ApiKeyTokensListResponse) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



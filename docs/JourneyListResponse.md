# JourneyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Journeys** | Pointer to [**[]JourneyListItem**](JourneyListItem.md) | Journeys ordered by creation time, newest first. | [optional] 
**HasMore** | Pointer to **bool** | true if more journeys exist beyond this page. | [optional] 
**NextCursor** | Pointer to **string** | Cursor for the next page. Present only when has_more is true. | [optional] 

## Methods

### NewJourneyListResponse

`func NewJourneyListResponse() *JourneyListResponse`

NewJourneyListResponse instantiates a new JourneyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyListResponseWithDefaults

`func NewJourneyListResponseWithDefaults() *JourneyListResponse`

NewJourneyListResponseWithDefaults instantiates a new JourneyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJourneys

`func (o *JourneyListResponse) GetJourneys() []JourneyListItem`

GetJourneys returns the Journeys field if non-nil, zero value otherwise.

### GetJourneysOk

`func (o *JourneyListResponse) GetJourneysOk() (*[]JourneyListItem, bool)`

GetJourneysOk returns a tuple with the Journeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJourneys

`func (o *JourneyListResponse) SetJourneys(v []JourneyListItem)`

SetJourneys sets Journeys field to given value.

### HasJourneys

`func (o *JourneyListResponse) HasJourneys() bool`

HasJourneys returns a boolean if a field has been set.

### GetHasMore

`func (o *JourneyListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *JourneyListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *JourneyListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *JourneyListResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNextCursor

`func (o *JourneyListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *JourneyListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *JourneyListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *JourneyListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)



# JourneyAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | Audience kind. Selects which other fields apply. | 
**IncludedSegmentIds** | Pointer to **[]string** | segment audiences: Segment UUIDs whose users enter the journey. | [optional] 
**ExcludedSegmentIds** | Pointer to **[]string** | segment audiences: Segment UUIDs whose users are excluded. | [optional] 
**FutureAdditionsOnly** | Pointer to **NullableBool** | segment audiences: when true, only users who newly match the segment after activation enter the journey. Defaults to false. | [optional] 
**Name** | Pointer to **string** | event_trigger audiences: event name that triggers entry, up to 255 characters. | [optional] 
**Attributes** | Pointer to [**[][]JourneyEventAttribute**]([]JourneyEventAttribute.md) | Event attribute matchers, as a list of condition groups. Send a single group whose conditions are AND&#39;d together. More than one group is rejected. | [optional] 

## Methods

### NewJourneyAudience

`func NewJourneyAudience(kind string, ) *JourneyAudience`

NewJourneyAudience instantiates a new JourneyAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyAudienceWithDefaults

`func NewJourneyAudienceWithDefaults() *JourneyAudience`

NewJourneyAudienceWithDefaults instantiates a new JourneyAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *JourneyAudience) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *JourneyAudience) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *JourneyAudience) SetKind(v string)`

SetKind sets Kind field to given value.


### GetIncludedSegmentIds

`func (o *JourneyAudience) GetIncludedSegmentIds() []string`

GetIncludedSegmentIds returns the IncludedSegmentIds field if non-nil, zero value otherwise.

### GetIncludedSegmentIdsOk

`func (o *JourneyAudience) GetIncludedSegmentIdsOk() (*[]string, bool)`

GetIncludedSegmentIdsOk returns a tuple with the IncludedSegmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSegmentIds

`func (o *JourneyAudience) SetIncludedSegmentIds(v []string)`

SetIncludedSegmentIds sets IncludedSegmentIds field to given value.

### HasIncludedSegmentIds

`func (o *JourneyAudience) HasIncludedSegmentIds() bool`

HasIncludedSegmentIds returns a boolean if a field has been set.

### GetExcludedSegmentIds

`func (o *JourneyAudience) GetExcludedSegmentIds() []string`

GetExcludedSegmentIds returns the ExcludedSegmentIds field if non-nil, zero value otherwise.

### GetExcludedSegmentIdsOk

`func (o *JourneyAudience) GetExcludedSegmentIdsOk() (*[]string, bool)`

GetExcludedSegmentIdsOk returns a tuple with the ExcludedSegmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSegmentIds

`func (o *JourneyAudience) SetExcludedSegmentIds(v []string)`

SetExcludedSegmentIds sets ExcludedSegmentIds field to given value.

### HasExcludedSegmentIds

`func (o *JourneyAudience) HasExcludedSegmentIds() bool`

HasExcludedSegmentIds returns a boolean if a field has been set.

### GetFutureAdditionsOnly

`func (o *JourneyAudience) GetFutureAdditionsOnly() bool`

GetFutureAdditionsOnly returns the FutureAdditionsOnly field if non-nil, zero value otherwise.

### GetFutureAdditionsOnlyOk

`func (o *JourneyAudience) GetFutureAdditionsOnlyOk() (*bool, bool)`

GetFutureAdditionsOnlyOk returns a tuple with the FutureAdditionsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureAdditionsOnly

`func (o *JourneyAudience) SetFutureAdditionsOnly(v bool)`

SetFutureAdditionsOnly sets FutureAdditionsOnly field to given value.

### HasFutureAdditionsOnly

`func (o *JourneyAudience) HasFutureAdditionsOnly() bool`

HasFutureAdditionsOnly returns a boolean if a field has been set.

### SetFutureAdditionsOnlyNil

`func (o *JourneyAudience) SetFutureAdditionsOnlyNil(b bool)`

 SetFutureAdditionsOnlyNil sets the value for FutureAdditionsOnly to be an explicit nil

### UnsetFutureAdditionsOnly
`func (o *JourneyAudience) UnsetFutureAdditionsOnly()`

UnsetFutureAdditionsOnly ensures that no value is present for FutureAdditionsOnly, not even an explicit nil
### GetName

`func (o *JourneyAudience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JourneyAudience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JourneyAudience) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JourneyAudience) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttributes

`func (o *JourneyAudience) GetAttributes() [][]JourneyEventAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *JourneyAudience) GetAttributesOk() (*[][]JourneyEventAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *JourneyAudience) SetAttributes(v [][]JourneyEventAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *JourneyAudience) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)



# JourneyBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Server-assigned branch identifier. Read-only on create; echo it on update to keep the branch. | [optional] 
**Condition** | Pointer to [**JourneyCondition**](JourneyCondition.md) |  | [optional] 
**Weight** | Pointer to **float32** | Branch weight for split_range nodes. Weights across a node&#39;s branches must sum to 100. | [optional] 
**Nodes** | Pointer to [**[]JourneyNode**](JourneyNode.md) | Nodes run when this branch is taken, before flow converges to the next sibling node. | [optional] 

## Methods

### NewJourneyBranch

`func NewJourneyBranch() *JourneyBranch`

NewJourneyBranch instantiates a new JourneyBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyBranchWithDefaults

`func NewJourneyBranchWithDefaults() *JourneyBranch`

NewJourneyBranchWithDefaults instantiates a new JourneyBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JourneyBranch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JourneyBranch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JourneyBranch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JourneyBranch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCondition

`func (o *JourneyBranch) GetCondition() JourneyCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *JourneyBranch) GetConditionOk() (*JourneyCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *JourneyBranch) SetCondition(v JourneyCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *JourneyBranch) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetWeight

`func (o *JourneyBranch) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *JourneyBranch) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *JourneyBranch) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *JourneyBranch) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetNodes

`func (o *JourneyBranch) GetNodes() []JourneyNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *JourneyBranch) GetNodesOk() (*[]JourneyNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *JourneyBranch) SetNodes(v []JourneyNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *JourneyBranch) HasNodes() bool`

HasNodes returns a boolean if a field has been set.


[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)



# JourneySchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | **string** | The requested operation. | 
**Document** | **string** | The document that the schema describes. For update_node, merged_node is the complete node after applying the patch. | 
**Schema** | **map[string]interface{}** | A self-contained JSON Schema Draft 6 document with local $defs references. | 

## Methods

### NewJourneySchemaResponse

`func NewJourneySchemaResponse(operation string, document string, schema map[string]interface{}, ) *JourneySchemaResponse`

NewJourneySchemaResponse instantiates a new JourneySchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneySchemaResponseWithDefaults

`func NewJourneySchemaResponseWithDefaults() *JourneySchemaResponse`

NewJourneySchemaResponseWithDefaults instantiates a new JourneySchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *JourneySchemaResponse) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *JourneySchemaResponse) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *JourneySchemaResponse) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetDocument

`func (o *JourneySchemaResponse) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *JourneySchemaResponse) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *JourneySchemaResponse) SetDocument(v string)`

SetDocument sets Document field to given value.


### GetSchema

`func (o *JourneySchemaResponse) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *JourneySchemaResponse) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *JourneySchemaResponse) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.



[[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference) [[Back to README]](https://github.com/OneSignal/onesignal-go-api)



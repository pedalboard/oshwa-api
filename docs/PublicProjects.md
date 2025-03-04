# PublicProjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of results | [optional] 
**Offset** | Pointer to **int32** | Number of results skipped | [optional] 
**Limit** | Pointer to **int32** | Number of results per request | [optional] 
**Items** | Pointer to [**[]PublicProjectsItemsInner**](PublicProjectsItemsInner.md) | List of results | [optional] 

## Methods

### NewPublicProjects

`func NewPublicProjects() *PublicProjects`

NewPublicProjects instantiates a new PublicProjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicProjectsWithDefaults

`func NewPublicProjectsWithDefaults() *PublicProjects`

NewPublicProjectsWithDefaults instantiates a new PublicProjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PublicProjects) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PublicProjects) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PublicProjects) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PublicProjects) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOffset

`func (o *PublicProjects) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PublicProjects) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PublicProjects) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PublicProjects) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PublicProjects) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PublicProjects) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PublicProjects) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PublicProjects) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetItems

`func (o *PublicProjects) GetItems() []PublicProjectsItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PublicProjects) GetItemsOk() (*[]PublicProjectsItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PublicProjects) SetItems(v []PublicProjectsItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *PublicProjects) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



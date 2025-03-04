# PublicProjectsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OshwaUid** | Pointer to **string** | Assigned by OSHWA admin | [optional] 
**ResponsibleParty** | Pointer to **string** | Name of Individual, Company, or Organization Responsible for the Certified Item. Required. | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**PublicContact** | Pointer to **string** |  | [optional] 
**ProjectName** | Pointer to **string** |  | [optional] 
**ProjectWebsite** | Pointer to **string** |  | [optional] 
**ProjectVersion** | Pointer to **string** | e.g. 1.0 | [optional] 
**PreviousVersions** | Pointer to **[]string** | An array of OSHWA UIDs | [optional] 
**ProjectDescription** | Pointer to **string** |  | [optional] 
**PrimaryType** | Pointer to **string** | Primary project type. The GET api/options endpoint provides all valid options. | [optional] 
**AdditionalType** | Pointer to **[]string** | Additional project types. The GET api/options endpoint provides all valid options. | [optional] 
**ProjectKeywords** | Pointer to **[]string** | Additional searchable keywords | [optional] 
**Citations** | Pointer to [**[]ProjectCitationsInner**](ProjectCitationsInner.md) | If the project incorporates or builds upon other open projects that are not currently certified by OSHWA, this field can be used to cite those projects. | [optional] 
**DocumentationUrl** | Pointer to **string** | URL for project documentation | [optional] 
**HardwareLicense** | Pointer to **string** | The GET api/options endpoint provides all valid options. | [optional] 
**SoftwareLicense** | Pointer to **string** | The GET api/options endpoint provides all valid options. | [optional] 
**DocumentationLicense** | Pointer to **string** | The GET api/options endpoint provides all valid options. | [optional] 
**CertificationDate** | Pointer to **string** | Certification date | [optional] 

## Methods

### NewPublicProjectsItemsInner

`func NewPublicProjectsItemsInner() *PublicProjectsItemsInner`

NewPublicProjectsItemsInner instantiates a new PublicProjectsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicProjectsItemsInnerWithDefaults

`func NewPublicProjectsItemsInnerWithDefaults() *PublicProjectsItemsInner`

NewPublicProjectsItemsInnerWithDefaults instantiates a new PublicProjectsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOshwaUid

`func (o *PublicProjectsItemsInner) GetOshwaUid() string`

GetOshwaUid returns the OshwaUid field if non-nil, zero value otherwise.

### GetOshwaUidOk

`func (o *PublicProjectsItemsInner) GetOshwaUidOk() (*string, bool)`

GetOshwaUidOk returns a tuple with the OshwaUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOshwaUid

`func (o *PublicProjectsItemsInner) SetOshwaUid(v string)`

SetOshwaUid sets OshwaUid field to given value.

### HasOshwaUid

`func (o *PublicProjectsItemsInner) HasOshwaUid() bool`

HasOshwaUid returns a boolean if a field has been set.

### GetResponsibleParty

`func (o *PublicProjectsItemsInner) GetResponsibleParty() string`

GetResponsibleParty returns the ResponsibleParty field if non-nil, zero value otherwise.

### GetResponsiblePartyOk

`func (o *PublicProjectsItemsInner) GetResponsiblePartyOk() (*string, bool)`

GetResponsiblePartyOk returns a tuple with the ResponsibleParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleParty

`func (o *PublicProjectsItemsInner) SetResponsibleParty(v string)`

SetResponsibleParty sets ResponsibleParty field to given value.

### HasResponsibleParty

`func (o *PublicProjectsItemsInner) HasResponsibleParty() bool`

HasResponsibleParty returns a boolean if a field has been set.

### GetCountry

`func (o *PublicProjectsItemsInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PublicProjectsItemsInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PublicProjectsItemsInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PublicProjectsItemsInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPublicContact

`func (o *PublicProjectsItemsInner) GetPublicContact() string`

GetPublicContact returns the PublicContact field if non-nil, zero value otherwise.

### GetPublicContactOk

`func (o *PublicProjectsItemsInner) GetPublicContactOk() (*string, bool)`

GetPublicContactOk returns a tuple with the PublicContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicContact

`func (o *PublicProjectsItemsInner) SetPublicContact(v string)`

SetPublicContact sets PublicContact field to given value.

### HasPublicContact

`func (o *PublicProjectsItemsInner) HasPublicContact() bool`

HasPublicContact returns a boolean if a field has been set.

### GetProjectName

`func (o *PublicProjectsItemsInner) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *PublicProjectsItemsInner) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *PublicProjectsItemsInner) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *PublicProjectsItemsInner) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetProjectWebsite

`func (o *PublicProjectsItemsInner) GetProjectWebsite() string`

GetProjectWebsite returns the ProjectWebsite field if non-nil, zero value otherwise.

### GetProjectWebsiteOk

`func (o *PublicProjectsItemsInner) GetProjectWebsiteOk() (*string, bool)`

GetProjectWebsiteOk returns a tuple with the ProjectWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectWebsite

`func (o *PublicProjectsItemsInner) SetProjectWebsite(v string)`

SetProjectWebsite sets ProjectWebsite field to given value.

### HasProjectWebsite

`func (o *PublicProjectsItemsInner) HasProjectWebsite() bool`

HasProjectWebsite returns a boolean if a field has been set.

### GetProjectVersion

`func (o *PublicProjectsItemsInner) GetProjectVersion() string`

GetProjectVersion returns the ProjectVersion field if non-nil, zero value otherwise.

### GetProjectVersionOk

`func (o *PublicProjectsItemsInner) GetProjectVersionOk() (*string, bool)`

GetProjectVersionOk returns a tuple with the ProjectVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectVersion

`func (o *PublicProjectsItemsInner) SetProjectVersion(v string)`

SetProjectVersion sets ProjectVersion field to given value.

### HasProjectVersion

`func (o *PublicProjectsItemsInner) HasProjectVersion() bool`

HasProjectVersion returns a boolean if a field has been set.

### GetPreviousVersions

`func (o *PublicProjectsItemsInner) GetPreviousVersions() []string`

GetPreviousVersions returns the PreviousVersions field if non-nil, zero value otherwise.

### GetPreviousVersionsOk

`func (o *PublicProjectsItemsInner) GetPreviousVersionsOk() (*[]string, bool)`

GetPreviousVersionsOk returns a tuple with the PreviousVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersions

`func (o *PublicProjectsItemsInner) SetPreviousVersions(v []string)`

SetPreviousVersions sets PreviousVersions field to given value.

### HasPreviousVersions

`func (o *PublicProjectsItemsInner) HasPreviousVersions() bool`

HasPreviousVersions returns a boolean if a field has been set.

### GetProjectDescription

`func (o *PublicProjectsItemsInner) GetProjectDescription() string`

GetProjectDescription returns the ProjectDescription field if non-nil, zero value otherwise.

### GetProjectDescriptionOk

`func (o *PublicProjectsItemsInner) GetProjectDescriptionOk() (*string, bool)`

GetProjectDescriptionOk returns a tuple with the ProjectDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDescription

`func (o *PublicProjectsItemsInner) SetProjectDescription(v string)`

SetProjectDescription sets ProjectDescription field to given value.

### HasProjectDescription

`func (o *PublicProjectsItemsInner) HasProjectDescription() bool`

HasProjectDescription returns a boolean if a field has been set.

### GetPrimaryType

`func (o *PublicProjectsItemsInner) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *PublicProjectsItemsInner) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *PublicProjectsItemsInner) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.

### HasPrimaryType

`func (o *PublicProjectsItemsInner) HasPrimaryType() bool`

HasPrimaryType returns a boolean if a field has been set.

### GetAdditionalType

`func (o *PublicProjectsItemsInner) GetAdditionalType() []string`

GetAdditionalType returns the AdditionalType field if non-nil, zero value otherwise.

### GetAdditionalTypeOk

`func (o *PublicProjectsItemsInner) GetAdditionalTypeOk() (*[]string, bool)`

GetAdditionalTypeOk returns a tuple with the AdditionalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalType

`func (o *PublicProjectsItemsInner) SetAdditionalType(v []string)`

SetAdditionalType sets AdditionalType field to given value.

### HasAdditionalType

`func (o *PublicProjectsItemsInner) HasAdditionalType() bool`

HasAdditionalType returns a boolean if a field has been set.

### GetProjectKeywords

`func (o *PublicProjectsItemsInner) GetProjectKeywords() []string`

GetProjectKeywords returns the ProjectKeywords field if non-nil, zero value otherwise.

### GetProjectKeywordsOk

`func (o *PublicProjectsItemsInner) GetProjectKeywordsOk() (*[]string, bool)`

GetProjectKeywordsOk returns a tuple with the ProjectKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKeywords

`func (o *PublicProjectsItemsInner) SetProjectKeywords(v []string)`

SetProjectKeywords sets ProjectKeywords field to given value.

### HasProjectKeywords

`func (o *PublicProjectsItemsInner) HasProjectKeywords() bool`

HasProjectKeywords returns a boolean if a field has been set.

### GetCitations

`func (o *PublicProjectsItemsInner) GetCitations() []ProjectCitationsInner`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *PublicProjectsItemsInner) GetCitationsOk() (*[]ProjectCitationsInner, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *PublicProjectsItemsInner) SetCitations(v []ProjectCitationsInner)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *PublicProjectsItemsInner) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *PublicProjectsItemsInner) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *PublicProjectsItemsInner) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *PublicProjectsItemsInner) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *PublicProjectsItemsInner) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetHardwareLicense

`func (o *PublicProjectsItemsInner) GetHardwareLicense() string`

GetHardwareLicense returns the HardwareLicense field if non-nil, zero value otherwise.

### GetHardwareLicenseOk

`func (o *PublicProjectsItemsInner) GetHardwareLicenseOk() (*string, bool)`

GetHardwareLicenseOk returns a tuple with the HardwareLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareLicense

`func (o *PublicProjectsItemsInner) SetHardwareLicense(v string)`

SetHardwareLicense sets HardwareLicense field to given value.

### HasHardwareLicense

`func (o *PublicProjectsItemsInner) HasHardwareLicense() bool`

HasHardwareLicense returns a boolean if a field has been set.

### GetSoftwareLicense

`func (o *PublicProjectsItemsInner) GetSoftwareLicense() string`

GetSoftwareLicense returns the SoftwareLicense field if non-nil, zero value otherwise.

### GetSoftwareLicenseOk

`func (o *PublicProjectsItemsInner) GetSoftwareLicenseOk() (*string, bool)`

GetSoftwareLicenseOk returns a tuple with the SoftwareLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareLicense

`func (o *PublicProjectsItemsInner) SetSoftwareLicense(v string)`

SetSoftwareLicense sets SoftwareLicense field to given value.

### HasSoftwareLicense

`func (o *PublicProjectsItemsInner) HasSoftwareLicense() bool`

HasSoftwareLicense returns a boolean if a field has been set.

### GetDocumentationLicense

`func (o *PublicProjectsItemsInner) GetDocumentationLicense() string`

GetDocumentationLicense returns the DocumentationLicense field if non-nil, zero value otherwise.

### GetDocumentationLicenseOk

`func (o *PublicProjectsItemsInner) GetDocumentationLicenseOk() (*string, bool)`

GetDocumentationLicenseOk returns a tuple with the DocumentationLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLicense

`func (o *PublicProjectsItemsInner) SetDocumentationLicense(v string)`

SetDocumentationLicense sets DocumentationLicense field to given value.

### HasDocumentationLicense

`func (o *PublicProjectsItemsInner) HasDocumentationLicense() bool`

HasDocumentationLicense returns a boolean if a field has been set.

### GetCertificationDate

`func (o *PublicProjectsItemsInner) GetCertificationDate() string`

GetCertificationDate returns the CertificationDate field if non-nil, zero value otherwise.

### GetCertificationDateOk

`func (o *PublicProjectsItemsInner) GetCertificationDateOk() (*string, bool)`

GetCertificationDateOk returns a tuple with the CertificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationDate

`func (o *PublicProjectsItemsInner) SetCertificationDate(v string)`

SetCertificationDate sets CertificationDate field to given value.

### HasCertificationDate

`func (o *PublicProjectsItemsInner) HasCertificationDate() bool`

HasCertificationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



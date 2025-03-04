# PublicProject

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

### NewPublicProject

`func NewPublicProject() *PublicProject`

NewPublicProject instantiates a new PublicProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicProjectWithDefaults

`func NewPublicProjectWithDefaults() *PublicProject`

NewPublicProjectWithDefaults instantiates a new PublicProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOshwaUid

`func (o *PublicProject) GetOshwaUid() string`

GetOshwaUid returns the OshwaUid field if non-nil, zero value otherwise.

### GetOshwaUidOk

`func (o *PublicProject) GetOshwaUidOk() (*string, bool)`

GetOshwaUidOk returns a tuple with the OshwaUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOshwaUid

`func (o *PublicProject) SetOshwaUid(v string)`

SetOshwaUid sets OshwaUid field to given value.

### HasOshwaUid

`func (o *PublicProject) HasOshwaUid() bool`

HasOshwaUid returns a boolean if a field has been set.

### GetResponsibleParty

`func (o *PublicProject) GetResponsibleParty() string`

GetResponsibleParty returns the ResponsibleParty field if non-nil, zero value otherwise.

### GetResponsiblePartyOk

`func (o *PublicProject) GetResponsiblePartyOk() (*string, bool)`

GetResponsiblePartyOk returns a tuple with the ResponsibleParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleParty

`func (o *PublicProject) SetResponsibleParty(v string)`

SetResponsibleParty sets ResponsibleParty field to given value.

### HasResponsibleParty

`func (o *PublicProject) HasResponsibleParty() bool`

HasResponsibleParty returns a boolean if a field has been set.

### GetCountry

`func (o *PublicProject) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PublicProject) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PublicProject) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PublicProject) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPublicContact

`func (o *PublicProject) GetPublicContact() string`

GetPublicContact returns the PublicContact field if non-nil, zero value otherwise.

### GetPublicContactOk

`func (o *PublicProject) GetPublicContactOk() (*string, bool)`

GetPublicContactOk returns a tuple with the PublicContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicContact

`func (o *PublicProject) SetPublicContact(v string)`

SetPublicContact sets PublicContact field to given value.

### HasPublicContact

`func (o *PublicProject) HasPublicContact() bool`

HasPublicContact returns a boolean if a field has been set.

### GetProjectName

`func (o *PublicProject) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *PublicProject) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *PublicProject) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *PublicProject) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetProjectWebsite

`func (o *PublicProject) GetProjectWebsite() string`

GetProjectWebsite returns the ProjectWebsite field if non-nil, zero value otherwise.

### GetProjectWebsiteOk

`func (o *PublicProject) GetProjectWebsiteOk() (*string, bool)`

GetProjectWebsiteOk returns a tuple with the ProjectWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectWebsite

`func (o *PublicProject) SetProjectWebsite(v string)`

SetProjectWebsite sets ProjectWebsite field to given value.

### HasProjectWebsite

`func (o *PublicProject) HasProjectWebsite() bool`

HasProjectWebsite returns a boolean if a field has been set.

### GetProjectVersion

`func (o *PublicProject) GetProjectVersion() string`

GetProjectVersion returns the ProjectVersion field if non-nil, zero value otherwise.

### GetProjectVersionOk

`func (o *PublicProject) GetProjectVersionOk() (*string, bool)`

GetProjectVersionOk returns a tuple with the ProjectVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectVersion

`func (o *PublicProject) SetProjectVersion(v string)`

SetProjectVersion sets ProjectVersion field to given value.

### HasProjectVersion

`func (o *PublicProject) HasProjectVersion() bool`

HasProjectVersion returns a boolean if a field has been set.

### GetPreviousVersions

`func (o *PublicProject) GetPreviousVersions() []string`

GetPreviousVersions returns the PreviousVersions field if non-nil, zero value otherwise.

### GetPreviousVersionsOk

`func (o *PublicProject) GetPreviousVersionsOk() (*[]string, bool)`

GetPreviousVersionsOk returns a tuple with the PreviousVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersions

`func (o *PublicProject) SetPreviousVersions(v []string)`

SetPreviousVersions sets PreviousVersions field to given value.

### HasPreviousVersions

`func (o *PublicProject) HasPreviousVersions() bool`

HasPreviousVersions returns a boolean if a field has been set.

### GetProjectDescription

`func (o *PublicProject) GetProjectDescription() string`

GetProjectDescription returns the ProjectDescription field if non-nil, zero value otherwise.

### GetProjectDescriptionOk

`func (o *PublicProject) GetProjectDescriptionOk() (*string, bool)`

GetProjectDescriptionOk returns a tuple with the ProjectDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDescription

`func (o *PublicProject) SetProjectDescription(v string)`

SetProjectDescription sets ProjectDescription field to given value.

### HasProjectDescription

`func (o *PublicProject) HasProjectDescription() bool`

HasProjectDescription returns a boolean if a field has been set.

### GetPrimaryType

`func (o *PublicProject) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *PublicProject) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *PublicProject) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.

### HasPrimaryType

`func (o *PublicProject) HasPrimaryType() bool`

HasPrimaryType returns a boolean if a field has been set.

### GetAdditionalType

`func (o *PublicProject) GetAdditionalType() []string`

GetAdditionalType returns the AdditionalType field if non-nil, zero value otherwise.

### GetAdditionalTypeOk

`func (o *PublicProject) GetAdditionalTypeOk() (*[]string, bool)`

GetAdditionalTypeOk returns a tuple with the AdditionalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalType

`func (o *PublicProject) SetAdditionalType(v []string)`

SetAdditionalType sets AdditionalType field to given value.

### HasAdditionalType

`func (o *PublicProject) HasAdditionalType() bool`

HasAdditionalType returns a boolean if a field has been set.

### GetProjectKeywords

`func (o *PublicProject) GetProjectKeywords() []string`

GetProjectKeywords returns the ProjectKeywords field if non-nil, zero value otherwise.

### GetProjectKeywordsOk

`func (o *PublicProject) GetProjectKeywordsOk() (*[]string, bool)`

GetProjectKeywordsOk returns a tuple with the ProjectKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKeywords

`func (o *PublicProject) SetProjectKeywords(v []string)`

SetProjectKeywords sets ProjectKeywords field to given value.

### HasProjectKeywords

`func (o *PublicProject) HasProjectKeywords() bool`

HasProjectKeywords returns a boolean if a field has been set.

### GetCitations

`func (o *PublicProject) GetCitations() []ProjectCitationsInner`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *PublicProject) GetCitationsOk() (*[]ProjectCitationsInner, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *PublicProject) SetCitations(v []ProjectCitationsInner)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *PublicProject) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *PublicProject) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *PublicProject) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *PublicProject) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *PublicProject) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetHardwareLicense

`func (o *PublicProject) GetHardwareLicense() string`

GetHardwareLicense returns the HardwareLicense field if non-nil, zero value otherwise.

### GetHardwareLicenseOk

`func (o *PublicProject) GetHardwareLicenseOk() (*string, bool)`

GetHardwareLicenseOk returns a tuple with the HardwareLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareLicense

`func (o *PublicProject) SetHardwareLicense(v string)`

SetHardwareLicense sets HardwareLicense field to given value.

### HasHardwareLicense

`func (o *PublicProject) HasHardwareLicense() bool`

HasHardwareLicense returns a boolean if a field has been set.

### GetSoftwareLicense

`func (o *PublicProject) GetSoftwareLicense() string`

GetSoftwareLicense returns the SoftwareLicense field if non-nil, zero value otherwise.

### GetSoftwareLicenseOk

`func (o *PublicProject) GetSoftwareLicenseOk() (*string, bool)`

GetSoftwareLicenseOk returns a tuple with the SoftwareLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareLicense

`func (o *PublicProject) SetSoftwareLicense(v string)`

SetSoftwareLicense sets SoftwareLicense field to given value.

### HasSoftwareLicense

`func (o *PublicProject) HasSoftwareLicense() bool`

HasSoftwareLicense returns a boolean if a field has been set.

### GetDocumentationLicense

`func (o *PublicProject) GetDocumentationLicense() string`

GetDocumentationLicense returns the DocumentationLicense field if non-nil, zero value otherwise.

### GetDocumentationLicenseOk

`func (o *PublicProject) GetDocumentationLicenseOk() (*string, bool)`

GetDocumentationLicenseOk returns a tuple with the DocumentationLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLicense

`func (o *PublicProject) SetDocumentationLicense(v string)`

SetDocumentationLicense sets DocumentationLicense field to given value.

### HasDocumentationLicense

`func (o *PublicProject) HasDocumentationLicense() bool`

HasDocumentationLicense returns a boolean if a field has been set.

### GetCertificationDate

`func (o *PublicProject) GetCertificationDate() string`

GetCertificationDate returns the CertificationDate field if non-nil, zero value otherwise.

### GetCertificationDateOk

`func (o *PublicProject) GetCertificationDateOk() (*string, bool)`

GetCertificationDateOk returns a tuple with the CertificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationDate

`func (o *PublicProject) SetCertificationDate(v string)`

SetCertificationDate sets CertificationDate field to given value.

### HasCertificationDate

`func (o *PublicProject) HasCertificationDate() bool`

HasCertificationDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



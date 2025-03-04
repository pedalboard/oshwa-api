# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponsiblePartyType** | **string** | Responsible party type. Must be either \&quot;Individual\&quot;, \&quot;Company\&quot;, or \&quot;Organization\&quot;. Required. | 
**ResponsibleParty** | **string** | Name of Individual, Company, or Organization Responsible for the Certified Item. Required. | 
**BindingParty** | **string** | If not an Individual, name of Individual with Authority to Bind the Company or Organization. Required only if responsiblePartyType is not \&quot;Individual\&quot;. | 
**Country** | **string** |  | 
**StreetAddress1** | Pointer to **string** |  | [optional] 
**StreetAddress2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**PrivateContact** | Pointer to **string** |  | [optional] 
**PublicContact** | Pointer to **string** |  | [optional] 
**ProjectName** | **string** |  | 
**ProjectWebsite** | Pointer to **string** |  | [optional] 
**ProjectVersion** | Pointer to **string** | e.g. 1.0 | [optional] 
**PreviousVersions** | Pointer to **[]string** | An array of OSHWA UIDs | [optional] 
**ProjectDescription** | Pointer to **string** |  | [optional] 
**PrimaryType** | **string** | Primary project type. The GET api/options endpoint provides all valid options. | 
**AdditionalType** | Pointer to **[]string** | Additional project types. The GET api/options endpoint provides all valid options. | [optional] 
**ProjectKeywords** | Pointer to **[]string** | Additional searchable keywords | [optional] 
**Citations** | Pointer to [**[]ProjectCitationsInner**](ProjectCitationsInner.md) | If the project incorporates or builds upon other open projects that are not currently certified by OSHWA, this field can be used to cite those projects. | [optional] 
**DocumentationUrl** | Pointer to **string** | URL for project documentation | [optional] 
**AvailableFileFormat** | Pointer to **bool** | All project documentation and design files are available in the preferred format for making changes. | [optional] 
**HardwareLicense** | **string** | The GET api/options endpoint provides all valid options. | 
**SoftwareLicense** | **string** | The GET api/options endpoint provides all valid options. | 
**DocumentationLicense** | **string** | The GET api/options endpoint provides all valid options. | 
**NoCommercialRestriction** | **string** | The project is licensed in a way to allow for modifications and derivative works without commercial restriction. | 
**ExplanationNcr** | **string** | Explanation is required if noCommercialRestriction is false. | 
**NoDocumentationRestriction** | **bool** | There is no restriction within my control to selling or giving away the project documentation. | 
**ExplanationNdr** | **string** | Explanation is required if noDocumentationRestriction is false. | 
**OpenHardwareComponents** | **bool** | Where possible, I have chosen to use components in my hardware that are openly licensed. | 
**ExplanationOhwc** | **string** | Explanation is required if openHardwareComponents is false. | 
**CreatorContribution** | **bool** | I understand and comply with the \&quot;Creator Contribution requirement,\&quot; explained in the Requirements for Certification. | 
**ExplanationCcr** | **string** | Explanation is required if creatorContribution is false. | 
**NoUseRestriction** | **bool** | There is no restriction on the use by persons or groups, or by the field of endeavor. | 
**ExplanationNur** | **string** | Explanation is required if noUseRestriction is false. | 
**RedistributedWork** | **bool** | The rights granted by any license on the project applies to all whom the work is redistributed to. | 
**ExplanationRwr** | **string** | Explanation is required if redistributedWork is false. | 
**NoSpecificProduct** | **bool** | The rights granted under any license on the project do not depend on the licensed work being part of a specific product. | 
**ExplanationNsp** | **string** | Explanation is required if noSpecificProduct is false. | 
**NoComponentRestriction** | **bool** | The rights granted under any license on the project do not restrict other hardware or software, for example by requiring that all other hardware or software sold with the item be open source. | 
**ExplanationNor** | **string** | Explanation is required if noComponentRestriction is false. | 
**TechnologyNeutral** | **bool** | The rights granted under any license on the project are technology neutral. | 
**ExplanationTn** | **string** | Explanation is required if technologyNeutral is false. | 
**CertificationMarkTerms** | [**ProjectCertificationMarkTerms**](ProjectCertificationMarkTerms.md) |  | 
**ExplanationCertificationTerms** | **string** | Explanation for certification mark terms | 
**Relationship** | **string** | Relationship | 
**AgreementTerms** | **bool** | Agreement to terms | 
**ParentName** | **string** | Parent name | 

## Methods

### NewProject

`func NewProject(responsiblePartyType string, responsibleParty string, bindingParty string, country string, projectName string, primaryType string, hardwareLicense string, softwareLicense string, documentationLicense string, noCommercialRestriction string, explanationNcr string, noDocumentationRestriction bool, explanationNdr string, openHardwareComponents bool, explanationOhwc string, creatorContribution bool, explanationCcr string, noUseRestriction bool, explanationNur string, redistributedWork bool, explanationRwr string, noSpecificProduct bool, explanationNsp string, noComponentRestriction bool, explanationNor string, technologyNeutral bool, explanationTn string, certificationMarkTerms ProjectCertificationMarkTerms, explanationCertificationTerms string, relationship string, agreementTerms bool, parentName string, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponsiblePartyType

`func (o *Project) GetResponsiblePartyType() string`

GetResponsiblePartyType returns the ResponsiblePartyType field if non-nil, zero value otherwise.

### GetResponsiblePartyTypeOk

`func (o *Project) GetResponsiblePartyTypeOk() (*string, bool)`

GetResponsiblePartyTypeOk returns a tuple with the ResponsiblePartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsiblePartyType

`func (o *Project) SetResponsiblePartyType(v string)`

SetResponsiblePartyType sets ResponsiblePartyType field to given value.


### GetResponsibleParty

`func (o *Project) GetResponsibleParty() string`

GetResponsibleParty returns the ResponsibleParty field if non-nil, zero value otherwise.

### GetResponsiblePartyOk

`func (o *Project) GetResponsiblePartyOk() (*string, bool)`

GetResponsiblePartyOk returns a tuple with the ResponsibleParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleParty

`func (o *Project) SetResponsibleParty(v string)`

SetResponsibleParty sets ResponsibleParty field to given value.


### GetBindingParty

`func (o *Project) GetBindingParty() string`

GetBindingParty returns the BindingParty field if non-nil, zero value otherwise.

### GetBindingPartyOk

`func (o *Project) GetBindingPartyOk() (*string, bool)`

GetBindingPartyOk returns a tuple with the BindingParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingParty

`func (o *Project) SetBindingParty(v string)`

SetBindingParty sets BindingParty field to given value.


### GetCountry

`func (o *Project) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Project) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Project) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetStreetAddress1

`func (o *Project) GetStreetAddress1() string`

GetStreetAddress1 returns the StreetAddress1 field if non-nil, zero value otherwise.

### GetStreetAddress1Ok

`func (o *Project) GetStreetAddress1Ok() (*string, bool)`

GetStreetAddress1Ok returns a tuple with the StreetAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress1

`func (o *Project) SetStreetAddress1(v string)`

SetStreetAddress1 sets StreetAddress1 field to given value.

### HasStreetAddress1

`func (o *Project) HasStreetAddress1() bool`

HasStreetAddress1 returns a boolean if a field has been set.

### GetStreetAddress2

`func (o *Project) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *Project) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *Project) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *Project) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### GetCity

`func (o *Project) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Project) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Project) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Project) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Project) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Project) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Project) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Project) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *Project) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Project) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Project) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *Project) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetPrivateContact

`func (o *Project) GetPrivateContact() string`

GetPrivateContact returns the PrivateContact field if non-nil, zero value otherwise.

### GetPrivateContactOk

`func (o *Project) GetPrivateContactOk() (*string, bool)`

GetPrivateContactOk returns a tuple with the PrivateContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateContact

`func (o *Project) SetPrivateContact(v string)`

SetPrivateContact sets PrivateContact field to given value.

### HasPrivateContact

`func (o *Project) HasPrivateContact() bool`

HasPrivateContact returns a boolean if a field has been set.

### GetPublicContact

`func (o *Project) GetPublicContact() string`

GetPublicContact returns the PublicContact field if non-nil, zero value otherwise.

### GetPublicContactOk

`func (o *Project) GetPublicContactOk() (*string, bool)`

GetPublicContactOk returns a tuple with the PublicContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicContact

`func (o *Project) SetPublicContact(v string)`

SetPublicContact sets PublicContact field to given value.

### HasPublicContact

`func (o *Project) HasPublicContact() bool`

HasPublicContact returns a boolean if a field has been set.

### GetProjectName

`func (o *Project) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *Project) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *Project) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetProjectWebsite

`func (o *Project) GetProjectWebsite() string`

GetProjectWebsite returns the ProjectWebsite field if non-nil, zero value otherwise.

### GetProjectWebsiteOk

`func (o *Project) GetProjectWebsiteOk() (*string, bool)`

GetProjectWebsiteOk returns a tuple with the ProjectWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectWebsite

`func (o *Project) SetProjectWebsite(v string)`

SetProjectWebsite sets ProjectWebsite field to given value.

### HasProjectWebsite

`func (o *Project) HasProjectWebsite() bool`

HasProjectWebsite returns a boolean if a field has been set.

### GetProjectVersion

`func (o *Project) GetProjectVersion() string`

GetProjectVersion returns the ProjectVersion field if non-nil, zero value otherwise.

### GetProjectVersionOk

`func (o *Project) GetProjectVersionOk() (*string, bool)`

GetProjectVersionOk returns a tuple with the ProjectVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectVersion

`func (o *Project) SetProjectVersion(v string)`

SetProjectVersion sets ProjectVersion field to given value.

### HasProjectVersion

`func (o *Project) HasProjectVersion() bool`

HasProjectVersion returns a boolean if a field has been set.

### GetPreviousVersions

`func (o *Project) GetPreviousVersions() []string`

GetPreviousVersions returns the PreviousVersions field if non-nil, zero value otherwise.

### GetPreviousVersionsOk

`func (o *Project) GetPreviousVersionsOk() (*[]string, bool)`

GetPreviousVersionsOk returns a tuple with the PreviousVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersions

`func (o *Project) SetPreviousVersions(v []string)`

SetPreviousVersions sets PreviousVersions field to given value.

### HasPreviousVersions

`func (o *Project) HasPreviousVersions() bool`

HasPreviousVersions returns a boolean if a field has been set.

### GetProjectDescription

`func (o *Project) GetProjectDescription() string`

GetProjectDescription returns the ProjectDescription field if non-nil, zero value otherwise.

### GetProjectDescriptionOk

`func (o *Project) GetProjectDescriptionOk() (*string, bool)`

GetProjectDescriptionOk returns a tuple with the ProjectDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDescription

`func (o *Project) SetProjectDescription(v string)`

SetProjectDescription sets ProjectDescription field to given value.

### HasProjectDescription

`func (o *Project) HasProjectDescription() bool`

HasProjectDescription returns a boolean if a field has been set.

### GetPrimaryType

`func (o *Project) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *Project) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *Project) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.


### GetAdditionalType

`func (o *Project) GetAdditionalType() []string`

GetAdditionalType returns the AdditionalType field if non-nil, zero value otherwise.

### GetAdditionalTypeOk

`func (o *Project) GetAdditionalTypeOk() (*[]string, bool)`

GetAdditionalTypeOk returns a tuple with the AdditionalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalType

`func (o *Project) SetAdditionalType(v []string)`

SetAdditionalType sets AdditionalType field to given value.

### HasAdditionalType

`func (o *Project) HasAdditionalType() bool`

HasAdditionalType returns a boolean if a field has been set.

### GetProjectKeywords

`func (o *Project) GetProjectKeywords() []string`

GetProjectKeywords returns the ProjectKeywords field if non-nil, zero value otherwise.

### GetProjectKeywordsOk

`func (o *Project) GetProjectKeywordsOk() (*[]string, bool)`

GetProjectKeywordsOk returns a tuple with the ProjectKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKeywords

`func (o *Project) SetProjectKeywords(v []string)`

SetProjectKeywords sets ProjectKeywords field to given value.

### HasProjectKeywords

`func (o *Project) HasProjectKeywords() bool`

HasProjectKeywords returns a boolean if a field has been set.

### GetCitations

`func (o *Project) GetCitations() []ProjectCitationsInner`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *Project) GetCitationsOk() (*[]ProjectCitationsInner, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *Project) SetCitations(v []ProjectCitationsInner)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *Project) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *Project) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *Project) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *Project) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *Project) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetAvailableFileFormat

`func (o *Project) GetAvailableFileFormat() bool`

GetAvailableFileFormat returns the AvailableFileFormat field if non-nil, zero value otherwise.

### GetAvailableFileFormatOk

`func (o *Project) GetAvailableFileFormatOk() (*bool, bool)`

GetAvailableFileFormatOk returns a tuple with the AvailableFileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFileFormat

`func (o *Project) SetAvailableFileFormat(v bool)`

SetAvailableFileFormat sets AvailableFileFormat field to given value.

### HasAvailableFileFormat

`func (o *Project) HasAvailableFileFormat() bool`

HasAvailableFileFormat returns a boolean if a field has been set.

### GetHardwareLicense

`func (o *Project) GetHardwareLicense() string`

GetHardwareLicense returns the HardwareLicense field if non-nil, zero value otherwise.

### GetHardwareLicenseOk

`func (o *Project) GetHardwareLicenseOk() (*string, bool)`

GetHardwareLicenseOk returns a tuple with the HardwareLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareLicense

`func (o *Project) SetHardwareLicense(v string)`

SetHardwareLicense sets HardwareLicense field to given value.


### GetSoftwareLicense

`func (o *Project) GetSoftwareLicense() string`

GetSoftwareLicense returns the SoftwareLicense field if non-nil, zero value otherwise.

### GetSoftwareLicenseOk

`func (o *Project) GetSoftwareLicenseOk() (*string, bool)`

GetSoftwareLicenseOk returns a tuple with the SoftwareLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareLicense

`func (o *Project) SetSoftwareLicense(v string)`

SetSoftwareLicense sets SoftwareLicense field to given value.


### GetDocumentationLicense

`func (o *Project) GetDocumentationLicense() string`

GetDocumentationLicense returns the DocumentationLicense field if non-nil, zero value otherwise.

### GetDocumentationLicenseOk

`func (o *Project) GetDocumentationLicenseOk() (*string, bool)`

GetDocumentationLicenseOk returns a tuple with the DocumentationLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLicense

`func (o *Project) SetDocumentationLicense(v string)`

SetDocumentationLicense sets DocumentationLicense field to given value.


### GetNoCommercialRestriction

`func (o *Project) GetNoCommercialRestriction() string`

GetNoCommercialRestriction returns the NoCommercialRestriction field if non-nil, zero value otherwise.

### GetNoCommercialRestrictionOk

`func (o *Project) GetNoCommercialRestrictionOk() (*string, bool)`

GetNoCommercialRestrictionOk returns a tuple with the NoCommercialRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoCommercialRestriction

`func (o *Project) SetNoCommercialRestriction(v string)`

SetNoCommercialRestriction sets NoCommercialRestriction field to given value.


### GetExplanationNcr

`func (o *Project) GetExplanationNcr() string`

GetExplanationNcr returns the ExplanationNcr field if non-nil, zero value otherwise.

### GetExplanationNcrOk

`func (o *Project) GetExplanationNcrOk() (*string, bool)`

GetExplanationNcrOk returns a tuple with the ExplanationNcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationNcr

`func (o *Project) SetExplanationNcr(v string)`

SetExplanationNcr sets ExplanationNcr field to given value.


### GetNoDocumentationRestriction

`func (o *Project) GetNoDocumentationRestriction() bool`

GetNoDocumentationRestriction returns the NoDocumentationRestriction field if non-nil, zero value otherwise.

### GetNoDocumentationRestrictionOk

`func (o *Project) GetNoDocumentationRestrictionOk() (*bool, bool)`

GetNoDocumentationRestrictionOk returns a tuple with the NoDocumentationRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDocumentationRestriction

`func (o *Project) SetNoDocumentationRestriction(v bool)`

SetNoDocumentationRestriction sets NoDocumentationRestriction field to given value.


### GetExplanationNdr

`func (o *Project) GetExplanationNdr() string`

GetExplanationNdr returns the ExplanationNdr field if non-nil, zero value otherwise.

### GetExplanationNdrOk

`func (o *Project) GetExplanationNdrOk() (*string, bool)`

GetExplanationNdrOk returns a tuple with the ExplanationNdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationNdr

`func (o *Project) SetExplanationNdr(v string)`

SetExplanationNdr sets ExplanationNdr field to given value.


### GetOpenHardwareComponents

`func (o *Project) GetOpenHardwareComponents() bool`

GetOpenHardwareComponents returns the OpenHardwareComponents field if non-nil, zero value otherwise.

### GetOpenHardwareComponentsOk

`func (o *Project) GetOpenHardwareComponentsOk() (*bool, bool)`

GetOpenHardwareComponentsOk returns a tuple with the OpenHardwareComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenHardwareComponents

`func (o *Project) SetOpenHardwareComponents(v bool)`

SetOpenHardwareComponents sets OpenHardwareComponents field to given value.


### GetExplanationOhwc

`func (o *Project) GetExplanationOhwc() string`

GetExplanationOhwc returns the ExplanationOhwc field if non-nil, zero value otherwise.

### GetExplanationOhwcOk

`func (o *Project) GetExplanationOhwcOk() (*string, bool)`

GetExplanationOhwcOk returns a tuple with the ExplanationOhwc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationOhwc

`func (o *Project) SetExplanationOhwc(v string)`

SetExplanationOhwc sets ExplanationOhwc field to given value.


### GetCreatorContribution

`func (o *Project) GetCreatorContribution() bool`

GetCreatorContribution returns the CreatorContribution field if non-nil, zero value otherwise.

### GetCreatorContributionOk

`func (o *Project) GetCreatorContributionOk() (*bool, bool)`

GetCreatorContributionOk returns a tuple with the CreatorContribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorContribution

`func (o *Project) SetCreatorContribution(v bool)`

SetCreatorContribution sets CreatorContribution field to given value.


### GetExplanationCcr

`func (o *Project) GetExplanationCcr() string`

GetExplanationCcr returns the ExplanationCcr field if non-nil, zero value otherwise.

### GetExplanationCcrOk

`func (o *Project) GetExplanationCcrOk() (*string, bool)`

GetExplanationCcrOk returns a tuple with the ExplanationCcr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationCcr

`func (o *Project) SetExplanationCcr(v string)`

SetExplanationCcr sets ExplanationCcr field to given value.


### GetNoUseRestriction

`func (o *Project) GetNoUseRestriction() bool`

GetNoUseRestriction returns the NoUseRestriction field if non-nil, zero value otherwise.

### GetNoUseRestrictionOk

`func (o *Project) GetNoUseRestrictionOk() (*bool, bool)`

GetNoUseRestrictionOk returns a tuple with the NoUseRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoUseRestriction

`func (o *Project) SetNoUseRestriction(v bool)`

SetNoUseRestriction sets NoUseRestriction field to given value.


### GetExplanationNur

`func (o *Project) GetExplanationNur() string`

GetExplanationNur returns the ExplanationNur field if non-nil, zero value otherwise.

### GetExplanationNurOk

`func (o *Project) GetExplanationNurOk() (*string, bool)`

GetExplanationNurOk returns a tuple with the ExplanationNur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationNur

`func (o *Project) SetExplanationNur(v string)`

SetExplanationNur sets ExplanationNur field to given value.


### GetRedistributedWork

`func (o *Project) GetRedistributedWork() bool`

GetRedistributedWork returns the RedistributedWork field if non-nil, zero value otherwise.

### GetRedistributedWorkOk

`func (o *Project) GetRedistributedWorkOk() (*bool, bool)`

GetRedistributedWorkOk returns a tuple with the RedistributedWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedistributedWork

`func (o *Project) SetRedistributedWork(v bool)`

SetRedistributedWork sets RedistributedWork field to given value.


### GetExplanationRwr

`func (o *Project) GetExplanationRwr() string`

GetExplanationRwr returns the ExplanationRwr field if non-nil, zero value otherwise.

### GetExplanationRwrOk

`func (o *Project) GetExplanationRwrOk() (*string, bool)`

GetExplanationRwrOk returns a tuple with the ExplanationRwr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationRwr

`func (o *Project) SetExplanationRwr(v string)`

SetExplanationRwr sets ExplanationRwr field to given value.


### GetNoSpecificProduct

`func (o *Project) GetNoSpecificProduct() bool`

GetNoSpecificProduct returns the NoSpecificProduct field if non-nil, zero value otherwise.

### GetNoSpecificProductOk

`func (o *Project) GetNoSpecificProductOk() (*bool, bool)`

GetNoSpecificProductOk returns a tuple with the NoSpecificProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSpecificProduct

`func (o *Project) SetNoSpecificProduct(v bool)`

SetNoSpecificProduct sets NoSpecificProduct field to given value.


### GetExplanationNsp

`func (o *Project) GetExplanationNsp() string`

GetExplanationNsp returns the ExplanationNsp field if non-nil, zero value otherwise.

### GetExplanationNspOk

`func (o *Project) GetExplanationNspOk() (*string, bool)`

GetExplanationNspOk returns a tuple with the ExplanationNsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationNsp

`func (o *Project) SetExplanationNsp(v string)`

SetExplanationNsp sets ExplanationNsp field to given value.


### GetNoComponentRestriction

`func (o *Project) GetNoComponentRestriction() bool`

GetNoComponentRestriction returns the NoComponentRestriction field if non-nil, zero value otherwise.

### GetNoComponentRestrictionOk

`func (o *Project) GetNoComponentRestrictionOk() (*bool, bool)`

GetNoComponentRestrictionOk returns a tuple with the NoComponentRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoComponentRestriction

`func (o *Project) SetNoComponentRestriction(v bool)`

SetNoComponentRestriction sets NoComponentRestriction field to given value.


### GetExplanationNor

`func (o *Project) GetExplanationNor() string`

GetExplanationNor returns the ExplanationNor field if non-nil, zero value otherwise.

### GetExplanationNorOk

`func (o *Project) GetExplanationNorOk() (*string, bool)`

GetExplanationNorOk returns a tuple with the ExplanationNor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationNor

`func (o *Project) SetExplanationNor(v string)`

SetExplanationNor sets ExplanationNor field to given value.


### GetTechnologyNeutral

`func (o *Project) GetTechnologyNeutral() bool`

GetTechnologyNeutral returns the TechnologyNeutral field if non-nil, zero value otherwise.

### GetTechnologyNeutralOk

`func (o *Project) GetTechnologyNeutralOk() (*bool, bool)`

GetTechnologyNeutralOk returns a tuple with the TechnologyNeutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologyNeutral

`func (o *Project) SetTechnologyNeutral(v bool)`

SetTechnologyNeutral sets TechnologyNeutral field to given value.


### GetExplanationTn

`func (o *Project) GetExplanationTn() string`

GetExplanationTn returns the ExplanationTn field if non-nil, zero value otherwise.

### GetExplanationTnOk

`func (o *Project) GetExplanationTnOk() (*string, bool)`

GetExplanationTnOk returns a tuple with the ExplanationTn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationTn

`func (o *Project) SetExplanationTn(v string)`

SetExplanationTn sets ExplanationTn field to given value.


### GetCertificationMarkTerms

`func (o *Project) GetCertificationMarkTerms() ProjectCertificationMarkTerms`

GetCertificationMarkTerms returns the CertificationMarkTerms field if non-nil, zero value otherwise.

### GetCertificationMarkTermsOk

`func (o *Project) GetCertificationMarkTermsOk() (*ProjectCertificationMarkTerms, bool)`

GetCertificationMarkTermsOk returns a tuple with the CertificationMarkTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificationMarkTerms

`func (o *Project) SetCertificationMarkTerms(v ProjectCertificationMarkTerms)`

SetCertificationMarkTerms sets CertificationMarkTerms field to given value.


### GetExplanationCertificationTerms

`func (o *Project) GetExplanationCertificationTerms() string`

GetExplanationCertificationTerms returns the ExplanationCertificationTerms field if non-nil, zero value otherwise.

### GetExplanationCertificationTermsOk

`func (o *Project) GetExplanationCertificationTermsOk() (*string, bool)`

GetExplanationCertificationTermsOk returns a tuple with the ExplanationCertificationTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanationCertificationTerms

`func (o *Project) SetExplanationCertificationTerms(v string)`

SetExplanationCertificationTerms sets ExplanationCertificationTerms field to given value.


### GetRelationship

`func (o *Project) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *Project) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *Project) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.


### GetAgreementTerms

`func (o *Project) GetAgreementTerms() bool`

GetAgreementTerms returns the AgreementTerms field if non-nil, zero value otherwise.

### GetAgreementTermsOk

`func (o *Project) GetAgreementTermsOk() (*bool, bool)`

GetAgreementTermsOk returns a tuple with the AgreementTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementTerms

`func (o *Project) SetAgreementTerms(v bool)`

SetAgreementTerms sets AgreementTerms field to given value.


### GetParentName

`func (o *Project) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *Project) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *Project) SetParentName(v string)`

SetParentName sets ParentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



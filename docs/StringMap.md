# StringMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**En** | **string** | Text in English.  Will be used as a fallback | 
**Ar** | Pointer to **string** | Text in Arabic. | [optional] 
**Bs** | Pointer to **string** | Text in Bosnian. | [optional] 
**Bg** | Pointer to **string** | Text in Bulgarian. | [optional] 
**Ca** | Pointer to **string** | Text in Catalan. | [optional] 
**ZhHans** | Pointer to **string** | Text in Chinese (Simplified). | [optional] 
**ZhHant** | Pointer to **string** | Text in Chinese (Traditional). | [optional] 
**Zh** | Pointer to **string** | Alias for zh-Hans. | [optional] 
**Hr** | Pointer to **string** | Text in Croatian. | [optional] 
**Cs** | Pointer to **string** | Text in Czech. | [optional] 
**Da** | Pointer to **string** | Text in Danish. | [optional] 
**Nl** | Pointer to **string** | Text in Dutch. | [optional] 
**Et** | Pointer to **string** | Text in Estonian. | [optional] 
**Fi** | Pointer to **string** | Text in Finnish. | [optional] 
**Fr** | Pointer to **string** | Text in French. | [optional] 
**Ka** | Pointer to **string** | Text in Georgian. | [optional] 
**De** | Pointer to **string** | Text in German. | [optional] 
**El** | Pointer to **string** | Text in Greek. | [optional] 
**Hi** | Pointer to **string** | Text in Hindi. | [optional] 
**He** | Pointer to **string** | Text in Hebrew. | [optional] 
**Hu** | Pointer to **string** | Text in Hungarian. | [optional] 
**Id** | Pointer to **string** | Text in Indonesian. | [optional] 
**It** | Pointer to **string** | Text in Italian. | [optional] 
**Ja** | Pointer to **string** | Text in Japanese. | [optional] 
**Ko** | Pointer to **string** | Text in Korean. | [optional] 
**Lv** | Pointer to **string** | Text in Latvian. | [optional] 
**Lt** | Pointer to **string** | Text in Lithuanian. | [optional] 
**Ms** | Pointer to **string** | Text in Malay. | [optional] 
**Nb** | Pointer to **string** | Text in Norwegian. | [optional] 
**Pl** | Pointer to **string** | Text in Polish. | [optional] 
**Fa** | Pointer to **string** | Text in Persian. | [optional] 
**Pt** | Pointer to **string** | Text in Portugese. | [optional] 
**Pa** | Pointer to **string** | Text in Punjabi. | [optional] 
**Ro** | Pointer to **string** | Text in Romanian. | [optional] 
**Ru** | Pointer to **string** | Text in Russian. | [optional] 
**Sr** | Pointer to **string** | Text in Serbian. | [optional] 
**Sk** | Pointer to **string** | Text in Slovak. | [optional] 
**Es** | Pointer to **string** | Text in Spanish. | [optional] 
**Sv** | Pointer to **string** | Text in Swedish. | [optional] 
**Th** | Pointer to **string** | Text in Thai. | [optional] 
**Tr** | Pointer to **string** | Text in Turkish. | [optional] 
**Uk** | Pointer to **string** | Text in Ukrainian. | [optional] 
**Vi** | Pointer to **string** | Text in Vietnamese. | [optional] 

## Methods

### NewStringMap

`func NewStringMap(en string, ) *StringMap`

NewStringMap instantiates a new StringMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringMapWithDefaults

`func NewStringMapWithDefaults() *StringMap`

NewStringMapWithDefaults instantiates a new StringMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEn

`func (o *StringMap) GetEn() string`

GetEn returns the En field if non-nil, zero value otherwise.

### GetEnOk

`func (o *StringMap) GetEnOk() (*string, bool)`

GetEnOk returns a tuple with the En field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEn

`func (o *StringMap) SetEn(v string)`

SetEn sets En field to given value.


### GetAr

`func (o *StringMap) GetAr() string`

GetAr returns the Ar field if non-nil, zero value otherwise.

### GetArOk

`func (o *StringMap) GetArOk() (*string, bool)`

GetArOk returns a tuple with the Ar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAr

`func (o *StringMap) SetAr(v string)`

SetAr sets Ar field to given value.

### HasAr

`func (o *StringMap) HasAr() bool`

HasAr returns a boolean if a field has been set.

### GetBs

`func (o *StringMap) GetBs() string`

GetBs returns the Bs field if non-nil, zero value otherwise.

### GetBsOk

`func (o *StringMap) GetBsOk() (*string, bool)`

GetBsOk returns a tuple with the Bs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBs

`func (o *StringMap) SetBs(v string)`

SetBs sets Bs field to given value.

### HasBs

`func (o *StringMap) HasBs() bool`

HasBs returns a boolean if a field has been set.

### GetBg

`func (o *StringMap) GetBg() string`

GetBg returns the Bg field if non-nil, zero value otherwise.

### GetBgOk

`func (o *StringMap) GetBgOk() (*string, bool)`

GetBgOk returns a tuple with the Bg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBg

`func (o *StringMap) SetBg(v string)`

SetBg sets Bg field to given value.

### HasBg

`func (o *StringMap) HasBg() bool`

HasBg returns a boolean if a field has been set.

### GetCa

`func (o *StringMap) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *StringMap) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *StringMap) SetCa(v string)`

SetCa sets Ca field to given value.

### HasCa

`func (o *StringMap) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetZhHans

`func (o *StringMap) GetZhHans() string`

GetZhHans returns the ZhHans field if non-nil, zero value otherwise.

### GetZhHansOk

`func (o *StringMap) GetZhHansOk() (*string, bool)`

GetZhHansOk returns a tuple with the ZhHans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZhHans

`func (o *StringMap) SetZhHans(v string)`

SetZhHans sets ZhHans field to given value.

### HasZhHans

`func (o *StringMap) HasZhHans() bool`

HasZhHans returns a boolean if a field has been set.

### GetZhHant

`func (o *StringMap) GetZhHant() string`

GetZhHant returns the ZhHant field if non-nil, zero value otherwise.

### GetZhHantOk

`func (o *StringMap) GetZhHantOk() (*string, bool)`

GetZhHantOk returns a tuple with the ZhHant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZhHant

`func (o *StringMap) SetZhHant(v string)`

SetZhHant sets ZhHant field to given value.

### HasZhHant

`func (o *StringMap) HasZhHant() bool`

HasZhHant returns a boolean if a field has been set.

### GetZh

`func (o *StringMap) GetZh() string`

GetZh returns the Zh field if non-nil, zero value otherwise.

### GetZhOk

`func (o *StringMap) GetZhOk() (*string, bool)`

GetZhOk returns a tuple with the Zh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZh

`func (o *StringMap) SetZh(v string)`

SetZh sets Zh field to given value.

### HasZh

`func (o *StringMap) HasZh() bool`

HasZh returns a boolean if a field has been set.

### GetHr

`func (o *StringMap) GetHr() string`

GetHr returns the Hr field if non-nil, zero value otherwise.

### GetHrOk

`func (o *StringMap) GetHrOk() (*string, bool)`

GetHrOk returns a tuple with the Hr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHr

`func (o *StringMap) SetHr(v string)`

SetHr sets Hr field to given value.

### HasHr

`func (o *StringMap) HasHr() bool`

HasHr returns a boolean if a field has been set.

### GetCs

`func (o *StringMap) GetCs() string`

GetCs returns the Cs field if non-nil, zero value otherwise.

### GetCsOk

`func (o *StringMap) GetCsOk() (*string, bool)`

GetCsOk returns a tuple with the Cs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCs

`func (o *StringMap) SetCs(v string)`

SetCs sets Cs field to given value.

### HasCs

`func (o *StringMap) HasCs() bool`

HasCs returns a boolean if a field has been set.

### GetDa

`func (o *StringMap) GetDa() string`

GetDa returns the Da field if non-nil, zero value otherwise.

### GetDaOk

`func (o *StringMap) GetDaOk() (*string, bool)`

GetDaOk returns a tuple with the Da field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDa

`func (o *StringMap) SetDa(v string)`

SetDa sets Da field to given value.

### HasDa

`func (o *StringMap) HasDa() bool`

HasDa returns a boolean if a field has been set.

### GetNl

`func (o *StringMap) GetNl() string`

GetNl returns the Nl field if non-nil, zero value otherwise.

### GetNlOk

`func (o *StringMap) GetNlOk() (*string, bool)`

GetNlOk returns a tuple with the Nl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNl

`func (o *StringMap) SetNl(v string)`

SetNl sets Nl field to given value.

### HasNl

`func (o *StringMap) HasNl() bool`

HasNl returns a boolean if a field has been set.

### GetEt

`func (o *StringMap) GetEt() string`

GetEt returns the Et field if non-nil, zero value otherwise.

### GetEtOk

`func (o *StringMap) GetEtOk() (*string, bool)`

GetEtOk returns a tuple with the Et field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEt

`func (o *StringMap) SetEt(v string)`

SetEt sets Et field to given value.

### HasEt

`func (o *StringMap) HasEt() bool`

HasEt returns a boolean if a field has been set.

### GetFi

`func (o *StringMap) GetFi() string`

GetFi returns the Fi field if non-nil, zero value otherwise.

### GetFiOk

`func (o *StringMap) GetFiOk() (*string, bool)`

GetFiOk returns a tuple with the Fi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFi

`func (o *StringMap) SetFi(v string)`

SetFi sets Fi field to given value.

### HasFi

`func (o *StringMap) HasFi() bool`

HasFi returns a boolean if a field has been set.

### GetFr

`func (o *StringMap) GetFr() string`

GetFr returns the Fr field if non-nil, zero value otherwise.

### GetFrOk

`func (o *StringMap) GetFrOk() (*string, bool)`

GetFrOk returns a tuple with the Fr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFr

`func (o *StringMap) SetFr(v string)`

SetFr sets Fr field to given value.

### HasFr

`func (o *StringMap) HasFr() bool`

HasFr returns a boolean if a field has been set.

### GetKa

`func (o *StringMap) GetKa() string`

GetKa returns the Ka field if non-nil, zero value otherwise.

### GetKaOk

`func (o *StringMap) GetKaOk() (*string, bool)`

GetKaOk returns a tuple with the Ka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKa

`func (o *StringMap) SetKa(v string)`

SetKa sets Ka field to given value.

### HasKa

`func (o *StringMap) HasKa() bool`

HasKa returns a boolean if a field has been set.

### GetDe

`func (o *StringMap) GetDe() string`

GetDe returns the De field if non-nil, zero value otherwise.

### GetDeOk

`func (o *StringMap) GetDeOk() (*string, bool)`

GetDeOk returns a tuple with the De field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDe

`func (o *StringMap) SetDe(v string)`

SetDe sets De field to given value.

### HasDe

`func (o *StringMap) HasDe() bool`

HasDe returns a boolean if a field has been set.

### GetEl

`func (o *StringMap) GetEl() string`

GetEl returns the El field if non-nil, zero value otherwise.

### GetElOk

`func (o *StringMap) GetElOk() (*string, bool)`

GetElOk returns a tuple with the El field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEl

`func (o *StringMap) SetEl(v string)`

SetEl sets El field to given value.

### HasEl

`func (o *StringMap) HasEl() bool`

HasEl returns a boolean if a field has been set.

### GetHi

`func (o *StringMap) GetHi() string`

GetHi returns the Hi field if non-nil, zero value otherwise.

### GetHiOk

`func (o *StringMap) GetHiOk() (*string, bool)`

GetHiOk returns a tuple with the Hi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHi

`func (o *StringMap) SetHi(v string)`

SetHi sets Hi field to given value.

### HasHi

`func (o *StringMap) HasHi() bool`

HasHi returns a boolean if a field has been set.

### GetHe

`func (o *StringMap) GetHe() string`

GetHe returns the He field if non-nil, zero value otherwise.

### GetHeOk

`func (o *StringMap) GetHeOk() (*string, bool)`

GetHeOk returns a tuple with the He field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHe

`func (o *StringMap) SetHe(v string)`

SetHe sets He field to given value.

### HasHe

`func (o *StringMap) HasHe() bool`

HasHe returns a boolean if a field has been set.

### GetHu

`func (o *StringMap) GetHu() string`

GetHu returns the Hu field if non-nil, zero value otherwise.

### GetHuOk

`func (o *StringMap) GetHuOk() (*string, bool)`

GetHuOk returns a tuple with the Hu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHu

`func (o *StringMap) SetHu(v string)`

SetHu sets Hu field to given value.

### HasHu

`func (o *StringMap) HasHu() bool`

HasHu returns a boolean if a field has been set.

### GetId

`func (o *StringMap) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StringMap) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StringMap) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StringMap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIt

`func (o *StringMap) GetIt() string`

GetIt returns the It field if non-nil, zero value otherwise.

### GetItOk

`func (o *StringMap) GetItOk() (*string, bool)`

GetItOk returns a tuple with the It field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIt

`func (o *StringMap) SetIt(v string)`

SetIt sets It field to given value.

### HasIt

`func (o *StringMap) HasIt() bool`

HasIt returns a boolean if a field has been set.

### GetJa

`func (o *StringMap) GetJa() string`

GetJa returns the Ja field if non-nil, zero value otherwise.

### GetJaOk

`func (o *StringMap) GetJaOk() (*string, bool)`

GetJaOk returns a tuple with the Ja field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJa

`func (o *StringMap) SetJa(v string)`

SetJa sets Ja field to given value.

### HasJa

`func (o *StringMap) HasJa() bool`

HasJa returns a boolean if a field has been set.

### GetKo

`func (o *StringMap) GetKo() string`

GetKo returns the Ko field if non-nil, zero value otherwise.

### GetKoOk

`func (o *StringMap) GetKoOk() (*string, bool)`

GetKoOk returns a tuple with the Ko field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKo

`func (o *StringMap) SetKo(v string)`

SetKo sets Ko field to given value.

### HasKo

`func (o *StringMap) HasKo() bool`

HasKo returns a boolean if a field has been set.

### GetLv

`func (o *StringMap) GetLv() string`

GetLv returns the Lv field if non-nil, zero value otherwise.

### GetLvOk

`func (o *StringMap) GetLvOk() (*string, bool)`

GetLvOk returns a tuple with the Lv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLv

`func (o *StringMap) SetLv(v string)`

SetLv sets Lv field to given value.

### HasLv

`func (o *StringMap) HasLv() bool`

HasLv returns a boolean if a field has been set.

### GetLt

`func (o *StringMap) GetLt() string`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *StringMap) GetLtOk() (*string, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *StringMap) SetLt(v string)`

SetLt sets Lt field to given value.

### HasLt

`func (o *StringMap) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetMs

`func (o *StringMap) GetMs() string`

GetMs returns the Ms field if non-nil, zero value otherwise.

### GetMsOk

`func (o *StringMap) GetMsOk() (*string, bool)`

GetMsOk returns a tuple with the Ms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMs

`func (o *StringMap) SetMs(v string)`

SetMs sets Ms field to given value.

### HasMs

`func (o *StringMap) HasMs() bool`

HasMs returns a boolean if a field has been set.

### GetNb

`func (o *StringMap) GetNb() string`

GetNb returns the Nb field if non-nil, zero value otherwise.

### GetNbOk

`func (o *StringMap) GetNbOk() (*string, bool)`

GetNbOk returns a tuple with the Nb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNb

`func (o *StringMap) SetNb(v string)`

SetNb sets Nb field to given value.

### HasNb

`func (o *StringMap) HasNb() bool`

HasNb returns a boolean if a field has been set.

### GetPl

`func (o *StringMap) GetPl() string`

GetPl returns the Pl field if non-nil, zero value otherwise.

### GetPlOk

`func (o *StringMap) GetPlOk() (*string, bool)`

GetPlOk returns a tuple with the Pl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPl

`func (o *StringMap) SetPl(v string)`

SetPl sets Pl field to given value.

### HasPl

`func (o *StringMap) HasPl() bool`

HasPl returns a boolean if a field has been set.

### GetFa

`func (o *StringMap) GetFa() string`

GetFa returns the Fa field if non-nil, zero value otherwise.

### GetFaOk

`func (o *StringMap) GetFaOk() (*string, bool)`

GetFaOk returns a tuple with the Fa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFa

`func (o *StringMap) SetFa(v string)`

SetFa sets Fa field to given value.

### HasFa

`func (o *StringMap) HasFa() bool`

HasFa returns a boolean if a field has been set.

### GetPt

`func (o *StringMap) GetPt() string`

GetPt returns the Pt field if non-nil, zero value otherwise.

### GetPtOk

`func (o *StringMap) GetPtOk() (*string, bool)`

GetPtOk returns a tuple with the Pt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPt

`func (o *StringMap) SetPt(v string)`

SetPt sets Pt field to given value.

### HasPt

`func (o *StringMap) HasPt() bool`

HasPt returns a boolean if a field has been set.

### GetPa

`func (o *StringMap) GetPa() string`

GetPa returns the Pa field if non-nil, zero value otherwise.

### GetPaOk

`func (o *StringMap) GetPaOk() (*string, bool)`

GetPaOk returns a tuple with the Pa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPa

`func (o *StringMap) SetPa(v string)`

SetPa sets Pa field to given value.

### HasPa

`func (o *StringMap) HasPa() bool`

HasPa returns a boolean if a field has been set.

### GetRo

`func (o *StringMap) GetRo() string`

GetRo returns the Ro field if non-nil, zero value otherwise.

### GetRoOk

`func (o *StringMap) GetRoOk() (*string, bool)`

GetRoOk returns a tuple with the Ro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRo

`func (o *StringMap) SetRo(v string)`

SetRo sets Ro field to given value.

### HasRo

`func (o *StringMap) HasRo() bool`

HasRo returns a boolean if a field has been set.

### GetRu

`func (o *StringMap) GetRu() string`

GetRu returns the Ru field if non-nil, zero value otherwise.

### GetRuOk

`func (o *StringMap) GetRuOk() (*string, bool)`

GetRuOk returns a tuple with the Ru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRu

`func (o *StringMap) SetRu(v string)`

SetRu sets Ru field to given value.

### HasRu

`func (o *StringMap) HasRu() bool`

HasRu returns a boolean if a field has been set.

### GetSr

`func (o *StringMap) GetSr() string`

GetSr returns the Sr field if non-nil, zero value otherwise.

### GetSrOk

`func (o *StringMap) GetSrOk() (*string, bool)`

GetSrOk returns a tuple with the Sr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSr

`func (o *StringMap) SetSr(v string)`

SetSr sets Sr field to given value.

### HasSr

`func (o *StringMap) HasSr() bool`

HasSr returns a boolean if a field has been set.

### GetSk

`func (o *StringMap) GetSk() string`

GetSk returns the Sk field if non-nil, zero value otherwise.

### GetSkOk

`func (o *StringMap) GetSkOk() (*string, bool)`

GetSkOk returns a tuple with the Sk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSk

`func (o *StringMap) SetSk(v string)`

SetSk sets Sk field to given value.

### HasSk

`func (o *StringMap) HasSk() bool`

HasSk returns a boolean if a field has been set.

### GetEs

`func (o *StringMap) GetEs() string`

GetEs returns the Es field if non-nil, zero value otherwise.

### GetEsOk

`func (o *StringMap) GetEsOk() (*string, bool)`

GetEsOk returns a tuple with the Es field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEs

`func (o *StringMap) SetEs(v string)`

SetEs sets Es field to given value.

### HasEs

`func (o *StringMap) HasEs() bool`

HasEs returns a boolean if a field has been set.

### GetSv

`func (o *StringMap) GetSv() string`

GetSv returns the Sv field if non-nil, zero value otherwise.

### GetSvOk

`func (o *StringMap) GetSvOk() (*string, bool)`

GetSvOk returns a tuple with the Sv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSv

`func (o *StringMap) SetSv(v string)`

SetSv sets Sv field to given value.

### HasSv

`func (o *StringMap) HasSv() bool`

HasSv returns a boolean if a field has been set.

### GetTh

`func (o *StringMap) GetTh() string`

GetTh returns the Th field if non-nil, zero value otherwise.

### GetThOk

`func (o *StringMap) GetThOk() (*string, bool)`

GetThOk returns a tuple with the Th field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTh

`func (o *StringMap) SetTh(v string)`

SetTh sets Th field to given value.

### HasTh

`func (o *StringMap) HasTh() bool`

HasTh returns a boolean if a field has been set.

### GetTr

`func (o *StringMap) GetTr() string`

GetTr returns the Tr field if non-nil, zero value otherwise.

### GetTrOk

`func (o *StringMap) GetTrOk() (*string, bool)`

GetTrOk returns a tuple with the Tr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTr

`func (o *StringMap) SetTr(v string)`

SetTr sets Tr field to given value.

### HasTr

`func (o *StringMap) HasTr() bool`

HasTr returns a boolean if a field has been set.

### GetUk

`func (o *StringMap) GetUk() string`

GetUk returns the Uk field if non-nil, zero value otherwise.

### GetUkOk

`func (o *StringMap) GetUkOk() (*string, bool)`

GetUkOk returns a tuple with the Uk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUk

`func (o *StringMap) SetUk(v string)`

SetUk sets Uk field to given value.

### HasUk

`func (o *StringMap) HasUk() bool`

HasUk returns a boolean if a field has been set.

### GetVi

`func (o *StringMap) GetVi() string`

GetVi returns the Vi field if non-nil, zero value otherwise.

### GetViOk

`func (o *StringMap) GetViOk() (*string, bool)`

GetViOk returns a tuple with the Vi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVi

`func (o *StringMap) SetVi(v string)`

SetVi sets Vi field to given value.

### HasVi

`func (o *StringMap) HasVi() bool`

HasVi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



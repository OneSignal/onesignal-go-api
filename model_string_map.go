/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.4.0
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// StringMap struct for StringMap
type StringMap struct {
	// Text in English.  Will be used as a fallback
	En *string `json:"en,omitempty"`
	// Text in Arabic.
	Ar *string `json:"ar,omitempty"`
	// Text in Bosnian.
	Bs *string `json:"bs,omitempty"`
	// Text in Bulgarian.
	Bg *string `json:"bg,omitempty"`
	// Text in Catalan.
	Ca *string `json:"ca,omitempty"`
	// Text in Chinese (Simplified).
	ZhHans *string `json:"zh-Hans,omitempty"`
	// Text in Chinese (Traditional).
	ZhHant *string `json:"zh-Hant,omitempty"`
	// Alias for zh-Hans.
	Zh *string `json:"zh,omitempty"`
	// Text in Croatian.
	Hr *string `json:"hr,omitempty"`
	// Text in Czech.
	Cs *string `json:"cs,omitempty"`
	// Text in Danish.
	Da *string `json:"da,omitempty"`
	// Text in Dutch.
	Nl *string `json:"nl,omitempty"`
	// Text in Estonian.
	Et *string `json:"et,omitempty"`
	// Text in Finnish.
	Fi *string `json:"fi,omitempty"`
	// Text in French.
	Fr *string `json:"fr,omitempty"`
	// Text in Georgian.
	Ka *string `json:"ka,omitempty"`
	// Text in German.
	De *string `json:"de,omitempty"`
	// Text in Greek.
	El *string `json:"el,omitempty"`
	// Text in Hindi.
	Hi *string `json:"hi,omitempty"`
	// Text in Hebrew.
	He *string `json:"he,omitempty"`
	// Text in Hungarian.
	Hu *string `json:"hu,omitempty"`
	// Text in Indonesian.
	Id *string `json:"id,omitempty"`
	// Text in Italian.
	It *string `json:"it,omitempty"`
	// Text in Japanese.
	Ja *string `json:"ja,omitempty"`
	// Text in Korean.
	Ko *string `json:"ko,omitempty"`
	// Text in Latvian.
	Lv *string `json:"lv,omitempty"`
	// Text in Lithuanian.
	Lt *string `json:"lt,omitempty"`
	// Text in Malay.
	Ms *string `json:"ms,omitempty"`
	// Text in Norwegian.
	Nb *string `json:"nb,omitempty"`
	// Text in Polish.
	Pl *string `json:"pl,omitempty"`
	// Text in Persian.
	Fa *string `json:"fa,omitempty"`
	// Text in Portugese.
	Pt *string `json:"pt,omitempty"`
	// Text in Punjabi.
	Pa *string `json:"pa,omitempty"`
	// Text in Romanian.
	Ro *string `json:"ro,omitempty"`
	// Text in Russian.
	Ru *string `json:"ru,omitempty"`
	// Text in Serbian.
	Sr *string `json:"sr,omitempty"`
	// Text in Slovak.
	Sk *string `json:"sk,omitempty"`
	// Text in Spanish.
	Es *string `json:"es,omitempty"`
	// Text in Swedish.
	Sv *string `json:"sv,omitempty"`
	// Text in Thai.
	Th *string `json:"th,omitempty"`
	// Text in Turkish.
	Tr *string `json:"tr,omitempty"`
	// Text in Ukrainian.
	Uk *string `json:"uk,omitempty"`
	// Text in Vietnamese.
	Vi *string `json:"vi,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StringMap StringMap

// NewStringMap instantiates a new StringMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringMap() *StringMap {
	this := StringMap{}
	return &this
}

// NewStringMapWithDefaults instantiates a new StringMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringMapWithDefaults() *StringMap {
	this := StringMap{}
	return &this
}

// GetEn returns the En field value if set, zero value otherwise.
func (o *StringMap) GetEn() string {
	if o == nil || o.En == nil {
		var ret string
		return ret
	}
	return *o.En
}

// GetEnOk returns a tuple with the En field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetEnOk() (*string, bool) {
	if o == nil || o.En == nil {
		return nil, false
	}
	return o.En, true
}

// HasEn returns a boolean if a field has been set.
func (o *StringMap) HasEn() bool {
	if o != nil && o.En != nil {
		return true
	}

	return false
}

// SetEn gets a reference to the given string and assigns it to the En field.
func (o *StringMap) SetEn(v string) {
	o.En = &v
}

// GetAr returns the Ar field value if set, zero value otherwise.
func (o *StringMap) GetAr() string {
	if o == nil || o.Ar == nil {
		var ret string
		return ret
	}
	return *o.Ar
}

// GetArOk returns a tuple with the Ar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetArOk() (*string, bool) {
	if o == nil || o.Ar == nil {
		return nil, false
	}
	return o.Ar, true
}

// HasAr returns a boolean if a field has been set.
func (o *StringMap) HasAr() bool {
	if o != nil && o.Ar != nil {
		return true
	}

	return false
}

// SetAr gets a reference to the given string and assigns it to the Ar field.
func (o *StringMap) SetAr(v string) {
	o.Ar = &v
}

// GetBs returns the Bs field value if set, zero value otherwise.
func (o *StringMap) GetBs() string {
	if o == nil || o.Bs == nil {
		var ret string
		return ret
	}
	return *o.Bs
}

// GetBsOk returns a tuple with the Bs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetBsOk() (*string, bool) {
	if o == nil || o.Bs == nil {
		return nil, false
	}
	return o.Bs, true
}

// HasBs returns a boolean if a field has been set.
func (o *StringMap) HasBs() bool {
	if o != nil && o.Bs != nil {
		return true
	}

	return false
}

// SetBs gets a reference to the given string and assigns it to the Bs field.
func (o *StringMap) SetBs(v string) {
	o.Bs = &v
}

// GetBg returns the Bg field value if set, zero value otherwise.
func (o *StringMap) GetBg() string {
	if o == nil || o.Bg == nil {
		var ret string
		return ret
	}
	return *o.Bg
}

// GetBgOk returns a tuple with the Bg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetBgOk() (*string, bool) {
	if o == nil || o.Bg == nil {
		return nil, false
	}
	return o.Bg, true
}

// HasBg returns a boolean if a field has been set.
func (o *StringMap) HasBg() bool {
	if o != nil && o.Bg != nil {
		return true
	}

	return false
}

// SetBg gets a reference to the given string and assigns it to the Bg field.
func (o *StringMap) SetBg(v string) {
	o.Bg = &v
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *StringMap) GetCa() string {
	if o == nil || o.Ca == nil {
		var ret string
		return ret
	}
	return *o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetCaOk() (*string, bool) {
	if o == nil || o.Ca == nil {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *StringMap) HasCa() bool {
	if o != nil && o.Ca != nil {
		return true
	}

	return false
}

// SetCa gets a reference to the given string and assigns it to the Ca field.
func (o *StringMap) SetCa(v string) {
	o.Ca = &v
}

// GetZhHans returns the ZhHans field value if set, zero value otherwise.
func (o *StringMap) GetZhHans() string {
	if o == nil || o.ZhHans == nil {
		var ret string
		return ret
	}
	return *o.ZhHans
}

// GetZhHansOk returns a tuple with the ZhHans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetZhHansOk() (*string, bool) {
	if o == nil || o.ZhHans == nil {
		return nil, false
	}
	return o.ZhHans, true
}

// HasZhHans returns a boolean if a field has been set.
func (o *StringMap) HasZhHans() bool {
	if o != nil && o.ZhHans != nil {
		return true
	}

	return false
}

// SetZhHans gets a reference to the given string and assigns it to the ZhHans field.
func (o *StringMap) SetZhHans(v string) {
	o.ZhHans = &v
}

// GetZhHant returns the ZhHant field value if set, zero value otherwise.
func (o *StringMap) GetZhHant() string {
	if o == nil || o.ZhHant == nil {
		var ret string
		return ret
	}
	return *o.ZhHant
}

// GetZhHantOk returns a tuple with the ZhHant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetZhHantOk() (*string, bool) {
	if o == nil || o.ZhHant == nil {
		return nil, false
	}
	return o.ZhHant, true
}

// HasZhHant returns a boolean if a field has been set.
func (o *StringMap) HasZhHant() bool {
	if o != nil && o.ZhHant != nil {
		return true
	}

	return false
}

// SetZhHant gets a reference to the given string and assigns it to the ZhHant field.
func (o *StringMap) SetZhHant(v string) {
	o.ZhHant = &v
}

// GetZh returns the Zh field value if set, zero value otherwise.
func (o *StringMap) GetZh() string {
	if o == nil || o.Zh == nil {
		var ret string
		return ret
	}
	return *o.Zh
}

// GetZhOk returns a tuple with the Zh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetZhOk() (*string, bool) {
	if o == nil || o.Zh == nil {
		return nil, false
	}
	return o.Zh, true
}

// HasZh returns a boolean if a field has been set.
func (o *StringMap) HasZh() bool {
	if o != nil && o.Zh != nil {
		return true
	}

	return false
}

// SetZh gets a reference to the given string and assigns it to the Zh field.
func (o *StringMap) SetZh(v string) {
	o.Zh = &v
}

// GetHr returns the Hr field value if set, zero value otherwise.
func (o *StringMap) GetHr() string {
	if o == nil || o.Hr == nil {
		var ret string
		return ret
	}
	return *o.Hr
}

// GetHrOk returns a tuple with the Hr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetHrOk() (*string, bool) {
	if o == nil || o.Hr == nil {
		return nil, false
	}
	return o.Hr, true
}

// HasHr returns a boolean if a field has been set.
func (o *StringMap) HasHr() bool {
	if o != nil && o.Hr != nil {
		return true
	}

	return false
}

// SetHr gets a reference to the given string and assigns it to the Hr field.
func (o *StringMap) SetHr(v string) {
	o.Hr = &v
}

// GetCs returns the Cs field value if set, zero value otherwise.
func (o *StringMap) GetCs() string {
	if o == nil || o.Cs == nil {
		var ret string
		return ret
	}
	return *o.Cs
}

// GetCsOk returns a tuple with the Cs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetCsOk() (*string, bool) {
	if o == nil || o.Cs == nil {
		return nil, false
	}
	return o.Cs, true
}

// HasCs returns a boolean if a field has been set.
func (o *StringMap) HasCs() bool {
	if o != nil && o.Cs != nil {
		return true
	}

	return false
}

// SetCs gets a reference to the given string and assigns it to the Cs field.
func (o *StringMap) SetCs(v string) {
	o.Cs = &v
}

// GetDa returns the Da field value if set, zero value otherwise.
func (o *StringMap) GetDa() string {
	if o == nil || o.Da == nil {
		var ret string
		return ret
	}
	return *o.Da
}

// GetDaOk returns a tuple with the Da field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetDaOk() (*string, bool) {
	if o == nil || o.Da == nil {
		return nil, false
	}
	return o.Da, true
}

// HasDa returns a boolean if a field has been set.
func (o *StringMap) HasDa() bool {
	if o != nil && o.Da != nil {
		return true
	}

	return false
}

// SetDa gets a reference to the given string and assigns it to the Da field.
func (o *StringMap) SetDa(v string) {
	o.Da = &v
}

// GetNl returns the Nl field value if set, zero value otherwise.
func (o *StringMap) GetNl() string {
	if o == nil || o.Nl == nil {
		var ret string
		return ret
	}
	return *o.Nl
}

// GetNlOk returns a tuple with the Nl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetNlOk() (*string, bool) {
	if o == nil || o.Nl == nil {
		return nil, false
	}
	return o.Nl, true
}

// HasNl returns a boolean if a field has been set.
func (o *StringMap) HasNl() bool {
	if o != nil && o.Nl != nil {
		return true
	}

	return false
}

// SetNl gets a reference to the given string and assigns it to the Nl field.
func (o *StringMap) SetNl(v string) {
	o.Nl = &v
}

// GetEt returns the Et field value if set, zero value otherwise.
func (o *StringMap) GetEt() string {
	if o == nil || o.Et == nil {
		var ret string
		return ret
	}
	return *o.Et
}

// GetEtOk returns a tuple with the Et field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetEtOk() (*string, bool) {
	if o == nil || o.Et == nil {
		return nil, false
	}
	return o.Et, true
}

// HasEt returns a boolean if a field has been set.
func (o *StringMap) HasEt() bool {
	if o != nil && o.Et != nil {
		return true
	}

	return false
}

// SetEt gets a reference to the given string and assigns it to the Et field.
func (o *StringMap) SetEt(v string) {
	o.Et = &v
}

// GetFi returns the Fi field value if set, zero value otherwise.
func (o *StringMap) GetFi() string {
	if o == nil || o.Fi == nil {
		var ret string
		return ret
	}
	return *o.Fi
}

// GetFiOk returns a tuple with the Fi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetFiOk() (*string, bool) {
	if o == nil || o.Fi == nil {
		return nil, false
	}
	return o.Fi, true
}

// HasFi returns a boolean if a field has been set.
func (o *StringMap) HasFi() bool {
	if o != nil && o.Fi != nil {
		return true
	}

	return false
}

// SetFi gets a reference to the given string and assigns it to the Fi field.
func (o *StringMap) SetFi(v string) {
	o.Fi = &v
}

// GetFr returns the Fr field value if set, zero value otherwise.
func (o *StringMap) GetFr() string {
	if o == nil || o.Fr == nil {
		var ret string
		return ret
	}
	return *o.Fr
}

// GetFrOk returns a tuple with the Fr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetFrOk() (*string, bool) {
	if o == nil || o.Fr == nil {
		return nil, false
	}
	return o.Fr, true
}

// HasFr returns a boolean if a field has been set.
func (o *StringMap) HasFr() bool {
	if o != nil && o.Fr != nil {
		return true
	}

	return false
}

// SetFr gets a reference to the given string and assigns it to the Fr field.
func (o *StringMap) SetFr(v string) {
	o.Fr = &v
}

// GetKa returns the Ka field value if set, zero value otherwise.
func (o *StringMap) GetKa() string {
	if o == nil || o.Ka == nil {
		var ret string
		return ret
	}
	return *o.Ka
}

// GetKaOk returns a tuple with the Ka field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetKaOk() (*string, bool) {
	if o == nil || o.Ka == nil {
		return nil, false
	}
	return o.Ka, true
}

// HasKa returns a boolean if a field has been set.
func (o *StringMap) HasKa() bool {
	if o != nil && o.Ka != nil {
		return true
	}

	return false
}

// SetKa gets a reference to the given string and assigns it to the Ka field.
func (o *StringMap) SetKa(v string) {
	o.Ka = &v
}

// GetDe returns the De field value if set, zero value otherwise.
func (o *StringMap) GetDe() string {
	if o == nil || o.De == nil {
		var ret string
		return ret
	}
	return *o.De
}

// GetDeOk returns a tuple with the De field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetDeOk() (*string, bool) {
	if o == nil || o.De == nil {
		return nil, false
	}
	return o.De, true
}

// HasDe returns a boolean if a field has been set.
func (o *StringMap) HasDe() bool {
	if o != nil && o.De != nil {
		return true
	}

	return false
}

// SetDe gets a reference to the given string and assigns it to the De field.
func (o *StringMap) SetDe(v string) {
	o.De = &v
}

// GetEl returns the El field value if set, zero value otherwise.
func (o *StringMap) GetEl() string {
	if o == nil || o.El == nil {
		var ret string
		return ret
	}
	return *o.El
}

// GetElOk returns a tuple with the El field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetElOk() (*string, bool) {
	if o == nil || o.El == nil {
		return nil, false
	}
	return o.El, true
}

// HasEl returns a boolean if a field has been set.
func (o *StringMap) HasEl() bool {
	if o != nil && o.El != nil {
		return true
	}

	return false
}

// SetEl gets a reference to the given string and assigns it to the El field.
func (o *StringMap) SetEl(v string) {
	o.El = &v
}

// GetHi returns the Hi field value if set, zero value otherwise.
func (o *StringMap) GetHi() string {
	if o == nil || o.Hi == nil {
		var ret string
		return ret
	}
	return *o.Hi
}

// GetHiOk returns a tuple with the Hi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetHiOk() (*string, bool) {
	if o == nil || o.Hi == nil {
		return nil, false
	}
	return o.Hi, true
}

// HasHi returns a boolean if a field has been set.
func (o *StringMap) HasHi() bool {
	if o != nil && o.Hi != nil {
		return true
	}

	return false
}

// SetHi gets a reference to the given string and assigns it to the Hi field.
func (o *StringMap) SetHi(v string) {
	o.Hi = &v
}

// GetHe returns the He field value if set, zero value otherwise.
func (o *StringMap) GetHe() string {
	if o == nil || o.He == nil {
		var ret string
		return ret
	}
	return *o.He
}

// GetHeOk returns a tuple with the He field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetHeOk() (*string, bool) {
	if o == nil || o.He == nil {
		return nil, false
	}
	return o.He, true
}

// HasHe returns a boolean if a field has been set.
func (o *StringMap) HasHe() bool {
	if o != nil && o.He != nil {
		return true
	}

	return false
}

// SetHe gets a reference to the given string and assigns it to the He field.
func (o *StringMap) SetHe(v string) {
	o.He = &v
}

// GetHu returns the Hu field value if set, zero value otherwise.
func (o *StringMap) GetHu() string {
	if o == nil || o.Hu == nil {
		var ret string
		return ret
	}
	return *o.Hu
}

// GetHuOk returns a tuple with the Hu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetHuOk() (*string, bool) {
	if o == nil || o.Hu == nil {
		return nil, false
	}
	return o.Hu, true
}

// HasHu returns a boolean if a field has been set.
func (o *StringMap) HasHu() bool {
	if o != nil && o.Hu != nil {
		return true
	}

	return false
}

// SetHu gets a reference to the given string and assigns it to the Hu field.
func (o *StringMap) SetHu(v string) {
	o.Hu = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StringMap) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StringMap) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StringMap) SetId(v string) {
	o.Id = &v
}

// GetIt returns the It field value if set, zero value otherwise.
func (o *StringMap) GetIt() string {
	if o == nil || o.It == nil {
		var ret string
		return ret
	}
	return *o.It
}

// GetItOk returns a tuple with the It field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetItOk() (*string, bool) {
	if o == nil || o.It == nil {
		return nil, false
	}
	return o.It, true
}

// HasIt returns a boolean if a field has been set.
func (o *StringMap) HasIt() bool {
	if o != nil && o.It != nil {
		return true
	}

	return false
}

// SetIt gets a reference to the given string and assigns it to the It field.
func (o *StringMap) SetIt(v string) {
	o.It = &v
}

// GetJa returns the Ja field value if set, zero value otherwise.
func (o *StringMap) GetJa() string {
	if o == nil || o.Ja == nil {
		var ret string
		return ret
	}
	return *o.Ja
}

// GetJaOk returns a tuple with the Ja field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetJaOk() (*string, bool) {
	if o == nil || o.Ja == nil {
		return nil, false
	}
	return o.Ja, true
}

// HasJa returns a boolean if a field has been set.
func (o *StringMap) HasJa() bool {
	if o != nil && o.Ja != nil {
		return true
	}

	return false
}

// SetJa gets a reference to the given string and assigns it to the Ja field.
func (o *StringMap) SetJa(v string) {
	o.Ja = &v
}

// GetKo returns the Ko field value if set, zero value otherwise.
func (o *StringMap) GetKo() string {
	if o == nil || o.Ko == nil {
		var ret string
		return ret
	}
	return *o.Ko
}

// GetKoOk returns a tuple with the Ko field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetKoOk() (*string, bool) {
	if o == nil || o.Ko == nil {
		return nil, false
	}
	return o.Ko, true
}

// HasKo returns a boolean if a field has been set.
func (o *StringMap) HasKo() bool {
	if o != nil && o.Ko != nil {
		return true
	}

	return false
}

// SetKo gets a reference to the given string and assigns it to the Ko field.
func (o *StringMap) SetKo(v string) {
	o.Ko = &v
}

// GetLv returns the Lv field value if set, zero value otherwise.
func (o *StringMap) GetLv() string {
	if o == nil || o.Lv == nil {
		var ret string
		return ret
	}
	return *o.Lv
}

// GetLvOk returns a tuple with the Lv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetLvOk() (*string, bool) {
	if o == nil || o.Lv == nil {
		return nil, false
	}
	return o.Lv, true
}

// HasLv returns a boolean if a field has been set.
func (o *StringMap) HasLv() bool {
	if o != nil && o.Lv != nil {
		return true
	}

	return false
}

// SetLv gets a reference to the given string and assigns it to the Lv field.
func (o *StringMap) SetLv(v string) {
	o.Lv = &v
}

// GetLt returns the Lt field value if set, zero value otherwise.
func (o *StringMap) GetLt() string {
	if o == nil || o.Lt == nil {
		var ret string
		return ret
	}
	return *o.Lt
}

// GetLtOk returns a tuple with the Lt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetLtOk() (*string, bool) {
	if o == nil || o.Lt == nil {
		return nil, false
	}
	return o.Lt, true
}

// HasLt returns a boolean if a field has been set.
func (o *StringMap) HasLt() bool {
	if o != nil && o.Lt != nil {
		return true
	}

	return false
}

// SetLt gets a reference to the given string and assigns it to the Lt field.
func (o *StringMap) SetLt(v string) {
	o.Lt = &v
}

// GetMs returns the Ms field value if set, zero value otherwise.
func (o *StringMap) GetMs() string {
	if o == nil || o.Ms == nil {
		var ret string
		return ret
	}
	return *o.Ms
}

// GetMsOk returns a tuple with the Ms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetMsOk() (*string, bool) {
	if o == nil || o.Ms == nil {
		return nil, false
	}
	return o.Ms, true
}

// HasMs returns a boolean if a field has been set.
func (o *StringMap) HasMs() bool {
	if o != nil && o.Ms != nil {
		return true
	}

	return false
}

// SetMs gets a reference to the given string and assigns it to the Ms field.
func (o *StringMap) SetMs(v string) {
	o.Ms = &v
}

// GetNb returns the Nb field value if set, zero value otherwise.
func (o *StringMap) GetNb() string {
	if o == nil || o.Nb == nil {
		var ret string
		return ret
	}
	return *o.Nb
}

// GetNbOk returns a tuple with the Nb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetNbOk() (*string, bool) {
	if o == nil || o.Nb == nil {
		return nil, false
	}
	return o.Nb, true
}

// HasNb returns a boolean if a field has been set.
func (o *StringMap) HasNb() bool {
	if o != nil && o.Nb != nil {
		return true
	}

	return false
}

// SetNb gets a reference to the given string and assigns it to the Nb field.
func (o *StringMap) SetNb(v string) {
	o.Nb = &v
}

// GetPl returns the Pl field value if set, zero value otherwise.
func (o *StringMap) GetPl() string {
	if o == nil || o.Pl == nil {
		var ret string
		return ret
	}
	return *o.Pl
}

// GetPlOk returns a tuple with the Pl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetPlOk() (*string, bool) {
	if o == nil || o.Pl == nil {
		return nil, false
	}
	return o.Pl, true
}

// HasPl returns a boolean if a field has been set.
func (o *StringMap) HasPl() bool {
	if o != nil && o.Pl != nil {
		return true
	}

	return false
}

// SetPl gets a reference to the given string and assigns it to the Pl field.
func (o *StringMap) SetPl(v string) {
	o.Pl = &v
}

// GetFa returns the Fa field value if set, zero value otherwise.
func (o *StringMap) GetFa() string {
	if o == nil || o.Fa == nil {
		var ret string
		return ret
	}
	return *o.Fa
}

// GetFaOk returns a tuple with the Fa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetFaOk() (*string, bool) {
	if o == nil || o.Fa == nil {
		return nil, false
	}
	return o.Fa, true
}

// HasFa returns a boolean if a field has been set.
func (o *StringMap) HasFa() bool {
	if o != nil && o.Fa != nil {
		return true
	}

	return false
}

// SetFa gets a reference to the given string and assigns it to the Fa field.
func (o *StringMap) SetFa(v string) {
	o.Fa = &v
}

// GetPt returns the Pt field value if set, zero value otherwise.
func (o *StringMap) GetPt() string {
	if o == nil || o.Pt == nil {
		var ret string
		return ret
	}
	return *o.Pt
}

// GetPtOk returns a tuple with the Pt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetPtOk() (*string, bool) {
	if o == nil || o.Pt == nil {
		return nil, false
	}
	return o.Pt, true
}

// HasPt returns a boolean if a field has been set.
func (o *StringMap) HasPt() bool {
	if o != nil && o.Pt != nil {
		return true
	}

	return false
}

// SetPt gets a reference to the given string and assigns it to the Pt field.
func (o *StringMap) SetPt(v string) {
	o.Pt = &v
}

// GetPa returns the Pa field value if set, zero value otherwise.
func (o *StringMap) GetPa() string {
	if o == nil || o.Pa == nil {
		var ret string
		return ret
	}
	return *o.Pa
}

// GetPaOk returns a tuple with the Pa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetPaOk() (*string, bool) {
	if o == nil || o.Pa == nil {
		return nil, false
	}
	return o.Pa, true
}

// HasPa returns a boolean if a field has been set.
func (o *StringMap) HasPa() bool {
	if o != nil && o.Pa != nil {
		return true
	}

	return false
}

// SetPa gets a reference to the given string and assigns it to the Pa field.
func (o *StringMap) SetPa(v string) {
	o.Pa = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *StringMap) GetRo() string {
	if o == nil || o.Ro == nil {
		var ret string
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetRoOk() (*string, bool) {
	if o == nil || o.Ro == nil {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *StringMap) HasRo() bool {
	if o != nil && o.Ro != nil {
		return true
	}

	return false
}

// SetRo gets a reference to the given string and assigns it to the Ro field.
func (o *StringMap) SetRo(v string) {
	o.Ro = &v
}

// GetRu returns the Ru field value if set, zero value otherwise.
func (o *StringMap) GetRu() string {
	if o == nil || o.Ru == nil {
		var ret string
		return ret
	}
	return *o.Ru
}

// GetRuOk returns a tuple with the Ru field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetRuOk() (*string, bool) {
	if o == nil || o.Ru == nil {
		return nil, false
	}
	return o.Ru, true
}

// HasRu returns a boolean if a field has been set.
func (o *StringMap) HasRu() bool {
	if o != nil && o.Ru != nil {
		return true
	}

	return false
}

// SetRu gets a reference to the given string and assigns it to the Ru field.
func (o *StringMap) SetRu(v string) {
	o.Ru = &v
}

// GetSr returns the Sr field value if set, zero value otherwise.
func (o *StringMap) GetSr() string {
	if o == nil || o.Sr == nil {
		var ret string
		return ret
	}
	return *o.Sr
}

// GetSrOk returns a tuple with the Sr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetSrOk() (*string, bool) {
	if o == nil || o.Sr == nil {
		return nil, false
	}
	return o.Sr, true
}

// HasSr returns a boolean if a field has been set.
func (o *StringMap) HasSr() bool {
	if o != nil && o.Sr != nil {
		return true
	}

	return false
}

// SetSr gets a reference to the given string and assigns it to the Sr field.
func (o *StringMap) SetSr(v string) {
	o.Sr = &v
}

// GetSk returns the Sk field value if set, zero value otherwise.
func (o *StringMap) GetSk() string {
	if o == nil || o.Sk == nil {
		var ret string
		return ret
	}
	return *o.Sk
}

// GetSkOk returns a tuple with the Sk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetSkOk() (*string, bool) {
	if o == nil || o.Sk == nil {
		return nil, false
	}
	return o.Sk, true
}

// HasSk returns a boolean if a field has been set.
func (o *StringMap) HasSk() bool {
	if o != nil && o.Sk != nil {
		return true
	}

	return false
}

// SetSk gets a reference to the given string and assigns it to the Sk field.
func (o *StringMap) SetSk(v string) {
	o.Sk = &v
}

// GetEs returns the Es field value if set, zero value otherwise.
func (o *StringMap) GetEs() string {
	if o == nil || o.Es == nil {
		var ret string
		return ret
	}
	return *o.Es
}

// GetEsOk returns a tuple with the Es field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetEsOk() (*string, bool) {
	if o == nil || o.Es == nil {
		return nil, false
	}
	return o.Es, true
}

// HasEs returns a boolean if a field has been set.
func (o *StringMap) HasEs() bool {
	if o != nil && o.Es != nil {
		return true
	}

	return false
}

// SetEs gets a reference to the given string and assigns it to the Es field.
func (o *StringMap) SetEs(v string) {
	o.Es = &v
}

// GetSv returns the Sv field value if set, zero value otherwise.
func (o *StringMap) GetSv() string {
	if o == nil || o.Sv == nil {
		var ret string
		return ret
	}
	return *o.Sv
}

// GetSvOk returns a tuple with the Sv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetSvOk() (*string, bool) {
	if o == nil || o.Sv == nil {
		return nil, false
	}
	return o.Sv, true
}

// HasSv returns a boolean if a field has been set.
func (o *StringMap) HasSv() bool {
	if o != nil && o.Sv != nil {
		return true
	}

	return false
}

// SetSv gets a reference to the given string and assigns it to the Sv field.
func (o *StringMap) SetSv(v string) {
	o.Sv = &v
}

// GetTh returns the Th field value if set, zero value otherwise.
func (o *StringMap) GetTh() string {
	if o == nil || o.Th == nil {
		var ret string
		return ret
	}
	return *o.Th
}

// GetThOk returns a tuple with the Th field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetThOk() (*string, bool) {
	if o == nil || o.Th == nil {
		return nil, false
	}
	return o.Th, true
}

// HasTh returns a boolean if a field has been set.
func (o *StringMap) HasTh() bool {
	if o != nil && o.Th != nil {
		return true
	}

	return false
}

// SetTh gets a reference to the given string and assigns it to the Th field.
func (o *StringMap) SetTh(v string) {
	o.Th = &v
}

// GetTr returns the Tr field value if set, zero value otherwise.
func (o *StringMap) GetTr() string {
	if o == nil || o.Tr == nil {
		var ret string
		return ret
	}
	return *o.Tr
}

// GetTrOk returns a tuple with the Tr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetTrOk() (*string, bool) {
	if o == nil || o.Tr == nil {
		return nil, false
	}
	return o.Tr, true
}

// HasTr returns a boolean if a field has been set.
func (o *StringMap) HasTr() bool {
	if o != nil && o.Tr != nil {
		return true
	}

	return false
}

// SetTr gets a reference to the given string and assigns it to the Tr field.
func (o *StringMap) SetTr(v string) {
	o.Tr = &v
}

// GetUk returns the Uk field value if set, zero value otherwise.
func (o *StringMap) GetUk() string {
	if o == nil || o.Uk == nil {
		var ret string
		return ret
	}
	return *o.Uk
}

// GetUkOk returns a tuple with the Uk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetUkOk() (*string, bool) {
	if o == nil || o.Uk == nil {
		return nil, false
	}
	return o.Uk, true
}

// HasUk returns a boolean if a field has been set.
func (o *StringMap) HasUk() bool {
	if o != nil && o.Uk != nil {
		return true
	}

	return false
}

// SetUk gets a reference to the given string and assigns it to the Uk field.
func (o *StringMap) SetUk(v string) {
	o.Uk = &v
}

// GetVi returns the Vi field value if set, zero value otherwise.
func (o *StringMap) GetVi() string {
	if o == nil || o.Vi == nil {
		var ret string
		return ret
	}
	return *o.Vi
}

// GetViOk returns a tuple with the Vi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringMap) GetViOk() (*string, bool) {
	if o == nil || o.Vi == nil {
		return nil, false
	}
	return o.Vi, true
}

// HasVi returns a boolean if a field has been set.
func (o *StringMap) HasVi() bool {
	if o != nil && o.Vi != nil {
		return true
	}

	return false
}

// SetVi gets a reference to the given string and assigns it to the Vi field.
func (o *StringMap) SetVi(v string) {
	o.Vi = &v
}

func (o StringMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.En != nil {
		toSerialize["en"] = o.En
	}
	if o.Ar != nil {
		toSerialize["ar"] = o.Ar
	}
	if o.Bs != nil {
		toSerialize["bs"] = o.Bs
	}
	if o.Bg != nil {
		toSerialize["bg"] = o.Bg
	}
	if o.Ca != nil {
		toSerialize["ca"] = o.Ca
	}
	if o.ZhHans != nil {
		toSerialize["zh-Hans"] = o.ZhHans
	}
	if o.ZhHant != nil {
		toSerialize["zh-Hant"] = o.ZhHant
	}
	if o.Zh != nil {
		toSerialize["zh"] = o.Zh
	}
	if o.Hr != nil {
		toSerialize["hr"] = o.Hr
	}
	if o.Cs != nil {
		toSerialize["cs"] = o.Cs
	}
	if o.Da != nil {
		toSerialize["da"] = o.Da
	}
	if o.Nl != nil {
		toSerialize["nl"] = o.Nl
	}
	if o.Et != nil {
		toSerialize["et"] = o.Et
	}
	if o.Fi != nil {
		toSerialize["fi"] = o.Fi
	}
	if o.Fr != nil {
		toSerialize["fr"] = o.Fr
	}
	if o.Ka != nil {
		toSerialize["ka"] = o.Ka
	}
	if o.De != nil {
		toSerialize["de"] = o.De
	}
	if o.El != nil {
		toSerialize["el"] = o.El
	}
	if o.Hi != nil {
		toSerialize["hi"] = o.Hi
	}
	if o.He != nil {
		toSerialize["he"] = o.He
	}
	if o.Hu != nil {
		toSerialize["hu"] = o.Hu
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.It != nil {
		toSerialize["it"] = o.It
	}
	if o.Ja != nil {
		toSerialize["ja"] = o.Ja
	}
	if o.Ko != nil {
		toSerialize["ko"] = o.Ko
	}
	if o.Lv != nil {
		toSerialize["lv"] = o.Lv
	}
	if o.Lt != nil {
		toSerialize["lt"] = o.Lt
	}
	if o.Ms != nil {
		toSerialize["ms"] = o.Ms
	}
	if o.Nb != nil {
		toSerialize["nb"] = o.Nb
	}
	if o.Pl != nil {
		toSerialize["pl"] = o.Pl
	}
	if o.Fa != nil {
		toSerialize["fa"] = o.Fa
	}
	if o.Pt != nil {
		toSerialize["pt"] = o.Pt
	}
	if o.Pa != nil {
		toSerialize["pa"] = o.Pa
	}
	if o.Ro != nil {
		toSerialize["ro"] = o.Ro
	}
	if o.Ru != nil {
		toSerialize["ru"] = o.Ru
	}
	if o.Sr != nil {
		toSerialize["sr"] = o.Sr
	}
	if o.Sk != nil {
		toSerialize["sk"] = o.Sk
	}
	if o.Es != nil {
		toSerialize["es"] = o.Es
	}
	if o.Sv != nil {
		toSerialize["sv"] = o.Sv
	}
	if o.Th != nil {
		toSerialize["th"] = o.Th
	}
	if o.Tr != nil {
		toSerialize["tr"] = o.Tr
	}
	if o.Uk != nil {
		toSerialize["uk"] = o.Uk
	}
	if o.Vi != nil {
		toSerialize["vi"] = o.Vi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StringMap) UnmarshalJSON(bytes []byte) (err error) {
	varStringMap := _StringMap{}

	if err = json.Unmarshal(bytes, &varStringMap); err == nil {
		*o = StringMap(varStringMap)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "en")
		delete(additionalProperties, "ar")
		delete(additionalProperties, "bs")
		delete(additionalProperties, "bg")
		delete(additionalProperties, "ca")
		delete(additionalProperties, "zh-Hans")
		delete(additionalProperties, "zh-Hant")
		delete(additionalProperties, "zh")
		delete(additionalProperties, "hr")
		delete(additionalProperties, "cs")
		delete(additionalProperties, "da")
		delete(additionalProperties, "nl")
		delete(additionalProperties, "et")
		delete(additionalProperties, "fi")
		delete(additionalProperties, "fr")
		delete(additionalProperties, "ka")
		delete(additionalProperties, "de")
		delete(additionalProperties, "el")
		delete(additionalProperties, "hi")
		delete(additionalProperties, "he")
		delete(additionalProperties, "hu")
		delete(additionalProperties, "id")
		delete(additionalProperties, "it")
		delete(additionalProperties, "ja")
		delete(additionalProperties, "ko")
		delete(additionalProperties, "lv")
		delete(additionalProperties, "lt")
		delete(additionalProperties, "ms")
		delete(additionalProperties, "nb")
		delete(additionalProperties, "pl")
		delete(additionalProperties, "fa")
		delete(additionalProperties, "pt")
		delete(additionalProperties, "pa")
		delete(additionalProperties, "ro")
		delete(additionalProperties, "ru")
		delete(additionalProperties, "sr")
		delete(additionalProperties, "sk")
		delete(additionalProperties, "es")
		delete(additionalProperties, "sv")
		delete(additionalProperties, "th")
		delete(additionalProperties, "tr")
		delete(additionalProperties, "uk")
		delete(additionalProperties, "vi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStringMap struct {
	value *StringMap
	isSet bool
}

func (v NullableStringMap) Get() *StringMap {
	return v.value
}

func (v *NullableStringMap) Set(val *StringMap) {
	v.value = val
	v.isSet = true
}

func (v NullableStringMap) IsSet() bool {
	return v.isSet
}

func (v *NullableStringMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringMap(val *StringMap) *NullableStringMap {
	return &NullableStringMap{value: val, isSet: true}
}

func (v NullableStringMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



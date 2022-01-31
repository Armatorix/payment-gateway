/*
Payment Gateway API

Payment gateway API for merchants usage.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// AuthorizeRequest struct for AuthorizeRequest
type AuthorizeRequest struct {
	CardNumber *string                     `json:"card_number,omitempty"`
	CardHolder *string                     `json:"card_holder,omitempty"`
	CardExpiry *AuthorizeRequestCardExpiry `json:"card_expiry,omitempty"`
	CardCvv    *int32                      `json:"card_cvv,omitempty"`
	// Amount of money in a form of least significant monetary unit.
	Amount   *int64  `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewAuthorizeRequest instantiates a new AuthorizeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeRequest() *AuthorizeRequest {
	this := AuthorizeRequest{}
	return &this
}

// NewAuthorizeRequestWithDefaults instantiates a new AuthorizeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeRequestWithDefaults() *AuthorizeRequest {
	this := AuthorizeRequest{}
	return &this
}

// GetCardNumber returns the CardNumber field value if set, zero value otherwise.
func (o *AuthorizeRequest) GetCardNumber() string {
	if o == nil || o.CardNumber == nil {
		var ret string
		return ret
	}
	return *o.CardNumber
}

// GetCardNumberOk returns a tuple with the CardNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeRequest) GetCardNumberOk() (*string, bool) {
	if o == nil || o.CardNumber == nil {
		return nil, false
	}
	return o.CardNumber, true
}

// HasCardNumber returns a boolean if a field has been set.
func (o *AuthorizeRequest) HasCardNumber() bool {
	if o != nil && o.CardNumber != nil {
		return true
	}

	return false
}

// SetCardNumber gets a reference to the given string and assigns it to the CardNumber field.
func (o *AuthorizeRequest) SetCardNumber(v string) {
	o.CardNumber = &v
}

// GetCardHolder returns the CardHolder field value if set, zero value otherwise.
func (o *AuthorizeRequest) GetCardHolder() string {
	if o == nil || o.CardHolder == nil {
		var ret string
		return ret
	}
	return *o.CardHolder
}

// GetCardHolderOk returns a tuple with the CardHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeRequest) GetCardHolderOk() (*string, bool) {
	if o == nil || o.CardHolder == nil {
		return nil, false
	}
	return o.CardHolder, true
}

// HasCardHolder returns a boolean if a field has been set.
func (o *AuthorizeRequest) HasCardHolder() bool {
	if o != nil && o.CardHolder != nil {
		return true
	}

	return false
}

// SetCardHolder gets a reference to the given string and assigns it to the CardHolder field.
func (o *AuthorizeRequest) SetCardHolder(v string) {
	o.CardHolder = &v
}

// GetCardExpiry returns the CardExpiry field value if set, zero value otherwise.
func (o *AuthorizeRequest) GetCardExpiry() AuthorizeRequestCardExpiry {
	if o == nil || o.CardExpiry == nil {
		var ret AuthorizeRequestCardExpiry
		return ret
	}
	return *o.CardExpiry
}

// GetCardExpiryOk returns a tuple with the CardExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeRequest) GetCardExpiryOk() (*AuthorizeRequestCardExpiry, bool) {
	if o == nil || o.CardExpiry == nil {
		return nil, false
	}
	return o.CardExpiry, true
}

// HasCardExpiry returns a boolean if a field has been set.
func (o *AuthorizeRequest) HasCardExpiry() bool {
	if o != nil && o.CardExpiry != nil {
		return true
	}

	return false
}

// SetCardExpiry gets a reference to the given AuthorizeRequestCardExpiry and assigns it to the CardExpiry field.
func (o *AuthorizeRequest) SetCardExpiry(v AuthorizeRequestCardExpiry) {
	o.CardExpiry = &v
}

// GetCardCvv returns the CardCvv field value if set, zero value otherwise.
func (o *AuthorizeRequest) GetCardCvv() int32 {
	if o == nil || o.CardCvv == nil {
		var ret int32
		return ret
	}
	return *o.CardCvv
}

// GetCardCvvOk returns a tuple with the CardCvv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeRequest) GetCardCvvOk() (*int32, bool) {
	if o == nil || o.CardCvv == nil {
		return nil, false
	}
	return o.CardCvv, true
}

// HasCardCvv returns a boolean if a field has been set.
func (o *AuthorizeRequest) HasCardCvv() bool {
	if o != nil && o.CardCvv != nil {
		return true
	}

	return false
}

// SetCardCvv gets a reference to the given int32 and assigns it to the CardCvv field.
func (o *AuthorizeRequest) SetCardCvv(v int32) {
	o.CardCvv = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AuthorizeRequest) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeRequest) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AuthorizeRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *AuthorizeRequest) SetAmount(v int64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *AuthorizeRequest) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *AuthorizeRequest) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *AuthorizeRequest) SetCurrency(v string) {
	o.Currency = &v
}

func (o AuthorizeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CardNumber != nil {
		toSerialize["card_number"] = o.CardNumber
	}
	if o.CardHolder != nil {
		toSerialize["card_holder"] = o.CardHolder
	}
	if o.CardExpiry != nil {
		toSerialize["card_expiry"] = o.CardExpiry
	}
	if o.CardCvv != nil {
		toSerialize["card_cvv"] = o.CardCvv
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizeRequest struct {
	value *AuthorizeRequest
	isSet bool
}

func (v NullableAuthorizeRequest) Get() *AuthorizeRequest {
	return v.value
}

func (v *NullableAuthorizeRequest) Set(val *AuthorizeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeRequest(val *AuthorizeRequest) *NullableAuthorizeRequest {
	return &NullableAuthorizeRequest{value: val, isSet: true}
}

func (v NullableAuthorizeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

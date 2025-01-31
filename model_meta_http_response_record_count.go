/*
Zinc Search engine API

Zinc Search engine API documents https://docs.zincsearch.com

API version: 0.2.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MetaHTTPResponseRecordCount struct for MetaHTTPResponseRecordCount
type MetaHTTPResponseRecordCount struct {
	Message *string `json:"message,omitempty"`
	RecordCount *int32 `json:"record_count,omitempty"`
}

// NewMetaHTTPResponseRecordCount instantiates a new MetaHTTPResponseRecordCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaHTTPResponseRecordCount() *MetaHTTPResponseRecordCount {
	this := MetaHTTPResponseRecordCount{}
	return &this
}

// NewMetaHTTPResponseRecordCountWithDefaults instantiates a new MetaHTTPResponseRecordCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaHTTPResponseRecordCountWithDefaults() *MetaHTTPResponseRecordCount {
	this := MetaHTTPResponseRecordCount{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MetaHTTPResponseRecordCount) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHTTPResponseRecordCount) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MetaHTTPResponseRecordCount) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MetaHTTPResponseRecordCount) SetMessage(v string) {
	o.Message = &v
}

// GetRecordCount returns the RecordCount field value if set, zero value otherwise.
func (o *MetaHTTPResponseRecordCount) GetRecordCount() int32 {
	if o == nil || o.RecordCount == nil {
		var ret int32
		return ret
	}
	return *o.RecordCount
}

// GetRecordCountOk returns a tuple with the RecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaHTTPResponseRecordCount) GetRecordCountOk() (*int32, bool) {
	if o == nil || o.RecordCount == nil {
		return nil, false
	}
	return o.RecordCount, true
}

// HasRecordCount returns a boolean if a field has been set.
func (o *MetaHTTPResponseRecordCount) HasRecordCount() bool {
	if o != nil && o.RecordCount != nil {
		return true
	}

	return false
}

// SetRecordCount gets a reference to the given int32 and assigns it to the RecordCount field.
func (o *MetaHTTPResponseRecordCount) SetRecordCount(v int32) {
	o.RecordCount = &v
}

func (o MetaHTTPResponseRecordCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.RecordCount != nil {
		toSerialize["record_count"] = o.RecordCount
	}
	return json.Marshal(toSerialize)
}

type NullableMetaHTTPResponseRecordCount struct {
	value *MetaHTTPResponseRecordCount
	isSet bool
}

func (v NullableMetaHTTPResponseRecordCount) Get() *MetaHTTPResponseRecordCount {
	return v.value
}

func (v *NullableMetaHTTPResponseRecordCount) Set(val *MetaHTTPResponseRecordCount) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaHTTPResponseRecordCount) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaHTTPResponseRecordCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaHTTPResponseRecordCount(val *MetaHTTPResponseRecordCount) *NullableMetaHTTPResponseRecordCount {
	return &NullableMetaHTTPResponseRecordCount{value: val, isSet: true}
}

func (v NullableMetaHTTPResponseRecordCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaHTTPResponseRecordCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



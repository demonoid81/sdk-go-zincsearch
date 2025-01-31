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

// MetaJSONIngest struct for MetaJSONIngest
type MetaJSONIngest struct {
	Index *string `json:"index,omitempty"`
	Records []map[string]interface{} `json:"records,omitempty"`
}

// NewMetaJSONIngest instantiates a new MetaJSONIngest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaJSONIngest() *MetaJSONIngest {
	this := MetaJSONIngest{}
	return &this
}

// NewMetaJSONIngestWithDefaults instantiates a new MetaJSONIngest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaJSONIngestWithDefaults() *MetaJSONIngest {
	this := MetaJSONIngest{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *MetaJSONIngest) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaJSONIngest) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *MetaJSONIngest) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *MetaJSONIngest) SetIndex(v string) {
	o.Index = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *MetaJSONIngest) GetRecords() []map[string]interface{} {
	if o == nil || o.Records == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaJSONIngest) GetRecordsOk() ([]map[string]interface{}, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *MetaJSONIngest) HasRecords() bool {
	if o != nil && o.Records != nil {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []map[string]interface{} and assigns it to the Records field.
func (o *MetaJSONIngest) SetRecords(v []map[string]interface{}) {
	o.Records = v
}

func (o MetaJSONIngest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableMetaJSONIngest struct {
	value *MetaJSONIngest
	isSet bool
}

func (v NullableMetaJSONIngest) Get() *MetaJSONIngest {
	return v.value
}

func (v *NullableMetaJSONIngest) Set(val *MetaJSONIngest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaJSONIngest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaJSONIngest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaJSONIngest(val *MetaJSONIngest) *NullableMetaJSONIngest {
	return &NullableMetaJSONIngest{value: val, isSet: true}
}

func (v NullableMetaJSONIngest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaJSONIngest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



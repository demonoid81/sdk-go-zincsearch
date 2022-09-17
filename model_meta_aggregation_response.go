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

type BucketsValue struct {
	DocCount    int    `json:"doc_count,omitempty"`
	Key         int    `json:"key,omitempty"`
	KeyAsString string `json:"key_as_string,omitempty"`
}

// MetaAggregationResponse struct for MetaAggregationResponse
type MetaAggregationResponse struct {
	// slice or map
	Buckets []*BucketsValue `json:"buckets,omitempty"`
	// support for auto_date_histogram_aggregation
	Interval *string `json:"interval,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewMetaAggregationResponse instantiates a new MetaAggregationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaAggregationResponse() *MetaAggregationResponse {
	this := MetaAggregationResponse{}
	return &this
}

// NewMetaAggregationResponseWithDefaults instantiates a new MetaAggregationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaAggregationResponseWithDefaults() *MetaAggregationResponse {
	this := MetaAggregationResponse{}
	return &this
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *MetaAggregationResponse) GetBuckets() []*BucketsValue {
	if o == nil || o.Buckets == nil {
		var ret []*BucketsValue
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationResponse) GetBucketsOk() ([]*BucketsValue, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *MetaAggregationResponse) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given map[string]interface{} and assigns it to the Buckets field.
func (o *MetaAggregationResponse) SetBuckets(v []*BucketsValue) {
	o.Buckets = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *MetaAggregationResponse) GetInterval() string {
	if o == nil || o.Interval == nil {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationResponse) GetIntervalOk() (*string, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *MetaAggregationResponse) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *MetaAggregationResponse) SetInterval(v string) {
	o.Interval = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MetaAggregationResponse) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregationResponse) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MetaAggregationResponse) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *MetaAggregationResponse) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o MetaAggregationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableMetaAggregationResponse struct {
	value *MetaAggregationResponse
	isSet bool
}

func (v NullableMetaAggregationResponse) Get() *MetaAggregationResponse {
	return v.value
}

func (v *NullableMetaAggregationResponse) Set(val *MetaAggregationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaAggregationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaAggregationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaAggregationResponse(val *MetaAggregationResponse) *NullableMetaAggregationResponse {
	return &NullableMetaAggregationResponse{value: val, isSet: true}
}

func (v NullableMetaAggregationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaAggregationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



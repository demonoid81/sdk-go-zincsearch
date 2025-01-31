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

// MetaAggregations struct for MetaAggregations
type MetaAggregations struct {
	// nested aggregations
	Aggs *map[string]MetaAggregations `json:"aggs,omitempty"`
	AutoDateHistogram *MetaAggregationAutoDateHistogram `json:"auto_date_histogram,omitempty"`
	Avg *MetaAggregationMetric `json:"avg,omitempty"`
	Cardinality *MetaAggregationMetric `json:"cardinality,omitempty"`
	Count *MetaAggregationMetric `json:"count,omitempty"`
	DateHistogram *MetaAggregationDateHistogram `json:"date_histogram,omitempty"`
	DateRange *MetaAggregationDateRange `json:"date_range,omitempty"`
	Histogram *MetaAggregationHistogram `json:"histogram,omitempty"`
	IpRange *MetaAggregationIPRange `json:"ip_range,omitempty"`
	Max *MetaAggregationMetric `json:"max,omitempty"`
	Min *MetaAggregationMetric `json:"min,omitempty"`
	Range *MetaAggregationRange `json:"range,omitempty"`
	Sum *MetaAggregationMetric `json:"sum,omitempty"`
	Terms *MetaAggregationsTerms `json:"terms,omitempty"`
	WeightedAvg *MetaAggregationMetric `json:"weighted_avg,omitempty"`
}

// NewMetaAggregations instantiates a new MetaAggregations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaAggregations() *MetaAggregations {
	this := MetaAggregations{}
	return &this
}

// NewMetaAggregationsWithDefaults instantiates a new MetaAggregations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaAggregationsWithDefaults() *MetaAggregations {
	this := MetaAggregations{}
	return &this
}

// GetAggs returns the Aggs field value if set, zero value otherwise.
func (o *MetaAggregations) GetAggs() map[string]MetaAggregations {
	if o == nil || o.Aggs == nil {
		var ret map[string]MetaAggregations
		return ret
	}
	return *o.Aggs
}

// GetAggsOk returns a tuple with the Aggs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetAggsOk() (*map[string]MetaAggregations, bool) {
	if o == nil || o.Aggs == nil {
		return nil, false
	}
	return o.Aggs, true
}

// HasAggs returns a boolean if a field has been set.
func (o *MetaAggregations) HasAggs() bool {
	if o != nil && o.Aggs != nil {
		return true
	}

	return false
}

// SetAggs gets a reference to the given map[string]MetaAggregations and assigns it to the Aggs field.
func (o *MetaAggregations) SetAggs(v map[string]MetaAggregations) {
	o.Aggs = &v
}

// GetAutoDateHistogram returns the AutoDateHistogram field value if set, zero value otherwise.
func (o *MetaAggregations) GetAutoDateHistogram() MetaAggregationAutoDateHistogram {
	if o == nil || o.AutoDateHistogram == nil {
		var ret MetaAggregationAutoDateHistogram
		return ret
	}
	return *o.AutoDateHistogram
}

// GetAutoDateHistogramOk returns a tuple with the AutoDateHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetAutoDateHistogramOk() (*MetaAggregationAutoDateHistogram, bool) {
	if o == nil || o.AutoDateHistogram == nil {
		return nil, false
	}
	return o.AutoDateHistogram, true
}

// HasAutoDateHistogram returns a boolean if a field has been set.
func (o *MetaAggregations) HasAutoDateHistogram() bool {
	if o != nil && o.AutoDateHistogram != nil {
		return true
	}

	return false
}

// SetAutoDateHistogram gets a reference to the given MetaAggregationAutoDateHistogram and assigns it to the AutoDateHistogram field.
func (o *MetaAggregations) SetAutoDateHistogram(v MetaAggregationAutoDateHistogram) {
	o.AutoDateHistogram = &v
}

// GetAvg returns the Avg field value if set, zero value otherwise.
func (o *MetaAggregations) GetAvg() MetaAggregationMetric {
	if o == nil || o.Avg == nil {
		var ret MetaAggregationMetric
		return ret
	}
	return *o.Avg
}

// GetAvgOk returns a tuple with the Avg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetAvgOk() (*MetaAggregationMetric, bool) {
	if o == nil || o.Avg == nil {
		return nil, false
	}
	return o.Avg, true
}

// HasAvg returns a boolean if a field has been set.
func (o *MetaAggregations) HasAvg() bool {
	if o != nil && o.Avg != nil {
		return true
	}

	return false
}

// SetAvg gets a reference to the given MetaAggregationMetric and assigns it to the Avg field.
func (o *MetaAggregations) SetAvg(v MetaAggregationMetric) {
	o.Avg = &v
}

// GetCardinality returns the Cardinality field value if set, zero value otherwise.
func (o *MetaAggregations) GetCardinality() MetaAggregationMetric {
	if o == nil || o.Cardinality == nil {
		var ret MetaAggregationMetric
		return ret
	}
	return *o.Cardinality
}

// GetCardinalityOk returns a tuple with the Cardinality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetCardinalityOk() (*MetaAggregationMetric, bool) {
	if o == nil || o.Cardinality == nil {
		return nil, false
	}
	return o.Cardinality, true
}

// HasCardinality returns a boolean if a field has been set.
func (o *MetaAggregations) HasCardinality() bool {
	if o != nil && o.Cardinality != nil {
		return true
	}

	return false
}

// SetCardinality gets a reference to the given MetaAggregationMetric and assigns it to the Cardinality field.
func (o *MetaAggregations) SetCardinality(v MetaAggregationMetric) {
	o.Cardinality = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *MetaAggregations) GetCount() MetaAggregationMetric {
	if o == nil || o.Count == nil {
		var ret MetaAggregationMetric
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetCountOk() (*MetaAggregationMetric, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *MetaAggregations) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given MetaAggregationMetric and assigns it to the Count field.
func (o *MetaAggregations) SetCount(v MetaAggregationMetric) {
	o.Count = &v
}

// GetDateHistogram returns the DateHistogram field value if set, zero value otherwise.
func (o *MetaAggregations) GetDateHistogram() MetaAggregationDateHistogram {
	if o == nil || o.DateHistogram == nil {
		var ret MetaAggregationDateHistogram
		return ret
	}
	return *o.DateHistogram
}

// GetDateHistogramOk returns a tuple with the DateHistogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetDateHistogramOk() (*MetaAggregationDateHistogram, bool) {
	if o == nil || o.DateHistogram == nil {
		return nil, false
	}
	return o.DateHistogram, true
}

// HasDateHistogram returns a boolean if a field has been set.
func (o *MetaAggregations) HasDateHistogram() bool {
	if o != nil && o.DateHistogram != nil {
		return true
	}

	return false
}

// SetDateHistogram gets a reference to the given MetaAggregationDateHistogram and assigns it to the DateHistogram field.
func (o *MetaAggregations) SetDateHistogram(v MetaAggregationDateHistogram) {
	o.DateHistogram = &v
}

// GetDateRange returns the DateRange field value if set, zero value otherwise.
func (o *MetaAggregations) GetDateRange() MetaAggregationDateRange {
	if o == nil || o.DateRange == nil {
		var ret MetaAggregationDateRange
		return ret
	}
	return *o.DateRange
}

// GetDateRangeOk returns a tuple with the DateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetDateRangeOk() (*MetaAggregationDateRange, bool) {
	if o == nil || o.DateRange == nil {
		return nil, false
	}
	return o.DateRange, true
}

// HasDateRange returns a boolean if a field has been set.
func (o *MetaAggregations) HasDateRange() bool {
	if o != nil && o.DateRange != nil {
		return true
	}

	return false
}

// SetDateRange gets a reference to the given MetaAggregationDateRange and assigns it to the DateRange field.
func (o *MetaAggregations) SetDateRange(v MetaAggregationDateRange) {
	o.DateRange = &v
}

// GetHistogram returns the Histogram field value if set, zero value otherwise.
func (o *MetaAggregations) GetHistogram() MetaAggregationHistogram {
	if o == nil || o.Histogram == nil {
		var ret MetaAggregationHistogram
		return ret
	}
	return *o.Histogram
}

// GetHistogramOk returns a tuple with the Histogram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetHistogramOk() (*MetaAggregationHistogram, bool) {
	if o == nil || o.Histogram == nil {
		return nil, false
	}
	return o.Histogram, true
}

// HasHistogram returns a boolean if a field has been set.
func (o *MetaAggregations) HasHistogram() bool {
	if o != nil && o.Histogram != nil {
		return true
	}

	return false
}

// SetHistogram gets a reference to the given MetaAggregationHistogram and assigns it to the Histogram field.
func (o *MetaAggregations) SetHistogram(v MetaAggregationHistogram) {
	o.Histogram = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *MetaAggregations) GetIpRange() MetaAggregationIPRange {
	if o == nil || o.IpRange == nil {
		var ret MetaAggregationIPRange
		return ret
	}
	return *o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetIpRangeOk() (*MetaAggregationIPRange, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *MetaAggregations) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given MetaAggregationIPRange and assigns it to the IpRange field.
func (o *MetaAggregations) SetIpRange(v MetaAggregationIPRange) {
	o.IpRange = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *MetaAggregations) GetMax() MetaAggregationMetric {
	if o == nil || o.Max == nil {
		var ret MetaAggregationMetric
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetMaxOk() (*MetaAggregationMetric, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *MetaAggregations) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given MetaAggregationMetric and assigns it to the Max field.
func (o *MetaAggregations) SetMax(v MetaAggregationMetric) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *MetaAggregations) GetMin() MetaAggregationMetric {
	if o == nil || o.Min == nil {
		var ret MetaAggregationMetric
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetMinOk() (*MetaAggregationMetric, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *MetaAggregations) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given MetaAggregationMetric and assigns it to the Min field.
func (o *MetaAggregations) SetMin(v MetaAggregationMetric) {
	o.Min = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *MetaAggregations) GetRange() MetaAggregationRange {
	if o == nil || o.Range == nil {
		var ret MetaAggregationRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetRangeOk() (*MetaAggregationRange, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *MetaAggregations) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given MetaAggregationRange and assigns it to the Range field.
func (o *MetaAggregations) SetRange(v MetaAggregationRange) {
	o.Range = &v
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *MetaAggregations) GetSum() MetaAggregationMetric {
	if o == nil || o.Sum == nil {
		var ret MetaAggregationMetric
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetSumOk() (*MetaAggregationMetric, bool) {
	if o == nil || o.Sum == nil {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *MetaAggregations) HasSum() bool {
	if o != nil && o.Sum != nil {
		return true
	}

	return false
}

// SetSum gets a reference to the given MetaAggregationMetric and assigns it to the Sum field.
func (o *MetaAggregations) SetSum(v MetaAggregationMetric) {
	o.Sum = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *MetaAggregations) GetTerms() MetaAggregationsTerms {
	if o == nil || o.Terms == nil {
		var ret MetaAggregationsTerms
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetTermsOk() (*MetaAggregationsTerms, bool) {
	if o == nil || o.Terms == nil {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *MetaAggregations) HasTerms() bool {
	if o != nil && o.Terms != nil {
		return true
	}

	return false
}

// SetTerms gets a reference to the given MetaAggregationsTerms and assigns it to the Terms field.
func (o *MetaAggregations) SetTerms(v MetaAggregationsTerms) {
	o.Terms = &v
}

// GetWeightedAvg returns the WeightedAvg field value if set, zero value otherwise.
func (o *MetaAggregations) GetWeightedAvg() MetaAggregationMetric {
	if o == nil || o.WeightedAvg == nil {
		var ret MetaAggregationMetric
		return ret
	}
	return *o.WeightedAvg
}

// GetWeightedAvgOk returns a tuple with the WeightedAvg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaAggregations) GetWeightedAvgOk() (*MetaAggregationMetric, bool) {
	if o == nil || o.WeightedAvg == nil {
		return nil, false
	}
	return o.WeightedAvg, true
}

// HasWeightedAvg returns a boolean if a field has been set.
func (o *MetaAggregations) HasWeightedAvg() bool {
	if o != nil && o.WeightedAvg != nil {
		return true
	}

	return false
}

// SetWeightedAvg gets a reference to the given MetaAggregationMetric and assigns it to the WeightedAvg field.
func (o *MetaAggregations) SetWeightedAvg(v MetaAggregationMetric) {
	o.WeightedAvg = &v
}

func (o MetaAggregations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aggs != nil {
		toSerialize["aggs"] = o.Aggs
	}
	if o.AutoDateHistogram != nil {
		toSerialize["auto_date_histogram"] = o.AutoDateHistogram
	}
	if o.Avg != nil {
		toSerialize["avg"] = o.Avg
	}
	if o.Cardinality != nil {
		toSerialize["cardinality"] = o.Cardinality
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.DateHistogram != nil {
		toSerialize["date_histogram"] = o.DateHistogram
	}
	if o.DateRange != nil {
		toSerialize["date_range"] = o.DateRange
	}
	if o.Histogram != nil {
		toSerialize["histogram"] = o.Histogram
	}
	if o.IpRange != nil {
		toSerialize["ip_range"] = o.IpRange
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.Sum != nil {
		toSerialize["sum"] = o.Sum
	}
	if o.Terms != nil {
		toSerialize["terms"] = o.Terms
	}
	if o.WeightedAvg != nil {
		toSerialize["weighted_avg"] = o.WeightedAvg
	}
	return json.Marshal(toSerialize)
}

type NullableMetaAggregations struct {
	value *MetaAggregations
	isSet bool
}

func (v NullableMetaAggregations) Get() *MetaAggregations {
	return v.value
}

func (v *NullableMetaAggregations) Set(val *MetaAggregations) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaAggregations) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaAggregations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaAggregations(val *MetaAggregations) *NullableMetaAggregations {
	return &NullableMetaAggregations{value: val, isSet: true}
}

func (v NullableMetaAggregations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaAggregations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



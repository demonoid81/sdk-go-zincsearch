# Go API client for client

Zinc Search engine API documents https://docs.zincsearch.com

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.2.4
- Package version: 0.2.4
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.zincsearch.com](https://www.zincsearch.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import client "github.com/zinclabs/sdk-go-zincsearch"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:4080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*Default* | [**Healthz**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Default.md#healthz) | **Get** /healthz | Get healthz
*Default* | [**Version**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Default.md#version) | **Get** /version | Get version
*Document* | [**Bulk**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Document.md#bulk) | **Post** /api/_bulk | Bulk documents
*Document* | [**Delete**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Document.md#delete) | **Delete** /api/{index}/_doc/{id} | Delete document
*Document* | [**ESBulk**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Document.md#esbulk) | **Post** /es/_bulk | ES bulk documents
*Document* | [**Index**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Document.md#index) | **Post** /api/{index}/_doc | Create or update document
*Document* | [**IndexWithID**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Document.md#indexwithid) | **Put** /api/{index}/_doc/{id} | Create or update document with id
*Document* | [**Update**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Document.md#update) | **Post** /api/{index}/_update/{id} | Update document with id
*Index* | [**Analyze**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#analyze) | **Post** /api/_analyze | Analyze
*Index* | [**AnalyzeIndex**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#analyzeindex) | **Post** /api/{index}/_analyze | Analyze
*Index* | [**Create**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#create) | **Post** /api/index | Create index
*Index* | [**CreateTemplate**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#createtemplate) | **Post** /es/_index_template | Create update index template
*Index* | [**Delete**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#delete) | **Delete** /api/index/{index} | Delete index
*Index* | [**DeleteTemplate**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#deletetemplate) | **Delete** /es/_index_template/{name} | Delete template
*Index* | [**GetMapping**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#getmapping) | **Get** /api/{index}/_mapping | Get index mappings
*Index* | [**GetSettings**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#getsettings) | **Get** /api/{index}/_settings | Get index settings
*Index* | [**GetTemplate**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#gettemplate) | **Get** /es/_index_template/{name} | Get index template
*Index* | [**List**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#list) | **Get** /api/index | List indexes
*Index* | [**ListTemplates**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#listtemplates) | **Get** /es/_index_template | List index teplates
*Index* | [**Refresh**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#refresh) | **Post** /api/index/{index}/refresh | Resfresh index
*Index* | [**SetMapping**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#setmapping) | **Put** /api/{index}/_mapping | Set index mappings
*Index* | [**SetSettings**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#setsettings) | **Put** /api/{index}/_settings | Set index Settings
*Index* | [**UpdateTemplate**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Index.md#updatetemplate) | **Put** /es/_index_template/{name} | Create update index template
*Search* | [**MSearch**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Search.md#msearch) | **Post** /es/_msearch | Search V2 MultipleSearch for compatible ES
*Search* | [**Search**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Search.md#search) | **Post** /es/{index}/_search | Search V2 DSL for compatible ES
*Search* | [**SearchV1**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/Search.md#searchv1) | **Post** /api/{index}/_search | Search V1
*User* | [**Create**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/User.md#create) | **Post** /api/user | Create user
*User* | [**Delete**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/User.md#delete) | **Delete** /api/user/{id} | Delete user
*User* | [**List**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/User.md#list) | **Get** /api/user | List user
*User* | [**Login**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/User.md#login) | **Post** /api/login | Login
*User* | [**Update**](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/User.md#update) | **Put** /api/user | Update user


## Documentation For Models

 - [AggregationHistogramBound](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/AggregationHistogramBound.md)
 - [AuthLoginRequest](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/AuthLoginRequest.md)
 - [AuthLoginResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/AuthLoginResponse.md)
 - [AuthLoginUser](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/AuthLoginUser.md)
 - [CoreIndex](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/CoreIndex.md)
 - [IndexAnalyzeResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/IndexAnalyzeResponse.md)
 - [IndexAnalyzeResponseToken](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/IndexAnalyzeResponseToken.md)
 - [MetaAggregationAutoDateHistogram](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationAutoDateHistogram.md)
 - [MetaAggregationDateHistogram](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationDateHistogram.md)
 - [MetaAggregationDateRange](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationDateRange.md)
 - [MetaAggregationHistogram](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationHistogram.md)
 - [MetaAggregationIPRange](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationIPRange.md)
 - [MetaAggregationMetric](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationMetric.md)
 - [MetaAggregationRange](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationRange.md)
 - [MetaAggregationResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationResponse.md)
 - [MetaAggregations](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregations.md)
 - [MetaAggregationsTerms](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAggregationsTerms.md)
 - [MetaAnalyzer](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaAnalyzer.md)
 - [MetaBoolQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaBoolQuery.md)
 - [MetaDateRange](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaDateRange.md)
 - [MetaExistsQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaExistsQuery.md)
 - [MetaFuzzyQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaFuzzyQuery.md)
 - [MetaHTTPResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHTTPResponse.md)
 - [MetaHTTPResponseDocument](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHTTPResponseDocument.md)
 - [MetaHTTPResponseError](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHTTPResponseError.md)
 - [MetaHTTPResponseID](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHTTPResponseID.md)
 - [MetaHTTPResponseIndex](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHTTPResponseIndex.md)
 - [MetaHTTPResponseRecordCount](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHTTPResponseRecordCount.md)
 - [MetaHTTPResponseTemplate](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHTTPResponseTemplate.md)
 - [MetaHealthzResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHealthzResponse.md)
 - [MetaHighlight](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHighlight.md)
 - [MetaHit](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHit.md)
 - [MetaHits](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaHits.md)
 - [MetaIPRange](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaIPRange.md)
 - [MetaIdsQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaIdsQuery.md)
 - [MetaIndexAnalysis](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaIndexAnalysis.md)
 - [MetaIndexSettings](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaIndexSettings.md)
 - [MetaIndexShard](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaIndexShard.md)
 - [MetaIndexSimple](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaIndexSimple.md)
 - [MetaIndexTemplate](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaIndexTemplate.md)
 - [MetaMappings](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaMappings.md)
 - [MetaMatchBoolPrefixQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaMatchBoolPrefixQuery.md)
 - [MetaMatchPhrasePrefixQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaMatchPhrasePrefixQuery.md)
 - [MetaMatchPhraseQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaMatchPhraseQuery.md)
 - [MetaMatchQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaMatchQuery.md)
 - [MetaMultiMatchQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaMultiMatchQuery.md)
 - [MetaPrefixQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaPrefixQuery.md)
 - [MetaProperty](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaProperty.md)
 - [MetaQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaQuery.md)
 - [MetaQueryStringQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaQueryStringQuery.md)
 - [MetaRange](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaRange.md)
 - [MetaRangeQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaRangeQuery.md)
 - [MetaRegexpQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaRegexpQuery.md)
 - [MetaSearchResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaSearchResponse.md)
 - [MetaShards](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaShards.md)
 - [MetaSimpleQueryStringQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaSimpleQueryStringQuery.md)
 - [MetaTemplate](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaTemplate.md)
 - [MetaTemplateTemplate](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaTemplateTemplate.md)
 - [MetaTermQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaTermQuery.md)
 - [MetaTotal](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaTotal.md)
 - [MetaUser](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaUser.md)
 - [MetaVersionResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaVersionResponse.md)
 - [MetaWildcardQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaWildcardQuery.md)
 - [MetaZincQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/MetaZincQuery.md)
 - [V1AggregationBucket](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1AggregationBucket.md)
 - [V1AggregationDateRange](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1AggregationDateRange.md)
 - [V1AggregationNumberRange](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1AggregationNumberRange.md)
 - [V1AggregationParams](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1AggregationParams.md)
 - [V1AggregationResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1AggregationResponse.md)
 - [V1Hit](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1Hit.md)
 - [V1Hits](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1Hits.md)
 - [V1QueryHighlight](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1QueryHighlight.md)
 - [V1QueryParams](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1QueryParams.md)
 - [V1SearchResponse](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1SearchResponse.md)
 - [V1Total](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1Total.md)
 - [V1ZincQuery](https://github.com/zinclabs/sdk-go-zincsearch/tree/main/docs/V1ZincQuery.md)


## Documentation For Authorization



### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), client.ContextBasicAuth, client.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




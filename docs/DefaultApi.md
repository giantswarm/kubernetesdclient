# \DefaultApi

All URIs are relative to *http://kubernetesd.giantswarm.g8s.fra-1.giantswarm.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**100ClustersPost**](DefaultApi.md#100ClustersPost) | **Post** /1.0.0/clusters | 
[**RootGet**](DefaultApi.md#RootGet) | **Get** / | 


# **100ClustersPost**
> CreatorSuccessResponse 100ClustersPost($cluster)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | [**CreatorRequest**](CreatorRequest.md)| Cluster creator specific request information. | 

### Return type

[**CreatorSuccessResponse**](CreatorSuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RootGet**
> ServiceInformation RootGet()



Returns service specific information.


### Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceInformation**](ServiceInformation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


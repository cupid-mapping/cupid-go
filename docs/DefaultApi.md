# {{classname}}

All URIs are relative to *http://34.219.111.155:8080/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryDetails**](DefaultApi.md#InventoryDetails) | **Get** /inventories/{inventory_id} | Inventory details
[**ListInventories**](DefaultApi.md#ListInventories) | **Get** /inventories | List inventories
[**MapHotels**](DefaultApi.md#MapHotels) | **Post** /map-hotels | Map hotel list
[**UploadInventory**](DefaultApi.md#UploadInventory) | **Post** /inventories/upload | Upload inventory

# **InventoryDetails**
> InlineResponse2001 InventoryDetails(ctx, inventoryId)
Inventory details

This endpoint can be used to check the requested inventory's details.  You can use it to get the inventory status ```mapping_process_status_id``` as outlined below. The status has to be 2 (Done) to start mapping.  **Possible values for the inventory's status ```mapping_process_status_id```:**  | value     | Name           | Description                                        | Action                                                     | | --------- | -------------- | -------------------------------------------------- | ---------------------------------------------------------- | | -1        | **Invalid**    | The inventory doesn't contain any valid hotels       | Correct your catalog and try again.                        | | 0         | **Pending**    | The inventory is being uploaded                      | Wait, no action required.                                  | | 1         | **Processing** | The inventory is being processed                     | Wait, no action required.                                  | | 2         | **Done**       | The process is complete                            | **You can start mapping.**                                 | | 3         | **Failed**     | There was an error while processing                | Retry and if the error persists, feel free to contact us.  |  When you have an inventory with ```active=true``` and ```mapping_process_status_id=2``` you can use the following endpoint to perfom mapping:  - [Map a list of properties](https://nuitee2.stoplight.io/docs/cupid/b3A6NDUxMTM0NTA-map-hotel-list) <!-- - [Map a CSV file](https://nuitee2.stoplight.io/docs/cupid/b3A6NDQ3ODY2OTY-upload-user-catalog)-->

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inventoryId** | **int32**| Inventory ID | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInventories**
> InlineResponse200 ListInventories(ctx, )
List inventories

This endpoint can be used to list the inventories uploaded to your workspace.<br>You can check their status and other basic info.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MapHotels**
> MapHotelsResponse MapHotels(ctx, optional)
Map hotel list

<!-- theme: warning --> > #### Requirement > > Active inventory with ```mapping_process_status_id=2```, you can check this using the [inventory details endpoint](https://nuitee2.stoplight.io/docs/cupid/b3A6NDg1OTM1MTA-inventory-details).  This endpoint allows you to send a list of properties and maps it against your active inventory. You will get the resulting mappings in the response.  ### Limit You can send up to 1000 properties per request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiMapHotelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiMapHotelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of []HotelItem**](HotelItem.md)| List of hotels to map | 
 **extendedInfo** | **optional.**| If true, returns properties info along with the mapped IDs.&lt;br&gt; Otherwise, returns only the user properties IDs along with the mapped partner IDs. | [default to false]

### Return type

[**MapHotelsResponse**](MapHotelsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadInventory**
> UploadInventoryResponse UploadInventory(ctx, file, name, headerId, headerName, headerAddress, headerCity, headerCountryCode, headerLatitude, headerLongitude)
Upload inventory

This end point can be used to upload an inventory with a provided CSV file.<br> <br> You will have to provide the inventory's name and for each column in your file, you'll need match its index with the standard field names provided below.<br>  Let's assume you have a CSV file that looks like the following table:  | HotelID     | Hotel_Name   | Hotel_Address | Hotel_Lat  | Hotel_Lon  | Country | City     | ... | | ----------- | ------------ | ------------- | ---------- | ---------- | ------- | -------- | --- | | 1408        | Smile Hotel  | 123 main st   | 5.419514   | 3.392694   | AU      | Adelaid  | ... | | 123         | Cupid resort | 456 north av  | 55.69254   | 37.3167887 | MA      | Zagoura  | ... |  In this case the form data would be: - **file**: The csv file to be uploaded - **name**: Cupid inventory - **header_id**: 0 - **header_name**: 1 - **header_address**: 2 - **header_city**: 6 - **header_country_code**: 5 - **header_latitude**: 3 - **header_longitude**: 4  <!-- theme: warning --> > #### Watch Out! > > When uploading your inventory it's important to make sure the process is complete before trying to start mapping. > Depending on the size of your file, this can take several minutes.<br><br> > You can use the [inventory details endpoint](https://nuitee2.stoplight.io/docs/cupid/b3A6NDg1OTM1MTA-inventory-details) to check the processing status. > If the inventory is still processing, the mapping endpoints will respond with an error 404.  When you have an inventory with ```active=true``` and ```mapping_process_status_id=2``` you can use the following endpoint to perfom mapping:  - [Map a list of properties](https://nuitee2.stoplight.io/docs/cupid/b3A6NDUxMTM0NTA-map-hotel-list) <!-- - [Map a CSV file](https://nuitee2.stoplight.io/docs/cupid/b3A6NDQ3ODY2OTY-upload-user-catalog) -->

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File*****os.File**|  | 
  **name** | **string**|  | 
  **headerId** | **int32**|  | 
  **headerName** | **int32**|  | 
  **headerAddress** | **int32**|  | 
  **headerCity** | **int32**|  | 
  **headerCountryCode** | **int32**|  | 
  **headerLatitude** | **int32**|  | 
  **headerLongitude** | **int32**|  | 

### Return type

[**UploadInventoryResponse**](UploadInventoryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


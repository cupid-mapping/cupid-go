package cupid

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/antihax/optional"
	"github.com/cupid-mapping/golang/models"
)

// Linger please
var (
	_ context.Context
)

type DefaultApiService service

/*
DefaultApiService Inventory details
This endpoint can be used to check the requested inventory&#x27;s details.  You can use it to get the inventory status &#x60;&#x60;&#x60;mapping_process_status_id&#x60;&#x60;&#x60; as outlined below. The status has to be 2 (Done) to start mapping.  **Possible values for the inventory&#x27;s status &#x60;&#x60;&#x60;mapping_process_status_id&#x60;&#x60;&#x60;:**  | value     | Name           | Description                                        | Action                                                     | | --------- | -------------- | -------------------------------------------------- | ---------------------------------------------------------- | | -1        | **Invalid**    | The inventory doesn&#x27;t contain any valid hotels       | Correct your catalog and try again.                        | | 0         | **Pending**    | The inventory is being uploaded                      | Wait, no action required.                                  | | 1         | **Processing** | The inventory is being processed                     | Wait, no action required.                                  | | 2         | **Done**       | The process is complete                            | **You can start mapping.**                                 | | 3         | **Failed**     | There was an error while processing                | Retry and if the error persists, feel free to contact us.  |  When you have an inventory with &#x60;&#x60;&#x60;active&#x3D;true&#x60;&#x60;&#x60; and &#x60;&#x60;&#x60;mapping_process_status_id&#x3D;2&#x60;&#x60;&#x60; you can use the following endpoint to perfom mapping:  - [Map a list of properties](https://nuitee2.stoplight.io/docs/cupid/b3A6NDUxMTM0NTA-map-hotel-list) &lt;!-- - [Map a CSV file](https://nuitee2.stoplight.io/docs/cupid/b3A6NDQ3ODY2OTY-upload-user-catalog)--&gt;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param inventoryId Inventory ID
@return InlineResponse200
*/
func (a *DefaultApiService) InventoryDetails(ctx context.Context, inventoryId int32) (models.InlineResponse200, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/inventories/{inventory_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"inventory_id"+"}", fmt.Sprintf("%v", inventoryId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v models.InlineResponse2001
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.UnauthorizedErrorResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v models.NotFoundResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ServerErrorResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefaultApiService List inventories
This endpoint can be used to list the inventories uploaded to your workspace.&lt;br&gt;You can check their status and other basic info.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return InlineResponse200
*/
func (a *DefaultApiService) ListInventories(ctx context.Context) (models.InlineResponse200, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/inventories"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v models.InlineResponse200
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.UnauthorizedErrorResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefaultApiService Map hotel list
&lt;!-- theme: warning --&gt; &gt; #### Requirement &gt; &gt; Active inventory with &#x60;&#x60;&#x60;mapping_process_status_id&#x3D;2&#x60;&#x60;&#x60;, you can check this using the [inventory details endpoint](https://nuitee2.stoplight.io/docs/cupid/b3A6NDg1OTM1MTA-inventory-details).  This endpoint allows you to send a list of properties and maps it against your active inventory. You will get the resulting mappings in the response.  ### Limit You can send up to 1000 properties per request.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *DefaultApiMapHotelsOpts - Optional Parameters:
     * @param "Body" (optional.Interface of []HotelItem) -  List of hotels to map
     * @param "ExtendedInfo" (optional.Bool) -  If true, returns properties info along with the mapped IDs.&lt;br&gt; Otherwise, returns only the user properties IDs along with the mapped partner IDs.
@return MapHotelsResponse
*/

type DefaultApiMapHotelsOpts struct {
	Body         optional.Interface
	ExtendedInfo optional.Bool
}

func (a *DefaultApiService) MapHotels(ctx context.Context, localVarOptionals *DefaultApiMapHotelsOpts) (models.MapHotelsResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.MapHotelsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/map-hotels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ExtendedInfo.IsSet() {
		localVarQueryParams.Add("extended_info", parameterToString(localVarOptionals.ExtendedInfo.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v models.MapHotelsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v models.BadRequestResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.UnauthorizedErrorResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v models.NotFoundResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 409 {
			var v models.ConflictResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ServerErrorResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
DefaultApiService Upload inventory
This end point can be used to upload an inventory with a provided CSV file.&lt;br&gt; &lt;br&gt; You will have to provide the inventory&#x27;s name and for each column in your file, you&#x27;ll need match its index with the standard field names provided below.&lt;br&gt;  Let&#x27;s assume you have a CSV file that looks like the following table:  | HotelID     | Hotel_Name   | Hotel_Address | Hotel_Lat  | Hotel_Lon  | Country | City     | ... | | ----------- | ------------ | ------------- | ---------- | ---------- | ------- | -------- | --- | | 1408        | Smile Hotel  | 123 main st   | 5.419514   | 3.392694   | AU      | Adelaid  | ... | | 123         | Cupid resort | 456 north av  | 55.69254   | 37.3167887 | MA      | Zagoura  | ... |  In this case the form data would be: - **file**: The csv file to be uploaded - **name**: Cupid inventory - **header_id**: 0 - **header_name**: 1 - **header_address**: 2 - **header_city**: 6 - **header_country_code**: 5 - **header_latitude**: 3 - **header_longitude**: 4  &lt;!-- theme: warning --&gt; &gt; #### Watch Out! &gt; &gt; When uploading your inventory it&#x27;s important to make sure the process is complete before trying to start mapping. &gt; Depending on the size of your file, this can take several minutes.&lt;br&gt;&lt;br&gt; &gt; You can use the [inventory details endpoint](https://nuitee2.stoplight.io/docs/cupid/b3A6NDg1OTM1MTA-inventory-details) to check the processing status. &gt; If the inventory is still processing, the mapping endpoints will respond with an error 404.  When you have an inventory with &#x60;&#x60;&#x60;active&#x3D;true&#x60;&#x60;&#x60; and &#x60;&#x60;&#x60;mapping_process_status_id&#x3D;2&#x60;&#x60;&#x60; you can use the following endpoint to perfom mapping:  - [Map a list of properties](https://nuitee2.stoplight.io/docs/cupid/b3A6NDUxMTM0NTA-map-hotel-list) &lt;!-- - [Map a CSV file](https://nuitee2.stoplight.io/docs/cupid/b3A6NDQ3ODY2OTY-upload-user-catalog) --&gt;
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param file
 * @param inventoryName
 * @param columnsIndexOpts
@return UploadInventoryResponse
*/

type ColumnsIndexOpts struct {
	Id          int32
	Name        int32
	Address     int32
	City        int32
	CountryCode int32
	Latitude    int32
	Longitude   int32
}

func (a *DefaultApiService) UploadInventory(ctx context.Context, file *os.File, inventoryName string, columnsIndexOpts ColumnsIndexOpts) (models.UploadInventoryResponse, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue models.UploadInventoryResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/inventories/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFile := file
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	localVarFormParams.Add("name", parameterToString(inventoryName, ""))
	localVarFormParams.Add("header_id", strconv.Itoa(int(columnsIndexOpts.Id)))
	localVarFormParams.Add("header_name", strconv.Itoa(int(columnsIndexOpts.Name)))
	localVarFormParams.Add("header_address", strconv.Itoa(int(columnsIndexOpts.Address)))
	localVarFormParams.Add("header_city", strconv.Itoa(int(columnsIndexOpts.City)))
	localVarFormParams.Add("header_country_code", strconv.Itoa(int(columnsIndexOpts.CountryCode)))
	localVarFormParams.Add("header_latitude", strconv.Itoa(int(columnsIndexOpts.Latitude)))
	localVarFormParams.Add("header_longitude", strconv.Itoa(int(columnsIndexOpts.Longitude)))
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v models.UploadInventoryResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v models.BadRequestResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v models.UnauthorizedErrorResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v models.ServerErrorResult
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

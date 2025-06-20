# \DefaultApi

All URIs are relative to *https://api.onesignal.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelNotification**](DefaultApi.md#CancelNotification) | **Delete** /notifications/{notification_id} | Stop a scheduled or currently outgoing notification
[**CreateAlias**](DefaultApi.md#CreateAlias) | **Patch** /apps/{app_id}/users/by/{alias_label}/{alias_id}/identity | 
[**CreateAliasBySubscription**](DefaultApi.md#CreateAliasBySubscription) | **Patch** /apps/{app_id}/subscriptions/{subscription_id}/user/identity | 
[**CreateApp**](DefaultApi.md#CreateApp) | **Post** /apps | Create an app
[**CreateNotification**](DefaultApi.md#CreateNotification) | **Post** /notifications | Create notification
[**CreateSegment**](DefaultApi.md#CreateSegment) | **Post** /apps/{app_id}/segments | Create Segment
[**CreateSubscription**](DefaultApi.md#CreateSubscription) | **Post** /apps/{app_id}/users/by/{alias_label}/{alias_id}/subscriptions | 
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /apps/{app_id}/users | 
[**DeleteAlias**](DefaultApi.md#DeleteAlias) | **Delete** /apps/{app_id}/users/by/{alias_label}/{alias_id}/identity/{alias_label_to_delete} | 
[**DeleteSegment**](DefaultApi.md#DeleteSegment) | **Delete** /apps/{app_id}/segments/{segment_id} | Delete Segment
[**DeleteSubscription**](DefaultApi.md#DeleteSubscription) | **Delete** /apps/{app_id}/subscriptions/{subscription_id} | 
[**DeleteUser**](DefaultApi.md#DeleteUser) | **Delete** /apps/{app_id}/users/by/{alias_label}/{alias_id} | 
[**ExportEvents**](DefaultApi.md#ExportEvents) | **Post** /notifications/{notification_id}/export_events?app_id&#x3D;{app_id} | Export CSV of Events
[**ExportSubscriptions**](DefaultApi.md#ExportSubscriptions) | **Post** /players/csv_export?app_id&#x3D;{app_id} | Export CSV of Subscriptions
[**GetAliases**](DefaultApi.md#GetAliases) | **Get** /apps/{app_id}/users/by/{alias_label}/{alias_id}/identity | 
[**GetAliasesBySubscription**](DefaultApi.md#GetAliasesBySubscription) | **Get** /apps/{app_id}/subscriptions/{subscription_id}/user/identity | 
[**GetApp**](DefaultApi.md#GetApp) | **Get** /apps/{app_id} | View an app
[**GetApps**](DefaultApi.md#GetApps) | **Get** /apps | View apps
[**GetNotification**](DefaultApi.md#GetNotification) | **Get** /notifications/{notification_id} | View notification
[**GetNotificationHistory**](DefaultApi.md#GetNotificationHistory) | **Post** /notifications/{notification_id}/history | Notification History
[**GetNotifications**](DefaultApi.md#GetNotifications) | **Get** /notifications | View notifications
[**GetOutcomes**](DefaultApi.md#GetOutcomes) | **Get** /apps/{app_id}/outcomes | View Outcomes
[**GetSegments**](DefaultApi.md#GetSegments) | **Get** /apps/{app_id}/segments | Get Segments
[**GetUser**](DefaultApi.md#GetUser) | **Get** /apps/{app_id}/users/by/{alias_label}/{alias_id} | 
[**TransferSubscription**](DefaultApi.md#TransferSubscription) | **Patch** /apps/{app_id}/subscriptions/{subscription_id}/owner | 
[**UnsubscribeEmailWithToken**](DefaultApi.md#UnsubscribeEmailWithToken) | **Post** /apps/{app_id}/notifications/{notification_id}/unsubscribe | Unsubscribe with token
[**UpdateApp**](DefaultApi.md#UpdateApp) | **Put** /apps/{app_id} | Update an app
[**UpdateLiveActivity**](DefaultApi.md#UpdateLiveActivity) | **Post** /apps/{app_id}/live_activities/{activity_id}/notifications | Update a Live Activity via Push
[**UpdateSubscription**](DefaultApi.md#UpdateSubscription) | **Patch** /apps/{app_id}/subscriptions/{subscription_id} | 
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Patch** /apps/{app_id}/users/by/{alias_label}/{alias_id} | 



## CancelNotification

> GenericSuccessBoolResponse CancelNotification(ctx, notificationId).AppId(appId).Execute()

Stop a scheduled or currently outgoing notification



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    notificationId := "notificationId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CancelNotification(restAuth, notificationId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CancelNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelNotification`: GenericSuccessBoolResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CancelNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** |  | 


### Return type

[**GenericSuccessBoolResponse**](GenericSuccessBoolResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAlias

> UserIdentityBody CreateAlias(ctx, appId, aliasLabel, aliasId).UserIdentityBody(userIdentityBody).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    aliasLabel := "aliasLabel_example" // string | 
    aliasId := "aliasId_example" // string | 
    userIdentityBody := *onesignal.NewUserIdentityBody() // UserIdentityBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateAlias(restAuth, appId, aliasLabel, aliasId).UserIdentityBody(userIdentityBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlias`: UserIdentityBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**aliasLabel** | **string** |  | 
**aliasId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userIdentityBody** | [**UserIdentityBody**](UserIdentityBody.md) |  | 

### Return type

[**UserIdentityBody**](UserIdentityBody.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAliasBySubscription

> UserIdentityBody CreateAliasBySubscription(ctx, appId, subscriptionId).UserIdentityBody(userIdentityBody).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    subscriptionId := "subscriptionId_example" // string | 
    userIdentityBody := *onesignal.NewUserIdentityBody() // UserIdentityBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateAliasBySubscription(restAuth, appId, subscriptionId).UserIdentityBody(userIdentityBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAliasBySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAliasBySubscription`: UserIdentityBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAliasBySubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAliasBySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userIdentityBody** | [**UserIdentityBody**](UserIdentityBody.md) |  | 

### Return type

[**UserIdentityBody**](UserIdentityBody.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApp

> App CreateApp(ctx).App(app).Execute()

Create an app



### Authorization

[organization_api_key](../README.md#organization_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    app := *onesignal.NewApp() // App | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.CreateApp(orgAuth).App(app).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApp`: App
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | [**App**](App.md) |  | 

### Return type

[**App**](App.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotification

> CreateNotificationSuccessResponse CreateNotification(ctx).Notification(notification).Execute()

Create notification



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    notification := *onesignal.NewNotification("AppId_example") // Notification | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateNotification(restAuth).Notification(notification).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotification`: CreateNotificationSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notification** | [**Notification**](Notification.md) |  | 

### Return type

[**CreateNotificationSuccessResponse**](CreateNotificationSuccessResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSegment

> CreateSegmentSuccessResponse CreateSegment(ctx, appId).Segment(segment).Execute()

Create Segment



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    segment := *onesignal.NewSegment("Name_example", []onesignal.FilterExpression{onesignal.FilterExpression{Filter: onesignal.NewFilter()}}) // Segment |  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateSegment(restAuth, appId).Segment(segment).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSegment`: CreateSegmentSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segment** | [**Segment**](Segment.md) |  | 

### Return type

[**CreateSegmentSuccessResponse**](CreateSegmentSuccessResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> SubscriptionBody CreateSubscription(ctx, appId, aliasLabel, aliasId).SubscriptionBody(subscriptionBody).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    aliasLabel := "aliasLabel_example" // string | 
    aliasId := "aliasId_example" // string | 
    subscriptionBody := *onesignal.NewSubscriptionBody() // SubscriptionBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateSubscription(restAuth, appId, aliasLabel, aliasId).SubscriptionBody(subscriptionBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: SubscriptionBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**aliasLabel** | **string** |  | 
**aliasId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **subscriptionBody** | [**SubscriptionBody**](SubscriptionBody.md) |  | 

### Return type

[**SubscriptionBody**](SubscriptionBody.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> User CreateUser(ctx, appId).User(user).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    user := *onesignal.NewUser() // User | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateUser(restAuth, appId).User(user).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) |  | 

### Return type

[**User**](User.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlias

> UserIdentityBody DeleteAlias(ctx, appId, aliasLabel, aliasId, aliasLabelToDelete).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    aliasLabel := "aliasLabel_example" // string | 
    aliasId := "aliasId_example" // string | 
    aliasLabelToDelete := "aliasLabelToDelete_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.DeleteAlias(restAuth, appId, aliasLabel, aliasId, aliasLabelToDelete).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlias`: UserIdentityBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**aliasLabel** | **string** |  | 
**aliasId** | **string** |  | 
**aliasLabelToDelete** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**UserIdentityBody**](UserIdentityBody.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSegment

> GenericSuccessBoolResponse DeleteSegment(ctx, appId, segmentId).Execute()

Delete Segment



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    segmentId := "segmentId_example" // string | The segment_id can be found in the URL of the segment when viewing it in the dashboard.

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.DeleteSegment(restAuth, appId, segmentId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSegment`: GenericSuccessBoolResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 
**segmentId** | **string** | The segment_id can be found in the URL of the segment when viewing it in the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GenericSuccessBoolResponse**](GenericSuccessBoolResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscription

> DeleteSubscription(ctx, appId, subscriptionId).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    subscriptionId := "subscriptionId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.DeleteSubscription(restAuth, appId, subscriptionId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, appId, aliasLabel, aliasId).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    aliasLabel := "aliasLabel_example" // string | 
    aliasId := "aliasId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.DeleteUser(restAuth, appId, aliasLabel, aliasId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**aliasLabel** | **string** |  | 
**aliasId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportEvents

> ExportEventsSuccessResponse ExportEvents(ctx, notificationId).AppId(appId).Execute()

Export CSV of Events



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    notificationId := "notificationId_example" // string | The ID of the notification to export events from.
    appId := "appId_example" // string | The ID of the app that the notification belongs to.

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.ExportEvents(restAuth, notificationId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportEvents`: ExportEventsSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExportEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The ID of the notification to export events from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** | The ID of the app that the notification belongs to. | 

### Return type

[**ExportEventsSuccessResponse**](ExportEventsSuccessResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSubscriptions

> ExportSubscriptionsSuccessResponse ExportSubscriptions(ctx, appId).ExportSubscriptionsRequestBody(exportSubscriptionsRequestBody).Execute()

Export CSV of Subscriptions



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | The app ID that you want to export devices from
    exportSubscriptionsRequestBody := *onesignal.NewExportSubscriptionsRequestBody() // ExportSubscriptionsRequestBody |  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.ExportSubscriptions(restAuth, appId).ExportSubscriptionsRequestBody(exportSubscriptionsRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSubscriptions`: ExportSubscriptionsSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExportSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The app ID that you want to export devices from | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportSubscriptionsRequestBody** | [**ExportSubscriptionsRequestBody**](ExportSubscriptionsRequestBody.md) |  | 

### Return type

[**ExportSubscriptionsSuccessResponse**](ExportSubscriptionsSuccessResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAliases

> UserIdentityBody GetAliases(ctx, appId, aliasLabel, aliasId).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    aliasLabel := "aliasLabel_example" // string | 
    aliasId := "aliasId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetAliases(restAuth, appId, aliasLabel, aliasId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAliases`: UserIdentityBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAliases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**aliasLabel** | **string** |  | 
**aliasId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserIdentityBody**](UserIdentityBody.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAliasesBySubscription

> UserIdentityBody GetAliasesBySubscription(ctx, appId, subscriptionId).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    subscriptionId := "subscriptionId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetAliasesBySubscription(restAuth, appId, subscriptionId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAliasesBySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAliasesBySubscription`: UserIdentityBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAliasesBySubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAliasesBySubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserIdentityBody**](UserIdentityBody.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApp

> App GetApp(ctx, appId).Execute()

View an app



### Authorization

[organization_api_key](../README.md#organization_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | An app id

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.GetApp(orgAuth, appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApp`: App
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | An app id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**App**](App.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApps

> []App GetApps(ctx).Execute()

View apps



### Authorization

[organization_api_key](../README.md#organization_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.GetApps(orgAuth).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApps`: []App
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


### Return type

[**[]App**](App.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotification

> NotificationWithMeta GetNotification(ctx, notificationId).AppId(appId).Execute()

View notification



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    notificationId := "notificationId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetNotification(restAuth, notificationId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotification`: NotificationWithMeta
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** |  | 


### Return type

[**NotificationWithMeta**](NotificationWithMeta.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationHistory

> NotificationHistorySuccessResponse GetNotificationHistory(ctx, notificationId).GetNotificationHistoryRequestBody(getNotificationHistoryRequestBody).Execute()

Notification History



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    notificationId := "notificationId_example" // string | The \"id\" of the message found in the Notification object
    getNotificationHistoryRequestBody := *onesignal.NewGetNotificationHistoryRequestBody() // GetNotificationHistoryRequestBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetNotificationHistory(restAuth, notificationId).GetNotificationHistoryRequestBody(getNotificationHistoryRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotificationHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationHistory`: NotificationHistorySuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNotificationHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | The \&quot;id\&quot; of the message found in the Notification object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getNotificationHistoryRequestBody** | [**GetNotificationHistoryRequestBody**](GetNotificationHistoryRequestBody.md) |  | 

### Return type

[**NotificationHistorySuccessResponse**](NotificationHistorySuccessResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> NotificationSlice GetNotifications(ctx).AppId(appId).Limit(limit).Offset(offset).Kind(kind).Execute()

View notifications



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | The app ID that you want to view notifications from
    limit := int32(56) // int32 | How many notifications to return.  Max is 50.  Default is 50. (optional)
    offset := int32(56) // int32 | Page offset.  Default is 0.  Results are sorted by queued_at in descending order.  queued_at is a representation of the time that the notification was queued at. (optional)
    kind := int32(56) // int32 | Kind of notifications returned:   * unset - All notification types (default)   * `0` - Dashboard only   * `1` - API only   * `3` - Automated only  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetNotifications(restAuth).AppId(appId).Limit(limit).Offset(offset).Kind(kind).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotifications`: NotificationSlice
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | The app ID that you want to view notifications from | 
 **limit** | **int32** | How many notifications to return.  Max is 50.  Default is 50. | 
 **offset** | **int32** | Page offset.  Default is 0.  Results are sorted by queued_at in descending order.  queued_at is a representation of the time that the notification was queued at. | 
 **kind** | **int32** | Kind of notifications returned:   * unset - All notification types (default)   * &#x60;0&#x60; - Dashboard only   * &#x60;1&#x60; - API only   * &#x60;3&#x60; - Automated only  | 

### Return type

[**NotificationSlice**](NotificationSlice.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutcomes

> OutcomesData GetOutcomes(ctx, appId).OutcomeNames(outcomeNames).OutcomeNames2(outcomeNames2).OutcomeTimeRange(outcomeTimeRange).OutcomePlatforms(outcomePlatforms).OutcomeAttribution(outcomeAttribution).Execute()

View Outcomes



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    outcomeNames := "outcomeNames_example" // string | Required Comma-separated list of names and the value (sum/count) for the returned outcome data. Note: Clicks only support count aggregation. For out-of-the-box OneSignal outcomes such as click and session duration, please use the \"os\" prefix with two underscores. For other outcomes, please use the name specified by the user. Example:os__session_duration.count,os__click.count,CustomOutcomeName.sum 
    outcomeNames2 := "outcomeNames_example" // string | Optional If outcome names contain any commas, then please specify only one value at a time. Example: outcome_names[]=os__click.count&outcome_names[]=Sales, Purchase.count where \"Sales, Purchase\" is the custom outcomes with a comma in the name.  (optional)
    outcomeTimeRange := "outcomeTimeRange_example" // string | Optional Time range for the returned data. The values can be 1h (for the last 1 hour data), 1d (for the last 1 day data), or 1mo (for the last 1 month data). Default is 1h if the parameter is omitted.  (optional)
    outcomePlatforms := "outcomePlatforms_example" // string | Optional Platform id. Refer device's platform ids for values. Example: outcome_platform=0 for iOS outcome_platform=7,8 for Safari and Firefox Default is data from all platforms if the parameter is omitted.  (optional)
    outcomeAttribution := "outcomeAttribution_example" // string | Optional Attribution type for the outcomes. The values can be direct or influenced or unattributed. Example: outcome_attribution=direct Default is total (returns direct+influenced+unattributed) if the parameter is omitted.  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetOutcomes(restAuth, appId).OutcomeNames(outcomeNames).OutcomeNames2(outcomeNames2).OutcomeTimeRange(outcomeTimeRange).OutcomePlatforms(outcomePlatforms).OutcomeAttribution(outcomeAttribution).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOutcomes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOutcomes`: OutcomesData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetOutcomes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutcomesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **outcomeNames** | **string** | Required Comma-separated list of names and the value (sum/count) for the returned outcome data. Note: Clicks only support count aggregation. For out-of-the-box OneSignal outcomes such as click and session duration, please use the \&quot;os\&quot; prefix with two underscores. For other outcomes, please use the name specified by the user. Example:os__session_duration.count,os__click.count,CustomOutcomeName.sum  | 
 **outcomeNames2** | **string** | Optional If outcome names contain any commas, then please specify only one value at a time. Example: outcome_names[]&#x3D;os__click.count&amp;outcome_names[]&#x3D;Sales, Purchase.count where \&quot;Sales, Purchase\&quot; is the custom outcomes with a comma in the name.  | 
 **outcomeTimeRange** | **string** | Optional Time range for the returned data. The values can be 1h (for the last 1 hour data), 1d (for the last 1 day data), or 1mo (for the last 1 month data). Default is 1h if the parameter is omitted.  | 
 **outcomePlatforms** | **string** | Optional Platform id. Refer device&#39;s platform ids for values. Example: outcome_platform&#x3D;0 for iOS outcome_platform&#x3D;7,8 for Safari and Firefox Default is data from all platforms if the parameter is omitted.  | 
 **outcomeAttribution** | **string** | Optional Attribution type for the outcomes. The values can be direct or influenced or unattributed. Example: outcome_attribution&#x3D;direct Default is total (returns direct+influenced+unattributed) if the parameter is omitted.  | 

### Return type

[**OutcomesData**](OutcomesData.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSegments

> GetSegmentsSuccessResponse GetSegments(ctx, appId).Offset(offset).Limit(limit).Execute()

Get Segments



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    offset := int32(56) // int32 | Segments are listed in ascending order of created_at date. offset will omit that number of segments from the beginning of the list. Eg offset 5, will remove the 5 earliest created Segments. (optional)
    limit := int32(56) // int32 | The amount of Segments in the response. Maximum 300. (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetSegments(restAuth, appId).Offset(offset).Limit(limit).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSegments`: GetSegmentsSuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Segments are listed in ascending order of created_at date. offset will omit that number of segments from the beginning of the list. Eg offset 5, will remove the 5 earliest created Segments. | 
 **limit** | **int32** | The amount of Segments in the response. Maximum 300. | 

### Return type

[**GetSegmentsSuccessResponse**](GetSegmentsSuccessResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, appId, aliasLabel, aliasId).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    aliasLabel := "aliasLabel_example" // string | 
    aliasId := "aliasId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetUser(restAuth, appId, aliasLabel, aliasId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**aliasLabel** | **string** |  | 
**aliasId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**User**](User.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferSubscription

> UserIdentityBody TransferSubscription(ctx, appId, subscriptionId).TransferSubscriptionRequestBody(transferSubscriptionRequestBody).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    subscriptionId := "subscriptionId_example" // string | 
    transferSubscriptionRequestBody := *onesignal.NewTransferSubscriptionRequestBody() // TransferSubscriptionRequestBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.TransferSubscription(restAuth, appId, subscriptionId).TransferSubscriptionRequestBody(transferSubscriptionRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TransferSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferSubscription`: UserIdentityBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TransferSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transferSubscriptionRequestBody** | [**TransferSubscriptionRequestBody**](TransferSubscriptionRequestBody.md) |  | 

### Return type

[**UserIdentityBody**](UserIdentityBody.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeEmailWithToken

> GenericSuccessBoolResponse UnsubscribeEmailWithToken(ctx, appId, notificationId).Token(token).Execute()

Unsubscribe with token



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    notificationId := "notificationId_example" // string | The id of the message found in the creation notification POST response, View Notifications GET response, or URL within the Message Report.
    token := "token_example" // string | The unsubscribe token that is generated via liquid syntax in {{subscription.unsubscribe_token}} when personalizing an email.

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UnsubscribeEmailWithToken(restAuth, appId, notificationId).Token(token).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UnsubscribeEmailWithToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsubscribeEmailWithToken`: GenericSuccessBoolResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UnsubscribeEmailWithToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 
**notificationId** | **string** | The id of the message found in the creation notification POST response, View Notifications GET response, or URL within the Message Report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeEmailWithTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **string** | The unsubscribe token that is generated via liquid syntax in {{subscription.unsubscribe_token}} when personalizing an email. | 

### Return type

[**GenericSuccessBoolResponse**](GenericSuccessBoolResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApp

> App UpdateApp(ctx, appId).App(app).Execute()

Update an app



### Authorization

[organization_api_key](../README.md#organization_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | An app id
    app := *onesignal.NewApp() // App | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.UpdateApp(orgAuth, appId).App(app).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApp`: App
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | An app id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **app** | [**App**](App.md) |  | 

### Return type

[**App**](App.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLiveActivity

> UpdateLiveActivitySuccessResponse UpdateLiveActivity(ctx, appId, activityId).UpdateLiveActivityRequest(updateLiveActivityRequest).Execute()

Update a Live Activity via Push



### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    activityId := "activityId_example" // string | Live Activity record ID
    updateLiveActivityRequest := *onesignal.NewUpdateLiveActivityRequest("Name_example", "Event_example", map[string]interface{}(123)) // UpdateLiveActivityRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UpdateLiveActivity(restAuth, appId, activityId).UpdateLiveActivityRequest(updateLiveActivityRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateLiveActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLiveActivity`: UpdateLiveActivitySuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateLiveActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 
**activityId** | **string** | Live Activity record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLiveActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateLiveActivityRequest** | [**UpdateLiveActivityRequest**](UpdateLiveActivityRequest.md) |  | 

### Return type

[**UpdateLiveActivitySuccessResponse**](UpdateLiveActivitySuccessResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> UpdateSubscription(ctx, appId, subscriptionId).SubscriptionBody(subscriptionBody).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    subscriptionId := "subscriptionId_example" // string | 
    subscriptionBody := *onesignal.NewSubscriptionBody() // SubscriptionBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UpdateSubscription(restAuth, appId, subscriptionId).SubscriptionBody(subscriptionBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subscriptionBody** | [**SubscriptionBody**](SubscriptionBody.md) |  | 

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> PropertiesBody UpdateUser(ctx, appId, aliasLabel, aliasId).UpdateUserRequest(updateUserRequest).Execute()





### Authorization

[rest_api_key](../README.md#rest_api_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api"
)

func main() {
    appId := "appId_example" // string | 
    aliasLabel := "aliasLabel_example" // string | 
    aliasId := "aliasId_example" // string | 
    updateUserRequest := *onesignal.NewUpdateUserRequest() // UpdateUserRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UpdateUser(restAuth, appId, aliasLabel, aliasId).UpdateUserRequest(updateUserRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: PropertiesBody
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**aliasLabel** | **string** |  | 
**aliasId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

### Return type

[**PropertiesBody**](PropertiesBody.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


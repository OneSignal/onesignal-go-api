# \DefaultApi

All URIs are relative to *https://onesignal.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelNotification**](DefaultApi.md#CancelNotification) | **Delete** /notifications/{notification_id} | Stop a scheduled or currently outgoing notification
[**CreateApp**](DefaultApi.md#CreateApp) | **Post** /apps | Create an app
[**CreateNotification**](DefaultApi.md#CreateNotification) | **Post** /notifications | Create notification
[**CreatePlayer**](DefaultApi.md#CreatePlayer) | **Post** /players | Add a device
[**CreateSegments**](DefaultApi.md#CreateSegments) | **Post** /apps/{app_id}/segments | Create Segments
[**DeletePlayer**](DefaultApi.md#DeletePlayer) | **Delete** /players/{player_id} | Delete a user record
[**DeleteSegments**](DefaultApi.md#DeleteSegments) | **Delete** /apps/{app_id}/segments/{segment_id} | Delete Segments
[**ExportPlayers**](DefaultApi.md#ExportPlayers) | **Post** /players/csv_export?app_id&#x3D;{app_id} | CSV export
[**GetApp**](DefaultApi.md#GetApp) | **Get** /apps/{app_id} | View an app
[**GetApps**](DefaultApi.md#GetApps) | **Get** /apps | View apps
[**GetNotification**](DefaultApi.md#GetNotification) | **Get** /notifications/{notification_id} | View notification
[**GetNotificationHistory**](DefaultApi.md#GetNotificationHistory) | **Post** /notifications/{notification_id}/history | Notification History
[**GetNotifications**](DefaultApi.md#GetNotifications) | **Get** /notifications | View notifications
[**GetOutcomes**](DefaultApi.md#GetOutcomes) | **Get** /apps/{app_id}/outcomes | View Outcomes
[**GetPlayer**](DefaultApi.md#GetPlayer) | **Get** /players/{player_id} | View device
[**GetPlayers**](DefaultApi.md#GetPlayers) | **Get** /players | View devices
[**UpdateApp**](DefaultApi.md#UpdateApp) | **Put** /apps/{app_id} | Update an app
[**UpdatePlayer**](DefaultApi.md#UpdatePlayer) | **Put** /players/{player_id} | Edit device
[**UpdatePlayerTags**](DefaultApi.md#UpdatePlayerTags) | **Put** /apps/{app_id}/users/{external_user_id} | Edit tags with external user id



## CancelNotification

> InlineResponse2001 CancelNotification(ctx, notificationId).AppId(appId).Execute()

Stop a scheduled or currently outgoing notification



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | 
    notificationId := "notificationId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.CancelNotification(appAuth, notificationId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CancelNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelNotification`: InlineResponse2001
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

[**InlineResponse2001**](InlineResponse2001.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApp

> App CreateApp(ctx).App(app).Execute()

Create an app



### Authorization

[user_key](../README.md#user_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    app := *openapiclient.NewApp("Id_example") // App | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    userAuth := context.WithValue(context.Background(), onesignal.UserAuth, "USER_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.CreateApp(userAuth).App(app).Execute()

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

> InlineResponse200 CreateNotification(ctx).Notification(notification).Execute()

Create notification



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    notification := *openapiclient.NewNotification("AppId_example") // Notification | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.CreateNotification(appAuth).Notification(notification).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotification`: InlineResponse200
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

[**InlineResponse200**](InlineResponse200.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlayer

> InlineResponse2004 CreatePlayer(ctx).Player(player).Execute()

Add a device



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    player := *openapiclient.NewPlayer("Id_example", "AppId_example", int32(123)) // Player | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.CreatePlayer(appAuth).Player(player).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlayer`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePlayer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **player** | [**Player**](Player.md) |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSegments

> InlineResponse201 CreateSegments(ctx, appId).Segment(segment).Execute()

Create Segments



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    segment := *openapiclient.NewSegment("Name_example", []openapiclient.FilterExpressions{*openapiclient.NewFilterExpressions("Field_example", "Relation_example")}) // Segment |  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.CreateSegments(appAuth, appId).Segment(segment).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSegments`: InlineResponse201
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **segment** | [**Segment**](Segment.md) |  | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlayer

> InlineResponse2001 DeletePlayer(ctx, playerId).AppId(appId).Execute()

Delete a user record



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    playerId := "playerId_example" // string | The OneSignal player_id

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.DeletePlayer(appAuth, playerId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePlayer`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeletePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | The OneSignal player_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 


### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSegments

> InlineResponse2003 DeleteSegments(ctx, appId, segmentId).Execute()

Delete Segments



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    segmentId := "segmentId_example" // string | The segment_id can be found in the URL of the segment when viewing it in the dashboard.

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.DeleteSegments(appAuth, appId, segmentId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSegments`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID for your app.  Available in Keys &amp; IDs. | 
**segmentId** | **string** | The segment_id can be found in the URL of the segment when viewing it in the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPlayers

> InlineResponse2005 ExportPlayers(ctx, appId).ExportPlayersRequestBody(exportPlayersRequestBody).Execute()

CSV export



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | The app ID that you want to export devices from
    exportPlayersRequestBody := *openapiclient.NewExportPlayersRequestBody() // ExportPlayersRequestBody |  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.ExportPlayers(appAuth, appId).ExportPlayersRequestBody(exportPlayersRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportPlayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPlayers`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExportPlayers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The app ID that you want to export devices from | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPlayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportPlayersRequestBody** | [**ExportPlayersRequestBody**](ExportPlayersRequestBody.md) |  | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApp

> App GetApp(ctx, appId).Execute()

View an app



### Authorization

[user_key](../README.md#user_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | An app id

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    userAuth := context.WithValue(context.Background(), onesignal.UserAuth, "USER_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.GetApp(userAuth, appId).Execute()

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

> string GetApps(ctx).Execute()

View apps



### Authorization

[user_key](../README.md#user_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    userAuth := context.WithValue(context.Background(), onesignal.UserAuth, "USER_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.GetApps(userAuth).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApps`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppsRequest struct via the builder pattern


### Return type

**string**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotification

> Notification GetNotification(ctx, notificationId).AppId(appId).Execute()

View notification



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | 
    notificationId := "notificationId_example" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.GetNotification(appAuth, notificationId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotification`: Notification
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

[**Notification**](Notification.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationHistory

> InlineResponse2002 GetNotificationHistory(ctx, notificationId).GetNotificationRequestBody(getNotificationRequestBody).Execute()

Notification History



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    notificationId := "notificationId_example" // string | The \"id\" of the message found in the Notification object
    getNotificationRequestBody := *openapiclient.NewGetNotificationRequestBody() // GetNotificationRequestBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.GetNotificationHistory(appAuth, notificationId).GetNotificationRequestBody(getNotificationRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotificationHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationHistory`: InlineResponse2002
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

 **getNotificationRequestBody** | [**GetNotificationRequestBody**](GetNotificationRequestBody.md) |  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

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

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | The app ID that you want to view notifications from
    limit := "limit_example" // string | How many notifications to return.  Max is 50.  Default is 50. (optional)
    offset := int32(56) // int32 | Page offset.  Default is 0.  Results are sorted by queued_at in descending order.  queued_at is a representation of the time that the notification was queued at. (optional)
    kind := int32(56) // int32 | Kind of notifications returned:   * unset - All notification types (default)   * `0` - Dashboard only   * `1` - API only   * `3` - Automated only  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.GetNotifications(appAuth).AppId(appId).Limit(limit).Offset(offset).Kind(kind).Execute()

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
 **limit** | **string** | How many notifications to return.  Max is 50.  Default is 50. | 
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

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
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

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.GetOutcomes(appAuth, appId).OutcomeNames(outcomeNames).OutcomeNames2(outcomeNames2).OutcomeTimeRange(outcomeTimeRange).OutcomePlatforms(outcomePlatforms).OutcomeAttribution(outcomeAttribution).Execute()

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


## GetPlayer

> Player GetPlayer(ctx, playerId).AppId(appId).EmailAuthHash(emailAuthHash).Execute()

View device



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | Your app_id for this device
    playerId := "playerId_example" // string | Player's OneSignal ID
    emailAuthHash := "emailAuthHash_example" // string | Email - Only required if you have enabled Identity Verification and device_type is email (11). (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.GetPlayer(appAuth, playerId).AppId(appId).EmailAuthHash(emailAuthHash).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayer`: Player
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | Player&#39;s OneSignal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | Your app_id for this device | 

 **emailAuthHash** | **string** | Email - Only required if you have enabled Identity Verification and device_type is email (11). | 

### Return type

[**Player**](Player.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlayers

> PlayerSlice GetPlayers(ctx).AppId(appId).Limit(limit).Offset(offset).Execute()

View devices



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | The app ID that you want to view players from
    limit := "limit_example" // string | How many devices to return. Max is 300. Default is 300 (optional)
    offset := int32(56) // int32 | Result offset. Default is 0. Results are sorted by id; (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.GetPlayers(appAuth).AppId(appId).Limit(limit).Offset(offset).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPlayers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlayers`: PlayerSlice
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPlayers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlayersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | The app ID that you want to view players from | 
 **limit** | **string** | How many devices to return. Max is 300. Default is 300 | 
 **offset** | **int32** | Result offset. Default is 0. Results are sorted by id; | 

### Return type

[**PlayerSlice**](PlayerSlice.md)

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

[user_key](../README.md#user_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | An app id
    app := *openapiclient.NewApp("Id_example") // App | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    userAuth := context.WithValue(context.Background(), onesignal.UserAuth, "USER_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.UpdateApp(userAuth, appId).App(app).Execute()

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


## UpdatePlayer

> InlineResponse2001 UpdatePlayer(ctx, playerId).Player(player).Execute()

Edit device



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    playerId := "playerId_example" // string | Player's OneSignal ID
    player := *openapiclient.NewPlayer("Id_example", "AppId_example", int32(123)) // Player | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.UpdatePlayer(appAuth, playerId).Player(player).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePlayer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlayer`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdatePlayer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **string** | Player&#39;s OneSignal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlayerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **player** | [**Player**](Player.md) |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlayerTags

> InlineResponse2003 UpdatePlayerTags(ctx, appId, externalUserId).UpdatePlayerTagsRequestBody(updatePlayerTagsRequestBody).Execute()

Edit tags with external user id



### Authorization

[app_key](../README.md#app_key)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "onesignal"
)

func main() {
    appId := "appId_example" // string | The OneSignal App ID the user record is found under.
    externalUserId := "externalUserId_example" // string | The External User ID mapped to teh device record in OneSignal.  Must be actively set on the device to be updated.
    updatePlayerTagsRequestBody := *openapiclient.NewUpdatePlayerTagsRequestBody() // UpdatePlayerTagsRequestBody |  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")

    resp, r, err := apiClient.DefaultApi.UpdatePlayerTags(appAuth, appId, externalUserId).UpdatePlayerTagsRequestBody(updatePlayerTagsRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdatePlayerTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlayerTags`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdatePlayerTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The OneSignal App ID the user record is found under. | 
**externalUserId** | **string** | The External User ID mapped to teh device record in OneSignal.  Must be actively set on the device to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlayerTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePlayerTagsRequestBody** | [**UpdatePlayerTagsRequestBody**](UpdatePlayerTagsRequestBody.md) |  | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


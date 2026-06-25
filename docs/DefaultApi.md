# \DefaultApi

All URIs are relative to *https://api.onesignal.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelNotification**](DefaultApi.md#CancelNotification) | **Delete** /notifications/{notification_id} | Stop a scheduled or currently outgoing notification
[**CopyTemplateToApp**](DefaultApi.md#CopyTemplateToApp) | **Post** /templates/{template_id}/copy_to_app | Copy template to another app
[**CreateAlias**](DefaultApi.md#CreateAlias) | **Patch** /apps/{app_id}/users/by/{alias_label}/{alias_id}/identity | 
[**CreateAliasBySubscription**](DefaultApi.md#CreateAliasBySubscription) | **Patch** /apps/{app_id}/subscriptions/{subscription_id}/user/identity | 
[**CreateApiKey**](DefaultApi.md#CreateApiKey) | **Post** /apps/{app_id}/auth/tokens | Create API key
[**CreateApp**](DefaultApi.md#CreateApp) | **Post** /apps | Create an app
[**CreateCustomEvents**](DefaultApi.md#CreateCustomEvents) | **Post** /apps/{app_id}/integrations/custom_events | Create custom events
[**CreateNotification**](DefaultApi.md#CreateNotification) | **Post** /notifications | Create notification
[**CreateSegment**](DefaultApi.md#CreateSegment) | **Post** /apps/{app_id}/segments | Create Segment
[**CreateSubscription**](DefaultApi.md#CreateSubscription) | **Post** /apps/{app_id}/users/by/{alias_label}/{alias_id}/subscriptions | 
[**CreateTemplate**](DefaultApi.md#CreateTemplate) | **Post** /templates | Create template
[**CreateUser**](DefaultApi.md#CreateUser) | **Post** /apps/{app_id}/users | 
[**DeleteAlias**](DefaultApi.md#DeleteAlias) | **Delete** /apps/{app_id}/users/by/{alias_label}/{alias_id}/identity/{alias_label_to_delete} | 
[**DeleteApiKey**](DefaultApi.md#DeleteApiKey) | **Delete** /apps/{app_id}/auth/tokens/{token_id} | Delete API key
[**DeleteSegment**](DefaultApi.md#DeleteSegment) | **Delete** /apps/{app_id}/segments/{segment_id} | Delete Segment
[**DeleteSubscription**](DefaultApi.md#DeleteSubscription) | **Delete** /apps/{app_id}/subscriptions/{subscription_id} | 
[**DeleteTemplate**](DefaultApi.md#DeleteTemplate) | **Delete** /templates/{template_id} | Delete template
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
[**RotateApiKey**](DefaultApi.md#RotateApiKey) | **Post** /apps/{app_id}/auth/tokens/{token_id}/rotate | Rotate API key
[**StartLiveActivity**](DefaultApi.md#StartLiveActivity) | **Post** /apps/{app_id}/activities/activity/{activity_type} | Start Live Activity
[**TransferSubscription**](DefaultApi.md#TransferSubscription) | **Patch** /apps/{app_id}/subscriptions/{subscription_id}/owner | 
[**UnsubscribeEmailWithToken**](DefaultApi.md#UnsubscribeEmailWithToken) | **Post** /apps/{app_id}/notifications/{notification_id}/unsubscribe | Unsubscribe with token
[**UpdateApiKey**](DefaultApi.md#UpdateApiKey) | **Patch** /apps/{app_id}/auth/tokens/{token_id} | Update API key
[**UpdateApp**](DefaultApi.md#UpdateApp) | **Put** /apps/{app_id} | Update an app
[**UpdateLiveActivity**](DefaultApi.md#UpdateLiveActivity) | **Post** /apps/{app_id}/live_activities/{activity_id}/notifications | Update a Live Activity via Push
[**UpdateSubscription**](DefaultApi.md#UpdateSubscription) | **Patch** /apps/{app_id}/subscriptions/{subscription_id} | 
[**UpdateSubscriptionByToken**](DefaultApi.md#UpdateSubscriptionByToken) | **Patch** /apps/{app_id}/subscriptions_by_token/{token_type}/{token} | Update subscription by token
[**UpdateTemplate**](DefaultApi.md#UpdateTemplate) | **Patch** /templates/{template_id} | Update template
[**UpdateUser**](DefaultApi.md#UpdateUser) | **Patch** /apps/{app_id}/users/by/{alias_label}/{alias_id} | 
[**ViewApiKeys**](DefaultApi.md#ViewApiKeys) | **Get** /apps/{app_id}/auth/tokens | View API keys
[**ViewTemplate**](DefaultApi.md#ViewTemplate) | **Get** /templates/{template_id} | View template
[**ViewTemplates**](DefaultApi.md#ViewTemplates) | **Get** /templates | View templates


## Common patterns

The per-endpoint examples below illustrate one specific call each. This section covers patterns that apply to most operations.

### Authentication

Every operation requires either a **REST API Key** (App-scoped, used by ~77% of endpoints) or an **Organization API Key** (used by the remaining ~23% — the app-management endpoints `GetApps` / `CreateApp` / `GetApp` / `UpdateApp` / `CopyTemplateToApp`, plus the API-key administration endpoints `ViewApiKeys` / `CreateApiKey` / `DeleteApiKey` / `UpdateApiKey` / `RotateApiKey`). The two are not interchangeable. The "Authorization" row on each endpoint below lists the exact scheme.

### Idempotency

`POST /notifications` accepts a top-level `idempotency_key` (UUIDv4) that the server uses for request dedup within a **30-day window**. Pass a freshly-generated UUID per logical send so that network-level retries are safe. Never reuse a key across distinct sends — the server returns the original response instead of acting on the new payload. The hero `CreateNotification` example below demonstrates the call.

Prefer the bundled `CreateNotificationWithRetry` helper over wiring this up by hand: it generates the key when absent (a caller-provided key is respected), retries 429 / 503 / transport errors with the **same** key (honoring `Retry-After`, exponential backoff otherwise; `MaxRetries` / `BaseDelay` configurable via the options struct), fails fast on other errors, and reports via `WasReplayed` whether the server answered from a previously completed request (`Idempotent-Replayed` response header). It is available as a `DefaultApi` method so the call mirrors `CreateNotification`:

```go
result, err := apiClient.DefaultApi.CreateNotificationWithRetry(context.Background(), *notification, nil)
if err == nil {
    fmt.Printf("%s replayed=%v\n", result.Response.GetId(), result.WasReplayed)
}
```

### Error handling

When a request fails, the call returns a non-nil `err`. The raw `*http.Response` (the second return value, conventionally `r`) is non-nil only when the failure occurred AFTER the HTTP round-trip completed — for pre-roundtrip failures (URL build, missing required parameter, prepareRequest error, transport errors like DNS/timeout/TLS) `r` is `nil`. Always nil-check `r` before reading `r.StatusCode` (int). For the parsed error body, type-assert `err` to `*onesignal.GenericOpenAPIError` (pointer — the SDK constructs every error as `&GenericOpenAPIError{...}`, so a value-type assertion silently returns `ok=false`) and call `apiErr.Body()` (returns `[]byte`, the raw JSON envelope). Most envelopes match `{ "errors": ["..."] }` (an array of strings) but a few endpoints return `{ "errors": [{"code": ..., "title": ..., "meta": {...}}] }` (an array of structured error objects — used by `POST /apps/{app_id}/users` 409 conflict, see `CreateUserConflictResponse`), `{ "errors": "..." }` (string), or `{ "success": false }` (no `errors` field at all). Robust error-handling code should tolerate all four shapes. The `apiErr.ErrorMessages()` method does this for you, normalizing every shape to a flat `[]string` (empty when the body carries no `errors`). To branch on a specific error without hard-coding message strings, test membership against the generated [`OneSignalErrors`](https://github.com/OneSignal/onesignal-go-api/blob/main/error_catalog.go) catalog — e.g. check whether `apiErr.ErrorMessages()` contains `onesignal.NO_TARGETING_SPECIFIED`.

### Polymorphic 200 from POST /notifications

`CreateNotificationSuccessResponse` has two distinct shapes that share the same schema; branch on `id`:
- **Success** — `id` is a non-empty UUID. `errors`, if present, is an object keyed by recipient-identifier type (`invalid_player_ids`, `invalid_external_user_ids`, `invalid_aliases`, ...) listing recipients that were skipped (partial-success path).
- **No-send** — `id` is the empty string `""`. `errors` is a string array with the sentinel reason, typically `["All included players are not subscribed"]`.

The hero `CreateNotification` example below demonstrates the branch pattern explicitly.

### Targeting users by External ID

Set `include_aliases.external_id` to a list of External IDs and set `target_channel` to `push` / `email` / `sms`. The alias label must be the literal string `external_id` — camelCase variants such as `externalId` are silently ignored and yield zero recipients. **Do not confuse** this with the deprecated top-level `external_id` notification field — a separate correlation/idempotency field with its own 30-day dedup keyspace (parallel to `idempotency_key`, not an alias) and no targeting effect.


## CancelNotification

> GenericSuccessBoolResponse CancelNotification(ctx, notificationId).AppId(appId).Execute()

Stop a scheduled or currently outgoing notification



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    notificationId := "b3a0c8bd-3a4c-4b22-9a73-3f1a8c2d1b88" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CancelNotification(restAuth, notificationId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CancelNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CopyTemplateToApp

> TemplateResource CopyTemplateToApp(ctx, templateId).AppId(appId).CopyTemplateRequest(copyTemplateRequest).Execute()

Copy template to another app



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    templateId := "e4d3c2b1-a09f-4f1e-8d7c-6b5a4f3e2d1c" // string | 
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    copyTemplateRequest := *onesignal.NewCopyTemplateRequest("TargetAppId_example") // CopyTemplateRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.CopyTemplateToApp(orgAuth, templateId).AppId(appId).CopyTemplateRequest(copyTemplateRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CopyTemplateToApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `CopyTemplateToApp`: TemplateResource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CopyTemplateToApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyTemplateToAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** |  | 
 **copyTemplateRequest** | [**CopyTemplateRequest**](CopyTemplateRequest.md) |  | 

### Return type

[**TemplateResource**](TemplateResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateAlias

> UserIdentityBody CreateAlias(ctx, appId, aliasLabel, aliasId).UserIdentityBody(userIdentityBody).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    aliasLabel := "external_id" // string | 
    aliasId := "YOUR_USER_EXTERNAL_ID" // string | 
    userIdentityBody := *onesignal.NewUserIdentityBody() // UserIdentityBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateAlias(restAuth, appId, aliasLabel, aliasId).UserIdentityBody(userIdentityBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateAliasBySubscription

> UserIdentityBody CreateAliasBySubscription(ctx, appId, subscriptionId).UserIdentityBody(userIdentityBody).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    subscriptionId := "7e4c5b9a-1f60-4d07-9b1a-2e8c8d2cae51" // string | 
    userIdentityBody := *onesignal.NewUserIdentityBody() // UserIdentityBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateAliasBySubscription(restAuth, appId, subscriptionId).UserIdentityBody(userIdentityBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAliasBySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateApiKey

> CreateApiKeyResponse CreateApiKey(ctx, appId).CreateApiKeyRequest(createApiKeyRequest).Execute()

Create API key



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    createApiKeyRequest := *onesignal.NewCreateApiKeyRequest() // CreateApiKeyRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.CreateApiKey(orgAuth, appId).CreateApiKeyRequest(createApiKeyRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `CreateApiKey`: CreateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createApiKeyRequest** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md) |  | 

### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateApp

> App CreateApp(ctx).App(app).Execute()

Create an app



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
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
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateCustomEvents

> map[string]interface{} CreateCustomEvents(ctx, appId).CustomEventsRequest(customEventsRequest).Execute()

Create custom events



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | Your OneSignal App ID in UUID v4 format.
    customEventsRequest := *onesignal.NewCustomEventsRequest([]onesignal.CustomEvent{*onesignal.NewCustomEvent("Name_example")}) // CustomEventsRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateCustomEvents(restAuth, appId).CustomEventsRequest(customEventsRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCustomEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `CreateCustomEvents`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCustomEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Your OneSignal App ID in UUID v4 format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customEventsRequest** | [**CustomEventsRequest**](CustomEventsRequest.md) |  | 

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateNotification

> CreateNotificationSuccessResponse CreateNotification(ctx).Notification(notification).Execute()

Create notification



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/google/uuid"
    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY")

    notification := onesignal.NewNotification("YOUR_APP_ID")
    contents := onesignal.NewLanguageStringMap()
    contents.SetEn("Hello from OneSignal!")
    notification.SetContents(*contents)
    notification.SetIncludeAliases(map[string][]string{"external_id": {"YOUR_USER_EXTERNAL_ID"}})
    notification.SetTargetChannel("push")
    // Idempotency key: a client-generated UUID that lets you safely retry on network failure.
    // If two requests arrive with the same key inside the 30-day window, only the first is
    // sent and the second returns the original response. The `github.com/google/uuid` module
    // is not a declared dep of this SDK; run `go get github.com/google/uuid` (or `go mod tidy`
    // after importing it) before building. DO NOT reuse keys across logically distinct sends.
    notification.SetIdempotencyKey(uuid.NewString())

    resp, r, err := apiClient.DefaultApi.CreateNotification(restAuth).Notification(*notification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
        return
    }
    // `resp.GetId()` discriminates the two HTTP 200 shapes. An empty string means no
    // notification was created (e.g. all targets were unreachable / not subscribed).
    // `resp.GetErrors()` is `interface{}` because the field is polymorphic: a `[]string` in
    // the no-subscribers case, or a map keyed by recipient-identifier type
    // (`invalid_player_ids`, `invalid_external_user_ids`, `invalid_aliases`, ...) when
    // the notification WAS created but some recipients were skipped.
    if resp.GetId() == "" {
        fmt.Fprintf(os.Stderr, "Notification was not sent: %v\n", resp.GetErrors())
    } else if errors := resp.GetErrors(); errors != nil {
        fmt.Fprintf(os.Stdout, "Notification created: %s (partial failures: %v)\n", resp.GetId(), errors)
    } else {
        fmt.Fprintf(os.Stdout, "Notification created: %s\n", resp.GetId())
    }
}
```

#### Using `CreateNotificationWithRetry` (preferred for safe, idempotent retries)

The `CreateNotificationWithRetry` method mirrors `CreateNotification` but generates the `IdempotencyKey` for you, transparently retries transient failures (HTTP 429 / 503 / transport errors) with the **same** key, and reports via `WasReplayed` whether the server answered from a previously-completed request.

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY")

    notification := onesignal.NewNotification("YOUR_APP_ID")
    contents := onesignal.NewLanguageStringMap()
    contents.SetEn("Hello from OneSignal!")
    notification.SetContents(*contents)
    notification.SetIncludeAliases(map[string][]string{"external_id": {"YOUR_USER_EXTERNAL_ID"}})
    notification.SetTargetChannel("push")
    // No idempotency key set: the helper generates a UUIDv4 and reuses it across retries.

    // opts is optional — pass nil for defaults (3 retries, 1s backoff base).
    opts := &onesignal.CreateNotificationWithRetryOptions{MaxRetries: 5}
    result, err := apiClient.DefaultApi.CreateNotificationWithRetry(restAuth, *notification, opts)
    if err != nil {
        fmt.Fprintf(os.Stderr, "CreateNotificationWithRetry failed: %v\n", err)
        return
    }
    if result.WasReplayed {
        fmt.Fprintf(os.Stdout, "Server replayed a prior send (no duplicate): %s\n", result.Response.GetId())
    } else {
        fmt.Fprintf(os.Stdout, "Notification created: %s\n", result.Response.GetId())
    }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateSegment

> CreateSegmentSuccessResponse CreateSegment(ctx, appId).Segment(segment).Execute()

Create Segment



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    segment := *onesignal.NewSegment("Name_example", []onesignal.FilterExpression{onesignal.FilterExpression{Filter: onesignal.NewFilter()}}) // Segment |  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateSegment(restAuth, appId).Segment(segment).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateSubscription

> SubscriptionBody CreateSubscription(ctx, appId, aliasLabel, aliasId).SubscriptionBody(subscriptionBody).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    aliasLabel := "external_id" // string | 
    aliasId := "YOUR_USER_EXTERNAL_ID" // string | 
    subscriptionBody := *onesignal.NewSubscriptionBody() // SubscriptionBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateSubscription(restAuth, appId, aliasLabel, aliasId).SubscriptionBody(subscriptionBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateTemplate

> TemplateResource CreateTemplate(ctx).CreateTemplateRequest(createTemplateRequest).Execute()

Create template



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    createTemplateRequest := *onesignal.NewCreateTemplateRequest("AppId_example", "Name_example", *onesignal.NewLanguageStringMap()) // CreateTemplateRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateTemplate(restAuth).CreateTemplateRequest(createTemplateRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `CreateTemplate`: TemplateResource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTemplateRequest** | [**CreateTemplateRequest**](CreateTemplateRequest.md) |  | 

### Return type

[**TemplateResource**](TemplateResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## CreateUser

> User CreateUser(ctx, appId).User(user).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    user := *onesignal.NewUser() // User | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.CreateUser(restAuth, appId).User(user).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `CreateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUser`: %v\n", resp)
}
```

### Reading the 409 conflict metadata

A `409` from this endpoint deserializes to `CreateUserConflictResponse`. `ErrorMessages()` flattens each error to its `title`/`code` and omits the structured `meta` object (currently `conflicting_aliases`); read it from the typed model via `Model()`:

```go
if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
    if conflict, ok := apiErr.Model().(onesignal.CreateUserConflictResponse); ok {
        for _, e := range conflict.GetErrors() {
            meta := e.GetMeta()
            fmt.Println(e.GetTitle(), meta.GetConflictingAliases())
        }
    }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## DeleteAlias

> UserIdentityBody DeleteAlias(ctx, appId, aliasLabel, aliasId, aliasLabelToDelete).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    aliasLabel := "external_id" // string | 
    aliasId := "YOUR_USER_EXTERNAL_ID" // string | 
    aliasLabelToDelete := "external_id" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.DeleteAlias(restAuth, appId, aliasLabel, aliasId, aliasLabelToDelete).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## DeleteApiKey

> map[string]interface{} DeleteApiKey(ctx, appId, tokenId).Execute()

Delete API key



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    tokenId := "0aa1b2c3-d4e5-46f7-8899-aabbccddeeff" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.DeleteApiKey(orgAuth, appId, tokenId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `DeleteApiKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## DeleteSegment

> GenericSuccessBoolResponse DeleteSegment(ctx, appId, segmentId).Execute()

Delete Segment



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    segmentId := "d6c5a3e1-9f17-44a1-9d10-7c0e4a2b1c8e" // string | The segment_id can be found in the URL of the segment when viewing it in the dashboard.

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.DeleteSegment(restAuth, appId, segmentId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## DeleteSubscription

> DeleteSubscription(ctx, appId, subscriptionId).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    subscriptionId := "7e4c5b9a-1f60-4d07-9b1a-2e8c8d2cae51" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    r, err := apiClient.DefaultApi.DeleteSubscription(restAuth, appId, subscriptionId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## DeleteTemplate

> GenericSuccessBoolResponse DeleteTemplate(ctx, templateId).AppId(appId).Execute()

Delete template



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    templateId := "e4d3c2b1-a09f-4f1e-8d7c-6b5a4f3e2d1c" // string | 
    appId := "00000000-0000-0000-0000-000000000000" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.DeleteTemplate(restAuth, templateId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `DeleteTemplate`: GenericSuccessBoolResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** |  | 

### Return type

[**GenericSuccessBoolResponse**](GenericSuccessBoolResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## DeleteUser

> DeleteUser(ctx, appId, aliasLabel, aliasId).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    aliasLabel := "external_id" // string | 
    aliasId := "YOUR_USER_EXTERNAL_ID" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    r, err := apiClient.DefaultApi.DeleteUser(restAuth, appId, aliasLabel, aliasId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## ExportEvents

> ExportEventsSuccessResponse ExportEvents(ctx, notificationId).AppId(appId).Execute()

Export CSV of Events



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    notificationId := "b3a0c8bd-3a4c-4b22-9a73-3f1a8c2d1b88" // string | The ID of the notification to export events from.
    appId := "00000000-0000-0000-0000-000000000000" // string | The ID of the app that the notification belongs to.

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.ExportEvents(restAuth, notificationId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## ExportSubscriptions

> ExportSubscriptionsSuccessResponse ExportSubscriptions(ctx, appId).ExportSubscriptionsRequestBody(exportSubscriptionsRequestBody).Execute()

Export CSV of Subscriptions



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | The app ID that you want to export devices from
    exportSubscriptionsRequestBody := *onesignal.NewExportSubscriptionsRequestBody() // ExportSubscriptionsRequestBody |  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.ExportSubscriptions(restAuth, appId).ExportSubscriptionsRequestBody(exportSubscriptionsRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetAliases

> UserIdentityBody GetAliases(ctx, appId, aliasLabel, aliasId).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    aliasLabel := "external_id" // string | 
    aliasId := "YOUR_USER_EXTERNAL_ID" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetAliases(restAuth, appId, aliasLabel, aliasId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetAliasesBySubscription

> UserIdentityBody GetAliasesBySubscription(ctx, appId, subscriptionId).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    subscriptionId := "7e4c5b9a-1f60-4d07-9b1a-2e8c8d2cae51" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetAliasesBySubscription(restAuth, appId, subscriptionId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAliasesBySubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetApp

> App GetApp(ctx, appId).Execute()

View an app



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | An app id

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.GetApp(orgAuth, appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetApps

> []App GetApps(ctx).Execute()

View apps



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.GetApps(orgAuth).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetNotification

> NotificationWithMeta GetNotification(ctx, notificationId).AppId(appId).Execute()

View notification



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    notificationId := "b3a0c8bd-3a4c-4b22-9a73-3f1a8c2d1b88" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetNotification(restAuth, notificationId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetNotificationHistory

> NotificationHistorySuccessResponse GetNotificationHistory(ctx, notificationId).GetNotificationHistoryRequestBody(getNotificationHistoryRequestBody).Execute()

Notification History



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    notificationId := "b3a0c8bd-3a4c-4b22-9a73-3f1a8c2d1b88" // string | The \"id\" of the message found in the Notification object
    getNotificationHistoryRequestBody := *onesignal.NewGetNotificationHistoryRequestBody() // GetNotificationHistoryRequestBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetNotificationHistory(restAuth, notificationId).GetNotificationHistoryRequestBody(getNotificationHistoryRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotificationHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetNotifications

> NotificationSlice GetNotifications(ctx).AppId(appId).Limit(limit).Offset(offset).Kind(kind).TimeOffset(timeOffset).Execute()

View notifications



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | The app ID that you want to view notifications from
    limit := int32(10) // int32 | How many notifications to return.  Max is 50.  Default is 50. (optional)
    offset := int32(0) // int32 | Page offset.  Default is 0.  Results are sorted by queued_at in descending order.  queued_at is a representation of the time that the notification was queued at. (optional)
    kind := int32(0) // int32 | Kind of notifications returned:   * unset - All notification types (default)   * `0` - Dashboard only   * `1` - API only   * `3` - Automated only  (optional)
    timeOffset := "2025-01-01T00:00:00.000Z" // string | Time-offset pagination cursor for sequential pulls of all messages.  Accepts either an ISO 8601 formatted timestamp (e.g. `2025-01-01T00:00:00.000Z`) or the opaque Base64 cursor token returned as `next_time_offset` in a prior response.  When set, results are sorted ascending by send_after and the standard `offset` parameter cannot be used.  Repeat the request with each `next_time_offset` until an empty notifications array is returned. (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetNotifications(restAuth).AppId(appId).Limit(limit).Offset(offset).Kind(kind).TimeOffset(timeOffset).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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
 **timeOffset** | **string** | Time-offset pagination cursor for sequential pulls of all messages.  Accepts either an ISO 8601 formatted timestamp (e.g. &#x60;2025-01-01T00:00:00.000Z&#x60;) or the opaque Base64 cursor token returned as &#x60;next_time_offset&#x60; in a prior response.  When set, results are sorted ascending by send_after and the standard &#x60;offset&#x60; parameter cannot be used.  Repeat the request with each &#x60;next_time_offset&#x60; until an empty notifications array is returned. | 

### Return type

[**NotificationSlice**](NotificationSlice.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetOutcomes

> OutcomesData GetOutcomes(ctx, appId).OutcomeNames(outcomeNames).OutcomeNames2(outcomeNames2).OutcomeTimeRange(outcomeTimeRange).OutcomePlatforms(outcomePlatforms).OutcomeAttribution(outcomeAttribution).Execute()

View Outcomes



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    outcomeNames := "os__session_duration.count,os__click.count" // string | Required Comma-separated list of names and the value (sum/count) for the returned outcome data. Note: Clicks only support count aggregation. For out-of-the-box OneSignal outcomes such as click and session duration, please use the \"os\" prefix with two underscores. For other outcomes, please use the name specified by the user. Example:os__session_duration.count,os__click.count,CustomOutcomeName.sum 
    outcomeNames2 := "os__session_duration.count" // string | Optional If outcome names contain any commas, then please specify only one value at a time. Example: outcome_names[]=os__click.count&outcome_names[]=Sales, Purchase.count where \"Sales, Purchase\" is the custom outcomes with a comma in the name.  (optional)
    outcomeTimeRange := "1d" // string | Optional Time range for the returned data. The values can be 1h (for the last 1 hour data), 1d (for the last 1 day data), or 1mo (for the last 1 month data). Default is 1h if the parameter is omitted.  (optional)
    outcomePlatforms := "0,1" // string | Optional Platform id. Refer device's platform ids for values. Example: outcome_platform=0 for iOS outcome_platform=7,8 for Safari and Firefox Default is data from all platforms if the parameter is omitted.  (optional)
    outcomeAttribution := "direct" // string | Optional Attribution type for the outcomes. The values can be direct or influenced or unattributed. Example: outcome_attribution=direct Default is total (returns direct+influenced+unattributed) if the parameter is omitted.  (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetOutcomes(restAuth, appId).OutcomeNames(outcomeNames).OutcomeNames2(outcomeNames2).OutcomeTimeRange(outcomeTimeRange).OutcomePlatforms(outcomePlatforms).OutcomeAttribution(outcomeAttribution).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetOutcomes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetSegments

> GetSegmentsSuccessResponse GetSegments(ctx, appId).Offset(offset).Limit(limit).Execute()

Get Segments



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    offset := int32(0) // int32 | Segments are listed in ascending order of created_at date. offset will omit that number of segments from the beginning of the list. Eg offset 5, will remove the 5 earliest created Segments. (optional)
    limit := int32(10) // int32 | The amount of Segments in the response. Maximum 300. (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetSegments(restAuth, appId).Offset(offset).Limit(limit).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSegments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## GetUser

> User GetUser(ctx, appId, aliasLabel, aliasId).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    aliasLabel := "external_id" // string | 
    aliasId := "YOUR_USER_EXTERNAL_ID" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.GetUser(restAuth, appId, aliasLabel, aliasId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## RotateApiKey

> CreateApiKeyResponse RotateApiKey(ctx, appId, tokenId).Execute()

Rotate API key



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    tokenId := "0aa1b2c3-d4e5-46f7-8899-aabbccddeeff" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.RotateApiKey(orgAuth, appId, tokenId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RotateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `RotateApiKey`: CreateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RotateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## StartLiveActivity

> StartLiveActivitySuccessResponse StartLiveActivity(ctx, appId, activityType).StartLiveActivityRequest(startLiveActivityRequest).Execute()

Start Live Activity



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | Your OneSignal App ID in UUID v4 format.
    activityType := "order_status" // string | The name of the Live Activity defined in your app. This should match the attributes struct used in your app's Live Activity implementation.
    startLiveActivityRequest := *onesignal.NewStartLiveActivityRequest("Name_example", "Event_example", "ActivityId_example", map[string]interface{}(123), map[string]interface{}(123), *onesignal.NewLanguageStringMap(), *onesignal.NewLanguageStringMap()) // StartLiveActivityRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.StartLiveActivity(restAuth, appId, activityType).StartLiveActivityRequest(startLiveActivityRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StartLiveActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `StartLiveActivity`: StartLiveActivitySuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StartLiveActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Your OneSignal App ID in UUID v4 format. | 
**activityType** | **string** | The name of the Live Activity defined in your app. This should match the attributes struct used in your app&#39;s Live Activity implementation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartLiveActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startLiveActivityRequest** | [**StartLiveActivityRequest**](StartLiveActivityRequest.md) |  | 

### Return type

[**StartLiveActivitySuccessResponse**](StartLiveActivitySuccessResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## TransferSubscription

> UserIdentityBody TransferSubscription(ctx, appId, subscriptionId).TransferSubscriptionRequestBody(transferSubscriptionRequestBody).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    subscriptionId := "7e4c5b9a-1f60-4d07-9b1a-2e8c8d2cae51" // string | 
    transferSubscriptionRequestBody := *onesignal.NewTransferSubscriptionRequestBody() // TransferSubscriptionRequestBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.TransferSubscription(restAuth, appId, subscriptionId).TransferSubscriptionRequestBody(transferSubscriptionRequestBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TransferSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## UnsubscribeEmailWithToken

> GenericSuccessBoolResponse UnsubscribeEmailWithToken(ctx, appId, notificationId).Token(token).Execute()

Unsubscribe with token



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    notificationId := "b3a0c8bd-3a4c-4b22-9a73-3f1a8c2d1b88" // string | The id of the message found in the creation notification POST response, View Notifications GET response, or URL within the Message Report.
    token := "YOUR_TOKEN_ID" // string | The unsubscribe token that is generated via liquid syntax in {{subscription.unsubscribe_token}} when personalizing an email.

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UnsubscribeEmailWithToken(restAuth, appId, notificationId).Token(token).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UnsubscribeEmailWithToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## UpdateApiKey

> map[string]interface{} UpdateApiKey(ctx, appId, tokenId).UpdateApiKeyRequest(updateApiKeyRequest).Execute()

Update API key



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    tokenId := "0aa1b2c3-d4e5-46f7-8899-aabbccddeeff" // string | 
    updateApiKeyRequest := *onesignal.NewUpdateApiKeyRequest() // UpdateApiKeyRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.UpdateApiKey(orgAuth, appId, tokenId).UpdateApiKeyRequest(updateApiKeyRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `UpdateApiKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateApiKeyRequest** | [**UpdateApiKeyRequest**](UpdateApiKeyRequest.md) |  | 

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## UpdateApp

> App UpdateApp(ctx, appId).App(app).Execute()

Update an app



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | An app id
    app := *onesignal.NewApp() // App | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.UpdateApp(orgAuth, appId).App(app).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## UpdateLiveActivity

> UpdateLiveActivitySuccessResponse UpdateLiveActivity(ctx, appId, activityId).UpdateLiveActivityRequest(updateLiveActivityRequest).Execute()

Update a Live Activity via Push



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | The OneSignal App ID for your app.  Available in Keys & IDs.
    activityId := "12345" // string | Live Activity record ID
    updateLiveActivityRequest := *onesignal.NewUpdateLiveActivityRequest("Name_example", "Event_example", map[string]interface{}(123)) // UpdateLiveActivityRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UpdateLiveActivity(restAuth, appId, activityId).UpdateLiveActivityRequest(updateLiveActivityRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateLiveActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## UpdateSubscription

> UpdateSubscription(ctx, appId, subscriptionId).SubscriptionBody(subscriptionBody).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    subscriptionId := "7e4c5b9a-1f60-4d07-9b1a-2e8c8d2cae51" // string | 
    subscriptionBody := *onesignal.NewSubscriptionBody() // SubscriptionBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    r, err := apiClient.DefaultApi.UpdateSubscription(restAuth, appId, subscriptionId).SubscriptionBody(subscriptionBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## UpdateSubscriptionByToken

> map[string]interface{} UpdateSubscriptionByToken(ctx, appId, tokenType, token).SubscriptionBody(subscriptionBody).Execute()

Update subscription by token



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | Your OneSignal App ID in UUID v4 format.
    tokenType := "Email" // string | The type of token to use when looking up the subscription. See Subscription Types.
    token := "user@example.com" // string | The value of the token to lookup by (e.g., email address, phone number).
    subscriptionBody := *onesignal.NewSubscriptionBody() // SubscriptionBody | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UpdateSubscriptionByToken(restAuth, appId, tokenType, token).SubscriptionBody(subscriptionBody).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSubscriptionByToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `UpdateSubscriptionByToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSubscriptionByToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | Your OneSignal App ID in UUID v4 format. | 
**tokenType** | **string** | The type of token to use when looking up the subscription. See Subscription Types. | 
**token** | **string** | The value of the token to lookup by (e.g., email address, phone number). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionByTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **subscriptionBody** | [**SubscriptionBody**](SubscriptionBody.md) |  | 

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## UpdateTemplate

> TemplateResource UpdateTemplate(ctx, templateId).AppId(appId).UpdateTemplateRequest(updateTemplateRequest).Execute()

Update template



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    templateId := "e4d3c2b1-a09f-4f1e-8d7c-6b5a4f3e2d1c" // string | 
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    updateTemplateRequest := *onesignal.NewUpdateTemplateRequest() // UpdateTemplateRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UpdateTemplate(restAuth, templateId).AppId(appId).UpdateTemplateRequest(updateTemplateRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `UpdateTemplate`: TemplateResource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** |  | 
 **updateTemplateRequest** | [**UpdateTemplateRequest**](UpdateTemplateRequest.md) |  | 

### Return type

[**TemplateResource**](TemplateResource.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## UpdateUser

> PropertiesBody UpdateUser(ctx, appId, aliasLabel, aliasId).UpdateUserRequest(updateUserRequest).Execute()





### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 
    aliasLabel := "external_id" // string | 
    aliasId := "YOUR_USER_EXTERNAL_ID" // string | 
    updateUserRequest := *onesignal.NewUpdateUserRequest() // UpdateUserRequest | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.UpdateUser(restAuth, appId, aliasLabel, aliasId).UpdateUserRequest(updateUserRequest).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
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

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## ViewApiKeys

> ApiKeyTokensListResponse ViewApiKeys(ctx, appId).Execute()

View API keys



### Authorization

[organization_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    orgAuth := context.WithValue(context.Background(), onesignal.OrganizationApiKey, "YOUR_ORGANIZATION_API_KEY") // Organization API key is only required for creating new apps and other top-level endpoints

    resp, r, err := apiClient.DefaultApi.ViewApiKeys(orgAuth, appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ViewApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `ViewApiKeys`: ApiKeyTokensListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ViewApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKeyTokensListResponse**](ApiKeyTokensListResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## ViewTemplate

> TemplateResource ViewTemplate(ctx, templateId).AppId(appId).Execute()

View template



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    templateId := "e4d3c2b1-a09f-4f1e-8d7c-6b5a4f3e2d1c" // string | 
    appId := "00000000-0000-0000-0000-000000000000" // string | 

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.ViewTemplate(restAuth, templateId).AppId(appId).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ViewTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `ViewTemplate`: TemplateResource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ViewTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appId** | **string** |  | 

### Return type

[**TemplateResource**](TemplateResource.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


## ViewTemplates

> TemplatesListResponse ViewTemplates(ctx).AppId(appId).Limit(limit).Offset(offset).Channel(channel).Execute()

View templates



### Authorization

[rest_api_key](https://github.com/OneSignal/onesignal-go-api#configuration)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/OneSignal/onesignal-go-api/v5"
)

func main() {
    appId := "00000000-0000-0000-0000-000000000000" // string | Your OneSignal App ID in UUID v4 format.
    limit := int32(10) // int32 | Maximum number of templates. Default and max is 50. (optional) (default to 50)
    offset := int32(0) // int32 | Pagination offset. (optional) (default to 0)
    channel := "push" // string | Filter by delivery channel. (optional)

    configuration := onesignal.NewConfiguration()
    apiClient := onesignal.NewAPIClient(configuration)

    restAuth := context.WithValue(context.Background(), onesignal.RestApiKey, "YOUR_REST_API_KEY") // App REST API key required for most endpoints

    resp, r, err := apiClient.DefaultApi.ViewTemplates(restAuth).AppId(appId).Limit(limit).Offset(offset).Channel(channel).Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ViewTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        if apiErr, ok := err.(*onesignal.GenericOpenAPIError); ok {
            // ErrorMessages() flattens any error-envelope shape to a []string;
            // the raw body remains on Body().
            fmt.Fprintf(os.Stderr, "Error Messages: %v\n", apiErr.ErrorMessages())
            fmt.Fprintf(os.Stderr, "Response Body: %s\n", apiErr.Body())
        }
    }
    // response from `ViewTemplates`: TemplatesListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ViewTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiViewTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string** | Your OneSignal App ID in UUID v4 format. | 
 **limit** | **int32** | Maximum number of templates. Default and max is 50. | [default to 50]
 **offset** | **int32** | Pagination offset. | [default to 0]
 **channel** | **string** | Filter by delivery channel. | 

### Return type

[**TemplatesListResponse**](TemplatesListResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](https://github.com/OneSignal/onesignal-go-api#full-api-reference)
[[Back to README]](https://github.com/OneSignal/onesignal-go-api)


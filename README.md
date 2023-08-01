<h1 align="center">Welcome to the official OneSignal Go Client 👋</h1>

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.2.2
- Package version: 2.0.3
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://onesignal.com](https://onesignal.com)

## Installation

```shell
go get github.com/OneSignal/onesignal-go-api
```

Install the following dependencies:

```shell
go get golang.org/x/oauth2
```

Put the package under your project folder and add the following in import:

```golang
import "github.com/OneSignal/onesignal-go-api"
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
ctx := context.WithValue(context.Background(), onesignal.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), onesignal.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), onesignal.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), onesignal.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://onesignal.com/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**BeginLiveActivity**](docs/DefaultApi.md#beginliveactivity) | **Post** /apps/{app_id}/live_activities/{activity_id}/token | Start Live Activity
*DefaultApi* | [**CancelNotification**](docs/DefaultApi.md#cancelnotification) | **Delete** /notifications/{notification_id} | Stop a scheduled or currently outgoing notification
*DefaultApi* | [**CreateApp**](docs/DefaultApi.md#createapp) | **Post** /apps | Create an app
*DefaultApi* | [**CreateNotification**](docs/DefaultApi.md#createnotification) | **Post** /notifications | Create notification
*DefaultApi* | [**CreatePlayer**](docs/DefaultApi.md#createplayer) | **Post** /players | Add a device
*DefaultApi* | [**CreateSegments**](docs/DefaultApi.md#createsegments) | **Post** /apps/{app_id}/segments | Create Segments
*DefaultApi* | [**CreateSubscription**](docs/DefaultApi.md#createsubscription) | **Post** /apps/{app_id}/users/by/{alias_label}/{alias_id}/subscriptions | 
*DefaultApi* | [**CreateUser**](docs/DefaultApi.md#createuser) | **Post** /apps/{app_id}/users | 
*DefaultApi* | [**DeleteAlias**](docs/DefaultApi.md#deletealias) | **Delete** /apps/{app_id}/users/by/{alias_label}/{alias_id}/identity/{alias_label_to_delete} | 
*DefaultApi* | [**DeletePlayer**](docs/DefaultApi.md#deleteplayer) | **Delete** /players/{player_id} | Delete a user record
*DefaultApi* | [**DeleteSegments**](docs/DefaultApi.md#deletesegments) | **Delete** /apps/{app_id}/segments/{segment_id} | Delete Segments
*DefaultApi* | [**DeleteSubscription**](docs/DefaultApi.md#deletesubscription) | **Delete** /apps/{app_id}/subscriptions/{subscription_id} | 
*DefaultApi* | [**DeleteUser**](docs/DefaultApi.md#deleteuser) | **Delete** /apps/{app_id}/users/by/{alias_label}/{alias_id} | 
*DefaultApi* | [**EndLiveActivity**](docs/DefaultApi.md#endliveactivity) | **Delete** /apps/{app_id}/live_activities/{activity_id}/token/{subscription_id} | Stop Live Activity
*DefaultApi* | [**ExportEvents**](docs/DefaultApi.md#exportevents) | **Post** /notifications/{notification_id}/export_events?app_id&#x3D;{app_id} | Export CSV of Events
*DefaultApi* | [**ExportPlayers**](docs/DefaultApi.md#exportplayers) | **Post** /players/csv_export?app_id&#x3D;{app_id} | Export CSV of Players
*DefaultApi* | [**FetchAliases**](docs/DefaultApi.md#fetchaliases) | **Get** /apps/{app_id}/subscriptions/{subscription_id}/user/identity | 
*DefaultApi* | [**FetchUser**](docs/DefaultApi.md#fetchuser) | **Get** /apps/{app_id}/users/by/{alias_label}/{alias_id} | 
*DefaultApi* | [**FetchUserIdentity**](docs/DefaultApi.md#fetchuseridentity) | **Get** /apps/{app_id}/users/by/{alias_label}/{alias_id}/identity | 
*DefaultApi* | [**GetApp**](docs/DefaultApi.md#getapp) | **Get** /apps/{app_id} | View an app
*DefaultApi* | [**GetApps**](docs/DefaultApi.md#getapps) | **Get** /apps | View apps
*DefaultApi* | [**GetEligibleIams**](docs/DefaultApi.md#geteligibleiams) | **Get** /apps/{app_id}/subscriptions/{subscription_id}/iams | 
*DefaultApi* | [**GetNotification**](docs/DefaultApi.md#getnotification) | **Get** /notifications/{notification_id} | View notification
*DefaultApi* | [**GetNotificationHistory**](docs/DefaultApi.md#getnotificationhistory) | **Post** /notifications/{notification_id}/history | Notification History
*DefaultApi* | [**GetNotifications**](docs/DefaultApi.md#getnotifications) | **Get** /notifications | View notifications
*DefaultApi* | [**GetOutcomes**](docs/DefaultApi.md#getoutcomes) | **Get** /apps/{app_id}/outcomes | View Outcomes
*DefaultApi* | [**GetPlayer**](docs/DefaultApi.md#getplayer) | **Get** /players/{player_id} | View device
*DefaultApi* | [**GetPlayers**](docs/DefaultApi.md#getplayers) | **Get** /players | View devices
*DefaultApi* | [**IdentifyUserByAlias**](docs/DefaultApi.md#identifyuserbyalias) | **Patch** /apps/{app_id}/users/by/{alias_label}/{alias_id}/identity | 
*DefaultApi* | [**IdentifyUserBySubscriptionId**](docs/DefaultApi.md#identifyuserbysubscriptionid) | **Patch** /apps/{app_id}/subscriptions/{subscription_id}/user/identity | 
*DefaultApi* | [**TransferSubscription**](docs/DefaultApi.md#transfersubscription) | **Patch** /apps/{app_id}/subscriptions/{subscription_id}/owner | 
*DefaultApi* | [**UpdateApp**](docs/DefaultApi.md#updateapp) | **Put** /apps/{app_id} | Update an app
*DefaultApi* | [**UpdateLiveActivity**](docs/DefaultApi.md#updateliveactivity) | **Post** /apps/{app_id}/live_activities/{activity_id}/notifications | Update a Live Activity via Push
*DefaultApi* | [**UpdatePlayer**](docs/DefaultApi.md#updateplayer) | **Put** /players/{player_id} | Edit device
*DefaultApi* | [**UpdatePlayerTags**](docs/DefaultApi.md#updateplayertags) | **Put** /apps/{app_id}/users/{external_user_id} | Edit tags with external user id
*DefaultApi* | [**UpdateSubscription**](docs/DefaultApi.md#updatesubscription) | **Patch** /apps/{app_id}/subscriptions/{subscription_id} | 
*DefaultApi* | [**UpdateUser**](docs/DefaultApi.md#updateuser) | **Patch** /apps/{app_id}/users/by/{alias_label}/{alias_id} | 


## Documentation For Models

 - [App](docs/App.md)
 - [BasicNotification](docs/BasicNotification.md)
 - [BasicNotificationAllOf](docs/BasicNotificationAllOf.md)
 - [BasicNotificationAllOfAndroidBackgroundLayout](docs/BasicNotificationAllOfAndroidBackgroundLayout.md)
 - [BeginLiveActivityRequest](docs/BeginLiveActivityRequest.md)
 - [Button](docs/Button.md)
 - [CancelNotificationSuccessResponse](docs/CancelNotificationSuccessResponse.md)
 - [CreateNotificationSuccessResponse](docs/CreateNotificationSuccessResponse.md)
 - [CreatePlayerSuccessResponse](docs/CreatePlayerSuccessResponse.md)
 - [CreateSegmentConflictResponse](docs/CreateSegmentConflictResponse.md)
 - [CreateSegmentSuccessResponse](docs/CreateSegmentSuccessResponse.md)
 - [CreateSubscriptionRequestBody](docs/CreateSubscriptionRequestBody.md)
 - [CreateUserConflictResponse](docs/CreateUserConflictResponse.md)
 - [CreateUserConflictResponseErrorsInner](docs/CreateUserConflictResponseErrorsInner.md)
 - [CreateUserConflictResponseErrorsItemsMeta](docs/CreateUserConflictResponseErrorsItemsMeta.md)
 - [DeletePlayerNotFoundResponse](docs/DeletePlayerNotFoundResponse.md)
 - [DeletePlayerSuccessResponse](docs/DeletePlayerSuccessResponse.md)
 - [DeleteSegmentNotFoundResponse](docs/DeleteSegmentNotFoundResponse.md)
 - [DeleteSegmentSuccessResponse](docs/DeleteSegmentSuccessResponse.md)
 - [DeliveryData](docs/DeliveryData.md)
 - [ExportEventsSuccessResponse](docs/ExportEventsSuccessResponse.md)
 - [ExportPlayersRequestBody](docs/ExportPlayersRequestBody.md)
 - [ExportPlayersSuccessResponse](docs/ExportPlayersSuccessResponse.md)
 - [Filter](docs/Filter.md)
 - [FilterExpressions](docs/FilterExpressions.md)
 - [GenericError](docs/GenericError.md)
 - [GenericErrorErrorsInner](docs/GenericErrorErrorsInner.md)
 - [GetNotificationRequestBody](docs/GetNotificationRequestBody.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse201](docs/InlineResponse201.md)
 - [InlineResponse202](docs/InlineResponse202.md)
 - [InvalidIdentifierError](docs/InvalidIdentifierError.md)
 - [Notification](docs/Notification.md)
 - [Notification200Errors](docs/Notification200Errors.md)
 - [NotificationAllOf](docs/NotificationAllOf.md)
 - [NotificationHistorySuccessResponse](docs/NotificationHistorySuccessResponse.md)
 - [NotificationSlice](docs/NotificationSlice.md)
 - [NotificationTarget](docs/NotificationTarget.md)
 - [NotificationWithMeta](docs/NotificationWithMeta.md)
 - [NotificationWithMetaAllOf](docs/NotificationWithMetaAllOf.md)
 - [Operator](docs/Operator.md)
 - [OutcomeData](docs/OutcomeData.md)
 - [OutcomesData](docs/OutcomesData.md)
 - [PlatformDeliveryData](docs/PlatformDeliveryData.md)
 - [PlatformDeliveryDataEmailAllOf](docs/PlatformDeliveryDataEmailAllOf.md)
 - [PlatformDeliveryDataSmsAllOf](docs/PlatformDeliveryDataSmsAllOf.md)
 - [Player](docs/Player.md)
 - [PlayerNotificationTarget](docs/PlayerNotificationTarget.md)
 - [PlayerNotificationTargetIncludeAliases](docs/PlayerNotificationTargetIncludeAliases.md)
 - [PlayerSlice](docs/PlayerSlice.md)
 - [PropertiesDeltas](docs/PropertiesDeltas.md)
 - [PropertiesObject](docs/PropertiesObject.md)
 - [Purchase](docs/Purchase.md)
 - [RateLimiterError](docs/RateLimiterError.md)
 - [Segment](docs/Segment.md)
 - [SegmentNotificationTarget](docs/SegmentNotificationTarget.md)
 - [StringMap](docs/StringMap.md)
 - [SubscriptionObject](docs/SubscriptionObject.md)
 - [TransferSubscriptionRequestBody](docs/TransferSubscriptionRequestBody.md)
 - [UpdateLiveActivityRequest](docs/UpdateLiveActivityRequest.md)
 - [UpdateLiveActivitySuccessResponse](docs/UpdateLiveActivitySuccessResponse.md)
 - [UpdatePlayerSuccessResponse](docs/UpdatePlayerSuccessResponse.md)
 - [UpdatePlayerTagsRequestBody](docs/UpdatePlayerTagsRequestBody.md)
 - [UpdatePlayerTagsSuccessResponse](docs/UpdatePlayerTagsSuccessResponse.md)
 - [UpdateSubscriptionRequestBody](docs/UpdateSubscriptionRequestBody.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [User](docs/User.md)
 - [UserIdentityRequestBody](docs/UserIdentityRequestBody.md)
 - [UserIdentityResponse](docs/UserIdentityResponse.md)
 - [UserSubscriptionOptions](docs/UserSubscriptionOptions.md)


## Documentation For Authorization
Use a OneSignal authentication context for each auth type:
- `AppAuth`
- `UserAuth`

### app_key
- **Type**: HTTP Bearer token authentication

Example

```golang
appAuth := context.WithValue(context.Background(), onesignal.AppAuth, "APP_KEY_STRING")
```

### user_key
- **Type**: HTTP Bearer token authentication

Example

```golang
userAuth := context.WithValue(context.Background(), onesignal.UserAuth, "USER_KEY_STRING")
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

devrel@onesignal.com


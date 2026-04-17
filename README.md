<h1 align="center">Welcome to the official OneSignal Go Client</h1>

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

- API version: 5.4.0
- Package version: 5.4.0

## Installation

```shell
go get github.com/OneSignal/onesignal-go-api/v5
```

## Configuration

Every SDK requires authentication via API keys. Two key types are available:

- **REST API Key** — required for most endpoints (sending notifications, managing users, etc.). Found in your app's **Settings > Keys & IDs**.
- **Organization API Key** — only required for organization-level endpoints like creating or listing apps. Found in **Organization Settings**.

> **Warning:** Store your API keys in environment variables or a secrets manager. Never commit them to source control.

```go
import onesignal "github.com/OneSignal/onesignal-go-api"

restAuth := context.WithValue(
    context.Background(),
    onesignal.RestApiKey,
    "YOUR_REST_API_KEY",
)

orgAuth := context.WithValue(
    restAuth,
    onesignal.OrganizationApiKey,
    "YOUR_ORGANIZATION_API_KEY",
)

apiClient := onesignal.NewAPIClient(onesignal.NewConfiguration())
```

## Send a push notification

```go
notification := *onesignal.NewNotification("YOUR_APP_ID")
notification.SetContents(onesignal.StringMap{En: onesignal.PtrString("Hello from OneSignal!")})
notification.SetHeadings(onesignal.StringMap{En: onesignal.PtrString("Push Notification")})
notification.SetIncludedSegments([]string{"Subscribed Users"})

response, _, err := apiClient.DefaultApi
    .CreateNotification(orgAuth)
    .Notification(notification)
    .Execute()

if err != nil {
    log.Fatal(err)
}
fmt.Println("Notification ID:", response.GetId())
```

## Send an email

```go
notification := *onesignal.NewNotification("YOUR_APP_ID")
notification.SetEmailSubject("Important Update")
notification.SetEmailBody("<h1>Hello!</h1><p>This is an HTML email.</p>")
notification.SetIncludedSegments([]string{"Subscribed Users"})
notification.SetChannelForExternalUserIds("email")

response, _, err := apiClient.DefaultApi
    .CreateNotification(orgAuth)
    .Notification(notification)
    .Execute()
```

## Send an SMS

```go
notification := *onesignal.NewNotification("YOUR_APP_ID")
notification.SetContents(onesignal.StringMap{En: onesignal.PtrString("Your SMS message content here")})
notification.SetIncludedSegments([]string{"Subscribed Users"})
notification.SetChannelForExternalUserIds("sms")
notification.SetSmsFrom("+15551234567")

response, _, err := apiClient.DefaultApi
    .CreateNotification(orgAuth)
    .Notification(notification)
    .Execute()
```

## Full API reference

The complete list of API endpoints and their parameters is available in the [DefaultApi documentation](https://github.com/OneSignal/onesignal-go-api/blob/main/docs/DefaultApi.md).

For the underlying REST API, see the [OneSignal API reference](https://documentation.onesignal.com/reference).

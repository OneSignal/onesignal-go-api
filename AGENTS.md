# AGENTS.md — OneSignal Go SDK

Integration guide for AI agents and LLMs using the OneSignal server SDK for Go. Human-oriented docs are in the [README](./README.md).

## What this SDK does

Server-side access to the OneSignal REST API: send push / email / SMS, manage users, subscriptions, segments, templates and live activities, and administer apps & API keys.

- Module: `github.com/OneSignal/onesignal-go-api/v5`
- Repository: https://github.com/OneSignal/onesignal-go-api

## Install

```sh
go get github.com/OneSignal/onesignal-go-api/v5
```

```go
import onesignal "github.com/OneSignal/onesignal-go-api/v5"
```

## Authentication — two key types

OneSignal uses two bearer credentials; each endpoint requires a specific one. In Go you attach the key to the request `context.Context`, then pass that context to the call:

- **REST API key** (`onesignal.RestApiKey`) — used by almost every endpoint (notifications, users, subscriptions, segments, templates, live activities, custom events). Found in **App Settings → Keys & IDs**.
- **Organization API key** (`onesignal.OrganizationApiKey`) — required *only* for organization-level endpoints: app management (list / create / get / update apps), API-key management (view / create / update / rotate / delete API keys), and copying a template to another app. Found in **Organization Settings**.

Never hard-code keys — read them from environment variables or a secrets manager.

```go
apiClient := onesignal.NewAPIClient(onesignal.NewConfiguration())

// Most endpoints: attach the REST API key.
restCtx := context.WithValue(context.Background(), onesignal.RestApiKey, os.Getenv("ONESIGNAL_REST_API_KEY"))

// Organization-level endpoints: attach the Organization API key instead.
orgCtx := context.WithValue(context.Background(), onesignal.OrganizationApiKey, os.Getenv("ONESIGNAL_ORGANIZATION_API_KEY"))
```

## Calling convention

Calls use the fluent request builder: start with the method (passing the auth context), chain the body, then `.Execute()`.

```go
notification := onesignal.NewNotification("YOUR_APP_ID")
contents := onesignal.NewLanguageStringMap()
contents.SetEn("Hello from OneSignal!")
notification.SetContents(*contents)
notification.SetIncludeAliases(map[string][]string{"external_id": {"YOUR_USER_EXTERNAL_ID"}})
notification.SetTargetChannel("push")

resp, r, err := apiClient.DefaultApi.CreateNotification(restCtx).Notification(*notification).Execute()
```

## Idempotent sends & retries

Set `idempotency_key` (a UUID) so a create-notification request can be safely retried — the server returns the original result instead of sending twice. The `CreateNotificationWithRetry` helper handles this for you: it generates an `idempotency_key` when absent, retries `429` / `503` / transport errors with the **same** key (honoring `Retry-After`), and reports via `WasReplayed` whether the server answered from a previously completed request.

```go
opts := &onesignal.CreateNotificationWithRetryOptions{MaxRetries: 5}
result, err := apiClient.DefaultApi.CreateNotificationWithRetry(restCtx, *notification, opts)
if err == nil {
    fmt.Printf("id: %s replayed: %t\n", result.Response.GetId(), result.WasReplayed)
}
```

> The notification-level `external_id` field is the **deprecated** idempotency mechanism — prefer `idempotency_key`. Don't confuse it with the `external_id` **alias label** (under `include_aliases`) used to target users.

## Full API reference

- [DefaultApi.md](https://github.com/OneSignal/onesignal-go-api/blob/main/docs/DefaultApi.md) — every endpoint, parameter, and model, with runnable examples.
- [OneSignal REST API reference](https://documentation.onesignal.com/reference)

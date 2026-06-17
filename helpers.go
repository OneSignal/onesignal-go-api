package onesignal

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// CreateNotificationWithRetryOptions configures CreateNotificationWithRetry.
// Zero or negative values fall back to the defaults.
type CreateNotificationWithRetryOptions struct {
	// MaxRetries is the number of retries after the initial attempt. Default 3.
	MaxRetries int
	// BaseDelay is the exponential-backoff base used when the response carries
	// no Retry-After header. Clamped to [1s, 60s]. Default 1s.
	BaseDelay time.Duration
}

const (
	minBaseDelay = time.Second
	maxBaseDelay = 60 * time.Second
)

// CreateNotificationWithRetryResult carries the create response plus whether
// the server replayed a previously completed request (Idempotent-Replayed
// response header).
type CreateNotificationWithRetryResult struct {
	Response    *CreateNotificationSuccessResponse
	WasReplayed bool
}

// CreateNotificationWithRetry creates a notification with safe, idempotent
// retries. It ensures notification.IdempotencyKey is set (generating a UUIDv4
// when absent) so the server can deduplicate, then calls CreateNotification.
// Transient failures (HTTP 429, HTTP 503, or transport-level errors) are
// retried with the SAME idempotency key, honoring the Retry-After response
// header when present and falling back to exponential backoff
// (BaseDelay * 2^attempt) otherwise. Other errors are returned immediately.
func CreateNotificationWithRetry(ctx context.Context, client *APIClient, notification Notification, opts *CreateNotificationWithRetryOptions) (*CreateNotificationWithRetryResult, error) {
	maxRetries := 3
	baseDelay := time.Second
	if opts != nil {
		if opts.MaxRetries > 0 {
			maxRetries = opts.MaxRetries
		}
		if opts.BaseDelay > 0 {
			baseDelay = opts.BaseDelay
		}
	}
	// Clamp the backoff base so a stray value can neither hammer the API
	// (too small) nor stall the caller for an unbounded stretch (too large).
	if baseDelay < minBaseDelay {
		baseDelay = minBaseDelay
	} else if baseDelay > maxBaseDelay {
		baseDelay = maxBaseDelay
	}

	if notification.IdempotencyKey.Get() == nil || *notification.IdempotencyKey.Get() == "" {
		key, err := newUUIDv4()
		if err != nil {
			return nil, err
		}
		notification.SetIdempotencyKey(key)
	}

	for attempt := 0; ; attempt++ {
		resp, httpResp, err := client.DefaultApi.CreateNotification(ctx).Notification(notification).Execute()
		if err == nil {
			return &CreateNotificationWithRetryResult{
				Response:    resp,
				WasReplayed: isReplayed(httpResp),
			}, nil
		}

		// A nil *http.Response means the request never completed (timeout,
		// connection failure) — transient by definition.
		retryable := httpResp == nil || httpResp.StatusCode == http.StatusTooManyRequests || httpResp.StatusCode == http.StatusServiceUnavailable
		if !retryable || attempt >= maxRetries {
			return nil, err
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryDelay(httpResp, attempt, baseDelay)):
		}
	}
}

// CreateNotificationWithRetry surfaces the idempotent-retry helper as a
// DefaultApiService method so the call mirrors CreateNotification. It
// delegates to the package-level CreateNotificationWithRetry (single source of
// truth), reusing this service's API client.
func (a *DefaultApiService) CreateNotificationWithRetry(ctx context.Context, notification Notification, opts *CreateNotificationWithRetryOptions) (*CreateNotificationWithRetryResult, error) {
	return CreateNotificationWithRetry(ctx, a.client, notification, opts)
}

func isReplayed(httpResp *http.Response) bool {
	if httpResp == nil {
		return false
	}
	return strings.EqualFold(strings.TrimSpace(httpResp.Header.Get("Idempotent-Replayed")), "true")
}

func retryDelay(httpResp *http.Response, attempt int, baseDelay time.Duration) time.Duration {
	if httpResp != nil {
		if retryAfter := strings.TrimSpace(httpResp.Header.Get("Retry-After")); retryAfter != "" {
			if seconds, err := strconv.Atoi(retryAfter); err == nil && seconds >= 0 {
				return time.Duration(seconds) * time.Second
			}
		}
	}
	return baseDelay * (1 << uint(attempt))
}

func newUUIDv4() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}

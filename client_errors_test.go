package onesignal

import (
	"reflect"
	"testing"
)

func TestErrorMessagesNormalizesEnvelopeShapes(t *testing.T) {
	newErr := func(body string) GenericOpenAPIError {
		return GenericOpenAPIError{body: []byte(body)}
	}

	cases := []struct {
		name string
		body string
		want []string
	}{
		{"scalar", `{"errors": "Invalid app_id"}`, []string{"Invalid app_id"}},
		{"array of strings", `{"errors": ["All included players are not subscribed"]}`, []string{"All included players are not subscribed"}},
		{"array of objects", `{"errors": [{"code": "conflict", "title": "Conflicting aliases"}, {"code": "no_title"}]}`, []string{"Conflicting aliases", "no_title"}},
		{"no errors field", `{"success": false}`, []string{}},
		{"unparseable", "not json at all", []string{}},
		{"null errors", `{"errors": null}`, []string{}},
		{"array with null", `{"errors": [null]}`, []string{}},
		{"mixed array", `{"errors": ["str1", {"code": "x", "title": "y"}, "str2"]}`, []string{"str1", "y", "str2"}},
		{"non-string title", `{"errors": [{"code": "c1", "title": "good"}, {"code": "c2", "title": 42}, {"code": "c3", "title": "later"}]}`, []string{"good", "c2", "later"}},
		{"empty title falls back to code", `{"errors": [{"title": "", "code": "C1"}]}`, []string{"C1"}},
		{"object map surfaces keys", `{"errors": {"invalid_aliases": {"external_id": ["111"]}}}`, []string{`invalid_aliases: {"external_id":["111"]}`}},
		{"object map sorts multiple keys", `{"errors": {"invalid_player_ids": ["pid"], "invalid_aliases": {"external_id": ["111"]}}}`, []string{`invalid_aliases: {"external_id":["111"]}`, `invalid_player_ids: ["pid"]`}},
		{"object map string value", `{"errors": {"detail": "boom"}}`, []string{"detail: boom"}},
	}

	for _, tc := range cases {
		if got := newErr(tc.body).ErrorMessages(); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: ErrorMessages() = %#v, want %#v", tc.name, got, tc.want)
		}
	}
}

// The generated OneSignalErrors catalog exposes stable API sentinel strings and
// pairs with GenericOpenAPIError.ErrorMessages() for membership checks. Pure
// offline assertions (no network), so this runs without API credentials.
func TestOneSignalErrorsCatalog(t *testing.T) {
	if NO_SUBSCRIBERS != "All included players are not subscribed" {
		t.Errorf("NO_SUBSCRIBERS = %q", NO_SUBSCRIBERS)
	}
	if SERVICE_UNAVAILABLE != "Service temporarily unavailable" {
		t.Errorf("SERVICE_UNAVAILABLE = %q", SERVICE_UNAVAILABLE)
	}
	// Pin the auth string verbatim, including the intentional double space after
	// "denied." — a whitespace-normalizing edit would silently break matching.
	const wantInvalidKey = "Access denied.  Please include an 'Authorization: ...' header with a valid API key (https://documentation.onesignal.com/docs/en/keys-and-ids#api-keys)."
	if INVALID_API_KEY != wantInvalidKey {
		t.Errorf("INVALID_API_KEY = %q, want %q", INVALID_API_KEY, wantInvalidKey)
	}

	err := GenericOpenAPIError{body: []byte(`{"errors": ["All included players are not subscribed"]}`)}
	found := false
	for _, m := range err.ErrorMessages() {
		if m == NO_SUBSCRIBERS {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("ErrorMessages() should contain NO_SUBSCRIBERS")
	}
}

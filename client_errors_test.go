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

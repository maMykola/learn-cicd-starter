package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

// GetAPIKey

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		name  string
		input http.Header
		want  string
		err   error
	}{
		"no auth header": {input: http.Header{}, want: "", err: ErrNoAuthHeaderIncluded},
		"bearer":         {input: http.Header{"Authorization": []string{"Bearer token_value"}}, want: "", err: errors.New("malformed authorization header")},
		"api key":        {input: http.Header{"Authorization": []string{"ApiKey api_key"}}, want: "api_key", err: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if !reflect.DeepEqual(tc.want, got) || !reflect.DeepEqual(tc.err, err) {
				t.Fatalf("%s: expected: (%#v, %#v), got: (%#v, %#v)", name, tc.want, tc.err, got, err)
			}
		})
	}
}

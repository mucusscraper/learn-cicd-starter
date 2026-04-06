package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	var APIKey1 []string
	APIKey1 = append(APIKey1, "ApiKey APIKey1")
	HeaderAPIKey1 := http.Header{
		"Authorization": APIKey1,
	}
	var APIKey2 []string
	APIKey2 = append(APIKey2, "ApiKey APIKey1")
	HeaderAPIKey2 := http.Header{
		"Authorization2": APIKey2,
	}
	var APIKey3 []string
	APIKey3 = append(APIKey3, "ApiKeyAPIKey1")
	HeaderAPIKey3 := http.Header{
		"Authorization": APIKey3,
	}
	tests := []struct {
		name  string
		input http.Header
		want  string
	}{
		{name: "simple", input: HeaderAPIKey1, want: "APIKey1"},
		{name: "authorization mal formed", input: HeaderAPIKey2, want: ""},
		{name: "no space between apikey and key", input: HeaderAPIKey3, want: ""},
	}

	for _, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

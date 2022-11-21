package util

import (
	"testing"
)

func TestValidateDomain(t *testing.T) {

	tests := []struct {
		domain   string
		expected bool
	}{
		{"https://example.com", true},
		{"https://e.com", true},
		{"https://my.example.com", true},
		{"http://142.250.74.206", true},
		{"http://localhost", true},
	}

	for _, test := range tests {
		actual := ValidateDomain(test.domain)
		if actual != test.expected {
			t.Errorf("wrong result, got %t, want %t. Domain: %s", actual, test.expected, test.domain)
		}
	}
}

func TestValidateDomainScheme(t *testing.T) {

	tests := []struct {
		domain   string
		expected bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"ftp://127.0.0.1", false},
		{"example.com", false},
		{"://example.com", false},
		{"//example.com", false},
		{"/example.com", false},
	}

	for _, test := range tests {
		actual := ValidateDomain(test.domain)
		if actual != test.expected {
			t.Errorf("wrong result, got %t, want %t. Domain: %s", actual, test.expected, test.domain)
		}
	}
}

func TestValidateDomainHost(t *testing.T) {

	tests := []struct {
		domain   string
		expected bool
	}{
		{"http://example.com", true},
		{"http://example.org", true},
		{"http://example.net", true},
		{"http://example.de", true},
		{"http://example.academy", true},
	}

	for _, test := range tests {
		actual := ValidateDomain(test.domain)
		if actual != test.expected {
			t.Errorf("wrong result, got %t, want %t. Domain: %s", actual, test.expected, test.domain)
		}
	}
}

func TestGenerateRouteKey(t *testing.T) {
	expected := 4
	random := GenerateRouteKey(expected)
	actual := len(random)

	if len(random) != expected {
		t.Errorf("wrong result, got %d, want %d.", actual, expected)
	}
}

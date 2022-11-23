package util

import "testing"

// Tests domains without focus on specific URL components
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

// Tests domains with focus on schemes
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

// Tests domains with focus on TLDs
func TestValidateDomainTld(t *testing.T) {

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

// Tests the route key generator for the length
func TestRandomStringLength(t *testing.T) {
	expected := 4
	random := RandomString(expected)
	actual := len(random)

	if len(random) != expected {
		t.Errorf("wrong result, got %d, want %d.", actual, expected)
	}
}

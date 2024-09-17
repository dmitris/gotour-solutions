package stringers

import "testing"

func TestIpAddrString(t *testing.T) {
	test := []struct {
		IPAddr IPAddr
		want   string
	}{
		{IPAddr{1, 2, 3, 4}, "1.2.3.4"},
		{IPAddr{0, 0, 0, 0}, "0.0.0.0"},
		{IPAddr{10, 3, 111, 60}, "10.3.111.60"},
	}
	for _, tt := range test {
		if got := tt.IPAddr.String(); got != tt.want {
			t.Errorf("IPAddr.String() = %v, want %v", got, tt.want)
		}
	}
}

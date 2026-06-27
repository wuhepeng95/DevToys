package services

import (
	"testing"
)

func TestCalculateSubnet(t *testing.T) {
	tests := []struct {
		name    string
		cidr    string
		check   func(t *testing.T, info SubnetInfo)
		wantErr bool
	}{
		{
			name: "/24 subnet",
			cidr: "192.168.1.0/24",
			check: func(t *testing.T, info SubnetInfo) {
				if info.Network != "192.168.1.0" {
					t.Errorf("Network = %q, want 192.168.1.0", info.Network)
				}
				if info.Broadcast != "192.168.1.255" {
					t.Errorf("Broadcast = %q, want 192.168.1.255", info.Broadcast)
				}
				if info.SubnetMask != "255.255.255.0" {
					t.Errorf("SubnetMask = %q", info.SubnetMask)
				}
				if info.UsableStart != "192.168.1.1" {
					t.Errorf("UsableStart = %q", info.UsableStart)
				}
				if info.UsableEnd != "192.168.1.254" {
					t.Errorf("UsableEnd = %q", info.UsableEnd)
				}
				if info.UsableCount != 254 {
					t.Errorf("UsableCount = %d, want 254", info.UsableCount)
				}
				if info.NetSize != 24 {
					t.Errorf("NetSize = %d, want 24", info.NetSize)
				}
			},
		},
		{
			name: "/16 subnet",
			cidr: "10.0.0.0/16",
			check: func(t *testing.T, info SubnetInfo) {
				if info.HostCount != 65536 {
					t.Errorf("HostCount = %d, want 65536", info.HostCount)
				}
				if info.UsableCount != 65534 {
					t.Errorf("UsableCount = %d, want 65534", info.UsableCount)
				}
			},
		},
		{
			name: "/32 single host",
			cidr: "192.168.1.1/32",
			check: func(t *testing.T, info SubnetInfo) {
				if info.HostCount != 1 {
					t.Errorf("HostCount = %d, want 1", info.HostCount)
				}
				if info.UsableCount != 1 {
					t.Errorf("UsableCount = %d, want 1", info.UsableCount)
				}
			},
		},
		{
			name:    "invalid CIDR",
			cidr:    "not-a-cidr",
			wantErr: true,
		},
		{
			name:    "out of range mask",
			cidr:    "192.168.1.0/33",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateSubnet(tt.cidr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateSubnet() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.check != nil {
				tt.check(t, got)
			}
		})
	}
}

func TestSplitSubnet(t *testing.T) {
	result, err := SplitSubnet("192.168.1.0/24", 26)
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Subnets) != 4 {
		t.Fatalf("expected 4 subnets, got %d", len(result.Subnets))
	}

	expected := []string{
		"192.168.1.0/26",
		"192.168.1.64/26",
		"192.168.1.128/26",
		"192.168.1.192/26",
	}

	for i, sub := range result.Subnets {
		if sub.CIDR != expected[i] {
			t.Errorf("subnet %d = %q, want %q", i, sub.CIDR, expected[i])
		}
	}

	if result.Original.Network != "192.168.1.0" {
		t.Errorf("original network = %q", result.Original.Network)
	}
}

func TestSplitSubnetErrors(t *testing.T) {
	_, err := SplitSubnet("192.168.1.0/24", 16)
	if err == nil {
		t.Error("expected error for target mask <= original mask")
	}

	_, err = SplitSubnet("192.168.1.0/24", 33)
	if err == nil {
		t.Error("expected error for mask > 32")
	}
}

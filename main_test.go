package main

import "testing"

func TestHitungLuasKubus(t *testing.T) {
	testCases := []struct {
		sisi     float64
		expected float64
	}{
		{2, 24},
		{3, 54},
		{4, 96},
	}

	for _, tc := range testCases {
		luas := hitungLuasKubus(tc.sisi)
		if luas != tc.expected {
			t.Errorf("HitungLuasKubus(%.2f) = %.2f; expected %.2f", tc.sisi, luas, tc.expected)
		}
	}
}

package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"zero", 0, 5, 5},
		{"negative", -1, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		x, y int
		want int
	}{
		{"positive", 3, 4, 12},
		{"zero", 0, 5, 0},
		{"negative", -2, 3, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.x, tt.y); got != tt.want {
				t.Errorf("Multiply(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
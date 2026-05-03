package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterSubset(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		actual   string
		want     string
	}{
		{
			name:     "handles top level maps as subset",
			expected: `{"foo": "bar"}`,
			actual:   `{"foo": "bar", "extra": "val"}`,
			want:     `{"foo": "bar"}`,
		},
		{
			name:     "handles nested maps as subset",
			expected: `{"outer": {"inner": "val"}}`,
			actual:   `{"outer": {"inner": "val", "extra": "val"}, "other": "val"}`,
			want:     `{"outer": {"inner": "val"}}`,
		},
		{
			name:     "handles slices with different order and extra fields",
			expected: `{"list": [{"id": 1, "val": "a"}, {"id": 2, "val": "b"}]}`,
			actual:   `{"list": [{"id": 2, "val": "b", "extra": 1}, {"id": 1, "val": "a", "extra": 2}, {"id": 3, "val": "c"}], "other": 1}`,
			want:     `{"list": [{"id": 1, "val": "a"}, {"id": 2, "val": "b"}]}`,
		},
		{
			name:     "handles consent_management pattern",
			expected: `{"consent_management": [{"provider": "OneTrust", "id": "123"}]}`,
			actual:   `{"consent_management": [{"provider": "OneTrust", "id": "123", "resolution": "match"}], "other": "val"}`,
			want:     `{"consent_management": [{"provider": "OneTrust", "id": "123"}]}`,
		},
		{
			name:     "returns actual if type mismatch",
			expected: `{"foo": []}`,
			actual:   `{"foo": "not-a-slice"}`,
			want:     `{"foo": "not-a-slice"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var e, a, w interface{}
			require.NoError(t, json.Unmarshal([]byte(tc.expected), &e))
			require.NoError(t, json.Unmarshal([]byte(tc.actual), &a))
			require.NoError(t, json.Unmarshal([]byte(tc.want), &w))

			got := filterSubset(e, a)
			assert.Equal(t, w, got)
		})
	}
}

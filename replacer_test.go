package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "пше курва!",
			want:  "git push --force",
		},
		{
			input: "пше рассол сурово",
			want:  "git reset --hard",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.want, Replace(tt.input)) //nolint:scopelint
		})
	}
}

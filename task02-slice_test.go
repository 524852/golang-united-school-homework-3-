package homework

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		input      []int64
		wantResult []int64
	}{
		{
			name:       "from task",
			input:      []int64{1, 2, 5, 15},
			wantResult: []int64{15, 5, 2, 1},
		},
	}
	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if gotResult := reverse(tt.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("reverse() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

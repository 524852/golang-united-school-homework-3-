package homework

import (
	"reflect"
	"testing"
)

func Test_sortMapValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		input      map[int]string
		wantResult []string
	}{
		{
			name:       "from task-1",
			input:      map[int]string{2: "a", 0: "b", 1: "c"},
			wantResult: []string{"b", "c", "a"},
		},
		{
			name:       "from task-2",
			input:      map[int]string{10: "aa", 0: "bb", 500: "cc"},
			wantResult: []string{"bb", "aa", "cc"},
		},
	}
	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if gotResult := sortMapValues(tt.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("sortMapValues() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

package homework

import "testing"

func Test_average(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		input      [15]float32
		wantResult float32
	}{
		{
			name:       "from task",
			input:      [15]float32{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			wantResult: 1.4,
		},
		{
			name:       "from task",
			input:      [15]float32{0, 0, 0, 0, 0, 0, 0, 15, 0, 0, 0, 0, 0, 0, 0},
			wantResult: 1,
		},
	}
	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if gotResult := average(tt.input); gotResult != tt.wantResult {
				t.Errorf("average() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

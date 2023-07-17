package main

import "testing"

func Test_verifyPasswordHasAllRequiredChar(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"digits-only", args{"123456"}, false},
		{"lower-case-only", args{"abcdefg"}, false},
		{"upper-case-only", args{"ABCDEFG"}, false},
		{"symbols-only", args{"~!@#$%^&*()_+-={}[]:<>?,./"}, false},
		{"missing-digits", args{"Abcdefg@"}, false},
		{"missing-lower-case", args{"ABCDEFG@1"}, false},
		{"missing-upper-case", args{"abcdefg@1"}, false},
		{"missing-symbols", args{"Abcdefg1"}, false},
		{"all-required", args{"Abc123!@#"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPasswordHasAllRequiredChar(tt.args.password); got != tt.want {
				t.Errorf("verifyPasswordHasAllRequiredChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

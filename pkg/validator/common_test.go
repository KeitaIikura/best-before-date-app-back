package validator

import (
	"testing"
)

func TestPasswordValidate(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "OK_minimum",
			password: "Aa000000",
			wantErr:  false,
		},
		{
			name:     "OK_full",
			password: "Zz9!\"#$%&'()*+,\\-./:;<=>?@[]^_{|}~`",
			wantErr:  false,
		},
		{
			name:     "NG_length",
			password: "Aa0",
			wantErr:  true,
		},
		{
			name:     "NG_lack_of_uppercase",
			password: "a1234567",
			wantErr:  true,
		},
		{
			name:     "NG_lack_of_lowercase",
			password: "A1234567",
			wantErr:  true,
		},
		{
			name:     "NG_lack_of_number",
			password: "aABCDEFG",
			wantErr:  true,
		},
		{
			name:     "NG_invalid_character",
			password: "あAa0000",
			wantErr:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := PasswordValidate(test.password)
			if (err != nil) != test.wantErr {
				t.Errorf("failed test %s: error want is %v but result %v", test.name, test.wantErr, (err != nil))
			}
		})
	}
}

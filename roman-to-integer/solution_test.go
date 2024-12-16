package roman_to_integer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"K", 0},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := romanToInt(tt.in)
			if got != tt.out {
				t.Errorf("expected=%d, got=%d", tt.out, got)
			}
		})
	}
}

func mustNewToken(b byte) token {
	t, err := newToken(b)
	if err != nil {
		panic(err)
	}

	return t
}

func Test_token_update(t *testing.T) {
	tests := []struct {
		token       token
		updates     string
		wantUpdated bool
		wantTotal   int
		wantErr     bool
		wantClosed  bool
	}{
		{
			token:       mustNewToken('I'),
			updates:     "",
			wantTotal:   1,
			wantUpdated: false,
			wantClosed:  false,
		},
		{
			token:       mustNewToken('I'),
			updates:     "I",
			wantTotal:   2,
			wantUpdated: true,
			wantClosed:  false,
		},
		{
			token:       mustNewToken('I'),
			updates:     "II",
			wantTotal:   3,
			wantUpdated: true,
			wantClosed:  true,
		},
		{
			token:       mustNewToken('I'),
			updates:     "V",
			wantTotal:   4,
			wantUpdated: true,
			wantClosed:  true,
		},
		{
			token:      mustNewToken('V'),
			updates:    "",
			wantTotal:  5,
			wantClosed: false,
		},
		{
			token:       mustNewToken('V'),
			updates:     "I",
			wantTotal:   6,
			wantUpdated: true,
			wantClosed:  false,
		},
		{
			token:       mustNewToken('V'),
			updates:     "II",
			wantTotal:   7,
			wantUpdated: true,
			wantClosed:  false,
		},
		{
			token:       mustNewToken('V'),
			updates:     "III",
			wantTotal:   8,
			wantUpdated: true,
			wantClosed:  true,
		},
		{
			token:       mustNewToken('I'),
			updates:     "X",
			wantTotal:   9,
			wantUpdated: true,
			wantClosed:  true,
		},
		{
			token:       mustNewToken('I'),
			updates:     "L",
			wantTotal:   1,
			wantUpdated: false,
			wantClosed:  true,
		},
		{
			token:      mustNewToken('M'),
			updates:    "",
			wantTotal:  1,
			wantClosed: false,
		},
		{
			token:       mustNewToken('M'),
			updates:     "M",
			wantTotal:   2,
			wantUpdated: true,
			wantClosed:  false,
		},
		{
			token:       mustNewToken('M'),
			updates:     "MM",
			wantTotal:   3,
			wantUpdated: true,
			wantClosed:  true,
		},
		{
			token:       mustNewToken('M'),
			updates:     "C",
			wantTotal:   1,
			wantUpdated: false,
			wantClosed:  true,
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("token %c updated with %v", tt.token.tokenClass.oneSymbol, tt.updates)

		t.Run(name, func(t *testing.T) {
			var (
				gotErr     error = nil
				gotUpdated       = false
			)

			for _, b := range []byte(tt.updates) {
				gotUpdated, gotErr = tt.token.update(b)
				if gotErr != nil {
					break
				}
			}

			if tt.wantErr {
				assert.Error(t, gotErr)
				return
			}

			if !assert.NoError(t, gotErr) {
				return
			}

			assert.Equal(t, tt.wantUpdated, gotUpdated)
			assert.Equal(t, tt.wantTotal, tt.token.total)
			assert.Equal(t, tt.wantClosed, tt.token.closed)
		})
	}
}

func Test_romanToInt1(t *testing.T) {
	tests := []struct {
		in      string
		want    int
		wantErr bool
	}{
		{in: "III", want: 3},
		{in: "IV", want: 4},
		{in: "IX", want: 9},
		{in: "LVIII", want: 58},
		{in: "MCMXCIV", want: 1994},
		{in: "MMMCMXCIV", want: 3994},
		{in: "DCXXI", want: 621},
		{in: "IIILV", wantErr: true},
		{in: "K", wantErr: true},
		{in: "MK", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, gotErr := romanToInt1(tt.in)

			if tt.wantErr {
				assert.Error(t, gotErr)
				return
			}

			if !assert.NoError(t, gotErr) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

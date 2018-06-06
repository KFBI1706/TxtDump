package crypto

import "testing"

func Test_sha256encode(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"testing with single byte", args{b: []byte{0x6e}}, "1b16b1df538ba12dc3f97edbb85caa7050d46c148134290feba80f8236c83db9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sha256encode(tt.args.b); got != tt.want {
				t.Errorf("sha256encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package semver

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    Semver
		wantErr bool
	}{
		{
			name:    "success",
			version: "1.0",
			want:    Semver{Major: 1, Minor: 0},
			wantErr: false,
		},
		{
			name:    "sem_ver_3",
			version: "1.1.3",
			want:    Semver{Major: 1, Minor: 1},
			wantErr: false,
		},
		{
			name:    "incorrect",
			version: "incorrect",
			want:    Semver{},
			wantErr: true,
		},
		{
			name:    "incorrect_major",
			version: "incorrect.0",
			want:    Semver{},
			wantErr: true,
		},
		{
			name:    "incorrect_minor",
			version: "0.incorrect",
			want:    Semver{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := Parse(tt.version)
			if err != nil != tt.wantErr {
				t.Fatalf("wantErr %v got %v", tt.wantErr, err)
			}

			if !reflect.DeepEqual(s, tt.want) {
				t.Fatalf("expected %v got %v", tt.want, s)
			}
		})
	}
}

func Test_MajorGreatThen(t *testing.T) {
	tests := []struct {
		name string
		a    Semver
		b    Semver
		want bool
	}{
		{
			name: "false",
			a:    Semver{Major: 1, Minor: 99},
			b:    Semver{Major: 2, Minor: 0},
			want: false,
		},
		{
			name: "true",
			a:    Semver{Major: 3, Minor: 0},
			b:    Semver{Major: 2, Minor: 99},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.a.MajorGreatThen(tt.b)
			if res != tt.want {
				t.Fatalf("Result want %v got %v", tt.want, res)
			}
		})
	}
}

func Test_MinorGreatThen(t *testing.T) {
	tests := []struct {
		name string
		a    Semver
		b    Semver
		want bool
	}{
		{
			name: "false",
			a:    Semver{Major: 2, Minor: 0},
			b:    Semver{Major: 1, Minor: 99},
			want: false,
		},
		{
			name: "true",
			a:    Semver{Major: 2, Minor: 99},
			b:    Semver{Major: 3, Minor: 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.a.MinorGreatThen(tt.b)
			if res != tt.want {
				t.Fatalf("Result want %v got %v", tt.want, res)
			}
		})
	}
}

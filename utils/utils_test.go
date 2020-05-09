package utils

import (
	"os"
	"runtime"
	"testing"
)

func TestFilePath(t *testing.T) {
	type args struct {
		name   string
		ext    string
		escape bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal test",
			args: args{
				name:   "hello",
				ext:    "txt",
				escape: false,
			},
			want: "hello.txt",
		},
		{
			name: "normal test",
			args: args{
				name:   "hello:world",
				ext:    "txt",
				escape: true,
			},
			want: "hello：world.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := FilePath(tt.args.name, tt.args.ext, tt.args.escape); got != tt.want {
				t.Errorf("FilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMkdir(t *testing.T) {
	tests := []struct {
		name     string
		path     []string
		wantWin  string
		wantUnix string
	}{
		{
			name:     "a",
			path:     []string{"a"},
			wantWin:  "a",
			wantUnix: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Mkdir(tt.path...)

			if runtime.GOOS == "windows" {
				if got != tt.wantWin {
					t.Errorf("Mkdir() = %v, want %v", got, tt.wantWin)
				}
			} else {
				if got != tt.wantUnix {
					t.Errorf("Mkdir() = %v, want %v", got, tt.wantUnix)
				}
			}
			os.RemoveAll(got)
		})
	}
}

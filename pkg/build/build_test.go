package build

import (
	"testing"

	vcs2 "github.com/acorn-io/runtime/pkg/vcs"
	"github.com/stretchr/testify/assert"
)

func TestVCS(t *testing.T) {
	vcs := vcs2.VCS(".")
	assert.NotEqual(t, "", vcs.Revision)
}

func Test_toContextCopyDockerFile(t *testing.T) {
	type args struct {
		baseImage   string
		contextDirs map[string]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				baseImage: "base",
				contextDirs: map[string]string{
					"/var/tmp":   "./files",
					"/var/tmp2/": "./files2",
				},
			},
			want: `FROM base
COPY --link "./files" "/var/tmp"
COPY --link "./files2" "/var/tmp2/"
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toContextCopyDockerFile(tt.args.baseImage, tt.args.contextDirs)
			assert.Equal(t, tt.want, got)
		})
	}
}

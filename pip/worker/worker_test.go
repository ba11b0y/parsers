package worker

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsRootModule(t *testing.T) {
	moduleName := "test-module"
	moduleVersion := "0.92.0"
	for name, tc := range map[string]struct {
		pkg      Packages
		pipType  string
		expected bool
	}{
		"darwin, poetry": {
			pkg: Packages{
				Name:      moduleName,
				Version:   moduleVersion,
				Installer: "poetry",
				Root:      false,
			},
			pipType:  "poetry",
			expected: true,
		},
		"darwin, pyenv": {
			pkg: Packages{
				Name:     moduleName,
				Version:  moduleVersion,
				Location: "/Users/test-user/Documents/test-module",
				Root:     false,
			},
			pipType:  "pyenv",
			expected: true,
		},
		"darwin, pipenv": {
			pkg: Packages{
				Name:     moduleName,
				Version:  moduleVersion,
				Location: "/Users/test-user/Documents/test-module",
				Root:     false,
			},
			pipType:  "pipenv",
			expected: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			actual := IsRootModule(tc.pkg, tc.pipType)
			require.Equal(t, tc.expected, actual)
		})
	}
}

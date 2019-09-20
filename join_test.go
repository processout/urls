package urls_test

import (
	"testing"

	"github.com/kaluza-tech/urls"
	"gotest.tools/assert"
)

func TestJoin(t *testing.T) {
	t.Parallel()
	table := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "Same as path.Join",
			args:     []string{"hello/world", "addition"},
			expected: "hello/world/addition",
		},
		{
			name:     "Same as path.Join with trailing slash",
			args:     []string{"hello/world/", "addition"},
			expected: "hello/world/addition",
		},
		{
			name:     "Multiple additions to path",
			args:     []string{"hello/world", "addition/", "/another/", "/again"},
			expected: "hello/world/addition/another/again",
		},
		{
			name:     "URL protocol",
			args:     []string{"https://hello/world/", "addition"},
			expected: "https://hello/world/addition",
		},
		{
			name:     "0 args",
			args:     []string{},
			expected: "",
		},
		{
			name:     "Query Params are dealt with elegantly",
			args:     []string{"https://hello/world?query=param&hello=world", "addition"},
			expected: "https://hello/world/addition?query=param&hello=world",
		},
		{
			name:     "Query Params without URL protocol are dealt with elegantly",
			args:     []string{"hello/world?query=param&hello=world", "addition"},
			expected: "hello/world/addition?query=param&hello=world",
		},
		{
			name:     "Continuations are dealt with elegantly",
			args:     []string{"https://hello/world#idOfSomeThing", "addition"},
			expected: "https://hello/world/addition#idOfSomeThing",
		},
	}

	for _, test := range table {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result, err := urls.Join(tc.args...)
			assert.NilError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestJoinError(t *testing.T) {
	_, err := urls.Join(string(0x7f))
	assert.ErrorContains(t, err, "url.Parse")
}

package urls_test

import (
	"net/url"
	"testing"

	"github.com/kaluza-tech/urls"
	"gotest.tools/assert"
)

func TestJoinURL(t *testing.T) {
	t.Parallel()
	table := []struct {
		name      string
		url       string
		additions []string
		expected  string
	}{
		{
			name:      "Regular URL",
			url:       "http://hello.world.com/",
			additions: []string{"addition", "again"},
			expected:  "http://hello.world.com/addition/again",
		},
		{
			name:      "0 args in additions",
			url:       "http://hello.world.com/",
			additions: []string{},
			expected:  "http://hello.world.com/",
		},
		{
			name:      "Query Params are dealt with elegantly",
			url:       "https://hello/world?query=param&hello=world",
			additions: []string{"addition"},
			expected:  "https://hello/world/addition?query=param&hello=world",
		},
		{
			name:      "Continuations are dealt with elegantly",
			url:       "https://hello/world#stuffAfterHash",
			additions: []string{"addition"},
			expected:  "https://hello/world/addition#stuffAfterHash",
		},
	}

	for _, test := range table {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			myURL, err := url.Parse(tc.url)
			if err != nil {
				t.Fatal(err)
			}
			urls.JoinURL(myURL, tc.additions...)
			assert.Equal(t, tc.expected, myURL.String())
		})
	}
}

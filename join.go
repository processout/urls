package urls

import (
	"net/url"
	"path"

	"github.com/pkg/errors"
)

// Join is equivalent to path.Join, except that it works for urls
func Join(paths ...string) (string, error) {
	if len(paths) == 0 {
		return "", nil
	}
	url, err := url.Parse(paths[0])
	if err != nil {
		return "", errors.Wrap(err, "url.Parse")
	}
	// overwrite path[0] with url.Path as that won't have URL protocol
	paths[0] = url.Path
	url.Path = path.Join(paths...)
	return url.String(), nil
}

package urls

import (
	"net/url"
	"path"
)

// Join is equivalent to path.Join, except that it works for urls
func Join(paths ...string) (string, error) {
	if len(paths) == 0 {
		return "", nil
	}
	url, err := url.Parse(paths[0])
	if err != nil {
		return "", err
	}
	// overwrite path[0] with url.Path as that won't have URL protocol
	paths[0] = url.Path
	url.Path = path.Join(paths...)
	return url.String(), nil
}

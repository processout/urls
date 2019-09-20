package urls

import (
	"net/url"
	"path"
)

// JoinURL is similar to Join, but it takes in a *url.URL pointer and updates it
// in place.
func JoinURL(origURL *url.URL, paths ...string) {
	pathsToJoin := []string{origURL.Path}
	pathsToJoin = append(pathsToJoin, paths...)
	origURL.Path = path.Join(pathsToJoin...)
}

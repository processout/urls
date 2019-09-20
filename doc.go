/*
Package urls provides some very helpful functions for manipulating URLs.

For example, you may want to simply add a path onto your url.
If you were to use the path pacakge you would end up with an unexpected result:
	newURL := path.Join("http://original.url.com/stuff", "addition")
	// newURL is "http:/original.url.com/stuff/addition", which is unexpected

With this package you can end up with the expected result:
	newURL, err := urls.Join("http://original.url.com/stuff", "addition")
	// newURL is "http://original.url.com/stuff/addition", as expected

If you are alredy starting with a url object then you can get an updated url object:
	origURL, _ := url.Parse("http://original.url.com/stuff")
	newURL := urls.JoinURL(origURL, "addition")
	// newURL.String() is "http://original.url.com/stuff/addition", as expected

*/
package urls

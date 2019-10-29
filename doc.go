/*
Package urls provides some very helpful functions for manipulating URLs.

For example, you may want to simply add a path onto your url.
If you were to use the path pacakge you would end up with an unexpected result:
	newURL := path.Join("http://original.url.com/stuff", "addition")
	// newURL is "http:/original.url.com/stuff/addition", which is unexpected

With this package you can end up with the expected result:
	newURL, err := urls.Join("http://original.url.com/stuff", "addition")
	// newURL is "http://original.url.com/stuff/addition", as expected

If you are alredy starting with a *url.URL pointer then you can get an updated *url.URL pointer:
	myURL, _ := url.Parse("http://original.url.com/stuff")
	urls.JoinURL(myURL, "addition")
	// myURL.String() is "http://original.url.com/stuff/addition", as expected

An example involving query parameters:
	params := url.Values{}
	params.Set("query", "param")
	newURL, err := urls.AppendQueryParams("http://original.url.com/stuff?hello=world", params)
	// newURL is "http://original.url.com/stuff?hello=world&query=param"

If you are alredy starting with a \*url.URL pointer then you can get an updated \*url.URL pointer:
	myURL, err := url.Parse("http://original.url.com/stuff?hello=world")
	params := url.Values{}
	params.Set("query", "param")
	urls.AppendQueryParamsURL(myURL, params)
	// myURL.String() is now "http://original.url.com/stuff?hello=world&query=param"


*/
package urls

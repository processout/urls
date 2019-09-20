# urls
Join URLs and manipulate query parameters.

For example, you may want to simply add a path onto your url.
If you were to use the path pacakge you would end up with unexpected result:


	fmt.Println(path.Join("<a href="http://original.url.com/stuff">http://original.url.com/stuff</a>", "addition"))
	// Results in "http:/original.url.com/stuff/addition", which is unexpected

With this package you can end up with the expected result:


	fmt.Println(urls.Join("<a href="http://original.url.com/stuff">http://original.url.com/stuff</a>", "addition"))
	// Results in "<a href="http://original.url.com/stuff/addition">http://original.url.com/stuff/addition</a>", as expected

If you are alredy starting with a url object then you can get an updated url object:


	origURL, _ := url.Parse("<a href="http://original.url.com/stuff">http://original.url.com/stuff</a>")
	newURL := urls.JoinURL(origURL, "addition")
	fmt.Println(newURL.String())
	// Results in "<a href="http://original.url.com/stuff/addition">http://original.url.com/stuff/addition</a>", as expected

## Documentation
See the documentation on [godoc.org](https://godoc.org/github.com/kaluza-tech/urls).

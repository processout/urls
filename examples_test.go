package urls_test

import (
	"fmt"
	"net/url"

	"github.com/kaluza-tech/urls"
)

func ExampleJoin_joinPaths() {
	newURL, err := urls.Join("http://original.url.com/stuff", "addition", "extrastuff")
	if err != nil {
		fmt.Println("unexpected error: ", err)
	}
	fmt.Println(newURL)
	// Output:
	// http://original.url.com/stuff/addition/extrastuff
}

func ExampleJoin_withQueryParameters() {
	newURL, err := urls.Join("http://original.url.com/stuff?query=param", "addition", "extrastuff")
	if err != nil {
		fmt.Println("unexpected error: ", err)
	}
	fmt.Println(newURL)
	// Output:
	// http://original.url.com/stuff/addition/extrastuff?query=param
}

func ExampleJoinURL_joinPaths() {
	newURL, err := url.Parse("http://original.url.com/stuff")
	if err != nil {
		fmt.Println("unexpected error: ", err)
	}
	urls.JoinURL(newURL, "addition", "extrastuff")
	fmt.Println(newURL.String())
	// Output:
	// http://original.url.com/stuff/addition/extrastuff
}

func ExampleJoinURL_withQueryParameters() {
	newURL, err := url.Parse("http://original.url.com/stuff?query=param")
	if err != nil {
		fmt.Println("unexpected error: ", err)
	}
	urls.JoinURL(newURL, "addition", "extrastuff")
	fmt.Println(newURL.String())
	// Output:
	// http://original.url.com/stuff/addition/extrastuff?query=param
}

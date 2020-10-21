package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// We’ll parse this example URL, which includes a scheme, authentication info, host, port, path,
	// query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("scheme:", u.Scheme)

	// User contains all authentication info; call Username and Password on this for individual values.
	fmt.Println("user:", u.User)
	fmt.Println("user.name:", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("passwd:", p)

	// The Host contains both the hostname and the port, if present. Use SplitHostPort to extract them.
	fmt.Println("host/port:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host:", host)
	fmt.Println("port:", port)

	// Here we extract the path and the fragment after the #.
	fmt.Println("path:", u.Path)
	fmt.Println("fragment:", u.Fragment)

	// To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map.
	// The parsed query param maps are from strings to slices of strings,
	// so index into [0] if you only want the first value.
	fmt.Println("raw query params:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("parsed query params:", m)
	fmt.Println("m[\"k\"][0:", m["k"][0])
}

package main

import "testing"

func TestReplaceHost(t *testing.T) {
	url := "http://127.0.0.1:8008/4ryl4Dc79RCz4n8ENjYD_08w3u7KoF401qpD47HR17JkrMrOStni5NHxhAFELCwfes8g5Gyz-o37Pe3DJ3xIubAfUU8eqILImTmc"
	url2 := replaceHost(url, "https://lh3.googleusercontent.com")
	if url2 != "https://lh3.googleusercontent.com/4ryl4Dc79RCz4n8ENjYD_08w3u7KoF401qpD47HR17JkrMrOStni5NHxhAFELCwfes8g5Gyz-o37Pe3DJ3xIubAfUU8eqILImTmc" {
		t.Error("replace error ", url2)
	}
}

func TestHttpGetHelper(t *testing.T) {
	url := "http://127.0.0.1:8008/4ryl4Dc79RCz4n8ENjYD_08w3u7KoF401qpD47HR17JkrMrOStni5NHxhAFELCwfes8g5Gyz-o37Pe3DJ3xIubAfUU8eqILImTmc"
	if data, err := httpGetHelper(url); err != nil {
		t.Error(err)
	} else {
		if len(data) != 78297 {
			t.Error("httpGetHelper return error length")
		}
	}
}

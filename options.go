package xeno

import (
	"net/url"
	"strconv"
)

type RequestOption func(*requestOptions)

type requestOptions struct {
	urlParams url.Values
}

func processOptions(options ...RequestOption) requestOptions {
	o := requestOptions{
		urlParams: url.Values{},
	}
	for _, opt := range options {
		opt(&o)
	}
	return o
}

func Page(amount int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("page", strconv.Itoa(amount))
	}
}

func NumPages(amount int) RequestOption {
	return func(o *requestOptions) {
		o.urlParams.Set("numPages", strconv.Itoa(amount))
	}
}

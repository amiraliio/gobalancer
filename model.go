package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

//Backend model
type Backend struct {
	IsAlive      bool
	URL          *url.URL
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

//ServerPool model
type SeverPool struct {
	backends []*Backend
	current  uint64
}

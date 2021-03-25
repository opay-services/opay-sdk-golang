package util

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type options struct {
	ReadTimeout    time.Duration
	ConnectTimeout time.Duration
	TLSConfig      *tls.Config
	Timeout        time.Duration  //connect and read time
	Transport      *http.Transport
}

type HttpOption func(*options)

func HttpConnectTimeout(t time.Duration) HttpOption {
	return func(o *options) {
		o.ConnectTimeout = t
	}
}

func TLSConfig(t *tls.Config) HttpOption {
	return func(o *options) {
		o.TLSConfig = t
	}
}

func HttpReadTimeout(t time.Duration) HttpOption {
	return func(o *options) {
		o.ReadTimeout = t
	}
}

func NewHttpClient(opts ...HttpOption) http.Client {
	httpOptions := options{}
	for _, o := range (opts) {
		o(&httpOptions)
	}

	client := http.Client{}
	if httpOptions.Timeout != time.Duration(0){
		client.Timeout = httpOptions.Timeout
	}

	if len(opts) > 0 && httpOptions.Transport != nil{
		transport := http.Transport{DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			if httpOptions.ConnectTimeout != time.Duration(0) {
				conn, e = net.DialTimeout(network, addr, httpOptions.ConnectTimeout)
				if e != nil {
					return conn, e
				}
			}

			if httpOptions.ReadTimeout != time.Duration(0) {
				conn.SetDeadline(time.Now().Add(httpOptions.ReadTimeout))
			}
			return conn, e
		}}

		client.Transport = http.RoundTripper(&transport)
		return client
	}

	if httpOptions.Transport != nil{
		client.Transport = httpOptions.Transport
	}

	return client
}


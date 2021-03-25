package util

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type httpOptions struct {
	ReadTimeout    time.Duration
	ConnectTimeout time.Duration
	TLSConfig      *tls.Config
}

type HttpOption func(*httpOptions)

func HttpConnectTimeout(t time.Duration) HttpOption {
	return func(o *httpOptions) {
		o.ConnectTimeout = t
	}
}

func TLSConfig(t *tls.Config) HttpOption {
	return func(o *httpOptions) {
		o.TLSConfig = t
	}
}

func HttpReadTimeout(t time.Duration) HttpOption {
	return func(o *httpOptions) {
		o.ReadTimeout = t
	}
}

func NewHttpClient(opts ...HttpOption) http.Client {
	if len(opts) > 0 {
		transport := http.Transport{DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
			httpOptions := httpOptions{}
			for _, o := range (opts) {
				o(&httpOptions)
			}

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

		return http.Client{Transport: http.RoundTripper(&transport)}
	}

	return http.Client{}
}


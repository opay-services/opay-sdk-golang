package sdk

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"crypto/tls"
	"encoding/hex"
	"github.com/opay-services/opay-sdk-golang/conf"
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

func Signature(data []byte) (ret string) {
	h := hmac.New(sha512.New, []byte(conf.GetSecretKey()))
	h.Write([]byte(data))
	ret = hex.EncodeToString(h.Sum(nil))
	return
}
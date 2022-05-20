// Author: yann
// Date: 2022/5/20
// Desc: hdfs_rest

package hdfsrest

import (
	"net"
	"net/http"
	"time"
)

type client struct {
	*http.Client
	DataNodes []string
	User      string
}

type ClientOption struct {
	DataNodes []string
	User      string
}

func NewClient(opt *ClientOption) *client {
	return &client{
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   3 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
		DataNodes: opt.DataNodes,
		User:      opt.User,
	}
}

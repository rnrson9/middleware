// Copyright 2017 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package handlers

import "net/http"

type host struct {
	domains []string
	handler http.Handler
}

// Host 声明一个限定域名的 handler。
//
// 若请求的域名不允许，会返回 403 错误。
// 若 domains 为空，则任何请求都将返回 403。
//
// 仅会将域名与 domains 进行比较，端口与协议都将不参写比较。
func Host(h http.Handler, domains ...string) http.Handler {
	return &host{
		domains: domains,
		handler: h,
	}
}

// HostFunc 将一个 http.HandlerFunc 包装成 http.Handler
func HostFunc(f func(http.ResponseWriter, *http.Request), domains ...string) http.Handler {
	return Host(http.HandlerFunc(f), domains...)
}

func (h *host) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostname := r.URL.Hostname()
	for _, domain := range h.domains {
		if domain == hostname {
			h.handler.ServeHTTP(w, r)
			return
		}
	}

	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

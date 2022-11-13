// Copyright(c) 2022 individual contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package harwp

import (
	"io"
	"net/http"
)

type responseProxyBase struct {
	proxied http.ResponseWriter

	statusCode    int
	headerWritten bool
	bytesWritten  int
	hijacked      bool
}

// Ensure responseProxyBase implements the required interfaces.
var (
	_ io.Writer           = &responseProxyBase{}
	_ http.ResponseWriter = &responseProxyBase{}
)

func (p *responseProxyBase) StatusCode() int             { return p.statusCode }
func (p *responseProxyBase) HeadersWritten() bool        { return p.headerWritten }
func (p *responseProxyBase) BytesWritten() int           { return p.bytesWritten }
func (p *responseProxyBase) Hijacked() bool              { return p.hijacked }
func (p *responseProxyBase) Unwrap() http.ResponseWriter { return p.proxied }

func (p *responseProxyBase) maybeWriteHeader() {
	if !p.headerWritten {
		p.WriteHeader(http.StatusOK)
	}
}

func (p *responseProxyBase) Write(data []byte) (n int, err error) {
	p.maybeWriteHeader()
	n, err = p.proxied.Write(data)
	p.bytesWritten += n
	return
}

func (p *responseProxyBase) Header() http.Header {
	return p.proxied.Header()
}

func (p *responseProxyBase) WriteHeader(statusCode int) {
	if !p.headerWritten {
		p.headerWritten = true

		p.statusCode = statusCode
		p.proxied.WriteHeader(statusCode)
	}
}

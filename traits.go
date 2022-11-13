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
	"bufio"
	"errors"
	"io"
	"net"
	"net/http"
)

type responseProxyFl struct{ *responseProxyBase }

func (p *responseProxyFl) Flush() {
	p.maybeWriteHeader()
	p.proxied.(http.Flusher).Flush()
}

type responseProxyHj struct{ *responseProxyBase }

func (p *responseProxyHj) Hijack() (rwc net.Conn, buf *bufio.ReadWriter, err error) {
	rwc, buf, err = p.proxied.(http.Hijacker).Hijack()
	if err == nil || errors.Is(err, http.ErrHijacked) {
		p.hijacked = true
	}
	return
}

type responseProxyPh struct{ *responseProxyBase }

func (p *responseProxyPh) Push(target string, opts *http.PushOptions) error {
	return p.proxied.(http.Pusher).Push(target, opts)
}

type responseProxyRf struct{ *responseProxyBase }

func (p *responseProxyRf) ReadFrom(r io.Reader) (n int64, err error) {
	p.maybeWriteHeader()
	n, err = p.proxied.(io.ReaderFrom).ReadFrom(r)
	p.bytesWritten += int(n)
	return
}

type responseProxySw struct{ *responseProxyBase }

func (p *responseProxySw) WriteString(s string) (n int, err error) {
	p.maybeWriteHeader()
	n, err = p.proxied.(io.StringWriter).WriteString(s)
	p.bytesWritten += n
	return
}

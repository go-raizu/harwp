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

package harwp_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	assertPkg "github.com/stretchr/testify/assert"
	requirePkg "github.com/stretchr/testify/require"

	"github.com/go-raizu/harwp"
)

func Test_Trait_Hijacked(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := harwp.NewResponseProxy(w)

		hj, ok := p.(http.Hijacker)
		if !assertPkg.True(t, ok) {
			return
		}

		rwc, buf, err := hj.Hijack()
		defer func() { _ = rwc.Close() }()
		assertPkg.NoError(t, err)

		assertPkg.True(t, p.Hijacked())

		_, _ = buf.WriteString("HTTP/1.0 200 OK\r\n")
		_, _ = buf.WriteString("Content-Type: text/plain\r\n")
		_, _ = buf.WriteString("Connection: close\r\n\r\n")
		_, _ = buf.WriteString("Hello World")
		_ = buf.Flush()
	}))
	defer srv.Close()

	resp, err := srv.Client().Get(srv.URL)
	requirePkg.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	requirePkg.NoError(t, err)

	assertPkg.Equal(t, "Hello World", string(body))
}

func Test_Trait_WriteString(t *testing.T) {
	assert := assertPkg.New(t)

	body := assertPkg.HTTPBody(func(w http.ResponseWriter, r *http.Request) {
		p := harwp.NewResponseProxy(w)

		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			n, err := w.(io.StringWriter).WriteString("Hello World!")
			requirePkg.NoError(t, err)

			assert.Equal(12, n)
		}).ServeHTTP(p, r)
	}, "GET", "", nil)

	assert.Equal("Hello World!", body)
}

type readFromMock struct{ readFromCalled bool }

func (m *readFromMock) ReadFrom(r io.Reader) (n int64, err error) {
	m.readFromCalled = true

	d, err := io.ReadAll(r)
	return int64(len(d)), err
}
func (m *readFromMock) Header() (h http.Header)           { return }
func (m *readFromMock) Write(i []byte) (n int, err error) { return }
func (m *readFromMock) WriteHeader(statusCode int)        { return }

type readWrapper struct{ io.Reader }

func Test_Trait_ReadFrom(t *testing.T) {
	assert := assertPkg.New(t)

	m := &readFromMock{}

	p := harwp.NewResponseProxy(m)

	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := bytes.NewBufferString("Hello World!")

		n, err := io.Copy(w, readWrapper{buf})
		requirePkg.NoError(t, err)

		assert.Equal(int64(12), n)
	}).ServeHTTP(p, nil)

	assert.True(m.readFromCalled)
	assert.Equal(p.BytesWritten(), 12)
}

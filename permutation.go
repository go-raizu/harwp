// Code generated by 'go generate'; DO NOT EDIT.

package harwp

import (
	"io"
	"net/http"
)

type traits int

const (
	traitsBase       traits = 0
	traitsFl         traits = 1
	traitsHj         traits = 2
	traitsFlHj       traits = 3
	traitsPh         traits = 4
	traitsFlPh       traits = 5
	traitsHjPh       traits = 6
	traitsFlHjPh     traits = 7
	traitsRf         traits = 8
	traitsFlRf       traits = 9
	traitsHjRf       traits = 10
	traitsFlHjRf     traits = 11
	traitsPhRf       traits = 12
	traitsFlPhRf     traits = 13
	traitsHjPhRf     traits = 14
	traitsFlHjPhRf   traits = 15
	traitsSw         traits = 16
	traitsFlSw       traits = 17
	traitsHjSw       traits = 18
	traitsFlHjSw     traits = 19
	traitsPhSw       traits = 20
	traitsFlPhSw     traits = 21
	traitsHjPhSw     traits = 22
	traitsFlHjPhSw   traits = 23
	traitsRfSw       traits = 24
	traitsFlRfSw     traits = 25
	traitsHjRfSw     traits = 26
	traitsFlHjRfSw   traits = 27
	traitsPhRfSw     traits = 28
	traitsFlPhRfSw   traits = 29
	traitsHjPhRfSw   traits = 30
	traitsFlHjPhRfSw traits = 31
)

var (
	_ http.ResponseWriter = (*responseProxyFl)(nil)
	_ http.Flusher        = (*responseProxyFl)(nil)
)

func newResponseProxyFl(b *responseProxyBase) *responseProxyFl {
	return &responseProxyFl{
		responseProxyBase: b,
	}
}

var (
	_ http.ResponseWriter = (*responseProxyHj)(nil)
	_ http.Hijacker       = (*responseProxyHj)(nil)
)

func newResponseProxyHj(b *responseProxyBase) *responseProxyHj {
	return &responseProxyHj{
		responseProxyBase: b,
	}
}

var (
	_ http.ResponseWriter = (*responseProxyPh)(nil)
	_ http.Pusher         = (*responseProxyPh)(nil)
)

func newResponseProxyPh(b *responseProxyBase) *responseProxyPh {
	return &responseProxyPh{
		responseProxyBase: b,
	}
}

var (
	_ http.ResponseWriter = (*responseProxyRf)(nil)
	_ io.ReaderFrom       = (*responseProxyRf)(nil)
)

func newResponseProxyRf(b *responseProxyBase) *responseProxyRf {
	return &responseProxyRf{
		responseProxyBase: b,
	}
}

var (
	_ http.ResponseWriter = (*responseProxySw)(nil)
	_ io.StringWriter     = (*responseProxySw)(nil)
)

func newResponseProxySw(b *responseProxyBase) *responseProxySw {
	return &responseProxySw{
		responseProxyBase: b,
	}
}

type responseProxyFlHj struct {
	*responseProxyBase
	*responseProxyFl
	*responseProxyHj
}

var (
	_ http.ResponseWriter = (*responseProxyFlHj)(nil)
	_ http.Flusher        = (*responseProxyFlHj)(nil)
	_ http.Hijacker       = (*responseProxyFlHj)(nil)
)

func newResponseProxyFlHj(b *responseProxyBase) *responseProxyFlHj {
	return &responseProxyFlHj{
		responseProxyBase: b,
		responseProxyFl:   newResponseProxyFl(b),
		responseProxyHj:   newResponseProxyHj(b),
	}
}

type responseProxyFlPh struct {
	*responseProxyBase
	*responseProxyFl
	*responseProxyPh
}

var (
	_ http.ResponseWriter = (*responseProxyFlPh)(nil)
	_ http.Flusher        = (*responseProxyFlPh)(nil)
	_ http.Pusher         = (*responseProxyFlPh)(nil)
)

func newResponseProxyFlPh(b *responseProxyBase) *responseProxyFlPh {
	return &responseProxyFlPh{
		responseProxyBase: b,
		responseProxyFl:   newResponseProxyFl(b),
		responseProxyPh:   newResponseProxyPh(b),
	}
}

type responseProxyFlRf struct {
	*responseProxyBase
	*responseProxyFl
	*responseProxyRf
}

var (
	_ http.ResponseWriter = (*responseProxyFlRf)(nil)
	_ http.Flusher        = (*responseProxyFlRf)(nil)
	_ io.ReaderFrom       = (*responseProxyFlRf)(nil)
)

func newResponseProxyFlRf(b *responseProxyBase) *responseProxyFlRf {
	return &responseProxyFlRf{
		responseProxyBase: b,
		responseProxyFl:   newResponseProxyFl(b),
		responseProxyRf:   newResponseProxyRf(b),
	}
}

type responseProxyFlSw struct {
	*responseProxyBase
	*responseProxyFl
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyFlSw)(nil)
	_ http.Flusher        = (*responseProxyFlSw)(nil)
	_ io.StringWriter     = (*responseProxyFlSw)(nil)
)

func newResponseProxyFlSw(b *responseProxyBase) *responseProxyFlSw {
	return &responseProxyFlSw{
		responseProxyBase: b,
		responseProxyFl:   newResponseProxyFl(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyHjPh struct {
	*responseProxyBase
	*responseProxyHj
	*responseProxyPh
}

var (
	_ http.ResponseWriter = (*responseProxyHjPh)(nil)
	_ http.Hijacker       = (*responseProxyHjPh)(nil)
	_ http.Pusher         = (*responseProxyHjPh)(nil)
)

func newResponseProxyHjPh(b *responseProxyBase) *responseProxyHjPh {
	return &responseProxyHjPh{
		responseProxyBase: b,
		responseProxyHj:   newResponseProxyHj(b),
		responseProxyPh:   newResponseProxyPh(b),
	}
}

type responseProxyHjRf struct {
	*responseProxyBase
	*responseProxyHj
	*responseProxyRf
}

var (
	_ http.ResponseWriter = (*responseProxyHjRf)(nil)
	_ http.Hijacker       = (*responseProxyHjRf)(nil)
	_ io.ReaderFrom       = (*responseProxyHjRf)(nil)
)

func newResponseProxyHjRf(b *responseProxyBase) *responseProxyHjRf {
	return &responseProxyHjRf{
		responseProxyBase: b,
		responseProxyHj:   newResponseProxyHj(b),
		responseProxyRf:   newResponseProxyRf(b),
	}
}

type responseProxyHjSw struct {
	*responseProxyBase
	*responseProxyHj
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyHjSw)(nil)
	_ http.Hijacker       = (*responseProxyHjSw)(nil)
	_ io.StringWriter     = (*responseProxyHjSw)(nil)
)

func newResponseProxyHjSw(b *responseProxyBase) *responseProxyHjSw {
	return &responseProxyHjSw{
		responseProxyBase: b,
		responseProxyHj:   newResponseProxyHj(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyPhRf struct {
	*responseProxyBase
	*responseProxyPh
	*responseProxyRf
}

var (
	_ http.ResponseWriter = (*responseProxyPhRf)(nil)
	_ http.Pusher         = (*responseProxyPhRf)(nil)
	_ io.ReaderFrom       = (*responseProxyPhRf)(nil)
)

func newResponseProxyPhRf(b *responseProxyBase) *responseProxyPhRf {
	return &responseProxyPhRf{
		responseProxyBase: b,
		responseProxyPh:   newResponseProxyPh(b),
		responseProxyRf:   newResponseProxyRf(b),
	}
}

type responseProxyPhSw struct {
	*responseProxyBase
	*responseProxyPh
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyPhSw)(nil)
	_ http.Pusher         = (*responseProxyPhSw)(nil)
	_ io.StringWriter     = (*responseProxyPhSw)(nil)
)

func newResponseProxyPhSw(b *responseProxyBase) *responseProxyPhSw {
	return &responseProxyPhSw{
		responseProxyBase: b,
		responseProxyPh:   newResponseProxyPh(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyRfSw struct {
	*responseProxyBase
	*responseProxyRf
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyRfSw)(nil)
	_ io.ReaderFrom       = (*responseProxyRfSw)(nil)
	_ io.StringWriter     = (*responseProxyRfSw)(nil)
)

func newResponseProxyRfSw(b *responseProxyBase) *responseProxyRfSw {
	return &responseProxyRfSw{
		responseProxyBase: b,
		responseProxyRf:   newResponseProxyRf(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyFlHjPh struct {
	*responseProxyBase
	*responseProxyFlHj
	*responseProxyPh
}

var (
	_ http.ResponseWriter = (*responseProxyFlHjPh)(nil)
	_ http.Flusher        = (*responseProxyFlHjPh)(nil)
	_ http.Hijacker       = (*responseProxyFlHjPh)(nil)
	_ http.Pusher         = (*responseProxyFlHjPh)(nil)
)

func newResponseProxyFlHjPh(b *responseProxyBase) *responseProxyFlHjPh {
	return &responseProxyFlHjPh{
		responseProxyBase: b,
		responseProxyFlHj: newResponseProxyFlHj(b),
		responseProxyPh:   newResponseProxyPh(b),
	}
}

type responseProxyFlHjRf struct {
	*responseProxyBase
	*responseProxyFlHj
	*responseProxyRf
}

var (
	_ http.ResponseWriter = (*responseProxyFlHjRf)(nil)
	_ http.Flusher        = (*responseProxyFlHjRf)(nil)
	_ http.Hijacker       = (*responseProxyFlHjRf)(nil)
	_ io.ReaderFrom       = (*responseProxyFlHjRf)(nil)
)

func newResponseProxyFlHjRf(b *responseProxyBase) *responseProxyFlHjRf {
	return &responseProxyFlHjRf{
		responseProxyBase: b,
		responseProxyFlHj: newResponseProxyFlHj(b),
		responseProxyRf:   newResponseProxyRf(b),
	}
}

type responseProxyFlHjSw struct {
	*responseProxyBase
	*responseProxyFlHj
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyFlHjSw)(nil)
	_ http.Flusher        = (*responseProxyFlHjSw)(nil)
	_ http.Hijacker       = (*responseProxyFlHjSw)(nil)
	_ io.StringWriter     = (*responseProxyFlHjSw)(nil)
)

func newResponseProxyFlHjSw(b *responseProxyBase) *responseProxyFlHjSw {
	return &responseProxyFlHjSw{
		responseProxyBase: b,
		responseProxyFlHj: newResponseProxyFlHj(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyFlPhRf struct {
	*responseProxyBase
	*responseProxyFlPh
	*responseProxyRf
}

var (
	_ http.ResponseWriter = (*responseProxyFlPhRf)(nil)
	_ http.Flusher        = (*responseProxyFlPhRf)(nil)
	_ http.Pusher         = (*responseProxyFlPhRf)(nil)
	_ io.ReaderFrom       = (*responseProxyFlPhRf)(nil)
)

func newResponseProxyFlPhRf(b *responseProxyBase) *responseProxyFlPhRf {
	return &responseProxyFlPhRf{
		responseProxyBase: b,
		responseProxyFlPh: newResponseProxyFlPh(b),
		responseProxyRf:   newResponseProxyRf(b),
	}
}

type responseProxyFlPhSw struct {
	*responseProxyBase
	*responseProxyFlPh
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyFlPhSw)(nil)
	_ http.Flusher        = (*responseProxyFlPhSw)(nil)
	_ http.Pusher         = (*responseProxyFlPhSw)(nil)
	_ io.StringWriter     = (*responseProxyFlPhSw)(nil)
)

func newResponseProxyFlPhSw(b *responseProxyBase) *responseProxyFlPhSw {
	return &responseProxyFlPhSw{
		responseProxyBase: b,
		responseProxyFlPh: newResponseProxyFlPh(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyFlRfSw struct {
	*responseProxyBase
	*responseProxyFlRf
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyFlRfSw)(nil)
	_ http.Flusher        = (*responseProxyFlRfSw)(nil)
	_ io.ReaderFrom       = (*responseProxyFlRfSw)(nil)
	_ io.StringWriter     = (*responseProxyFlRfSw)(nil)
)

func newResponseProxyFlRfSw(b *responseProxyBase) *responseProxyFlRfSw {
	return &responseProxyFlRfSw{
		responseProxyBase: b,
		responseProxyFlRf: newResponseProxyFlRf(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyHjPhRf struct {
	*responseProxyBase
	*responseProxyHjPh
	*responseProxyRf
}

var (
	_ http.ResponseWriter = (*responseProxyHjPhRf)(nil)
	_ http.Hijacker       = (*responseProxyHjPhRf)(nil)
	_ http.Pusher         = (*responseProxyHjPhRf)(nil)
	_ io.ReaderFrom       = (*responseProxyHjPhRf)(nil)
)

func newResponseProxyHjPhRf(b *responseProxyBase) *responseProxyHjPhRf {
	return &responseProxyHjPhRf{
		responseProxyBase: b,
		responseProxyHjPh: newResponseProxyHjPh(b),
		responseProxyRf:   newResponseProxyRf(b),
	}
}

type responseProxyHjPhSw struct {
	*responseProxyBase
	*responseProxyHjPh
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyHjPhSw)(nil)
	_ http.Hijacker       = (*responseProxyHjPhSw)(nil)
	_ http.Pusher         = (*responseProxyHjPhSw)(nil)
	_ io.StringWriter     = (*responseProxyHjPhSw)(nil)
)

func newResponseProxyHjPhSw(b *responseProxyBase) *responseProxyHjPhSw {
	return &responseProxyHjPhSw{
		responseProxyBase: b,
		responseProxyHjPh: newResponseProxyHjPh(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyHjRfSw struct {
	*responseProxyBase
	*responseProxyHjRf
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyHjRfSw)(nil)
	_ http.Hijacker       = (*responseProxyHjRfSw)(nil)
	_ io.ReaderFrom       = (*responseProxyHjRfSw)(nil)
	_ io.StringWriter     = (*responseProxyHjRfSw)(nil)
)

func newResponseProxyHjRfSw(b *responseProxyBase) *responseProxyHjRfSw {
	return &responseProxyHjRfSw{
		responseProxyBase: b,
		responseProxyHjRf: newResponseProxyHjRf(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyPhRfSw struct {
	*responseProxyBase
	*responseProxyPhRf
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyPhRfSw)(nil)
	_ http.Pusher         = (*responseProxyPhRfSw)(nil)
	_ io.ReaderFrom       = (*responseProxyPhRfSw)(nil)
	_ io.StringWriter     = (*responseProxyPhRfSw)(nil)
)

func newResponseProxyPhRfSw(b *responseProxyBase) *responseProxyPhRfSw {
	return &responseProxyPhRfSw{
		responseProxyBase: b,
		responseProxyPhRf: newResponseProxyPhRf(b),
		responseProxySw:   newResponseProxySw(b),
	}
}

type responseProxyFlHjPhRf struct {
	*responseProxyBase
	*responseProxyFlHjPh
	*responseProxyRf
}

var (
	_ http.ResponseWriter = (*responseProxyFlHjPhRf)(nil)
	_ http.Flusher        = (*responseProxyFlHjPhRf)(nil)
	_ http.Hijacker       = (*responseProxyFlHjPhRf)(nil)
	_ http.Pusher         = (*responseProxyFlHjPhRf)(nil)
	_ io.ReaderFrom       = (*responseProxyFlHjPhRf)(nil)
)

func newResponseProxyFlHjPhRf(b *responseProxyBase) *responseProxyFlHjPhRf {
	return &responseProxyFlHjPhRf{
		responseProxyBase:   b,
		responseProxyFlHjPh: newResponseProxyFlHjPh(b),
		responseProxyRf:     newResponseProxyRf(b),
	}
}

type responseProxyFlHjPhSw struct {
	*responseProxyBase
	*responseProxyFlHjPh
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyFlHjPhSw)(nil)
	_ http.Flusher        = (*responseProxyFlHjPhSw)(nil)
	_ http.Hijacker       = (*responseProxyFlHjPhSw)(nil)
	_ http.Pusher         = (*responseProxyFlHjPhSw)(nil)
	_ io.StringWriter     = (*responseProxyFlHjPhSw)(nil)
)

func newResponseProxyFlHjPhSw(b *responseProxyBase) *responseProxyFlHjPhSw {
	return &responseProxyFlHjPhSw{
		responseProxyBase:   b,
		responseProxyFlHjPh: newResponseProxyFlHjPh(b),
		responseProxySw:     newResponseProxySw(b),
	}
}

type responseProxyFlHjRfSw struct {
	*responseProxyBase
	*responseProxyFlHjRf
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyFlHjRfSw)(nil)
	_ http.Flusher        = (*responseProxyFlHjRfSw)(nil)
	_ http.Hijacker       = (*responseProxyFlHjRfSw)(nil)
	_ io.ReaderFrom       = (*responseProxyFlHjRfSw)(nil)
	_ io.StringWriter     = (*responseProxyFlHjRfSw)(nil)
)

func newResponseProxyFlHjRfSw(b *responseProxyBase) *responseProxyFlHjRfSw {
	return &responseProxyFlHjRfSw{
		responseProxyBase:   b,
		responseProxyFlHjRf: newResponseProxyFlHjRf(b),
		responseProxySw:     newResponseProxySw(b),
	}
}

type responseProxyFlPhRfSw struct {
	*responseProxyBase
	*responseProxyFlPhRf
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyFlPhRfSw)(nil)
	_ http.Flusher        = (*responseProxyFlPhRfSw)(nil)
	_ http.Pusher         = (*responseProxyFlPhRfSw)(nil)
	_ io.ReaderFrom       = (*responseProxyFlPhRfSw)(nil)
	_ io.StringWriter     = (*responseProxyFlPhRfSw)(nil)
)

func newResponseProxyFlPhRfSw(b *responseProxyBase) *responseProxyFlPhRfSw {
	return &responseProxyFlPhRfSw{
		responseProxyBase:   b,
		responseProxyFlPhRf: newResponseProxyFlPhRf(b),
		responseProxySw:     newResponseProxySw(b),
	}
}

type responseProxyHjPhRfSw struct {
	*responseProxyBase
	*responseProxyHjPhRf
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyHjPhRfSw)(nil)
	_ http.Hijacker       = (*responseProxyHjPhRfSw)(nil)
	_ http.Pusher         = (*responseProxyHjPhRfSw)(nil)
	_ io.ReaderFrom       = (*responseProxyHjPhRfSw)(nil)
	_ io.StringWriter     = (*responseProxyHjPhRfSw)(nil)
)

func newResponseProxyHjPhRfSw(b *responseProxyBase) *responseProxyHjPhRfSw {
	return &responseProxyHjPhRfSw{
		responseProxyBase:   b,
		responseProxyHjPhRf: newResponseProxyHjPhRf(b),
		responseProxySw:     newResponseProxySw(b),
	}
}

type responseProxyFlHjPhRfSw struct {
	*responseProxyBase
	*responseProxyFlHjPhRf
	*responseProxySw
}

var (
	_ http.ResponseWriter = (*responseProxyFlHjPhRfSw)(nil)
	_ http.Flusher        = (*responseProxyFlHjPhRfSw)(nil)
	_ http.Hijacker       = (*responseProxyFlHjPhRfSw)(nil)
	_ http.Pusher         = (*responseProxyFlHjPhRfSw)(nil)
	_ io.ReaderFrom       = (*responseProxyFlHjPhRfSw)(nil)
	_ io.StringWriter     = (*responseProxyFlHjPhRfSw)(nil)
)

func newResponseProxyFlHjPhRfSw(b *responseProxyBase) *responseProxyFlHjPhRfSw {
	return &responseProxyFlHjPhRfSw{
		responseProxyBase:     b,
		responseProxyFlHjPhRf: newResponseProxyFlHjPhRf(b),
		responseProxySw:       newResponseProxySw(b),
	}
}

func detectTraits(w http.ResponseWriter) (out traits) {
	if _, has := w.(http.Flusher); has {
		out |= 1 << 0
	}
	if _, has := w.(http.Hijacker); has {
		out |= 1 << 1
	}
	if _, has := w.(http.Pusher); has {
		out |= 1 << 2
	}
	if _, has := w.(io.ReaderFrom); has {
		out |= 1 << 3
	}
	if _, has := w.(io.StringWriter); has {
		out |= 1 << 4
	}
	return
}

func newResponseProxy(proxied http.ResponseWriter) ResponseWriterProxier {
	b := &responseProxyBase{proxied: proxied}

	switch detectTraits(proxied) {
	case traitsBase:
		return b
	case traitsFl:
		return newResponseProxyFl(b)
	case traitsHj:
		return newResponseProxyHj(b)
	case traitsPh:
		return newResponseProxyPh(b)
	case traitsRf:
		return newResponseProxyRf(b)
	case traitsSw:
		return newResponseProxySw(b)
	case traitsFlHj:
		return newResponseProxyFlHj(b)
	case traitsFlPh:
		return newResponseProxyFlPh(b)
	case traitsFlRf:
		return newResponseProxyFlRf(b)
	case traitsFlSw:
		return newResponseProxyFlSw(b)
	case traitsHjPh:
		return newResponseProxyHjPh(b)
	case traitsHjRf:
		return newResponseProxyHjRf(b)
	case traitsHjSw:
		return newResponseProxyHjSw(b)
	case traitsPhRf:
		return newResponseProxyPhRf(b)
	case traitsPhSw:
		return newResponseProxyPhSw(b)
	case traitsRfSw:
		return newResponseProxyRfSw(b)
	case traitsFlHjPh:
		return newResponseProxyFlHjPh(b)
	case traitsFlHjRf:
		return newResponseProxyFlHjRf(b)
	case traitsFlHjSw:
		return newResponseProxyFlHjSw(b)
	case traitsFlPhRf:
		return newResponseProxyFlPhRf(b)
	case traitsFlPhSw:
		return newResponseProxyFlPhSw(b)
	case traitsFlRfSw:
		return newResponseProxyFlRfSw(b)
	case traitsHjPhRf:
		return newResponseProxyHjPhRf(b)
	case traitsHjPhSw:
		return newResponseProxyHjPhSw(b)
	case traitsHjRfSw:
		return newResponseProxyHjRfSw(b)
	case traitsPhRfSw:
		return newResponseProxyPhRfSw(b)
	case traitsFlHjPhRf:
		return newResponseProxyFlHjPhRf(b)
	case traitsFlHjPhSw:
		return newResponseProxyFlHjPhSw(b)
	case traitsFlHjRfSw:
		return newResponseProxyFlHjRfSw(b)
	case traitsFlPhRfSw:
		return newResponseProxyFlPhRfSw(b)
	case traitsHjPhRfSw:
		return newResponseProxyHjPhRfSw(b)
	case traitsFlHjPhRfSw:
		return newResponseProxyFlHjPhRfSw(b)
	default:
		panic("unreachable")
	}
}
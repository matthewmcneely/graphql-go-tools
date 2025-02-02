package pool

import (
	"sync"

	"github.com/matthewmcneely/graphql-go-tools/v2/pkg/fastbuffer"
)

var FastBuffer = fastBufferPool{
	pool: sync.Pool{
		New: func() interface{} {
			return fastbuffer.New()
		},
	},
}

type fastBufferPool struct {
	pool sync.Pool
}

func (f *fastBufferPool) Get() *fastbuffer.FastBuffer {
	return f.pool.Get().(*fastbuffer.FastBuffer)
}

func (f *fastBufferPool) Put(buf *fastbuffer.FastBuffer) {
	buf.Reset()
	f.pool.Put(buf)
}

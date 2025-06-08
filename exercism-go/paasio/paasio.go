package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	count  int64
	ops    int
	reader io.Reader
	mu     sync.Mutex
}

type writeCounter struct {
	count  int64
	ops    int
	writer io.Writer
	mu     sync.Mutex
}

type counter struct {
	r *readCounter
	w *writeCounter
}

func NewWriteCounter(writer io.Writer) IWriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) IReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) IReadWriteCounter {
	return &counter{
		r: &readCounter{reader: readwriter},
		w: &writeCounter{writer: readwriter},
	}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)
	if n > 0 {
		rc.mu.Lock()
		rc.count += int64(n)
		rc.ops++
		rc.mu.Unlock()
	}
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.count, rc.ops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)
	if n > 0 {
		wc.mu.Lock()
		wc.count += int64(n)
		wc.ops++
		wc.mu.Unlock()
	}
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.count, wc.ops
}

func (c *counter) Read(p []byte) (int, error) {
	return c.r.Read(p)
}

func (c *counter) ReadCount() (int64, int) {
	return c.r.ReadCount()
}

func (c *counter) Write(p []byte) (int, error) {
	return c.w.Write(p)
}

func (c *counter) WriteCount() (int64, int) {
	return c.w.WriteCount()
}

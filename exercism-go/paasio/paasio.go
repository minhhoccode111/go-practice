package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	reader io.Reader
	count  int64
	ops    int
	mu     sync.Mutex
}

type writeCounter struct {
	writer io.Writer
	count  int64
	ops    int
	mu     sync.Mutex
}

type counter struct {
	reader io.Reader
	writer io.Writer
	r      readCounter
	w      writeCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &counter{
		reader: readwriter,
		writer: readwriter,
		r:      readCounter{reader: readwriter},
		w:      writeCounter{writer: readwriter},
	}
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

package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	// a slice to store all values of the buffer
	bytes []byte
	// a read and write pointer
	read, write int
	// distance between the read and write pointer
	distance int
}

// return new buffer of given length
func NewBuffer(size int) *Buffer {
	return &Buffer{bytes: make([]byte, size)}
}

// read current byte in buffer
func (b *Buffer) ReadByte() (byte, error) {
	// nothing in buffer
	if b.distance == 0 {
		return 0, errors.New("buffer is empty")
	}
	// read pointer is greater than last index
	if b.read == len(b.bytes) {
		b.read = 0
	}
	// read data
	curr := b.bytes[b.read]
	// increase read pointer
	b.read++
	// decrease distance
	b.distance--
	// NOTE: no need to reset the value back to 0 because we manage empty by distance
	return curr, nil
}

func (b *Buffer) WriteByte(curr byte) error {
	// buffer is full
	if b.distance == len(b.bytes) {
		return errors.New("buffer is full")
	}
	// write pointer is greater than last index
	if b.write == len(b.bytes) {
		b.write = 0
	}
	// write data
	b.bytes[b.write] = curr
	// increase write pointer
	b.write++
	// increase distance
	b.distance++
	return nil
}

func (b *Buffer) Overwrite(curr byte) {
	// try write to buffer normally
	err := b.WriteByte(curr)
	// if there is an error
	if err != nil {
		// read the current byte, ignore the value returned
		b.ReadByte()
		// try again
		b.Overwrite(curr)
	}
}

func (b *Buffer) Reset() {
	// just dereference the b pointer and reinitialize new buffer :)
	*b = *NewBuffer(len(b.bytes))
}

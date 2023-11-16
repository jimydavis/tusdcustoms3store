package s3store

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

type Semaphore chan struct{}

// New creates a semaphore with the given concurrency limit.
func NewSemaphore(concurrency int) Semaphore {
	return make(chan struct{}, concurrency)
}

// Acquire will block until the semaphore can be acquired.
func (s Semaphore) Acquire() {
	s <- struct{}{}
}

// Release frees the acquired slot in the semaphore.
func (s Semaphore) Release() {
	<-s
}

// uid returns a unique id. These ids consist of 128 bits from a
// cryptographically strong pseudo-random generator and are like uuids, but
// without the dashes and significant bits.
// See: http://en.wikipedia.org/wiki/UUID#Random_UUID_probability_of_duplicates
func Uid() string {
	id := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, id)
	if err != nil {
		// This is probably an appropriate way to handle errors from our source
		// for random bits.
		panic(err)
	}
	return hex.EncodeToString(id)
}

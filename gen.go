package avatar

import (
	"crypto/md5"
	saferand "crypto/rand"
	"encoding/binary"
	"image"
	"math/rand"
	"sync"
)

type Generator interface {
	Generate(cfg Config) image.Image
}

type Config struct {
	Seed   string
	Width  int
	Height int
}

type DefaultDigest struct {
	Seed string
	//
	init    bool
	mutex   sync.Mutex
	r       *rand.Rand
	i64seed int64
}

func (w *DefaultDigest) Read(p []byte) (n int, err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if !w.init {
		if len(w.Seed) < 1 {
			bb := make([]byte, 8)
			_, err := saferand.Read(bb)
			if err != nil {
				return 0, err
			}
			w.i64seed, _ = binary.Varint(bb)
		} else {
			bb := md5.Sum([]byte(w.Seed))
			w.i64seed, _ = binary.Varint(bb[:]) // will read the fisrt 8
		}
		source := rand.NewSource(w.i64seed)
		w.r = rand.New(source)
		//
		w.init = true
	}
	return w.r.Read(p)
}

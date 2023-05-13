package image

import (
	"sync"
	"time"
)

const (
	getByTgUserError = "can't get user by tg user"
)

type ImageRegistry struct {
	sync.RWMutex
	images map[string]*Image
}

type Image struct {
	path           string
	fileId         string
	lastAccessTime time.Time
}

func New() *ImageRegistry {
	return &ImageRegistry{
		images: make(map[string]*Image),
	}
}

func (ir *ImageRegistry) RemoveInactive(minutes int) {
	ir.Lock()
	defer ir.Unlock()

	for path, image := range ir.images {
		if time.Since(image.lastAccessTime).Minutes() > float64(minutes) {
			delete(ir.images, path)
		}
	}
}

func (ir *ImageRegistry) Sync() {} // Losing file ids in case of restart is not a big deal

func (ir *ImageRegistry) GetByPath(path string) (image string) {
	ir.RLock()
	defer ir.RUnlock()

	if i, ok := ir.images[path]; ok {
		image = i.fileId
		i.lastAccessTime = time.Now()
	} else {
		image = path
	}

	return
}

func (ir *ImageRegistry) Save(path, fileId string) {
	ir.Lock()
	defer ir.Unlock()

	ir.images[path] = &Image{
		path:           path,
		fileId:         fileId,
		lastAccessTime: time.Now(),
	}
}

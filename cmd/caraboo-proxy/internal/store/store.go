package store

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Store struct {
	rwMutex sync.RWMutex
	saved   map[sig]*fasthttp.Response
}

type sig struct {
	Path   string
	Method string
}

func NewStore() *Store {
	return &Store{
		saved: make(map[sig]*fasthttp.Response),
	}
}

func (store *Store) Set(request *fiber.Request, copyFrom *fiber.Response) {
	s := store.asSig(request)
	r := fasthttp.AcquireResponse()
	copyFrom.CopyTo(r)

	store.rwMutex.Lock()
	defer store.rwMutex.Unlock()
	store.saved[s] = r
}

func (store *Store) Get(request *fiber.Request, into *fiber.Response) bool {
	s := store.asSig(request)

	store.rwMutex.RLock()
	found, ok := store.saved[s]
	store.rwMutex.RUnlock()

	if !ok {
		return false
	}

	found.CopyTo(into)
	return true
}

func (store *Store) asSig(request *fiber.Request) sig {
	return sig{
		Path:   string(request.URI().Path()),
		Method: string(request.Header.Method()),
	}
}

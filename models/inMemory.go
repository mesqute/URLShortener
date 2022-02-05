package models

import (
	u "URLShortener/utilites"
	"errors"
	"sync"
)

type value struct {
	url string
}

type inMemory struct {
	items map[string]value
	lock  sync.RWMutex
}

// SetDB активирует БД
func (i *inMemory) SetDB() *inMemory {
	db := inMemory{items: make(map[string]value)}
	return &db
}

// Get возвращает полный URL, соответствующий заданной ссылке
func (i *inMemory) Get(link string) (URL string, err error) {
	i.lock.RLock()
	defer i.lock.RUnlock()

	val, ok := i.items[link]
	if !ok {
		err = errors.New("in memory: no rows in result set")
		return
	}
	URL = val.url
	return
}

// Insert добавляет в БД заданный URL и возвращает сгенерированную ссылку
func (i *inMemory) Insert(URL string) (link string, err error) {
	_link, ok := i.findURL(URL)
	if ok { // если такой URL уже существует в БД
		link = _link // возвращает сгенерированную ранее ссылку
		return
	}
	link = i.findFreeToken()
	i.items[link] = value{url: URL}
	return
}

// findURL проверяет наличие URL в БД
func (i *inMemory) findURL(URL string) (link string, ok bool) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	for s, v := range i.items {
		if v.url == URL {
			ok = true
			link = s
		}
	}
	return
}

// findFreeToken ищет свободный токен
func (i *inMemory) findFreeToken() (token string) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	for {
		link := u.GenerateToken(10)
		_, ok := i.items[link]
		if !ok {
			token = link
			return
		}
	}
}

package impl

import "sync"

type MenuServiceImpl struct {
	wg sync.WaitGroup
}

func NewMenuServiceImpl() *MenuServiceImpl {
	return &MenuServiceImpl{}
}

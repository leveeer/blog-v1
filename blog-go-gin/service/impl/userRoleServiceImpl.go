package impl

import (
	"sync"
)

type UserRoleServiceImpl struct {
	wg sync.WaitGroup
}

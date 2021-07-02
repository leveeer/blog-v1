package common

import (
	"blog-go-gin/logging"
	"sync"
)

// Set Set集合
type Set struct {
	set map[interface{}]struct{}
	len int // 集合的大小
	*sync.RWMutex
}

func NewSet(cap int64) *Set {
	temp := make(map[interface{}]struct{}, cap)
	return &Set{
		set:     temp,
		RWMutex: &sync.RWMutex{},
	}
}

func (s *Set) Add(k interface{}) {
	s.Lock()
	defer s.Unlock()
	s.set[k] = struct{}{}
	logging.Logger.Debug(s.set)
	s.len = len(s.set) // 重新计算元素数量
}

func (s *Set) Remove(k interface{}) {
	s.Lock()
	s.Unlock()

	// 集合没元素直接返回
	if s.len == 0 {
		return
	}
	delete(s.set, k)   // 实际从字典删除这个键
	s.len = len(s.set) // 重新计算元素数量
}

func (s *Set) isExit(k interface{}) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.set[k]
	return ok
}

// Len 查看集合大小
func (s *Set) Len() int {
	return s.len
}

// IsEmpty 集合是否为空
func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// Clear 清除集合所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.set = map[interface{}]struct{}{} // 字典重新赋值
	s.len = 0                          // 大小归零
}

// ToList 将集合转化为切片
func (s *Set) ToList() []interface{} {
	s.RLock()
	defer s.RUnlock()
	list := make([]interface{}, 0, s.len)
	for item := range s.set {
		list = append(list, item)
	}
	return list
}

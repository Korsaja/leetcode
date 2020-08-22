package lru

import (
	"container/list"
)

type LRUCache struct {
	queue    *list.List
	storage  map[int]*list.Element
	//ptr      unsafe.Pointer
	capacity int


}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		queue:    list.New(),
		storage:  make(map[int]*list.Element),
		capacity: capacity,
	}
	//atomic.StorePointer(&lru.ptr,unsafe.Pointer(&lru.storage))
	return lru
}

func (this *LRUCache) Get(key int) int {

	//ptr := (*map[int]*list.Element)(atomic.LoadPointer(&this.ptr))
	el, ok := this.storage[key]
	if !ok{
		return -1
	}
	this.queue.MoveToFront(el)
	return el.Value.(*entry).value


}

func (this *LRUCache) Put(key int, value int) {


	_, ok := this.storage[key]
	if ok{
		el := this.storage[key]
		this.queue.MoveToFront(el)
		el.Value = value
	}else {
		if this.queue.Len() == this.capacity{
			el := this.queue.Back()
			this.queue.Remove(el)
			delete(this.storage,el.Value.(*entry).key)
		}
		this.queue.PushFront(&entry{key: key,value: value})
		this.storage[key] = this.queue.Front()
	}




}

package main

import (
	"container/list"
	"fmt"
)

/*
{
	from:"https://leetcode-cn.com/problems/lru-cache",
	reference:[],
	description:"LRU cache",
	solved:true,
}
*/

type LRUCache struct {
	Dict map[int]*list.Element
	List *list.List
	Cap  int
}

type KV struct {
	K int
	V int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Dict: make(map[int]*list.Element),
		List: list.New(),
		Cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Dict[key]; ok {
		var value = v.Value.(KV)
		this.List.PushFront(KV{
			K: key,
			V: value.V,
		})
		this.List.Remove(v)
		this.Dict[key] = this.List.Front()
		return value.V
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.Dict[key]; ok {
		this.List.Remove(v)
		this.List.PushFront(KV{
			K: key,
			V: value,
		})
		this.Dict[key] = this.List.Front()
	} else {
		if this.List.Len() == this.Cap {
			value := this.List.Back().Value.(KV)
			this.List.Remove(this.List.Back())
			delete(this.Dict,value.K)
		}
		this.List.PushFront(KV{
			K: key,
			V: value,
		})
		this.Dict[key] = this.List.Front()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(2, 4)
	obj.Put(1, 2)
	fmt.Println(obj.Get(2))
	obj.Put(3, 5)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
}

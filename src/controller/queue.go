// 用队列的形式存储获得的地址
package controller

import (
	"container/list"
	"strings"
)

// 处理URL的队列
// 记录处理过的URL防止重复处理，造成死循环
type queue struct {
	// 待处理的URL
	urls *list.List
	// 已处理的URL
	used *list.List
}

// 创建一个新的queue
func NewQueue() *queue {
	q := new(queue)

	q.urls = list.New()
	q.used = list.New()

	return q
}

// 检查指定的url是否已经处理过了
// 如果地址已经在待处理的队列中时，也不需要该地址放入队列
func (q *queue) IsExist(url string) bool {
	// 检查是否在已处理的队列中
	for e := q.used.Front(); e != nil; e = e.Next() {
		if e.Value == url {
			return true
		}
	}
	// 检查是否在待处理的队列中
	for e := q.urls.Front(); e != nil; e = e.Next() {
		if e.Value == url {
			return true
		}
	}

	return false
}

// 将url放入队列中
func (q *queue) Put(url string) {
	// 对入队的URL处理成统一的格式，并过滤掉不需要的地址
	// 过滤掉以http开头并且不为http://Root的站外地址
	if strings.Index(url, "http") == 0 &&
		strings.Index(url, spider.Root) == -1 {
		return
	}
	// 补全以/开头的地址
	if strings.Index(url, "/") == 0 {
		url = spider.Root + url
	}
	// 如果这个url不存在，就将它放到队列当中去
	if !q.IsExist(url) {
		q.urls.PushBack(url)
	}
}

// 返回队首的url，并将它标记为已处理
func (q *queue) Get() string {
	// 处理链表为空的情况，防止空指针
	if q.Len_urls() == 0 {
		return ""
	}
	front := q.urls.Remove(q.urls.Front())
	if str, ok := front.(string); ok {
		q.used.PushFront(str)
		return str
	}

	return ""
}

func (q *queue) Len_urls() int { return q.urls.Len() }

func (q *queue) Len_used() int { return q.used.Len() }

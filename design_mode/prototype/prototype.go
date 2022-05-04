package prototype

import (
	"encoding/json"
	"time"
)

type Keyword struct {
	Word     string
	Visit    int
	UpdateAt *time.Time
}

// 通过序列化和反序列化深拷贝
func (k *Keyword) Clone() *Keyword {
	var nk Keyword
	b, _ := json.Marshal(k)
	_ = json.Unmarshal(b, &nk)
	return &nk
}

// Keywords 关键字 map
type Keywords map[string]*Keyword

// Clone 复制一个新的keywords
// updateWords: 需要更新的关键词列表，由于从数据库中获取常常是数组的方式
func (words Keywords) Clone(updateWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	// 这里是浅拷贝，直接拷贝了地址
	for k, v := range words {
		newKeywords[k] = v
	}

	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updateWords {
		newKeywords[word.Word] = word.Clone()
	}

	return newKeywords
}

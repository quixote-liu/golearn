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

type Keywords map[string]*Keyword

func (k *Keywords) Clone()

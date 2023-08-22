package sensitive

import (
	"github.com/Tohrusky/chinese-sensitive-go/sensitive"
	"sync"
)

var (
	sensitiveInstance *sensitive.Filter
	once              sync.Once
)

// Init 初始化敏感词库，单例模式
func Init() {
	once.Do(func() {
		sensitiveInstance = sensitive.DefaultNew()
	})
}

// Validate 敏感词校验，校验一个句子是否包含敏感词，合法返回true，有敏感内容返回false
func Validate(s string) bool {
	b, _ := sensitiveInstance.Validate(s)
	return b
}

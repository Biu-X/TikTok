package sensitive

import (
	"sync"

	"github.com/Tohrusky/sensitive-go/sensitive"
)

var (
	sensitiveInstance     *sensitive.Filter
	sensitiveBossInstance *sensitive.Filter
	once                  sync.Once
)

// Init 初始化敏感词库，单例模式
func Init() {
	once.Do(func() {
		sensitiveInstance = sensitive.NewWithDefaultSDict()
		sensitiveBossInstance = sensitive.NewWithBossSDict()
	})
}

// ValidateBoss 重点敏感词校验，校验一个句子是否包含“重点”敏感词，合法返回true，有敏感内容返回false
func ValidateBoss(s string) bool {
	b, _ := sensitiveBossInstance.Validate(s)
	return b
}

// Replace 和谐敏感词，将敏感词替换为*号
func Replace(s string) string {
	return sensitiveInstance.Replace(s, '*')
}

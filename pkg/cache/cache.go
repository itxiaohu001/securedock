package cache

import (
	"context"
	"time"
)

// Cache 定义了一个通用的缓存接口，用于存储镜像层的检测结果
type Cache[T any] interface {
	// Set 将检测结果存储到缓存中
	// key: 缓存键（如镜像层ID）
	// value: 要缓存的检测结果
	// expiration: 过期时间，如果为0则表示永不过期
	// 返回error表示存储过程中的错误
	Set(ctx context.Context, key string, value T, expiration time.Duration) error

	// Get 从缓存中获取检测结果
	// key: 缓存键（如镜像层ID）
	// 返回缓存的检测结果和是否找到的标志
	// 如果发生错误，将通过error返回
	Get(ctx context.Context, key string) (T, bool, error)

	// Delete 从缓存中删除指定的检测结果
	// key: 要删除的缓存键
	// 返回error表示删除过程中的错误
	Delete(ctx context.Context, key string) error
}
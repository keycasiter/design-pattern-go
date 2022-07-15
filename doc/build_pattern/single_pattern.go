package build_pattern

import (
	"context"
	"time"
)

type Instance struct {
	Ctx context.Context
}

var hungryInstance *Instance = new(Instance)
var lazyInstance *Instance

//懒汉式
func GetInstanceByLazyModeWithBug(ctx context.Context) *Instance {
	if nil == lazyInstance {
		//模拟耗时
		time.Sleep(time.Second)
		lazyInstance = new(Instance)
	}
	return lazyInstance
}

//饿汉式
func GetInstanceByHungryMode(ctx context.Context) *Instance {
	return hungryInstance
}

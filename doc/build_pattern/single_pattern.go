package build_pattern

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Instance struct {
	Ctx context.Context
}

var hungryInstance *Instance = new(Instance)
var lazyInstance *Instance
var once sync.Once

//懒汉式-Bug版
func GetInstanceByLazyModeWithBug(ctx context.Context) *Instance {
	if nil == lazyInstance {
		//模拟耗时
		time.Sleep(time.Second)
		lazyInstance = new(Instance)
	}
	return lazyInstance
}

//懒汉式-Fix版
func GetInstanceByLazyModeWithFix(ctx context.Context) *Instance {
	if nil == lazyInstance {
		fmt.Println("初始化...")
		once.Do(func() {
			//模拟耗时
			time.Sleep(3 * time.Second)
			lazyInstance = new(Instance)
		})
	}
	return lazyInstance
}

//饿汉式
func GetInstanceByHungryMode(ctx context.Context) *Instance {
	return hungryInstance
}

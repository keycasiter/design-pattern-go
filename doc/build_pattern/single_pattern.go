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
var mutex sync.Mutex

//懒汉式-Bug版
func GetInstanceByLazyModeWithBug(ctx context.Context) *Instance {
	if nil == lazyInstance {
		//模拟耗时
		time.Sleep(time.Second)
		lazyInstance = new(Instance)
	}
	return lazyInstance
}

//懒汉式-Sync.Once版
func GetInstanceByLazyModeWithSyncOnce(ctx context.Context) *Instance {
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

//懒汉式-Sync.Mutex版
func GetInstanceByLazyModeWithSyncMutex(ctx context.Context) *Instance {
	mutex.Lock()
	defer mutex.Unlock()
	if nil == lazyInstance {
		fmt.Println("初始化...")
		//模拟耗时
		time.Sleep(3 * time.Second)
		lazyInstance = new(Instance)
	}
	return lazyInstance
}

//饿汉式
func GetInstanceByHungryMode(ctx context.Context) *Instance {
	return hungryInstance
}

package share_memory

import (
	"sync"
	"testing"
	"time"
)

// 线程不安全
func TestCounter(t *testing.T) {

	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d\n", counter) //counter = 4698
}

// safe
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second) //若是去掉此语句，test程序会比counter执行的快导致得出不正确结果
	t.Logf("counter = %d\n", counter)
}
/*
  大大加快运行时间
  先前并不知道5000个协程会运行多久，所以先睡眠1秒钟等待其完成再进行test协程
  现在增添了wait，只有协程任务都完成，才不会阻塞下面的程序运行
 */
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			//先上锁 再改数据
			mut.Lock()
			counter++

			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d\n", counter) //counter = 4698
}

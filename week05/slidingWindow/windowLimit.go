package slidingWindow

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)
// windows  is map[time.now()]window{value: int}
//
type SlidingWindow struct {
	Windows map[int64]*window
	Mutex   *sync.RWMutex
}

type window struct {
	Value int64
}

func NewSlidingWindow() *SlidingWindow {
	// 创建滑动窗口结构体
	return &SlidingWindow{
		Windows: make(map[int64]*window),
		Mutex:   &sync.RWMutex{},
	}
}

func (sw *SlidingWindow) getCurrentWindow() *window {
	// 当前时间时间戳
	now := time.Now().Unix()

	var w *window
	var ok bool
	if w, ok = sw.Windows[now]; !ok {
		w = &window{}
		sw.Windows[now] = w
	}

	return w
}

func (sw *SlidingWindow) removeOldWindows() {
	// 当前时间时间戳- 10
	now := time.Now().Unix() - 10
	// 循环 滑动窗口结构体
	for timestamp := range sw.Windows {
		if timestamp <= now {
			delete(sw.Windows, timestamp)
		}
	}
}

func (sw *SlidingWindow) Increment(i int64) {
	if i == 0 {
		return
	}

	sw.Mutex.Lock()
	defer sw.Mutex.Unlock()

	b := sw.getCurrentWindow()
	b.Value += i
	sw.removeOldWindows()
}

func (sw *SlidingWindow) Sum(now time.Time) (sum int64) {
	sw.Mutex.Lock()
	defer sw.Mutex.Unlock()

	for timestamp, window := range sw.Windows {
		if timestamp > now.Unix()-10 {
			fmt.Println(sum)
			sum += window.Value
		}
	}
	fmt.Println("sum" + strconv.Itoa(int(sum)))
	return sum
}

func (sw *SlidingWindow) Avg(now time.Time) int64 {
	return sw.Sum(now) / 10
}

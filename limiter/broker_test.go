package limiter

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestName(t *testing.T) {

}

func LimiterService(appid int) {
	rl := NewRequestLimier("HTTP GET", WithTimeOut(2*time.Second), WithMaxRequests(1))
	data, err := rl.Broker(GetGameData(appid))
	fmt.Println(data, err)
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetGameData(appid int) func() (interface{}, error) {
	return func() (interface{}, error) {
		// 在这里使用 appid 进行操作
		td := r.Intn(100)
		if td < appid {
			fmt.Println("error")
			return nil, errors.New("error")
		}
		return []byte("success"), nil
	}
}

func TestName111(t *testing.T) {
	// 创建一个 Calculator 实例
	calc := &Calculator{}
	// 链式调用
	result, err := calc.Add(5).Subtract(3).Add(10).Result()
	fmt.Println(result, err)

}

// 定义一个结构体
type Calculator struct {
	result int
	err    error
}

// 定义一个自定义错误类型
var ErrNegativeResult = errors.New("negative result not allowed")

// 定义结构体的方法，用于执行加法操作并返回新的结构体
func (c *Calculator) Add(x int) *Calculator {
	if x < 0 {
		c.err = ErrNegativeResult
		return c
	}
	c.result += x
	return c
}

// 定义结构体的方法，用于执行减法操作并返回新的结构体
func (c *Calculator) Subtract(x int) *Calculator {
	if x < 0 {
		c.err = ErrNegativeResult
		return c
	}
	c.result -= x
	return c
}

// 定义结构体的方法，用于获取最终计算结果和可能的错误
func (c *Calculator) Result() (int, error) {
	if c.result < 0 {
		c.err = ErrNegativeResult
	}
	return c.result, c.err
}

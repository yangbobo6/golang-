package func1

import "testing"

func TestPanic(t *testing.T){
	panicTest()

}

func panicTest() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()

	panic("panic error!")
}

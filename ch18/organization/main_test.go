package add

import (
	"fmt"
	"os"
	"testing"
)

// setupGlobal函数在整个测试执行之前运行
func setupGlobal() {
	fmt.Println("running setupGlobal...")
}

// teardownGlobal函数在整个测试执行之后运行
func teardownGlobal() {
	fmt.Println("running teardownGlobal...")
}

func TestMain(m *testing.M) {
	setupGlobal()
	r := m.Run()
	teardownGlobal()
	os.Exit(r)
}

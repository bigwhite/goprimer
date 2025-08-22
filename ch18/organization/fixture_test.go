package add

import (
	"testing"
)

// setup 函数在整个测试套件执行之前运行
func setupSuite(t *testing.T) {
	// 执行一些初始化操作，例如连接数据库或设置测试环境
	t.Log("running setupSuite...")
}

// teardown 函数在整个测试套件执行之后运行
func teardownSuite(t *testing.T) {
	// 清理测试环境，关闭数据库连接等
	t.Log("running teardownSuite...")
}

// setup 函数在每个测试用例执行之前运行
func setup(t *testing.T) {
	// 执行一些初始化操作，例如准备测试数据
	t.Log("running setup...")
}

// teardown 函数在每个测试用例执行之后运行
func teardown(t *testing.T) {
	// 清理测试环境
	t.Log("running teardown...")
}

func MyFunction(a int) int {
	return a * 2
}

func TestMyFunction(t *testing.T) {
	// 执行整个测试套件的 setup 函数
	setupSuite(t)

	// 定义测试用例的数据表格
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"testcase1", 1, 2},  // 测试用例1
		{"testcase2", 3, 6},  // 测试用例2
		{"testcase3", 5, 10}, // 测试用例3
	}

	// 遍历测试用例表格，逐个执行测试
	for _, tc := range testCases {
		// 执行每个测试用例的 setup 函数
		t.Run(tc.name, func(t *testing.T) {
			setup(t)

			result := MyFunction(tc.input)

			// 断言实际结果是否与预期结果相符
			if result != tc.expected {
				t.Errorf("want: %d, got: %d", tc.expected, result)
			}

			// 执行每个测试用例的 teardown 函数
			teardown(t)
		})
	}

	// 执行整个测试套件的 teardown 函数
	teardownSuite(t)
}

package test
import(
	"testing"
)

// Parallel 测试
func TestWrite(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		WriteToMap(tt.k,tt.v)
	}
}
func TestRead(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
	    actual := ReadFromMap(tt.k)
		if actual != tt.v {
			t.Errorf("错误的k %v",tt.k)
		}
	}
}

var pairs = []struct {
    k string
    v string
}{
    {"polaris", " 徐新华 "},
    {"studygolang", "Go 语言中文网 "},
    {"stdlib", "Go 语言标准库 "},
    {"polaris1", " 徐新华 1"},
    {"studygolang1", "Go 语言中文网 1"},
    {"stdlib1", "Go 语言标准库 1"},
    {"polaris2", " 徐新华 2"},
    {"studygolang2", "Go 语言中文网 2"},
    {"stdlib2", "Go 语言标准库 2"},
    {"polaris3", " 徐新华 3"},
    {"studygolang3", "Go 语言中文网 3"},
    {"stdlib3", "Go 语言标准库 3"},
    {"polaris4", " 徐新华 4"},
    {"studygolang4", "Go 语言中文网 4"},
    {"stdlib4", "Go 语言标准库 4"},
}


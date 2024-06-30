package behavior

import "fmt"

// ServerConfig 定义服务器配置
type ServerConfigV2 struct {
	Port    int
	Timeout int
}

// OptionV2 定义函数选项类型
type OptionV2 func(*ServerConfigV2)

// WithPort 设置服务器端口
func WithPortV2(port int) OptionV2 {
	return func(cfg *ServerConfigV2) {
		cfg.Port = port
	}
}

// WithTimeout 设置超时时间
func WithTimeoutV2(timeout int) OptionV2 {
	return func(cfg *ServerConfigV2) {
		cfg.Timeout = timeout
	}
}

// NewServer 创建一个新的服务器实例
func NewServerV2(options ...OptionV2) *ServerConfigV2 {
	cfg := &ServerConfigV2{
		Port:    8080,
		Timeout: 30,
	}
	for _, opt := range options {
		opt(cfg)
	}
	return cfg
}

func main() {
	// 创建服务器实例并指定选项
	server := NewServerV2(
		WithPortV2(9090),
		WithTimeoutV2(60),
	)

	fmt.Printf("Server Port: %d, Timeout: %d\n", server.Port, server.Timeout)
}

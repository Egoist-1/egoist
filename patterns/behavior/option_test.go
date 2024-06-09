package behavior

import "fmt"

// ServerConfig 定义服务器配置
type ServerConfig struct {
	Port    int
	Timeout int
}

// Option 定义函数选项类型
type Option func(*ServerConfig)

// WithPort 设置服务器端口
func WithPort(port int) Option {
	return func(cfg *ServerConfig) {
		cfg.Port = port
	}
}

// WithTimeout 设置超时时间
func WithTimeout(timeout int) Option {
	return func(cfg *ServerConfig) {
		cfg.Timeout = timeout
	}
}

// NewServer 创建一个新的服务器实例
func NewServer(options ...Option) *ServerConfig {
	cfg := &ServerConfig{
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
	server := NewServer(
		WithPort(9090),
		WithTimeout(60),
	)

	fmt.Printf("Server Port: %d, Timeout: %d\n", server.Port, server.Timeout)
}

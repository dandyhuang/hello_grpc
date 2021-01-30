package config

// Load 根据参数读取指定配置
func Load(path string) (*HelloConfig, error) {
	return DefaultConfig.load(path)
}

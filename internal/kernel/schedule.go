package kernel

type Schedule interface {
	Runtime

	// Data 数据
	Data() map[string]any
}

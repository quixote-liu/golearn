package builder

import "fmt"

// 参数传递法，这个参数是什么样的调用者决定，建造者可以对参数进行基础的校验
// 在大多数的场景下都可以使用这种方法，最大程度保证了公共库的兼容性，避免在后续更新的时候出现
// 破坏性更新的情况

type EndpointOptions struct {
	Type    string
	Region  string
	Service string
}

type EndpointOptionsFunc func(option *EndpointOptions)

func NewEndpointOptions(opts ...EndpointOptionsFunc) (*EndpointOptions, error) {
	eo := &EndpointOptions{
		// 如果有默认值，可以在这里进行设置
		Type:   "public",
		Region: "CHINA",
	}

	// 在这里进行设置属性值
	for _, opt := range opts {
		opt(eo)
	}

	// 进行基础的校验
	if eo.Region == "" {
		return nil, fmt.Errorf("the region is empty")
	}
	if eo.Service == "" {
		return nil, fmt.Errorf("the service is empty")
	}
	if eo.Type == "" {
		return nil, fmt.Errorf("the type is empty")
	}

	return eo, nil
}

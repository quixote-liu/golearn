package factory

// 椅子腿，实现了椅子腿的功能：支撑椅子功能
// 只要实现了椅子功能不管它的圆形、方形还是其他类型
type ChairLeg interface {
	Support() error
}

// 椅子坐垫，实现了坐的功能
// 只要实现了坐垫功能，不管它是软垫、硬垫还是其他类型
type ChairCushion interface {
	Sit() error
}

// 椅子背靠，实现了背靠的功能
// 只要实现了背靠功能，不管它是网靠、木头靠还是其他类型
type ChairBack interface {
	Back() error
}

// 椅子的组成需要：椅子腿、椅子坐垫、椅子背靠
type Chair interface {
	ChairLeg
	ChairCushion
	ChairBack
}

// 椅子工厂
type ChairFactory interface {
	CreateChairLeg() ChairLeg
	CreateChairCushion() ChairCushion
	CreateChairBack() ChairBack
}

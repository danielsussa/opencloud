package ipush

type Push struct {
	LoadPresets LoadPresets
}

var push Push

func Setup(p Push){
	push = p
}

func Get()Push {
	return push
}

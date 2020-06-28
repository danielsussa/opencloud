package program

type program struct {
}

type command interface {
	Run(chain, key, port string)
}

func New() {

}

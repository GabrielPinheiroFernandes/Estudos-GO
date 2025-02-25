package interfaces


type Document interface {
	Doc() string
}


type Controle interface {
	ProximoCanal()
	CanalAnterior()
}



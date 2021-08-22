package common

type Generator interface{
	Create() Instance
}

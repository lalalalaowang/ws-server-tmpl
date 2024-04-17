package net

type groupFunc func()

type router struct {
	groups *[]group
}

type group struct {
	prefix       string
	groupFuncMap map[string]groupFunc
}

package higo

/**
框架内方法
*/

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

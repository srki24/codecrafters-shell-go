package main

var functions []go_function

type go_function struct {
	name string
	impl func(argv ...string)
}

func getAllFunctions() []go_function {
	if functions == nil {
		functions = append(functions,
			go_function{name: "exit", impl: go_exit},
			go_function{name: "echo", impl: go_echo},
			go_function{name: "type", impl: go_type},
			go_function{name: "pwd", impl: go_pwd},
			go_function{name: "cd", impl: go_cd},
		)
	}

	return functions
}

func getFunction(name string) (go_func func(argv ...string), ok bool) {
	go_func = nil
	ok = false

	if functions == nil {
		functions = getAllFunctions()
	}

	for _, f := range functions {
		if f.name == name {
			go_func = f.impl
			ok = true
			break
		}
	}
	return
}

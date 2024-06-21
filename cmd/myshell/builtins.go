package main

var functions []go_function

type go_function struct {
	name string
	impl func(argv []string)
}

func getAllFunctions() []go_function {
	if functions == nil {
		functions = append(functions,
			go_function{name: "exit", impl: go_exit},
			go_function{name: "echo", impl: go_echo},
			go_function{name: "type", impl: go_type},
		)
	}

	return functions
}

func getFunction(name string) (go_func func(argv []string), ok bool) {
	if functions == nil {
		functions = getAllFunctions()
	}

	go_func = nil
	ok = false
	for _, f := range functions {
		if f.name == name {
			go_func = f.impl
			ok = true
			break
		}
	}
	return
}

func hasFunction(name string) (ok bool) {
	if functions == nil {
		functions = getAllFunctions()
	}

	ok = false
	for _, f := range functions {
		if f.name == name {
			ok = true
			break
		}
	}
	return
}

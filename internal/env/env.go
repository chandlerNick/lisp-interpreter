// internal/env/env.go
package env

type Environment struct {
    store map[string]interface{}
}

func NewEnvironment() *Environment {
    env := &Environment{store: make(map[string]interface{})}
    // Register built-ins immediately
    env.RegisterBuiltins()
    return env
}

func (e *Environment) Get(name string) interface{} {
	return e.store[name]
}

func (e *Environment) RegisterBuiltins() {
    e.store["+"] = func(args []interface{}) interface{} {
        sum := int64(0)
        for _, arg := range args {
            sum += arg.(int64)
        }
        return sum
    }
    
    e.store["*"] = func(args []interface{}) interface{} {
        prod := int64(1)
        for _, arg := range args {
            prod *= arg.(int64)
        }
        return prod
    }

	e.store["-"] = func(args []interface{}) interface{} {
        // Lisp: (- 10 2) = 8; (- 10) = -10 (negation)
        if len(args) == 1 {
            return -args[0].(int64)
        }
        res := args[0].(int64)
        for _, arg := range args[1:] {
            res -= arg.(int64)
        }
        return res
    }

    e.store["/"] = func(args []interface{}) interface{} {
        // Lisp: (/ 10 2) = 5
        res := args[0].(int64)
        for _, arg := range args[1:] {
            res /= arg.(int64)
        }
        return res
    }
}
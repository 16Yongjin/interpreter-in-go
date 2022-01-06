package evaluator

import "monkey/object"

const (
	BUILTIN_FUNC_NAME_LEN   = "len"
	BUILTIN_FUNC_NAME_FIRST = "first"
	BUILTIN_FUNC_NAME_LAST  = "last"
	BUILTIN_FUNC_NAME_REST  = "rest"
	BUILTIN_FUNC_NAME_PUSH  = "push"
	BUILTIN_FUNC_NAME_PUTS  = "puts"
)

var builtins = map[string]*object.Builtin{
	BUILTIN_FUNC_NAME_LEN: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to %q not supported, got %s", BUILTIN_FUNC_NAME_LEN, args[0].Type())
			}
		},
	},
}

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
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to %q not supported, got %s", BUILTIN_FUNC_NAME_LEN, args[0].Type())
			}
		},
	},
	BUILTIN_FUNC_NAME_FIRST: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to %q must be %s, got %s", BUILTIN_FUNC_NAME_FIRST, object.ARRAY_OBJ, args[0].Type())
			}

			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[0]
			}

			return NULL
		},
	},
	BUILTIN_FUNC_NAME_LAST: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to %q must be %s, got %s", BUILTIN_FUNC_NAME_LAST, object.ARRAY_OBJ, args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)
			if length > 0 {
				return array.Elements[length-1]
			}

			return NULL
		},
	},
	BUILTIN_FUNC_NAME_REST: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to %q must be %s, got %s", BUILTIN_FUNC_NAME_REST, object.ARRAY_OBJ, args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, array.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},

	BUILTIN_FUNC_NAME_PUSH: {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to %q must be %s, got %s", BUILTIN_FUNC_NAME_PUSH, object.ARRAY_OBJ, args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, array.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
}

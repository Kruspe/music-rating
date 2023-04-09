package api

import "strings"

func match(path, pattern string, vars ...interface{}) bool {
	for ; pattern != "" && path != ""; pattern = pattern[1:] {
		switch pattern[0] {
		case '+':
			slash := strings.IndexByte(path, '/')
			if slash < 0 {
				slash = len(path)
			}
			segment := path[:slash]
			path = path[slash:]
			switch p := vars[0].(type) {
			case *string:
				*p = segment
			default:
				panic("vars must be *string")
			}
			vars = vars[1:]
		case path[0]:
			path = path[1:]
		default:
			return false
		}
	}

	return path == "" && pattern == ""
}

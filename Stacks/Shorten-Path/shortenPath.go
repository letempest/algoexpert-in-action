package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/foo/../test/../test/../foo//bar/./baz"
	fmt.Printf("%#v\n", shortenPath(path))
	path2 := "foo/../../../something/././//hello/../byebye/."
	fmt.Printf("%#v\n", shortenPath(path2))
}

// Big O: O(n) time | O(n) space, n is the length of the path string
func shortenPath(path string) string {
	isAbsolute := strings.HasPrefix(path, "/")
	splitPath := strings.Split(path, "/")
	stack := make([]string, 0)
	if isAbsolute {
		stack = append(stack, "")
	}

	filteredTokens := []string{}
	for _, token := range splitPath {
		if len(token) > 0 && token != "." {
			filteredTokens = append(filteredTokens, token)
		}
	}

	for _, token := range filteredTokens {
		if token == ".." {
			if len(stack) == 0 || stack[len(stack)-1] == ".." {
				// only when (path is relative and stack is empty, i.e. at the beginning of relative path) OR last element in stack is "..", we would append this ".."
				// last element in stack is ".." only possible for relative path case, for absolute path, we cannot reach beyond the root directory
				// so for absolute path string, it's impossible to have trailing ".." in the stack, as it would either cancel against previous token, or resolved to root path (which we ignore it)
				stack = append(stack, token)
			} else if stack[len(stack)-1] != "" {
				// if we make into this block, it doesn't matter whether the src path is absolute or relative,
				// so long as last token is a normal directory but not the root (!= ""), current token, which is "..", should cancel out the last token in the stack
				stack = stack[:len(stack)-1]
			}
		} else {
			// for normal directory, e.g. "foo", "bar", it's safe to append it to the stack
			stack = append(stack, token)
		}
	}
	// edge case, if stack =[""], then return root path directly, as join method will only return empty string for this case
	if len(stack) == 1 && stack[0] == "" {
		return "/"
	}
	return strings.Join(stack, "/")
}

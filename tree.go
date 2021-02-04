package gRouter

import "regexp"

var (
	methodList = []string{
		"POST",
		"GET",
		"HEAD",
		"OPTIONS",
		"PUT",
		"DELETE",
		"PATCH",
		"CONNECT",
		"TRACE",
	}
)

type tree struct {
	method string
	root   *node
}

func newTree(method string) *tree {
	if _isDebug {
		checkMethod(method)
	}
	return &tree{
		method: method,
		root:   newNode("", true),
	}
}

func checkMethod(method string) {
	for i := 0; i < len(methodList); i++ {
		if methodList[i] == method {
			return
		}
	}
	panic("method is error")
}

func checkUrl(url string) {
	rule := "[a-zA-Z0-9-._~!$&'/()*+,;=:]*"
	reg, err := regexp.Compile(rule)
	if err != nil {
		panic(err)
	}
	if !reg.Match([]byte(url)) {
		panic(errorUrlFormat)
	}
}

func (tree *tree) Add(url string, handlers HandlersChain) error {
	if _isDebug {
		checkUrl(url)
	}
	return tree.root.Add(url, handlers)
}

func (tree *tree) Find(url string) (HandlersChain, error) {
	node, err := tree.find(url)
	if err != nil {
		return nil, err
	}
	return node.handlers, nil
}

func (tree *tree) find(url string) (*node, error) {
	return tree.root.Find(url)
}

func (tree *tree) PathList() []string {
	list := make([]string, 0)
	tree.pathList(&list, tree.root, "")
	return list
}

func (tree *tree) pathList(list *[]string, root *node, subPath string) {
	if root == nil {
		return
	}
	if root.path != "" {
		subPath += "/" + root.path
	}
	if root.isLeaf {
		if root.isRoot {
			*list = append(*list, "/")
		} else {
			*list = append(*list, subPath)
		}
	}
	for i := 0; i < len(root.children); i++ {
		tree.pathList(list, root.children[i], subPath)
	}
}

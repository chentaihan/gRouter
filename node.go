package gRouter

import (
	"errors"
	"strings"
)

var (
	errorUrlNotFound = errors.New("url not found")
	errorUrlFormat   = errors.New("url must start with /")
)

type nodeType uint8

const (
	nodeTypeNormal nodeType = 0 //uri普通节点
	nodeTypeParam  nodeType = 1 //uri中restful参数
	nodeTypeAll    nodeType = 2 //uri中*全部匹配
)

//路径按照这个顺序匹配
var nodeTypeList = [3]nodeType{
	nodeTypeNormal,
	nodeTypeParam,
	nodeTypeAll,
}

type node struct {
	path      string
	pathMatch string
	isRoot    bool
	isLeaf    bool
	nType     nodeType
	children  []*node
	parent    *node
	handlers  HandlersChain
}

func newNode(path string, isRoot bool) *node {
	return &node{
		path:   path,
		isRoot: isRoot,
	}
}

func (node *node) Find(url string) (*node, error) {
	url = strings.TrimSpace(url)
	if url == "" || url[0] != '/' {
		return nil, errorUrlFormat
	}
	if url == "/" {
		return node, nil
	}
	paths := node.removeEmptyPath(strings.Split(url, "/"))
	return node.find(paths)
}

//删除为空的路径，第0个元素是空的，且不删除
func (node *node) removeEmptyPath(paths []string) []string {
	var list []string
	for i := 1; i < len(paths); i++ {
		if paths[i] == "" {
			if list == nil {
				list = make([]string, i, len(paths)-1)
				copy(list, paths[:i])
			}
		} else if list != nil {
			list = append(list, paths[i])
		}
	}
	if len(list) == 0 {
		return paths
	}
	return list
}

func (node *node) find(paths []string) (*node, error) {
	if len(paths) == 0 {
		return nil, errorUrlNotFound
	}

	if err := node.checkPath(paths[0]); err != nil {
		return nil, err
	}
	if len(paths) == 1 {
		if node.isLeaf {
			return node, nil
		}
		//当前节点不是叶子节点，说明路径没有完整匹配，地址没找到
		return nil, errorUrlNotFound
	}

	//1：nodeTypeNormal 优先级高
	for i := 0; i < len(node.children); i++ {
		if node.children[i].path == paths[1] && node.children[i].nType == nodeTypeNormal {
			return node.children[i].find(paths[1:])
		}
	}

	//2：nodeTypeParam 优先级中
	//3：nodeTypeAll 优先级低
	for _, nType := range []nodeType{nodeTypeParam, nodeTypeAll} {
		for i := 0; i < len(node.children); i++ {
			if node.children[i].nType == nType {
				result, err := node.children[i].find(paths[1:])
				if err == nil {
					return result, err
				}
			}
		}
	}

	return nil, errorUrlNotFound
}

func (node *node) checkPath(path string) error {
	if node.isNodeTypeNormal() && node.path != path {
		return errorUrlNotFound
	}
	return nil
}

func (node *node) Add(url string, handlers HandlersChain) error {
	url = strings.TrimSpace(url)
	if url == "" || url[0] != '/' {
		return errorUrlFormat
	}
	if url == "/" {
		node.handlers = append(node.handlers, handlers...)
		node.isLeaf = true
		return nil
	}
	paths := node.removeEmptyPath(strings.Split(url, "/"))
	return node.add(paths[1:], handlers)
}

func (node *node) isNodeTypeNormal() bool {
	return node.nType == nodeTypeNormal
}

func (node *node) getNodeType(path string) nodeType {
	nType := nodeTypeNormal
	if len(path) > 0 && path[0] == ':' {
		nType = nodeTypeParam
	} else if path == "*" {
		nType = nodeTypeAll
	}
	return nType
}

func (node *node) add(paths []string, handlers HandlersChain) error {
	path := paths[0]
	nType := node.getNodeType(path)
	if nType == nodeTypeParam {
		path = path[1:]
	}
	for index := 0; index < len(nodeTypeList); index++ {
		if nodeTypeList[index] != nType {
			continue
		}
		for i := 0; i < len(node.children); i++ {
			if node.children[i].path == path {
				return node.children[i].add(paths[1:], handlers)
			}
		}
	}

	sonNode := newNode(path, false)
	sonNode.parent = node
	sonNode.nType = nType
	node.children = append(node.children, sonNode)
	if len(paths) == 1 {
		sonNode.handlers = handlers
		sonNode.isLeaf = true
	} else {
		sonNode.add(paths[1:], handlers)
	}
	return nil
}

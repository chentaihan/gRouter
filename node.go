package gRouter

import (
	"errors"
	"strings"
)

var (
	errorUrlNotFound = errors.New("url not found")
	errorUrlFormat   = errors.New("url must start with /")
)

type node struct {
	path     string
	isRoot   bool
	isLeaf   bool
	children []*node
	parent   *node
	handlers HandlersChain
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
	paths := strings.Split(url, "/")
	return node.find(paths)
}

func (node *node) find(paths []string) (*node, error) {
	if len(paths) == 0 {
		return nil, errorUrlNotFound
	}
	if node.path != paths[0] {
		return nil, errorUrlNotFound
	}
	if len(paths) == 1 {
		return node, nil
	}
	for i := 0; i < len(node.children); i++ {
		if node.children[i].path == paths[1] {
			return node.children[i].find(paths[1:])
		}
	}
	return nil, errorUrlNotFound
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
	paths := strings.Split(url, "/")
	return node.add(paths[1:], handlers)
}

func (node *node) add(paths []string, handlers HandlersChain) error {
	for i := 0; i < len(node.children); i++ {
		if node.children[i].path == paths[0] {
			return node.children[i].add(paths[1:], handlers)
		}
	}
	sonNode := newNode(paths[0], false)
	sonNode.parent = node
	node.children = append(node.children, sonNode)
	if len(paths) == 1 {
		sonNode.handlers = handlers
		sonNode.isLeaf = true
	} else {
		sonNode.add(paths[1:], handlers)
	}
	return nil
}

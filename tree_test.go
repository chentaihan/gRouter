package gRouter

import (
	"testing"
)

func handlerTest(c *Context) {

}

var urlsTest = []string{
	"/",
	"/0",
	"/0/1",
	"/0/2",
	"/0/1/3",
	"/0/1/3/7",
	"/0/1/3/8",
	"/0/1/4",
	"/0/1/4/9",
	"/0/1/4/10",
	"/0/2/5",
	"/0/2/5/11",
	"/0/2/5/12",
	"/0/2/6",
	"/0/2/6/13",
	"/0/2/6/14",
}

func createTree() *tree {
	tree := newTree("POST")
	for _, url := range urlsTest {
		tree.Add(url, handlerTest)
	}
	return tree
}

func TestTree_Add(t *testing.T) {
	tree := createTree()
	pathList := tree.PathList()
	if !StringSortEqual(urlsTest, pathList) {
		t.Fatal("createTree error")
	}
}

func TestTree_Find(t *testing.T) {
	tree := createTree()
	for _, url := range urlsTest {
		handlers, err := tree.Find(url)
		if err != nil {
			t.Fatalf("url=%v, error=%v", url, err)
		}
		if len(handlers) == 0 {
			t.Fatalf("url=%v, error=not found", url)
		}
	}
	t.Logf("%v success", len(urlsTest))
}

package gRouter

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

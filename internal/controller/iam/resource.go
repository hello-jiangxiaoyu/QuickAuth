package iam

type TreeNode struct {
	Code  string     `json:"code"`
	Value any        `json:"value"`
	Child []TreeNode `json:"child"`
}

type ResourceValue struct {
	Json   []TreeNode `json:"json"`
	String string     `json:"string"`
	Number int        `json:"number"`
}

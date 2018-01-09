package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type ClassMethod struct {
	ReturnsRef    bool
	PhpDocComment string
	MethodName    node.Node
	Modifiers     []node.Node
	Params        []node.Node
	ReturnType    node.Node
	Stmts         []node.Node
}

func NewClassMethod(MethodName node.Node, Modifiers []node.Node, ReturnsRef bool, Params []node.Node, ReturnType node.Node, Stmts []node.Node, PhpDocComment string) *ClassMethod {
	return &ClassMethod{
		ReturnsRef,
		PhpDocComment,
		MethodName,
		Modifiers,
		Params,
		ReturnType,
		Stmts,
	}
}

func (n *ClassMethod) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"PhpDocComment": n.PhpDocComment,
	}
}

func (n *ClassMethod) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.MethodName != nil {
		vv := v.GetChildrenVisitor("MethodName")
		n.MethodName.Walk(vv)
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			nn.Walk(vv)
		}
	}

	if n.Params != nil {
		vv := v.GetChildrenVisitor("Params")
		for _, nn := range n.Params {
			nn.Walk(vv)
		}
	}

	if n.ReturnType != nil {
		vv := v.GetChildrenVisitor("ReturnType")
		n.ReturnType.Walk(vv)
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}

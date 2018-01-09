package expr

import (
	"github.com/z7zmey/php-parser/node"
)

type UnaryPlus struct {
	Expr node.Node
}

func NewUnaryPlus(Expression node.Node) *UnaryPlus {
	return &UnaryPlus{
		Expression,
	}
}

func (n *UnaryPlus) Attributes() map[string]interface{} {
	return nil
}

func (n *UnaryPlus) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}

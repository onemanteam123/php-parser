package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type SmallerOrEqual struct {
	BinaryOp
}

func NewSmallerOrEqual(Variable node.Node, Expression node.Node) *SmallerOrEqual {
	return &SmallerOrEqual{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *SmallerOrEqual) Attributes() map[string]interface{} {
	return nil
}

func (n *SmallerOrEqual) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		vv := v.GetChildrenVisitor("Left")
		n.Left.Walk(vv)
	}

	if n.Right != nil {
		vv := v.GetChildrenVisitor("Right")
		n.Right.Walk(vv)
	}

	v.LeaveNode(n)
}

package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type Mul struct {
	BinaryOp
}

func NewMul(Variable node.Node, Expression node.Node) *Mul {
	return &Mul{
		BinaryOp{
			Variable,
			Expression,
		},
	}
}

func (n *Mul) Attributes() map[string]interface{} {
	return nil
}

func (n *Mul) Walk(v node.Visitor) {
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

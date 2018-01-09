package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type ShiftLeft struct {
	AssignOp
}

func NewShiftLeft(Variable node.Node, Expression node.Node) *ShiftLeft {
	return &ShiftLeft{
		AssignOp{
			Variable,
			Expression,
		},
	}
}

func (n *ShiftLeft) Attributes() map[string]interface{} {
	return nil
}

func (n *ShiftLeft) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Expression != nil {
		vv := v.GetChildrenVisitor("Expression")
		n.Expression.Walk(vv)
	}

	v.LeaveNode(n)
}

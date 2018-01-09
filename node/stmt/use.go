package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Use struct {
	UseType node.Node
	Use     node.Node
	Alias   node.Node
}

func NewUse(UseType node.Node, use node.Node, Alias node.Node) *Use {
	return &Use{
		UseType,
		use,
		Alias,
	}
}

func (n *Use) Attributes() map[string]interface{} {
	return nil
}

func (n *Use) SetUseType(UseType node.Node) node.Node {
	n.UseType = UseType
	return n
}

func (n *Use) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.UseType != nil {
		vv := v.GetChildrenVisitor("UseType")
		n.UseType.Walk(vv)
	}

	if n.Use != nil {
		vv := v.GetChildrenVisitor("Use")
		n.Use.Walk(vv)
	}

	if n.Alias != nil {
		vv := v.GetChildrenVisitor("Alias")
		n.Alias.Walk(vv)
	}

	v.LeaveNode(n)
}

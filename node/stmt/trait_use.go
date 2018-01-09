package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type TraitUse struct {
	Traits      []node.Node
	Adaptations []node.Node
}

func NewTraitUse(Traits []node.Node, Adaptations []node.Node) *TraitUse {
	return &TraitUse{
		Traits,
		Adaptations,
	}
}

func (n *TraitUse) Attributes() map[string]interface{} {
	return nil
}

func (n *TraitUse) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Traits != nil {
		vv := v.GetChildrenVisitor("Traits")
		for _, nn := range n.Traits {
			nn.Walk(vv)
		}
	}

	if n.Adaptations != nil {
		vv := v.GetChildrenVisitor("Adaptations")
		for _, nn := range n.Adaptations {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}

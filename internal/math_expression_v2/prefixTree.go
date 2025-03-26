package math_expression_v2

//type PrefixTreeExpression interface {
//	Insert()
//}
//
//type prefixTreeExpression struct {
//	prefix   Expression
//	subTrees map[Expression]prefixTreeExpression
//}
//
//func (p *prefixTreeExpression) Insert(e Expression) {
//	if p.prefix.Compare(&e) == Equal {
//		return
//	}
//	if e.Compare(&p.prefix) == Prefix {
//		for expression, treeExpression := range p.subTrees {
//			if e.Compare(&expression) != Different {
//				treeExpression.Insert(e)
//				return
//			}
//		}
//	}
//}
//
//func New() PrefixTreeExpression {
//
//}

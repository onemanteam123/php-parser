package ast

type Vertex interface {
	Accept(v NodeVisitor)
	GetNode() *Node
}

type Traverser interface {
	Traverse(n Vertex)
}

type Visitor interface {
	Enter(key string, singleNode bool)
	Leave(key string, singleNode bool)

	EnterNode(n Vertex) bool
	LeaveNode(n Vertex)
}

type NodeVisitor interface {
	Root(n *Root)
	Nullable(n *Nullable)
	Reference(n *Reference)
	Variadic(n *Variadic)
	Parameter(n *Parameter)
	Identifier(n *Identifier)
	ArgumentList(n *ArgumentList)
	Argument(n *Argument)

	StmtAltForeach(n *StmtAltForeach)
	StmtAltSwitch(n *StmtAltSwitch)
	StmtBreak(n *StmtBreak)
	StmtCase(n *StmtCase)
	StmtCaseList(n *StmtCaseList)
	StmtCatch(n *StmtCatch)
	StmtClass(n *StmtClass)
	StmtClassConstList(n *StmtClassConstList)
	StmtClassExtends(n *StmtClassExtends)
	StmtClassImplements(n *StmtClassImplements)
	StmtClassMethod(n *StmtClassMethod)
	StmtConstList(n *StmtConstList)
	StmtConstant(n *StmtConstant)
	StmtContinue(n *StmtContinue)
	StmtDeclare(n *StmtDeclare)
	StmtDefault(n *StmtDefault)
	StmtDo(n *StmtDo)
	StmtEcho(n *StmtEcho)
	StmtElse(n *StmtElse)
	StmtElseIf(n *StmtElseIf)
	StmtExpression(n *StmtExpression)
	StmtFinally(n *StmtFinally)
	StmtFor(n *StmtFor)
	StmtForeach(n *StmtForeach)
	StmtFunction(n *StmtFunction)
	StmtGlobal(n *StmtGlobal)
	StmtGoto(n *StmtGoto)
	StmtHaltCompiler(n *StmtHaltCompiler)
	StmtIf(n *StmtIf)
	StmtInlineHtml(n *StmtInlineHtml)
	StmtInterface(n *StmtInterface)
	StmtInterfaceExtends(n *StmtInterfaceExtends)
	StmtLabel(n *StmtLabel)
	StmtNamespace(n *StmtNamespace)
	StmtNop(n *StmtNop)
	StmtProperty(n *StmtProperty)
	StmtPropertyList(n *StmtPropertyList)
	StmtReturn(n *StmtReturn)
	StmtStatic(n *StmtStatic)
	StmtStaticVar(n *StmtStaticVar)
	StmtStmtList(n *StmtStmtList)
	StmtSwitch(n *StmtSwitch)
	StmtThrow(n *StmtThrow)
	StmtTrait(n *StmtTrait)
	StmtTraitAdaptationList(n *StmtTraitAdaptationList)
	StmtTraitMethodRef(n *StmtTraitMethodRef)
	StmtTraitUse(n *StmtTraitUse)
	StmtTraitUseAlias(n *StmtTraitUseAlias)
	StmtTraitUsePrecedence(n *StmtTraitUsePrecedence)
	StmtTry(n *StmtTry)
	StmtUnset(n *StmtUnset)
	StmtUse(n *StmtUse)
	StmtGroupUse(n *StmtGroupUse)
	StmtUseDeclaration(n *StmtUseDeclaration)
	StmtWhile(n *StmtWhile)

	ExprArray(n *ExprArray)
	ExprArrayDimFetch(n *ExprArrayDimFetch)
	ExprArrayItem(n *ExprArrayItem)
	ExprArrowFunction(n *ExprArrowFunction)
	ExprBitwiseNot(n *ExprBitwiseNot)
	ExprBooleanNot(n *ExprBooleanNot)
	ExprClassConstFetch(n *ExprClassConstFetch)
	ExprClone(n *ExprClone)
	ExprClosure(n *ExprClosure)
	ExprClosureUse(n *ExprClosureUse)
	ExprConstFetch(n *ExprConstFetch)
	ExprEmpty(n *ExprEmpty)
	ExprErrorSuppress(n *ExprErrorSuppress)
	ExprEval(n *ExprEval)
	ExprExit(n *ExprExit)
	ExprFunctionCall(n *ExprFunctionCall)
	ExprInclude(n *ExprInclude)
	ExprIncludeOnce(n *ExprIncludeOnce)
	ExprInstanceOf(n *ExprInstanceOf)
	ExprIsset(n *ExprIsset)
	ExprList(n *ExprList)
	ExprMethodCall(n *ExprMethodCall)
	ExprNew(n *ExprNew)
	ExprPostDec(n *ExprPostDec)
	ExprPostInc(n *ExprPostInc)
	ExprPreDec(n *ExprPreDec)
	ExprPreInc(n *ExprPreInc)
	ExprPrint(n *ExprPrint)
	ExprPropertyFetch(n *ExprPropertyFetch)
	ExprReference(n *ExprReference)
	ExprRequire(n *ExprRequire)
	ExprRequireOnce(n *ExprRequireOnce)
	ExprShellExec(n *ExprShellExec)
	ExprShortArray(n *ExprShortArray)
	ExprShortList(n *ExprShortList)
	ExprStaticCall(n *ExprStaticCall)
	ExprStaticPropertyFetch(n *ExprStaticPropertyFetch)
	ExprTernary(n *ExprTernary)
	ExprUnaryMinus(n *ExprUnaryMinus)
	ExprUnaryPlus(n *ExprUnaryPlus)
	ExprVariable(n *ExprVariable)
	ExprYield(n *ExprYield)
	ExprYieldFrom(n *ExprYieldFrom)

	ExprAssign(n *ExprAssign)
	ExprAssignReference(n *ExprAssignReference)
	ExprAssignBitwiseAnd(n *ExprAssignBitwiseAnd)
	ExprAssignBitwiseOr(n *ExprAssignBitwiseOr)
	ExprAssignBitwiseXor(n *ExprAssignBitwiseXor)
	ExprAssignCoalesce(n *ExprAssignCoalesce)
	ExprAssignConcat(n *ExprAssignConcat)
	ExprAssignDiv(n *ExprAssignDiv)
	ExprAssignMinus(n *ExprAssignMinus)
	ExprAssignMod(n *ExprAssignMod)
	ExprAssignMul(n *ExprAssignMul)
	ExprAssignPlus(n *ExprAssignPlus)
	ExprAssignPow(n *ExprAssignPow)
	ExprAssignShiftLeft(n *ExprAssignShiftLeft)
	ExprAssignShiftRight(n *ExprAssignShiftRight)

	ExprBinaryBitwiseAnd(n *ExprBinaryBitwiseAnd)
	ExprBinaryBitwiseOr(n *ExprBinaryBitwiseOr)
	ExprBinaryBitwiseXor(n *ExprBinaryBitwiseXor)
	ExprBinaryBooleanAnd(n *ExprBinaryBooleanAnd)
	ExprBinaryBooleanOr(n *ExprBinaryBooleanOr)
	ExprBinaryCoalesce(n *ExprBinaryCoalesce)
	ExprBinaryConcat(n *ExprBinaryConcat)
	ExprBinaryDiv(n *ExprBinaryDiv)
	ExprBinaryEqual(n *ExprBinaryEqual)
	ExprBinaryGreater(n *ExprBinaryGreater)
	ExprBinaryGreaterOrEqual(n *ExprBinaryGreaterOrEqual)
	ExprBinaryIdentical(n *ExprBinaryIdentical)
	ExprBinaryLogicalAnd(n *ExprBinaryLogicalAnd)
	ExprBinaryLogicalOr(n *ExprBinaryLogicalOr)
	ExprBinaryLogicalXor(n *ExprBinaryLogicalXor)
	ExprBinaryMinus(n *ExprBinaryMinus)
	ExprBinaryMod(n *ExprBinaryMod)
	ExprBinaryMul(n *ExprBinaryMul)
	ExprBinaryNotEqual(n *ExprBinaryNotEqual)
	ExprBinaryNotIdentical(n *ExprBinaryNotIdentical)
	ExprBinaryPlus(n *ExprBinaryPlus)
	ExprBinaryPow(n *ExprBinaryPow)
	ExprBinaryShiftLeft(n *ExprBinaryShiftLeft)
	ExprBinaryShiftRight(n *ExprBinaryShiftRight)
	ExprBinarySmaller(n *ExprBinarySmaller)
	ExprBinarySmallerOrEqual(n *ExprBinarySmallerOrEqual)
	ExprBinarySpaceship(n *ExprBinarySpaceship)

	ExprCastArray(n *ExprCastArray)
	ExprCastBool(n *ExprCastBool)
	ExprCastDouble(n *ExprCastDouble)
	ExprCastInt(n *ExprCastInt)
	ExprCastObject(n *ExprCastObject)
	ExprCastString(n *ExprCastString)
	ExprCastUnset(n *ExprCastUnset)

	ScalarDnumber(n *ScalarDnumber)
	ScalarEncapsed(n *ScalarEncapsed)
	ScalarEncapsedStringPart(n *ScalarEncapsedStringPart)
	ScalarHeredoc(n *ScalarHeredoc)
	ScalarLnumber(n *ScalarLnumber)
	ScalarMagicConstant(n *ScalarMagicConstant)
	ScalarString(n *ScalarString)

	NameName(n *NameName)
	NameFullyQualified(n *NameFullyQualified)
	NameRelative(n *NameRelative)
	NameNamePart(n *NameNamePart)

	ParserAs(n *ParserAs)
	ParserNsSeparator(n *ParserNsSeparator)
	ParserBrackets(n *ParserBrackets)
}

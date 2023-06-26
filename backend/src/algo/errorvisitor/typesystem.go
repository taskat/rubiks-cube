package errorvisitor

type typeSystem struct {
	types map[string]iType
}

func newTypeSystem() *typeSystem {
	ts := &typeSystem{types: make(map[string]iType)}
	errorType := newBasicType("error")
	pieceType := newBasicType("piece")
	positionType := newBasicType("position")
	nodeType := newBasicType("node", pieceType, positionType)
	coordType := newBasicType("coord")
	pieceListType := newListType(pieceType)
	positionListType := newListType(positionType)
	nodeListType := newListType(nodeType)
	coordListType := newListType(coordType)
	errorType.addParents(nodeType, coordType, pieceListType, positionListType,
		nodeListType, coordListType)
	ts.addType(errorType)
	ts.addType(pieceType)
	ts.addType(positionType)
	ts.addType(nodeType)
	ts.addType(coordType)
	ts.addType(pieceListType)
	ts.addType(positionListType)
	ts.addType(nodeListType)
	ts.addType(coordListType)

	return ts
}

func (ts *typeSystem) addType(t iType) {
	ts.types[t.String()] = t
}

func (ts *typeSystem) getType(name string) iType {
	return ts.types[name]
}

package patch

type MoveOp struct {
	Path Pointer
	From Pointer
}

func (op MoveOp) Apply(doc interface{}) (interface{}, error) {
	value, err := FindOp{Path: op.From}.Apply(doc)
	if err != nil {
		return nil, err
	}

	doc, err = ReplaceOp{Path: op.Path, Value: value}.Apply(doc)
	if err != nil {
		return nil, err
	}
	return RemoveOp{Path: op.From}.Apply(doc)
}

package util

func MapJoinedRows[TRow, TParent, TChild any](
	rows []TRow,
	parentMapper func(TRow, []TChild) TParent,
	childMapper func(TRow) (parentID string, child TChild),
) []TParent {
	children := make([]TChild, 0, len(rows))
	var currentParentID string
	parents := make([]TParent, 0, len(rows))
	parentLastIndex := 0
	for index, row := range rows {
		parentID, child := childMapper(row)
		children = append(children, child)
		if index == 0 {
			currentParentID = parentID
		}
		if currentParentID != parentID {
			parent := parentMapper(rows[parentLastIndex], children[parentLastIndex:index])
			parents = append(parents, parent)
			currentParentID = parentID
			parentLastIndex = index
		}
	}
	if len(rows) == 0 {
		return parents
	}
	parent := parentMapper(rows[parentLastIndex], children[parentLastIndex:])
	parents = append(parents, parent)

	return parents
}

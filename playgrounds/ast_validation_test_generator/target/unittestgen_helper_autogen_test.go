package target

type EquivalencyType string

const (
	EquivalencyNone    EquivalencyType = "EquivalencyNone"
	EquivalencyValid   EquivalencyType = "EquivalencyValid"
	EquivalencyInvalid EquivalencyType = "EquivalencyInvalid"
)

type BoundaryType string

const (
	BoundaryNone       BoundaryType = "BoundaryNone"
	BoundaryLowerBelow BoundaryType = "BoundaryLowerBelow"
	BoundaryLower      BoundaryType = "BoundaryLower"
	BoundaryLowerAbove BoundaryType = "BoundaryLowerAbove"
	BoundaryUpperBelow BoundaryType = "BoundaryUpperBelow"
	BoundaryUpper      BoundaryType = "BoundaryUpper"
	BoundaryUpperAbove BoundaryType = "BoundaryUpperAbove"
)

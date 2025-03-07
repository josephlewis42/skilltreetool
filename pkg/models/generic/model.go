package generic

// Number of rows in a skill tree.
const NumRows = 10

// Total number of skills in a tree.
const NumSkills = 68

// SkillTree is a generic representation of a skill tree.
type SkillTree struct {
	// Title of the tree.
	Title string
	// Footer for the tree.
	Footer string
	// Skills in the tree. Missing skills are assumed to be blank.
	Skills []Skill
}

// A skill in the skill tree.
type Skill struct {
	Row  int
	Col  int
	Text string
}

type Converter interface {
	ToGeneric() *SkillTree
}

// ColsInRow returns the number of expected columns in a 0-indexed row.
// If row is outside the standard skill tree range, 0 is returned.
func ColsInRow(row int) int {
	switch {
	case row >= 0 && row <= 8:
		return 7
	case row == 9:
		return 5
	default:
		return 0
	}
}

// Converts an internal column to the physical column for layout.
func LayoutCol(row int, col int) int {
	switch {
	case row == 9:
		return col + 1
	default:
		return col
	}
}

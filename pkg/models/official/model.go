package official

import (
	"fmt"
	"iter"

	"github.com/josephlewis42/skilltreetool/pkg/models/generic"
	"sigs.k8s.io/yaml"
)

// SkillTree is the official Maker Skill Tree format from
// https://github.com/sjpiper145/MakerSkillTree/tree/main
type SkillTree struct {
	Title  string              `json:"title"`
	Footer string              `json:"footer"`
	Rows   map[string][]string `json:"row"`
}

func (tree *SkillTree) Skills() iter.Seq[generic.Skill] {
	return func(yield func(generic.Skill) bool) {
		for rowIdx := 0; rowIdx < 10; rowIdx++ {
			cols, ok := tree.Rows[fmt.Sprintf("%d", rowIdx)]
			if !ok {
				continue
			}

			for colIdx, text := range cols {
				skill := generic.Skill{
					Col:  colIdx,
					Row:  rowIdx,
					Text: text,
				}
				if !yield(skill) {
					return
				}
			}
		}
	}
}

var _ generic.Converter = (*SkillTree)(nil)

func NewFromYaml(yamlBytes []byte) (*SkillTree, error) {
	var out SkillTree

	if err := yaml.Unmarshal(yamlBytes, &out); err != nil {
		return nil, fmt.Errorf("decoding error: %w", err)
	}

	return &out, nil
}

func NewFromGeneric(input *generic.SkillTree) *SkillTree {
	var out SkillTree

	out.Title = input.Title
	out.Footer = input.Footer
	out.Rows = make(map[string][]string)
	for row := 0; row < generic.NumRows; row++ {
		out.Rows[fmt.Sprintf("%d", row)] = make([]string, generic.ColsInRow(row))
	}

	for _, skill := range input.Skills {
		out.Rows[fmt.Sprintf("%d", skill.Row)][skill.Col] = skill.Text
	}

	return &out
}

// ToGeneric implements generic.Converter.
func (tree *SkillTree) ToGeneric() *generic.SkillTree {
	var out generic.SkillTree
	out.Title = tree.Title
	out.Footer = tree.Footer

	for skill := range tree.Skills() {
		out.Skills = append(out.Skills, skill)
	}

	return &out
}

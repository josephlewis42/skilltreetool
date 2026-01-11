package commands

import (
	"bytes"
	"fmt"
	"math"
	"regexp"

	"github.com/hbollon/go-edlib"
	"github.com/josephlewis42/skilltreetool/pkg/models/generic"
)

type SkillTreeDiff struct {
	Added []string

	Changed []string

	Removed []string

	Moved []string
}

func (diff *SkillTreeDiff) ToMarkdown() string {
	out := &bytes.Buffer{}

	for _, section := range []struct {
		Name string
		Data []string
	}{
		{"Added", diff.Added},
		{"Changed", diff.Changed},
		{"Removed", diff.Removed},
		{"Moved", diff.Moved},
	} {

		fmt.Fprintf(out, "## %s\n\n", section.Name)

		if len(section.Data) == 0 {
			fmt.Fprintf(out, "_No change_\n")
		}

		for _, item := range section.Data {
			fmt.Fprintf(out, "- %q\n", item)
		}

		fmt.Fprintln(out)
	}

	return out.String()
}

var (
	whitespaceRegex = regexp.MustCompile(`\s+`)
)

func collapseWhitespace(in string) string {
	return whitespaceRegex.ReplaceAllString(in, " ")
}

// Diff creates a diff between two skill trees.
func Diff(before, after *generic.SkillTree) *SkillTreeDiff {
	var diff SkillTreeDiff

	beforeSkills := make(map[string]generic.RowCol)

	for _, skill := range before.Skills {
		cleanedText := collapseWhitespace(skill.Text)
		beforeSkills[cleanedText] = skill.RowCol()
	}

	needsMatch := make(map[string]generic.RowCol)

	for _, skill := range after.Skills {
		cleanedText := collapseWhitespace(skill.Text)

		// If skill is in beforeSkills, skip it
		if beforePos, ok := beforeSkills[cleanedText]; ok {
			delete(beforeSkills, cleanedText)
			if beforePos == skill.RowCol() {
				// Same position, no change.
				continue
			} else {
				diff.Moved = append(diff.Moved, cleanedText)
				continue
			}
		}

		needsMatch[cleanedText] = skill.RowCol()
	}

	var remainingKeys []string
	for key := range needsMatch {
		remainingKeys = append(remainingKeys, key)
	}

	for _, afterText := range remainingKeys {
		bestMatchDistance := math.MaxInt
		bestMatch := ""

		for beforeText := range beforeSkills {
			if distance := edlib.LevenshteinDistance(beforeText, afterText); distance < bestMatchDistance {
				bestMatch = beforeText
				bestMatchDistance = distance
			}
		}

		// Values here are best estimates.
		if bestMatch != "" && (bestMatchDistance < 5 || float32(bestMatchDistance) < .25*float32(len(afterText))) {
			delete(beforeSkills, bestMatch)
			delete(needsMatch, afterText)

			diff.Changed = append(diff.Changed, fmt.Sprintf("%q to %q", bestMatch, afterText))
		}
	}

	for key := range beforeSkills {
		diff.Removed = append(diff.Removed, key)
	}

	for key := range needsMatch {
		diff.Added = append(diff.Added, key)
	}

	return &diff
}

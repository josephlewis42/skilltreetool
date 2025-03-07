package skilltreegenerator

import (
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"net/url"
	"strings"

	"github.com/josephlewis42/skilltreetool/pkg/models/generic"
	"github.com/josephlewis42/skilltreetool/third_party/makerskilltreegenerator"
	"github.com/rivo/uniseg"
)

// This mapping is flipped from what you'd usually see
// visually.
var rowColToIndex = [][]int{
	{0, 9, 19, 29, 39, 49, 59},
	{1, 10, 20, 30, 40, 50, 60},
	{2, 11, 21, 31, 41, 51, 61},
	{3, 12, 22, 32, 42, 52, 62},
	{4, 13, 23, 33, 43, 53, 63},
	{5, 14, 24, 34, 44, 54, 64},
	{6, 15, 25, 35, 45, 55, 65},
	{7, 16, 26, 36, 46, 56, 66},
	{8, 17, 27, 37, 47, 57, 67},
	{18, 28, 38, 48, 58},
}

type RowCol struct {
	Row int
	Col int
}

var indexToRowCol = map[string]RowCol{}

func init() {
	for rowIdx, row := range rowColToIndex {
		for colIdx, index := range row {
			indexToRowCol[fmt.Sprintf("%d", index)] = RowCol{rowIdx, colIdx}
		}
	}
}

// SkillTree represents a skill tree from the generator
// https://github.com/schme16/MakerSkillTree-Generator
type SkillTree struct {
	Title   string `json:"title"`
	Credits string `json:"credits"`

	// Items have string keys 0-67, starting bottom left and going up each column
	// Items without an entry are blank.
	Items map[string]string `json:"items"`
}

// Asserts that SkillTree is a generic.Converter
var _ generic.Converter = (*SkillTree)(nil)

func NewFromSVG(svgBytes []byte) (*SkillTree, error) {
	type SkillTreeSVG struct {
		XMLName xml.Name `xml:"svg"`
		Json    *string  `xml:"json"`
	}

	var inner SkillTreeSVG
	if err := xml.Unmarshal(svgBytes, &inner); err != nil {
		return nil, fmt.Errorf("invalid XML: %w", err)
	}

	if inner.Json == nil {
		return nil, errors.New("SVG doesn't have an embedded <json> tag, was it generated with MakerSkillTree-Generator?")
	}

	// Data is wrapped: base64(urlencode(json))
	// https://github.com/schme16/MakerSkillTree-Generator/blob/d033178251adeabb3815c1223bffc32be8a8eb96/sys/js/main.js#L231C28-L231C32
	decodedBase64 := make([]byte, len(*inner.Json))
	if _, err := base64.StdEncoding.Decode(decodedBase64, []byte(*inner.Json)); err != nil {
		return nil, fmt.Errorf("<json> tag has invalid bas64: %w", err)
	}

	unescaped, err := url.PathUnescape(string(decodedBase64))
	if err != nil {
		return nil, fmt.Errorf("<json> tag has invalid URL Encoding: %w", err)
	}

	unescaped = strings.TrimFunc(unescaped, func(r rune) bool {
		return r == 0x00
	})

	var out SkillTree
	if err := json.Unmarshal([]byte(unescaped), &out); err != nil {
		return nil, fmt.Errorf("<json> tag has invalid JSON: %w", err)
	}

	return &out, nil
}

// NewFromGeneric converts a generic skill tree into a SkillTreeGenerator version.
func NewFromGeneric(tree *generic.SkillTree) *SkillTree {
	var out SkillTree

	out.Title = tree.Title
	out.Credits = tree.Footer
	out.Items = make(map[string]string)

	for _, skill := range tree.Skills {
		idx := rowColToIndex[skill.Row][skill.Col]
		out.Items[fmt.Sprintf("%d", idx)] = skill.Text
	}

	return &out
}

// ToGeneric implements generic.Converter
func (tree *SkillTree) ToGeneric() *generic.SkillTree {
	var out generic.SkillTree

	out.Title = tree.Title
	out.Footer = tree.Credits

	for idx, skillText := range tree.Items {
		rowCol := indexToRowCol[idx]

		out.Skills = append(out.Skills, generic.Skill{
			Row:  rowCol.Row,
			Col:  rowCol.Col,
			Text: skillText,
		})
	}

	return &out
}

// ToSVG converts the skill tree to an SVG format.
func (tree *SkillTree) ToSVG() (string, error) {

	jsonBytes, err := json.Marshal(tree)
	if err != nil {
		return "", fmt.Errorf("couldn't create embedded JSON: %w", err)
	}

	escaped := url.PathEscape(string(jsonBytes))
	embeddedJson := base64.StdEncoding.EncodeToString([]byte(escaped))

	var replacements []string

	setReplacement := func(key, value string) {
		replacements = append(replacements, key, value)
	}
	setReplacement("{{ json }}", html.EscapeString(embeddedJson))
	setReplacement("{{ title }}", html.EscapeString(tree.Title))
	setReplacement("{{ credits }}", html.EscapeString(tree.Credits))

	for i := 0; i < generic.NumSkills; i++ {
		skillIdx := fmt.Sprintf("%d", i)
		skill, ok := tree.Items[skillIdx]
		if !ok {
			skill = ""
		}

		setReplacement(fmt.Sprintf("{{ skill_%d }}", i), wrapToTextSpans(skill, skillIdx))
	}

	replacer := strings.NewReplacer(replacements...)

	// This could be done using the HTML template package, but for binary size it's easier to
	// use this.
	return replacer.Replace(makerskilltreegenerator.SVGTemplate), nil
}

var colXOffsets = []string{
	"135",
	"239",
	"341",
	"445",
	"547",
	"651",
	"754",
}

func wrapToTextSpans(str string, skillIdx string) string {

	rowCol := indexToRowCol[skillIdx]
	layoutCol := generic.LayoutCol(rowCol.Row, rowCol.Col)
	offset := colXOffsets[layoutCol]

	var svg string
	for _, line := range strings.Split(wrap(str), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		svg += fmt.Sprintf(
			`<tspan x="%s" dy="1em">%s</tspan>`,
			offset,
			html.EscapeString(line),
		)
	}

	return svg
}

var rowWidthsEn = []int{12, 18, 18, 18, 18, 12}

// TODO: use a real layout engine like:
// https://github.com/go-text/typesetting
func wrap(str string) string {
	var built string

	var row int
	remainingWidth := rowWidthsEn[0]
	advanceRow := func() bool {
		row += 1
		if row >= len(rowWidthsEn) {
			built += "…"
			return true
		}
		built += "\n"
		remainingWidth = rowWidthsEn[row]
		return false
	}

	state := -1
	var word string
	for len(str) > 0 {
		word, str, state = uniseg.FirstWordInString(str, state)
		approxDisplayWidth := uniseg.StringWidth(word)

		if approxDisplayWidth > remainingWidth {
			if outOfSpace := advanceRow(); outOfSpace {
				return built
			}
		}

		remainingWidth -= approxDisplayWidth
		built += word
	}

	return built
}

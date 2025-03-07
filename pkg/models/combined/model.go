package combined

import (
	"fmt"

	"github.com/josephlewis42/skilltreetool/pkg/models/generic"
)

type CombinedFormat struct {
	Language string                         `json:"language"`
	Title    TranslationString              `json:"title"`
	Footer   TranslationString              `json:"footer"`
	Row      map[string][]TranslationString `json:"row"`
}

func NewCombinedFormat(original generic.SkillTree, languageCode string) *CombinedFormat {
	var out CombinedFormat
	out.Language = languageCode
	out.Title = TranslationString{
		Original: original.Title,
	}
	out.Footer = TranslationString{
		Original: original.Footer,
	}

	for row := 0; row < generic.NumRows; row++ {
		out.Row[fmt.Sprintf("%d", row)] = make([]TranslationString, generic.ColsInRow(row))
	}

	for _, skill := range original.Skills {
		out.Row[fmt.Sprintf("%d", skill.Row)][skill.Col].Original = skill.Text
	}

	return &out
}

// AddTranslation merges in translations for title, footer, and all skills.
func (combined *CombinedFormat) AddTranslation(languageCode string, toMerge generic.SkillTree) {
	combined.Title.AddTranslation(languageCode, toMerge.Title)
	combined.Footer.AddTranslation(languageCode, toMerge.Footer)

	for _, skill := range toMerge.Skills {
		combined.Row[fmt.Sprintf("%d", skill.Row)][skill.Col].AddTranslation(languageCode, skill.Text)
	}
}

type TranslationString struct {
	Original     string            `json:"original"`
	Note         string            `json:"note,omitempty"`
	Translations map[string]string `json:"translations"`
}

// AddTranslation adds or overwrites a translation for the given language code.
// If the translation is blank, nothing is changed.
func (ts *TranslationString) AddTranslation(lang, translation string) {
	if translation == "" {
		return
	}

	if ts.Translations == nil {
		ts.Translations = make(map[string]string)
	}

	ts.Translations[lang] = translation
}

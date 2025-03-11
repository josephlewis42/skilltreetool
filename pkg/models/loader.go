package models

import (
	"errors"
	"fmt"
	"os"

	"github.com/josephlewis42/skilltreetool/pkg/models/generic"
	"github.com/josephlewis42/skilltreetool/pkg/models/official"
	"github.com/josephlewis42/skilltreetool/pkg/models/skilltreegenerator"
)

func LoadFromFile(name string) (*generic.SkillTree, error) {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("couldn't read file %s: %w", name, err)
	}

	return LoadFromBytes(bytes)
}

func LoadFromString(data string) (*generic.SkillTree, error) {
	return LoadFromBytes([]byte(data))
}

func LoadFromBytes(data []byte) (*generic.SkillTree, error) {
	var errs []error

	svg, svgErr := skilltreegenerator.NewFromSVG(data)
	if svgErr != nil {
		errs = append(errs, fmt.Errorf("svg error: %w", svgErr))
	} else {
		return svg.ToGeneric(), nil
	}

	yaml, yamlErr := official.NewFromYaml(data)
	if yamlErr != nil {
		errs = append(errs, fmt.Errorf("yaml error: %w", yamlErr))
	} else {
		return yaml.ToGeneric(), nil
	}

	return nil, fmt.Errorf("couldn't decode file as YAML or SVG:\n%w", errors.Join(errs...))
}

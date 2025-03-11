package commands

import (
	"fmt"
	"io"

	"github.com/josephlewis42/skilltreetool/pkg/models/official"
	"github.com/josephlewis42/skilltreetool/pkg/models/skilltreegenerator"
	"sigs.k8s.io/yaml"
)

func SVG2Yaml(input []byte, output io.Writer) error {
	decoded, err := skilltreegenerator.NewFromSVG(input)
	if err != nil {
		return fmt.Errorf("couldn't decode embedded data in SVG: %w", err)
	}

	officialTree := official.NewFromGeneric(decoded.ToGeneric())

	converted, err := yaml.Marshal(officialTree)
	if err != nil {
		return fmt.Errorf("couldn't convert to YAML: %w", err)
	}

	if _, err := output.Write(converted); err != nil {
		return fmt.Errorf("couldn't write output: %w", err)
	}

	return nil
}

package commands

import (
	"fmt"
	"io"

	"github.com/josephlewis42/skilltreetool/pkg/models/official"
	"github.com/josephlewis42/skilltreetool/pkg/models/skilltreegenerator"
)

func Yaml2SVG(input []byte, output io.Writer) error {
	official, err := official.NewFromYaml(input)
	if err != nil {
		return fmt.Errorf("couldn't decode embedded data in SVG: %w", err)
	}

	svg := skilltreegenerator.NewFromGeneric(official.ToGeneric())

	data, err := svg.ToSVG()
	if err != nil {
		return fmt.Errorf("couldn't generate SVG: %w", err)
	}

	if _, err := fmt.Fprintln(output, data); err != nil {
		return err
	}

	return nil
}

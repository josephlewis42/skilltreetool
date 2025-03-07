package browser

import (
	"fmt"

	"github.com/josephlewis42/skilltreetool/pkg/models/official"
	"github.com/josephlewis42/skilltreetool/pkg/models/skilltreegenerator"
	"sigs.k8s.io/yaml"
)

func SVG2Yaml(svg string) string {
	decoded, err := skilltreegenerator.NewFromSVG([]byte(svg))
	if err != nil {
		return fmt.Sprintf("couldn't decode embedded data in SVG: %s", err)
	}

	converted := official.NewFromGeneric(decoded.ToGeneric())

	yamlBytes, err := yaml.Marshal(converted)
	if err != nil {
		return fmt.Sprintf("couldn't convert to YAML: %w", err)
	}

	return string(yamlBytes)
}

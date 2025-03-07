package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/josephlewis42/skilltreetool/pkg/models/official"
	"github.com/josephlewis42/skilltreetool/pkg/models/skilltreegenerator"
	"github.com/urfave/cli/v3"
	"sigs.k8s.io/yaml"
)

var version = "0.0.0"

func svg2yaml() *cli.Command {
	var path string
	var outPath string

	return &cli.Command{
		Name:      "svg2yaml",
		Usage:     "Convert a SkillTreeGenerator SVG file to the Maker Skill Tree YAML format.",
		ArgsUsage: "SVG_FILE",
		Arguments: []cli.Argument{
			&cli.StringArg{Name: "SVG_FILE", Destination: &path, UsageText: "Path to the SkillTreeGenerator SVG", Min: 1, Max: 1},
			&cli.StringArg{Name: "YAML_FILE", Destination: &outPath, UsageText: "Path to write to or - for stdout", Min: 1, Max: 1},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fileBytes, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("couldn't read file %s: %w", path, err)
			}

			decoded, err := skilltreegenerator.NewFromSVG(fileBytes)
			if err != nil {
				return fmt.Errorf("couldn't decode embedded data in SVG: %w", err)
			}

			officialTree := official.NewFromGeneric(decoded.ToGeneric())

			converted, err := yaml.Marshal(officialTree)
			if err != nil {
				return fmt.Errorf("couldn't convert to YAML: %w", err)
			}

			writer := cmd.Writer
			if path != "-" {
				fd, err := os.OpenFile(path, os.O_CREATE, 0700)
				if err != nil {
					return fmt.Errorf("couldn't open file %q: %w", path, err)
				}
				defer fd.Close()
				writer = fd
			}

			fmt.Fprintln(writer, string(converted))
			return nil
		},
	}
}

func yaml2svg() *cli.Command {
	var path string

	return &cli.Command{
		Name:      "yaml2svg",
		Usage:     "Convert a Maker Skill Tree YAML file to an SVG that can be edited in SkillTreeGenerator.",
		ArgsUsage: "YAML_FILE",
		Arguments: []cli.Argument{
			&cli.StringArg{Name: "YAML_FILE", Destination: &path, UsageText: "Path to the Maker Skill Tree YAML file", Max: 1},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fileBytes, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("couldn't read file %s: %w", path, err)
			}

			official, err := official.NewFromYaml(fileBytes)
			if err != nil {
				return fmt.Errorf("couldn't decode embedded data in SVG: %w", err)
			}

			svg := skilltreegenerator.NewFromGeneric(official.ToGeneric())

			data, err := svg.ToSVG()
			if err != nil {
				return fmt.Errorf("couldn't generate SVG: %w", err)
			}

			fmt.Fprintln(cmd.Writer, data)
			return nil
		},
	}
}

func main() {
	cmd := &cli.Command{
		Usage:   "Convert between Maker Skill Tree formats",
		Version: version,
		Commands: []*cli.Command{
			svg2yaml(),
			yaml2svg(),

			// TODO: Add the following features
			// diff -- difference between old and new in a KeepAChangelog format
			// sync -- sync all the files in a directory (translation, SVG, YAML)
			// TODO: Figure out how to handle versioning of SVGs
			// generateSVGs
			// generatePDFs
			// validate/lint/check
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

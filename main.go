package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/josephlewis42/skilltreetool/pkg/commands"
	"github.com/josephlewis42/skilltreetool/pkg/models"
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
		ArgsUsage: "SVG_FILE_IN YAML_FILE_OUT",
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
			if outPath != "-" {
				fd, err := os.Create(outPath)
				if err != nil {
					return fmt.Errorf("couldn't create file %q: %w", outPath, err)
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
	var outPath string

	return &cli.Command{
		Name:      "yaml2svg",
		Usage:     "Convert a Maker Skill Tree YAML file to an SVG that can be edited in SkillTreeGenerator.",
		ArgsUsage: "YAML_FILE_IN SVG_FILE_OUT",
		Arguments: []cli.Argument{
			&cli.StringArg{Name: "YAML_FILE", Destination: &path, UsageText: "Path to the Maker Skill Tree YAML file", Min: 1, Max: 1},
			&cli.StringArg{Name: "SVG_FILE", Destination: &outPath, UsageText: "Path to write to or - for stdout", Min: 1, Max: 1},
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

			writer := cmd.Writer
			if outPath != "-" {
				fd, err := os.Create(outPath)
				if err != nil {
					return fmt.Errorf("couldn't create file %q: %w", outPath, err)
				}
				defer fd.Close()
				writer = fd
			}

			fmt.Fprintln(writer, data)
			return nil
		},
	}
}

func diff() *cli.Command {
	var beforePath string
	var afterPath string

	return &cli.Command{
		Name:      "diff",
		Usage:     "Create a markdown diff between two skill tree files.",
		ArgsUsage: "BEFORE AFTER",
		Arguments: []cli.Argument{
			&cli.StringArg{Name: "BEFORE", Destination: &beforePath, UsageText: "The original file", Min: 1, Max: 1},
			&cli.StringArg{Name: "AFTER", Destination: &afterPath, UsageText: "The new file", Min: 1, Max: 1},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			before, err := models.LoadFromFile(beforePath)
			if err != nil {
				return fmt.Errorf("couldn't read file %s: %w", beforePath, err)
			}

			after, err := models.LoadFromFile(afterPath)
			if err != nil {
				return fmt.Errorf("couldn't read file %s: %w", afterPath, err)
			}

			diff := commands.Diff(before, after)

			fmt.Fprintln(cmd.Writer, diff.ToMarkdown())
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
			diff(),

			// TODO: Add the following features
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

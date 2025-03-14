# SkillTreeTool

SkillTreeTool provides utilities for working with collections of the various [Maker Skill Tree](https://github.com/sjpiper145/MakerSkillTree/tree/main) formats.

## Why SkillTreeTool?

SkillTreeTool reduces some of the effort required to maintain a set of skill trees by
automating common tasks so maintainers can focus on community engagement, content, and quality.

## Installing

Install using `go`:

```bash
go install github.com/josephlewis42/skilltreetool@latest
```

Or, build the tool locally:

```bash
# Clone the repository
git clone github.com/josephlewis42/skilltreetool && cd skilltreetool

# Build the tool
make

# Run the tool
./build/skilltreetool
```

## Using the tool

Examples:

```sh
# Convert YAML to SVG format:
skilltreetool yaml2svg input.yaml output.svg

# Convert SVG to YAML format:
skilltreetool svg2yaml input.svg output.yaml

# Convert SVG to YAML, output directly in the console
skilltreetool svg2yaml input.svg -

# Diff two trees and output a Keep A Changelog style changelog
skilltreetool diff original.svg new.svg
```

## Scope and future

The goals of this tool are to fill a similar role to a traditional programming language toolchain:

* Convert skill trees into different formats.
* Check trees for errors.
* Compile them for publishing e.g. to GitHub pages.

Because not everyone is comfortable with CLI tools, the tool can also be 
compiled to WASM to run in the browser.

## License

Code licensed under the Apache 2.0 License: [LICENSE](LICENSE)

Resources in [third_party](third_party) are under their own licenses.
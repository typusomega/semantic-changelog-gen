[![GoDoc](https://godoc.org/github.com/typusomega/semantic-changelog-gen?status.svg)](http://godoc.org/github.com/typusomega/semantic-changelog-gen)
[![Go Report](https://goreportcard.com/badge/github.com/typusomega/semantic-changelog-gen)](https://goreportcard.com/report/github.com/typusomega/semantic-changelog-gen)

# Semantic Changelog Gen

## How to install
If you have [Go](https://golang.org/) installed, it's as simple as running:

```bash
go get github.com/typusomega/semantic-changelog-gen/cmd/semantic-changelog-gen
```

## Usage

To generate your changelog simply navigate into the git repository and run:

```bash
semantic-changelog-gen generate
```

You can also specify the repository and the output file like this:

```bash
semantic-changelog-gen generate \
   --dir /path/to/repo \
   --out /path/to/changelog.md
```

## Contribution

[Create a new issue](https://github.com/typusomega/semantic-changelog-gen/issues/new) if you want to:
- have a new feature
- report a bug
- something else

[Pull requests](https://github.com/typusomega/semantic-changelog-gen/compare) are welcome

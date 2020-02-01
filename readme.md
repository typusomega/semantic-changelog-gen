[![GoDoc](https://godoc.org/github.com/typusomega/semantic-changelog-gen/pkg?status.svg)](http://godoc.org/github.com/typusomega/semantic-changelog-gen/pkg)
[![Go Report](https://goreportcard.com/badge/github.com/typusomega/semantic-changelog-gen)](https://goreportcard.com/report/github.com/typusomega/semantic-changelog-gen)

# Semantic Changelog Gen

A lightweight and extensible changelog generator working with semantic commits.
This generator sticks to the rules defined in [Karma Git Commit Msg](http://karma-runner.github.io/4.0/dev/git-commit-msg.html).

## How to install
If you have [Go](https://golang.org/) installed, it's as simple as running:

```bash
go get github.com/typusomega/semantic-changelog-gen
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

If you are tired of the default markdown format used by the generator, feel free to use your own style:

```bash
semantic-changelog-gen generate \
   --dir /path/to/repo \
   --out /path/to/changelog.vnext \
   --format custom \
   --template /path/to/your/go.tpl
```

## Contribution

[Create a new issue](https://github.com/typusomega/semantic-changelog-gen/issues/new) if you want to:
- have a new feature
- report a bug
- something else

[Pull requests](https://github.com/typusomega/semantic-changelog-gen/compare) are welcome

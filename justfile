default:
    @just --list

changelog:
    git cliff --bump -c cliff.toml > docs/CHANGELOG.md

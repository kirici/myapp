default:
    @just --list

changelog:
    git cliff --bump -c cliff.toml > docs/CHANGELOG.md

up:
    podman compose up -d

down:
    podman compose down --volumes

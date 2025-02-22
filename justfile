container := ```
    if command -v docker >/dev/null 2>&1 ; then
        echo "docker"
    else
        echo "podman"
    fi
    ```

default:
    @just --list

up:
    {{container}} compose up -d

up-all:
    {{container}} compose --profile test up -d

down:
    {{container}} compose --profile test down --volumes

build:
    {{container}} compose build

refresh: build && up

test:
    {{container}} compose --profile test up -d k6

release:
    git cliff -c cliff.toml --bump > docs/CHANGELOG.md
    git add docs/CHANGELOG.md
    git commit -S -m "release: add changelog for $(git-cliff --bumped-version)"
    git tag $(git-cliff --bumped-version) -m "Release $(git-cliff --bumped-version)"

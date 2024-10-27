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

changelog:
    git cliff --bump -c cliff.toml > docs/CHANGELOG.md

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

down:
    {{container}} compose down --volumes

build:
    {{container}} compose build

refresh: build && up

restart-load:
    {{container}} compose restart k6

changelog:
    git cliff --bump -c cliff.toml > docs/CHANGELOG.md

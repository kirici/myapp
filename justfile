container := ```
    if command -v docker >/dev/null 2>&1 ; then
        echo "docker"
    else
        echo "podman"
    fi
    ```

default:
    @just --list

changelog:
    git cliff --bump -c cliff.toml > docs/CHANGELOG.md

up:
    {{container}} compose up -d

down:
    {{container}} compose down --volumes

refresh:
    {{container}} compose up --build --force-recreate -d

load-low:
    vegeta attack -rate=2000/1s -duration=600s -targets targets.txt &>/dev/null

load-high:
    vegeta attack -rate=9001/1s -duration=300s -targets targets.txt &>/dev/null

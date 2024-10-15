default:
    @just --list

changelog:
    git cliff --bump -c cliff.toml > docs/CHANGELOG.md

up: fix-grafana
    podman compose up -d

down:
    podman compose down --volumes

refresh: fix-grafana
    podman compose up --build --force-recreate -d

vegeta:
    vegeta attack -rate=2000/1s -duration=600s -targets targets.txt &>/dev/null

ddos:
    vegeta attack -rate=9001/1s -duration=300s -targets targets.txt &>/dev/null

fix-grafana:
    chmod 664 ./grafana.db

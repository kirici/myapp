{{/* Platform-side tweaks for the local kind cluster; nothing here belongs in the Score files. */}}
{{ range $i, $m := .Manifests }}
{{ $wl := dig "metadata" "annotations" "k8s.score.dev/workload-name" "" $m }}
{{/* Expose workload Services via cloud-provider-kind */}}
{{ if and (eq $m.kind "Service") (has $wl (list "web" "prometheus" "grafana")) }}
- op: set
  path: {{ $i }}.spec.type
  value: LoadBalancer
{{ end }}
{{/* Label k6's objects as the "test" profile so `up` can skip them with a selector */}}
{{ if or (eq $wl "k6") (hasPrefix "k6-k6-file-" $m.metadata.name) }}
- op: set
  path: {{ $i }}.metadata.labels.profile
  value: test
{{ end }}
{{/* score-k8s only emits Deployments/StatefulSets; a load test should run once */}}
{{ if and (eq $m.kind "Deployment") (eq $wl "k6") }}
- op: set
  path: {{ $i }}.apiVersion
  value: batch/v1
- op: set
  path: {{ $i }}.kind
  value: Job
- op: set
  path: {{ $i }}.spec.backoffLimit
  value: 0
- op: set
  path: {{ $i }}.spec.template.spec.restartPolicy
  value: Never
- op: delete
  path: {{ $i }}.spec.selector
- op: delete
  path: {{ $i }}.spec.strategy
{{ end }}
{{ end }}

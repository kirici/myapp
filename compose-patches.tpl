{{/* Platform-side tweaks for local compose; nothing here belongs in the Score files. */}}
{{ range $name, $svc := .Compose.services }}
{{/* Keep the load test out of a plain `compose up`, like the original `test` profile */}}
{{ if eq (dig "annotations" "compose.score.dev/workload-name" "" $svc) "k6" }}
- op: set
  path: services.{{ $name }}.profiles
  value: [test]
{{ end }}
{{/* Relabel bind mounts for rootless podman on SELinux hosts (same as the original :z) */}}
{{ range $i, $v := (dig "volumes" list $svc) }}
{{ if eq $v.type "bind" }}
- op: set
  path: services.{{ $name }}.volumes.{{ $i }}.bind.selinux
  value: z
{{ end }}
{{ end }}
{{ end }}

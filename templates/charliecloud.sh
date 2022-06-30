#!/bin/sh

set -ex

th started {{.User.JobId}}

# --------------------------------------------

IMAGE={{.Cluster.Path}}{{.Cluster.LocalData}}{{.User.JobId}}
TMPDIR={{.Cluster.Path}}{{.Cluster.LocalScratch}}{{.User.JobId}}
ARTIFACTS={{.Cluster.Path}}{{.Cluster.GlobalScratch}}{{.User.JobId}}

# srun
ch-convert -i docker -o dir {{.User.Image}} $IMAGE

{{range .User.PreScript -}}
    ch-run $IMAGE -- {{. -}}
{{end}}

# srun
ch-run \
    --bind=$TMPDIR:{{.User.TmpDir}} \
    {{range $src, $dst := .User.Artifacts -}}
        --bind=$ARTIFACTS/{{$src}}:{{$dst}} \
    {{end -}}
    $IMAGE -- {{.User.JobScript}}

{{range .User.PostScript -}}
    ch-run $IMAGE -- {{. -}}
{{end}}

# --------------------------------------------

th completed {{.User.JobId}}

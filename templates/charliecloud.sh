#!/bin/sh

set -ex

JOB_ID=$1

th started $JOB_ID

# --------------------------------------------

# unpack image
IMAGE={{.Cluster.Path}}{{.Cluster.LocalData}}$JOB_ID
# tmp store results
TMPDIR={{.Cluster.Path}}{{.Cluster.LocalScratch}}$JOB_ID
# artifacts to be uploaded to cloud
ARTIFACTS={{.Cluster.Path}}{{.Cluster.GlobalScratch}}$JOB_ID

# srun
ch-convert -i docker -o dir {{.Image}} $IMAGE

{{range .User.PreScripts -}}
    ch-run $IMAGE -- {{.}}
{{end}}

ch-run \
    --bind=$TMPDIR:{{.User.TmpDir}} \
    {{range $src, $dst := .User.Artifacts -}}
        --bind=$ARTIFACTS/{{$src}}:{{$dst}} \
    {{end -}}
    $IMAGE -- {{.User.JobScript}}

{{range .User.PostScripts -}}
    ch-run $IMAGE -- {{.}}
{{end}}

# --------------------------------------------

th completed $JOB_ID

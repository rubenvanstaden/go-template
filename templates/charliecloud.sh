#!/bin/sh

set -ex

JOB_ID=$1

th started $JOB_ID

# --------------------------------------------

MASTER_IMAGE={{.Master.Path}}{{.Master.Name}}{{.Master.TmpDir}}$JOB_ID

ch-convert -i docker -o dir {{.Master.Image}} $MASTER_IMAGE

{{range .Master.PreScripts -}}
    ch-run $MASTER_IMAGE -- {{.}}
{{end}}

{{range .Nodes -}}

    workspace={{.Path}}{{.Name}}{{.TmpDir}}$JOB_ID

    IMAGE=$workspace/image
    TMP=$workspace/tmp
    ARTIFACTS=$workspace/artifacts

    ch-convert -i docker -o dir {{.Image}} $IMAGE

    ch-run \
        --bind=$TMP:{{.Scratch}} \
        {{range $src, $dst := .Artifacts -}}
            --bind=$ARTIFACTS/{{$src}}:{{$dst}} \
        {{end -}}
        $IMAGE -- {{.Script}}

{{end}}

{{range .Master.PostScripts -}}
    ch-run $MASTER_IMAGE -- {{.}}
{{end}}

# --------------------------------------------

th completed $JOB_ID

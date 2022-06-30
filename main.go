package main

import (
	"os"
	"log"
	"text/template"
)

var temp *template.Template

type ClusterConfig struct {
    Path          string
    GlobalData    string
    GlobalScratch string
    LocalData     string
    LocalScratch  string
}

type UserConfig struct {
    JobId      string
    Image      string
    TmpDir     string
    Artifacts  map[string]string
    PreScript  []string
    JobScript  string
    PostScript []string
}

type Config struct {
    Cluster ClusterConfig
    User    UserConfig
}

func init() {
    temp = template.Must(template.ParseFiles("templates/charliecloud.sh"))
}

func main() {

    // 1. one job equals one image
    // 2. one job equals one srun script (job script)

    // $GLOBAL_SCRATCH/<JOB_ID>/cfd-results  # store artifacts from local scratch
    // $GLOBAL_SCRATCH/<JOB_ID>/em-results   # store artifacts
    // $LOCAL_DATA/<JOB_ID>                  # unpack charlie cloud image
    // $LOCAL_SCRATCH/<JOB_ID>               # store temp job run data

    variables := map[string]string{
        "TMPDIR": "/var/tmp/scratch",
    }

    artifacts := map[string]string{
        "cfd-results": "/var/tmp/cfd-results",
        "em-results": "/var/tmp/em-results",
    }

    cluster := ClusterConfig{
        Path:          "/var/tmp/cluster/",
        GlobalData:    "share/thunderhead/data/",
        GlobalScratch: "share/thunderhead/scratch/",
        LocalData:     "local/thunderhead/data/",
        LocalScratch:  "local/thunderhead/scratch/tmp/",
    }

    user := UserConfig{
        JobId: "12345667789",
        Image: "mongo:latest",
        TmpDir: variables["TMPDIR"],
        Artifacts: artifacts,
        PreScript: []string{"echo 'run pre-script...'"},
        JobScript: "hello/hello.sh",
        PostScript: []string{"echo 'run post-script...'"},
    }

    config := Config{
        Cluster: cluster,
        User: user,
    }

    err := temp.Execute(os.Stdout, config)
    if err != nil {
        log.Fatalln(err)
    }
}

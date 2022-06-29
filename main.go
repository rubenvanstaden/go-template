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
    Path        string
    Name        string
    Image       string
    TmpDir      string
    Artifacts   map[string]string
    PreScripts  []string
    JobScript   string
    PostScripts []string
}

type Config struct {
    Cluster ClusterConfig
    User    UserConfig
}

func init() {
    temp = template.Must(template.ParseFiles("templates/charliecloud.sh"))
}

func main() {

    variables := map[string]string{
        "TMPDIR": "/var/tmp/scratch",
    }

    // artifacts are automatically mount to SCRATCH
    artifacts := map[string]string{
        "cfd-results": "/var/tmp/cfd-results",
        "em-results": "/var/tmp/em-results",
    }

    cluster := ClusterConfig{
        Path: "/var/tmp/cluster",

        GlobalData: "/share/thunderhead/data",

        // - $GLOBAL_SCRATCH/<JOB_ID>/cfd-results
        // - $GLOBAL_SCRATCH/<JOB_ID>/em-results
        GlobalScratch: "/share/thunderhead/scratch",

        // - $LOCAL_DATA/<JOB_ID>
        LocalData:  "/local/thunderhead/data",

        // - $LOCAL_SCRATCH/<JOB_ID>
        LocalScratch:  "/local/thunderhead/scratch",
    }

    user := UserConfig{
        Name: "node",
        Image: "mongo:latest",
        TmpDir: variables["TMPDIR"],
        Artifacts: artifacts,
        PreScripts: []string{"echo 'run pre-script...'"},
        JobScript: "hello/hello.sh",
        PostScripts: []string{"echo 'run post-script...'"},
    }

    config := Config{
        Cluster: cluster,
        User:    user,
    }

    err := temp.Execute(os.Stdout, config)
    if err != nil {
        log.Fatalln(err)
    }
}

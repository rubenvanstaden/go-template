package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

type MasterNode struct {
    Path        string
    Name        string
    Image       string
    TmpDir      string
    PreScripts  []string
    PostScripts []string
}

type ComputeNode struct {
    Path      string
    Name      string
    Image     string
    TmpDir    string
    Script    string
    Scratch   string
    Artifacts map[string]string
}

type Data struct {
    Master MasterNode
    Nodes  []ComputeNode
}

func init() {
    temp = template.Must(template.ParseFiles("templates/charliecloud.sh"))
}

func main() {

    master := MasterNode{
        Path: "/var/tmp/cluster/",
        Name: "master",
        Image: "mongo:latest",
        TmpDir: "/scratch/local/",
        PreScripts: []string{"echo 'run pre-script...'"},
        PostScripts: []string{"echo 'run post-script...'"},
    }

    artifacts := map[string]string{
        "cfd-results": "/var/tmp/cfd-results",
        "em-results": "/var/tmp/em-results",
    }

    node1 := ComputeNode{
        Path: "/var/tmp/cluster/",
        Name: "node1",
        Image: "mongo:latest",
        TmpDir: "/scratch/local/",
        Script: "hello/hello.sh",
        Scratch: "/var/tmp/scratch",
        Artifacts: artifacts,
    }

    node2 := ComputeNode{
        Path: "/var/tmp/cluster/",
        Name: "node2",
        Image: "mongo:latest",
        TmpDir: "/scratch/local/",
        Script: "hello/hello.sh",
        Scratch: "/var/tmp/scratch",
        Artifacts: artifacts,
    }

    data := Data{
        Master: master,
        Nodes:  []ComputeNode{node1, node2},
    }

    err := temp.Execute(os.Stdout, data)
    if err != nil {
        log.Fatalln(err)
    }
}

package main


import (
    "os"
    "github.com/codegangsta/cli"
    "os/exec"
    "strings"
    path "path/filepath"
)

var HOME string

func main() {
    app := cli.NewApp()
    app.Name = "wg"
    app.Usage = "wg is a git tool for downloads git repo project"
    app.Action = func(c *cli.Context) {
        println("Welcome to wg.")
    }
    app.Commands = []cli.Command{
        {
            Name:      "test",
            Aliases:     []string{"t"},
            Usage:     "wg test",
            Action:  func(c *cli.Context) {
                test(c)
                home()
            },
        },
        {
            Name:      "get",
            Aliases:     []string{"g"},
            Usage:     "wg get",
            Action:  func(c *cli.Context) {
                home()
                get(c)
            },
        },
    }

    app.Run(os.Args)
}

func test(c *cli.Context) {
//    f, err := exec.LookPath("git")
//    if err != nil {
//        println("not install git")
//    }
//    println(f)

//    cmd :=exec.Command("git", "help")
//    out, err := cmd.Output()
//    if err != nil {
//        println(err)
//    }
//    println(string(out))
//    cd("~/workspace")

//    cmd := exec.Command("mkdir", "-p", "/Users/bluesand/workspace/github.com/well")
//    out,err := cmd.Output()
//    if err != nil {
//        println(err)
//    }
//    println(string(out))

//    os.Mkdir("/Users/bluesand/workspace/github.com/well", 0755)


    println()


}


func home() {
//    well_home := os.Getenv("WELL_HOME")
//    if well_home != "" {
//        println("["+well_home+"]")
//    } else {
//        os.Setenv("WELL_HOME", "~/wellspace")
//    }
//
//    well_home = os.Getenv("WELL_HOME")
//    println("["+well_home+"]")
    HOME = os.Getenv("HOME")

}

func pwd() {

    cmd := exec.Command("pwd")
    out, err := cmd.Output()
    if err != nil {
        println(err)
    }
    println("pwd:" + string(out))
}

func cd(path string) {
    println("cd:"+ path)
    exec.Command("cd", path)
}

func mkdir(path string) {
    println("mkdir:" + path)
    exec.Command("mkdir", "-p", path)
}

func git_get(host, user, repo string) {

    pwd()

    userpath := HOME + string(path.Separator) + "workspace"+ string(path.Separator)+ host + string(path.Separator) + user
    mkdir(userpath)

    url := "https://" + host + string(path.Separator) + user + string(path.Separator) + repo + ".git"

    local := userpath + string(path.Separator) + repo

    cmd := exec.Command("git", "clone", url, local)
    out, err := cmd.Output()
    if err != nil {
        println(err)
    }
    println(string(out))
}


func get(c *cli.Context) {
    url := c.Args().First()
    urls := strings.Split(url, "/")
    if len(urls)  < 3 {
        println("sub params wrong")
    } else {
        host := urls[0]
        user := urls[1]
        repo := urls[2]
        println( host + ":" + user + ":" + repo)
        git_get(host, user,repo)
    }
}

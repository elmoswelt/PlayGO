package git

import (
	"github/elmoswelt/Geekon-2016/Source/model"
	"os/exec"
	"bytes"
	"fmt"
)

type Git struct  {

	Project model.Project
}

const path = "/tmp/localizationServer/projects/"
const cmd_git = "git"
const cmd_rm = "rm"


func (g *Git)Clone() error {

	fmt.Print("[Clone:] \t " + g.Project.RepoURL)

	_, err := g.executeCommand(cmd_git, []string{"clone",g.Project.RepoURL, g.path()})

	return err
}


func (g *Git)Diff() error  {

	cmd := "git -C " + g.path() + " diff master..origin/" + g.Project.BranchName + ` -U0 | egrep "^\+\""`

	out, err := exec.Command("bash","-c", cmd).Output()
	if err != nil {
		fmt.Println("Failed to execute command: %s", cmd)
		return err
	}

	fmt.Println(string(out))

	return err
}


func (g *Git)Pull() error  {

	fmt.Print("[Pull:] \t" + g.Project.BranchName)

	_, err := g.executeCommand(cmd_git, []string{"-C", g.path(), "pull"})

	return err
}


func (g *Git)Cleanup() error  {

	fmt.Print("[Cleanup:] \t" + g.path())

	_, err := g.executeCommand(cmd_rm, []string{"-rf", g.path()})

	return err
}


func (g *Git)executeCommand(cmdName string, cmdArgs []string) (string, error) {

	cmd := exec.Command(cmdName, cmdArgs...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {

		fmt.Println("\t [ERROR]" + fmt.Sprint(err) + ": " + stderr.String() + "\n")
		return "", err
	}

	fmt.Println("\t [OK] ")

	return out.String(), nil
}


func (g *Git)path() string  {

	return path + g.Project.Name + "_" + g.Project.Platform + "_" + g.Project.ProjectID
}

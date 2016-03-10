package routes

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"os/exec"
	"bytes"
	"github/elmoswelt/Geekon-2016/Source/model"
	"github/elmoswelt/Geekon-2016/Source/git"
)

var fileCounter int64 = 1
var folderCounter int64 = 1

const path = "/tmp/localizationServer/diff"


func Diff(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Start processing...")

	projects := model.ProjectList()
	project := projects[0]
	fmt.Println(project)
	fmt.Println("\n")

	g := git.Git{Project:project}

	g.Cleanup()
	g.Clone()
	g.Diff()
}


func diffStringFromUrl(url string) (string, error) {

	response , err := http.Get(url)

	if err != nil {
		return "", err
	} else {

		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return  "", err
		} else {
			return string(contents), err
		}
	}
}

func saveDiffStringToFile(str string) (error) {

	err := os.MkdirAll(path, 0777)

	if err != nil {
		fmt.Println("Error on folder creation", err)
		return err
	}

	f, err := os.Create(path + "/" + strconv.FormatInt(fileCounter, 16) + "tmp.diff")

	if err != nil {
		fmt.Println("Error on file creation", err)
		return err
	} else {

		defer f.Close()

		_, err := f.WriteString(str)

		if err != nil {
			fmt.Println("Error on file writing", err)
			return err
		}

		fileCounter++

		f.Sync()
	}

	return nil
}

func cloneRepo(rp string)  {

	destPath := path + strconv.FormatInt(folderCounter, 16)

	cmdName := "git"
	cmdArgs := []string{"clone",rp, destPath}

	cmd := exec.Command(cmdName, cmdArgs...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	folderCounter++

	fmt.Println("Result: " + out.String())
}


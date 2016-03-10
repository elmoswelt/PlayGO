package model

import (
	"strconv"
)

type Project struct  {

	ProjectID, Platform, Name, RepoURL, BranchName string
	BinaryFormat bool
}

// This is how a project model could look like. This needed to be set up somehow somewhere and
// be used as configuration for the process.
func ProjectList() []Project {

	projectList := []Project{

		Project{
			ProjectID:"2342",
			Platform:"iOS",
			Name:"MCV3",
			RepoURL:"git@github.groupondev.com:etampe/LocalizationStrings.git",
			BranchName:"testBranch",
			BinaryFormat:true},
		Project{
			ProjectID:"112342",
			Platform:"Android",
			Name:"MCV3",
			RepoURL:"git@github.groupondev.com:etampe/LocalizationStrings.git",
			BranchName:"testBranch",
			BinaryFormat:false},
	}

	return projectList
}

func (p Project) String() string {

	return "+++ Project Description +++" + "\n" +
		"ID: " + p.ProjectID + "\n" +
		"Name: " + p.Name + "\n" +
		"Platform: " + p.Platform + "\n" +
		"URL: " + p.RepoURL + "\n" +
		"Branch: " + p.BranchName + "\n" +
		"Binary: " + strconv.FormatBool(p.BinaryFormat)
}


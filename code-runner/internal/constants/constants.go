package constants

import "time"

const (
	SessionName          = "code-runner"
	UserData             = "user-data"
	SessionSecret        = "ed5173116fcb6eea062bf86d464fd697badffa9c"
	SessionUserKey       = "githubID"
	SessionUserName      = "githubNAME"
	SessionUserToken     = "githubTOKEN"
	Database             = "gcm"
	WorkspacesCollection = "workspaces"
	AppsCollection = "apps"
	SourceCode           = "https://github.com/hugobarzano/GCM"
	DockerRegistry       = "docker.pkg.github.com"
)

type Nature string

const (
	StaticApp Nature = "Single-Page"
	ApiRest Nature = "Api-Rest"
	DataService Nature = "Data-Service"
	DevOpsService Nature = "DevOps-Service"
)


var (
	Version = "0.0.0"
	GeneratedBanner =  `
***This App has been generated***`+"\n"+`
***Timestamp*** `+ time.Now().String()+ ``+"\n"+`
gcm/`+Version+``+"\n"+`[source-code](`+SourceCode+`)`+"\n"+`
***Powered by CesarCorp***`+"\n"+``)
package constants

import "time"

const (
	SessionName          = "code-runner"
	SessionSecret        = "ed5173116fcb6eea062bf86d464fd697badffa9c"
	SessionUserKey       = "githubID"
	SessionUserName      = "githubNAME"
	SessionUserToken     = "githubTOKEN"
	HttpAddress          = "localhost:8080"
	Database             = "generative-cloud"
	WorkspacesCollection = "workspaces"
	SourceCode = "https://github.com/hugobarzano/GCM"
)



var (
	Version = "0.0.0"
	GeneratorBanner =  `
***This App has been generated***`+"\n"+`
***Timestamp*** `+ time.Now().String()+ ``+"\n"+`
gcm/`+Version+``+"\n"+`[source-code](`+SourceCode+`)`+"\n"+`
***Powered by CesarCorp***`+"\n"+``)
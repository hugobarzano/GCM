package handlers

//func createApp(w http.ResponseWriter, req *http.Request) {
//
//	ctx := req.Context()
//
//	if err := req.ParseForm(); err != nil {
//		http.Error(w,
//			fmt.Sprintf("Error parsing form %s", err.Error()),
//			http.StatusInternalServerError)
//		return
//	}
//
//	session, err := sessionStore.Get(req, constants.SessionName)
//	if err != nil {
//		http.Error(w,
//			fmt.Sprintf("Session Error: %s", err.Error()),
//			http.StatusInternalServerError)
//		return
//	}
//	accessToken := session.Values[constants.SessionUserToken].(string)
//	user := session.Values[constants.SessionUserName].(string)
//
//	app := &models.App{
//		Name: strings.Replace(
//			req.FormValue("name"), "\"", "", -1),
//		Des: strings.Replace(
//			req.FormValue("description"), "\"", "", -1),
//		Owner: user,
//	}
//
//	gitApp := repos.GitApp{
//		App:app,
//	}
//	gitApp.Init(ctx,accessToken)
//	repo,_:=gitApp.CreateRepo(ctx)
//	if repo == nil{
//		fmt.Print("no repo")
//	}
//	app.Repository = repo.GetCloneURL()
//	dao:=store.InitMongoStore(ctx)
//	app,err=dao.CreateApp(ctx,app)
//	if err != nil {
//		http.Error(w,
//			fmt.Sprintf("Create App Error: %s", err.Error()),
//			http.StatusInternalServerError)
//		return
//	}
//
//	readme := generator.GenerateAppReadme(app)
//	fileOptions := repos.BuilFileOptions("Starting app...", user, readme)
//	_, err = gitApp.CommitFile(ctx, "README.md", fileOptions)
//	if err != nil {
//		http.Error(w,
//			fmt.Sprintf("CommitFile Error: %s", err.Error()),
//			http.StatusInternalServerError)
//		return
//	}
//	http.Redirect(w, req, "/workspace", http.StatusFound)
//}



//func generateApp(w http.ResponseWriter, r *http.Request) {
//
//	switch method := r.Method; method {
//	case http.MethodGet:
//		vars := mux.Vars(r)
//		app := vars["app"]
//		app = strings.Replace(app, "\"", "", -1)
//		session, _ := sessionStore.Get(r, constants.SessionName)
//		user := session.Values[constants.SessionUserName].(string)
//
//		ctx:=r.Context()
//		dao:=store.InitMongoStore(ctx)
//		appObj, err := dao.GetApp(ctx, user, app)
//		if err != nil {
//			http.Error(w,
//				fmt.Sprintf("Error getting app %s", err.Error()),
//				http.StatusNotFound)
//			return
//		}
//		if err := userAccessViews["generate"].Render(w, appObj); err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//	case http.MethodPost:
//		if err := r.ParseForm(); err != nil {
//			http.Error(w,
//				fmt.Sprintf("Error parsing form %s", err.Error()),
//				http.StatusInternalServerError)
//			return
//		}
//
//		fmt.Println(r.FormValue("nature"))
//
//		vars := mux.Vars(r)
//		app := vars["app"]
//		app = strings.Replace(app, "\"", "", -1)
//		session, _ := sessionStore.Get(r, constants.SessionName)
//		user := session.Values[constants.SessionUserName].(string)
//		accessToken := session.Values[constants.SessionUserToken].(string)
//		ctx:=r.Context()
//		dao:=store.InitMongoStore(ctx)
//		appObj, err :=dao.GetApp(ctx,user,app)
//		if err != nil {
//			http.Error(w,
//				fmt.Sprintf("Error getting app %s", err.Error()),
//				http.StatusNotFound)
//			return
//		}
//		dockerfile := generator.GenerateApacheDockerfile(appObj)
//		fileOptions := repos.BuilFileOptions("Generating Dockerfile", user, dockerfile)
//		appGit := repos.GitApp{
//			App:appObj,
//		}
//		appGit.Init(ctx,accessToken)
//		_,err=appGit.CommitFile(ctx, "Dockerfile", fileOptions)
//		ciFileData, err := ioutil.ReadFile("internal/resources/ci/imageBuilder.yml")
//		if err != nil {
//			fmt.Println("Error Reading")
//			fmt.Println(err)
//		}
//
//		ciFileOptions := repos.BuilFileOptions(
//			"Generating CI workflow action to build docker image",
//			user,
//			ciFileData)
//
//		_,err=appGit.CommitFile(ctx,		".github/workflows/ci.yml", ciFileOptions)
//		if err != nil {
//			fmt.Println("Error commit")
//			fmt.Println(err)
//		}
//		dockerApp:=deploy.DockerApp{
//			App:appObj,
//		}
//		dockerApp.Start(accessToken)
//		http.Redirect(w, r, "/workspace", http.StatusFound)
//	default:
//		fmt.Println("not supported")
//	}
//}

package generator

func (app *GenApp) generateStaticAppCode() {
	app.Data = make(map[string][]byte)
	app.Data["html/index.html"] = app.genIndexHtml()
	app.Data["html/js/index.js"] = app.genIndexJs()
	app.Data["html/css/style.css"] = app.genStyleCss()
}

func (app *GenApp) genIndexHtml() []byte {
	index :=
		`
<html>
<header>
	<title>` + app.App.Name + `</title>
		<link rel="stylesheet" type="text/css" href="css/style.css" />
</header>
<body>
	<br>
	<div>
		<h2>` + app.App.Name + `</h2>
		<br>
		<br>
		<h3><p>` + app.App.Des + `<p></h3>
	</div>
</body>
	<script src="js/index.js"></script>
</html>`

	return []byte(index)
}

func (app *GenApp) genIndexJs() []byte {
	index :=
		` 
alert( 'Hi, this is your site `+ app.App.Owner +`!' );
`
	return []byte(index)
}

func (app *GenApp) genStyleCss() []byte {
	index :=
		` 
h1, h2 {
	color: #800080;
	background-color: lightyellow;
	align:center
}
h3 {
	align:center
}
div {
	background-color:lightblue
}
`
	return []byte(index)
}

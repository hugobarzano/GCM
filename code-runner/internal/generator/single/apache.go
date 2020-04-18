package single

func GenStyleCss() []byte {
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

func GenIndexJs(owner string) []byte {
	index :=
		` 
alert( 'Hi, this is your site `+ owner +`!' );
`
	return []byte(index)
}

func GenIndexHtml(name, des string) []byte {
	index :=
		`
<html>
<header>
	<title>` + name + `</title>
		<link rel="stylesheet" type="text/css" href="css/style.css" />
</header>
<body>
	<br>
	<div>
		<h2>` + name + `</h2>
		<br>
		<br>
		<h3><p>` + des + `<p></h3>
	</div>
</body>
	<script src="js/index.js"></script>
</html>`

	return []byte(index)
}



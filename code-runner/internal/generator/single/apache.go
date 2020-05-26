package single

func GenStyleCss() []byte {
	index :=
		` 
h1, h2 {
	align:center
}
h3 {
	align:center
}
div {
	background-color:lightgreen
}
`
	return []byte(index)
}

func GenIndexJs(owner string) []byte {
	index :=
		` 
alert( 'Hi, this is your site ` + owner + `!' );
`
	return []byte(index)
}

func GenIndexHtml(name, des string) []byte {
	index := `
<!DOCTYPE html>
<html lang="en">

	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
		<meta name="description" content="">
		<meta name="cesar hugo" content="">
		<title>` + name + `</title>
		<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
		<link rel="stylesheet" type="text/css" href="css/style.css" />
	</head>

	<body>
		<nav class="navbar navbar-expand-lg navbar-dark bg-dark static-top">
		<div class="container">
			<a class="navbar-brand"> </a>
			<button class="navbar-toggler" type="button" data-toggle="collapse" aria-controls="navbarResponsive" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
		</div>
		</nav>

	<div class="container">
		<div class="row">
			<div class="col-lg-12 text-center">
				<h1 class="mt-5">Single Page Generated: ` + name + `</h1>
				<br>
				<hr>
				<br>
				<h3><p>` + des + `<p></h3>
				<br>
				<hr>
				<br>
			</div>
		</div>
	</div>
	</body>
		<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
		<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
		<script src="js/index.js"></script>	
</html>
`
	return []byte(index)

}

package main

var indexPage string = `
<!DOCTYPE html>
<html lang="ko">
	<head>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>UADL</title>
	</head>

	<body>
		<h1>Upload And Direct-Link</h1>
		<form enctype="multipart/form-data" action="/upload" method="post">
			<input type="file" name="fileName" />
			<input type="submit" value="upload" />
		</form>
		<ul>
			{{range .}}
			<li><a href="/files/{{.}}" download>{{.}}</a></li>
			{{end}}
		</ul>
	</body>
</html>`
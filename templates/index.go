package templates

const Index string = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Goupfile - Secure and anonymous file upload from the command line</title>
	<link rel="stylesheet" href="https://unpkg.com/@bafs/mu@0.3/mu.min.css" />
	<link rel="icon" type="image/png" href="/static/favicon-32x32.png" sizes="32x32">
	<style>
		body {
			font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen-Sans,Ubuntu,Cantarell,"Helvetica Neue",sans-serif;
		}
		small {
			font-size: 0.7rem !important;
		}
	</style>
</head>
<body>
	<h1>Goupfile <small>alpha</small></h1>
	<p>Goupfile is secure and anonymous file upload from the command line. It allows you to share files and get a links for them without leaving your terminal.</p>
	<h2>How does it work?</h2>
	<p>The goup CLI tool is currently under development. You can try out Goupfile by sending an HTTP POST request with a multipart form body to upload a file. The key should be named 'file'.</p>
	<pre>
curl -F "file=@/home/user/file.txt" https://goupfile.com
</pre>
	<p>To download a file:</p>
<pre>
curl https://goupfile.com/OVn1C1
</pre>
	<p>In the future, you will be able to use Goupfile through the <a href="https://github.com/goupfile/goup">goup CLI tool</a>.</p>
	<p>For example:</p>
	<pre>
$ goup file.txt
Uploading file...
https://goupfile.com/aw9kzm
$ goup download aw9kzm
</pre>
	<hr>
	<a href="https://github.com/goupfile">GitHub</a>
</body>
</html>
`

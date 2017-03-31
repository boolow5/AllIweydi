<!DOCTYPE html>
<html {{if eq .Lang "ar-SA"}}lang="ar" DIR="RTL"{{else}}lang="en"{{end}}>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>Bolow Template</title>

    {{.favicons}}

    {{.css}}
    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  <body>
    <div class="top-menu">
      {{if eq .Lang "ar-SA"}}
        {{.navbarAr}}
      {{else}}
        {{.navbar}}
      {{end}}
    </div>

    {{.featured}}

    <div class="container  main-wrapper">
      <div class="row">
        {{.sidebarLeft}}
        <div class="col-sm-5 main-content min-height">
          {{.LayoutContent}}
        </div>
        {{.sidebarRight}}

      </div>
    </div>

    {{.footer}}

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <script src="static/bootstrap/js/bootstrap.min.js"></script>
    <script src="/static/js/jquery.cookie.js"></script>
    <script src="static/js/main.js"></script>
  </body>
</html>

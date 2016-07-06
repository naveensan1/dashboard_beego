<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <title>Regression Dashboard</title>
  <link rel="stylesheet" href="../static/css/bootstrap.min.css">
  <link rel="stylesheet" href="../static/css/bootstrap-theme.min.css">
  <link rel="stylesheet" href="../static/css/site.css">
</head>
<body>
  <div id="page">
    <header class="container">
      <div id="menu" class="navbar navbar-default navbar-fixed-top">
        <div class="navbar-header">
          <button class="btn btn-success navbar-toggle"
                  data-toggle="collapse"
                  data-target=".navbar-collapse"><span class="glyphicon glyphicon-chevron-down"></span></button>
          <div id="logo">
            <a href='.'><h4>Regression Dashboard</h4></a>
          </div>
        </div>
        <div class="navbar-collapse collapse">
          <ul class="nav navbar-nav navbar-right">
           <li class="nav active"><a href=".">Home</a></li>
           <li class="nav"><a href="contact.html">Contact</a></li>
           <li class="nav"><a href="contact.html"></a></li>
          </ul>
       </div>
      </div>
    </header>
    <section id="body" class="container">
      {{block "content" . }}{{end}}
    </section>
    <!--<img src="../img/openstack-logo.jpeg" />-->
    <hr />
    <footer class="container">
      <p>© 2016 Nokia. All rights reserved. <a href="http://www.nuagenetworks.com/">Nuage Networks</a> is a Nokia venture.</p>

    </footer>
  </div>
  <script src="../static/js/jquery-2.2.4.min.js"></script>
  <script src="../static/js/bootstrap.min.js"></script>
</body>
</html>

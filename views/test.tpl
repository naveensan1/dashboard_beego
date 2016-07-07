{{ define "content" }}
<div class="col-md-10">
  <div class="row">
    <h4>hi {{ .Name }}</h4>
  </div>
</div>
{{ end }}
{{ template "_layout.tpl" }}

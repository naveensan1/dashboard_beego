{{define "content"}}
<div class="col-md-10">
  <div class="row">
    <table class="table table-striped table-bordered table-hover">
      <thead>
        <tr>
          <th></th><th colspan="2">Build</th>
        </tr>
        <tr>
          <th>Subtopology</th>
          <th>3.2/253</th>
          <th>3.2/252</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>
            openstackIcehouseUbuntu1404
          </td>
          <td>
            Pass
          </td>
          <td>
            Pass
          </td>
        </tr>
        <tr>
          <td>
            openstackIcehouseEl7
          </td>
          <td>
            Fail
          </td>
          <td>
            Fail
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
{{end}}
{{template "_layout.tpl"}}

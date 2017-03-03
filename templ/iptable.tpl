<table class="table">
<thead>
<tr>
<th>Index</th>
<th>Public Port</th>
<th>Local Port</th>
<th>Local IP</th>
<th>Proto</th>
<th>Operate</th>
</tr>
</thead>
<tbody>

{{range .}}
<tr class="table .table-striped">
<td>{{.Index}}</td>
<td>{{.Public_port}}</td>
<td>{{.Local_port}}</td>
<td>{{.Local_ip}}</td>
<td>{{.Proto}}</td>
<td>
<p>
<button id="modal-{{.Index}}" href="#modal-container-{{.Index}}" role="button" class="btn btn-mini btn-primary" data-toggle="modal">详情</button>
<button class="btn btn-mini" type="button">编辑</button>
<button class="btn btn-mini btn-danger" type="button">删除</button>
</p>
</td>
</tr>
{{end}}
</tbody>
</table>

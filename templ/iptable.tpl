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

<form method="post" id="form-for-entry-{{.Index}}">
<input type="hidden" name="operate" value="del"/>
<input type="hidden" name="new_public_port" value="{{.Public_port}}"/>
<input type="hidden" name="new_local_port" value="{{.Local_port}}"/>
<input type="hidden" name="new_local_ip" value="{{.Local_ip}}"/>
<input type="hidden" name="new_proto" value="{{.Proto}}"/>
</form>
<tr class="table .table-striped" id="entry-{{.Index}}">
<td>{{.Index}}</td>
<td>{{.Public_port}}</td>
<td>{{.Local_port}}</td>
<td>{{.Local_ip}}</td>
<td>{{.Proto}}</td>
<td>
<p>
<button id="modal-{{.Index}}" href="#modal-container-{{.Index}}" role="button" class="btn btn-mini btn-primary" data-toggle="modal">详情</button>
<button class="btn btn-mini btn-danger" onclick="ondel('{{.Index}}')"   type="button">删除</button>
<script>
function ondel(index){

	if (confirm("real?") == true){
		var which = "form-for-entry-"+index;
		document.getElementById(which).submit();
	}else{
	}
}	
</script>
</p>
</td>
</tr>
{{end}}
</tbody>
</table>

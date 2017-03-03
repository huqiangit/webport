{{range .}}
<div class="modal fade" id="modal-container-{{.Index}}" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
<div class="modal-dialog">
<div class="modal-content">
<div class="modal-header">
<button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
<h4 class="modal-title" id="myModalLabel">
NAT详情
</h4>
</div>
<div class="modal-body">
<form class="form-inline">
  <input type="text" class="input-small" data-toggle="tooltip" placeholder="{{.Public_port}}">
  <input type="text" class="input-small" placeholder="{{.Local_port}}">
  <input type="text" class="input-small" placeholder="{{.Local_ip}}">
  <input type="text" class="input-small" placeholder="{{.Proto}}">
  <button type="submit" class="btn">Set</button>
</form>
</div>
<div class="modal-footer">
<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button> <button type="button" class="btn btn-primary">保存</button>
</div>
</div>
</div>
</div>
{{end}}

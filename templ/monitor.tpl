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
<div style="font-size:1px" class="modal-body">
<p size="2">
POSTROUTING -d {{.Local_ip}}/32 -p {{.Proto}} -m {{.Proto}} --dport {{.Local_port}} -j SNAT --to-source [public_ip]
</p>
<p size="2">
PREROUTING -d [public_ip] -p {{.Proto}} -m {{.Proto}} --dport {{.Public_port}} -j DNAT --to-destination {{.Local_ip}}:{{.Local_port}}
</p>
<p size="2">
FORWARD -d {{.Local_ip}}/32 -p {{.Proto}} -m {{.Proto}} --dport {{.Local_port}} -j ACCEPT
</p>
<!--
<form class="form-inline">
  <input type="text" class="input-small" data-toggle="tooltip" placeholder="{{.Public_port}}">
  <input type="text" class="input-small" placeholder="{{.Local_port}}">
  <input type="text" class="input-small" placeholder="{{.Local_ip}}">
  <input type="text" class="input-small" placeholder="{{.Proto}}">
  <button type="submit" class="btn">Set</button>
</form>
-->
</div>
<div class="modal-footer">
<!--
<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button> <button type="button" class="btn btn-primary">保存</button>
-->
</div>
</div>
</div>
</div>
{{end}}

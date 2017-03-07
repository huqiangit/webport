<script>
<!--document.getElementById("sessionkey").value = getCookie("key"); -->

function getCookie(cname){
	var name = cname + "=";
	var ca = document.cookie.split(';');
	for(var i=0; i<ca.length; i++) {
		var c = ca[i].trim();
		if (c.indexOf(name)==0) return c.substring(name.length,c.length);
	}
	return "";
}


</script>
<div class="container">
	<div class="row clearfix">
		<div class="col-md-12 column">
				
			<button id="modal-newentry" href="#modal-container-newentry" role="button" class="btn" data-toggle="modal">添加</button>
			
			<div class="modal fade" id="modal-container-newentry" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
				<div class="modal-dialog">
					<div class="modal-content">
						<div class="modal-header">
							 <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
							<h4 class="modal-title" id="myModalLabel">
								new entry
							</h4>
						</div>
						<form class="form-horizontal" role="form" method="post">
							<input type="hidden" name="operate" value="add"/>
							<input type="hidden" name="sessionkey" value="" id="sessionkey"/>

							<div class="modal-body">
								<div class="form-group">
									 <label for="new_public_port" class="col-sm-2 control-label">公网端口</label>
									<div class="col-sm-10">
										<input type="text" name="new_public_port" class="form-control" id="new_public_port" />
									</div>
								</div>
								<div class="form-group">
									 <label for="new_local_port" class="col-sm-2 control-label">内网端口</label>
									<div class="col-sm-10">
										<input type="text" name="new_local_port" class="form-control" id="new_local_port" />
									</div>
								</div>
								<div class="form-group">
									 <label for="new_local_ip" class="col-sm-2 control-label">内网ip</label>
									<div class="col-sm-10">
										<input type="text" name="new_local_ip" class="form-control" id="new_local_ip" />
									</div>
								</div>

								<div class="form-group">
									<div class="col-sm-offset-2 col-sm-10">
										<div class="checkbox">
											 <label><input name="new_record" type="checkbox" />record</label>
										</div>
									</div>
								</div>
								<div class="form-group">
									 <label for="new_proto" class="col-sm-2 control-label">协议类型</label>
									<div class="col-sm-10">
										<select name="new_proto" id="new_proto">
										<option>tcp</option>
										<option>udp</option> 
										</select> 
									</div>
								</div>  
							</div>
							<div class="modal-footer">
								 <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button> 
								 <button type="submit" class="btn btn-primary">添加</button>
							</div>
						</form>
					</div>
					
				</div>
				
			</div>
			
		</div>
	</div>
</div>

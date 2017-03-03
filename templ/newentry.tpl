

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
						<form class="form-horizontal" role="form">
							<div class="modal-body">
								<div class="form-group">
									 <label for="new_public_port" class="col-sm-2 control-label">公网端口</label>
									<div class="col-sm-10">
										<input type="text" class="form-control" id="new_public_port" />
									</div>
								</div>
								<div class="form-group">
									 <label for="new_local_port" class="col-sm-2 control-label">内网端口</label>
									<div class="col-sm-10">
										<input type="text" class="form-control" id="new_local_port" />
									</div>
								</div>
								<div class="form-group">
									 <label for="new_local_ip" class="col-sm-2 control-label">内网ip</label>
									<div class="col-sm-10">
										<input type="text" class="form-control" id="new_local_ip" />
									</div>
								</div>

								<div class="form-group">
									<div class="col-sm-offset-2 col-sm-10">
										<div class="checkbox">
											 <label><input type="checkbox" />record</label>
										</div>
									</div>
								</div>
								<div class="form-group">
									 <label for="new_proto" class="col-sm-2 control-label">协议类型</label>
									<div class="col-sm-10">
										<select id="new_proto">
										<option>tcp</option>
										<option>udp</option> 
										</select> 
									</div>
								</div>  
							</div>
							<div class="modal-footer">
								 <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button> 
								 <button type="submit" class="btn btn-primary">保存</button>
							</div>
						</form>
					</div>
					
				</div>
				
			</div>
			
		</div>
	</div>
</div>

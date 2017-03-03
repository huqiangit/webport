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
						<div class="modal-body">
							<form class="form-horizontal" role="form">
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

										<select>
										<option>tcp</option>
										<option>udp</option> 
										</select>
								<div class="form-group">
									<div class="col-sm-offset-2 col-sm-10">
										 <button type="submit" class="btn btn-default">Add</button>
									</div>
								</div>
							</form>
						</div>
						<div class="modal-footer">
							 <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button> 
							 <button type="button" class="btn btn-primary">保存</button>
						</div>
					</div>
					
				</div>
				
			</div>
			
		</div>
	</div>
</div>

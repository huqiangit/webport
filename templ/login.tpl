<!DOCTYPE html>
<html>
<body>

<!-- 新 Bootstrap 核心 CSS 文件 -->
<link href="/script/bootstrap-3.3.7-dist/css/bootstrap.min.css" rel="stylesheet">

<!-- 可选的Bootstrap主题文件（一般不使用） -->
<script src="/script/bootstrap-3.3.7-dist/css/bootstrap-theme.min.css"></script>

<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
<script src="/script/jquery-3.1.1.min.js"></script>


<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
<script src="/script/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
<div class="container">
<div class="row clearfix">
<div class="col-md-12 column">
<script>
        $(function () { $("[data-toggle='tooltip']").tooltip(); });
</script>

<form  method="post" class="form-horizontal">

  <div class="control-group">
    <label class="control-label" for="inputUsername">Username</label>
    <div class="controls">
      <input type="text" id="inputUsername" placeholder="Username" name="username">
    </div>
  </div>

  <div class="control-group">
    <label class="control-label" for="inputPassword">Password</label>
    <div class="controls">
      <input type="password" id="inputPassword" placeholder="Password" name="password">
    </div>
  </div>

  <div class="control-group">
    <div class="controls">
      <label class="checkbox">
        <input type="checkbox"> Remember me
      </label>
      <button type="submit" class="btn">Sign in</button>
    </div>
  </div>

</form>


</div>
</div>
</div>
</body>
</html>

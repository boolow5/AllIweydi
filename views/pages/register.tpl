<h1>{{i18n $.Lang "register"}}</h1>
<div id="form-container" class="list-item-container q-list">
  <form id="register-form" class="form" method="post" action="/register">
    {{ .xsrfdata }}
    <div class="input-group">
      <input type="text" class="form-control" name="full_name" placeholder='{{i18n .Lang "full_name"}}' id="register-full_name" autofocus>
      <input type="text" class="form-control" name="username" placeholder='{{i18n .Lang "username"}}' id="register-username">
      <input type="password" class="form-control" name="password" placeholder='{{i18n .Lang "password"}}' id="login-email">
    </div>
    <button class="btn btn-primary" id="register-user-btn">{{i18n .Lang "login_btn"}}</button>
  </form>

</div>

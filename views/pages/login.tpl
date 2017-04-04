<h1>{{i18n $.Lang "login"}}</h1>
<div id="form-container" class="list-item-container q-list">
  <form id="login-form" class="form" method="post" action="/login">
    {{ .xsrfdata }}
    <div class="input-group">
      <input type="text" class="form-control" name="username" placeholder='{{i18n .Lang "username"}}' id="login-username" autofocus>
      <input type="password" class="form-control" name="password" placeholder='{{i18n .Lang "password"}}' id="login-email">
    </div>
    <button class="btn btn-primary" id="login-user-btn">{{i18n .Lang "login_btn"}}</button>
  </form>
</div>

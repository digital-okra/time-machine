<Router url="{url}">
  <Route path="/admin"><AdminMainView /></Route>
    <Route path="/admin/user/:userid" let:params><AdminTaskView userid="{params.userid}"/></Route>
  <Route path="/tasks"><NormalMainView /></Route>
  <Route path="/"><LoginView /></Route>
</Router>

<script>
  import { Router, Route } from "svelte-routing";
  import LoginView from './views/LoginView.svelte';
  import NormalMainView from './views/NormalMainView.svelte';
  import AdminMainView from './views/AdminMainView.svelte';
  import AdminTaskView from './views/AdminTaskView.svelte';

  let url = "";
</script>

<style>
</style>

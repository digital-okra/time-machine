<div class="login-view">
  <Paper color="primary" square style="flex: 1;">
    <div class="header-container">
      <i class="material-icons" style="font-size: 7rem;">av_timer</i>
    </div>
  </Paper>
  <ul class="login-form">
    <li class="spacing">
      <Textfield withLeadingIcon variant="filled" bind:value={username} label="Username" style="width: 100%">
        <Icon class="material-icons">person</Icon>
      </Textfield>
    </li>
    <li class="spacing">
      <Textfield withLeadingIcon variant="filled" bind:value={password} label="Password" style="width: 100%" type="password">
        <Icon class="material-icons">fiber_pin</Icon>
      </Textfield>
    </li>
    <li class="spacing align-bottom">
      <Button on:click={onLogin} variant="raised" style="width: 100%;">
        {#if !loading}
          Login
        {/if}
        {#if loading}
          <Spinner
            size="30"
            speed="750"
            color="#FFFFFF"
            thickness="2"
            gap="40"
          />
        {/if}
      </Button>
    </li>
    <Snackbar bind:this={errorSnackbar}>
      <Label>{errorMessage}</Label>
      <Actions>
        <IconButton class="material-icons" title="Dismiss">close</IconButton>
      </Actions>
    </Snackbar>
  </ul>
</div>

<script>
  import Paper from '@smui/paper';
  import { navigate } from "svelte-routing";
  import Textfield from '@smui/textfield';
  import Icon from '@smui/textfield/icon/index';
  import Button from '@smui/button';

  import Snackbar, {Actions, Label} from '@smui/snackbar';
  import IconButton from '@smui/icon-button';

  import Spinner from 'svelte-spinner';

  import { onMount } from 'svelte';

  import { loginUser } from '../services/LoginService.js';

  let username = "";
  let password = "";  
  let loading = false;

  let errorSnackbar;
  let errorMessage = "";

  async function onLogin() {
    try {
      loading = true;
      let token = await loginUser(username, password);

      let storage = window.localStorage;
      storage.setItem("jwt", token.jwt);
      storage.setItem("type", token.type);
      storage.setItem("id", token.id);

      // Navigate to user page
      if(token.type === "normal") {
        navigate("/tasks", { replace: true });
      } else if(token.type === "admin") {
        navigate("/admin", { replace: true });
      } else {
        throw "Something went wrong";
      }
    } catch(err) {
      if(err === 400 || err === 401) {
        errorMessage = "Invalid username or password";
      } else {
        errorMessage = `Unknown error ${err}. Try again!`;
      }
      errorSnackbar.open();
      loading = false;
    }
  }

  onMount(() => {
    let storage = window.localStorage;
    let jwt = storage.getItem("jwt");
    let type = storage.getItem("type");
    let id = storage.getItem("id");

    // If there is already an existed logged in user, log them in directly
    if(jwt != null && type != null && id != null) { 
      // Navigate to user page
      if(type === "normal") {
        navigate("/tasks", { replace: true });
      } else if(type === "admin") {
        navigate("/admin", { replace: true });
      }
    }
  });
</script>

<style>
  .login-view {
    display: flex;
    flex-direction: column;
  }

  .header-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    width: 100%;
    height: 100%;

    color: white;
  }

  .login-form {
    flex: 2;
    padding: 0.1rem 2rem;
    list-style-type: none;
    display: flex;
    flex-direction: column;
    align-items: center;

    height: 100%;
  }

  .spacing {
    margin: 0.5rem 0rem;
    width: 100%;
  }

  .align-bottom {
    position: absolute;
    bottom: 10px;
    width: 80%;
  }
</style>


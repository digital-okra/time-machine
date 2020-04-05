<div class="login-view">
  <Paper color="primary" square style="flex: 1;">
    <div class="header-container">
      <h1 class="mdc-typography--headline4">Time Machine</h1>
    </div>
  </Paper>
  <ul class="login-form">
    <li class="spacing">
      <Textfield withLeadingIcon variant="filled" bind:value={valueFilledB} label="Username" style="width: 100%">
        <Icon class="material-icons">person</Icon>
      </Textfield>
    </li>
    <li class="spacing">
      <Textfield withLeadingIcon variant="filled" bind:value={valueFilledB} label="Password" style="width: 100%">
        <Icon class="material-icons">fiber_pin</Icon>
      </Textfield>
    </li>
    <li class="spacing align-bottom">
      <Button on:click={() => alert('Clicked!')} variant="raised" style="width: 100%;">Login</Button>
    </li>
    <li class="spacing">
      <Button on:click={() => alert('Clicked!')} variant="raised" style="width: 100%;">Register</Button>
    </li>
  </ul>
</div>

<script>
  import Paper from '@smui/paper';
  import Textfield from '@smui/textfield';
  import Icon from '@smui/textfield/icon/index';
  import Button from '@smui/button';
  
  let valueFilledB = "";
</script>

<style>
  .login-view {
    display: flex;
    flex-direction: column;

    height: 100vh;
  }

  .header-container {
    display: flex;
    justify-content: center;
    align-items: center;

    width: 100%;
    height: 100%;

    color: white;
  }

  .login-form {
    flex: 1;
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
    margin-top: auto;
  }
</style>


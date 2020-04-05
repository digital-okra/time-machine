<div class="login-view">
  <Paper color="primary" square style="flex: 1;">
    <div class="header-container">
      <i class="material-icons" style="font-size: 7rem;">av_timer</i>
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
      <Button on:click={() => {}} variant="raised" style="width: 100%;">Login</Button>
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
    margin-top: auto;
  }
</style>


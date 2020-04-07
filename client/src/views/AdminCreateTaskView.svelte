<div class="wrapper">
  <TopAppBar variant="static" color="primary" class="fixed-top-bar">
    <Row>
      <Section>
        <IconButton class="material-icons" on:click={onReturn}>keyboard_backspace</IconButton>
        <Title>Create Task</Title>
      </Section>
    </Row>
  </TopAppBar>

  <ul class="create-form">
    <li class="spacing">
      <Textfield variant="standard" bind:value={name} label="Task name" style="width: 100%">
      </Textfield>
    </li>
    <li class="spacing">
      <Select bind:value={userChoice} label="Assign to" style="width: 100%">
        <Option value=""></Option>
        {#each users as user}
          <Option value={user.id}>{user.rank} {user.first_name} {user.last_name}</Option>
        {/each}
      </Select>
    </li>
    <li class="spacing align-bottom">
      <Button on:click={onCreate} variant="raised" style="width: 100%;">
        {#if !loading}
          Create
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
  </ul>
  <Snackbar bind:this={errorSnackbar}>
    <Label>{errorMessage}</Label>
    <Actions>
      <IconButton class="material-icons" title="Dismiss">close</IconButton>
    </Actions>
  </Snackbar>
</div>

<script>
  import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
  import IconButton from '@smui/icon-button';
  import { navigate } from "svelte-routing";
  import Button from '@smui/button';
  import Textfield from '@smui/textfield';
  import Spinner from 'svelte-spinner';
  import Select, {Option} from '@smui/select';
  import Snackbar, {Actions, Label} from '@smui/snackbar';
  
  import { getAllUsers } from '../services/UserService.js';
  import { createTask } from '../services/TaskService.js';
  import { jwt_store } from '../store.js';
  import { onMount } from 'svelte';
  
  let name = "";
  let userChoice = "";
  let loading = false;
  let jwt = $jwt_store;
  let errorSnackbar;
  let errorMessage = "";
  
  let users = [];
  
  onMount(async () => {
    users = await getAllUsers(jwt)

    if(users == null) {
      tasks = [];
    }
  });

  async function onCreate() {
    try {
      loading = true;
      const found = users.find(user => `${user.rank} ${user.first_name} ${user.last_name}` === userChoice)

      if(name == "") {
        throw "Task name cannot be empty";
      }

      await createTask(jwt, name, found.id);
      navigate("/admin", {replace: true});
    } catch(err) {
        errorMessage = `Unknown error ${err}. Try again!`;
        errorSnackbar.open();

      loading = false;
    }
  }

  function onReturn() {
    // Do Nothing
    navigate("/admin", {replace: true});
  }
</script>

<style>
  .wrapper {
    height: 80vh;
  }
  .create-form {
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

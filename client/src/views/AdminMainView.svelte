<TopAppBar variant="fixed" color="primary" class="fixed-top-bar">
  <Row>
    <Section>
      {#if !showTaskView}
      <Title>Technician List</Title>
      {/if}
      {#if showTaskView}
        <IconButton class="material-icons" on:click={onReturn}>keyboard_backspace</IconButton>
        <Title>{currentActiveUser}'s Tasks</Title>
      {/if}
    </Section>
    <Section align="end" toolbar>
      {#if showTaskView}
        <IconButton class="material-icons" aria-label="Create task" on:click={() => dialog.open()}>library_add</IconButton>
      {/if}
      <IconButton class="material-icons" aria-label="Logout" on:click={onLogout}>exit_to_app</IconButton>
    </Section>
  </Row>
  
  {#if showTaskView}
    <TabBar tabs={['Completed', 'Active', 'Verified']} let:tab bind:active class="nav-tab-bar">
      <!-- Notice that the `tab` property is required! -->
      <Tab {tab} class="nav-tab">
        <Label>{tab}</Label>
      </Tab>
    </TabBar>
  {/if}
</TopAppBar>

{#if !showTaskView}
  <div class="users-view">
    <List>
      {#each users as user}
        <Item on:click="{viewUserTasks(user)}">
          <Text>
            {user.rank} {user.first_name} {user.last_name}
          </Text>
        </Item>
      {/each}
    </List>
  </div>
{/if}

{#if showTaskView}
  {#if showCompleted}
    <div class="tab-view completed-tab">
      {#if completedTasks.length == 0}
        <List>
          <Item>
            <Text><span style="color: #9e9e9e;">No completed tasks</span></Text>
          </Item>
        </List>
      {:else}
      <List checklist twoLine>
        {#each completedTasks as task}
          <Item on:click={toggleVerify(task)}>
            <Text>
              <PrimaryText>{task.name}</PrimaryText>
              <SecondaryText><span class="mdc-typography--overline">{task.assigned_by}</span></SecondaryText>
            </Text>
              <Meta>
                <Checkbox on:click={toggleVerify(task)} bind:checked={task.verified} />
              </Meta>
          </Item>
        {/each}
      </List>
      {/if}
    </div>
  {/if}

  {#if showActive}
    <div class="tab-view completed-tab">
    {#if activeTasks.length == 0}
      <List>
        <Item>
          <Text><span style="color: #9e9e9e;">No active tasks</span></Text>
        </Item>
      </List>
    {:else}
      <List twoLine>
        {#each activeTasks as task}
          <Item>
            <Text>
              <PrimaryText>{task.name}</PrimaryText>
                <SecondaryText><span class="mdc-typography--overline">{task.assigned_by}</span></SecondaryText>
            </Text>
          </Item>
        {/each}
      </List>
    {/if}
    </div>
  {/if}

  {#if showVerified}
    <div class="tab-view verified-tab">
    {#if verifiedTasks.length == 0}
      <List>
        <Item>
          <Text><span style="color: #9e9e9e;">No verified tasks</span></Text>
        </Item>
      </List>
    {:else}
      <List checklist threeLine>
        {#each verifiedTasks as task}
          <Item on:click={toggleVerify(task)}>
            <Text>
              <PrimaryText>{task.name}</PrimaryText>
              <SecondaryText><span class="mdc-typography--overline">{task.assigned_by}</span></SecondaryText>
              <SecondaryText>Verified by {task.verified_by}</SecondaryText>
            </Text>
              <Meta>
                <Checkbox on:click={toggleVerify(task)} bind:checked={task.verified} />
              </Meta>
          </Item>
        {/each}
      </List>
    {/if}
    </div>
  {/if}
{/if}

<!-- CREATE TASK DIALOG -->
<Dialog
  bind:this={dialog}
  aria-labelledby="dialog-title"
  aria-describedby="dialog-content"
  on:MDCDialog:closed={closeHandler}
>
  <DialogTitle id="dialog-title">Create Task</DialogTitle>
  <Content id="dialog-content">
      <Textfield variant="standard" bind:value={taskname} label="Task name" style="width: 100%">
      </Textfield>
  </Content>
  <Actions>
    <Button>
      <Label>Cancel</Label>
    </Button>
    <Button on:click={onCreateTask}>
      <Label>Create</Label>
    </Button>
  </Actions>
</Dialog>

<Snackbar bind:this={errorSnackbar}>
  <SnackLabel>{errorMessage}</SnackLabel>
  <SnackActions>
    <IconButton class="material-icons" title="Dismiss">close</IconButton>
  </SnackActions>
</Snackbar>

<script>
  import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
  import List, {Group, Subheader, Meta, Label, Item, Text, PrimaryText, SecondaryText} from '@smui/list';
  import Textfield from '@smui/textfield';
  import IconButton from '@smui/icon-button';
  import Checkbox from '@smui/checkbox';
  import { navigate } from "svelte-routing";

  import Dialog, {Title as DialogTitle, Content, Actions} from '@smui/dialog';
  import Button from '@smui/button';
  import Snackbar, {Actions as SnackActions, Label as SnackLabel} from '@smui/snackbar';

  import Tab from '@smui/tab';
  import TabBar from '@smui/tab-bar';
  
  import { getAllUsers } from '../services/UserService.js';
  import { getTasks, toggleVerifyTask, createTask } from '../services/TaskService.js';
  
  import { onMount } from 'svelte';

  let showTaskView = false;
  let active = "Completed";
  let currentActiveUser = "";
  let currentActiveUserId = "";
  let jwt;
  let userid;

  let dialog;
  let taskname = "";
  let errorSnackbar;
  let errorMessage = "";

  let tasks = [];
  let users = [];

  $: activeTasks = tasks.filter(task => task.assigned_to == currentActiveUserId && !task.completed && !task.verified);
  $: completedTasks = tasks.filter(task => task.assigned_to == currentActiveUserId && !task.verified && task.completed);
  $: verifiedTasks = tasks.filter(task => task.assigned_to == currentActiveUserId && task.verified && task.completed);
  
  $: showVerified = (active == "Verified");
  $: showActive = (active == "Active");
  $: showCompleted = (active == "Completed");

  onMount(async () => {
    let storage = window.localStorage;
    jwt = storage.getItem("jwt");
    userid = storage.getItem("id");

    if(jwt == null || userid == null) {
      navigate("/", { replace: true });
    }

    users = await getAllUsers(jwt)
    tasks = await getTasks(jwt)

    if(users == null) {
      tasks = [];
    }
    if(tasks == null) {
      tasks = [];
    }
  });
 
  function onReturn() {
    showTaskView = false;
    
    currentActiveUser = "";
    currentActiveUserId = "";
  }

  function closeHandler() {
    // Do nothing
  }

  async function onCreateTask() {
    try {
      if(taskname == "") {
        throw "Task name cannot be empty";
      }

      await createTask(jwt, taskname, currentActiveUserId);
      tasks = await getTasks(jwt)
    } catch(err) {
      console.log(err)
        errorMessage = `${err}. Try again!`;
        errorSnackbar.open();
    }
  }

  function viewUserTasks(user) {
    currentActiveUser = user.first_name;
    currentActiveUserId = user.id;

    showTaskView = true;
  }

  async function toggleVerify(task) {
    console.log("verified")
    toggleVerifyTask(jwt, task, userid);
    tasks = await getTasks(jwt);
  }
  
  async function onLogout() {
    // Clear localStorage
    let storage = window.localStorage;
    storage.clear();

    navigate("/", { replace: true});
  }
</script>

<style>
  .users-view {
    padding-top: 3rem;
  }
  
  :global(.nav-tab-bar) {
    background-color: #6200ea;
    color: white;
  }
  
  :global(.fixed-top-bar) {
    box-shadow: 0 2px 4px -1px rgba(0,0,0,.2),0 4px 5px 0 rgba(0,0,0,.14),0 1px 10px 0 rgba(0,0,0,.12);
  }

  :global(.nav-tab) {
    color: white;
  }
  
  :global(.mdc-tab-indicator .mdc-tab-indicator__content--underline) {
    border-color: white;
    border-top-width: 5px;
  }
  
  .tab-view {
    padding-top: 6.5rem;
  }
</style>

<TopAppBar variant="fixed" color="primary" class="fixed-top-bar">
  <Row>
    <Section>
      <IconButton class="material-icons" on:click={onReturn}>keyboard_backspace</IconButton>
        <Title>{user.first_name}'s Tasks</Title>
    </Section>
    <Section align="end" toolbar>
      <IconButton class="material-icons" aria-label="Create task" on:click={() => dialog.open()}>library_add</IconButton>
    </Section>
  </Row>
  
  <TabBar tabs={['Completed', 'Active', 'Verified']} let:tab bind:active class="nav-tab-bar">
    <!-- Notice that the `tab` property is required! -->
    <Tab {tab} class="nav-tab">
      <Label>{tab}</Label>
    </Tab>
  </TabBar>
</TopAppBar>

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

<!-- Error message snackbar -->
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
  
  import { getUserById } from '../services/UserService.js';
  import { getTasks, toggleVerifyTask, createTask } from '../services/TaskService.js';
  
  import { onMount } from 'svelte';

  export let userid;

  let active = "Completed";
  let jwt;
  let selfUserId;

  let dialog;
  let taskname = "";
  let errorSnackbar;
  let errorMessage = "";

  let user = {};
  let tasks = [];

  $: activeTasks = tasks.filter(task => task.assigned_to == userid && !task.completed && !task.verified);
  $: completedTasks = tasks.filter(task => task.assigned_to == userid && !task.verified && task.completed);
  $: verifiedTasks = tasks.filter(task => task.assigned_to == userid && task.verified && task.completed);
  
  $: showVerified = (active == "Verified");
  $: showActive = (active == "Active");
  $: showCompleted = (active == "Completed");

  onMount(async () => {
    let storage = window.localStorage;
    jwt = storage.getItem("jwt");
    selfUserId = storage.getItem("id");

    if(jwt == null || selfUserId == null) {
      navigate("/", { replace: true });
    }

    user = await getUserById(jwt, userid)
    tasks = await getTasks(jwt)

    if(tasks == null) {
      tasks = [];
    }
  });
 
  function onReturn() {
    navigate("/admin", { replace: true });
  }

  function closeHandler() {
    // Do nothing
  }

  async function onCreateTask() {
    try {
      if(taskname == "") {
        throw "Task name cannot be empty";
      }

      await createTask(jwt, taskname, userid);
      tasks = await getTasks(jwt)
    } catch(err) {
      console.log(err)
        errorMessage = `${err}. Try again!`;
        errorSnackbar.open();
    }
  }

  async function toggleVerify(task) {
    console.log("verified")
    toggleVerifyTask(jwt, task, selfUserId);
    tasks = await getTasks(jwt);
  }  
</script>

<style>  
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

 <TopAppBar variant="fixed" color="primary" class="fixed-top-bar">
  <Row>
    <Section>
      <Title>Your Tasks</Title>
    </Section>
    <Section align="end" toolbar>
      <IconButton class="material-icons" aria-label="Logout" on:click={onLogout}>exit_to_app</IconButton>
    </Section>
  </Row>
  <TabBar tabs={['Active', 'Completed']} let:tab bind:active class="nav-tab-bar">
    <!-- Notice that the `tab` property is required! -->
    <Tab {tab} class="nav-tab">
      <Label>{tab}</Label>
    </Tab>
  </TabBar>
</TopAppBar>

{#if showActive}
<div class="active-tab">
  <List twoLine checklist>
    {#each activeTasks as task}
      <Item>
        <Text>
          <PrimaryText>{task.name}</PrimaryText>
          <SecondaryText><span class="mdc-typography--body2">{task.assigned_to}</span></SecondaryText>
        </Text>
          <Meta>
            <Checkbox on:click="{toggleCompleted(task)}" bind:checked={task.completed} />
          </Meta>
      </Item>
    {/each}
  </List>
</div>
{/if}

{#if showCompleted}
<div class="completed-tab">
  <div>
    <div class="divider waiting mdc-typography--overline">
      Pending verification
    </div>
    {#if completedTasks.length == 0}
      <List>
        <Item>
          <Text><span style="color: #9e9e9e;">No completed tasks</span></Text>
        </Item>
      </List>
    {:else}
      <List twoLine checklist>
        {#each completedTasks as task}
          <Item>
            <Text>
              <PrimaryText>{task.name}</PrimaryText>
                <SecondaryText><span class="mdc-typography--body2">{task.assigned_to}</span></SecondaryText>
            </Text>
            <Meta>
              <Checkbox on:click="{toggleCompleted(task)}" bind:checked={task.completed}/>
            </Meta>
          </Item>
        {/each}
      </List>
    {/if}
  <div class="divider verified mdc-typography--overline">
    Verified
  </div>
    {#if verifiedTasks.length == 0}
      <List>
        <Item>
          <Text><span style="color: #9e9e9e;">No verified tasks</span></Text>
        </Item>
      </List>
    {:else}
      <List twoLine>
        {#each verifiedTasks as task}
          <Item>
            <Text>
              <PrimaryText>{task.name}</PrimaryText>
                <SecondaryText><span class="mdc-typography--body2">{task.assigned_to}</span></SecondaryText>
            </Text>
          </Item>
        {/each}
      </List>
    {/if}
  </div>
</div>
{/if}

<script>
  import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
  import { navigate } from "svelte-routing";
  import List, {Group, Subheader, Meta, Label, Item, Text, PrimaryText, SecondaryText} from '@smui/list';
  import IconButton from '@smui/icon-button';
  import Checkbox from '@smui/checkbox';

  import Tab from '@smui/tab';
  import TabBar from '@smui/tab-bar';
  
  import { jwt_store } from '../store.js';
  import { getTasks, toggleCompletedTask } from '../services/TaskService.js';

  import { onMount } from 'svelte';

  let selectedCheckbox = "";
  let active = "Active";
  let jwt = $jwt_store;
  let tasks = [];

  $: showActive = (active == "Active");
  $: showCompleted = (active == "Completed");

  $: activeTasks = tasks.filter(task => !task.completed);
  $: completedTasks = tasks.filter(task => task.completed);
  $: verifiedTasks = tasks.filter(task => task.verified);

  onMount(async () => {
    // Retrive the tasks
    tasks = await getTasks(jwt);
	});

  async function toggleCompleted(task) {
    console.log(task);
    toggleCompletedTask(jwt, task);
  }

  async function onLogout() {
    // Clear the store
    jwt_store.update(old_jwt => "");

    // Clear localStorage
    let storage = window.localStorage;
    storage.clear();

    navigate("/", { replace: true});
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

  .active-tab {
    padding-top: 6rem;
  }

  .completed-tab {
    padding-top: 6.5rem;
  }

  .divider {
    color: white;
    position: sticky;
    top: 6.5rem;
    height: 2rem;

    padding-left: 1rem;

    display: flex;
    flex-direction: column;
    justify-content: center;

    z-index: 1000;
  }

  .waiting {
    background-color: #ffb300;
  }
  
  .verified {
    background-color: #00897b;
  }
</style>

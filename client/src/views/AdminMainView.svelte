<TopAppBar variant="fixed" color="primary" class="fixed-top-bar">
  <Row>
    <Section>
      {#if !showTaskView}
      <Title>Technician List</Title>
      {/if}
      {#if showTaskView}
      <IconButton class="material-icons">keyboard_backspace</IconButton>
        <Title>{currentActiveUser}'s Tasks</Title>
      {/if}
    </Section>
    <Section align="end" toolbar>
      <IconButton class="material-icons" aria-label="Logout">exit_to_app</IconButton>
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
    <List twoLine>
      {#each users as user}
        <Item>
          <Text>
            <PrimaryText>{user.name}</PrimaryText>
            <SecondaryText><span class="mdc-typography--body2">{user.assigned_task_count} tasks left</span></SecondaryText>
          </Text>
        </Item>
      {/each}
    </List>
  </div>
{/if}

{#if showTaskView}
  {#if showCompleted}
    <div class="tab-view completed-tab">
      <List checklist twoLine>
        {#each tasks as task}
          <Item>
            <Text>
              <PrimaryText>{task.name}</PrimaryText>
              <SecondaryText><span class="mdc-typography--body2">{task.assigned_by}</span></SecondaryText>
            </Text>
              <Meta>
                <Checkbox bind:group={selectedCheckbox} value="{task.verified}" />
              </Meta>
          </Item>
        {/each}
      </List>
    </div>
  {/if}

  {#if showActive}
    <div class="tab-view completed-tab">
      <List twoLine>
        {#each tasks as task}
          <Item>
            <Text>
              <PrimaryText>{task.name}</PrimaryText>
              <SecondaryText><span class="mdc-typography--body2">{task.assigned_by}</span></SecondaryText>
            </Text>
          </Item>
        {/each}
      </List>
    </div>
  {/if}

  {#if showVerified}
    <div class="tab-view verified-tab">
      <List checklist twoLine>
        {#each tasks as task}
          <Item>
            <Text>
              <PrimaryText>{task.name}</PrimaryText>
              <SecondaryText><span class="mdc-typography--body2">{task.assigned_by}</span></SecondaryText>
            </Text>
              <Meta>
                <Checkbox bind:group={selectedCheckbox} value="{task.verified}" />
              </Meta>
          </Item>
        {/each}
      </List>
    </div>
  {/if}
{/if}

<script>
  import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
  import List, {Group, Subheader, Meta, Label, Item, Text, PrimaryText, SecondaryText} from '@smui/list';
  import IconButton from '@smui/icon-button';
  import Checkbox from '@smui/checkbox';

  import Tab from '@smui/tab';
  import TabBar from '@smui/tab-bar';

  let showTaskView = true;
  let active = "Completed";
  let currentActiveUser = "Sudharshan";

  let selectedCheckbox = "";
  
  $: showVerified = (active == "Verified");
  $: showActive = (active == "Active");
  $: showCompleted = (active == "Completed");
  
  let tasks = [
    {
      name: "Fix tank 1",
      assigned_by: "person 1",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 5",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
      completed: false,
      verified: false,
    }
  ];

  let users = [
    {
      name: "LCP Sundaramahalingam Sudharshan",
      assigned_task_count: 5
    },
    {
      name: "LCP Sanil Kuman Ashwin",
      assigned_task_count: 3
    },
    {
      name: "PTE Derwin Yan",
      assigned_task_count: 10
    }
  ];
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

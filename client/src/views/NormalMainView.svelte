 <TopAppBar variant="fixed" color="primary" class="fixed-top-bar">
  <Row>
    <Section>
      <Title>Your Tasks</Title>
    </Section>
    <Section align="end" toolbar>
      <IconButton class="material-icons" aria-label="Logout">exit_to_app</IconButton>
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

{#if showCompleted}
<div class="completed-tab">
  <div>
    <div class="divider waiting mdc-typography--overline">
      Pending verification
    </div>
    <List twoLine checklist>
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
  <div class="divider verified mdc-typography--overline">
    Verified
  </div>
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
</div>
{/if}

<script>
  import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
  import List, {Group, Subheader, Meta, Label, Item, Text, PrimaryText, SecondaryText} from '@smui/list';
  import IconButton from '@smui/icon-button';
  import Checkbox from '@smui/checkbox';

  import Tab from '@smui/tab';
  import TabBar from '@smui/tab-bar';

  let selectedCheckbox = "";
  let active = "Active";
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
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
      completed: false,
      verified: false,
    },
    {
      name: "Fix tank 2",
      assigned_by: "person 3",
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

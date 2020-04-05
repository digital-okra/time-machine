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

<List checklist style="padding-top: 7rem;">
  {#each tasks as task}
    <Item>
      <Label>{task.name}</Label>
        <Meta>
          <Checkbox bind:group={selectedCheckbox} value="{task.verified}" />
        </Meta>
    </Item>
  {/each}
</List>

<script>
  import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
  import List, {Meta, Label, Item} from '@smui/list';
  import IconButton from '@smui/icon-button';
  import Checkbox from '@smui/checkbox';

  import Tab from '@smui/tab';
  import TabBar from '@smui/tab-bar';

  let selectedCheckbox = "";
  let active = "Active"
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
</style>

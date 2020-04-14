
<TopAppBar variant="fixed" color="primary" class="fixed-top-bar">
  <Row>
    <Section>
      <Title>Technician List</Title>
    </Section>
    <Section align="end" toolbar>
      <IconButton class="material-icons" aria-label="Logout" on:click={onLogout}>exit_to_app</IconButton>
    </Section>
  </Row>  
</TopAppBar>

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

<script>
  import TopAppBar, {Row, Section, Title} from '@smui/top-app-bar';
  import List, {Item, Text} from '@smui/list';
  import IconButton from '@smui/icon-button';
  import { navigate } from "svelte-routing";
  
  import { getAllUsers } from '../services/UserService.js';

  import { onMount } from 'svelte';

  let jwt;
  let userid;
  let users = [];
  
  onMount(async () => {
    let storage = window.localStorage;
    jwt = storage.getItem("jwt");
    userid = storage.getItem("id");

    if(jwt == null || userid == null) {
      navigate("/", { replace: true });
    }

    users = await getAllUsers(jwt)

    if(users == null) {
      tasks = [];
    }
  });

  function viewUserTasks(user) {
    // Navigate to AdminTaskView and pass in the user
    navigate(`/admin/user/${user.id}`, { replace: true });
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
</style>

import { baseUrl } from './HttpService.js';

export async function getTasks(jwt) {
  const rawResponse = await fetch(`${baseUrl}/tasks`, {
    method: 'GET',
    headers: {
      'Accept': 'application/json',
      'Authorization': `Bearer ${jwt}`
    }
  });

  if(!rawResponse.ok) {
    throw Error(rawResponse.statusText);
  }

  // Return the User object
  let tasks = await rawResponse.json();
  return tasks;
}

export async function createTask(jwt, name, assignedTo) {
  // Create json
  let req = JSON.stringify({
    name,
    assigned_to: assignedTo
  }); 
  
  const rawResponse = await fetch(`${baseUrl}/tasks`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${jwt}`
    },
    body: req
  });

  if(!rawResponse.ok) {
    throw Error(rawResponse.statusText);
  }

  return await rawResponse.text()
}

export async function deleteTask(jwt, id) {
  // Create json
  let req = JSON.stringify({
    id
  }); 
  
  const rawResponse = await fetch(`${baseUrl}/tasks`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${jwt}`
    },
    body: req
  });

  if(!rawResponse.ok) {
    throw Error(rawResponse.statusText);
  }

  return await rawResponse.text()
}

export async function toggleVerifyTask(jwt, task, userid) {
  // Create json
  let req = JSON.stringify({
    id: task.id,
    name: task.name,
    assigned_to: task.assigned_to,
    assigned_by: task.assigned_by,
    completed: task.completed,
    verified: !task.verified,
    verified_by: userid
  }); 

  const rawResponse = await fetch(`${baseUrl}/tasks`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${jwt}`
    },
    body: req
  });

  if(!rawResponse.ok) {
    throw Error(rawResponse.statusText);
  }

  return await rawResponse.text()
}

export async function toggleCompletedTask(jwt, task) {
  // Create json
  let req = JSON.stringify({
    id: task.id,
    name: task.name,
    assigned_to: task.assigned_to,
    assigned_by: task.assigned_by,
    completed: !task.completed,
    verified: task.verified,
    verfied_by: task.verified_by
  }); 

  const rawResponse = await fetch(`${baseUrl}/tasks`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${jwt}`
    },
    body: req
  });

  if(!rawResponse.ok) {
    throw Error(rawResponse.statusText);
  }

  return await rawResponse.text()
}

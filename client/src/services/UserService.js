import baseUrl from './HttpService.js';

export async function getSelf(jwt) {
  const rawResponse = await fetch(`${baseUrl}/users/self`, {
    method: 'GET',
    headers: {
      'Accept': 'application/json',
      'Authorization': `Bearer ${jwt}`
    }
  });

  if(!rawResponse.ok) {
    throw Error(response.statusText);
  }

  // Return the User object
  let user = await response.json();
  console.log(user);
  return user;
}

export async function getAllUsers(jwt) {
  const rawResponse = await fetch(`${baseUrl}/users`, {
    method: 'GET',
    headers: {
      'Accept': 'application/json',
      'Authorization': `Bearer ${jwt}`
    }
  });

  if(!rawResponse.ok) {
    throw Error(response.statusText);
  }

  // Return the User object
  let users = await response.json();
  console.log(users);
  return users;
}

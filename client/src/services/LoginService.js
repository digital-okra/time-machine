import { baseUrl } from './HttpService.js';

export async function loginUser(username, password) {
  // Create json
  let req = JSON.stringify({
    username: username,
    password: password
  });

  console.log(req)

  const rawResponse = await fetch(`${baseUrl}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: req
  });

  console.log(rawResponse);

  if(!rawResponse.ok) {
    throw rawResponse.status;
  }

  // Return the text which is a JWT
  return await rawResponse.text();
}

export async function registerUser(username, password, type, amb, depot, platoon, section, man, name) {
  // Create json
  let req = JSON.stringify({
    username,
    password,
    type,
    amb,
    depot,
    platoon,
    section,
    man,
    name
  });
  
  const rawResponse = await fetch(`${baseUrl}/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: req
  });

  if(!rawResponse.ok) {
    throw Error(response.statusText);
  }

  // Return the text which is a JWT
  return await rawResponse.text() 
}

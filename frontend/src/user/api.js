import * as api from '../api';

export async function put(user) {
  const res = await fetch(api.withHost("/users"), {
    method: "PUT",
    body: JSON.stringify(user),
  }).then((r) => r.json());

  return res;
}

export async function get(user) {
  return fetch(api.withHost(`/users/${user.email()}`)).then((r) =>
    r.json()
  );
}

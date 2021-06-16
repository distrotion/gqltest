mutation {
  createUser(input: {username: "user1", password: "123"})
}

mutation {
  login(input: {username: "user1", password: "123"})
}

query {
	links{
    title
    address,
    user{
      name
    }
  }
}
# A much more sophisticated, true to the original Twitter clone implementing much more of the core features

## Description

Full stack chat application built using Golang and Postgres in the backend, and Javascript, HTML and CSS in the frontend. Current features include:

 - Users are able to chat in a public chat and private chat groups.
 - Global messages are able to be deleted.
 - Message history is tracked.
 - Group chat and single chat creation.
 - Add other users to specific group chats.

 There's more to come!

 ## Installation

```
   cd api
   go build 
   go mod vender //if you desire to have a folder including the dependencies, otherwise ignore
   .\fresh //to restart server on change or
   go run main.go //to run the server normally
```

  ## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

  ## License
[MIT](https://choosealicense.com/licenses/mit/)
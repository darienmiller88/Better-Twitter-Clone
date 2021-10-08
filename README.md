# Better Twitter Clone!
  ![](https://img.shields.io/badge/made%20by-DarienMiller-blue)
  ![](https://img.shields.io/badge/Golang-1.17-yellow)
  
## Description

Full stack chat application built using Golang and Postgres in the backend, and VueJS and Bootswatch in the frontend. This is a much more sophisticated, true to the original Twitter clone implementing much more of the core features. The functionality I wish to add later on include:

 - JWT + Google OAuth2 authentication.
 - The ability to for users to tweet.
 - The ability for users to like, retweet and quote tweet other users tweets.
 - The ability for users to reply to other tweets.
 - A Private messaging system.
 - Use News API to display current news.

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
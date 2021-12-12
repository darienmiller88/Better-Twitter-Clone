# Better Twitter Clone!
  ![](https://img.shields.io/badge/made%20by-DarienMiller-blue)
  ![](https://img.shields.io/badge/Golang-1.17-yellow)
  ![](https://img.shields.io/badge/build-passing-green)
  ![](https://img.shields.io/badge/Netlify-fail-red)

## Description

Full stack chat application built using Golang and Postgres in the backend, and React in the frontend. This is a much more sophisticated, true to the original Twitter clone implementing much more of the core features. The functionality I wish to add later on include:

 - JWT + Google OAuth2 authentication.
 - The ability to for users to tweet.
 - The ability for users to like, retweet and quote tweet other users tweets.
 - The ability for users to reply to other tweets.
 - A Private messaging system.
 - Use News API to display current news.

 There's more to come!

 ### Built With

* [Gin](https://github.com/gin-gonic/gin)
* [Gorm](http://gorm.io/index.html)
* [React](https://reactjs.org/)
* [PostgreSQL](https://www.postgresql.org/)
* [Netlify](https://bit.ly/3q4pcJz)

 ### Requirements
* Clone the repository using `git clone https://github.com/darienmiller88/Better-Twitter-Clone`
* Migrate the necessary information to your local `.env` as described in the `.env_sample` file
* Run `go build` to create a`Better-Twitter-Clone.exe` file, and then run`.\Better-Twitter-Clone` to run the executable. If an executable is not needed, simply input `go run main.go` instead, or `.\fresh` to enable a server restart on change.
* Run `cd client` to change directories to the client folder, and run `npm run start` to start the react server. Happy coding!

  ## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

  ## License
[MIT](https://choosealicense.com/licenses/mit/)
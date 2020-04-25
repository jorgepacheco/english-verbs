# english-verbs
English verb api


# Build

```
docker build -t jorgepacheco/server-english-verbs .
```

# Heroku

https://english-verbs.herokuapp.com/ | https://git.heroku.com/english-verbs.git

```
λ heroku login
heroku: Press any key to open up the browser to login or q to exit: 
Opening browser to https://cli-auth.heroku.com/auth/cli/browser/256e6904-b09b-4cf7-b6e5-86e9e2e1a94e
Logging in... done
Logged in as jorge.pacheco12@gmail.com
➜ english-verbs master ✓ 
λ heroku create english-verbs
Creating ⬢ english-verbs... done
https://english-verbs.herokuapp.com/ | https://git.heroku.com/english-verbs.git
`
```
heroku container:login

docker tag jorgepacheco/server-english-verbs registry.heroku.com/english-verbs/web

docker push registry.heroku.com/english-verbs/web
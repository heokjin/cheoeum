###build
go build

---
"# cheoeum" 

docker build -t heokjin/cheoeum .  <br>
docker run -p 80:80 -it --rm --name cheoeum cheoeum:latest
<p>
docker login <br>
docker push heokjin/cheoeum:latest
위에는 일반 docker 빌드 아래는 heroku용 빌
###

# go-getting-started

A barebones Go app, which can easily be deployed to Heroku.

This application supports the [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) article - check it out.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.12 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ git clone https://github.com/heroku/go-getting-started.git
$ cd go-getting-started
$ go build -o bin/cheoeum -v . # or `go build -o bin/cheoeum.exe -v .` in git bash
github.com/mattn/go-colorable
gopkg.in/bluesuncorp/validator.v5
golang.org/x/net/context
github.com/heroku/x/hmetrics
github.com/gin-gonic/gin/render
github.com/manucorporat/sse
github.com/heroku/x/hmetrics/onload
github.com/gin-gonic/gin/binding
github.com/gin-gonic/gin
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku
처음프로젝트를 만들면 아래메뉴얼을 참고하자
https://dashboard.heroku.com/apps/protected-reaches-14112/deploy/heroku-git
```sh
heroku login
$ heroku create
$ git push heroku master
$ heroku open

Create a new Git repository
Initialize a git repository in a new or existing directory

$ cd my-project/
$ git init
$ heroku git:remote -a cheoeum
```
```
$ git add .
$ git commit -am "make it better"
$ git push heroku master
```
or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Documentation

For more information about using Go on Heroku, see these Dev Center articles:

- [Go on Heroku](https://devcenter.heroku.com/categories/go)

#데이터베이스 만들기
해당 프로젝트 안에서 실행해야 한다.<br>
heroku addons:create heroku-postgresql:hobby-dev<br>
heroku addons:docs heroku-postgresql<br>
This database is empty. If upgrading, you can transfer data from another database with pg:copy <br>
heroku pg:info <br>

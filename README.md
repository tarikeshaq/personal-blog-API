# Go REST API for my personal blog

My first Golang project, a REST API to support my personal blog (Under development :wink: ) 


## Getting Started

Although this was meant as project for myself to use on my personal blog, feel free to clone it as boilerplate code
for a MUX-MongoDB project.

### Prerequisites

Of course, you'll need to have golang installed https://golang.org/dl/ (I use version 1.12.1, if you're using a different version
please change the "go-version" in the Gopkg.toml" . Additionally, for dependency management,
this project uses dep https://github.com/golang/dep

```
$ go get github.com/tarikeshaq/personal-blog-API
```

If you're unfamiliar with Go, this will download the repository in the GOPATH/src
Find where your GOPATH variable points, then navigate to src/github.com/tarikeshaq/personal-blog-API

NOTE: If you plan to deploy your application, you may want to fork my repository first, then clone that to avoid complication with 
 the user name :no_mouth:


### Installing

With dep installed, you'll be able to run in no time!
navigate to the cloned repo, and pull the dependencies using dep ensure

```
$ dep ensure
```

This will make sure all the dependencies are correct (I have my vendors as a part of the repository, so dep ensure is more of a precausion)

Now, you'll have the project ready, but you do need to set some Environment (OS) variables

```
* "PORT": Environment variable to point to the port where you'd like the server to run
* "MONGO_HOST": Your mongoDB host, find the URI for your database and the host would be after the '@' and before the '/'
* "MONGO_USERNAME": Your mongoDB database username
* "MONGO_PASSWORD": Your mongoDB database password for the above username
* "MONGO_DATABASE": The name of the database (I have mine as blogs)
```


With that the application should be ready for you take on a ride!!
Make sure that your MongoDB server is running, then run the following:
```
$ go run app.go
```

With that the server should be running and ready to be queried

## API End points
```
* GET /blogs : Returns all the blog posts in the database
* GET /blogs/{blogId} : Returns one post from the database with id = blogID
* POST /blogs : Adds a blog post to the database
* DELETE /blogs/{blogId} : Deletes a blog post from the database with id = blogID
```
To check the structure of a blog post, check the /models directory


## Deployment

I have my application deployed on Heroku, (it plays nice with dep, which I found difficulties with gcloud :sweat_smile: )

You can deploy on heroku by adding your repository from the heroku site, or using the CLI, just make sure all the 
dependencies and the Gopkg.toml are correct.

## Built With

* [Mux](https://github.com/gorilla/mux) - Routing management
* [Dep](https://github.com/golang/dep) - Dependency Management
* [Mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - For interacting with the database

## Contributing

This a personal project, although, if you feel like you see a way to improve it, feel free to create a PR :) 

## Authors

* **Tarik Eshaq** - *Project Owner*


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Mad respect to everyone who's code I sto.. got influenced by :angel: 
* Go has been an amazing techology to learn, definitly will build more stuff using it! 


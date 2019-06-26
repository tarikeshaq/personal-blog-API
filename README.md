# Go REST API for my personal blog

My first Golang project, a REST API to support my personal blog (Under development :wink: ) 


## Getting Started

Although this was meant as project for myself to use on my personal blog, feel free to clone it as boilerplate code
for a MUX-MongoDB project.

### Prerequisites

Of course, you'll need to have golang installed https://golang.org/dl/ (I use version 1.12.1, if you're using a different version
please change the "go-version" in the go.mod file. Additionally, for dependency management,
this project uses go modules!

```
$ git clone https://github.com/tarikeshaq/personal-blog-API.git
```

Since this project used go modules, you don't need to place it in your GOPATH :smile: Simply clone the project and run it!


NOTE: You'd need Go 1.11+ in order to run this using go modules


### Installing

With the repo downloaded, you'll be able to run in no time!
navigate to the cloned repo, and pull the dependencies by running go build

```
$ go build
```

This will create the binary and make sure all the dependencies are correct (Will download any dependencies you don't have, yay go modules)

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
$ go run main.go
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

I have my application deployed on Heroku

You can deploy on heroku by adding your repository from the heroku site, or using the CLI

## Built With

* [Mux](https://github.com/gorilla/mux) - Routing management
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


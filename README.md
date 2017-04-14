[![Build Status](https://travis-ci.org/husniadil/gsm.svg?branch=master)](https://travis-ci.org/husniadil/gsm)

# GSMArena API

GSMArena phone specification and finder.
This project is still in early development.

The API basically reads from GSMArena website and results JSON data. 

#### Requirements:

- [Go](https://golang.org/) 1.7.x or up
- [GNU Make](https://www.gnu.org/software/make/)


#### Quick Start
```bash
# setup dependencies (you should run this once before compiling/running the app)
make setup

# compile and run the app
make serve
```


#### Available commands

```
Command              : Description
-------------------- : --------------------------------------------
make setup           : Install all necessary dependencies
make build           : Generate production build for current OS
make test            : Run tests
make serve           : Run the app locally
make image           : Create docker image
make run-image       : Run the app on docker
make clean           : Remove dependencies and compiled files
```


#### Implemented Features
- [x] Get all brands
- [x] Get devices by brand
- [x] Get device specification
- [ ] Find devices by keyword
- [ ] Find devices by advanced filters


#### Online Demo
- [Get all brands](https://gsmarena-obrisiswox.now.sh/brands)
- [Get Samsung devices](https://gsmarena-obrisiswox.now.sh/devices/samsung-phones-9)
- [Get Samsung Galaxy S8 specification](https://gsmarena-obrisiswox.now.sh/specs/samsung_galaxy_s8-8161)


#### Try on Local
* **Using Postman**

  [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/555bfc6fbe025e73641e)

* **Using Swagger UI**

  This API doesn't have embedded Swagger UI. You may need to spawn Swagger UI docker container and use provided `swagger.json` file in this repo.
Detail of setting up Swagger UI can be [read from here](https://github.com/swagger-api/swagger-ui).

---


#### Contribution
If you want to contribute code, please consider creating a pull request after forking this project.

If you have any questions, found bugs, or ideas, please open a ticket (issue).


---


#### Credits
- [github.com/constabulary/gb](github.com/constabulary/gb)
- [github.com/PuerkitoBio/goquery](github.com/PuerkitoBio/goquery)
- [github.com/fatih/color](github.com/fatih/color)
- [github.com/gorilla/mux](github.com/gorilla/mux)
- [github.com/justinas/alice](github.com/justinas/alice)
- [github.com/rainycape/unidecode](github.com/rainycape/unidecode)
- [github.com/azer/snakecase](github.com/azer/snakecase)


---


#### License

This project is licensed under the terms of the [MIT license](https://husniadil.mit-license.org/).

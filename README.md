# GSMArena API

GSMArena phone specification and finder.
This project is still in early development.

The API basically reads from GSMArena website and results JSON data. 

#### Requirements:

- [Go](https://golang.org/) - Preferably version 1.8.x
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


#### API Docs
- [ ] Create API documentation


#### Demo
Demo server built on Apr 9, 2017 at 12:15am. Not all features has been implemented yet.
- [Get all brands](https://gsm-wcavkgnhzr.now.sh/brands)
- [Get Samsung devices](https://gsm-wcavkgnhzr.now.sh/devices/samsung-phones-9)
- [Get Samsung Galaxy S8 specification](https://gsm-wcavkgnhzr.now.sh/specs/samsung_galaxy_s8-8161)


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

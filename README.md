# GSMArena API

GSMArena phone specification and finder.
This project is still in early development.

The API basically reads from GSMArena website and results JSON data. 

#### Requirements:

- [Go](https://golang.org/) - Preferably version 1.8.x
- [GNU Make](https://www.gnu.org/software/make/)


#### Setup
```bash
# setup (one time only)
make setup
```

#### Build

```bash
# build binary
make build
```

#### Run
```bash
# run the binary
bin/gsm
```

#### Test
```bash
# test all functionality
make test
```

#### Implemented Features
- [x] Get all brands
- [x] Get devices by brand
- [ ] Get device specification
- [ ] Find devices by keyword
- [ ] Find devices by advanced filters


#### API Docs
- [ ] Create API documentation


---

#### Contribution
Please create a pull request after forking this project.

If you have any questions, found bugs, or ideas, please open a ticket (issue).


#### License

This project is licensed under the terms of the [MIT license](https://husniadil.mit-license.org/).

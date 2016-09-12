#Golang Imageboard - WIP 
[![Build Status](https://travis-ci.org/ewhal/gchan.svg?branch=master)](https://travis-ci.org/ewhal/gchan) [![GoDoc](https://godoc.org/github.com/ewhal/gchan?status.svg)](https://godoc.org/github.com/ewhal/gchan) [![Go Report Card](https://goreportcard.com/badge/github.com/ewhal/gchan)](https://goreportcard.com/report/github.com/ewhal/gchan) [![MIT
licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/ewhal/gchan/master/LICENSE.md)

#Motivation
Many implementations of imageboard software do exist but the most popular imageboard software is licensed under a partially non-free license and has a terrible and large code base.

## Getting started
### Prerequisities
* go
* mariadb

```
sudo yum install -y go mariadb-server mariadb
```

### Installing

* go get https://github.com/ewhal/gchan
* make will automatically download the dependencies for gchan
* cp config.example.json config.json
* nano config.json
* Configure port and database details

## Todo
- [ ] Posts
  - [ ] Tripcode
  - [ ] capcode
  - [ ] green text
  - [ ] red text
  - [ ] backlinks
  
- [ ] Frontend
  - [ ] thumbnails
  - [ ] Image uploading
  - [ ] webm and mp4
  - [ ] Alternative styles
  
- [ ] Javascript
  - [ ] auto thread refresh
  - [ ] captcha
  - [ ] thumbnail expansion

- [ ] Moderation tools
  - [ ] board creation
  - [ ] Admin panel
  - [ ] sticky thread
  - [ ] delete thread
  - [ ] ban and unban
  - [ ] ban appeals
  
- [ ] Documentation
  - [ ] api
  - [ ] comment code
  
## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details


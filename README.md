# About
Bookstore - REST API application written with Go

# Index
* System requirements
* Installation
* Administrating

# System requirements
### OS
* Any with x64 architecture
### RAM
* 2 Gb
### CPU
* Any 2-Core
### Software
* Go 1.17+
* `systemctl` or `server`

# Installation
* Build
```bash
    $go build ./cmd/bookstore
```

# Administrating

## Commands
* Start server
```bash
    $./cmd/bookstore
```
* Stop server by closing application (via bash or by systemctl)

## Configurations
* `./configs/app.xml` contains application main params, like host, port, environment
* `/.configs/database.xml` contains database configs, like host, name, lgoin, password

## API Documentation
* `https://documenter.getpostman.com/view/11110995/UVkvLDg9`
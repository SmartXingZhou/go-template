# go-template
Architecture template for Golang services

# Overview
This is a personal Go project template that you can use as a starting point for your project. You will need to replace the placeholder variables/values/names with your own though.   

The purpose is to scientifically organize the code structure, ensure that the project directory structure is clear, logically separated, the code is readable, and spaghetti code is avoided.

Referenced from the [Go Project Layout](https://github.com/golang-standards/project-layout) project.

# Quick start
`make dev`

# Structure introduction
## api
Service entry for protocols such as http, https, websocket, etc., may also include OpenAPI/Swagger specs, JSON schema files, protocol definition files.

## cmd
The startup file of the project, usually calling the code in the pkg、internet、web packages.
## configs
Configuration file, and read the configuration.
## docker
Docker-related configuration to distinguish multiple environments
## internal
Private application and library code.
## pkg
Code available for external reference.
## test
Test code and test data.
## tools
Supporting tools for this project.
## web
Web application specific components: static web assets, server side templates and SPAs.

# todo-app-golang

## 1. Introduction
Well, at the time of making this project, I was kind of a newbie in Golang. So the purpose was to learn how to code a API backend with Go. For other newbies in this matter, I hope that this project can help.

In addition to this project, I have started another one called "todo-app-frontend". Since I was also a newbie in Angular, that one was also meant to learn a couple of stuff. Long story short, you can run both frontend and backend codes to see how it actually looks like.

## 2. Setting up GO environment
Unlike other common languages, Go has one specific working directory which has to be defined in the environment variable $GOPATH. It should contain 3 typical folders inside: bin, pkg and src.

- bin folder contains the binary executables, not only the executables that belong to your applications but also some others which help you compile your apps in general.
- pkg folder contains the shared libraries which you import to your application and compile together.
- src folder contains all the magical source code of yours where you probably spend most of your time.

### 2.1 Linux & Mac
Go to your _~/.bash_profile_ and add the following line:

export GOPATH=_your working directory_

### 2.2 Windows
Open up your windows start menu and type _environment_. You'll probably see -> _Edit the system environment variables_. This will lead you automatically to the _Advanced_ tab on _System Properties_ window.

Click _Environment Variables_. On top you'll see _User variables for YOUR USER_.

Click _New_. Type _GOPATH_ for variable name and type _your working directory_ for variable value.

### 2.3 Useful Links
- https://golang.org/doc/gopath_code
- https://www.callicoder.com/golang-installation-setup-gopath-workspace/
- https://medium.com/rungo/working-in-go-workspace-3b0576e0534a

## 3. Requirements
### 3.1 MySql Database



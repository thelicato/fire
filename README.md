<h1 align="center">
  <br>
  <a href=""><img src="logo.png" alt="" width="300px;"></a>
  <br><b>fi</b>lter <b>re</b>solved<br>
</h1>

<p align="center">
  <a href="https://github.com/thelicato/fire/blob/main/README.md"><img src="https://img.shields.io/badge/Documentation-complete-green.svg?style=flat"></a>
  <a href="https://github.com/thelicato/fire/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-blue.svg"></a>
  <a href="https://twitter.com/intent/follow?screen_name=thelicato"><img src="https://img.shields.io/twitter/follow/thelicato?style=flat&logo=twitter"></a>
</p>

<!-- omit in toc -->
## Table of Contents
- [🛸 Overview](#-overview)
- [🛸 Usage](#-usage)
  - [Normal execution](#normal-execution)
  - [Verbose execution](#verbose-execution)
  - [Docker execution](#docker-execution)
- [🚀 Installation](#-installation)
- [🪪 License](#-license)

## 🛸 Overview

**fire** is a simple tool meant to work in a pipeline of other scripts. It takes domains on ``stdin`` and outputs them on ``stdout`` if they resolve. The inspiration for this work is the ``filter-resolved`` Golang code in [this repository](https://github.com/tomnomnom/hacks/tree/master/filter-resolved). That repo is not updated in a long time and I thought that it was time to switch to the new Golang paradigm of modules. I also added a ``Dockerfile``.

## 🛸 Usage

**Available Options:**
```
  c     The concurrency of the execution (default=20)
  v     Set this flag to run fire in VERBOSE mode
```

### Normal execution
Just pipe the domain (or the list of domains) to ``fire``:
```
echo thelicato.io | fire
```

Here is an example with multiple hosts:
```
# Don't forget to set the HOST variable!
subfinder -silent -d $HOST | fire 
```

### Verbose execution
You can also run with the ``-v`` option to have some other output. This mode is not preferred when ``fire`` is not the last script of a **pipeline** of scripts:
```
subfinder -silent -d $HOST | fire -v # This is OK

subfinder -silent -d $HOST | fire -v | sort -u # This is NOT OK because all the output will be sorted
```

### Docker execution
You can also execute ``fire`` in a Docker container if you prefer running containers:
```
echo thelicato.io | docker run -i --rm thelicato/fire
```

## 🚀 Installation
**Normal installation (and update)**:
```
go install github.com/thelicato/fire@latest
```

**Docker installation**:
```
docker pull thelicato/fire:latest
```


## 🪪 License
**fire** is an open-source and free software released under the [MIT License](https://github.com/thelicato/fire/blob/main/LICENSE).
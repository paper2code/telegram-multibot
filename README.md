# Paper2code - Modular Telegram Bot

![Paper2code feat telegram](./docs/img/paper2code_telegram.png "Paper2code feat telegram")

## Quick start


## Install

### Docker

#### Pre-requisites
If you don't have Docker/Docker-Compose check **Setup Docker** section

<details>
<summary><b>Setup Docker</b></summary>
<p>
<h3>Docker</h3>
MacOS:&nbsp;<a href="https://docs.docker.com/docker-for-mac/install/">https://docs.docker.com/docker-for-mac/install/</a><br />
Linux:&nbsp;<a href="https://docs.docker.com/install/linux/docker-ce/ubuntu/">https://docs.docker.com/install/linux/docker-ce/ubuntu/</a><br />
<hr />
<h3>Docker Compose</h3>
Linux:&nbsp;<a href="https://docs.docker.com/compose/install/">https://docs.docker.com/compose/install/</a><br />
<br />
</p>
</details>

#### (Optional) Docker-Compose with Nvidia support
There is a hack for making your gpu(s) available with docker-compose, for doing so, you will have to add the following to your /etc/docker/daemon.json file:
```json
{
    "default-runtime": "nvidia",
    "runtimes": {
        "nvidia": {
            "path": "nvidia-container-runtime",
            "runtimeArgs": []
        }
    }
}
```
and restart the docker service

#### Quick Start

```sh
mv .env-example .env # change value accordingly to your setup
docker-compose up -d
```

### Source

#### Pre-requisites
Requires Golang 1.1x and GNU Make.

```sh
make
make plugins
```

## Credits
@lucmichalski
@x0rzkov

## Contributions


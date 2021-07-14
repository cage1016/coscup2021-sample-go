<h1 align="center">Welcome to coscup2021-tw-sample-go 👋</h1>
<p>
  <a href="https://github.com/cage1016/coscup2021-tw-sample-go/blob/master/LICENSE" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/CageChung" target="_blank">
    <img alt="Twitter: CageChung" src="https://img.shields.io/twitter/follow/CageChung.svg?style=social" />
  </a>
  
</p>

[![Build](https://github.com/cage1016/coscup2021-tw-sample-go/actions/workflows/build.yml/badge.svg)](https://github.com/cage1016/coscup2021-tw-sample-go/actions/workflows/build.yml)

> [coscup 2021 tw - Google Cloud Buildpacks 剖析與實踐](https://coscup.org/2021/zh-TW/session/J3X8SE) source code

## Author

👤 **KAI CHU CHUNG**

* Website: https://kaichu.io/
* Twitter: [@CageChung](https://twitter.com/CageChung)
* Github: [@cage1016](https://github.com/cage1016)
* LinkedIn: [@kaichuchung](https://linkedin.com/in/kaichuchung)

## Usage

### build container image by pack

1. [Install Docker](https://store.docker.com/search?type=edition&offering=community)
1. [Install the pack tool (a CLI for running Buildpacks)](https://buildpacks.io/docs/install-pack/)
1. build container image by pack & gcr.io/buildpacks/builder:v1 builder
    ```bash
    # You could change container image as you want
    pack build --builder gcr.io/buildpacks/builder:v1 ghcr.io/cage1016/coscup2021-tw-sample-go
    ```
### Build contaiern image by skaffold

1. [Install skaffold](https://skaffold.dev/docs/install/)
1. build container image by skaffold
    ```bash
    skaffold build 
    ```
### Deploy to Google Cloud Run

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)


### Build Cloud Function by pack

1. build container image by pack & gcr.io/buildpacks/builder:v1 builder
    ```bash
    # You could change container image as you want
    pack build hello-coscup-2021-tw-get --builder gcr.io/buildpacks/builder:v1 --env GOOGLE_FUNCTION_TARGET=HelloCocsup2021tw -p functions
    ```
1. run container image locally
    ```bash
    docker run --rm -p 8080:8080 hello-coscup-2021-tw-get
    ```
2. curl hello-coscup-2021-tw-get
    ```bash
    curl localhost:8080
    ```

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2021 [KAI CHU CHUNG](https://github.com/cage1016).<br />
This project is [MIT](./LICENSE) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_
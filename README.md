# dokcli

Open Source API documentation in 1 file

## dokcli is an open source API Documentation written in Go, consisting of

- built-in Rapidoc web component
- support Swagger & OpenAPI Spec
- convenient way of hosting Api Documentation in 1 file

For documentation and examples, please visit <https://github.com/Shoaibashk/dokcli>
> :construction: Please keep in mind that Dokcli is still under active development :alembic: and therefore full backward compatibility is not guaranteed before reaching v1.0.0

## Overview ✨

Dokcli could be [downloaded directly as a standalone app](https://github.com/Shoaibashk/dokcli/releases) for API Documentation using Swagger & OpenAPI spec file.

## Quickstart 🚀

- Download the dokcli from [release](https://github.com/Shoaibashk/dokcli/releases).
- Store it in a folder `doc/` as a executable.
- Now, open the terminal, cd to the newly created folder location and simply type:
  - `./dokcli serve`

- You can provide server port and openAPI spec url.
  - `./dokcli serve -p 2323 -u https://petstore.swagger.io/v2/swagger.json`

## Customization 💅 (coming soon)

>❗ Dokcli uses rapidoc web component under the hood 🌺. Its still under active development.

Dokcli can be Customized using config file `config.yaml`.

You can look at the help section in `./dokcli help`.

## Support 🤗

The supported build targets by the non-cgo driver at the moment are:

```shell
darwin  arm64
darwin  amd64
linux   arm64
linux   386
linux   amd64
windows 386
windows amd64
windows arm64
```

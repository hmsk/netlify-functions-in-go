# Netlify Functions in Go

Deployable template of multiple Go lambda functions for [Netlify Functions](https://www.netlify.com/products/functions/).

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/hmsk/netlify-functions-in-go)

## File structure

```
├── Makefile
├── cmd
│   ├── goodbye
│   │   └── main.go
│   └── hello
│       └── main.go
├── dashboard
│   └── index.html
├── internal
│   └── pkg
│       └── utils
│           └── utils.go
└── netlify.toml
```

### `cmd`

Place your functions. A dir matches to an end-point: `/.netlify/functions/hello`.
Each dir should have `main.go` as `package main`.

### `internal/pkg`

Place your common package for sharing among multiple functions.
In this sample, `utils/utils.go` provides `utils.IntroductionYourself` for `hello`, `goodbye` endpoints.

### `dashboard`

This dir will be deployed as a website. The root path for your Netlify app brings visitors here.
Put HTML/assets...etc to support your functions 💪

## Development

### Build

```
$ make
```

Try to build packages and saves Go binaries into `functions` dir.

### Run tests

```
$ make test
```

## License

[MIT License](LICENSE) Copyright (c) 2020 Kengo Hamasaki

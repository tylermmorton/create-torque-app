# create-torque-app

This project is the official template project for creating a new torque app. The template is designed to get you up and running quickly.

Get started by running the official Golang tool `gonew`

```shell
$ go install golang.org/x/tools/cmd/gonew@latest
$ gonew github.com/tylermmorton/create-torque-app github.com/<your-username>/<your-app-name>
```

or simply fork this repository and clone it locally!

# The Stack
The `create-torque-app` template project is preconfigured with the following technologies:
- [torque](https://lbft.dev) - Webserver framework
- [htmx](https://htmx.org/) - Frontend framework
- [tmpl](https://github.com/tylermmorton/tmpl) - Go `html/template` compiler and renderer
- [TailwindCSS](https://tailwindcss.com/) - CSS framework

It also comes preconfigured with the following tooling:
- [Docker](https://www.docker.com/) - Container runtime
- [eslint](https://eslint.org/) - JavaScript & HTML linter
- [prettier](https://prettier.io/) - JavaScript & HTML formatter
- [Taskfile](https://taskfile.dev/) - Task runner & mini build system

# Project Structure

- `go.mod` & `go.sum` - Golang module files for managing Go dependencies.
- `package.json` & `package-lock.json` - Node module files for managing Node dependencies.

## App directories

| Directory     | Purpose |
|---------------|---------|
| `assets/`     |         |
| `elements/`   |         |
| `middleware/` |         |
| `model/`      |         |
| `routes/`     |         |
| `services/`   |         |
| `styles/`     |         |
| `templates/`  |         |
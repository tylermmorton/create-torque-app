# create-torque-app

This project is the official template project for creating a new torque app. The template includes a one-time setup script to get you up and running quickly.

Get started by clicking 'Use this template' above. 

After you've created a new repository, clone it locally into your `$GOPATH` and run the one time setup script:

```shell
git clone <your-repo-url>
cd <your-repo-name>
./one_time_setup.sh
```

The `one_time_setup.sh` script does the following:
- Changes necessary project files to match your new project name
- Resolves go dependencies via `go mod tidy`
- Install the build tool [Taskfile](https://taskfile.dev/) via `go install`
- Provides next steps for getting started
- Deletes itself

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
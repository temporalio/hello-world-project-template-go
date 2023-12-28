# hello-world-project-template-go

This is the Hello World project which has all the basic files explained in our [Hello World Tutorial](https://learn.temporal.io/getting_started/go/hello_world_in_go/).

## Instructions

Ensure you have Go 1.16 or later installed locally, and access to a Temporal Cluster for development.

The fastest way to get a development cluster running on your local machine is to use [Temporal CLI](https://docs.temporal.io/cli#install).

Temporal CLI is a tool for interacting with a Temporal Cluster from the command-line interface. It includes a self-contained distribution of the Temporal Server and [Web UI](https://docs.temporal.io/web-ui) as well.

Once you've installed Temporal CLI on your platform of choice and added it to your ```PATH```, open a new Terminal window and run the following command:

```temporal server start-dev```

*  The Temporal Server will be available on ```localhost:7233```.
*  The Temporal Web UI will be available at ```http://localhost:8233```.

Next, clone this repository and switch to the cloned directory:

```bash
git clone https://github.com/temporalio/hello-world-project-template-go
cd hello-world-project-template-go
```

Now you can run the worker and starter included in the project.

```bash
go run worker/main.go
go run start/main.go
```

If you have [`nodemon`](https://nodemon.io/) installed, you can automatically reload when you change any files: `nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run worker/main.go`

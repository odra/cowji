# Kubeji

This is an attempt to run [Koji](https://pagure.io/koji) in [Kubernetes](https://kubernetes.io),
as native as possible.

It will focus on using Kubernetes based frameworks and tools, such as kubebuilder, operators and so on.

It may port some upstream Koji libraries to `go` but just enough to run it in Kubernetes.

## Project Structure

This is a single repository project which will contains several projects
that will depend on each other.

* `examples`: code samples and examples
* `pkg`: this where reusable APIs packages will be placed;
* `pkg/koji-client`: koji client written in go to be used by other programs, such as an operator.

## License

[MIT](./LICENSE)

# Cowji

This is an attempt to run [Koji](https://pagure.io/koji)
and related tools in a different way.

It may port some upstream Koji libraries to `go`
but just enough to enable specific third-party
tools integrations.

## Project Structure

This is a single repository project which will
contain several projects that will depend on
each other:

* `examples`: code samples and examples;
* `images`: container images files;
* `kojiclient`: a "just enough" implementation of akoji client library written in go;
* `koji-operator`: a kubernetes operator that deploys koji.

## License

[MIT](./LICENSE)

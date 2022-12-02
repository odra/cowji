name: Go Unit Tests

on: [pull_request]

jobs:
  build:

    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Setup Go
        uses: actions/setup-go@v3

        with:
          go-version: 1.18

      - name: Unit Tests
        run: go test -v ./pkg/.../

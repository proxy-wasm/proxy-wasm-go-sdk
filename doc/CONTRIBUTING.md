# Contributing

We welcome contributions from the community. Please read the following guidelines carefully to maximize the chances of your PR being merged.

## Coding Style

The code is linted using a stringent golang-ci. To run this linter (and a few others) use run `make check`. To format your files, you can run `make format`.

## Running tests

```
# Run local tests without running envoy processes.
make test

# Run all e2e tests.
# This requires you to have Envoy binary locally.
make test.e2e

# Run e2e tests for a specific example.
# This requires you to have Envoy binary locally.
make test.e2e.single name=helloworld
```

## Contributor License Agreement

Contributions to this project must be accompanied by a Contributor License
Agreement. You (or your employer) retain the copyright to your contribution;
this simply gives us permission to use and redistribute your contributions as
part of the project. Head over to <https://cla.developers.google.com/> to see
your current agreements on file or to sign a new one.

You generally only need to submit a CLA once, so if you've already submitted one
(even if it was for a different project), you probably don't need to do it
again.

## Code reviews

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

## Community Guidelines

This project follows [Google's Open Source Community
Guidelines](https://opensource.google/conduct/).

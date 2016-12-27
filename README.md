# amelia

[![license](https://img.shields.io/github/license/hackebrot/amelia.svg)][amelia]
[![amelia on GitHub](https://img.shields.io/github/forks/hackebrot/amelia.svg?style=social&label=Fork)][amelia]

Create GitHub Gists from your CLI

## Installation

``go get github.com/hackebrot/amelia``


## GitHub API

Connecting to the [GitHub API][github-api] with **amelia** requires
authentication via [oauth-tokens][github-auth].

Please [generate tokens][github-tokens] (with scope "Create gists") and set the
following environment variables:

- ``AMELIA_USERNAME`` - your GitHub user name
- ``AMELIA_TOKEN`` - your personal GitHub API access token


## Usage

With auth environment variables in place you can now run **amelia**.

```
-description string
    A description of the gist. (default "Amelia Gist")

-public
    Indicates whether the gist is public. (default false)
```

To create a public Gist with description "Hello World" and a file "helloworld.md" run the following:

    $ amelia -public -description "Hello World" helloworld.md

To create a secret Gist with the default description and multiple files run the following:

    $ amelia file1.txt file2.txt


## License

Distributed under the terms of the [MIT License][MIT], amelia is free and open
source software.

[amelia]: https://github.com/hackebrot/amelia
[github-api]: https://developer.github.com/v3/
[github-auth]: https://developer.github.com/v3/auth/#via-oauth-tokens
[github-tokens]: https://github.com/settings/tokens
[MIT]: LICENSE

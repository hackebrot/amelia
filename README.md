# amelia

Create [GitHub gists][gists] from your CLI. 📄

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

With auth environment variables in place you can now use **amelia**. 💻

```
  -d string
        description of the gist (default "Amelia Gist")
  -description string
        description of the gist (default "Amelia Gist")

  -p	indicates whether the gist is public
  -public
        indicates whether the gist is public
```

To create a public gist with description "Hello World" and a file
"helloworld.md" run the following:

```text
$ amelia -public -description "Hello World" helloworld.md
```

To create a secret gist with the default description and multiple files run the following:

```text
$ amelia file1.txt file2.txt
```


## Community

Would you like to contribute to **amelia**? You're awesome! 😃

Please check out the [good first issue][good first issue] label for tasks, that
are good candidates for your first contribution to **amelia**. Your
contributions are greatly appreciated! Every little bit helps, and credit will
always be given! Find out who has already contributed to **amelia**
[here][community]!  🌍🌏🌎

Please note that **amelia** is released with a [Contributor Code of
Conduct][code of conduct]. By participating in this project you agree to abide
by its terms.


## License

Distributed under the terms of the [MIT License][MIT], **amelia** is free and
open source software.


## Project name

> A Saboscrivner can write about food so accurately that people get the
> sensation of taste when they read about the food. Amelia Mintz is a
> Saboscrivner.

[Wikipedia: Chew Comic Book Series][wikipedia] 📚

[MIT]: https://github.com/hackebrot/amelia/blob/master/LICENSE
[code of conduct]: https://github.com/hackebrot/amelia/blob/master/.github/CODE_OF_CONDUCT.md
[community]: https://github.com/hackebrot/amelia/blob/master/CONTRIBUTORS.md
[gists]: https://help.github.com/en/articles/about-gists
[github-api]: https://developer.github.com/v3/
[github-auth]: https://developer.github.com/v3/auth/#via-oauth-tokens
[github-tokens]: https://github.com/settings/tokens
[good first issue]: https://github.com/hackebrot/amelia/labels/good%20first%20issue
[wikipedia]: https://en.wikipedia.org/wiki/Chew_(comics)#Saboscrivner

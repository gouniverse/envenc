# envenc 

<a href="https://gitpod.io/#https://github.com/gouniverse/envenc" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

![tests](https://github.com/gouniverse/envenc/workflows/tests/badge.svg)

## Description

Secures your .env files with a password.

Works similarly to ansible-vault.

## Installation

- Download the binary for your platform from the latest release

- You may install it globally, or use as standalone executable

## Example Usage:

- To create a new vault file

```bash
$> ./envenc init .env.vault
```

- To set a new key-value pair

```bash
$> ./envenc key-set .env.vault
```

- To list all key-value pairs
```bash
$> ./envenc key-list .env.vault
```

- To remove a key-value pair
```bash
$> ./envenc key-remove .env.vault
```

- To obfuscate a string
```bash
$> ./envenc obfuscate
```

- To deobfuscate a string
```bash
$> ./envenc deobfuscate
```


## TODO

- https://github.com/burrowers/garble

- https://github.com/marketplace/actions/go-release-binaries

# nix-gitignore - filterSource by .gitignore

**status: experimental**

Nix v1.11.9 introduced a new builtin called `builtins.exec` which allows to run code during the evaluation phase. This project's goal is to play with it while trying to solve a problem at the same time:

When working in a nixified project, it can become quite cumbersome to maintain the filterSource matches. The matching capabilities are quite crude and don't support globbing and so the matches can be quite inexact. When the project changes, the filterSource rules have to be updated as well. What if filterSource could re-use the .gitignore of the project instead?

## Usage

Assuming filterGitignore is part of the top-level:

```nix
{ stdenv, filterGitignore }:
stdenv.mkDerivation {
    name = "my-package";
    src = filterGitignore ./.gitignore ./.;
    # ...
}
```

To build the package, the `--option allow-unsafe-native-code-during-evaluation true` flag is required.

## About `builtins.exec`

Nix v1.11.9 introduced a new builtin called `builtins.exec` which allows to run
code during the evaluation phase. The idea is that it allows to do cool things like dynamically calculate the sha256 of a git repo.

Because this option is quite dangerous, it is hidden behind a flag.

```
$ nix-repl
Welcome to Nix version 1.11.9. Type :? for help.

nix-repl> builtins.exec
error: attribute ‘exec’ missing, at (string):1:1
```

```
$ nix-repl --option allow-unsafe-native-code-during-evaluation true
Welcome to Nix version 1.11.9. Type :? for help.

nix-repl> builtins.exec
«primop»
```

### Usage

The function takes a list of arguments and parses the returned stdout output. The stderr isn't intercepted.

```
nix-repl> builtins.exec ["echo" "[1 2 3]"]
[ 1 2 3 ]
```

## TODO

* Example where this project is imported and built with goBuildPackage

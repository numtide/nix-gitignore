{ pkgs ? import <nixpkgs> {} }:
let
  filterGitignore = import ../. {};
in
  pkgs.stdenv.mkDerivation {
    name = "fixtures";

    nativeBuildInputs = with pkgs; [ tree ];

    src = filterGitignore ./test-gitignore ./.;

    buildPhase = ''
      tree

      false
    '';
  }

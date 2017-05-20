{ nix-gitignore ? ./nix-gitignore }:
let
  isNotIgnored = gitignore: path: type:
     builtins.exec [./nix-gitignore gitignore path type];

  filterGitignore = gitignore: path:
      builtins.filterSource (isNotIgnored gitignore) path;
in
  # needs nix v1.11.9+
  assert builtins.compareVersions "1.11.9" builtins.nixVersion >= 0;
  # needs --option allow-unsafe-native-code-during-evaluation true
  assert builtins.hasAttr "exec" builtins;
  filterGitignore

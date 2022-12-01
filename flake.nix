{
  description = "Advent of Code 2022 -- Nix Flake Edition";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      devShells.default = pkgs.mkShell {
        nativeBuildInputs = with pkgs; [
          bashInteractive

          # Go development
          go
          gopls
        ];
        buildInputs = with pkgs; [ ];
      };
    });
}

{
  description = "Web application for editing and managing audio metadata remotely";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-24.11";
  };

  outputs =
    { self, nixpkgs }:
    let
      armPkgs = nixpkgs.legacyPackages.aarch-linux;
      pkgs = nixpkgs.legacyPackages.x86_64-linux;

      armDeps = with armPkgs; [
        go
        nodejs_18
        tageditor
      ];

      deps = with pkgs; [
        go
        nodejs_18
        tageditor
      ];

    in
    {

      devShells.aarch-linux.default = armPkgs.mkShell {
        packages = armDeps;
      };

      devShells.x86_64-linux.default = pkgs.mkShell {
        packages = deps;
      };

    };
}

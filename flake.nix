{
  description = "Web application for editing and managing audio metadata remotely";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-24.11";
  };

  outputs =
    { self, nixpkgs }:
    let
      armPkgs = nixpkgs.legacyPackages.aarch64-linux;
      pkgs = nixpkgs.legacyPackages.x86_64-linux;

      armDeps = with armPkgs; [
        go
        nodejs_18
        tageditor
        ffmpeg
      ];

      deps = with pkgs; [
        go
        nodejs_18
        tageditor
        ffmpeg
      ];

    in
    {

      devShells.aarch64-linux.default = armPkgs.mkShell {
        packages = armDeps;
      };

      devShells.x86_64-linux.default = pkgs.mkShell {
        packages = deps;
      };

      defaultPackage.aarch64-linux = let 
      yggdrasil = armPkgs.buildGoModule {
        name = "Yggdrasil";

        src = ./.;
        vendorHash = null;
      };
      in armPkgs.runCommand "yggdrasil-with-deps" {
        nativeBuildInputs = [ armPkgs.makeWrapper ];
      } ''
      mkdir -p $out/bin
      cp ${yggdrasil}/bin/webtagger $out/bin/webtagger
      ln -s ${armPkgs.tageditor}/bin/tageditor $out/bin/tageditor
      ln -s ${armPkgs.ffmpeg}/bin/ffmpeg $out/bin/ffmpeg
        '';


    };
}

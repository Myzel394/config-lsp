{
  description = "Build config-lsp";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "utils";
    };
  };

  outputs = { self, nixpkgs, utils, gomod2nix }: 
    utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux"

      "x86_64-darwin"
      "aarch64-darwin"

      "x86_64-windows"
      "aarch64-windows"
    ] (system: 
      let
        version = "0.2.1"; # CI:CD-VERSION
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            (final: prev: {
              go = prev.go_1_24;
              buildGoModule = prev.buildGo124Module;
            })
            gomod2nix.overlays.default
          ];
        };
        inputs = [
          pkgs.go_1_24
        ];
        serverUncompressed = pkgs.buildGoModule {
          nativeBuildInputs = inputs;
          pname = "github.com/Myzel394/config-lsp";
          version = version;
          src = ./server;
          vendorHash = "sha256-ttr45N8i86mSJX9Scy/Cf+YlxU5wAKMVb0YhKg28JKM=";
          ldflags = [ "-s" "-w" ];
          checkPhase = ''
            go test -v $(pwd)/...
          '';
        };
        server = pkgs.stdenv.mkDerivation {
          name = "config-lsp-${version}";
          src = serverUncompressed;
          buildInputs = [
            pkgs.upx
          ];
          buildPhase = ''
            mkdir -p $out/bin
            cp $src/bin/config-lsp $out/bin/
            chmod +rw $out/bin/config-lsp

            # upx is currently not supported for darwin
            if [ "${system}" != "x86_64-darwin" ] && [ "${system}" != "aarch64-darwin" ]; then
              upx --ultra-brute $out/bin/config-lsp
            fi
          '';
        };
      in {
        packages = {
          default = server;
          "vs-code-extension-bare" = let
            name = "config-lsp";
            node-modules = pkgs.mkYarnPackage {
              src = ./vs-code-extension;
              name = name;
              packageJSON = ./vs-code-extension/package.json;
              yarnLock = ./vs-code-extension/yarn.lock;
              yarnNix = ./vs-code-extension/yarn.nix;

              buildPhase = ''
                yarn --offline run compile:prod
              '';
              installPhase = ''
                mkdir -p extension

                # No idea why this is being created
                rm deps/${name}/config-lsp

                cp -rL deps/${name}/. extension

                mkdir -p $out
                cp -r extension/. $out
              '';
              distPhase = "true";

              buildInputs = [
                pkgs.vsce
              ];
            };
          in node-modules;
          "vs-code-extension" = let
            name = "config-lsp";
            node-modules = pkgs.mkYarnPackage {
              src = ./vs-code-extension;
              name = name;
              packageJSON = ./vs-code-extension/package.json;
              yarnLock = ./vs-code-extension/yarn.lock;
              yarnNix = ./vs-code-extension/yarn.nix;

              buildPhase = ''
                yarn --offline run compile:prod
              '';
              installPhase = ''
                mkdir -p extension

                # No idea why this is being created
                rm deps/${name}/config-lsp

                cp -rL deps/${name}/. extension
                cp ${server}/bin/config-lsp extension/out/config-lsp

                cd extension && ${pkgs.vsce}/bin/vsce package
                mkdir -p $out
                cp *.vsix $out
              '';
              distPhase = "true";

              buildInputs = [
                pkgs.vsce
              ];
            };
          in node-modules;
        };

        devShells.default = let 
          version = "0.18.1";
          ourGopls = pkgs.buildGoModule {
              pname = "gopls";
              inherit version;
              modRoot = "gopls";
              vendorHash = "sha256-ta94xPboFtSxFeuMtPX76XiC1O7osNl4oLk64wIyyz4=";

              # https://github.com/golang/tools/blob/9ed98faa/gopls/main.go#L27-L30
              ldflags = [ "-X main.version=v${version}" ];

              doCheck = false;

              # Only build gopls, and not the integration tests or documentation generator.
              subPackages = [ "." ];

              src = pkgs.fetchFromGitHub {
                owner = "golang";
                repo = "tools";
                rev = "gopls/v${version}";
                hash = "sha256-amy00VMUcmyjDoZ4d9/+YswfcZ+1/cGvFsA4sAmc1dA=";
              };
          };
        in
          pkgs.mkShell {
            buildInputs = inputs ++ (with pkgs; [
              mailutils
              wireguard-tools
              antlr
              just
              ourGopls
            ]) ++ (if pkgs.stdenv.isLinux then with pkgs; [
              postfix
            ] else []);
          };

        devShells."vs-code-extension" = pkgs.mkShell {
          buildInputs = with pkgs; [
            nodejs
            vsce
            yarn2nix
          ];
        };
      }
    );
}

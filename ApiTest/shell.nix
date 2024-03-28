
{pkgs ? import <nixpkgs> {}
}:
pkgs.mkShell{
    nam = "ambiente de desenvolvimento";
    buildInputs = [
        pkgs.go
    ];
    shellHook = ''
        echo "Comece a desenvolver!" 
    '';
}
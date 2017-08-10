# provisioner
Provisioning scripts for starting your machine up


# Opinions

* No third-party tools for managing coexisting versions of other tools which do not naturally coexist.
  EG: `rvm`, `pyenv`, `sdkman` or `nvm`.
  Tools for changing aliasing versions of software which do coexist are allowed.
  EG: `virtualenv`, or `direnv`.
  This tool will attempt to maintain simple forks and modern versions only.
  Python2/3 is a good example of supporting multiple versions. golang1.8/1.9 is not.
  `rustup` is an interesting exception.
  While it works similarly to `nvm`, it's also supported by the language as the official tool, making it the preferred install process.
* The only assumptions which should be made globally are:
  - Full internet network access. Proxy settings should be obeyed but set externally to this script.
* Assumptions per platform are:
  - OSX:
    * Command line tools are installed
    * Xcode license is in the "Accepted" state
    * User is admin
    * Brew, Brew Cask, Brew Bundle and mas-cli are installed
  - Ubuntu:
    * The system is in "normal" state.
      No changes outside apt/dpkg which aren't normal.
      No software which is/isn't present on the distributed image.
    * Full distro repos are available and assume all needed tools are not installed if they are not base tools.
* Modules should stand alone to prevent dependency hell.
  Install all dependencies every time and rely on tools reporting "already installed" caching notices.
 

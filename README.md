# upm - Universal Package Manager 1.0.0
## Features
- Traditional package managers, flatpak and snapd are freely combinable
- Uniform commands for pacman, apt, dnf, zypper, flatpak and snapd
- Very user-friendly and hassle-free to install and use
- Powerful, expandable and customisable
- Hardly any dependencies
- Self-propelling

## Tech
- [Go](https://go.dev/)
- [GNU Bash](https://www.gnu.org/software/bash/)
- [Pacman](https://archlinux.org/pacman/)
- [APT](https://wiki.debian.org/Apt)
- [DNF](https://rpm-software-management.github.io/)
- [Zypper](https://en.opensuse.org/Portal:Libzypp)
- [Flatpak](https://flatpak.org/)
- [snapd](https://snapcraft.io/)

## Installation
```sh
wget https://raw.githubusercontent.com/libalis/upm/main/upm.go && sudo go build -o /bin/upm upm.go && rm -rf upm.go
```

## Uninstall
```sh
sudo rm -rf /bin/upm ~/.config/upm.xml
```

## Usage
| shell command | abbreviation |
| ------ | ------ |
| --copyright | -c | none |
| --help | -h |
| --reset | -r |
| --version | -v |

| shell command | abbreviation |
| ------ | ------ |
| --traditional | -t |
| --flatpak | -f |
| --snapd | -s |

| upm command | abbreviation | parameter(s) |
| ------ | ------ | ------ |
| autoremove |
| info | | package(s) |
| install | in | package(s) |
| remove | rm | package(s) |
| search | se | package(s) |
| update | up |

## Examples
```sh
upm -f se audacity # Search a package only from one source
```

```sh
upm up # Update the system
```

```sh
upm -h # Get help
```

## License
[MIT](https://raw.githubusercontent.com/libalis/upm/main/LICENSE)

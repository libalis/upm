# upm - Universal Package Manager 0.0.5
## Features
- Traditional package managers, flatpak and snapd are freely combinable
- Uniform commands for pacman, apt, dnf, zypper, flatpak and snapd
- Very user-friendly and hassle-free to install and use
- Powerful, expandable and customisable
- Hardly any dependencies
- Self-propelling

## Tech
- [GNU bash 5](https://www.gnu.org/software/bash/)
- [Python 3](https://www.python.org/)

## Installation
```sh
wget https://raw.githubusercontent.com/libalis/upm/main/upm && chmod +x upm && sudo mv upm /bin/
```

## Uninstall
```sh
sudo rm /bin/upm ~/.config/upm
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
upm se -f -s audacity ffmpeg # Search packages only from flatpak and snapd
```

```sh
upm up # Update the system
```

```sh
upm -h # Get help
```

## License
MIT

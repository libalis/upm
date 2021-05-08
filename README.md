# upm - Universal Package Manager 0.0.2
## Features
- Traditional package managers, flatpak and snapd are freely combinable
- Uniform commands for pacman, apt, dnf, zypper, flatpak and snapd
- Very user-friendly and hassle-free to install and use
- Powerful, expandable and customisable
- Hardly any dependencies
- Self-propelling

## To-do
Support for:
- [x] Flatpak
- [x] Snapd
- [ ] AppImage

## Tech
- [GNU bash 5](https://www.gnu.org/software/bash/)
- [Python 3](https://www.python.org/)

## Installation
```sh
wget https://raw.githubusercontent.com/libalis/upm/main/upm && chmod +x upm && sudo mv upm /bin/
```

## Usage
| bash command | abbreviation |
| ------ | ------ |
| --copyright | -c | none |
| --help | -h |
| --reset | -r |
| --version | -v |

| bash command | abbreviation |
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
upm search -t -f audacity ffmpeg # Search packages only from flatpak and snapd
```

```sh
upm update # Update the system
```

```sh
upm --help # Get help
```

## License
MIT

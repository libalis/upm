# upm - Universal Package Manager
## Features
- Uniform commands for pacman, apt, dnf and zypper
- Very user-friendly to install and use
- Expandable and customisable
- Hardly any dependencies

## Tech
- [GNU bash 5](https://www.gnu.org/software/bash/)
- [Python 3](https://www.python.org/)

## Installation
```sh
wget https://raw.githubusercontent.com/libalis/upm/main/upm
chmod +x upm
sudo mv upm /bin/
```

## Usage
| bash command | abbreviation |
| ------ | ------ |
| --copyright | -c | none |
| --help | -h |
| --version | -v |

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
upm install gimp audacity # Install packages
upm update # Update the system
upm -help # Get help
```

## License
MIT

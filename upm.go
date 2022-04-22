package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "time"
)

var version string = "1.0.0"
var traditional, flatpak, snapd, hold string

//

var db []string = []string{"--copyright", "-c", "--help", "-h", "--reset", "-r", "--version", "-v", "--traditional", "-t", "--flatpak", "-f", "--snapd", "-s", "autoremove", "info", "install", "in", "remove", "rm", "search", "se", "update", "up"}

func checks() bool {
    for i := 0; i < len(os.Args); i++ {
        if os.Args[i][0] == '-' && contains(db, os.Args[i]) == -1 {
            return false
        }
    }
    return true
}

func contains(x []string, y string) int {
    for i := 0; i < len(x); i++ {
        if x[i] == y {
            return i
        }
    }
    return -1
}

func counts(x []string) int {
    count := 0
    for i := 0; i < len(os.Args); i++ {
        if contains(x, os.Args[i]) != -1 {
            count++
        }
    }
    return count
}

func gets(x []string, y []string) []string {
    index := -1
    for i := 0; i < len(x); i++ {
        tmp := contains(removes(os.Args), x[i])
        if tmp != -1 && (index == -1 || tmp < index) {
            index = tmp
        }
    }
    output := make([]string, len(removes(os.Args)) + len(y) - index - 1)
    for i := 0; i < len(y); i++ {
        output[i] = y[i]
    }
    for i := len(y); i < len(output); i++ {
        output[i] = removes(os.Args)[index + i - len(y) + 1]
    }
    if len(output) == len(y) {
        return []string{}
    } else {
        return output
    }
}

func holds(x string) int {
    output := 0
    if hold == "true" {
        for {
            fmt.Print("Do you want to skip the \"" + x + "\" module? (y/N): ")
            input := ""
            fmt.Scanln(&input)
            if input == "N" || input == "n" || input == "" {
                break
            } else if input == "Y" || input == "y" {
                output = 1
                break
            }
        }
    }
    return output
}

func removes(x []string) []string {
    count := 0
    for i := 0; i < len(x); i++ {
        if x[i] == "--traditional" || x[i] == "-t" || x[i] == "--flatpak" || x[i] == "-f" || x[i] == "--snapd" || x[i] == "-s" {
            count++
        }
    }
    output := make([]string, len(x) - count)
    i := 0
    for j := 0; j < len(x); j++ {
        if x[j] != "--traditional" && x[j] != "-t" && x[j] != "--flatpak" && x[j] != "-f" && x[j] != "--snapd" && x[j] != "-s" {
            output[i] = x[j]
            i++
        }
    }
    return output
}

func runs(x []string) {
    command := exec.Command("sudo", x...)
    command.Stdin = os.Stdin
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr
    command.Run()
}

//

func read() {
    home, _:= os.UserHomeDir()
    file, e := os.Open(home + "/.config/upm.xml")
    if e != nil {
        write()
    } else {
        defer file.Close()
        scanner := bufio.NewScanner(file)
        lines := 0
        for scanner.Scan() {
            lines++
        }
        if lines == 8 {
            file, _ := os.Open(home + "/.config/upm.xml")
            defer file.Close()
            scanner := bufio.NewScanner(file)
            scanner.Scan()
            if scanner.Text() == "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" {
                scanner.Scan()
                if scanner.Text() == "<upm>" {
                    scanner.Scan()
                    if scanner.Text() == "\t<version>1.0.0</version>" {
                        scanner.Scan()
                        if scanner.Text() == "\t<traditional>true</traditional>" {
                            traditional = "true"
                        } else if scanner.Text() == "\t<traditional>false</traditional>" {
                            traditional = "false"
                        } else {
                            write()
                        }
                        if traditional == "true" || traditional == "false" {
                            scanner.Scan()
                            if scanner.Text() == "\t<flatpak>true</flatpak>" {
                                flatpak = "true"
                            } else if scanner.Text() == "\t<flatpak>false</flatpak>" {
                                flatpak = "false"
                            } else {
                                write()
                            }
                            if flatpak == "true" || flatpak == "false" {
                                scanner.Scan()
                                if scanner.Text() == "\t<snapd>true</snapd>" {
                                    snapd = "true"
                                } else if scanner.Text() == "\t<snapd>false</snapd>" {
                                    snapd = "false"
                                } else {
                                    write()
                                }
                                if snapd == "true" || snapd == "false" {
                                    scanner.Scan()
                                    if scanner.Text() == "\t<hold>true</hold>" {
                                        hold = "true"
                                    } else if scanner.Text() == "\t<hold>false</hold>" {
                                        hold = "false"
                                    } else {
                                        write()
                                    }
                                    if hold == "true" || hold == "false" {
                                        scanner.Scan()
                                        if scanner.Text() != "</upm>" {
                                            write()
                                        }
                                    } else {
                                        write()
                                    }
                                } else {
                                    write()
                                }
                            } else {
                                write()
                            }
                        } else {
                            write()
                        }
                    } else {
                        write()
                    }
                } else {
                    write()
                }
            } else {
                write()
            }
        } else {
            write()
        }
    }
}

func write() {
    home, _:= os.UserHomeDir()
    file, _ := os.Create(home + "/.config/upm.xml")
    defer file.Close()
    writer := bufio.NewWriter(file)
    writer.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
    writer.WriteString("<upm>\n")
    writer.WriteString("\t<version>" + version + "</version>\n")
    for {
        fmt.Print("Do you want to activate the \"traditional\" module? (true/false): ")
        fmt.Scanln(&traditional)
        if traditional == "true" || traditional == "false" {
            break
        }
    }
    writer.WriteString("\t<traditional>" + traditional + "</traditional>\n")
    for {
        fmt.Print("Do you want to activate the \"flatpak\" module? (true/false): ")
        fmt.Scanln(&flatpak)
        if flatpak == "true" || flatpak == "false" {
            break
        }
    }
    writer.WriteString("\t<flatpak>" + flatpak + "</flatpak>\n")
    for {
        fmt.Print("Do you want to activate the \"snapd\" module? (true/false): ")
        fmt.Scanln(&snapd)
        if snapd == "true" || snapd == "false" {
            break
        }
    }
    writer.WriteString("\t<snapd>" + snapd + "</snapd>\n")
    for {
        fmt.Print("Do you want to activate the \"hold\" module? (true/false): ")
        fmt.Scanln(&hold)
        if hold == "true" || hold == "false" {
            break
        }
    }
    writer.WriteString("\t<hold>" + hold + "</hold>\n")
    writer.WriteString("</upm>\n")
    writer.Flush()
}

//

func copyright() {
    fmt.Println("MIT License")
    fmt.Println()
    fmt.Println("Copyright (c)", time.Now().Year(), "Robert Kagan")
    fmt.Println()
    fmt.Println("Permission is hereby granted, free of charge, to any person obtaining a copy")
    fmt.Println("of this software and associated documentation files (the \"Software\"), to deal")
    fmt.Println("in the Software without restriction, including without limitation the rights")
    fmt.Println("to use, copy, modify, merge, publish, distribute, sublicense, and/or sell")
    fmt.Println("copies of the Software, and to permit persons to whom the Software is")
    fmt.Println("furnished to do so, subject to the following conditions:")
    fmt.Println()
    fmt.Println("The above copyright notice and this permission notice shall be included in all")
    fmt.Println("copies or substantial portions of the Software.")
    fmt.Println()
    fmt.Println("THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR")
    fmt.Println("IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,")
    fmt.Println("FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE")
    fmt.Println("AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER")
    fmt.Println("LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,")
    fmt.Println("OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE")
    fmt.Println("SOFTWARE.")
}

func help() {
    fmt.Println("--copyright\t-c")
    fmt.Println("--help\t\t-h")
    fmt.Println("--reset\t\t-r")
    fmt.Println("--version\t-v")
    fmt.Println()
    fmt.Println("--traditional\t-t")
    fmt.Println("--flatpak\t-f")
    fmt.Println("--snapd\t\t-s")
    fmt.Println()
    fmt.Println("autoremove")
    fmt.Println("info\t\t\tpackage(s)")
    fmt.Println("install\t\tin\tpackage(s)")
    fmt.Println("remove\t\trm\tpackage(s)")
    fmt.Println("search\t\tse\tpackage(s)")
    fmt.Println("update\t\tup")
}

//

func autoremove() {
    if len(gets([]string{"autoremove"}, []string{})) == 0 {
        if traditional == "true" {
            if holds("traditional") == 0 {
                if exec.Command("pacman").Run() == nil {
                    runs([]string{"pacman", "-Rscn", "$(pacman -Qdtq)"})
                    runs([]string{"pacman", "-Scc"})
                } else if exec.Command("apt").Run() == nil {
                    runs([]string{"apt", "autoremove", "--purge"})
                    runs([]string{"apt", "clean"})
                } else if exec.Command("dnf").Run() == nil {
                    runs([]string{"dnf", "autoremove"})
                    runs([]string{"dnf", "clean", "all"})
                } else if exec.Command("zypper").Run() == nil {
                    file, _ := os.Create("/tmp/upm.sh")
                    defer file.Close()
                    writer := bufio.NewWriter(file)
                    writer.WriteString("zypper packages --unneeded | awk -F'|' 'NR==0 || NR==1 || NR==2 || NR==3 || NR==4 {next} {print $3}' > /tmp/zyp\n")
                    writer.WriteString("while read p; do zypper remove --clean-deps \"$p\"; done < /tmp/zyp\n")
                    writer.WriteString("rm -rf /tmp/zyp\n")
                    writer.WriteString("rm -rf /tmp/upm.sh\n")
                    writer.Flush()
                    runs([]string{"bash", "/tmp/upm.sh"})
                    runs([]string{"zypper", "clean"})
                }
            }
        }
        if flatpak == "true" {
            if holds("flatpak") == 0 {
                runs([]string{"flatpak", "remove", "--delete-data", "--unused"})
            }
        }
        if snapd == "true" {
            if holds("snapd") == 0 {
                file, _ := os.Create("/tmp/upm.sh")
                defer file.Close()
                writer := bufio.NewWriter(file)
                writer.WriteString("snap list --all | awk '/disabled/{print $1, $3}' |\n")
                writer.WriteString("\twhile read snapname revision; do\n")
                writer.WriteString("\t\tsnap remove --purge \"$snapname\" --revision=\"$revision\"\n")
                writer.WriteString("\tdone\n")
                writer.WriteString("rm -rf /tmp/upm.sh\n")
                writer.Flush()
                runs([]string{"bash", "/tmp/upm.sh"})
            }
        }
    } else {
        help()
    }
}

func info() {
    if len(gets([]string{"info"}, []string{})) != 0 {
        if traditional == "true" {
            if holds("traditional") == 0 {
                if exec.Command("pacman").Run() == nil {
                    runs(gets([]string{"info"}, []string{"pacman", "-Si"}))
                } else if exec.Command("apt").Run() == nil {
                    runs(gets([]string{"info"}, []string{"apt", "info"}))
                } else if exec.Command("dnf").Run() == nil {
                    runs(gets([]string{"info"}, []string{"dnf", "info"}))
                } else if exec.Command("zypper").Run() == nil {
                    runs(gets([]string{"info"}, []string{"zypper", "info"}))
                }
            }
        }
        if flatpak == "true" {
            if holds("flatpak") == 0 {
                runs(gets([]string{"info"}, []string{"flatpak", "info"}))
            }
        }
        if snapd == "true" {
            if holds("snapd") == 0 {
                runs(gets([]string{"info"}, []string{"snap", "info"}))
            }
        }
    } else {
        help()
    }
}

func install() {
    if len(gets([]string{"install", "in"}, []string{})) != 0 {
        if traditional == "true" {
            if holds("traditional") == 0 {
                if exec.Command("pacman").Run() == nil {
                    runs(gets([]string{"install", "in"}, []string{"pacman", "-S"}))
                } else if exec.Command("apt").Run() == nil {
                    runs(gets([]string{"install", "in"}, []string{"apt", "install"}))
                } else if exec.Command("dnf").Run() == nil {
                    runs(gets([]string{"install", "in"}, []string{"dnf", "install"}))
                } else if exec.Command("zypper").Run() == nil {
                    runs(gets([]string{"install", "in"}, []string{"zypper", "install"}))
                }
            }
        }
        if flatpak == "true" {
            if holds("flatpak") == 0 {
                runs(gets([]string{"install", "in"}, []string{"flatpak", "install"}))
            }
        }
        if snapd == "true" {
            if holds("snapd") == 0 {
                runs(gets([]string{"install", "in"}, []string{"snap", "install"}))
            }
        }
    } else {
        help()
    }
}

func remove() {
    if len(gets([]string{"remove", "rm"}, []string{})) != 0 {
        if traditional == "true" {
            if holds("traditional") == 0 {
                if exec.Command("pacman").Run() == nil {
                    runs(gets([]string{"remove", "rm"}, []string{"pacman", "-Rscn"}))
                } else if exec.Command("apt").Run() == nil {
                    runs(gets([]string{"remove", "rm"}, []string{"apt", "purge"}))
                } else if exec.Command("dnf").Run() == nil {
                    runs(gets([]string{"remove", "rm"}, []string{"dnf", "remove"}))
                } else if exec.Command("zypper").Run() == nil {
                    runs(gets([]string{"remove", "rm"}, []string{"zypper", "remove", "--clean-deps"}))
                }
            }
        }
        if flatpak == "true" {
            if holds("flatpak") == 0 {
                runs(gets([]string{"remove", "rm"}, []string{"flatpak", "remove"}))
                for i := 0; i < len(gets([]string{"remove", "rm"}, []string{})); i++ {
                    home, _:= os.UserHomeDir()
                    runs([]string{"rm", "-rf", home + "/.var/app/" + gets([]string{"remove", "rm"}, []string{})[i]})
                }
            }
        }
        if snapd == "true" {
            if holds("snapd") == 0 {
                runs(gets([]string{"remove", "rm"}, []string{"snap", "remove", "--purge"}))
                for i := 0; i < len(gets([]string{"remove", "rm"}, []string{})); i++ {
                    home, _:= os.UserHomeDir()
                    runs([]string{"rm", "-rf", home + "/snap/" + gets([]string{"remove", "rm"}, []string{})[i]})
                }
            }
        }
    } else {
        help()
    }
}

func search() {
    if len(gets([]string{"search", "se"}, []string{})) != 0 {
        if traditional == "true" {
            if holds("traditional") == 0 {
                if exec.Command("pacman").Run() == nil {
                    runs(gets([]string{"search", "se"}, []string{"pacman", "-Ss"}))
                } else if exec.Command("apt").Run() == nil {
                    runs(gets([]string{"search", "se"}, []string{"apt", "search"}))
                } else if exec.Command("dnf").Run() == nil {
                    runs(gets([]string{"search", "se"}, []string{"dnf", "search"}))
                } else if exec.Command("zypper").Run() == nil {
                    runs(gets([]string{"search", "se"}, []string{"zypper", "search"}))
                }
            }
        }
        if flatpak == "true" {
            if holds("flatpak") == 0 {
                runs(gets([]string{"search", "se"}, []string{"flatpak", "search"}))
            }
        }
        if snapd == "true" {
            if holds("snapd") == 0 {
                runs(gets([]string{"search", "se"}, []string{"snap", "search"}))
            }
        }
    } else {
        help()
    }
}

func update() {
    if len(gets([]string{"update", "up"}, []string{})) == 0 {
        if traditional == "true" {
            if holds("traditional") == 0 {
                if exec.Command("pacman").Run() == nil {
                    runs([]string{"pacman", "-Syyu"})
                } else if exec.Command("apt").Run() == nil {
                    runs([]string{"apt", "update"})
                    runs([]string{"apt", "full-upgrade"})
                } else if exec.Command("dnf").Run() == nil {
                    runs([]string{"dnf", "upgrade", "--refresh"})
                } else if exec.Command("zypper").Run() == nil {
                    runs([]string{"zypper", "refresh", "--force"})
                    runs([]string{"zypper", "dist-upgrade"})
                }
            }
        }
        if flatpak == "true" {
            if holds("flatpak") == 0 {
                runs([]string{"flatpak", "update"})
            }
        }
        if snapd == "true" {
            if holds("snapd") == 0 {
                runs([]string{"snap", "refresh"})
            }
        }
    } else {
        help()
    }
}

//

func main() {
    read()
    if !checks() || counts([]string{"--traditional", "-t"}) > 1 ||  counts([]string{"--flatpak", "-f"}) > 1 ||  counts([]string{"--snapd", "-s"}) > 1 {
        help()
    } else {
        if contains(os.Args, "--traditional") != -1 || contains(os.Args, "-t") != -1 || contains(os.Args, "--flatpak") != -1 || contains(os.Args, "-f") != -1 || contains(os.Args, "--snapd") != -1 || contains(os.Args, "-s") != -1 {
            traditional = "false"
            flatpak = "false"
            snapd = "false"
            hold = "false"
            if contains(os.Args, "--traditional") != -1 || contains(os.Args, "-t") != -1 {
                traditional = "true"
            }
            if contains(os.Args, "--flatpak") != -1 || contains(os.Args, "-f") != -1 {
                flatpak = "true"
            }
            if contains(os.Args, "--snapd") != -1 || contains(os.Args, "-s") != -1 {
                snapd = "true"
            }
        }
        if contains(os.Args, "--copyright") != -1 || contains(os.Args, "-c") != -1 {
            copyright();
        } else if contains(os.Args, "--help") != -1 || contains(os.Args, "-h") != -1 {
            help()
        } else if contains(os.Args, "--reset") != -1 || contains(os.Args, "-r") != -1 {
            write()
        } else if contains(os.Args, "--version") != -1 || contains(os.Args, "-v") != -1 {
            fmt.Println("upm", version, "(" + runtime.GOARCH + ")")
        } else if contains(os.Args, "autoremove") != -1 {
            autoremove()
        } else if contains(os.Args, "info") != -1 {
            info()
        } else if contains(os.Args, "install") != -1 || contains(os.Args, "in") != -1 {
            install()
        } else if contains(os.Args, "remove") != -1 || contains(os.Args, "rm") != -1 {
            remove()
        } else if contains(os.Args, "search") != -1 || contains(os.Args, "se") != -1 {
            search()
        } else if contains(os.Args, "update") != -1 || contains(os.Args, "up") != -1 {
            update()
        } else {
            help()
        }
    }
}

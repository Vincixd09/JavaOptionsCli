# **JavaOptionsCli**

## JavaOptionsCli is a simple interactive CLI tool written in Go to install and manage Java versions on Debian-based systems.

## It provides a clean terminal interface to:

- **Install a new Java version from a .tar.gz**

- **Change the active Java version**

- **List installed versions**

- **Manage update-alternatives**

## ⚠ **Requirements**

- Debian-based Linux distribution (uses update-alternatives)

- sudo privileges

- Go (if building from source)

## Installation Target

##### All Java versions are installed in:

```bash
/opt/java
```

## How to Use
### Place the tar.gz in a folder

#### Example:

```
jdk-21_linux-x64_bin.tar.gz
```

## Run the CLI from the same directory

### The CLI must be executed in the same folder where the .tar.gz file is located.

#### Example:

```bash
JavaOptionsCli
```

## Provide the tar.gz file name when prompted

### You must enter the exact name of the .tar.gz file.

## Changing Java Version


## Make It Globally Available (Optional but Recommended)

### For easier usage, it is recommended to make the binary globally accessible.

### You can either:

---

## Option 1 — Move it to a directory in your PATH

#### For example:
```bash
sudo mv JavaOptionsCli /usr/bin/JavaOptionsCli
```
#### After that, you can run it from anywhere:
```
JavaOptionsCli
```
---

## Option 2 — Create an Alias

### Add this to your ~/.bashrc or ~/.zshrc:
```bash
alias javaoptionscli="/path/to/JavaOptionsCli"
```
#### Then reload your shell:
```bash
source ~/.bashrc
```
## Why?

## This allows you to run the CLI from any directory instead of manually navigating to its location every time.

---

### The CLI will:

- List installed versions in /opt/java

- Allow you to select one

- Configure update-alternatives accordingly

## Features

- Interactive menu (powered by pterm)

- Automatic extraction of .tar.gz

- Installation to /opt/java

- Simple version switching

- Minimal and focused workflow

# Current Limitations

- Only tested on Debian-based systems

- Assumes update-alternatives is available

- Requires running from the same directory as the tar file

## Future Improvements

- Support for other Linux distributions

- Automatic download of Java releases

- Better version detection

- Configurable install path

## 🤝 Contributions

Contributions are welcome.

## 👤 Vinci

[@Vincixd09](https://github.com/Vincixd09)

---
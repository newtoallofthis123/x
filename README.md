# Exec

A command line tool to execute commands easily.
It supports various features like global alias, local file based execution alias,
command history and more.

It is meant to replace the need to put all of your command alias in the 
shell profile or use something like a Makefile

It works by having a multi alias system, where you can declare terminal alias at multiple levels
- Persistent: Persistent alias are stored in a sqlite database
- Global: Global alias are stored in a file in the config directory
- Home / User: User alias are stored in a file in the user home directory
- Local: Local alias are stored in a file in the current directory

The order is given in an ascending order, so if you have a local alias that is the same as a global alias, the local alias will be used.

## Features

- Persistent alias
- Modern config system with support for
    - pointers
    - pipes
    - multiple commands
- Command history [TODO]
- Output to file

## Usage

```bash
x [flags] command
```

Note: x is the recommended alias for exec
Note: exec does not support args since it would be counter intuitive to the purpose of the tool. Only flags are supported.

### Flags

- `-h` or `--help`: Show help
- `-v` or `--version`: Show version
- `-l`: List all alias
- `-a name cmd1,cmd2,cmd3`: Add global alias in persistent storage
- `-r name|id`: Remove alias
- `-c`: Clear all alias
- `-s`: Suppress output
- `-o <file>`: Output to file
- `-sync`: Sync the current parsed alias with the database for persistent storage

## Installation

```bash
go install github.com/newtoallofthis123/exec
```

You can optionally set `alias x=exec` for additional cool points. 

## Configuration

The config file can be located in the following locations in ascending order of priority
- `$XDG_CONFIG_HOME/exec/exec.conf`
- `$HOME/.exec.conf`
- `$PWD/.exec.conf`

The config file itself uses a simple syntax to define alias:

```conf
# For executing a single command
name = command

# For executing multiple commands
name = [command1, command2, command3]

# For pointers
name = *alias
```

## Security

- Any and all errors in the config file will result in an error because of the sensitive nature of file execution.
- The program delegates any `sudo` or privilege escalation to the shell, so it is recommended to not use this tool for such commands.
- Direct commands cannot be passed to the tool, only the name of the alias can be passed.
- The parser is very strict with the syntax, leaving little room for error.

## Examples

```bash
x -a gg git status
```

```bash
x gg
```

This example hopefully illustrates the power of the tool. It is meant to be a simple and easy to use tool for executing commands.

```bash
x -a gs git add .,git commit,git push
```

## Use Cases

### 1. Persistent Alias

I for one like to set a lot of alias in my shell profile, so this tool
essentially simplifies that process by allowing me to set alias simply by executing a command.
This command can then be used in any terminal session and can be easily managed.

### 2. Local Alias

I sometimes need to set alias for a specific project, so I can set a local alias in the project directory
by simply creating a `exec.conf` file in the directory and setting the alias there.

For example, I can set a local alias for a project that has a complex build process, even including sensitive
ENV variables that would not be recommended to set in the shell profile.

```bash
# .exec.conf
build = DB_USER=postgres DB_PASS=postgres go build .
b = *build
```

Now I can simply run `x b` to build the project.

### 3. Replace Makefile

I can use this tool to replace the need for a Makefile in a project. I can set alias for all the commands
that I would normally put in a Makefile in the `exec.conf` file.

```bash
# .exec.conf
build = [go build ., go test .]
t = *b
```

Now I can simply run `x t` to build and test the project.

So hence, the aliases are smaller, localized and easier to manage.

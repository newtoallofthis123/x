package parser

import (
	"os"
	"strings"
)

type Parser struct {
	configFile string
	content    string
	cmds       map[string][]string
}

// MakeParser creates and returns a new Parser instance using the provided configuration file.
func MakeParser(configFile string) (Parser, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return Parser{}, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return Parser{}, err
	}

	content := make([]byte, stat.Size())
	_, err = file.Read(content)

	// remove the last newline
	if content[len(content)-1] == '\n' {
		content = content[:len(content)-1]
	}

	return Parser{configFile, string(content), make(map[string][]string)}, nil
}

// Parse processes the configuration and returns an error if any issues are encountered.
func (p *Parser) Parse() error {
	lines := strings.Split(p.content, "\n")

	pointers := make(map[string]string, 0)

	for _, line := range lines {
		// comments start with #
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		subs := strings.Split(line, "=")
		if len(subs) != 2 {
			return &ParserError{line, "Invalid line"}
		}

		name := strings.TrimSpace(subs[0])
		cmdsRaw := strings.TrimSpace(subs[1])
		cmds := make([]string, 0)
		if cmdsRaw[0] == '*' {
			pointers[name] = cmdsRaw[1:]
			continue
		}
		if cmdsRaw[0] == '[' && cmdsRaw[len(cmdsRaw)-1] == ']' {
			cmds = strings.Split(cmdsRaw[1:len(cmdsRaw)-1], ",")
		} else {
			cmds = []string{cmdsRaw}
		}

		p.cmds[name] = cmds
	}

	for k, v := range pointers {
		cmds, ok := p.cmds[v]
		if !ok {
			return &ParserError{k, "Pointer not found"}
		}

		p.cmds[k] = cmds
	}

	return nil
}

// GetCmd retrieves the command associated with the given name.
// It returns the command as a slice of strings and a boolean indicating if the command was found.
func (p *Parser) GetCmd(name string) ([]string, bool) {
	cmds, ok := p.cmds[name]
	return cmds, ok
}

func (p *Parser) GetCmds() map[string][]string {
	return p.cmds
}

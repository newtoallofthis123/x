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

func (p *Parser) Parse() error {
	lines := strings.Split(p.content, "\n")
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
		if cmdsRaw[0] == '[' && cmdsRaw[len(cmdsRaw)-1] == ']' {
			cmds = strings.Split(cmdsRaw[1:len(cmdsRaw)-1], ",")
		} else {
			cmds = []string{cmdsRaw}
		}

		p.cmds[name] = cmds
	}

	return nil
}

func (p *Parser) GetCmd(name string) ([]string, bool) {
	cmds, ok := p.cmds[name]
	return cmds, ok
}

func (p *Parser) GetCmds() map[string][]string {
	return p.cmds
}

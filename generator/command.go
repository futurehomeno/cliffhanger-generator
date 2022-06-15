package generator

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CommandSet map[string][]*Command

type Command struct {
	Cmd  string
	Args []string
	Dir  string
}

func (c *Command) String() string {
	return fmt.Sprintf("%s %s", c.Cmd, strings.Join(c.Args, " "))
}

func (c *Command) Execute(cfg Config) error {
	if err := c.prepare(cfg); err != nil {
		return err
	}

	return c.execute()
}

func (c *Command) execute() error {
	cmd := exec.Command(c.Cmd, c.Args...) //nolint:gosec
	cmd.Dir = c.Dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute the command %s: %w", c.String(), err)
	}

	return nil
}

func (c *Command) prepare(cfg Config) error {
	var err error

	c.Cmd, err = renderString("command", c.Cmd, cfg)
	if err != nil {
		return fmt.Errorf("failed to prepare command %s: %w", c.String(), err)
	}

	c.Dir, err = renderString("directory", c.Dir, cfg)
	if err != nil {
		return fmt.Errorf("failed to prepare command %s: %w", c.String(), err)
	}

	for i := range c.Args {
		c.Args[i], err = renderString("argument", c.Args[i], cfg)
		if err != nil {
			return fmt.Errorf("failed to prepare command %s: %w", c.String(), err)
		}
	}

	return nil
}

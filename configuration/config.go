package configuration

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/futurehomeno/cliffhanger-generator/generator"
)

const (
	DomainEdge = "edge"
	DomainCore = "core"

	TypeApp     = "app"
	TypeAdapter = "adapter"
)

var _ generator.Config = (*Config)(nil)

type Config struct {
	Domain          string
	Type            string
	Name            string
	Description     string
	Path            string
	ServiceName     string
	PackageName     string
	RepositoryName  string
	RepositoryPath  string
	SourcePath      string
	IncludeComments bool
	IncludeTesting  bool
}

func (c *Config) Configure() error {
	c.ServiceName = strings.ReplaceAll(strings.ToLower(c.Name), " ", "_")
	c.PackageName = strings.ReplaceAll(strings.ToLower(c.Name), " ", "-")
	c.RepositoryName = fmt.Sprintf("%s-%s", c.Domain, c.PackageName)

	if c.Domain == DomainEdge {
		c.RepositoryName = fmt.Sprintf("%s-%s", c.RepositoryName, c.Type)
	}

	c.RepositoryPath = path.Join(c.Path, c.RepositoryName)
	c.SourcePath = path.Join(c.RepositoryPath, "src")

	return nil
}

func (c *Config) Validate() error {
	err := ValidateDomain(c.Domain)
	if err != nil {
		return err
	}

	err = ValidateType(c.Type)
	if err != nil {
		return err
	}

	err = ValidateName(c.Name)
	if err != nil {
		return err
	}

	err = ValidateDescription(c.Description)
	if err != nil {
		return err
	}

	err = ValidatePath(c.Path)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) RootPath() string {
	return c.RepositoryPath
}

func (c *Config) TemplateConfig() interface{} {
	return &Config{
		ServiceName: "service_name",
		PackageName: "package-name",
	}
}

func (c *Config) FileSets() []string {
	sets := []string{"common"}

	if c.Domain == DomainEdge {
		sets = append(sets, c.edgeFileSets()...)
	} else {
		sets = append(sets, "core")
	}

	if c.IncludeTesting {
		sets = append(sets, "testing")
	}

	return sets
}

func (c *Config) edgeFileSets() []string {
	sets := []string{"edge"}

	if c.Type == TypeAdapter {
		sets = append(sets, "edge_adapter")
	}

	if c.IncludeTesting {
		sets = append(sets, "testing_edge")

		if c.Type == TypeAdapter {
			sets = append(sets, "testing_edge_adapter")
		}
	}

	return sets
}

func (c *Config) CommandSets() []string {
	return []string{"common"}
}

var nameIsValid = regexp.MustCompile(`^[A-Za-z\d ]+$`).MatchString

func ValidateDomain(domain string) error {
	if domain != DomainEdge && domain != DomainCore {
		return fmt.Errorf("domain must be either '%s' or '%s'", DomainEdge, DomainCore)
	}

	return nil
}

func ValidateType(resourceType string) error {
	if resourceType != TypeApp && resourceType != TypeAdapter {
		return fmt.Errorf("type must be either '%s' or '%s'", TypeApp, TypeAdapter)
	}

	return nil
}

func ValidateName(name string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return fmt.Errorf("name cannot be empty")
	}

	if !nameIsValid(name) {
		return fmt.Errorf("name must be an alphanumeric phrase")
	}

	if strings.Count(name, " ") > 2 {
		return fmt.Errorf("name must consist of no more than three words")
	}

	return nil
}

func ValidateDescription(description string) error {
	if len(strings.TrimSpace(description)) == 0 {
		return fmt.Errorf("description cannot be empty")
	}

	return nil
}

func ValidatePath(p string) error {
	info, err := os.Stat(p)
	if os.IsNotExist(err) {
		return fmt.Errorf("provided path does not exist")
	} else if err != nil {
		return fmt.Errorf("cannot check the path %s: %w", p, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("provided path %s is not a directory", p)
	}

	return nil
}

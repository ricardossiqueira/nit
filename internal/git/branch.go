/*Package git
 */
package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Branch struct {
	Name string
}

func GetBranch() (*Branch, error) {
	var cmdOut bytes.Buffer
	var cmdErr bytes.Buffer

	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("git branch: %s", cmdErr.String())
	}

	return &Branch{Name: strings.TrimSpace(cmdOut.String())}, nil
}

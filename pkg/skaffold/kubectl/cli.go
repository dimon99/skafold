/*
Copyright 2019 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kubectl

import (
	"context"
	"io"
	"os/exec"
	"sync"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/runner/runcontext"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/util"
)

// CLI holds parameters to run kubectl.
type CLI struct {
	KubeContext string
	Namespace   string

	version     ClientVersion
	versionOnce sync.Once
}

func NewFromRunContext(runCtx *runcontext.RunContext) *CLI {
	return &CLI{
		KubeContext: runCtx.KubeContext,
		Namespace:   runCtx.Opts.Namespace,
	}
}

// Command creates the underlying exec.CommandContext. This allows low-level control of the executed command.
func (c *CLI) Command(ctx context.Context, command string, arg ...string) *exec.Cmd {
	args := c.args(command, arg...)
	return exec.CommandContext(ctx, "kubectl", args...)
}

// Run shells out kubectl CLI.
func (c *CLI) Run(ctx context.Context, in io.Reader, out io.Writer, command string, arg ...string) error {
	cmd := c.Command(ctx, command, arg...)
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = out
	return util.RunCmd(cmd)
}

// Run shells out kubectl CLI.
func (c *CLI) RunOut(ctx context.Context, command string, arg ...string) ([]byte, error) {
	cmd := c.Command(ctx, command, arg...)
	return util.RunCmdOut(cmd)
}

// args builds an argument list for calling kubectl and consistently
// adds the `--context` and `--namespace` flags.
func (c *CLI) args(command string, arg ...string) []string {
	args := []string{"--context", c.KubeContext}
	if c.Namespace != "" {
		args = append(args, "--namespace", c.Namespace)
	}
	args = append(args, command)
	args = append(args, arg...)
	return args
}

/*
   Copyright The containerd Authors.

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

package runc

import (
	"context"
	"os"
	"os/exec"
	"syscall"
	"strings"
)

func (r *Runc) command(context context.Context, args ...string) *exec.Cmd {
	command := r.Command
	if command == "" {
		command = DefaultCommand
	}

	//	need to print commands here
	f, _ := os.OpenFile("/home/debian/go-runc.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	f.WriteString("command with args are " + command + " " + strings.Join(r.args()," ") + " " + strings.Join(args," "))
	f.WriteString(" ----end of invocation\n\n")
	f.Close()

	cmd := exec.CommandContext(context, command, append(r.args(), args...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: r.Setpgid,
	}
	cmd.Env = os.Environ()
	if r.PdeathSignal != 0 {
		cmd.SysProcAttr.Pdeathsig = r.PdeathSignal
	}

	return cmd
}

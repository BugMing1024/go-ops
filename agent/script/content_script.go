package script

import (
	"os"
	"osp/agent/cmdrunner"
	"osp/internal/model"

	"osp/agent/script/cmd"
	"path"
	"strings"
)

type ContentScript struct {
	GenericScript
	cmd   string
	input string
}

func NewContentScript(
	runner cmdrunner.CmdRunner,
	jobid string,
	path string,
	cmd string,
	content string,
	env map[string]string,
	timeout int,
	input string,
	user string,
	args []string,
) ContentScript {

	if cmd == "" {
		cmd = Cmder
	}

	s := ContentScript{cmd: cmd, input: input}
	s.GenericScript.runner = runner
	s.GenericScript.path = path
	s.GenericScript.content = content
	s.GenericScript.env = env
	s.GenericScript.jobid = jobid
	s.GenericScript.timeout = timeout
	s.GenericScript.user = user
	s.GenericScript.args = args

	return s
}

func (s ContentScript) Run() (r model.ResCmd) {

	if s.path == "" {
		s.path = ScriptPath
	}
	runpath := path.Join(s.path, s.jobid)
	runpath = path.Join(runpath, s.jobid+ScriptExt)
	err := s.ensureContainingDir(runpath)

	if err != nil {
		return s.getResCmd(nil, err)
	}

	f, err := os.OpenFile(runpath, fileOpenFlag, fileOpenPerm)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.WriteString(s.content)
	if err != nil {
		return
	}
	err = f.Close()
	if err != nil {
		return
	}

	cmdstr, args := getCmdArgs(s.cmd)

	command := cmd.BuildCommand(cmdstr)

	command.Args = append(command.Args, args...)
	command.Args = append(command.Args, runpath)
	command.Args = append(command.Args, s.args...)
	command.Timeout = s.timeout
	command.User = s.user

	if s.input != "" {
		command.Stdin = strings.NewReader(s.input)
	}

	for key, val := range s.env {
		command.Env[key] = val
	}

	res, err := s.runner.RunCommand(s.jobid, command)
	return s.getResCmd(res, err)
}

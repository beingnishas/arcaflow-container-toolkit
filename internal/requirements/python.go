package requirements

import (
	"bytes"
	"fmt"
	"os"

	"go.arcalot.io/arcaflow-container-toolkit/internal/util"
	"go.arcalot.io/log"
)

type ExternalProgramOnFile func(executable_filepath string, stdout *bytes.Buffer, logger log.Logger) error

func PythonRequirements(abspath string, filenames []string, name string, version string, logger log.Logger,
	pythonCodeStyleChecker func(abspath string, stdout *bytes.Buffer, logger log.Logger) error) (bool, error) {
	meets_reqs, err := PythonFileRequirements(filenames, logger)
	if err != nil {
		return false, err
	}
	good_style, err := PythonCodeStyle(abspath, name, version, pythonCodeStyleChecker, logger)
	if err != nil {
		return false, err
	} else if !good_style {
		meets_reqs = false
	}
	return meets_reqs, nil
}

func PythonCodeStyle(abspath string, name string, version string, checkPythonCodeStyle ExternalProgramOnFile, logger log.Logger) (bool, error) {
	stdout := &bytes.Buffer{}
	if err := checkPythonCodeStyle(abspath, stdout, logger); err != nil {
		if len(stdout.String()) > 0 {
			logger.Infof("python code style and quality check found these issues for %s version %s", name, version)
			logger.Infof("(%w)", stdout.String())
			return false, nil
		} else {
			return false, fmt.Errorf("error in executing python code style check for %s (%w)", name, err)
		}
	}
	return true, nil
}

func PythonFileRequirements(filenames []string, logger log.Logger) (bool, error) {
	meets_reqs := true
	has_reqs_txt, err := HasFilename(filenames, "requirements.txt")
	if err != nil {
		return false, err
	}
	has_pyproject, err := HasFilename(filenames, "pyproject.toml")
	if err != nil {
		return false, err
	}
	if !has_reqs_txt && !has_pyproject {
		logger.Infof("Missing a dependency manager: either add requirements.txt or pyproject.toml")
		meets_reqs = false
	}
	return meets_reqs, nil
}

func Flake8PythonCodeStyle(abspath string, stdout *bytes.Buffer, logger log.Logger) error {
	err := os.Chdir(abspath)
	if err != nil {
		return err
	}
	return util.RunExternalProgram(
		"python3",
		[]string{
			"-m",
			"flake8",
			"--show-source",
			abspath,
		},
		nil,
		nil,
		stdout,
		stdout,
	)
}

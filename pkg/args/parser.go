package args

import (
	"os"
	"strings"
	"time"
)

type CommandType string

const (
	Create  CommandType = "create"
	Destroy CommandType = "destroy"
	On      CommandType = "on"
	Off     CommandType = "off"
)

type OnParams struct {
	Duration time.Duration
}

type Args struct {
	Command CommandType
	Params  OnParams
}

func ParseArgs() (Args, error) {
	default_args := Args{
		Command: On,
		Params: OnParams{
			Duration: 1 * time.Hour,
		},
	}
	args := os.Args[1:]

	if len(args) == 0 {
		return default_args, nil
	}

	if args[0] == "with" {
		return Args{
			Command: Create,
		}, nil
	}

	if args[0] == "for" {
		duration, err := time.ParseDuration(strings.Join(args[1:], " "))
		if err != nil {
			return default_args, err
		}
		return Args{
			Command: On,
			Params: OnParams{
				Duration: duration,
			},
		}, nil
	}

	if args[0] == "no" || args[0] == "not" {
		return Args{
			Command: Off,
		}, nil
	}

	if args[0] == "never" {
		return Args{
			Command: Destroy,
		}, nil
	}

	return default_args, nil
}

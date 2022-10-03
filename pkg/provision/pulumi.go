package provision

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optdestroy"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
)

type PulumiProvisioner struct {
	stack auto.Stack
}

type PulumiConfig struct {
	Region        string
	Profile       string
	SSHPubKeyPath string
	InstanceType  string
}

func NewPulumiProvisioner() *PulumiProvisioner {
	ctx := context.Background()

	stackName := "dev"
	// stackName := auto.FullyQualifiedStackName("myOrgOrUser", projectName, stackName)

	workDir := filepath.Join(".", "infra")

	s, err := auto.UpsertStackLocalSource(ctx, stackName, workDir)
	if err != nil {
		panic(fmt.Errorf("failed to create or select stack: %v", err))
	}

	return &PulumiProvisioner{
		stack: s,
	}
}

func (p *PulumiProvisioner) SetConfig(config PulumiConfig) error {
	ctx := context.Background()
	if err := p.stack.SetConfig(ctx, "aws:region", auto.ConfigValue{Value: config.Region}); err != nil {
		return err
	}
	if err := p.stack.SetConfig(ctx, "aws:profile", auto.ConfigValue{Value: config.Profile}); err != nil {
		return err
	}
	if err := p.stack.SetConfig(ctx, "i-wanna-play-minecraft:SSH_PUB_KEY", auto.ConfigValue{Value: config.SSHPubKeyPath}); err != nil {
		return err
	}
	if err := p.stack.SetConfig(ctx, "i-wanna-play-minecraft:EC2_INSTANCE_TYPE", auto.ConfigValue{Value: config.InstanceType}); err != nil {
		return err
	}
	return nil
}

func (p *PulumiProvisioner) Refresh() error {
	_, err := p.stack.Refresh(context.Background())
	if err != nil {
		return fmt.Errorf("failed to refresh stack: %v", err)
	}
	return nil
}

func (p *PulumiProvisioner) Up() (auto.UpResult, error) {
	stdoutStreamer := optup.ProgressStreams(os.Stdout)
	return p.stack.Up(context.Background(), stdoutStreamer)
}

func (p *PulumiProvisioner) Destroy() (auto.DestroyResult, error) {
	stdoutStreamer := optdestroy.ProgressStreams(os.Stdout)
	return p.stack.Destroy(context.Background(), stdoutStreamer)
}

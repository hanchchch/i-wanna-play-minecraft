package main

import (
	"fmt"

	"github.com/hanchchch/i-wanna-play-minecraft/pkg/args"
	"github.com/hanchchch/i-wanna-play-minecraft/pkg/instance"
	"github.com/hanchchch/i-wanna-play-minecraft/pkg/provision"
	"github.com/hanchchch/i-wanna-play-minecraft/pkg/utils"
)

func getInstance(p *provision.PulumiProvisioner, c *args.Config) *instance.AWSInstance {
	r, err := p.Up()
	if err != nil {
		panic(err)
	}
	instanceId := r.Outputs["serverId"].Value.(string)
	return instance.NewAWSInstance(instance.AWSInstanceOptions{
		InstanceID: instanceId,
		Profile:    c.Profile,
		Region:     c.Region,
	})
}

func main() {
	a, err := args.ParseArgs()
	if err != nil {
		panic(err)
	}

	c, err := args.LoadConfig()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", a.Command)

	fmt.Println("preparing pulumi provider...")
	p := provision.NewPulumiProvisioner()
	p.SetConfig(provision.PulumiConfig{
		Region:        c.Region,
		Profile:       c.Profile,
		SSHPubKeyPath: c.SSHPubKeyPath,
		InstanceType:  c.InstanceType,
	})
	if err := p.Refresh(); err != nil {
		panic(err)
	}

	switch a.Command {
	case args.Create:
		fmt.Println("creating...")
		if _, err := p.Up(); err != nil {
			panic(err)
		}
	case args.Destroy:
		fmt.Println("destroying...")
		if _, err := p.Destroy(); err != nil {
			panic(err)
		}
	case args.On:
		fmt.Println("turning on...")
		i := getInstance(p, &c)
		go utils.PrintElapsedTime("running")
		if err := i.RunFor(a.Params.Duration); err != nil {
			panic(err)
		}
	case args.Off:
		fmt.Println("turning off...")
		if _, err := getInstance(p, &c).Stop(); err != nil {
			panic(err)
		}
	}
}

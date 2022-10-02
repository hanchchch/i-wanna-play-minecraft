package main

import (
	"fmt"

	"github.com/hanchchch/i-wanna-play-minecraft/pkg/args"
	"github.com/hanchchch/i-wanna-play-minecraft/pkg/provision"
)

func main() {
	a, err := args.ParseArgs()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", a.Command)

	fmt.Printf("preparing pulumi provider...")
	p := provision.NewPulumiProvisioner()
	fmt.Printf(" done\n")

	p.SetConfig(provision.PulumiConfig{
		Region:        "ap-northeast-2",
		Profile:       "personal",
		SSHPubKeyPath: "minecraft-pub.pem",
		InstanceType:  "c6i.large",
	})

	fmt.Printf("refreshing...")
	p.Refresh()
	fmt.Printf(" done\n")

	switch a.Command {
	case args.Create:
		fmt.Printf("creating...")
		p.Up()
	case args.Destroy:
		fmt.Printf("destroying...")
		p.Destroy()
	case args.On:
		fmt.Printf("turning on...")
	case args.Off:
		fmt.Printf("turning off...")
	}
	fmt.Println("done")

}

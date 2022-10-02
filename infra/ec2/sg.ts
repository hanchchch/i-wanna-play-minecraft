import * as aws from "@pulumi/aws";
import { vpc } from "./vpc";

export const allowMCServer = new aws.ec2.SecurityGroup("allow-minecraft", {
  vpcId: vpc.id,
  description: "Allow Minecraft connection inbound traffic",
  ingress: [
    {
      description: "Minecraft connection",
      fromPort: 25565,
      toPort: 25565,
      protocol: "tcp",
      cidrBlocks: ["0.0.0.0/0"],
      ipv6CidrBlocks: ["::/0"],
    },
  ],
  egress: [
    {
      fromPort: 0,
      toPort: 0,
      protocol: "-1",
      cidrBlocks: ["0.0.0.0/0"],
      ipv6CidrBlocks: ["::/0"],
    },
  ],
});

export const allowSSH = new aws.ec2.SecurityGroup("allow-ssh", {
  vpcId: vpc.id,
  description: "Allow SSH connection inbound traffic",
  ingress: [
    {
      description: "SSH connection",
      fromPort: 22,
      toPort: 22,
      protocol: "tcp",
      cidrBlocks: ["0.0.0.0/0"],
      ipv6CidrBlocks: ["::/0"],
    },
  ],
  egress: [
    {
      fromPort: 0,
      toPort: 0,
      protocol: "-1",
      cidrBlocks: ["0.0.0.0/0"],
      ipv6CidrBlocks: ["::/0"],
    },
  ],
});

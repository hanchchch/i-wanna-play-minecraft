import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import { allowMCServer, allowSSH } from "./sg";
import { getUserData } from "./user-data";
import { vpc } from "./vpc";
import { keyPair } from "./key-pair";
import { config } from "../utils/config";

const ubuntu = aws.ec2.getAmi({
  mostRecent: true,
  filters: [
    {
      name: "name",
      values: ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"],
    },
    {
      name: "virtualization-type",
      values: ["hvm"],
    },
  ],
  owners: ["099720109477"],
});

const server = new aws.ec2.Instance("mc-server", {
  ami: ubuntu.then((ubuntu) => ubuntu.id),
  keyName: keyPair.keyName,
  subnetId: pulumi
    .all([vpc.publicSubnetIds])
    .apply(([publicSubnetIds]) => publicSubnetIds[0]),
  instanceType: config.EC2_INSTANCE_TYPE,
  vpcSecurityGroupIds: [allowMCServer.id, allowSSH.id],
  userData: getUserData(),
  tags: {
    Name: "mc-server",
  },
});

export { server };

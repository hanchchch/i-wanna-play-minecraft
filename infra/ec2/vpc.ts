import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";

const vpc = new awsx.ec2.Vpc("minecraft-vpc", {
  cidrBlock: "10.0.0.0/16",
});

export { vpc };

import * as pulumi from "@pulumi/pulumi";

const cfg = new pulumi.Config();

export const config = {
  SSH_PUB_KEY: cfg.require("SSH_PUB_KEY"),
  EC2_INSTANCE_TYPE: cfg.require("EC2_INSTANCE_TYPE"),
};

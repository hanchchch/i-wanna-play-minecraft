import * as aws from "@pulumi/aws";
import { readFileContent } from "../utils/read";
import { config } from "../utils/config";

const keyPair = new aws.ec2.KeyPair("minecraft-keypair", {
  keyName: "minecraft",
  publicKey: readFileContent(config.SSH_PUB_KEY),
});

export { keyPair };

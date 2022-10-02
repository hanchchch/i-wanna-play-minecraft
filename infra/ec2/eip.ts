import * as aws from "@pulumi/aws";
import { server } from "./instance";

const eip = new aws.ec2.Eip("minecraft-eip", {
  vpc: true,
  instance: server.id,
});

export { eip };

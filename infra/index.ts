import { server } from "./ec2/instance";
import { eip } from "./ec2/eip";

export const serverId = server.id;
export const eipPublicIp = eip.publicIp;

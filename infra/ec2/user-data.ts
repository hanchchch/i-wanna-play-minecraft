import { readFileContent } from "../utils/read";

const getUserData = async () =>
  `Content-Type: multipart/mixed; boundary="==BOUNDARY=="
MIME-Version: 1.0

--==BOUNDARY==
Content-Type: text/cloud-config; charset="us-ascii"
MIME-Version: 1.0
Content-Disposition: attachment; filename="cloud-config.txt"

#cloud-config
cloud_final_modules:
- [scripts-user, always]

--==BOUNDARY==
Content-Type: text/x-shellscript; charset="us-ascii"
MIME-Version: 1.0
Content-Disposition: attachment; filename="userdata.txt"

${await readFileContent("scripts/user-data.sh")}
--==BOUNDARY==--`;

export { getUserData };

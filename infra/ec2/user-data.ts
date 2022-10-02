import { readFileContent } from "../utils/read";

const getUserData = async () =>
  `Content-Type: multipart/mixed; boundary="==BOUNDARY=="
MIME-Version: 1.0

--==BOUNDARY==
Content-Type: text/x-shellscript; charset="us-ascii"

${await readFileContent("scripts/user-data.sh")}
--==BOUNDARY==--`;

export { getUserData };

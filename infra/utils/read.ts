import { readFile } from "fs";

export const readFileContent = (filename: string) =>
  new Promise<string>((res, rej) =>
    readFile(filename, (err, data) => {
      if (err) {
        rej(err);
      } else {
        res(data.toString("utf-8"));
      }
    })
  );

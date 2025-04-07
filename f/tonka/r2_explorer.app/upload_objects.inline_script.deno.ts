import { Resource } from "https://deno.land/x/windmill@v1.85.0/mod.ts";
import { S3Client } from "https://deno.land/x/s3_lite_client@0.5.0/mod.ts";
import { decode } from "https://deno.land/std@0.181.0/encoding/base64.ts";

export async function main(
  s3: Resource<"s3">,
  objects: {name: string, data: string}[],
  acl: string,
) {
  const s3client = new S3Client(s3);
  const promises = objects.map(({name, data}) => {
    const contentType = data.slice(5, data.indexOf(";"))
    return s3client.putObject(
      name,
      contentType.startsWith('image') ? decode(data.slice(data.indexOf(',') + 1)) : data,
      {
        metadata: {
          "x-amz-acl": acl,
          "Content-Type": contentType
        },
      },
    );
  });
  return await Promise.all(promises);
}

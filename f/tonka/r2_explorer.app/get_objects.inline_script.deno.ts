import { Resource } from "https://deno.land/x/windmill@v1.85.0/mod.ts";
import { removeObjectEmptyFields } from "https://deno.land/x/windmill_helpers@v1.1.1/mod.ts";
import { S3Client } from "https://deno.land/x/s3_lite_client@0.5.0/mod.ts";

export async function main(
  s3: Resource<"s3">,
  prefix?: string,
  bucketName?: string,
  maxResults?: number,
  pageSize?: number,
) {
  if(!s3) {
    return []
  }
  const s3client = new S3Client(s3);
  const options = removeObjectEmptyFields({
    prefix,
    bucketName,
    maxResults,
    pageSize,
  })
  const objects = s3client.listObjects(options);

  const result: any[] = []
  for await (const obj of objects) {
    result.push(obj)
  }

  return result
}
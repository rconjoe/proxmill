import { Resource } from "https://deno.land/x/windmill@v1.85.0/mod.ts";
import { S3Client } from "https://deno.land/x/s3_lite_client@0.5.0/mod.ts";

export async function main(s3: Resource<"s3">, object_name: string) {
  const s3client = new S3Client(s3);
  return await s3client.getPresignedUrl('GET', object_name)
}
import { Resource } from "https://deno.land/x/windmill@v1.85.0/mod.ts";
import { S3Client } from "https://deno.land/x/s3_lite_client@0.5.0/mod.ts";
import { encode, decode } from "https://deno.land/std@0.181.0/encoding/base64.ts";

export async function main(s3: Resource<"s3">, object: {key: string}) {
  if(!(s3 && object)) {
    return 'Select a file in the table'
  }
  const s3client = new S3Client(s3);
  const res = await s3client.getObject(object.key)
  const ct = res.headers.get('Content-Type')
  const type = ct?.slice(ct.indexOf('/') + 1)
  if(!type) {
    return object
  }
  if(['png', 'jpeg', 'gif', 'svg'].includes(type)) {
    return {
      [type]: encode(await res.arrayBuffer())
    }
  } else if(ct?.startsWith('text/')) {
    const textDecoder = new TextDecoder('utf-8')
    const encodedText = await res.text()
    const binaryText = decode(encodedText.slice(encodedText.indexOf(',') + 1))
    const decodedText = textDecoder.decode(binaryText)
    return type === 'html' ? { html: decodedText } : decodedText
  } else if(type === 'json') {
    const textDecoder = new TextDecoder('utf-8')
    const encodedText = await res.text()
    const binaryText = decode(encodedText.slice(encodedText.indexOf(',') + 1))
    const decodedText = textDecoder.decode(binaryText)
    return JSON.parse(decodedText)
  }
  return object
}
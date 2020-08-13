# Figport

A simple utility to make mor easy publish your design assets. The main goal of Figport is help to publish or export your Figma designs (e.g. icons, logos, illustrations, etc) directly into your S3 object storage. You only need two things to launch the Figport service:

1. **Figma Auth Token**, at the moment Figport only support your Personal Figma Access Token. To know how you can obtain your credential see [the official documentation](https://www.figma.com/developers/api#access-tokens).
2. **S3 Credentials**, you can obtain these credentials from your S3 storage (e.g. AWS S3, Digital Ocean Spaces, Minio).
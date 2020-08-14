# Figport

A simple utility to make easy publish your design assets. The main goal of Figport is help to publish or export your Figma designs (e.g. icons, logos, illustrations, etc) directly into your S3 object storage. You only need two things to launch the Figport service:

1. **Figma Auth Token**, at the moment Figport only support your Personal Figma Access Token. To know how you can obtain your credential see [the official documentation](https://www.figma.com/developers/api#access-tokens).
2. **S3 Credentials**, you can obtain these credentials from your S3 storage (e.g. AWS S3, Digital Ocean Spaces, Minio).



## Getting started

1. Launch the Figport service. Figport service is a docker image and it's delivery through the official Docker hub. The name of the image is `minskylab/figport`. In order to up the service you need to define your configuration using env variables or a simple yaml file kinda:

   ```yaml
   secret: figportsecret # that's important because
   figma:
     accessToken: <YOUR_FIGMA_ACCESS_TOKEN>
   s3:
     endpoint: <YOUR_S3_ENDPOINT>
     accessKey: <YOUR_S3_ACCESS_KEY>
     secretKey: <YOUR_S3_SECRET_KEY>
     region: <YOUR_S3_REGION>
     bucket: <YOUR_S3_BUCKET_NAME>
   ```

   





## How it works

Use Figport for export your figma assets is very easy, you only need know the basic annotation principle for your exports:

For every **component** that you want to export to your s3 storage you need to annotate its name with the following nomenclature.

`{prefix}/{path/to/your/asset}{mods}{scales}`

Where:

- **prefix:** by default is figport, but you can modify this using env variables or configuration file of figport.
- **path/your/asset:** you need to indicate where will be saved your asset in your bucket, the last part of the path is the filename. Let's see an example: "/icons/feather/mail". Keep reading to understand how Figport adds the extension to your asset.
- **mods:** the mods are simple "add-ons" for your exportation routine. For add a mod in your asset annotation you should do with a **":"** prefix. (e.g. :png), and you can specify attributes for each with **"()"** (e.g. :svg(simplifyStroke=true)). Actually, Figport has 4 built-in: **:svg**, **:png**, **:jpg**, **:pdf**.
- **scales:** finally, the scales allow the ability to scale your component in a range between x0.01and x4. To add a scale for your component to assets exportation you just add **"@"** as a prefix following the scale (e.g. @1, @2, @0.5).

---

### Let's see an example:

In this figma file you can see two components being exported.

![image-20200813115422210](/Users/bregy/Library/Application Support/typora-user-images/image-20200813115422210.png)



The squares generate one file:

- squares.svg

And the circles generate three files:

- circles.png
- circles@2.png
- circles@4.png

The four files will be saved in your s3 bucket path. In this case (for the example) the file destination is **\<your-bucket-name\>/examples/**. And you can use these routes to access to your assets with your S3, for example, if you use Digital Ocean Spaces, you will have something like this: https://assets-minsky.sfo2.digitaloceanspaces.com/examples/circles.png.

#### Ok, but how perform a deployment


import type { NextApiRequest, NextApiResponse } from "next";
import nextConnect from "next-connect";
import multer from "multer";
import sharp from "sharp";
import fs from "fs";
import path from "path";

interface MulterRequest extends NextApiRequest {
  file: Express.Multer.File;
}

// Configure multer
const upload = multer({ dest: "/tmp" });

const apiRoute = nextConnect({
  onNoMatch(req: NextApiRequest, res: NextApiResponse) {
    res.status(405).json({ error: `Method '${req.method}' Not Allowed` });
  },
});

apiRoute.use(upload.single("file"));

apiRoute.post(async (req: MulterRequest, res: NextApiResponse) => {
  const file = req.file;
  if (!file) {
    return res.status(400).json({ error: "File is required." });
  }

  try {
    const outputPath = path.join(
      "public",
      "uploads",
      `${Date.now()}-${file.originalname}`
    );
    if (file.mimetype === "image/tiff") {
      await sharp(file.path).toFormat("png").toFile(outputPath);
      fs.unlinkSync(file.path); // Clean up tmp file
      return res.status(200).json({ url: outputPath.replace("public", "") });
    } else {
      return res.status(400).json({ error: "Unsupported file format." });
    }
  } catch (error) {
    return res.status(500).json({ error: "Internal server error." });
  }
});

export default apiRoute;
export const config = {
  api: {
    bodyParser: false,
  },
};

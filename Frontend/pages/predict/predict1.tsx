import Image from "next/image";
import { Inter } from "next/font/google";
import UploadFile from "@/components/uploadFile";
import Navbar from "@/components/Navbar";
import Transition from "@/components/Transitions";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    <div>
      <Transition />
      <div>
        <Navbar />
      </div>
    <div className="bg-gradient-background h-screen flex items-center justify-center">
      <UploadFile />
    </div>
    </div>
  );
}

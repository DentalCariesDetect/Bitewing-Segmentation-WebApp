import { Inter } from "next/font/google";
import UploadFile from "@/components/uploadFile";
import Navbar from "@/components/Navbar";
import Transition from "@/components/Transitions";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    <div className="h-full ">
      <Transition />
      <div >
        <Navbar />
      </div>
      <div className="bg-gradient-background h-full flex flex-col items-center justify-center ">
        <div className="flex flex-row space-x-2">
          <input
            className="rounded-md px-5 py-2 bg-inherit border mb-6"
            type="name"
            name=""
            placeholder="Name"
            required
          />

          <input
            className="rounded-md px-2 py-2 bg-inherit border mb-6"
            type="date"
            name=""
            placeholder="ชื่อ"
            required
          />

          {/* check box  */}
          <div className="flex items-center justify-center flex-cols -translate-y-3">
            <input
              type="checkbox"
              className="mr-2"
            />
            <p className="">วันนี้</p>
          </div>
        </div>
        <UploadFile />

      </div>
    </div>
  );
}

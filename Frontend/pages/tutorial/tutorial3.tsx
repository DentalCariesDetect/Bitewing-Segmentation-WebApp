import Navbar from "@/components/Navbar";
import Transition from "@/components/Transitions";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    // Full screen background
    <div className="bg-gradient-background h-full">
      <Transition />
      <div>
        <Navbar />
      </div>
      <div className=" flex justify-center items-center h-screen flex-col space-y-3">
        <h1 className="text-2xl font-thin w-[750px] text-center">Second begin uploading the bitewing file by pressing the &lsquo;Upload File&lsquo; button. 
        The files that can be uploaded should be in 
        the .jpeg, .png, or .tiff format.</h1>
      <div className=" w-[450px] h-[400px] bg-white rounded-2xl">
        <div className=" rounded-full shadow-white shadow-sm flex justify-center items-center h-full w-full">
            <div className=  "bg-white w-[200px] h-[50px] text-black text-custom rounded-full font-bold leading-normal text-shadow text-center flex flex-cols justify-center ">
                <h1 className="h-full flex justify-center items-center font-thin">Upload <br/>Bitewing X-ray</h1>
            </div>
        </div>
      </div>
      <div className="w-[450px] h-8 bg-white rounded-xl">

      </div>
     
      </div>
    </div>
  );
}

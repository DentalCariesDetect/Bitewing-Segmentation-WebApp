import { Inter } from "next/font/google";
import UploadFile from "@/components/uploadFile";
import Transition from "@/components/Transitions";
import NavbarMobile from "@/components/NavbarMobile";
import NavbarDesktop from "@/components/NavbarDesktop";


export default function Predict() {
  return (
    <div className="h-full">
      <Transition />
      <div className="block md:hidden">
        <NavbarMobile />
      </div>
      <div className="hidden md:block">
        <NavbarDesktop />
      </div>

      <div className="bg-gradient-background h-full flex flex-col  sm:flex-col  items-center justify-center">
        <div className="flex flex-col sm:flex-row space-x-2 mt-5 w-full items-center justify-center">

          {/* check box  */}
        </div >
        <div className="flex w-full h-screen items-center justify-center ">
          <UploadFile />
        </div>

      </div>
    </div>
  );
}

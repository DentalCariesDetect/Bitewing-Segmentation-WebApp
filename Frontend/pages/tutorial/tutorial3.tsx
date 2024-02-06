import Navbar from "@/components/Navbar";
import Image from 'next/image';
import Transition from "@/components/fastTransition";
import { Inter } from "next/font/google";
import Link from "next/link";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    // Full screen background
    <div className="flex flex-col h-screen bg-gradient-background">
      {/* <Transition /> */}
      <div>
        <Navbar />
      </div>

      <div className="flex justify-between items-center flex-row space-y-3 h-screen lg:m-16 sm:m-8 nn:m-2">
        <Link href="/tutorial/tutorial2" className="sm:invisible xl:visible nn:invisible">
          <div className="rounded-full bg-white xl:w-24 xl:h-24 sm:invisible xl:visible flex justify-center items-center">
            Previous
          </div>
        </Link>

        <div className=" flex justify-center items-center flex-col space-y-3">
          <h1 className="text-2xl nn:text-sm font-thin text-center text-white">Third once the file has been uploaded, press the `&apos;`START PREDICT`&apos;` button.</h1>
          <div className="xl:w-[450px] xl:h-[400px] sm:w-full sm:h-[450px] bg-white rounded-2xl">
            <div className=" rounded-full shadow-white shadow-sm flex justify-center items-center h-full w-full">
              <Image
                className=" items-center justify-center rounded-2xl"
                src="/image/tu3.png"
                alt="Picture of the author"
                width={400}
                height={400}
              />
            </div>
          </div>
          <div className="w-[450px] h-8 bg-white rounded-xl nn:w-full">
            <h1 className="h-full flex justify-center items-center font-thin">FileBitewing.tiff</h1>
          </div>

          <div className="bg-slate-300 w-[200px] h-[50px] text-black text-custom rounded-full font-bold leading-normal text-shadow text-center flex justify-center ">
            <h1 className="h-full flex justify-center items-center">START PREDICT</h1>
          </div>

          <div className="flex justify-between p-8 w-11/12">
            <Link href="/tutorial/tutorial2" className="sm:visible xl:invisible">
              <div className="rounded-full bg-white sm:visible xl:invisible flex justify-center items-center sm:h-[100px] sm:w-[140px] nn:px-4 nn:py-2">
                Previous
              </div>
            </Link>

            <Link href="/tutorial/tutorial4" className="sm:visible xl:invisible">
              <div className="rounded-full bg-white sm:visible xl:invisible flex justify-center items-center sm:h-[100px] sm:w-[140px] nn:px-4 nn:py-2">
                Next
              </div>
            </Link>
          </div>

        </div>

        <Link href="/tutorial/tutorial4" className="nn:invisible sm:invisible xl:visible">
          <div className="rounded-full bg-white xl:w-24 xl:h-24 sm:invisible xl:visible flex justify-center items-center">
            Next
          </div>
        </Link>

      </div>


    </div>
  );
}

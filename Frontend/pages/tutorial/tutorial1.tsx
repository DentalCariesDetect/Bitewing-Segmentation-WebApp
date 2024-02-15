import NavbarDesktop from "@/components/NavbarDesktop";
import NavbarMobile from "@/components/NavbarMobile";
import Transition from "@/components/Transitions";
import { Inter } from "next/font/google";
import Link from "next/link";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    // Full screen background
    <div className="flex flex-col h-screen bg-gradient-background">
      <Transition />

      <div className="sticky top-0">
        <div className="block md:hidden">
          <NavbarMobile />
        </div>
        <div className="hidden md:block">
          <NavbarDesktop />
        </div>
      </div>

      <div className="flex justify-between items-center flex-row space-y-3 h-screen lg:m-16 sm:m-8">


        <div className="rounded-full bg-white xl:w-24 xl:h-24 invisible flex justify-center items-center">
          invisible
        </div>


        <div className="flex justify-center items-center flex-col space-y-3">

          <h1 className="text-2xl font-thin text-white">
            First Press the &quot;Start Prediction&quot; button on our webpage.
          </h1>

          <div className="xl:w-[450px] xl:h-[400px] sm:w-full sm:h-[450px] bg-white rounded-2xl">
            <div className=" rounded-full shadow-white shadow-sm flex justify-center items-center h-full w-full">

              <div className="bg-slate-300 w-[200px] h-[50px] text-black text-custom rounded-full font-bold leading-normal text-shadow text-center flex justify-center ">
                <h1 className="h-full flex justify-center items-center">START PREDICT</h1>
              </div>

            </div>
          </div>

          <div className="flex justify-between p-8 w-11/12 mt-1">
            <Link href="/tutorial/tutorial1" className="invisible">
              <div className="rounded-full bg-white invisible flex justify-center items-center sm:h-[66px] sm:w-[82px] text-black">
                Next
              </div>
            </Link>

            <Link href="/tutorial/tutorial2" className="sm:visible xl:invisible">
              <div className="rounded-full bg-white sm:visible xl:invisible flex justify-center items-center sm:h-[100px] sm:w-[140px] text-black">
                Next
              </div>
            </Link>
          </div>


        </div>

        <Link href="/tutorial/tutorial3" className="sm:invisible xl:visible">
          <div className="rounded-full bg-white xl:w-24 xl:h-24 sm:invisible xl:visible flex justify-center items-center text-black">
            Next
          </div>
        </Link>

      </div>



    </div>
  );
}

import Navbar from "@/components/Navbar";
import Transition from "@/components/Transitions";
import { Inter } from "next/font/google";
import Link from "next/link";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    // Full screen background
    <div className="flex flex-col h-screen w-screen bg-gradient-background">
      <Transition />

      <div className="sticky top-0">
        <Navbar />
      </div>

      <div className="flex justify-between items-center flex-row space-y-3 h-screen lg:m-16 sm:m-8">


        <div className="rounded-full bg-white invisible flex justify-center items-center xl:w-24 xl:h-24">
          invisible
        </div>


        <div className="flex justify-center items-center flex-col space-y-3">

          <h1 className="sm:text-2xl nn:text-lg font-thin text-white">
            First Press the &quot;Start Prediction&quot; button on our webpage.
          </h1>

          <div className="bg-white rounded-2xl xl:w-[450px] xl:h-[400px] sm:w-full sm:h-[450px] nn:w-full nn:py-24">
            <div className=" rounded-full shadow-white shadow-sm flex justify-center items-center h-full w-full">

              <div className="bg-slate-300 w-[200px] h-[50px] nn:w-[150px] nn:h-[35px] text-black text-custom rounded-full font-bold leading-normal text-shadow text-center flex justify-center ">
                <h1 className="h-full flex justify-center items-center nn:text-xs">START PREDICT</h1>
              </div>

            </div>
          </div>

          <div className="flex justify-between p-8 w-11/12 mt-1">
            <Link href="/tutorial/tutorial1" className="invisible">
              <div className="rounded-full bg-white invisible flex justify-center items-center nn:px-4 nn:py-2">
                Next
              </div>
            </Link>

            <Link href="/tutorial/tutorial2" className="sm:visible xl:invisible">
              <div className="rounded-full bg-white nn:visible sm:visible xl:invisible flex justify-center items-center sm:h-[100px] sm:w-[140px] nn:px-4 nn:py-2">
                Next
              </div>
            </Link>
          </div>


        </div>

        <Link href="/tutorial/tutorial2" className="nn:invisible sm:invisible md:invisible xl:visible">
          <div className="rounded-full bg-white xl:w-24 xl:h-24 flex justify-center items-center">
            Next
          </div>
        </Link>

      </div>



    </div>
  );
}

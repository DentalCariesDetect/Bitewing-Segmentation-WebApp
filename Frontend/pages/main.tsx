import Image from 'next/image';
import Link from 'next/link';
import Navbar from '@/components/Navbar';
import { motion } from 'framer-motion';
import { fadeIn } from '@/variants';
import Transition from '@/components/Transitions';

// import HoverCard from "@/components/HoverCard";
// import Modal from "@/components/Modal";

// const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return (
    // Full screen background
    <div className="h-full bg-gradient-background ">
      <Transition />
      {/* Navbar  */}
      <div>
        <Navbar />
      </div>
      <div className=" lg:grid lg:grid-cols-2">
        <div className=" ml-20 flex flex-col">
          <motion.h1
            variants={fadeIn('right', 0.5)}
            initial="hidden"
            animate="show"
            exit="hidden"
            className="text-5xl font-normal leading-normal text-white "
          >
            Convolution  Neural Network for Dental Caries Classification
          </motion.h1>

          <motion.div
            variants={fadeIn('right', 0.5)}
            initial="hidden"
            animate="show"
            exit="hidden"
            className=""
          >
            <div className="text-xl font-normal leading-normal text-white">
              Convolutional Neural Network (CNN) technology designed to
              revolutionize the way dental caries (tooth decay) is detected and
              classified. Our system is engineered to provide rapid, accurate,
              and user-friendly detection and categorization of dental caries,
              from early onset to advanced decay, aiding you and your dentist to
              make the best-informed decisions for your dental treatments.
            </div>

            <div className=" mt-5 flex flex-row justify-center space-x-10">
              <Link href="/predict/predict1">
                <button className="h-[51px] w-[300px] rounded-full bg-white shadow-sm shadow-white ">
                  <div className="text-custom text-shadow font-bold leading-normal text-black">
                    START PREDICT
                  </div>
                </button>
              </Link>

              <Link href="/tutorial/tutorial1">
                <button className="h-[51px] w-[300px] rounded-full bg-white shadow-sm shadow-white">
                  <div className="text-custom text-shadow font-bold leading-normal text-black">
                    TUTORIAL
                  </div>
                </button>
              </Link>
            </div>
          </motion.div>
        </div>

        <div className="flex items-center justify-center">
          <Image
            className=" hidden items-center justify-center rounded-2xl lg:flex "
            src="/image/dentistDoJob.png"
            alt="Picture of the author"
            width={500}
            height={500}
          />
        </div>
      </div>
      <div className=" grid-cols-10 ">
        <motion.div
          variants={fadeIn('up', 0.5)}
          initial="hidden"
          animate="show"
          exit="hidden"
          className=" xl:flex-rows col-span-9 ml-20 mr-20 mt-10 justify-center space-x-28 md:grid md:grid-cols-2 xl:flex "
        >
          {/* box1 */}
          <div className="mb-10 h-[280px] w-[350px] rounded-3xl bg-white bg-opacity-70 shadow-lg shadow-white">
            <h1 className="mt-5 text-center text-2xl text-black opacity-100">
              Privacy-Centric
            </h1>
            <p className="text-shadow m-5 text-base font-light leading-normal text-black">
              Your privacy is paramount. All uploads are handled with the
              highest level of security and confidentiality.
            </p>
            <div className="flex items-center justify-center">
              <Image
                className=" items-center justify-center rounded-2xl"
                src="/image/cyberImage.png"
                alt="Picture of the author"
                width={130}
                height={130}
              />
            </div>
          </div>
          {/* box2 */}
          <div className="h-[280px] w-[350px] rounded-3xl bg-white bg-opacity-70 shadow-lg shadow-white">
            <h1 className="mt-5 text-center text-2xl text-black">
              Easy Accessibility
            </h1>
            <p className="text-shadow m-5 text-base font-light leading-normal text-black">
              Access our services from anywhere, anytime, with just an internet
              connection.
            </p>
            <div className="flex items-center justify-center">
              <Image
                className=" mt-5 items-center justify-center rounded-2xl"
                src="/image/easyUse.png"
                alt="Picture of the author"
                width={90}
                height={90}
              />
            </div>
          </div>
          {/* box3 */}
          <div className="h-[280px] w-[350px] rounded-3xl bg-white bg-opacity-70 shadow-lg shadow-white">
            <h1 className="mt-5 text-center text-2xl text-black">
              Support for Dental
            </h1>
            <p className="text-shadow mb-7 ml-5 mr-5 mt-5 text-base font-light leading-normal text-black">
              This website can serve as an adjunct tool for dentists, aiding in
              their diagnostic processes and ensuring a comprehensive approach
              to oral health care.
            </p>
            <div className="flex items-center justify-center">
              <Image
                className=" -mt-6 items-center justify-center rounded-2xl"
                src="/image/dentistImage.png"
                alt="Picture of the author"
                width={90}
                height={90}
              />
            </div>
          </div>

          <div className=" right-24 flex items-end justify-end">
            <Image
              className=" "
              src="/image/tooth.png"
              alt="Picture of the author"
              width={150}
              height={150}
            />
          </div>
        </motion.div>
        {/* <div>
          <HoverCard
            title={"Support for Dental Professionals"}
            message={
              "This website can serve as an adjunct tool for dentists, aiding in their diagnostic processes and ensuring a comprehensive approach to oral health care."
            }
          />
        </div> */}
        {/* <Modal /> */}
      </div>
    </div>
  );
}

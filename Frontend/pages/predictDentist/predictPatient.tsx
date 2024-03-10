import UploadPatient from "@/components/uploadPatient";
import Transition from "@/components/Transitions";
import NavbarMobile from "@/components/NavbarMobile";
import NavbarDesktop from "@/components/NavbarDesktop";


export default function PredictPatient() {
    return (
        <div className="h-full">
            <Transition />
            <div className="block md:hidden">
                <NavbarMobile />
            </div>
            <div className="hidden md:block">
                <NavbarDesktop />
            </div>

            <div className="bg-gradient-background h-screen flex flex-col  sm:flex-col  items-center justify-center">
                <div className="flex w-full items-center justify-center ">
                    <UploadPatient />
                </div>

            </div>
        </div>
    );
} 

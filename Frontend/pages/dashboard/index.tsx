import Link from "next/link";
import { headers, cookies } from "next/headers";
import { redirect } from "next/navigation";
import { SideBar } from "@/components/SideBar";
import Transition from "@/components/Transitions";

export default function profileSetting() {

    return (
        <div className="w-full h-screen">
            <Transition />
            <div className="bg-gradient-background h-full flex items-center justify-between">
                <SideBar />
                <div>

                </div>
            </div>
        </div>

    );
}

import React, { useState } from 'react'; // Ensure React is imported for JSX usage and useState is available
import Link from "next/link";
import { SideBar } from "@/components/SideBar";
import Transition from "@/components/Transitions";



export default function ProfileSetting() {

    const [open, setOpen] = React.useState(false);

    return (
        <div className="w-full h-screen">
            <div className="bg-gradient-background h-full flex items-center justify-between">
                <SideBar />
                <div className=' flex flex-col space-y-5'>
                    <div>
                        <div className=' w-full h-60 bg-black rounded-xl'>
                            Hello
                        </div>
                    </div>
                    <div className=' flex flex-row space-x-5 '>
                        <div className=' w-[300px] h-32 bg-black rounded-xl'>
                            Hello
                        </div>
                        <div className=' w-[300px] h-32 bg-black rounded-xl'>
                            Hello
                        </div>
                        <div className=' w-[300px] h-32 bg-black rounded-xl'>
                            Hello
                        </div>
                    </div>

                </div>
                <div></div>
            </div>
        </div>

    );
}

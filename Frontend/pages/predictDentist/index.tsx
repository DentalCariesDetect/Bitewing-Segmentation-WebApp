import { useState } from 'react';
import router from 'next/router';
import Transition from "@/components/Transitions";
import NavbarMobile from "@/components/NavbarMobile";
import NavbarDesktop from "@/components/NavbarDesktop";
import { motion } from "framer-motion";
import Image from 'next/image';
import { fadeIn } from "@/variants";
import Modal from "@/components/Modal";


export default function Predict() {
    const [patientId, setPatientId] = useState('');
    const [birthOfDate, setBirthOfDate] = useState('');
    const [gender, setGender] = useState('');
    const [age, setAge] = useState('');
    const [profileImage, setProfileImage] = useState(null);
    const [previewImage, setPreviewImage] = useState('');
    const [openConfirm, setOpenConfirm] = useState(false)
    const [openPredict, setOpenPredict] = useState(false)

    // Handle form submission
    const handleSubmit = (e: any) => {
        e.preventDefault();
        // Here you would typically handle the form submission, including the image file
        console.log("Submitted: ", { patientId, birthOfDate, gender, age, });

        // To upload the image, you'll need to use FormData and append the file
        // This is typically done when sending the data to your backend server
    };

    // Handle profile image change
    const handleImageChange = (e: any) => {
        const file = e.target.files[0];
        if (file) {
            setProfileImage(file);
            setPreviewImage(URL.createObjectURL(file));
        }
    };
    return (
        <div className="h-full">
            <Modal
                isOpen={openConfirm}
                setIsOpen={setOpenConfirm}
                title="Add patient detail"
                message="Are you sure you want to add patient detail?"
                onUnderstood={() => setOpenPredict(true)}
                status={"info"}
            />

            <Modal
                isOpen={openPredict}
                setIsOpen={setOpenPredict}
                title="Predict now"
                message="Are you want to predict now"
                onUnderstood={() => router.push('/predictDentist/predict')}
                status={"info"}
            />
            <Transition />
            <div className="block md:hidden">
                <NavbarMobile />
            </div>
            <div className="hidden md:block">
                <NavbarDesktop />
            </div>

            <div className="bg-gradient-background h-full flex flex-col  sm:flex-col  items-center justify-center">
                <div className="flex flex-col sm:flex-row space-x-2 mt-5 w-full h-screen items-center justify-center">
                    <motion.div
                        variants={fadeIn('right', 0.05)}
                        initial="hidden"
                        animate="show"
                        exit="hidden"
                        className="w-[500px] h-[570px] bg-blue-900 p-8 rounded-xl">
                        {/* Profile setting page */}
                        <form onSubmit={handleSubmit} className="space-y-5">
                            {/* Profile Image Upload */}
                            <div className="flex flex-col items-center">
                                <label htmlFor="profileImage" className="cursor-pointer">
                                    {previewImage ? (
                                        <Image src={previewImage} alt="Profile preview" width={24} height={24} className="w-24 h-24 rounded-full object-cover" />
                                    ) : (
                                        <div className="w-24 h-24 rounded-full flex items-center justify-center border-2 border-dashed border-gray-300">
                                            <span className="text-gray-200">Add Image</span>
                                        </div>
                                    )}
                                </label>
                                <input id="profileImage" name="profileImage" type="file" accept="image/*" onChange={handleImageChange} className="hidden" />
                            </div>

                            {/* Firstname Field */}
                            <div>
                                <label htmlFor="patientId" className="block text-sm font-medium text-gray-200">Patient ID</label>
                                <input type="text" id="name" name="name" value={patientId} placeholder=" Patient ID" onChange={(e) => setPatientId(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                            </div>

                            {/* Lastname Field */}
                            <div>
                                <label htmlFor="birthOfDate" className="block text-sm font-medium text-gray-200">Birth of date</label>
                                <input type="text" id="birthOfDate" name="birthOfDate" value={birthOfDate} placeholder=" Birth of date" onChange={(e) => setBirthOfDate(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                            </div>

                            {/* Hospital Field */}
                            <div>
                                <label htmlFor="gender" className="block text-sm font-medium text-gray-200">Gender</label>
                                <input type="text" id="gender" name="gender" value={gender} placeholder="Gender" onChange={(e) => setGender(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                            </div>

                            {/* Email Field */}
                            <div className="">
                                <label htmlFor="age" className="block text-sm font-medium text-gray-200">Age</label>
                                <input type="number" id="age" name="age" value={age} placeholder="Age" onChange={(e) => setAge(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                            </div>

                            {/* Submit Button */}
                            <div className=" translate-y-5 flex flex-row space-x-5">
                                <button className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                                    Decline
                                </button>
                                <button onClick={() => setOpenConfirm(true)} type="submit" className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                                    Accept
                                </button>
                            </div>


                        </form>
                    </motion.div>
                </div>

            </div>
        </div>
    );
}

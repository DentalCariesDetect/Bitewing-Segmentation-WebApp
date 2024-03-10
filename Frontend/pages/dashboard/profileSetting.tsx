import { useState } from 'react';
import { SideBar } from '@/components/SideBar';
import { motion } from 'framer-motion';
import Image from 'next/image';
import { fadeIn } from '@/variants';
import Transition from '@/components/Transitions';


export default function ProfileSetting() {
    const [firstname, setFirstname] = useState('');
    const [lastname, setLastname] = useState('');
    const [hostpital, setHospital] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [profileImage, setProfileImage] = useState(null);
    const [previewImage, setPreviewImage] = useState('');

    // Handle form submission
    const handleSubmit = (e: any) => {
        e.preventDefault();
        // Here you would typically handle the form submission, including the image file
        console.log("Submitted: ", { firstname, lastname, hostpital, email, password });

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
        <div className="w-full h-screen">
            <div className="bg-gradient-background h-full flex items-center justify-between">
                <SideBar />
                <motion.div
                    variants={fadeIn('right', 0.05)}
                    initial="hidden"
                    animate="show"
                    exit="hidden"
                    className="w-[500px] h-[620px] bg-blue-900 p-8 rounded-xl">
                    {/* Profile setting page */}
                    <form onSubmit={handleSubmit} className="space-y-4">
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
                            <label htmlFor="name" className="block text-sm font-medium text-gray-200">Firstname</label>
                            <input type="text" id="name" name="name" value={firstname} placeholder=" Your firstname" onChange={(e) => setFirstname(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                        </div>

                        {/* Lastname Field */}
                        <div>
                            <label htmlFor="name" className="block text-sm font-medium text-gray-200">Lastname</label>
                            <input type="text" id="name" name="name" value={lastname} placeholder=" Your lastname" onChange={(e) => setLastname(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                        </div>

                        {/* Hospital Field */}
                        <div>
                            <label htmlFor="hospital" className="block text-sm font-medium text-gray-200">Hospital</label>
                            <input type="text" id="hospital" name="hospital" value={hostpital} placeholder="Hospital" onChange={(e) => setHospital(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                        </div>

                        {/* Email Field */}
                        <div>
                            <label htmlFor="email" className="block text-sm font-medium text-gray-200">Username</label>
                            <input type="email" id="email" name="email" value={email} placeholder="Your username" onChange={(e) => setEmail(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                        </div>
                        {/* Password Field */}
                        <div>
                            <label htmlFor="password" className="block text-sm font-medium text-gray-200">Password</label>
                            <input type="password" id="password" name="password" value={password} placeholder="Your password" onChange={(e) => setPassword(e.target.value)} className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 text-black h-10" />
                        </div>


                        {/* Submit Button */}
                        <div>
                            <button type="submit" className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                                Update Profile
                            </button>
                        </div>
                    </form>
                </motion.div>
                <div></div>
            </div>
        </div>
    );
}

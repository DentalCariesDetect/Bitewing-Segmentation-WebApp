import { AnimatePresence, motion } from "framer-motion";
import { FiLoader } from "react-icons/fi";
import React from "react";
import { SideBar } from '@/components/SideBar';

interface LoadingPopupProps {
    isOpen: boolean;
}

const Loading: React.FC<LoadingPopupProps> = ({ isOpen }) => {
    return (
        <AnimatePresence>
            {isOpen && (
                <motion.div
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    exit={{ opacity: 0 }}
                    className="bg-slate-900/30 backdrop-blur-lg flex justify-center items-center fixed inset-0 z-50"
                >
                    <motion.div
                        initial={{ scale: 0 }}
                        animate={{ scale: 1 }}
                        exit={{ scale: 0 }}
                        className="flex flex-col justify-center items-center bg-white p-6 rounded-lg shadow-xl cursor-default"
                    >
                        <motion.div
                            animate={{ rotate: 360 }}
                            transition={{ repeat: Infinity, duration: 1, ease: "linear" }}
                            className="text-indigo-600 text-4xl"
                        >
                            <FiLoader />
                        </motion.div>
                        <p className="mt-2 text-black">Loading...</p>
                    </motion.div>
                </motion.div>
            )}
        </AnimatePresence>
    );
};

export default Loading;

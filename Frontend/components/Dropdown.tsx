import React, { useState } from 'react';
import { motion, AnimatePresence } from 'framer-motion';

interface DropdownProps {
    options: string[];
    selectedOption: string;
    onOptionChange: (option: string) => void;
}

const Dropdown: React.FC<DropdownProps> = ({ options, selectedOption, onOptionChange }) => {
    const [isOpen, setIsOpen] = useState(false);

    const dropdownVariants = {
        hidden: {
            opacity: 0,
            scaleY: 0,
        },
        visible: {
            opacity: 1,
            scaleY: 1,
            transition: {
                duration: 0.2,
                ease: 'easeInOut',
            }
        }
    };

    return (
        <div className="relative inline-block w-full md:w-64 text-gray-700">
            <div
                className="w-full bg-gray-50 border border-gray-300 py-2 px-4 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500 hover:border-gray-400 shadow-md cursor-pointer"
                onClick={() => setIsOpen(!isOpen)}
            >
                {selectedOption}
                <span className="absolute inset-y-0 right-0 flex items-center px-2 pointer-events-none">
                    <svg className="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" /></svg>
                </span>
            </div>
            <AnimatePresence>
                {isOpen && (
                    <motion.div
                        variants={dropdownVariants}
                        initial="hidden"
                        animate="visible"
                        exit="hidden"
                        className="absolute z-10 w-full bg-white mt-1 border border-gray-300 rounded shadow-lg"
                    >
                        {options.map((option) => (
                            <div
                                key={option}
                                className="py-2 px-4 hover:bg-gray-100 cursor-pointer"
                                onClick={() => {
                                    onOptionChange(option);
                                    setIsOpen(false);
                                }}
                            >
                                {option}
                            </div>
                        ))}
                    </motion.div>
                )}
            </AnimatePresence>
        </div>
    );
};

export default Dropdown;

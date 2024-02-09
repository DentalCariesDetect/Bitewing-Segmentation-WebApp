import React, { useState } from "react";
import Link from "next/link";
import { motion } from 'framer-motion';

const NavbarMobile = () => {
  // State for managing the dropdown menu visibility
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  // Variants for the dropdown animation
  const dropdownVariants = {
    hidden: {
      opacity: 0,
      height: 0,
      overflow: 'hidden',
      transition: {
        duration: 0.2,
        ease: 'easeInOut'
      }
    },
    visible: {
      opacity: 1,
      height: 'auto',
      overflow: 'visible',
      transition: {
        duration: 0.2,
        ease: 'easeInOut'
      }
    }
  };

  return (
    <nav className="bg-navbar-purple p-4">
      <div className="container mx-auto flex items-center justify-between flex-wrap">
        <div className="text-white flex-1">
          <Link href="https://sites.google.com/view/sf341smokybite/home">
            <h1 className="text-lg font-semibold lg:text-xl">SMOKYBIKE</h1>
          </Link>
        </div>
        {/* Mobile Menu Button */}
        <div className="lg:hidden">
          <button onClick={() => setIsDropdownOpen(!isDropdownOpen)} className="focus:outline-none">
            <svg className="w-6 h-6 text-white" fill="none" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" viewBox="0 0 24 24" stroke="currentColor">
              {/* Adjusted path for the third line to match the full width of the first two lines */}
              <path d="M4 6h16 M4 12h16 M4 18h16"></path>
            </svg>
          </button>
        </div>


        <motion.div
          variants={dropdownVariants}
          initial="hidden"
          animate={isDropdownOpen ? "visible" : "hidden"}
          className="w-full lg:flex lg:items-center lg:w-auto"
        >
          <ul className="flex flex-col lg:flex-row lg:space-x-4 mt-4 lg:mt-0">
            <li>
              <Link href="../main">
                <h1 className="text-gray-300 hover:text-white block mt-4 lg:inline-block lg:mt-0 text-sm md:text-base lg:text-lg">HOME</h1>
              </Link>
            </li>
            <li>
              <Link href="../tutorial/tutorial1">
                <h1 className="text-gray-300 hover:text-white block mt-4 lg:inline-block lg:mt-0 text-sm md:text-base lg:text-lg">TUTORIAL</h1>
              </Link>
            </li>
            <li>
              <Link href="../predict/">
                <h1 className="text-gray-300 hover:text-white block mt-4 lg:inline-block lg:mt-0 text-sm md:text-base lg:text-lg">PREDICT</h1>
              </Link>
            </li>
          </ul>
        </motion.div>
      </div>
    </nav>
  );
};

export default NavbarMobile;

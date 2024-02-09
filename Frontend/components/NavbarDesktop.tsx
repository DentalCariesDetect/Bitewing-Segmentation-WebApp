import React from "react";
import Link from "next/link";

const NavbarDesktop = () => {
    return (
        <nav className="bg-navbar-purple p-4">
            <div className="container mx-auto flex items-center justify-between flex-wrap">
                <div className="text-white flex-1">
                    <Link href="https://sites.google.com/view/sf341smokybite/home">
                        <h1 className="text-lg font-semibold lg:text-xl">SMOKYBIKE</h1>
                    </Link>
                </div>
                <div>
                    <ul className="flex space-x-2 md:space-x-4">
                        <li>
                            <Link href="../main">
                                <h1 className="text-gray-300 hover:text-white text-sm md:text-base lg:text-lg">HOME</h1>
                            </Link>
                        </li>
                        <li>
                            <Link href="../tutorial/tutorial1">
                                <h1 className="text-gray-300 hover:text-white text-sm md:text-base lg:text-lg">TUTORIAL</h1>
                            </Link>
                        </li>
                        <li>
                            <Link href="../predict/">
                                <h1 className="text-gray-300 hover:text-white text-sm md:text-base lg:text-lg">PREDICT</h1>
                            </Link>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    );
};

export default NavbarDesktop;
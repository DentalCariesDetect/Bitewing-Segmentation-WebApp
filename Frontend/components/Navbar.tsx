import React from "react";
import Link from "next/link";

const Navbar = () => {
  return (
    <nav className="bg-navbar-purple p-4">
      <div className="container mx-auto flex items-center justify-between">
        <div className="text-white">
          <Link href="https://sites.google.com/view/sf341smokybite/home">
            <div className="text-lg font-semibold">SMOKYBIKE</div>
          </Link>
        </div>
        <div>
          <ul className="flex space-x-4">
            <li>
              <Link href="../main">
                <div className="text-gray-300 hover:text-white">HOME</div>
              </Link>
            </li>
            <li>
              <Link href="../tutorial/tutorial1">
                <div className="text-gray-300 hover:text-white">TOTORIAL</div>
              </Link>
            </li>
            <li>
              <Link href="../predict/predict1">
                <div className="text-gray-300 hover:text-white">PREDICT</div>
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;

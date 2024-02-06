import React from "react";
import Image from "next/image";

// create dropdown component code
interface DropdownProps {
    options: string[];
    selectedOption: string;
    onOptionChange: (newOption: string) => void;
}

const Dropdown: React.FC<DropdownProps> = ({ options, selectedOption, onOptionChange }) => {
    return (
        <select
            className="rounded-md px-4 py-1 border mb-6"
            value={selectedOption}
            onChange={(e) => onOptionChange(e.target.value)}
            required
        >
            {options.map((option) => (
                <option key={option} value={option}>
                    {option}
                </option>
            ))}
        </select>
    );
};

export default Dropdown;

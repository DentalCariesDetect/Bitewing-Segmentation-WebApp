import React, { useEffect, useState } from "react";
import Image from "next/image";
import Dropdown from "./Dropdown";

interface ImageTableProps {
    images: string[];
}

const ImageTable: React.FC<ImageTableProps> = ({ images }) => {
    const [selectedImages, setSelectedImages] = useState<boolean[]>(Array(images.length).fill(false));

    useEffect(() => {
        setSelectedImages(Array(images.length).fill(false));
    }, [images]);

    const onImageSelect = (index: number, checked: boolean) => {
        const newSelectedImages = [...selectedImages];
        newSelectedImages[index] = checked;
        setSelectedImages(newSelectedImages);
    };

    return (
        <div className="bg-indigo-600 p-4 w-full ">
            <div className="flex flex-row space-x-5">
                <div className=" ">
                    {images.map((img, index) => (
                        <div key={index} className="flex  mb-4 border border-white w-80 p-3 rounded-lg">
                            <input
                                type="checkbox"
                                checked={selectedImages[index] || false} // Handle potential undefined values
                                onChange={(e) => onImageSelect(index, e.target.checked)}
                                className="mr-2"
                            />
                            <Image
                                src={img}
                                alt="Preview"
                                width={100}
                                height={100}
                                className="max-w-xs rounded-md"
                            />
                            <div className=" w-full ml-2">
                                <h1>Class 31</h1>
                                <p>confident : 0.99</p>
                            </div>
                            <button className="flex justify-end items-end underline">แก้ไข</button>

                        </div>
                    ))}
                </div>
                <div className=" flex flex-col">
                    {/* Dropdown example code */}
                    <div className=" flex flex-row space-x-3">
                        <h1 className=" text-center flex justify-center ">Carious Diagnosis</h1>
                        <select
                            className="rounded-md px-5 py-2 bg-inherit border mb-6 "
                            required
                        >
                            <option value="ra1">RA1</option>
                            <option value="ra2">RA2</option>
                            <option value="ra3">RA3</option>
                            <option value="rb4">RB4</option>
                            <option value="rc5">RC5</option>
                            <option value="rd6">RD6</option>
                        </select>
                    </div>
                    {/* <Dropdown options={["RA1", "RA2", "RA3", "RB4", "RC5", "RD6"]} selectedOption="RA1" onOptionChange={(newOption) => console.log(newOption)} /> */}


                    {/* Description box */}
                    <div className=" sm:w-[400px] w-full">
                        <h1>Description</h1>
                        {/* multiple line input4*/}
                        <textarea
                            className="rounded-md px-5 py-2 bg-inherit border mb-6 h-80 w-full"
                            placeholder="Description"
                            required
                        />
                    </div>

                    <div className=" sm:w-[400px] w-full">
                        <h1>Treatment approach</h1>
                        {/* multiple line input */}
                        <textarea
                            className="rounded-md px-5 py-2 bg-inherit border mb-6 h-80 w-full"
                            placeholder="Treatment approach"
                            required
                        />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default ImageTable;

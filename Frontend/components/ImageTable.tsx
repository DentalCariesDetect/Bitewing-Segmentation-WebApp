import React, { useEffect, useState } from "react";
import Image from "next/image";

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
        <div className="bg-indigo-600 p-4">
            <div>
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
            </div>
        </div>
    );
};

export default ImageTable;

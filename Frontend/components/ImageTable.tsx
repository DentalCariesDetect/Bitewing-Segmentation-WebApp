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
                <div className=" grid grid-cols-6">
                    {images.map((img, index) => (
                        <div key={index} className="flex items-center mb-4 ">
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
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
};

export default ImageTable;

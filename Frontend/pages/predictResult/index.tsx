import React, { useState } from 'react';
import Transition from '@/components/Transitions';
import NavbarMobile from '@/components/NavbarMobile';
import NavbarDesktop from '@/components/NavbarDesktop';

// ข้อมูลสำหรับรูปภาพหลักและรูปย่อย
const mainImages = [
    { id: 'main1', imageUrl: 'https://via.placeholder.com/500x300?text=Main+Image+1' },
    { id: 'main2', imageUrl: 'https://via.placeholder.com/500x300?text=Main+Image+2' },
    { id: 'main3', imageUrl: 'https://via.placeholder.com/500x300?text=Main+Image+3' }
];

const thumbnailImages = [
    { id: 'thumb1', relatedTo: 'main1', imageUrl: 'https://via.placeholder.com/100?text=Thumb+1', details: 'Detail for Thumb 1' },
    { id: 'thumb2', relatedTo: 'main1', imageUrl: 'https://via.placeholder.com/100?text=Thumb+2', details: 'Detail for Thumb 2' },
    { id: 'thumb3', relatedTo: 'main1', imageUrl: 'https://via.placeholder.com/100?text=Thumb+3', details: 'Detail for Thumb 3' },
    { id: 'thumb4', relatedTo: 'main1', imageUrl: 'https://via.placeholder.com/100?text=Thumb+4', details: 'Detail for Thumb 4' },
    { id: 'thumb5', relatedTo: 'main1', imageUrl: 'https://via.placeholder.com/100?text=Thumb+5', details: 'Detail for Thumb 5' },
    // เพิ่มเติมตามจำนวนรูปภาพที่คุณมี

    { id: 'thumb4', relatedTo: 'main2', imageUrl: 'https://via.placeholder.com/100?text=Thumb+4', details: 'Detail for Thumb 1' },
    { id: 'thumb5', relatedTo: 'main2', imageUrl: 'https://via.placeholder.com/100?text=Thumb+5', details: 'Detail for Thumb 2' },

    { id: 'thumb4', relatedTo: 'main3', imageUrl: 'https://via.placeholder.com/100?text=Thumb+4', details: 'Detail for Thumb 1' },
    { id: 'thumb5', relatedTo: 'main3', imageUrl: 'https://via.placeholder.com/100?text=Thumb+5', details: 'Detail for Thumb 2' },
    { id: 'thumb5', relatedTo: 'main3', imageUrl: 'https://via.placeholder.com/100?text=Thumb+5', details: 'Detail for Thumb 3' },
];

export default function PredictResult() {
    const [selectedMainImage, setSelectedMainImage] = useState('');
    const [selectedThumbDetail, setSelectedThumbDetail] = useState('');

    // Handler สำหรับการคลิกที่รูปภาพหลัก
    const handleSelectMainImage = (imageId: string) => {
        setSelectedMainImage(imageId);
        // Reset details to default text
        setSelectedThumbDetail('Click on a thumbnail to see the details.');
    };

    return (
        <div className="h-full">
            <Transition />
            <div className="block md:hidden"><NavbarMobile /></div>
            <div className="hidden md:block"><NavbarDesktop /></div>
            <div className="bg-gradient-background flex flex-row items-start justify-start min-h-screen p-5">
                <div className="flex flex-col space-y-5 overflow-auto max-h-screen">
                    {mainImages.map((image) => (
                        <img
                            key={image.id}
                            src={image.imageUrl}
                            alt="Main"
                            className="cursor-pointer rounded-md"
                            onClick={() => handleSelectMainImage(image.id)}
                        />
                    ))}
                </div>
                <div className="flex flex-col">
                    <div className="flex flex-wrap justify-start items-start ml-5 space-x-2">
                        {thumbnailImages.filter(img => img.relatedTo === selectedMainImage).map((image) => (
                            <img
                                key={image.id}
                                src={image.imageUrl}
                                alt="Thumbnail"
                                className="m-1 rounded-md cursor-pointer"
                                onClick={() => setSelectedThumbDetail(image.details)}
                            />
                        ))}
                    </div>
                    <div className="flex-1 bg-white text-black p-4 m-4 rounded-md shadow-md">
                        <h3 className="font-bold">Details:</h3>
                        <p>{selectedThumbDetail}</p>
                    </div>
                </div>
            </div>
        </div>
    );
}

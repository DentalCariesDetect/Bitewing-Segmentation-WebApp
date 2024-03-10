import React from "react";
import Image from "next/image";

interface DetailsModalProps {
    isOpen: boolean;
    imageUrl?: string; // Optional prop for image URL
    details: string;
    onSave: () => void;
    onCancel: () => void;
    onDetailsChange: (details: string) => void;
}

const DetailsModal: React.FC<DetailsModalProps> = ({
    isOpen,
    imageUrl,
    details,
    onSave,
    onCancel,
    onDetailsChange,
}) => {
    if (!isOpen) return null;

    return (
        <>
            <div className="fixed inset-0 bg-black bg-opacity-50 z-40 w-full -translate-x-5"></div>
            <div className="modal fixed z-50 left-1/2 top-1/2 transform -translate-x-1/2 -translate-y-1/2 bg-white shadow-lg rounded-lg p-6 text-black max-w-lg w-11/12 md:w-1/2">
                {imageUrl && (
                    <div className="mb-4">
                        <Image
                            src={imageUrl}
                            width={200}
                            height={300}
                            alt="Preview"
                            className=" max-h-60 rounded-md object-cover"
                        />
                    </div>
                )}
                <textarea
                    className="w-full p-3 border border-gray-300 rounded mb-4"
                    value={details}
                    onChange={(e) => onDetailsChange(e.target.value)}
                    rows={5}
                    placeholder="Enter details here..."
                />
                <div className="flex justify-end gap-4">
                    <button
                        className="py-2 px-4 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 transition-colors duration-150"
                        onClick={onCancel}
                    >
                        Cancel
                    </button>
                    <button
                        className="py-2 px-4 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors duration-150"
                        onClick={onSave}
                    >
                        Save
                    </button>
                </div>
            </div>
        </>
    );
};

export default DetailsModal;

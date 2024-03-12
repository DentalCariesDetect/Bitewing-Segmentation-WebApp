import React, { useEffect, useState } from 'react';
import Image from 'next/image';

interface UserDetails {
    name: string;
    email: string;
    gender: string;
    phonenumber: string;
}

interface EditPredict {
    isOpen: boolean;
    setIsOpen: (isOpen: boolean) => void;
    initialUserDetails: UserDetails;
    onSave: (newDetails: UserDetails) => void;
}



const EditPredict: React.FC<EditPredict> = ({ isOpen, setIsOpen, initialUserDetails, onSave }) => {
    const [editedUser, setEditedUser] = useState<UserDetails>(initialUserDetails);

    useEffect(() => {
        if (isOpen) {
            setEditedUser(initialUserDetails); // Update state when initialUserDetails changes
        }
    }, [initialUserDetails, isOpen]);

    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black bg-opacity-50">

            <div className="w-full max-w-lg p-8 space-y-6 bg-white flex items-center justify-center flex-col rounded-lg shadow-2xl">
                <div className="text-black">
                    <h1>Predict page</h1>
                </div>
                {/* <div className="flex flex-col gap-4">
                    {Object.entries(editedUser).map(([key, value]) => (
                        <div key={key}>
                            <label htmlFor={key} className="block mb-2 text-sm font-medium text-gray-900 capitalize">{key}</label>
                            <input
                                type="text"
                                id={key}
                                name={key}
                                value={value}
                                onChange={e => setEditedUser(prev => ({ ...prev, [key]: e.target.value }))}
                                className="block w-full p-2.5 text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
                                placeholder={`Enter ${key}`}
                            />
                        </div>
                    ))}
                </div> */}
                <div className='w-96 h-96 bg-gray-500 rounded-full flex items-center justify-center text-2xl'>
                    <Image
                        src="/image/cropped-image.jpeg"
                        width={100}
                        height={100}
                        alt="Preview"
                        className=" w-96 h-96 rounded-lg"

                    />
                </div>
                <div className="flex justify-end gap-2 mt-4">
                    <button
                        onClick={() => setIsOpen(false)}
                        className="px-5 py-2 text-sm font-medium text-gray-500 bg-white border border-gray-300 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                        Cancel
                    </button>
                    <button
                        onClick={() => {
                            onSave(editedUser);
                            setIsOpen(false);
                        }}
                        className="px-5 py-2 text-sm font-medium text-white bg-blue-600 rounded-lg hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                    >
                        Predict again
                    </button>
                </div>
            </div>
        </div>
    );
};

export default EditPredict;
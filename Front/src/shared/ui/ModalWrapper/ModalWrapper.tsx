import React from 'react';
import Modal from '@mui/material/Modal';


interface UniversalModalProps {
    open: boolean;
    onClose: () => void;
    title?: string;
    children: React.ReactNode;
}

export const UniversalModal = ({ open, onClose, children }: UniversalModalProps) => {
    return (
        <Modal
            open={open}
            onClose={onClose}
            aria-labelledby="universal-modal-title"
            aria-describedby="universal-modal-description"
        >
            <div className="absolute translate-x-[-50%] translate-y-[-50%] top-1/2 left-1/2 bg-black bg-opacity-25">
                <div className="bg-gray-900 p-5 rounded-md text-white">
                    {children}
                </div>
            </div>
        </Modal>
    );
};
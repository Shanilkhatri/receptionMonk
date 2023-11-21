// import React, { useState, useRef, useEffect } from "react";

// interface ModalProps {
//   isOpen: boolean;
//   hasCloseBtn?: boolean;
//   onClose?: () => void;
// }

// const Modal: React.FC<ModalProps> = ({ isOpen, hasCloseBtn = true, onClose, children }) => {
//   const modalRef = useRef<HTMLDialogElement | null>(null);

//   useEffect(() => {
//     setModalOpen(isOpen);
//   }, [isOpen]);

//   const handleClose = () => {
//     if (onClose) {
//       onClose();
//     }
//     setModalOpen(false);
//   };

//   return (
//     <dialog ref={modalRef} open={isOpen} className="modal">
//       {children}
//     </dialog>
//   );
// };

// export default Modal;
// function setModalOpen(isOpen: boolean) {
//     throw new Error("Function not implemented.");
// }

import React, { useState, useRef, useEffect } from "react";

interface ModalProps {
  isOpen: boolean;
  hasCloseBtn?: boolean;
  onClose?: () => void;
  children: React.ReactNode; // Add this line to define the 'children' prop
}

const Modal: React.FC<ModalProps> = ({ isOpen, hasCloseBtn = true, onClose, children }) => {
  const modalRef = useRef<HTMLDialogElement | null>(null);

  useEffect(() => {
    // You can use the ref to manipulate the modal if needed
    if (modalRef.current) {
      modalRef.current.showModal();
    }
  }, [isOpen]);

  const handleClose = () => {
    if (onClose) {
      onClose();
    }
    // Closing the modal by updating the parent state
  };

  return (
    <dialog ref={modalRef} open={isOpen} className="modal">
      {hasCloseBtn}
      {children}
    </dialog>
  );
};

export default Modal;

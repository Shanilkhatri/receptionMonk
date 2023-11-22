
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
// import React, { useState } from 'react';

// const MyModal = () => {
//   const [profileData, setProfileData] = useState({
//     name: '',
//     email: '',
//     phone: '',
//     address: '',
//   });

//   const [companyDetails, setCompanyDetails] = useState({
//     companyName: '',
//     industry: '',
//     revenue: '',
//     employees: '',
//   });

//   const [kyc, setKyc] = useState({
//     adharCard: '',
//     uploadButton: '',
//   });

//   const [personal, setPersonal] = useState({
//     firstName: '',
//     lastName: '',
//     dateOfBirth: '',
//     gender: '',
//   });

//   const [currentStep, setCurrentStep] = useState(1);

//   const handleNext = () => {
//     setCurrentStep((prevStep) => prevStep + 1);
//   };

//   const renderStepContent = () => {
//     switch (currentStep) {
//       case 1:
//         return (
//           <>
//             <h2>Profile Data</h2>
//             <input
//               type="text"
//               placeholder="Name"
//               value={profileData.name}
//               onChange={(e) => setProfileData({ ...profileData, name: e.target.value })}
//             />
//             {/* Add other input fields for Profile Data */}
//             <button onClick={handleNext} disabled={!isProfileDataValid()}>
//               Next
//             </button>
//           </>
//         );
//       case 2:
//         return (
//           <>
//             <h2>Company Details</h2>
//             <input
//               type="text"
//               placeholder="Company Name"
//               value={companyDetails.companyName}
//               onChange={(e) => setCompanyDetails({ ...companyDetails, companyName: e.target.value })}
//             />
//             {/* Add other input fields for Company Details */}
//             <button onClick={handleNext} disabled={!isCompanyDetailsValid()}>
//               Next
//             </button>
//           </>
//         );
//       case 3:
//         return (
//           <>
//             <h2>KYC</h2>
//             {/* Add file input for Adhar Card */}
//             {/* <input type="file" onChange={(e) => setKyc({ ...kyc, adharCard: e.target.files[0]})} /> */}
//             <input
//               type="text"
//               placeholder="adharno "
//               value={kyc.adharCard}
//               onChange={(e) => setKyc({ ...kyc, adharCard: e.target.value })}
//             />
//             {/* Add other input fields for KYC */}
//             <button onClick={handleNext} disabled={!isKycValid()}>
//               Next
//             </button>
//           </> 
//         );
//       case 4:
//         return (
//           <>
//             <h2>Personal</h2>
//             <input
//               type="text"
//               placeholder="First Name"
//               value={personal.firstName}
//               onChange={(e) => setPersonal({ ...personal, firstName: e.target.value })}
//             />
//             {/* Add other input fields for Personal */}
//             <button onClick={handleSubmit} disabled={!isPersonalValid()}>
//               Submit
//             </button>
//           </>
//         );
//       default:
//         return null;
//     }
//   };

//   const isProfileDataValid = () => {
//     // Implement your validation logic for Profile Data
//     return profileData.name !== '' && profileData.email !== '' && profileData.phone !== '' && profileData.address !== '';
//   };

//   const isCompanyDetailsValid = () => {
//     // Implement your validation logic for Company Details
//     return (
//       companyDetails.companyName !== '' && companyDetails.industry !== '' && companyDetails.revenue !== '' && companyDetails.employees !== ''
//     );
//   };

//   const isKycValid = () => {
//     // Implement your validation logic for KYC
//     return kyc.adharCard !== "" && kyc.uploadButton !== '';
//   };

//   const isPersonalValid = () => {
//     // Implement your validation logic for Personal
//     return personal.firstName !== '' && personal.lastName !== '' && personal.dateOfBirth !== '' && personal.gender !== '';
//   };

//   const handleSubmit = () => {
//     // Implement logic to submit the data
//     // You can send the data to the server or perform any necessary action
//     // Reset state or close the modal as needed
//     console.log('Data submitted:', { profileData, companyDetails, kyc, personal });
//     // Reset state or close the modal as needed
//   };

//   return (
//     <div className="modal">
//       <div className="modal-content">
//         {renderStepContent()}
//         <button onClick={() => setCurrentStep((prevStep) => prevStep - 1)} disabled={currentStep === 1}>
//           Back
//         </button>
//       </div>
//     </div>
//   );
// };

// export default MyModal;

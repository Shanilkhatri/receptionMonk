import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { useEffect } from 'react';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const AddUser = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Add User'));
    });

const submitForm = () => {
    const toast = Swal.mixin({
        toast: true,
        position: 'top',
        showConfirmButton: false,
        timer: 3000,
    });
    toast.fire({
        icon: 'success',
        title: 'Form submitted successfully',
        padding: '10px 20px',
    });
};

const SubmittedForm = Yup.object().shape({
    userName: Yup.string().required('Please fill User Name'),
    fullName: Yup.string().required('Please fill the Name'),
    email: Yup.string().email('Invalid email').required('Please fill the Email'),
    color: Yup.string().required('Please Select the field'),
    firstname: Yup.string().required('Please fill the first name'),
    lastname: Yup.string().required('Please fill the last name'),
    username: Yup.string().required('Please choose a userName'),
    city: Yup.string().required('Please provide a valid city'),
    state: Yup.string().required('Please provide a valid state'),
    zip: Yup.string().required('Please provide a valid zip'),
    agree: Yup.string().required('You must agree before submitting.'),
    select: Yup.string().required('Please Select the field'),
});

const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;
    
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        User
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Add</span>
                </li>
            </ul>

            <div className='flex xl:flex-row flex-col gap-2.5 py-5'>
                <div className="panel px-0 flex-1 py-6 ltr:xl:mr-6 rtl:xl:ml-6">                
                    <div className="mt-8 px-4">          
                        <div className='text-xl font-bold text-dark dark:text-white text-center'>Add User</div>                      
        
                            <div className="flex justify-between lg:flex-row flex-col">
                                <div className="lg:w-1/2 w-full m-8">  
                                    <div className="flex items-center my-8">
                                        <label htmlFor="userName" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                            Name
                                        </label>
                                        <input id="user-email" type="email" name="user-email" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Email" />
                                    </div>
                                    <div className="my-8 flex items-center">
                                        <label htmlFor="user-email" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                            Email
                                        </label>
                                        <input id="user-email" type="email" name="user-email" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Email" />
                                    </div>
                                    <div className="my-8 flex items-center">
                                        <label htmlFor="pwd" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                            Password
                                        </label>
                                        <input id="pwd" type="text" name="pwd" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Password" />
                                    </div>
                                    <div className="my-8 flex items-center">
                                        <label htmlFor="phn-number" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                            Phone Number
                                        </label>
                                        <input id="phn-number" type="text" name="phn-number" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Phone number" />
                                    </div>
                                </div>
                                <div className="lg:w-1/2 w-full m-8">
                                    <div className="flex items-center my-8">
                                        <label htmlFor="dob" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                            Date of Birth
                                        </label>
                                        <input id="dob" type="date" name="dob" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Account Number" />
                                    </div>
                                    <div className="flex items-center my-8">
                                        <label htmlFor="acctype" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                            Account Type
                                        </label>
                                        <select id="acctype" name="acctype" className="form-select flex-1 border border-gray-400 focus:border-orange-400">
                                            <option value="">Select</option>
                                            <option value="United States">User</option>
                                            <option value="United Kingdom">Owner</option>                   
                                        </select>
                                    </div>
                                    <div className="flex items-center my-8">
                                        <label htmlFor="comp-name" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                            Company Name
                                        </label>
                                        <input id="comp-name" type="text" name="comp-name" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Company Name" />
                                    </div>                            
                                    <div className="flex items-center my-8">
                                        <label htmlFor="status" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                            Status
                                        </label>
                                        <select id="status" name="status" className="form-select flex-1 border border-gray-400 focus:border-orange-400">
                                            <option value="">Select</option>
                                            <option value="United States">Active</option>
                                            <option value="United Kingdom">Pending</option>
                                            <option value="Canada">Deactive</option>                                  
                                        </select>
                                    </div>
                                </div>
                            </div>
                            <div className="flex w-full pb-12">
                                <div className="flex items-center lg:justify-end lg:mr-6 w-1/2">                        
                                    <button type="reset" className="btn btn-outline-dark rounded-xl px-6 font-bold">RESET</button>
                                </div>
                                <div className="flex items-center lg:ml-6 w-1/2">
                                    <button                                 
                                        type="submit" 
                                        className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:border-black font-bold"
                                    >
                                        ADD
                                    </button>
                                </div>
                            </div>
                            
                         {/* <Formik
                                initialValues={{
                                    userName: '',
                                }}
                                validationSchema={SubmittedForm}
                                onSubmit={() => {}}
                            >
                                {({ errors, submitCount, touched }) => (
                                    <Form className="space-y-5">
                                        <div className="mt-8 px-4">
                                            <div className="flex justify-between lg:flex-row flex-col">
                                                <div className="lg:w-1/2 w-full m-8"> 
                                                    <div className={submitCount ? (errors.userName ? 'has-error' : 'has-success') : ''}>
                                                        <div className="flex items-center my-8">   
                                                            <label htmlFor="userName" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">Full Name </label>
                                                            <Field name="userName" type="text" id="userName" placeholder="Enter User Name" className="form-input border border-gray-400 focus:border-orange-400" /><br />
                                                            {submitCount ? errors.userName ? <div className="text-danger mt-1">{errors.userName}</div> : <div className="text-success mt-1">Looks Good!</div> : ''}
                                                        
                                                        </div>

                                                    </div> 
                                                    <button
                                                        type="submit"
                                                        className="btn btn-primary !mt-6"
                                                        onClick={() => {
                                                            if (touched.userName && !errors.userName) {
                                                                submitForm();
                                                            }
                                                        }}
                                                    >
                                                        Submit Form
                                                    </button>
                                                </div>
                                            </div>
                                        </div>
                                    </Form>
                                )}
                            </Formik> */}   
                    </div>
                </div>
            </div>
        </div>
    );
};

export default AddUser;
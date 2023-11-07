import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { useEffect } from 'react';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const EditUser = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Update User'));
    });

    const navigate = useNavigate();

    const submitForm = () => {
        // navigate('/');
        const toast = Swal.mixin({
            toast: true,
            position: 'top-end',
            showConfirmButton: false,
            timer: 3000,
        });
        toast.fire({
            icon: 'success',
            title: 'User Updated Successfully',
            padding: '10px 20px',
        });
    };

    const SubmittedForm = Yup.object().shape({
        userName: Yup.string().required('Please fill User Name'),
        userEmail: Yup.string().email('Invalid Email Address').required('Please fill Email'),
        userPassword: Yup.string().required('Please fill User Password'),
        userPhone: Yup.string().required('Please fill User Phone Number'),
        userDob: Yup.string().required('Please fill User DOB'),
        userAccType: Yup.string().required('Please select User Type'),
        userCompName: Yup.string().required('Please fill User Company Name'),
        userStatus: Yup.string().required('Please select User Status'),
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
                    <span>Update</span>
                </li>
            </ul>

            <Formik
                    initialValues={{
                        userName: '',
                        userEmail: '',
                        userPassword: '',
                        userPhone: '',
                        userDob: '',
                        userAccType: '',
                        userCompName: '',
                        userStatus: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched })  => (   
                        <Form className="space-y-5">
                            <div className='panel py-6 mt-6'>
                                <div className='p-4 xl:m-12 lg:m-0 md:m-8 xl:my-18 xl:my-8'>  
                                    <div className='text-xl font-bold text-dark dark:text-white text-center pb-12'>Update User</div>      
                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userName" className='py-2'>
                                                    Name <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userName ? 'has-error' : 'has-success') : ''}>
                                                <Field name="userName" type="text" id="userName" placeholder="Enter User Name" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.userName ? <div className="text-danger mt-1">{errors.userName}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="user-email" className="py-2">
                                                    Email <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userEmail ? 'has-error' : 'has-success') : ''}>
                                                <Field name="userEmail" type="text" id="userEmail" placeholder="Enter User Email" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.userEmail ? <div className="text-danger mt-1">{errors.userEmail}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>                                            
                                        </div>
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userPassword" className='py-2'>
                                                    Password <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userPassword ? 'has-error' : 'has-success') : ''}>
                                                <Field name="userPassword" type="password" id="userPassword" placeholder="Enter User Password" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.userPassword ? <div className="text-danger mt-1">{errors.userPassword}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userPhone" className="py-2">
                                                    Phone Number <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userPhone ? 'has-error' : 'has-success') : ''}>
                                                <Field name="userPhone" type="text" id="userPhone" placeholder="Enter User Phone Number" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.userPhone ? <div className="text-danger mt-1">{errors.userPhone}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>                                            
                                        </div>                                     
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userDob" className='py-2'>
                                                    Date of Birth <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userDob ? 'has-error' : 'has-success') : ''}>
                                                <Field name="userDob" type="date" id="userDob" className="form-input border border-gray-400 focus:border-orange-400 text-gray-600" placeholder="Enter User Birth Date"  />

                                                {submitCount ? errors.userDob ? <div className="text-danger mt-1">{errors.userDob}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userAccType" className="py-2">
                                                    Account Type <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userAccType ? 'has-error' : 'has-success') : ''}>                                            

                                                <Field as="select" name="userAccType" className="form-select border border-gray-400 focus:border-orange-400 text-gray-600">
                                                    <option value="">Select</option>
                                                    <option value="User">User</option>
                                                    <option value="Owner">Owner</option> 
                                                </Field>
                                                {submitCount ? (
                                                    errors.userAccType ? (
                                                        <div className=" text-danger mt-1">{errors.userAccType}</div>
                                                    ) : (
                                                        <div className="text-success mt-1">It is fine!</div>
                                                    )
                                                ) : (
                                                    ''
                                                )}
                                               
                                            </div>                                            
                                        </div>
                                    </div>

                                    <div className='grid lg:grid-cols-2 lg:space-x-12 lg:space-y-0 lg:my-5'>
                                        <div className='grid md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userCompName" className='py-2'>
                                                    Company Name <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userCompName ? 'has-error' : 'has-success') : ''}>
                                                <Field name="userCompName" type="text" id="userCompName" placeholder="Enter Company Name" className="form-input border border-gray-400 focus:border-orange-400" />

                                                {submitCount ? errors.userCompName ? <div className="text-danger mt-1">{errors.userCompName}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                                            </div>
                                        </div>
                                        
                                        <div className='grid  md:grid-cols-2 my-3 lg:my-0'>
                                            <div className=''>
                                                <label htmlFor="userStatus" className="py-2">
                                                    Status <span className='text-red-600'>*</span>
                                                </label>
                                            </div>
                                            <div className={submitCount ? (errors.userStatus ? 'has-error' : 'has-success') : ''}>                                            

                                                <Field as="select" name="userStatus" className="form-select border border-gray-400 focus:border-orange-400 text-gray-600">
                                                    <option value="">Select</option>
                                                    <option value="United States">Active</option>
                                                    <option value="United Kingdom">Pending</option>
                                                    <option value="Canada">Deactive</option>           
                                                </Field>
                                                {submitCount ? (
                                                    errors.userStatus ? (
                                                        <div className=" text-danger mt-1">{errors.userStatus}</div>
                                                    ) : (
                                                        <div className="text-success mt-1">It is fine!</div>
                                                    )
                                                ) : (
                                                    ''
                                                )}
                                               
                                            </div>                                            
                                        </div>
                                    </div>


                                    <div className="flex justify-center py-6 mt-12">
                                        <button 
                                            type="reset" 
                                            className="btn btn-outline-dark rounded-xl px-8 mx-3 font-bold">
                                            RESET
                                        </button>

                                        <button
                                            type="submit"
                                            className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828] mx-3"
                                            onClick={() => {
                                                if (touched.userName && !errors.userName) {
                                                    submitForm();
                                                }
                                                else if (touched.userEmail && !errors.userEmail) {
                                                    submitForm();
                                                }
                                                else if (touched.userPassword && !errors.userPassword) {
                                                    submitForm();
                                                }
                                                else if (touched.userPhone && !errors.userPhone) {
                                                    submitForm();
                                                }
                                                else if (touched.userDob && !errors.userDob) {
                                                    submitForm();
                                                }
                                                else if (touched.userAccType && !errors.userAccType) {
                                                    submitForm();
                                                } 
                                                else if (touched.userCompName && !errors.userCompName) {
                                                    submitForm();
                                                }
                                                else if (touched.userStatus && !errors.userStatus) {
                                                    submitForm();
                                                }
                                            }}
                                        >
                                                UPDATE
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </Form>
                     )}
            </Formik> 
        </div>
    );
};

export default EditUser;
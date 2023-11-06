import { useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const SignIn = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Sign In'));
    });
    const navigate = useNavigate();
    
    const submitForm = () => {
        navigate('/auth/SigninOTP');
        const toast = Swal.mixin({
            toast: true,
            position: 'top',
            showConfirmButton: false,
            timer: 3000,
        });
        toast.fire({
            icon: 'success',
            title: 'OTP has been sent to registered mobile number successfully',
            padding: '10px 20px',
        });
    };

    const SubmittedForm = Yup.object().shape({
        authPhoneNo: Yup.string().required('It must be 10-digit Number'),
    });

    return (
        <div className="flex justify-center items-center min-h-screen bg-cover bg-center bg-[url('/assets/images/bg/bgcurve.svg')] dark:bg-[url('/assets/images/bg/bgcurve.svg')]">
            <div className="panel sm:w-[480px] m-6 max-w-lg w-full shadow-md">
                <div className='text-center'>
                    <div className="flex justify-center">
                        <div><img className="h-20" src="/assets/images/logo/rm.png" alt="logo" /></div>
                    </div>
                    <div className="text-center pb-8">
                        <h1 className="font-bold text-2xl text-black pb-12">Reception <span className="text-orange-700">Monk</span></h1>
                        <h2 className="font-semibold text-xl mb-3 text-neutral-800">Sign In</h2>
                    </div>
                </div>

                <Formik
                    initialValues={{
                        authPhoneNo: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched }) => (
                        <Form className="space-y-5">
                            <p className="mb-7">Enter your phone number to complete Registration</p>
                            
                            <div className={submitCount ? (errors.authPhoneNo ? 'has-error' : 'has-success') : ''}>
                                <label htmlFor="authPhoneNo">Phone No. <span className='text-red-600'>*</span></label>
                                <Field name="authPhoneNo" type="text" id="authPhoneNo" placeholder="Enter Phone Number" className="form-input border border-gray-400 focus-border-orange-400" />
                                
                                {submitCount ? errors.authPhoneNo ? <div className="text-danger mt-1">{errors.authPhoneNo}</div> : <div className="text-success mt-1">It is fine now!</div> : ''}
                            </div>

                            <div className="flex justify-center py-6">
                                <button
                                    type="submit"
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover-bg-[#282828]"
                                    onClick=
                                    {() => {
                                        if (touched.authPhoneNo && !errors.authPhoneNo) {
                                            submitForm();
                                        }
                                    }}
                                >
                                    REGISTER
                                </button>
                            </div>
                        </Form>
                    )}
                </Formik>
            </div>
        </div>
    );
};

export default SignIn;
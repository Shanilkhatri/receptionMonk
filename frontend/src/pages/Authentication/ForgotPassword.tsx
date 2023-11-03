import { Link, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const ForgotPassword = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Forgot Password'));
    });
    const navigate = useNavigate();
   
    const submitForm = () => {
        // navigate('/');
        const toast = Swal.mixin({
            toast: true,
            position: 'top',
            showConfirmButton: false,
            timer: 3000,
        });
        toast.fire({
            icon: 'success',
            title: 'Please Check your Email ID',
            padding: '10px 20px',
        });
    };

    const SubmittedForm = Yup.object().shape({
        authForgotPassword: Yup.string().email('Invalid Email Address').required('Please fill email address to recover your account'),
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
                        <h2 className="font-semibold text-xl mb-3 text-neutral-800">Forgot Password</h2>
                    </div>
                </div>

                <Formik
                    initialValues={{
                        authForgotPassword: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched })  => (   
                        <Form className="space-y-5">
                            
                            <div className={submitCount ? (errors.authForgotPassword ? 'has-error' : 'has-success') : ''}>
                                <label htmlFor="authForgotPassword">Email <span className='text-red-600'>*</span></label>
                                <Field name="authForgotPassword" type="email" id="authForgotPassword" placeholder="Enter Email Address" className="form-input border border-gray-400 focus:border-orange-400" />

                                {submitCount ? errors.authForgotPassword ? <div className="text-danger mt-1">{errors.authForgotPassword}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                            </div>

                            <div className="flex justify-center pt-6">
                                <button
                                    type="submit"
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]"
                                    onClick={() => {
                                        if (touched.authForgotPassword && !errors.authForgotPassword) {
                                            submitForm();
                                        }                                        
                                    }}
                                >
                                        RECOVER
                                </button>
                            </div>

                            <p className="text-center font-semibold pb-2">
                                Remember account ?
                                <Link to="/auth/Login" className="font-bold text-orange-700 hover:underline ltr:ml-1 rtl:mr-1 hover:text-gray-900">
                                    Sign In
                                </Link>                     
                            </p> 
                        </Form>
                    )}
                </Formik>   

            </div>
        </div>
    );
};

export default ForgotPassword;

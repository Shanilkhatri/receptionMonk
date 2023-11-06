import { Link, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';


const SignInOTP = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('SignIn OTP Verification'));
    });
    const navigate = useNavigate();
   
    const submitForm = () => {
        navigate('/');
        const toast = Swal.mixin({
            toast: true,
            position: 'top',
            showConfirmButton: false,
            timer: 3000,
        });
        toast.fire({
            icon: 'success',
            title: 'OTP has been verified successfully',
            padding: '10px 20px',
        });
    };

    const SubmittedForm = Yup.object().shape({
        authSignInOTP: Yup.string().required('Please fill all fields'),
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
                        <h2 className="font-semibold text-xl mb-3 text-neutral-800">OTP</h2>
                    </div>
                </div> 

                 <Formik
                    initialValues={{
                        authSignInOTP: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched })  => (   
                        <Form className="space-y-5">
                            <p className="mb-7 text-center">Enter 6-digit OTP to complete Registration</p>
                            <div className={submitCount ? (errors.authSignInOTP ? 'has-error' : 'has-success') : ''}>
                                <div className="grid grid-cols-6 gap-4 m-8">
                                    <div>
                                        <Field name="authSignInOTP1" type="text" className="form-input border border-gray-400 focus:border-orange-400 text-center" tabIndex={1} />
                                    </div>
                                    <div>
                                        <Field name="authSignInOTP2" type="text" className="form-input border border-gray-400 focus:border-orange-400 text-center" tabIndex={2} />
                                    </div>
                                    <div>
                                        <Field name="authSignInOTP3" type="text" className="form-input border border-gray-400 focus:border-orange-400 text-center" tabIndex={3} />
                                    </div>
                                    <div>
                                        <Field name="authSignInOTP4" type="text" className="form-input border border-gray-400 focus:border-orange-400 text-center" tabIndex={4} />
                                    </div>
                                    <div>
                                        <Field name="authSignInOTP5" type="text" className="form-input border border-gray-400 focus:border-orange-400 text-center" tabIndex={5} />
                                    </div>
                                    <div>
                                        <Field name="authSignInOTP6" type="text" className="form-input border border-gray-400 focus:border-orange-400 text-center" tabIndex={6} />
                                    </div>
                                </div>
                                {/* {submitCount ? errors.authSignInOTP ? <div className="text-danger mx-8">{errors.authSignInOTP}</div> : <div className="text-success mt-1">It is fine!</div> : ''} */}
                            </div>

                            <div className="flex justify-center py-6">
                                <button
                                    type="submit"
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]"
                                    onClick={() => {
                                        if (touched.authSignInOTP && !errors.authSignInOTP) {
                                            submitForm();
                                        }                                        
                                    }}
                                >
                                        SUBMIT
                                </button>
                            </div>
                        </Form>
                    )}
                </Formik>      
            </div>
        </div>
    );
};

export default SignInOTP;

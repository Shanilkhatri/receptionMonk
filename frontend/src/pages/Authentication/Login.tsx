import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const Login = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Login'));
    });
    const navigate = useNavigate();
    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;

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
            title: 'Login successfully',
            padding: '10px 20px',
        });
    };

    const SubmittedForm = Yup.object().shape({
        authEmail: Yup.string().email('Invalid Email Address').required('Please fill email address'),
        authPassword: Yup.string().required('Please fill password for Email ID'),
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
                        authEmail: '',
                        authPassword: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched })  => (   
                        <Form className="space-y-5">
                            <div className={submitCount ? (errors.authEmail ? 'has-error' : 'has-success') : ''}>
                                <label htmlFor="authEmail">Email <span className='text-red-600'>*</span></label>
                                <Field name="authEmail" type="email" id="authEmail" placeholder="Enter Email Address" className="form-input border border-gray-400 focus:border-orange-400" />

                                {submitCount ? errors.authEmail ? <div className="text-danger mt-1">{errors.authEmail}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                            </div>

                            <div className={submitCount ? (errors.authPassword ? 'has-error' : 'has-success') : ''}>
                                <label htmlFor="authPassword">Password <span className='text-red-600'>*</span></label>
                                <Field name="authPassword" type="Password" id="authPassword" placeholder="Enter Password" className="form-input border border-gray-400 focus:border-orange-400" />

                                {submitCount ? errors.authPassword ? <div className="text-danger mt-1">{errors.authPassword}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                            </div>

                            <div className="flex justify-center pt-6">
                                <button
                                    type="submit"
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]"
                                    onClick={() => {
                                        if (touched.authEmail && !errors.authEmail) {
                                            submitForm();
                                        }
                                        else if (touched.authPassword && !errors.authPassword) {
                                            submitForm();
                                        }
                                    }}
                                >
                                        SIGN IN 
                                </button>
                            </div>

                            <div className="relative my-7 h-5 text-center before:w-full before:h-[1px] before:absolute before:inset-0 before:m-auto before:bg-[#ebedf2] dark:before:bg-[#253b5c]">
                                <div className="font-bold text-white-dark bg-white dark:bg-black px-2 relative z-[1] inline-block">
                                    <span>OR</span>
                                </div>
                            </div>

                            <p className="text-center py-2">                    
                                <Link to="/auth/ForgotPassword" className="font-semibold hover:underline hover:text-orange-700 ltr:ml-1 rtl:mr-1">
                                    Forgot Password
                                </Link>                    
                            </p>         
                         
                        </Form>
                    )}
                </Formik>               
            </div>
        </div>
    );
};

export default Login;

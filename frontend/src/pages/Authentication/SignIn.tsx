import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
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
    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;

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
            title: 'Form submitted successfully',
            padding: '10px 20px',
        });
    };

    const SubmittedForm = Yup.object().shape({
        phoneNo: Yup.string().required('It must be a 10-digit Number'),
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
                        phoneNo: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched }) => (
                        <Form className="space-y-5" onSubmit={submitForm}>
                            <p className="mb-7">Enter your phone number to complete Registration</p>
                            
                            <div className={submitCount ? (errors.phoneNo ? 'has-error' : 'has-success') : ''}>
                                <label htmlFor="phoneNo">Phone No. <span className='text-red-600'>*</span></label>
                                <Field name="phoneNo" type="text" id="phoneNo" placeholder="Enter User Name" className="form-input border border-gray-400 focus-border-orange-400" /><br />
                                {submitCount ? errors.phoneNo ? <div className="text-danger mt-1">{errors.phoneNo}</div> : <div className="text-success mt-1">Looks Okay!</div> : ''}
                            </div>

                            <div className="flex justify-center py-6">
                                <button
                                    type="submit"
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover-bg-[#282828]"
                                    onClick=
                                    {() => {
                                        if (touched.phoneNo && !errors.phoneNo) {
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
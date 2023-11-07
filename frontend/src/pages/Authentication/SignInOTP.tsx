import { Link, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

interface FormValues {
    authSignInOTP1: string
    authSignInOTP2: string
    authSignInOTP3: string
    authSignInOTP4: string
    authSignInOTP5: string
    authSignInOTP6: string

    // Define other fields if needed
}
const SignInOTP = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('SignIn OTP Verification'));
    });
    const navigate = useNavigate();


    const submitForm = (values: FormValues, { setSubmitting }: any) => {
        console.log("values: ", values)
        // navigate('/');
        // const toast = Swal.mixin({
        //     toast: true,
        //     position: 'top',
        //     showConfirmButton: false,
        //     timer: 3000,
        // });
        // toast.fire({
        //     icon: 'success',
        //     title: 'OTP has been verified successfully',
        //     padding: '10px 20px',
        // });
        setSubmitting(false)
    };

    
    function hasAnyError(errors:any) {
        for (let i = 1; i <= 6; i++) {
            if (errors[`authSignInOTP${i}`]) {
                document.getElementById(`authSignInOTP${i}`)?.classList.remove("border-gray-400")
                document.getElementById(`authSignInOTP${i}`)?.classList.add("border-red-400")
                return true; // has err
            }else{
                document.getElementById(`authSignInOTP${i}`)?.classList.remove("border-red-400")
                document.getElementById(`authSignInOTP${i}`)?.classList.add("border-green-400")
            }
        }
        return false; //no err
    }
    const validationSchema = Yup.object().shape({
        authSignInOTP1: Yup.string().matches(/^\d$/, 'Enter a digit').required('Required'),
        authSignInOTP2: Yup.string().matches(/^\d$/, 'Enter a digit').required('Required'),
        authSignInOTP3: Yup.string().matches(/^\d$/, 'Enter a digit').required('Required'),
        authSignInOTP4: Yup.string().matches(/^\d$/, 'Enter a digit').required('Required'),
        authSignInOTP5: Yup.string().matches(/^\d$/, 'Enter a digit').required('Required'),
        authSignInOTP6: Yup.string().matches(/^\d$/, 'Enter a digit').required('Required'),
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
                        authSignInOTP1: '',
                        authSignInOTP2: '',
                        authSignInOTP3: '',
                        authSignInOTP4: '',
                        authSignInOTP5: '',
                        authSignInOTP6: '',
                    }}
                    validationSchema={validationSchema}
                    onSubmit={(values, actions) => {
                        submitForm(values, actions);
                    }}
                >
                    {({ errors, submitCount, touched }) => (
                        <Form className="space-y-5">
                            <p className="mb-7 text-center">Enter 6-digit OTP to complete Registration</p>
                            <div className={submitCount && (errors.authSignInOTP1 || errors.authSignInOTP2 || errors.authSignInOTP3 || errors.authSignInOTP4 || errors.authSignInOTP5 || errors.authSignInOTP6) ? 'has-error' : ''}>
                                <div className="grid grid-cols-6 gap-4 m-8">
                                    {Array.from({ length: 6 }, (_, i) => (
                                        <div key={i}>
                                            <Field
                                                id = {`authSignInOTP${i + 1}`}
                                                name={`authSignInOTP${i + 1}`}
                                                maxLength="1"
                                                type="text"
                                                className="form-input border border-gray-400 focus:border-orange-400 text-center"
                                                tabIndex={i + 1}
                                            // initialValue={undefined}
                                            />

                                        </div>

                                    ))}
                                </div>
                                {submitCount >= 0 && hasAnyError(errors) && (
                                    <div className="text-danger mx-8">
                                        Please fill all the fields correctly with numbers.
                                    </div>
                                )}
                                {/* {submitCount >= 0 && !hasAnyError(errors) && } */}
                                

                            </div>
                            <div className="flex justify-center py-6">
                                <button
                                    type="submit"
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]"

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

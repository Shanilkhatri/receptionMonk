import { useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setEmailVerToken, setPageTitle } from '../../store/themeConfigSlice';
import { setEmail } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';
// Define the type for form values
interface FormValues {
    authEmailId: string;
    // Define other fields if needed
}
const appUrl = import.meta.env.VITE_APPURL
const SignIn = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Sign In'));
    });
    const navigate = useNavigate();

    const submitForm = async (values: FormValues, { setSubmitting }: any) => {
        const authEmailId = values.authEmailId
        console.log(JSON.stringify(values)) // authEmailId: "hi@gmail"
        // console.log(authEmailId) // "hi@gmail"

        // use fetch to get to route
        try {
            // Send the form data to a server endpoint for validation and processing
            const response = await fetch(appUrl+'loginbyemail', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(values),
                
            });
            const responseData = await response.json(); // Parse the response JSON
            console.log("responseData.payload",responseData.Payload)
            if (response.ok) {
                // dispatch now to set email in state
                dispatch(setEmail(authEmailId));
                // dispatch now to set emailVerfToken in state
                dispatch(setEmailVerToken(responseData.Payload));

                // Successful response
                navigate('/auth/SigninOTP');
                const toast = Swal.mixin({
                    toast: true,
                    position: 'top',
                    showConfirmButton: false,
                    timer: 3000,
                });
                toast.fire({
                    icon: 'success',
                    title: 'OTP has been sent to registered E-mail successfully',
                    padding: '10px 20px',
                });

                // setting submit button disabled to false
                setSubmitting(false)
                
                }else{
                const toast = Swal.mixin({
                    toast: true,
                    position: 'top',
                    showConfirmButton: false,
                    timer: 3000,
                });
                toast.fire({
                    icon: 'error', 
                    title: 'Sending OTP failed. Please try again.',
                    padding: '10px 20px',
                });
            }
        } catch (error) {
            console.error('Error:', error);
        }
        // setting submit button disabled to false
        setSubmitting(false)
    };

    const SubmittedForm = Yup.object().shape({
        authEmailId: Yup.string().email('Invalid Email Address').required('Please enter a valid E-mail'),
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
                        <h2 className="font-semibold text-xl mb-3 text-neutral-800">Login / SignUp</h2>
                    </div>
                </div>

                <Formik
                    initialValues={{
                        authEmailId: '',
                    }}
                    validationSchema={SubmittedForm}
                    validateOnChange={true}
                    validateOnBlur={true}
                    onSubmit={(values, actions) => {
                        submitForm(values, actions);
                    }}
                >
                    {({ errors, submitCount, touched }) => (
                        <Form className="space-y-5">
                            {/* <p className="mb-7">Enter your E-mail to Login / SignUp</p> */}

                            <div className={submitCount ? (errors.authEmailId ? 'has-error' : 'has-success') : ''}>
                                <label htmlFor="authEmailId">E-mail <span className='text-red-600'>*</span></label>
                                <Field name="authEmailId" type="text" id="authEmailId" placeholder="Enter E-mail" className="form-input border border-gray-400 focus-border-orange-400" />

                                {submitCount ? errors.authEmailId ? <div className="text-danger mt-1">{errors.authEmailId}</div> : <div className="text-success mt-1">It is fine now!</div> : ''}
                            </div>

                            <div className="flex justify-center py-6">
                                <button
                                    type="submit"
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover-bg-[#282828]"
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

export default SignIn;
import { Link, useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useEffect, useState } from 'react';
import { setPageTitle,setEmailVerToken } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';
import store from '../../store';
interface FormValues {
    authSignInOTP1: string
    authSignInOTP2: string
    authSignInOTP3: string
    authSignInOTP4: string
    authSignInOTP5: string
    authSignInOTP6: string

    // Define other fields if needed
}
const appUrl = import.meta.env.VITE_APPURL
const SignInOTP = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    // otp timer code ----------
    const [timer, setTimer] = useState(10); // Initial timer value in seconds
    const [isTimerVisible, setIsTimerVisible] = useState(true);

    const startTimer = () => {
        setTimer(10);
        setIsTimerVisible(true);
    };

    const handleResendClick = async () => {
        // resend OTP 
        // jsonObj to send email 
        var jsonObj = {
            "authEmailId": store.getState().themeConfig.email  // email at redux-store
        }
        // writing a request to hit the send email again endpoint
        let response = await fetch(appUrl+"loginbyemail",{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                // adding header for email-verf-token
                'emailVerfToken': store.getState().themeConfig.emailVerfToken
            },
            body: JSON.stringify(jsonObj),
        })
        // response > json
        let data = await response.json()
        
        if (response.ok){
            // For demonstration purposes, let's just reset the timer
            startTimer();
            dispatch(setEmailVerToken(data.Payload));
            return
        }
        const toast = Swal.mixin({
            toast: true,
            position: 'top',
            showConfirmButton: false,
            timer: 3000,
        });
        toast.fire({
            icon: 'error', 
            title: 'Maximum resend attempts reached. Re-directing!',
            padding: '10px 20px',
        });
        setTimer(3)
        navigate("/auth/signin")
        //add a exception for 10 times resending the otp.
    };
    // otp timer code finish----------

    useEffect(() => {
        dispatch(setPageTitle('SignIn OTP Verification'));
        // if state doesn't have email we throw user back to login
        if (store.getState().themeConfig.email == "") {
            navigate("/auth/SignIn")
        }
        // otp timer code -----------
        let interval: any;

        if (timer > 0) {
            interval = setInterval(() => {
                setTimer((prevTimer) => prevTimer - 1);
            }, 1000);
        } else {
            setIsTimerVisible(false);
        }

        return () => clearInterval(interval);
        // otp timer code finish----------
    }, [timer]);



    // submitting otp
    const submitForm = async (values: FormValues, { setSubmitting }: any) => {
        var otpToSend = ""
        for (var keys in values) {
            if (values.hasOwnProperty(keys)) {
                otpToSend += values[keys as keyof FormValues]
            }
        }


        var jsonObj = {
            "otp": otpToSend,
            "authEmailId": store.getState().themeConfig.email  // email at redux-store
        }

        const response = await fetch(appUrl + "matchotp", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                // adding header for email-verf-token
                'emailVerfToken': store.getState().themeConfig.emailVerfToken
            },
            body: JSON.stringify(jsonObj),

        });
        var responseData = await response.json() // wait for response > json

        if (response.ok) {
            // if (true) {

            // otp is validated successfully
            // next we'll get a token from server which 
            // will be stored in cookies pointing to other info about the user

            // Setting it into cookies with an expiry time of 6 months (in seconds)
            var expirationDate = new Date();
            expirationDate.setMonth(expirationDate.getMonth() + 6);

            // setting token into cookies with expiry time 6months or 12
            // -> right now we won't use "secure" & "httpOnly" flags as we want to read cookies

            // for production:
            // -> document.cookie = "myCookie=myValue; secure; HttpOnly; path=/; SameSite=Strict";

            // for development :
            document.cookie = `exampleToken=${responseData.Payload}; secure;  expires=${expirationDate.toUTCString()}; path=/`;
            // document.cookie = `exampleData=${dataString}; expires=${expirationDate.toUTCString()}; path=/`;

        }

        navigate("/");
        setSubmitting(false)
    };


    function hasAnyError(errors: any) {
        for (let i = 1; i <= 6; i++) {
            if (errors[`authSignInOTP${i}`]) {
                document.getElementById(`authSignInOTP${i}`)?.classList.remove("border-gray-400")
                document.getElementById(`authSignInOTP${i}`)?.classList.add("border-red-400")
                return true; // has err
            } else {
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
                                                id={`authSignInOTP${i + 1}`}
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
                                    <div>
                                    {isTimerVisible ? (
                                        <p>Resend OTP in {timer} seconds</p>
                                    ) : (
                                        <button onClick={handleResendClick}>Resend OTP</button>
                                    )}
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

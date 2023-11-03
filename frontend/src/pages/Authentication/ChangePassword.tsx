import { useNavigate } from 'react-router-dom';
import { useDispatch} from 'react-redux';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';
import * as Yup from 'yup';
import { Field, Form, Formik } from 'formik';
import Swal from 'sweetalert2';

const ChangePassword = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Change Password'));
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
            title: 'Password has been changed successfully',
            padding: '10px 20px',
        });
    };

    const SubmittedForm = Yup.object().shape({
        authOldPassword: Yup.string().required('Please fill last used password'),
        authNewPassword: Yup.string().required('Please fill new password'),
    });


    return (
        <div className="flex justify-center items-center min-h-screen bg-cover bg-center bg-[url('/assets/images/bg/bgcurve.svg')] dark:bg-[url('/assets/images/bg/bgcurve.svg.jpg')]">
            <div className="panel sm:w-[480px] m-6 max-w-lg w-full shadow-md">
                <div className='text-center'>
                    <div className="flex justify-center">
                        <div><img className="h-20" src="/assets/images/logo/rm.png" alt="logo" /></div>                   
                    </div>
                    <div className="text-center pb-8">
                        <h1 className="font-bold text-2xl text-black pb-12">Reception <span className="text-orange-700">Monk</span></h1>
                        <h2 className="font-semibold text-xl mb-3 text-neutral-800">Change Password</h2>
                    </div>
                </div>

                <Formik
                    initialValues={{
                        authOldPassword: '',
                        authNewPassword: '',
                    }}
                    validationSchema={SubmittedForm}
                    onSubmit={() => {}}
                >
                    {({ errors, submitCount, touched })  => (   
                        <Form className="space-y-5">                           
                            
                            <div className={submitCount ? (errors.authOldPassword ? 'has-error' : 'has-success') : ''}>
                                <label htmlFor="authOldPassword">Old Password <span className='text-red-600'>*</span></label>
                                <Field name="authOldPassword" type="text" id="authOldPassword" placeholder="Enter Old Password" className="form-input border border-gray-400 focus:border-orange-400" />

                                {submitCount ? errors.authOldPassword ? <div className="text-danger mt-1">{errors.authOldPassword}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                            </div>

                            <div className={submitCount ? (errors.authNewPassword ? 'has-error' : 'has-success') : ''}>
                                <label htmlFor="authNewPassword">New Password <span className='text-red-600'>*</span></label>
                                <Field name="authNewPassword" type="password" id="authNewPassword" placeholder="Enter New Password" className="form-input border border-gray-400 focus:border-orange-400" />

                                {submitCount ? errors.authNewPassword ? <div className="text-danger mt-1">{errors.authNewPassword}</div> : <div className="text-success mt-1">It is fine!</div> : ''}
                            </div>

                            <div className="flex justify-center pt-6">
                                <button
                                    type="submit"
                                    className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]"
                                    onClick={() => {
                                        if (touched.authOldPassword && !errors.authOldPassword) {
                                            submitForm();
                                        }    
                                        else if (touched.authNewPassword && !errors.authNewPassword) {
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

export default ChangePassword;

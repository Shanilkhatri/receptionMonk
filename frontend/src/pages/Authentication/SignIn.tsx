import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';

const SignIn = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Sign In'));
    });
    const navigate = useNavigate();
    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;

    const submitForm = () => {
        navigate('/');
    };

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
                <form className="space-y-5" onSubmit={submitForm}>
                    <p className="mb-7">Enter your phone number to complete Registration</p>
                    <div>
                        <label htmlFor="changepwd">Phone No. <span className='text-red-600'>*</span></label>
                        <input id="phoneno" type="text" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter Phone Number" />
                    </div>

                    <div className="flex justify-center py-6">
                        <Link to="/auth/SignInOTP">
                            <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]">
                                REGISTER
                            </button>
                        </Link>
                    </div>        

                    {/* <p className="text-center font-semibold pb-2">
                        Don't have an account ?
                        <Link to="/auth/SignUp" className="font-bold text-orange-700 hover:underline ltr:ml-1 rtl:mr-1 hover:text-gray-900">
                            Sign Up
                        </Link>                     
                    </p>*/}
                </form>                
            </div>
        </div>
    );
};

export default SignIn;

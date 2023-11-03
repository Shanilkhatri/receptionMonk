import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';

const SignInOTP = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Sign IN OTP'));
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
                    <h2 className="font-semibold text-xl mb-3 text-neutral-800">OTP</h2>
                 </div>
                </div>  
                <form className="space-y-5" onSubmit={submitForm}>
                    <p className="mb-7 text-center">Enter 6-digit OTP to complete Registration</p>
                    <div className="grid grid-cols-6 gap-4 m-8">
                        <div>
                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400" tabIndex={1} />
                        </div>
                        <div>
                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400" tabIndex={2}/>
                        </div>
                        <div>
                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400" tabIndex={3}/>
                        </div>
                        <div>
                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400" tabIndex={4}/>
                        </div>
                        <div>
                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400" tabIndex={5}/>
                        </div>
                        <div>
                            <input type="type" className="form-input border border-gray-400 focus:border-orange-400" tabIndex={6}/>
                        </div>
                    </div>

                    <div className="flex justify-center py-6">
                        <Link to="/">
                            <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]">
                                SUBMIT
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

export default SignInOTP;

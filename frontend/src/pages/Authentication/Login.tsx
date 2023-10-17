import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';

const Login = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Login Boxed'));
    });
    const navigate = useNavigate();
    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;

    const submitForm = () => {
        navigate('/');
    };

    return (
        <div className="flex justify-center items-center min-h-screen bg-cover bg-center bg-[url('/assets/images/bg/IVR-bg.svg')] dark:bg-[url('/assets/images/bg/IVR-bg.svg')]">
            <div className="panel sm:w-[480px] m-6 max-w-lg w-full">
                <div className='text-center'>
                    <div className="flex justify-center">
                        <div><img className="h-20" src="/assets/images/logo/rm.png" alt="logo" /></div>                   
                    </div>
                    <div className="text-center pb-8">
                    <h1 className="font-bold text-2xl text-black pb-12">Reception <span className="text-orange-600">Monk</span></h1>
                    <h2 className="font-semibold text-xl mb-3 text-neutral-800">Sign In</h2>
                 </div>
                </div>                 
              
                <form className="space-y-5" onSubmit={submitForm}>
                    <div>
                        <label htmlFor="email">Email</label> 
                        <input id="email" type="email" className="form-input" placeholder="Enter Email" />
                    </div>
                    <div>
                        <label htmlFor="password">Password</label>
                        <input id="password" type="password" className="form-input" placeholder="Enter Password" />
                    </div>
                    
                    <div className="flex justify-center pt-6">
                        <button type="submit" className="btn btn-primary rounded-full px-8 hover:btn-dark">
                            SIGN IN
                        </button>
                    </div>
                </form>
                <div className="relative my-7 h-5 text-center before:w-full before:h-[1px] before:absolute before:inset-0 before:m-auto before:bg-[#ebedf2] dark:before:bg-[#253b5c]">
                    <div className="font-bold text-white-dark bg-white dark:bg-black px-2 relative z-[1] inline-block">
                        <span>OR</span>
                    </div>
                </div>

                <p className="text-center py-2">                    
                    <Link to="/auth/ForgotPassword" className="font-semibold hover:underline hover:text-primary ltr:ml-1 rtl:mr-1">
                        Forgot Password
                    </Link>                    
                </p>
                
                <p className="text-center font-semibold">
                    Dont&apos;t have an account ?
                    <Link to="/" className="font-bold text-primary hover:underline ltr:ml-1 rtl:mr-1">
                        Sign Up
                    </Link>                    
                </p>
            </div>
        </div>
    );
};

export default Login;

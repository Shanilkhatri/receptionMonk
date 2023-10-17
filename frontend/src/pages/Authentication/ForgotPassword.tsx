import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';

const ForgotPassword = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Recover Id Box'));
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
                    <h2 className="font-semibold text-xl mb-3 text-neutral-800">Forgot Password</h2>
                 </div>
                </div>  
                <form className="space-y-5" onSubmit={submitForm}>
                    <p className="mb-7">Enter your email to recover your Password</p>
                    <div>
                        <label htmlFor="changepwd">Password</label>
                        <input id="changepwd" type="password" className="form-input" placeholder="Enter Email" />
                    </div>

                    <div className="flex justify-center py-6">
                        <button type="submit" className="btn btn-primary rounded-full px-8 hover:btn-dark">
                            Recover
                        </button>
                    </div>                    
                </form>                
            </div>
        </div>
    );
};

export default ForgotPassword;

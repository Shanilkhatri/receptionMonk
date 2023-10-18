import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';

const ChangePassword = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Change Password'));
    });
    const navigate = useNavigate();
    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;

    const submitForm = () => {
        navigate('/');
    };

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
                <form className="space-y-5" onSubmit={submitForm}>
                    <p className="mb-7">Enter your new password to change password</p>
                    <div>
                        <label htmlFor="changepwd">Password</label>
                        <input id="changepwd" type="password" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter New Password" />
                    </div>

                    <div className="flex justify-center py-6">
                        <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]">
                            SUBMIT
                        </button>
                    </div>                    
                </form>                
            </div>
        </div>
    );
};

export default ChangePassword;

import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';

const Setup2FA = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('2FA Setup'));
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
                    <h2 className="font-semibold text-xl mb-3 text-neutral-800">Setup Two Factor Authentication</h2>
                    </div>
                </div>  
                    <p className="m-2 text-base"> Two Factor Authentication provides an extra protection for your account by requiring a special code. Protect your account in just two steps:</p>
                <div className="mx-6 my-3  text-base">
                    <ol className='list-decimal'>
                        <li className='pb-2'>Link a supported authentication app (such as Authenticator , Google Authenticator etc.)</li>
                        <li className='pb-2'>Enter the Confirmation Code</li>
                    </ol>
                </div>
                <p className="m-2 text-base"> Your 2FA verified successfully</p>

                <p className="m-2 text-sm"> <span className='font-semibold'>Note:</span> You are only activating Two Factor Authentication for Owner Account only.</p>

                <div className="flex justify-center py-6">
                    <Link to="/auth/QR2FA">
                        <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]">
                            NEXT
                        </button>
                    </Link>
                </div>                           
            </div>
        </div>
    );
};

export default Setup2FA;

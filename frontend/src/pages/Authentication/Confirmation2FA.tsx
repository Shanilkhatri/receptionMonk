import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import withReactContent from 'sweetalert2-react-content';
import { setPageTitle } from '../../store/themeConfigSlice';

const Confirmation2FA = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('2FA Confirmation'));
    });
    const navigate = useNavigate();
    const isDark = useSelector((state: IRootState) => state.themeConfig.theme) === 'dark' ? true : false;

    const submitForm = () => {
        navigate('/');
    };

    const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;

    const coloredToast = (color: any) => {
        const toast = Swal.mixin({
            toast: true,
            position: isRtl ? 'top-start' : 'top-end',
            showConfirmButton: false,
            timer: 7000,
            showCloseButton: true,
            customClass: {
                popup: `color-${color}`,
            },
        });
        toast.fire({
            title: 'Now whenever you sign in, we will ask you for a code after you enter your email address and password',
        });
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
                    <h2 className="font-semibold text-xl mb-3 text-neutral-800">Confirmation Two Factor Authentication</h2>
                    </div>
                </div>  
                    <p className="m-2 text-base">Enter 6-digit code generated by your authenticator app to activate your 2FA</p>
                
                <div className='grid grid-cols-6 gap-4 m-8'>
                    <div>
                        <input type="type" className="form-input border border-gray-400 focus:border-orange-400" />
                    </div>
                    <div>
                        <input type="type" className="form-input border border-gray-400 focus:border-orange-400" />
                    </div>
                    <div>
                        <input type="type" className="form-input border border-gray-400 focus:border-orange-400" />
                    </div>
                    <div>
                        <input type="type" className="form-input border border-gray-400 focus:border-orange-400" />
                    </div>
                    <div>
                        <input type="type" className="form-input border border-gray-400 focus:border-orange-400" />
                    </div>
                    <div>
                        <input type="type" className="form-input border border-gray-400 focus:border-orange-400" />
                    </div>
                </div>
                    
                <div className="flex justify-center py-6">
                    <Link to="">
                        <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]" onClick={() => coloredToast('dark')}>
                            CONFIRM
                        </button>
                        <div id="secondary-toast"></div>
                    </Link>
                </div>                           
            </div>
        </div>
    );
};

export default Confirmation2FA;

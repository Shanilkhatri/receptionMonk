import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';

const QR2FA = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('2FA QR Scan'));
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
                    <div className="text-center pb-3">
                    <h1 className="font-bold text-2xl text-black pb-12">Reception <span className="text-orange-700">Monk</span></h1>
                    <h2 className="font-semibold text-xl mb-3 text-neutral-800">Two Factor Authentication</h2>
                    </div>
                </div>  
                    <p className="m-2 text-base"> Scan the QR code using supported authenticator app</p>
                    <div className="flex justify-center">
                        <div><img className="h-40" src="/assets/images/logo/qr-code.svg" alt="qr code img" /></div>                   
                    </div>
                    <p className="m-2 text-base"> Can't scan the QR code?</p>
                    <p className="m-2 text-base"> Copy this "Recovery codes" to your authenticator app. Enter this code into your authenticator app instead. </p>
                <div className="mt-6 text-base flex justify-between font-semibold">
                    <div><p className='text-black'>Recovery Codes:</p></div>
                    <div>
                        <p className='text-orange-600 flex'>
                            <Link to="/" className=' flex'>
                                Download <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className='mx-1 mt-1 font-bold'>
                                    <path d="M12.5535 16.5061C12.4114 16.6615 12.2106 16.75 12 16.75C11.7894 16.75 11.5886 16.6615 11.4465 16.5061L7.44648 12.1311C7.16698 11.8254 7.18822 11.351 7.49392 11.0715C7.79963 10.792 8.27402 10.8132 8.55352 11.1189L11.25 14.0682V3C11.25 2.58579 11.5858 2.25 12 2.25C12.4142 2.25 12.75 2.58579 12.75 3V14.0682L15.4465 11.1189C15.726 10.8132 16.2004 10.792 16.5061 11.0715C16.8118 11.351 16.833 11.8254 16.5535 12.1311L12.5535 16.5061Z" fill="#c8400d"/>
                                    <path d="M3.75 15C3.75 14.5858 3.41422 14.25 3 14.25C2.58579 14.25 2.25 14.5858 2.25 15V15.0549C2.24998 16.4225 2.24996 17.5248 2.36652 18.3918C2.48754 19.2919 2.74643 20.0497 3.34835 20.6516C3.95027 21.2536 4.70814 21.5125 5.60825 21.6335C6.47522 21.75 7.57754 21.75 8.94513 21.75H15.0549C16.4225 21.75 17.5248 21.75 18.3918 21.6335C19.2919 21.5125 20.0497 21.2536 20.6517 20.6516C21.2536 20.0497 21.5125 19.2919 21.6335 18.3918C21.75 17.5248 21.75 16.4225 21.75 15.0549V15C21.75 14.5858 21.4142 14.25 21 14.25C20.5858 14.25 20.25 14.5858 20.25 15C20.25 16.4354 20.2484 17.4365 20.1469 18.1919C20.0482 18.9257 19.8678 19.3142 19.591 19.591C19.3142 19.8678 18.9257 20.0482 18.1919 20.1469C17.4365 20.2484 16.4354 20.25 15 20.25H9C7.56459 20.25 6.56347 20.2484 5.80812 20.1469C5.07435 20.0482 4.68577 19.8678 4.40901 19.591C4.13225 19.3142 3.9518 18.9257 3.85315 18.1919C3.75159 17.4365 3.75 16.4354 3.75 15Z" fill="#c8400d"/>
                                    </svg>
                            </Link>
                            <span className='px-3 text-gray-500'> | </span> 
                            <Link to="/"  className=' flex'>
                                Copy <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className='mx-1 mt-1 font-bold'>
                                <path fill-rule="evenodd" clip-rule="evenodd" d="M19.5 16.5L19.5 4.5L18.75 3.75H9L8.25 4.5L8.25 7.5L5.25 7.5L4.5 8.25V20.25L5.25 21H15L15.75 20.25V17.25H18.75L19.5 16.5ZM15.75 15.75L15.75 8.25L15 7.5L9.75 7.5V5.25L18 5.25V15.75H15.75ZM6 9L14.25 9L14.25 19.5L6 19.5L6 9Z" fill="#c8400d"/>
                                </svg>
                            </Link>                            
                        </p>
                    </div>
                </div>

                <pre>
                    <div className='border border-gray-500 p-3 my-3 grid grid-cols-3 gap-2 text-center'>                    
                        <div>123456</div>
                        <div>456789</div>
                        <div>789123</div>
                        <div>147258</div>
                        <div>258369</div>
                        <div>369147</div>
                        <div>159236</div>
                        <div>159478</div>
                        <div>784512</div>
                        <div>895623</div>
                        <div>741963</div>
                        <div>852741</div>
                    </div>
                </pre>
              
                <div className="flex justify-center py-6">
                    <Link to="/auth/Confirmation2FA">
                        <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]">
                            ENTER CONFIRMATION CODE
                        </button>
                    </Link>
                </div>                           
            </div>
        </div>
    );
};

export default QR2FA;

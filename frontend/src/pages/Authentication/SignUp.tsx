import { Link, useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { useEffect } from 'react';
import { setPageTitle } from '../../store/themeConfigSlice';

const SignUp = () => {
    const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Sign Up Process'));
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
                    <h1 className="font-bold text-2xl text-black pb-8">Reception <span className="text-orange-700">Monk</span></h1>
                    <h2 className="font-semibold text-xl mb-3 text-neutral-800">Sign Up</h2>
                 </div>
                </div>  
                <form className="space-y-5" onSubmit={submitForm}>
                    <p className="mb-7">Please fill all details to complete Registration</p>
                    
                    <div>
                        <label htmlFor="changepwd">Name <span className='text-red-600'>*</span></label>
                        <input id="phoneno" type="text" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter Name" />
                    </div>

                    <div>
                        <label htmlFor="email">Email <span className='text-red-600'>*</span></label> 
                        <input id="email" type="email" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter Email" />
                    </div>

                    <div>
                        <label htmlFor="password">Password <span className='text-red-600'>*</span></label>
                        <input id="password" type="password" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter Password" />
                    </div>

                    <div>
                        <label htmlFor="changepwd">Date of Birth <span className='text-red-600'>*</span></label>
                        <input id="phoneno" type="text" className="form-input border border-gray-400 focus:border-orange-400" placeholder="Enter DOB" />
                    </div>

                    <div>
                        <label htmlFor="changepwd">Account Type <span className='text-red-600'>*</span></label>
                        <select className="form-select text-white-dark border border-gray-400 focus:border-orange-400">
                            <option value="" selected disabled>Select</option>
                            <option value="owner">Owner</option>
                            <option value="client">Client</option>
                        </select>
                    </div>

                    <div className="flex justify-center py-6">
                    <Link to="/">
                        <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:bg-[#282828]">
                            SIGN UP
                        </button>
                    </Link>
                    </div>        

                </form>                
            </div>
        </div>
    );
};

export default SignUp;

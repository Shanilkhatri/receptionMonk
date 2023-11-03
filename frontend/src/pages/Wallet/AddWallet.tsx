import { Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { IRootState } from '../../store';
import { setPageTitle } from '../../store/themeConfigSlice';
import { useEffect } from 'react';

const AddWallet = () => {

const dispatch = useDispatch();
    useEffect(() => {
        dispatch(setPageTitle('Add Wallet'));
    });

const isRtl = useSelector((state: IRootState) => state.themeConfig.rtlClass) === 'rtl' ? true : false;
    
    return (
        <div>
            <ul className="flex space-x-2 rtl:space-x-reverse">
                <li>
                    <Link to="#" className="text-primary hover:underline">
                        Wallet
                    </Link>
                </li>
                <li className="before:content-['/'] ltr:before:mr-2 rtl:before:ml-2">
                    <span>Add</span>
                </li>
            </ul>

            <div className='flex xl:flex-row flex-col gap-2.5 py-5'>
                <div className="panel px-0 flex-1 py-6 ltr:xl:mr-6 rtl:xl:ml-6">
                    <div className="mt-8 px-4">
                        <div className='text-xl font-bold text-dark dark:text-white text-center'>Add Wallet</div>
                        <div className="flex justify-between lg:flex-row flex-col">
                            <div className="lg:w-1/2 w-full m-8">                           
                                <div className="flex items-center my-8">
                                    <label htmlFor="wallet-charge" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Charge Plan
                                    </label>
                                    <select id="wallet-charge" name="wallet-charge" className="form-select flex-1 border border-gray-400 focus:border-orange-400 text-gray-500">
                                        <option value="">Select Credit</option>
                                        <option value="United States">Plan 1</option>
                                        <option value="United Kingdom">Plan 2</option>                   
                                    </select>
                                </div>                         
                                <div className="my-8 flex items-center">
                                    <label htmlFor="wallet-cost" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Amount
                                    </label>
                                    <input id="wallet-cost" type="text" name="wallet-cost" className="form-input flex-1 border border-gray-400 focus:border-orange-400" placeholder="Enter Amount" />
                                </div>   
                                <div className="flex items-center my-8 ">
                                    <label htmlFor="charge-date" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Date
                                    </label>
                                    <input id="charge-date" type="date" name="charge-date" className="form-input flex-1 border border-gray-400 focus:border-orange-400  text-gray-600" />
                                </div>                                  
                            </div>
                            <div className="lg:w-1/2 w-full m-8">                                              
                                <div className="flex items-center my-8">
                                    <label htmlFor="company-name" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Company Name
                                    </label>
                                    <select id="company-name" name="company-name" className="form-select flex-1 border border-gray-400 focus:border-orange-400  text-gray-500">
                                        <option value="">Select</option>
                                        <option value="United States">HP</option>
                                        <option value="United Kingdom">Dell</option>                   
                                    </select>
                                </div>
                                <div className="my-8 flex items-center">
                                    <label htmlFor="reason" className="ltr:mr-2 rtl:ml-2 w-1/3 mb-0">
                                        Reason
                                    </label>
                                    <textarea className="form-textarea border border-gray-400 focus:border-orange-400 flex-1" placeholder="Enter Description"></textarea>
                                </div>
                            </div>
                        </div>
                        <div className="flex w-full pb-12">
                            <div className="flex items-center lg:justify-end lg:mr-6 w-1/2">                        
                                <button type="reset" className="btn btn-outline-dark rounded-xl px-6 font-bold">RESET</button>
                            </div>
                            <div className="flex items-center lg:ml-6 w-1/2">
                                <button type="submit" className="btn bg-[#c8400d] rounded-xl text-white font-bold shadow-none px-8 hover:border-black font-bold">ADD</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default AddWallet;